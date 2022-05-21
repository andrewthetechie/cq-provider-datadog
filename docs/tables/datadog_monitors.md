
# Table: datadog_monitors
Datadog Monitors
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|meta_id|text|The Meta_ID of this monitor, combines datadog account name from config with ID and monitor name and hashes them to provide a unique PK.|
|id|integer|The ID of this monitor.|
|name|text|The name of this monitor.|
|type|text|The type of this monitor.|
|created|timestamp without time zone|Timestamp of when the monitor was created|
|modified|timestamp without time zone|Timestamp of when the monitor was last modified|
|creator|jsonb|Info about the creator of this monitor including name, handle, and email|
|message|text|Monitor's message|
|multi|boolean|If a monitor is a multi alert|
|options|jsonb|Options for this monitor as a json blob|
|overall_state|text|Overall State of the monitor, string summary|
|state|jsonb|State of the monitor as json|
|priority|integer|Priority of the monitor|
|query|text|Query of the monitor|
|restricted_roles|text[]|Restricted roles of the monitor|
|tags|text[]|Tags of the monitor|
|additional_properties|jsonb|Additional monitor properties|
