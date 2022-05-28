service = "datadog"
output_directory = "."
add_generate = true

resource "datadog" "" "synthetics"{
  path = "github.com/DataDog/datadog-api-client-go/api/v1/datadog.SyntheticsTestDetails"
  limit_depth = 1
  options {
    primary_keys = ["id"]
  }

  column "config" {
    type ="json"
  }

  column "options" {
    type ="json"
  }

}