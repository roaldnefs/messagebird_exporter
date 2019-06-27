# MessageBird exporter

[![Travis CI](https://img.shields.io/travis/roaldnefs/messagebird_exporter.svg?style=for-the-badge)](https://travis-ci.org/roaldnefs/messagebird_exporter)
[![GoDoc](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=for-the-badge)](https://godoc.org/github.com/roaldnefs/messagebird_exporter)
[![Github All Releases](https://img.shields.io/github/downloads/roaldnefs/messagebird_exporter/total.svg?style=for-the-badge)](https://github.com/roaldnefs/messagebird_exporter/releases)
[![GitHub](https://img.shields.io/github/license/roaldnefs/messagebird_exporter.svg?style=for-the-badge)](https://github.com/roaldnefs/messagebird_exporter/blob/master/LICENSE)

Prometheus exporter for MessageBird metrics, written in Go.

* [Installation](README.md#installation)
     * [Binaries](README.md#binaries)
	 * [Via Go](README.md#via-go)
* [Usage](README.md#usage)

## Installation

### Binaries

For installation instructions from binaries please visit the [Release Page](https://github.com/roaldnefs/messagebird_exporter/releases).

### Via Go

```console
$ go get github.com/roaldnefs/messagebird_exporter
```

## Usage

```console
$ messagebird_exporter --help
usage: messagebird_exporter --messagebird.api-key=MESSAGEBIRD.API-KEY [<flags>]

Flags:
  -h, --help  Show context-sensitive help (also try --help-long and --help-man).
      --web.listen-address=":9601"  
              Address on which to expose metrics and web interface.
      --web.telemetry-path="/metrics"  
              Path under which to expose metrics.
      --messagebird.api-key=MESSAGEBIRD.API-KEY  
              MessageBird Api Key.

$ messagebird_exporter --messagebird.api-key=test_myapikey
INFO[0000] Starting MessageBird Exporter                 version=unknown
INFO[0000] Listening on :9601                            listen_address=":9601" metrics_path=/metrics
```
