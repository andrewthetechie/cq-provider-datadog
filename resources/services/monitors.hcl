service = "datadog"
output_directory = "."
add_generate = true

resource "datadog" "" "monitors"{
  path = "github.com/DataDog/datadog-api-client-go/api/v1/datadog.Monitor"
  limit_depth = 1
  options {
    primary_keys = ["id"]
  }

}