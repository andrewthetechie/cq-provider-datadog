
# Table: datadog_monitors
Monitor Object describing a monitor.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|created|timestamp without time zone|Timestamp of the monitor creation.|
|creator_email|text|Email of the creator.|
|creator_handle|text|Handle of the creator.|
|creator_additional_properties|jsonb||
|id|bigint|ID of this monitor.|
|message|text|A message to include with notifications for this monitor.|
|modified|timestamp without time zone|Last timestamp when the monitor was edited.|
|multi|boolean|Whether or not the monitor is broken down on different groups.|
|name|text|The monitor name.|
|options_aggregation_group_by|text|Group to break down the monitor on.|
|options_aggregation_metric|text|Metric name used in the monitor.|
|options_aggregation_type|text|Metric type used in the monitor.|
|options_aggregation_additional_properties|jsonb||
|options_device_ids|text[]|IDs of the device the Synthetics monitor is running on. Deprecated|
|options_enable_logs_sample|boolean|Whether or not to send a log sample when the log monitor triggers.|
|options_escalation_message|text|We recommend using the [is_renotify](https://docs.datadoghq.com/monitors/notify/?tab=is_alert#renotify), block in the original message instead. A message to include with a re-notification|
|options_groupby_simple_monitor|boolean|Whether the log alert monitor triggers a single alert or multiple alerts when any group breaches a threshold.|
|options_include_tags|boolean|A Boolean indicating whether notifications from this monitor automatically inserts its triggering tags into the title.  **Examples** - If `True`, `[Triggered on {host:h1}] Monitor Title` - If `False`, `[Triggered] Monitor Title`|
|options_locked|boolean|Whether or not the monitor is locked (only editable by creator and admins)|
|options_notify_audit|boolean|A Boolean indicating whether tagged users is notified on changes to this monitor.|
|options_notify_no_data|boolean|A Boolean indicating whether this monitor notifies when data stops reporting.|
|options_renotify_statuses|text[]|The types of monitor statuses for which re-notification messages are sent.|
|options_require_full_window|boolean|A Boolean indicating whether this monitor needs a full window of data before it’s evaluated. We highly recommend you set this to `false` for sparse metrics, otherwise some evaluations are skipped|
|options_silenced|jsonb|Information about the downtime applied to the monitor. Deprecated|
|options_threshold_windows_additional_properties|jsonb||
|options_thresholds_critical|float|The monitor `CRITICAL` threshold.|
|options_thresholds_additional_properties|jsonb||
|options_additional_properties|jsonb||
|overall_state|text|The different states your monitor can be in.|
|query|text|The monitor query.|
|restricted_roles|text[]|A list of unique role identifiers to define which roles are allowed to edit the monitor|
|state_groups|jsonb|Dictionary where the keys are groups (comma separated lists of tags) and the values are the list of groups your monitor is broken down on.|
|state_additional_properties|jsonb||
|tags|text[]|Tags associated to your monitor.|
|type|text|The type of the monitor|
|additional_properties|jsonb||
