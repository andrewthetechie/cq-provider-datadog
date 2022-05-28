
# Table: datadog_hosts
Host Object representing a host.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|aliases|text[]|Host aliases collected by Datadog.|
|apps|text[]|The Datadog integrations reporting metrics for the host.|
|aws_name|text|AWS name of your host.|
|host_name|text|The host name.|
|id|bigint|The host ID.|
|is_muted|boolean|If a host is muted or unmuted.|
|last_reported_time|bigint|Last time the host reported a metric data point.|
|meta_agent_version|text|The Datadog Agent version.|
|meta_cpu_cores|bigint|The number of cores.|
|meta_fbsd_v|text[]|An array of Mac versions.|
|meta_gohai|text|JSON string containing system information.|
|meta_install_method_installer_version|text|The installer version.|
|meta_install_method_tool|text|Tool used to install the agent.|
|meta_install_method_tool_version|text|The tool version.|
|meta_install_method_additional_properties|jsonb||
|meta_mac_v|text[]|An array of Mac versions.|
|meta_machine|text|The machine architecture.|
|meta_nix_v|text[]|Array of Unix versions.|
|meta_platform|text|The OS platform.|
|meta_processor|text|The processor.|
|meta_python_v|text|The Python version.|
|meta_win_v|text[]|An array of Windows versions.|
|meta_additional_properties|jsonb||
|metrics_cpu|float|The percent of CPU used (everything but idle).|
|metrics_iowait|float|The percent of CPU spent waiting on the IO (not reported for all platforms).|
|metrics_load|float|The system load over the last 15 minutes.|
|metrics_additional_properties|jsonb||
|mute_timeout|bigint|Timeout of the mute applied to your host.|
|name|text|The host name.|
|sources|text[]|Source or cloud provider associated with your host.|
|tags_by_source|jsonb|List of tags for each source (AWS, Datadog Agent, Chef..).|
|up|boolean|Displays UP when the expected metrics are received and displays `???` if no metrics are received.|
|additional_properties|jsonb||
