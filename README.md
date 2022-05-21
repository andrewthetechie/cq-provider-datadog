<p align="center">
<a href="https://cloudquery.io">
<img alt="cloudquery logo" width=75% src="https://github.com/cloudquery/cloudquery/raw/main/docs/images/logo.png" />
</a>
</p>

CloudQuery Datadog Provider ![BuildStatus](https://img.shields.io/github/workflow/status/andrewthetechie/cq-provider-datadog/test?style=flat-square) ![License](https://img.shields.io/github/license/andrewthetechie/cq-provider-datadog?style=flat-square)
==================================

This [CloudQuery](https://github.com/cloudquery/cloudquery)
provider transforms Datadog resources to a relational database.

## What is CloudQuery

The [open-source](https://github.com/cloudquery/cloudquery) cloud asset inventory powered by SQL.

CloudQuery extracts, transforms, and loads your cloud assets into normalized PostgreSQL tables. CloudQuery enables you to assess, audit, and evaluate the configurations of your cloud assets.

### Links
* Homepage: https://cloudquery.io
* Documentation: https://docs.cloudquery.io
* CloudQuery Hub (providers & policies documentation): https://hub.cloudquery.io/
* Discord: https://cloudquery.io/discord

# Example Config
```
cloudquery {
  plugin_directory = "./cq/providers"
  policy_directory = "./cq/policies"

  provider "datadog" {
    version = "latest"
  }

  connection {
  }
}

provider "datadog" {
  configuration {
    // currently only supports a single account, future plans for multi-account support
    accounts "main" {
      api_key = "DATADOG API KEY"
      app_key = "DATADOG APP KEY"
    }

  }
  // list of resources to fetch
  resources = ["*"]
}
```

# Resources this Provider Covers

* [dashboard_lists](https://docs.datadoghq.com/api/latest/dashboard-lists/#get-all-dashboard-lists)
* [dashboards](https://docs.datadoghq.com/api/latest/dashboards/#get-all-dashboards)
* [downtimes](https://docs.datadoghq.com/api/latest/downtimes/#get-all-downtimes)
* [hosts](https://docs.datadoghq.com/api/latest/hosts/#get-all-hosts-for-your-organization)
* [incidents](https://docs.datadoghq.com/api/latest/incidents/)
* [monitors](https://docs.datadoghq.com/api/latest/monitors/#get-all-monitor-details)
* [notebooks](https://docs.datadoghq.com/api/latest/notebooks/#get-all-notebooks)
* [permissions](https://docs.datadoghq.com/api/latest/roles/#list-permissions)
* [roles](https://docs.datadoghq.com/api/latest/roles/#list-roles)
* [synthetics](https://docs.datadoghq.com/api/latest/synthetics/#get-the-list-of-all-tests)
* [users](https://docs.datadoghq.com/api/latest/users/)

# Current Limitations

This is a very early provider and has some rough edges. Table layout should be expected to change. 

Most resources are just cq-gen's best approximation of a good structure for the object and have not been optimized.

Account multiplexing is not currently available. If you add more than one account to the datadog provider, only the first will be fetched.

