service = "datadog"
output_directory = "."
add_generate = true

resource "datadog" "" "incidents"{
  path = "github.com/DataDog/datadog-api-client-go/api/v2/datadog.IncidentResponseData"
  limit_depth = 1
  options {
    primary_keys = ["id"]
  }
}