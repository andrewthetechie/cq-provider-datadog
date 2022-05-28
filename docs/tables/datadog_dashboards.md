
# Table: datadog_dashboards
Dashboard A dashboard is Datadog’s tool for visually tracking, analyzing, and displaying key performance metrics, which enable you to monitor the health of your infrastructure.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|author_handle|text|Identifier of the dashboard author.|
|created_at|timestamp without time zone|Creation date of the dashboard.|
|id|text|ID of the dashboard.|
|is_read_only|boolean|Whether this dashboard is read-only|
|layout_type|text|Layout type of the dashboard.|
|modified_at|timestamp without time zone|Modification date of the dashboard.|
|notify_list|text[]|List of handles of users to notify when changes are made to this dashboard.|
|reflow_type|text|Reflow type for a **new dashboard layout** dashboard|
|restricted_roles|text[]|A list of role identifiers|
|title|text|Title of the dashboard.|
|url|text|The URL of the dashboard.|
|additional_properties|jsonb||
