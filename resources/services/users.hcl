service = "datadog"
output_directory = "."
add_generate = true

resource "datadog" "" "users"{
  path = "github.com/DataDog/datadog-api-client-go/api/v2/datadog.User"
  limit_depth = 1
  options {
    primary_keys = ["id"]
  }

}