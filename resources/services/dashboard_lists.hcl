service = "datadog"
output_directory = "."
add_generate = true

resource "datadog" "" "dashboard_lists"{
  path = "github.com/DataDog/datadog-api-client-go/api/v1/datadog.DashboardList"
  limit_depth = 1
  options {
    primary_keys = ["meta_id"]
  }

  userDefinedColumn "meta_id" {
    type        = "string"
    description = "Meta ID for this dashboard"
    generate_resolver = true
  }

}