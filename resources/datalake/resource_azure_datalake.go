// Copyright 2023 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package datalake

import (
	"context"
	"time"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/client/operations"
	datalakemodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource = &azureDatalakeResource{}
)

type azureDatalakeResource struct {
	client *cdp.Client
}

func NewAzureDatalakeResource() resource.Resource {
	return &azureDatalakeResource{}
}

func (r *azureDatalakeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_datalake_azure_datalake"
}

func (r *azureDatalakeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = azureDatalakeResourceSchema
}

func (r *azureDatalakeResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.client = utils.GetCdpClientForResource(req, resp)
}

func toAzureDatalakeRequest(ctx context.Context, model *azureDatalakeResourceModel) *datalakemodels.CreateAzureDatalakeRequest {
	req := &datalakemodels.CreateAzureDatalakeRequest{}
	req.CloudProviderConfiguration = &datalakemodels.AzureConfigurationRequest{
		ManagedIdentity: model.ManagedIdentity.ValueStringPointer(),
		StorageLocation: model.StorageLocation.ValueStringPointer(),
	}
	req.DatalakeName = model.DatalakeName.ValueStringPointer()
	req.EnableRangerRaz = model.EnableRangerRaz.ValueBool()
	req.EnvironmentName = model.EnvironmentName.ValueStringPointer()
	if model.Image != nil {
		req.Image = &datalakemodels.ImageRequest{
			CatalogName: model.Image.CatalogName.ValueStringPointer(),
			ID:          model.Image.ID.ValueStringPointer(),
		}
	}
	req.JavaVersion = int32(model.JavaVersion.ValueInt64())
	req.Recipes = make([]*datalakemodels.InstanceGroupRecipeRequest, len(model.Recipes))
	for i, v := range model.Recipes {
		req.Recipes[i] = &datalakemodels.InstanceGroupRecipeRequest{
			InstanceGroupName: v.InstanceGroupName.ValueStringPointer(),
			RecipeNames:       utils.FromSetValueToStringList(v.RecipeNames),
		}
	}
	req.Runtime = model.Runtime.ValueString()
	req.Scale = datalakemodels.DatalakeScaleType(model.Scale.ValueString())
	req.Tags = make([]*datalakemodels.DatalakeResourceTagRequest, len(model.Tags.Elements()))
	i := 0
	for k, v := range model.Tags.Elements() {
		val, diag := v.(basetypes.StringValuable).ToStringValue(ctx)
		if !diag.HasError() {
			req.Tags[i] = &datalakemodels.DatalakeResourceTagRequest{
				Key:   &k,
				Value: val.ValueStringPointer(),
			}
		}
	}
	return req
}

func (r *azureDatalakeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var state azureDatalakeResourceModel
	diags := req.Plan.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Got error while trying to set plan")
		return
	}

	client := r.client.Datalake

	params := operations.NewCreateAzureDatalakeParamsWithContext(ctx)
	params.WithInput(toAzureDatalakeRequest(ctx, &state))
	responseOk, err := client.Operations.CreateAzureDatalake(params)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Azure Datalake",
			"Got the following error creating Azure Datalake: "+err.Error(),
		)
		return
	}

	datalakeResp := responseOk.Payload
	toAzureDatalakeResourceModel(datalakeResp, &state)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := waitForDatalakeToBeRunning(ctx, state.DatalakeName.ValueString(), time.Hour, r.client.Datalake); err != nil {
		resp.Diagnostics.AddError(
			"Error creating Azure Data Lake",
			"Failed to poll creating Azure Data Lake: "+err.Error(),
		)
		return
	}

	descParams := operations.NewDescribeDatalakeParamsWithContext(ctx)
	descParams.WithInput(&datalakemodels.DescribeDatalakeRequest{DatalakeName: state.DatalakeName.ValueStringPointer()})
	descResponseOk, err := client.Operations.DescribeDatalake(descParams)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting Azure Datalake",
			"Got the following error getting Azure Datalake: "+err.Error(),
		)
		return
	}

	descDlResp := descResponseOk.Payload
	datalakeDetailsToAzureDatalakeResourceModel(ctx, descDlResp.Datalake, &state)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func toAzureDatalakeResourceModel(resp *datalakemodels.CreateAzureDatalakeResponse, model *azureDatalakeResourceModel) {
	model.ID = types.StringPointerValue(resp.Datalake.DatalakeName)
	model.CertificateExpirationState = types.StringValue(resp.Datalake.CertificateExpirationState)
	model.CreationDate = types.StringValue(resp.Datalake.CreationDate.String())
	model.Crn = types.StringPointerValue(resp.Datalake.Crn)
	model.DatalakeName = types.StringPointerValue(resp.Datalake.DatalakeName)
	model.EnableRangerRaz = types.BoolValue(resp.Datalake.EnableRangerRaz)
	model.EnvironmentCrn = types.StringValue(resp.Datalake.EnvironmentCrn)
	model.Status = types.StringValue(resp.Datalake.Status)
	model.StatusReason = types.StringValue(resp.Datalake.StatusReason)
}

