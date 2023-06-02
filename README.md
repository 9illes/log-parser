# Log viewer

## Description

This project is a log viewer for cloudfront logs. It allows to view the logs in a web interface and to filter
the requests by project, date, status code, googlebot, etc.

It is an experimental project ported to Go for testing purpose and is not crafted for large log volume.

## Steps

1. Download all cloudfront log files from an AWS S3 bucket
2. Aggregate the log files into a single file using the `scripts\aggregate.sh` script
3. Preprocess the main log file (split into multiple files)
4. Start the web interface and open [http://localhost:8080](http://localhost:8080)

## Usage

### Build

```sh
make build
```

```sh
build/logs preprocess -s default -i var/23-05-25.log -o var
```

## Development

### Help

```sh
make
```

### Preprocess log requests

```sh
# Create preprocessed log files for the 'default' project using the cloudfront log file 'var/23-05-25.log'
go run main.go preprocess -s default -i var/23-05-25.log -o var
```

### Start the web interface

Autorebuild and restart the web interface when a file is changed. Require [Air](github.com/cosmtrek/air) to be installed (`go install github.com/cosmtrek/air@latest`).
Install `air` to automatically rebuild and restart the web interface when a file is changed.

```sh
make live
```

### Enable profiling

pprof is disabled by default. To enalbe it, use the `-i` flag on the serve sub-command.

```sh
go run main.go serve -i
```

Open [pprof Web UI](http://localhost:6060/debug/pprof/)

## Resources

* [Bulma](https://bulma.io/) CSS framework
* [Tabulator](https://tabulator.info/) JS table library
* [Air](github.com/cosmtrek/air) Live reload for Go apps