# Configuring Threatseer

# agent

Nothing required except server endpoint (currently hardcoded to localhost)

# server

## Daemon

Daemon-level configuration options are in [`/config/daemon.yaml`](/config/daemon.yaml).
The daemon config is self-documented.

## Logging
Configure [`beats.yml`](/beats.yml) per [the docs](https://www.elastic.co/guide/en/beats/filebeat/current/configuring-output.html) or the documentation in [the file](/beats.yml).

## Analysis Engines

Configure the `yaml` files in the `config` folder to your needs.

### Dynamic Rules Engine syntax

Example queries tested [here](https://github.com/caibirdme/yql/blob/master/yql_test.go#L901)