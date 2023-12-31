# golerta-cli

## (This repo is a mirror)

golerta-cli is a simple CLI client that implements a subset of the alerta CLI functionality.

Currently sending alerts and sending heartbeats are supported. Even there some parameters might be missing since I've been testing the client based on my own needs. The official alerta CLI is a Python program.
I've written golerta-cli in Go to make the deployment of the client on several hosts easier. Although the resulting binary is about 7mb (without compression with upx, which brings it down to around 3mb), static linking ensures that the resulting single binary will be very easy to distribute. 

Please refer to the following help output for supported parameters.

# Configuration

All configuration options and parameters can come from three different sources. These are, command line parameters (`--environment`), environment variables (`GOLERTA_CLI_ENVIRONMENT`) or a configuration file (`environment=xxxx`). The precedence order from higher to lower is, command line parameters, environment variables and finally the configuration file. The default configuration file is named '.golerta-cli' and is looked for in the current directory. This can be overridden with the `--config` parameter.

# For sending alerts

```
golerta-cli send --help

Send an alert to alerta endpoint.

Usage:
  golerta-cli send [flags]

Flags:
      --environment string    Environment string
  -e, --event string          Event string (mandatory)
  -g, --group string          Group string
  -h, --help                  help for send
  -o, --origin string         Origin string
      --raw-data string       Raw data string
  -r, --resource string       Resource string (mandatory)
  -x, --service stringArray   Service (multiple invokation allowed)
  -s, --severity string       Severity ('ok', 'normal', 'major', 'minor', 'critical') (default "normal")
      --tag stringArray       Tags (multiple invokation allowed)
  -T, --text string           Text string
      --timeout int           Timeout (integer)
  -t, --type string           Event type string
      --value int             Integer value

Global Flags:
  -a, --apikey string     Apikey (Mandatory)
  -c, --config string     config file (default is ./.golerta-cli)
      --curl              Generate a curl command representation of gathered parameters for testing
  -d, --debug             Display info useful for debugging
      --dryrun            Display info but don't post the endpoint
  -E, --endpoint string   Endpoint (Mandatory)
```

# For sending heartbeats

```
golerta-cli heartbeat --help

Send a heartbeat to alerta endpoint

Usage:
  golerta-cli heartbeat [flags]

Flags:
  -h, --help              help for heartbeat
  -o, --origin string     Origin string
      --tag stringArray   Tags
      --timeout int       Timeout (integer)

Global Flags:
  -a, --apikey string     Apikey (Mandatory)
  -c, --config string     config file (default is ./.golerta-cli)
      --curl              Generate a curl command representation of gathered parameters for testing
  -d, --debug             Display info useful for debugging
      --dryrun            Display info but don't post the endpoint
  -E, --endpoint string   Endpoint (Mandatory)
```


# Details

The program uses the [Cobra](https://github.com/spf13/cobra) and [Viper](https://github.com/spf13/viper) Go libraries which basically do all the heavy lifting. The remainder is just making sure the JSON format is correct and pushing to the alerta server api. I was using a bash script before writing golerta-cli to do basically the same thing but handling JSON variables in bash quickly turns into escape character hell even when using jq ([see my beginner post about that here](https://xpufx.com/posts/building_a_json_string_with_shell_variables_explained/)) to do the formatting.
