
# Table: datadog_hosts
Datadog hosts
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|meta_id|text|The Meta_ID of this host, combines datadog account name from config with ID hashes them to provide a unique PK.|
|id|bigint|The ID of this host.|
|aws_name|text|The aws name of this host.|
|host_name|text|The name of this host.|
|aliases|jsonb|host's aliases|
|apps|jsonb|host's apps|
|sources|jsonb|Sources for this host|
|tags_by_source|jsonb|Tags for this host|
|last_reported_time|bigint|Timestamp of when the host has last reported as a unix timestamp|
|is_muted|boolean|Whether this host is muted|
|mute_timeout|integer|Whether this host is muted|
|up|boolean|Whether this host is up|
|meta|jsonb|Meta info about this host|
|metrics|jsonb|Meta info about this host|
