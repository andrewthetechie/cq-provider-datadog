service = "datadog"
output_directory = "."
add_generate = true

resource "datadog" "" "permissions"{
  path = "github.com/DataDog/datadog-api-client-go/api/v2/datadog.Role"
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