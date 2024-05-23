---
page_title: "cdp_opdb_operational_database Resource - terraform-provider-cdp"
subcategory: "Operational Database"
description: |-
  Creates an Operational DataBase.
---

# cdp_opdb_operational_database (Resource)

Creates an Operational DataBase.

## Example Usage
### Required parameters for Operational Database
Operational Database can be created by providing just the `environment_name` and the desired `database_name`.

Below example uses the most simple configuration.
```terraform
// Copyright 2024 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

terraform {
  required_providers {
    cdp = {
      source = "cloudera/cdp"
    }
  }
}

provider "cdp" {
  cdp_config_file             = "/Users/<value>/.cdp/config"
  cdp_shared_credentials_file = "/Users/<value>/.cdp/credentials"
}

resource "cdp_opdb_operational_database" "opdb-simple-example" {
  environment_name = "<value>"
  database_name    = "<value>"
}
```

### Optional parameters
Operational Database can also be created with custom configuration.

Below example uses every parameter available to configure.

```terraform
// Copyright 2024 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

terraform {
  required_providers {
    cdp = {
      source = "cloudera/cdp"
    }
  }
}

provider "cdp" {
  cdp_config_file             = "/Users/<value>/.cdp/config"
  cdp_shared_credentials_file = "/Users/<value>/.cdp/credentials"
}

resource "cdp_opdb_operational_database" "opdb-detailed-example" {
  environment_name = "<value>"
  database_name    = "<value>"

  scale_type   = "HEAVY"                // valid options are "MICRO","LIGHT","HEAVY"
  storage_type = "CLOUD_WITH_EPHEMERAL" // valid options are "CLOUD_WITH_EPHEMERAL","CLOUD","HDFS"

  java_version = 8

  disable_external_db = true

  disable_multi_az = true
  subnet_id        = "<value>"

  num_edge_nodes = 1

  auto_scaling_parameters = {
    targeted_value_for_metric = 249
    max_workers_for_database  = 50
    max_workers_per_batch     = 4
    min_workers_for_database  = 15
    evaluation_period         = 2400
    minimum_block_cache_gb    = 1
    # beta only
    max_cpu_utilization            = -1
    max_compute_nodes_for_database = -1
    min_compute_nodes_for_database = -1
    max_hdfs_usage_percentage      = 80
    max_regions_per_region_server  = 200
  }

  attached_storage_for_workers = {
    volume_count = 3 // min 1 max 8
    volume_size  = 1024
    volume_type  = "SSD" // valid options are "HDD", "SSD", "LOCAL_SSD"
  }

  image = {
    id      = "<value>"
    catalog = "<value>"
  }

  disable_kerberos = true
  disable_jwt_auth = true

  enable_grafana = true

  custom_user_tags = [
    {
      key   = "key1"
      value = "value1"
    },
    {
      key   = "key2"
      value = "value2"
    },
    {
      key   = "key3"
      value = "value3"
    }
  ]

  enable_region_canary = true

  recipes = [
    {
      names          = ["<value>"],
      instance_group = "<value>"
    },
    {
      names          = ["<value>", "<value>"],
      instance_group = "<value>"
    }
  ]
  storage_location = "s3a://<value>/"

  volume_encryptions = [
    {
      encryption_key = "<value>",
      instance_group = "<value>"
    },
    {
      encryption_key = "<value>",
      instance_group = "<value>"
    }
  ]

}
```

## Update Database
It is possible to update the autoscaling parameters without deleting and creating a new Operational Database.
It is also possible to change Catalog name for the image.

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `database_name` (String) The name of the database.
- `environment_name` (String) The name of the environment where the cluster will belong to.

### Optional

