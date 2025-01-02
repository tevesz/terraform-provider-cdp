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
	"github.com/cloudera/terraform-provider-cdp/utils"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type databaseCatalog struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	LastUpdated types.String `tfsdk:"last_updated"`
	Status      types.String `tfsdk:"status"`
}

type customRegistryOptions struct {
	RegistryType  types.String `tfsdk:"registry_type"`
	RepositoryURL types.String `tfsdk:"repository_url"`
}

type instanceResourceModel struct {
	EnableAZ             types.Bool `tfsdk:"enable_az"`
	EnableSpotInstances  types.Bool `tfsdk:"enable_spot_instances"`
	ComputeInstanceTypes types.List `tfsdk:"compute_instance_types"`
}

type networkResourceModel struct {
	OutboundType                     types.String `tfsdk:"outbound_type"`
	EnablePrivateSQL                 types.Bool   `tfsdk:"enable_private_sql"`
	PrivateDNSZoneAKS                types.String `tfsdk:"private_dns_zone_aks"`
	EnablePrivateAKS                 types.Bool   `tfsdk:"enable_private_aks"`
	UseOverlayNetwork                types.Bool   `tfsdk:"use_overlay_network"`
	WhitelistK8sClusterAccessIpCidrs types.List   `tfsdk:"whitelist_k8s_cluster_access_ip_cidrs"`
	WhitelistWorkloadAccessIpCidrs   types.List   `tfsdk:"whitelist_workload_access_ip_cidrs"`
	UsePrivateLoadBalancer           types.Bool   `tfsdk:"use_private_load_balancer"`
	UsePublicWorkerNode              types.Bool   `tfsdk:"use_public_worker_node"`
}

type resourceModel struct {
	ID                          types.String           `tfsdk:"id"`
	Crn                         types.String           `tfsdk:"crn"`
	Name                        types.String           `tfsdk:"name"`
	ClusterID                   types.String           `tfsdk:"cluster_id"`
	LastUpdated                 types.String           `tfsdk:"last_updated"`
	Status                      types.String           `tfsdk:"status"`
	UserAssignedManagedIdentity types.String           `tfsdk:"user_assigned_managed_identity"`
	LogAnalyticsWorkspaceID     types.String           `tfsdk:"log_analytics_workspace_id"`
	DatabaseBackupRetentionDays types.Int64            `tfsdk:"database_backup_retention_days"`
	CustomRegistryOptions       *customRegistryOptions `tfsdk:"custom_registry_options"`
	CustomSubdomain             types.String           `tfsdk:"custom_subdomain"`
	NetworkSettings             *networkResourceModel  `tfsdk:"network_settings"`
	InstanceSettings            *instanceResourceModel `tfsdk:"instance_settings"`
	DatabaseCatalog             *databaseCatalog       `tfsdk:"database_catalog"`
	PollingOptions              *utils.PollingOptions  `tfsdk:"polling_options"`
}

func (p *resourceModel) GetPollingOptions() *utils.PollingOptions {
	return p.PollingOptions
}