func (r *azureDatalakeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state azureDatalakeResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	client := r.client.Datalake

	params := operations.NewDescribeDatalakeParamsWithContext(ctx)
	params.WithInput(&datalakemodels.DescribeDatalakeRequest{DatalakeName: state.DatalakeName.ValueStringPointer()})
	responseOk, err := client.Operations.DescribeDatalake(params)
	if err != nil {
		if dlErr, ok := err.(*operations.DescribeDatalakeDefault); ok {
			if cdp.IsDatalakeError(dlErr.GetPayload(), "NOT_FOUND", "") {
				resp.Diagnostics.AddWarning("Resource not found on provider", "Data lake not found, removing from state.")
				tflog.Warn(ctx, "Data lake not found, removing from state", map[string]interface{}{
					"id": state.ID.ValueString(),
				})
				resp.State.RemoveResource(ctx)
				return
			}
		}
		resp.Diagnostics.AddError(
			"Error getting Azure Datalake",
			"Got the following error getting Azure Datalake: "+err.Error(),
		)
		return
	}

	datalakeResp := responseOk.Payload
	datalakeDetailsToAzureDatalakeResourceModel(ctx, datalakeResp.Datalake, &state)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func datalakeDetailsToAzureDatalakeResourceModel(ctx context.Context, resp *datalakemodels.DatalakeDetails, model *azureDatalakeResourceModel) {
	model.ID = types.StringPointerValue(resp.Crn)
	if resp.AzureConfiguration != nil {
		model.ManagedIdentity = types.StringValue(resp.AzureConfiguration.ManagedIdentity)
	}
	model.CloudStorageBaseLocation = types.StringValue(resp.CloudStorageBaseLocation)
	if resp.ClouderaManager != nil {
		model.ClouderaManager, _ = types.ObjectValueFrom(ctx, map[string]attr.Type{
			"cloudera_manager_repository_url": types.StringType,
			"cloudera_manager_server_url":     types.StringType,
			"version":                         types.StringType,
		}, &clouderaManagerDetails{
			ClouderaManagerRepositoryURL: types.StringPointerValue(resp.ClouderaManager.ClouderaManagerRepositoryURL),
			ClouderaManagerServerURL:     types.StringValue(resp.ClouderaManager.ClouderaManagerServerURL),
			Version:                      types.StringPointerValue(resp.ClouderaManager.Version),
		})
	}
	model.CreationDate = types.StringValue(resp.CreationDate.String())
	model.CredentialCrn = types.StringValue(resp.CredentialCrn)
	model.Crn = types.StringPointerValue(resp.Crn)
	model.DatalakeName = types.StringPointerValue(resp.DatalakeName)
	model.EnableRangerRaz = types.BoolValue(resp.EnableRangerRaz)
	endpoints := make([]*endpoint, len(resp.Endpoints.Endpoints))
	for i, v := range resp.Endpoints.Endpoints {
		endpoints[i] = &endpoint{
			DisplayName: types.StringPointerValue(v.DisplayName),
			KnoxService: types.StringPointerValue(v.KnoxService),
			Mode:        types.StringPointerValue(v.Mode),
			Open:        types.BoolPointerValue(v.Open),
			ServiceName: types.StringPointerValue(v.ServiceName),
			ServiceURL:  types.StringPointerValue(v.ServiceURL),
		}
	}
	model.Endpoints, _ = types.SetValueFrom(ctx, types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"display_name": types.StringType,
			"knox_service": types.StringType,
			"mode":         types.StringType,
			"open":         types.BoolType,
			"service_name": types.StringType,
			"service_url":  types.StringType,
		},
	}, endpoints)
	model.EnvironmentCrn = types.StringValue(resp.EnvironmentCrn)
	instanceGroups := make([]*instanceGroup, len(resp.InstanceGroups))
	for i, v := range resp.InstanceGroups {
		instanceGroups[i] = &instanceGroup{
			Name: types.StringPointerValue(v.Name),
		}

		instances := make([]*instance, len(v.Instances))
		for j, ins := range v.Instances {
			instances[j] = &instance{
				DiscoveryFQDN:   types.StringValue(ins.DiscoveryFQDN),
				ID:              types.StringPointerValue(ins.ID),
				InstanceGroup:   types.StringValue(ins.InstanceGroup),
				InstanceStatus:  types.StringValue(string(ins.InstanceStatus)),
				InstanceTypeVal: types.StringValue(string(ins.InstanceTypeVal)),
				PrivateIP:       types.StringValue(ins.PrivateIP),
				PublicIP:        types.StringValue(ins.PublicIP),
				SSHPort:         types.Int64Value(int64(ins.SSHPort)),
				State:           types.StringPointerValue(ins.State),
				StatusReason:    types.StringValue(ins.StatusReason),
			}
		}
		instanceGroups[i].Instances, _ = types.SetValueFrom(ctx, types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"ambari_server":     types.BoolType,
				"discovery_fqdn":    types.StringType,
				"id":                types.StringType,
				"instance_group":    types.StringType,
				"instance_status":   types.StringType,
				"instance_type_val": types.StringType,
				"life_cycle":        types.StringType,
				"mounted_volumes": types.SetType{
					ElemType: types.ObjectType{
						AttrTypes: map[string]attr.Type{
							"device":      types.StringType,
							"volume_id":   types.StringType,
							"volume_size": types.StringType,
							"volume_type": types.StringType,
						},
					},
				},
				"private_ip":    types.StringType,
				"public_ip":     types.StringType,
				"ssh_port":      types.Int64Type,
				"state":         types.StringType,
				"status_reason": types.StringType,
			},
		}, instances)
	}
	model.InstanceGroups, _ = types.SetValueFrom(ctx, types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"instances": types.SetType{
				ElemType: types.ObjectType{
					AttrTypes: map[string]attr.Type{
						"ambari_server":     types.BoolType,
						"discovery_fqdn":    types.StringType,
						"id":                types.StringType,
						"instance_group":    types.StringType,
						"instance_status":   types.StringType,
						"instance_type_val": types.StringType,
						"life_cycle":        types.StringType,
						"mounted_volumes": types.SetType{
							ElemType: types.ObjectType{
								AttrTypes: map[string]attr.Type{
									"device":      types.StringType,
									"volume_id":   types.StringType,
									"volume_size": types.StringType,
									"volume_type": types.StringType,
								},
							},
						},
						"private_ip":    types.StringType,
						"public_ip":     types.StringType,
						"ssh_port":      types.Int64Type,
						"state":         types.StringType,
						"status_reason": types.StringType,
					},
				},
			},
			"name": types.StringType,
		},
	}, instanceGroups)
	productVersions := make([]*productVersion, len(resp.ProductVersions))
	for i, v := range resp.ProductVersions {
		productVersions[i] = &productVersion{
			Name:    types.StringPointerValue(v.Name),
			Version: types.StringPointerValue(v.Version),
		}
	}
	model.ProductVersions, _ = types.SetValueFrom(ctx, types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":    types.StringType,
			"version": types.StringType,
		},
	}, productVersions)
	model.Region = types.StringValue(resp.Region)
	model.Scale = types.StringValue(string(resp.Shape))
	model.Status = types.StringValue(resp.Status)
	model.StatusReason = types.StringValue(resp.StatusReason)
	if model.CertificateExpirationState.IsUnknown() {
		model.CertificateExpirationState = types.StringNull()
	}
	if model.Tags.IsUnknown() {
		model.Tags = types.MapNull(types.StringType)
	}
}

func (r *azureDatalakeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

func (r *azureDatalakeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state azureDatalakeResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	client := r.client.Datalake
	params := operations.NewDeleteDatalakeParamsWithContext(ctx)
	params.WithInput(&datalakemodels.DeleteDatalakeRequest{
		DatalakeName: state.DatalakeName.ValueStringPointer(),
		Force:        false,
	})
	_, err := client.Operations.DeleteDatalake(params)
	if err != nil {
		if dlErr, ok := err.(*operations.DescribeDatalakeDefault); ok {
			if cdp.IsDatalakeError(dlErr.GetPayload(), "NOT_FOUND", "") {
				tflog.Info(ctx, "Data lake already deleted", map[string]interface{}{
					"id": state.ID.ValueString(),
				})
				return
			}
		}
		resp.Diagnostics.AddError(
			"Error Deleting Azure Datalake",
			"Could not delete Azure Datalake unexpected error: "+err.Error(),
		)
		return
	}

	if err := waitForDatalakeToBeDeleted(ctx, state.DatalakeName.ValueString(), time.Hour, r.client.Datalake); err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Azure Data Lake",
			"Failed to poll delete Azure Data Lake, unexpected error: "+err.Error(),
		)
		return
	}
}