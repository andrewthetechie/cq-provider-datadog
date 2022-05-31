
# Table: datadog_notebooks
NotebooksResponseData The data for a notebook in get all response.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_name|text|The name of this datadog account from your config.|
|attributes_author_created_at|timestamp without time zone|Creation time of the user.|
|attributes_author_disabled|boolean|Whether the user is disabled.|
|attributes_author_email|text|Email of the user.|
|attributes_author_handle|text|Handle of the user.|
|attributes_author_icon|text|URL of the user's icon.|
|attributes_author_status|text|Status of the user.|
|attributes_author_verified|boolean|Whether the user is verified.|
|attributes_author_additional_properties|jsonb||
|attributes_created|timestamp without time zone|UTC time stamp for when the notebook was created.|
|attributes_metadata_is_template|boolean|Whether or not the notebook is a template.|
|attributes_metadata_take_snapshots|boolean|Whether or not the notebook takes snapshot image backups of the notebook's fixed-time graphs.|
|attributes_metadata_additional_properties|jsonb||
|attributes_modified|timestamp without time zone|UTC time stamp for when the notebook was last modified.|
|attributes_name|text|The name of the notebook.|
|attributes_status|text|Publication status of the notebook|
|attributes_time_notebook_relative_time_live_span|text|The available timeframes depend on the widget you are using.|
|attributes_time_notebook_relative_time_additional_properties|jsonb||
|attributes_time_notebook_absolute_time_end|timestamp without time zone|The end time.|
|attributes_time_notebook_absolute_time_live|boolean|Indicates whether the timeframe should be shifted to end at the current time.|
|attributes_time_notebook_absolute_time_start|timestamp without time zone|The start time.|
|attributes_time_notebook_absolute_time_additional_properties|jsonb||
|attributes_additional_properties|jsonb||
|id|bigint|Unique notebook ID, assigned when you create the notebook.|
|type|text|Type of the Notebook resource.|
|additional_properties|jsonb||
