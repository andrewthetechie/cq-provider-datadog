
# Table: datadog_synthetics
SyntheticsTestDetails Object containing details about your Synthetic test.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|config|jsonb|Configuration object for a Synthetic test.|
|creator_email|text|Email of the creator.|
|creator_handle|text|Handle of the creator.|
|creator_additional_properties|jsonb||
|locations|text[]|Array of locations used to run the test.|
|message|text|Notification message associated with the test.|
|monitor_id|bigint|The associated monitor ID.|
|name|text|Name of the test.|
|options|jsonb|Object describing the extra options for a Synthetic test.|
|public_id|text|The test public ID.|
|status|text|Define whether you want to start (`live`) or pause (`paused`) a Synthetic test.|
|subtype|text|The subtype of the Synthetic API test, `http`, `ssl`, `tcp`, `dns`, `icmp`, `udp`, `websocket` or `multi`.|
|tags|text[]|Array of tags attached to the test.|
|type|text|Type of the Synthetic test, either `api` or `browser`.|
|additional_properties|jsonb||
