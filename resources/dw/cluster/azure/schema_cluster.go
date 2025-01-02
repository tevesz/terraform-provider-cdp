// Copyright 2025 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package azure

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var dwClusterSchema = schema.Schema{
	MarkdownDescription: "Creates an Azure Data Warehouse cluster.",
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"crn": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "The cloudera resource name of the environment that the cluster will read from.",
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The name of the cluster matches the environment name.",
		},
		"cluster_id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The id of the cluster.",
		},
		"last_updated": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Timestamp of the last Terraform update of the order.",
		},
		"status": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The status of the cluster.",
		},
		"user_assigned_managed_identity": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "Resource ID of the managed identity used by AKS.",
		},
		"log_analytics_workspace_id": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: "Enable monitoring of Azure Kubernetes Service (AKS) cluster. Workspace ID for Azure log analytics.",
		},
		"database_backup_retention_days": schema.Int64Attribute{
			Optional:            true,
			MarkdownDescription: "The number of days to retain Postgres database backups.",
		},
		"custom_registry_options": schema.SingleNestedAttribute{
			Optional: true,
			Attributes: map[string]schema.Attribute{
				"registry_type": schema.StringAttribute{
					Required:            true,
					MarkdownDescription: "Registry type, supported values are ECR or ACR.",
				},
				"repository_url": schema.StringAttribute{
					Required:            true,
					MarkdownDescription: "The URL of the registry.",
				},
			},
		},
		"custom_subdomain": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: "The custom subdomain to keep compatibility with old URL format.",
		},
		"network_settings": schema.SingleNestedAttribute{
			Required:   true,
			Attributes: networkSettings,
		},
		"instance_settings": schema.SingleNestedAttribute{
			Optional:   true,
			Attributes: instanceSettings,
		},
		"polling_options": schema.SingleNestedAttribute{
			MarkdownDescription: "Polling related configuration options that could specify various values that will be used during CDP resource creation.",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"async": schema.BoolAttribute{
					MarkdownDescription: "Boolean value that specifies if Terraform should wait for resource creation/deletion.",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(false),
					PlanModifiers: []planmodifier.Bool{
						boolplanmodifier.UseStateForUnknown(),
					},
				},
				"polling_timeout": schema.Int64Attribute{
					MarkdownDescription: "Timeout value in minutes that specifies for how long should the polling go for resource creation/deletion.",
					Default:             int64default.StaticInt64(40),
					Computed:            true,
					Optional:            true,
				},
				"call_failure_threshold": schema.Int64Attribute{
					MarkdownDescription: "Threshold value that specifies how many times should a single call failure happen before giving up the polling.",
					Default:             int64default.StaticInt64(3),
					Computed:            true,
					Optional:            true,
				},
			},
		},
	},
}

var networkSettings = map[string]schema.Attribute{
	"outbound_type": schema.StringAttribute{
		Required:            true, //TODO add validation, allowed values: LoadBalancer UserAssignedNATGateway UserDefinedRouting
		MarkdownDescription: "Network outbound type. This setting controls the egress traffic for cluster nodes in Azure Kubernetes Service. Please refer to the following AKS documentation on the Azure portal. https://learn.microsoft.com/en-us/azure/aks/egress-outboundtype, https://learn.microsoft.com/en-us/azure/aks/nat-gateway",
	},
	"enable_private_sql": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Enable Azure Private AKS mode.",
	},
	"private_dns_zone_aks": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Private DNS zone resource ID for AKS.",
	},
}
