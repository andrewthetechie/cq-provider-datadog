
# Table: datadog_incidents
IncidentResponseData Incident data from a response.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|attributes_created|timestamp without time zone|Timestamp when the incident was created.|
|attributes_customer_impact_duration|bigint|Length of the incident's customer impact in seconds. Equals the difference between `customer_impact_start` and `customer_impact_end`.|
|attributes_customer_impacted|boolean|A flag indicating whether the incident caused customer impact.|
|attributes_fields|jsonb|A condensed view of the user-defined fields attached to incidents.|
|attributes_modified|timestamp without time zone|Timestamp when the incident was last modified.|
|attributes_postmortem_id|text|The UUID of the postmortem object attached to the incident.|
|attributes_public_id|bigint|The monotonically increasing integer ID for the incident.|
|attributes_time_to_detect|bigint|The amount of time in seconds to detect the incident. Equals the difference between `customer_impact_start` and `detected`.|
|attributes_time_to_internal_response|bigint|The amount of time in seconds to call incident after detection|
|attributes_time_to_repair|bigint|The amount of time in seconds to resolve customer impact after detecting the issue|
|attributes_time_to_resolve|bigint|The amount of time in seconds to resolve the incident after it was created|
|attributes_title|text|The title of the incident, which summarizes what happened.|
|attributes_additional_properties|jsonb||
|id|text|The incident's ID.|
|relationships_commander_user_additional_properties|jsonb||
|relationships_created_by_user_data_id|text|A unique identifier that represents the user.|
|relationships_created_by_user_data_type|text|Users resource type.|
|relationships_created_by_user_data_additional_properties|jsonb||
|relationships_created_by_user_additional_properties|jsonb||
|relationships_integrations_additional_properties|jsonb||
|relationships_last_modified_by_user_data_id|text|A unique identifier that represents the user.|
|relationships_last_modified_by_user_data_type|text|Users resource type.|
|relationships_last_modified_by_user_data_additional_properties|jsonb||
|relationships_last_modified_by_user_additional_properties|jsonb||
|relationships_postmortem_data_id|text|A unique identifier that represents the postmortem.|
|relationships_postmortem_data_type|text|Incident postmortem resource type.|
|relationships_postmortem_data_additional_properties|jsonb||
|relationships_postmortem_additional_properties|jsonb||
|relationships_additional_properties|jsonb||
|type|text|Incident resource type.|
|additional_properties|jsonb||
