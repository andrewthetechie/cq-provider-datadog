
# Table: datadog_notebooks
Datadog Notebooks
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|meta_id|text|The Meta_ID of this notebook, combines datadog account name from config with ID and notebook name and hashes them to provide a unique PK.|
|id|integer|The ID of this notebook.|
|name|text|The name of this notebook.|
|created|timestamp without time zone|Timestamp of when the notebook was created|
|modified|timestamp without time zone|Timestamp of when the author was last modified|
|author|jsonb|Info about the author of this notebook including name, handle, and email|
|cells|jsonb|notebook's cells|
|metadata|jsonb|Notebook metadata|
|status|text|Status of the notebook |
|time|jsonb|notebook time data|
