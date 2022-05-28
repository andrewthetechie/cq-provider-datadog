service = "datadog"
output_directory = "."
add_generate = true

resource "datadog" "" "hosts"{
  path = "github.com/DataDog/datadog-api-client-go/api/v1/datadog.Host"
  limit_depth = 1
  options {
    primary_keys = ["id"]
  }

  column "meta_agent_checks" {
    skip = true
  }

}