- `attached_storage_for_workers` (Attributes) Attached storage for the worker nodes for AWS, Azure, and GCP cloud providers. (see [below for nested schema](#nestedatt--attached_storage_for_workers))
- `auto_scaling_parameters` (Attributes) (see [below for nested schema](#nestedatt--auto_scaling_parameters))
- `custom_user_tags` (Attributes Set) Optional tags to apply to launched infrastructure resources (see [below for nested schema](#nestedatt--custom_user_tags))
- `disable_external_db` (Boolean) Disable external database creation or not. It is only available in the BETA api.
- `disable_jwt_auth` (Boolean) Disable OAuth Bearer (JWT) authentication scheme.
- `disable_kerberos` (Boolean) Disable Kerberos authentication.
- `disable_multi_az` (Boolean) Disable deployment to multiple availability zones or not
- `enable_grafana` (Boolean) To enable grafana server for the database.
- `enable_region_canary` (Boolean) To enable the region canary for the database.
- `image` (Attributes) Details of an Image. (see [below for nested schema](#nestedatt--image))
- `java_version` (Number) Java version. It is only available in the BETA api.
- `num_edge_nodes` (Number) Number of edge nodes
- `polling_options` (Attributes) Polling related configuration options that could specify various values that will be used during CDP resource creation. (see [below for nested schema](#nestedatt--polling_options))
- `recipes` (Attributes Set) Custom recipes for the database. (see [below for nested schema](#nestedatt--recipes))
- `scale_type` (String) Scale type, MICRO, LIGHT or HEAVY
- `storage_location` (String) Storage Location for OPDB. It is only available in the BETA api.
- `storage_type` (String) Storage type for clusters, CLOUD_WITH_EPHEMERAL, CLOUD or HDFS
- `subnet_id` (String) ID of the subnet to deploy to
- `volume_encryptions` (Attributes Set) Specifies encryption key to encrypt volume for instance group. It is currently supported for AWS cloud provider only. It is only available in the BETA api. (see [below for nested schema](#nestedatt--volume_encryptions))

### Read-Only

- `crn` (String) The CRN of the cluster.
- `status` (String) The last known state of the cluster

<a id="nestedatt--attached_storage_for_workers"></a>
### Nested Schema for `attached_storage_for_workers`

Optional:

- `volume_count` (Number) The number of Volumes. Default is 4. Valid Range: Minimum value of 1, maximum value 8.
- `volume_size` (Number) The target size of the volume, in GiB. Default is 2048.
- `volume_type` (String) Volume Type. HDD - Hard disk drives (HDD) volume type. Default is HDD. SSD - Solid disk drives (SSD) volume type. LOCAL_SSD - Local SSD volume type.


<a id="nestedatt--auto_scaling_parameters"></a>
### Nested Schema for `auto_scaling_parameters`

Optional:

- `evaluation_period` (Number) Period of metrics(in seconds) needs to be considered.
- `max_compute_nodes_for_database` (Number) The maximum number of compute nodes, as per these metrics, that can be scaled up to. It is only available in the BETA api.
- `max_cpu_utilization` (Number) The maximum percentage threshold for the CPU utilization of the worker nodes. The CPU utilization is obtained from the Cloudera Manager metric ‘cpu_percent’ across worker nodes. Set 100 or more to disable the CPU metrics. It is only available in the BETA api.
- `max_hdfs_usage_percentage` (Number) The maximum percentage of HDFS utilization for the database before we trigger the scaling. It is only available in the BETA api.
- `max_regions_per_region_server` (Number) The maximum number of regions per region server. It is only available in the BETA api.
- `max_workers_for_database` (Number) Maximum number of worker nodes as per this metrics can be scaled up to.
- `max_workers_per_batch` (Number) Maximum number of worker nodes as per this metrics can be scaled up to in one batch.
- `min_compute_nodes_for_database` (Number) The minimum number of compute nodes, as per these metrics, that can be scaled down to. It is only available in the BETA api.
- `min_workers_for_database` (Number) Minimum number of worker nodes as per this metrics can be scaled down to.
- `minimum_block_cache_gb` (Number) The amount of block cache, in Gigabytes, which the database should have.
- `targeted_value_for_metric` (Number) The target value of the metric a user expect to maintain for the cluster


<a id="nestedatt--custom_user_tags"></a>
### Nested Schema for `custom_user_tags`

Required:

- `key` (String)
- `value` (String)


<a id="nestedatt--image"></a>
### Nested Schema for `image`

Required:

- `catalog` (String) Catalog name for the image.
- `id` (String) Image ID for the database.


<a id="nestedatt--polling_options"></a>
### Nested Schema for `polling_options`

Optional:

- `polling_timeout` (Number) Timeout value in minutes that specifies for how long should the polling go for resource creation/deletion.


<a id="nestedatt--recipes"></a>
### Nested Schema for `recipes`

Required:

- `instance_group` (String) The name of the designated instance group.
- `names` (Set of String) The set of recipe names that are going to be applied on the given instance group.


<a id="nestedatt--volume_encryptions"></a>
### Nested Schema for `volume_encryptions`

Required:

- `encryption_key` (String) Encryption key to encrypt volume.
- `instance_group` (String) The name of the designated instance group.