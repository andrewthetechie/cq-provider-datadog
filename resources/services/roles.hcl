service = "datadog"
output_directory = "."
add_generate = true

resource "datadog" "" "roles"{
  path = "github.com/DataDog/datadog-api-client-go/api/v2/datadog.Role"
  limit_depth = 1
  options {
    primary_keys = ["id"]
  }

}