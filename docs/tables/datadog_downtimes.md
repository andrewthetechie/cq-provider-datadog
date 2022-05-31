
# Table: datadog_downtimes
Downtime Downtiming gives you greater control over monitor notifications by allowing you to globally exclude scopes from alerting. Downtime settings, which can be scheduled with start and end times, prevent all alerting related to specified Datadog tags.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_name|text|The name of this datadog account from your config.|
|active|boolean|If a scheduled downtime currently exists.|
|active_child|jsonb||
|creator_id|integer|User ID of the downtime creator.|
|disabled|boolean|If a downtime has been disabled.|
|downtime_type|integer|`0` for a downtime applied on `*` or all, `1` when the downtime is only scoped to hosts, or `2` when the downtime is scoped to anything but hosts.|
|id|bigint|The downtime ID.|
|message|text|A message to include with notifications for this downtime. Email notifications can be sent to specific users by using the same `@username` notation as events.|
|monitor_tags|text[]|A comma-separated list of monitor tags|
|mute_first_recovery_notification|boolean|If the first recovery notification during a downtime should be muted.|
|recurrence|jsonb||
|scope|text[]|The scope(s) to which the downtime applies|
|start|bigint|POSIX timestamp to start the downtime. If not provided, the downtime starts the moment it is created.|
|timezone|text|The timezone in which to display the downtime's start and end times in Datadog applications.|
|additional_properties|jsonb||
