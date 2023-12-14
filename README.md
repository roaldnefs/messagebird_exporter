<a href="https://github.com/roaldnefs/messagebird_exporter" style="color: black;">
    <h1 align="center">MessageBird Exporter</h1>
</a>
<p align="center">
    <a href="https://github.com/roaldnefs/messagebird_exporter/releases">
        <img src="https://img.shields.io/github/v/release/roaldnefs/messagebird_exporter?style=for-the-badge&color=blue"
            alt="Latest release version">
    </a>
    <a href="https://godoc.org/github.com/roaldnefs/messagebird_exporter">
        <img src="https://img.shields.io/badge/godoc-reference-5272B4.svg?style=for-the-badge&color=blue"
            alt="GoDoc">
    </a>
    <a href="https://github.com/roaldnefs/messagebird_exporter/blob/master/LICENSE">
        <img src="https://img.shields.io/github/license/roaldnefs/messagebird_exporter.svg?style=for-the-badge&color=blue"
            alt="GitHub - License">
    </a>
    <a href="https://github.com/roaldnefs/messagebird_exporter/actions">
        <img src="https://img.shields.io/github/actions/workflow/status/roaldnefs/messagebird_exporter/build.yaml?style=for-the-badge&color=blue"
            alt="GitHub Workflow Status">
    </a>
    <a href="https://github.com/roaldnefs/messagebird_exporter/graphs/contributors">
        <img src="https://img.shields.io/github/contributors/roaldnefs/messagebird_exporter?style=for-the-badge&color=blue"
            alt="GitHub contributors">
    </a>
    </br>
    <b>messagebird_exporter</b> is a Prometheus exporter for <a href="https://messagebird.com/en/">MessageBird</a> metrics, written in Go.
    <br />
    <a href="https://godoc.org/github.com/roaldnefs/messagebird_exporter"><strong>Explore the docs »</strong></a>
    <br />
    <a href="https://github.com/roaldnefs/messagebird_exporter/issues/new?title=Bug%3A">Report Bug</a>
    ·
    <a href="https://github.com/roaldnefs/messagebird_exporter/issues/new?title=Feature+Request%3A">Request Feature</a>
</p>

## Introduction
Prometheus exporter for [MessageBird](https://messagebird.com/en/) metrics, written in Go.

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
      --version  Show application version.

$ messagebird_exporter --messagebird.api-key=test_myapikey
INFO[0000] Starting MessageBird Exporter                 version=unknown
INFO[0000] Listening on :9601                            listen_address=":9601" metrics_path=/metrics
```
