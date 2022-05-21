service = "datadog"
output_directory = "."
add_generate = true

resource "datadog" "" "notebooks"{
  path = "github.com/DataDog/datadog-api-client-go/api/v1/datadog.NotebooksResponseData"
  limit_depth = 1
  options {
    primary_keys = ["id"]
  }

  column "additional_properties" {
    type ="json"
  }

  column "attributes_time_unparsed_object" {
    skip = true
  }

}