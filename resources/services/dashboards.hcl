service = "datadog"
output_directory = "."
add_generate = true

resource "datadog" "" "dashboards"{
  path = "github.com/DataDog/datadog-api-client-go/api/v1/datadog.Dashboard"
  limit_depth = 1
  options {
    primary_keys = ["id"]
  }

}