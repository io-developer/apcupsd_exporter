apcupsd_exporter
[![Build Status](https://travis-ci.org/mdlayher/apcupsd_exporter.svg?branch=master)](https://travis-ci.org/mdlayher/apcupsd_exporter) [![GoDoc](http://godoc.org/github.com/mdlayher/apcupsd_exporter?status.svg)](http://godoc.org/github.com/mdlayher/apcupsd_exporter)
[![Docker Automated build](https://img.shields.io/docker/automated/iodeveloper/prom_apcupsdexporter.svg)](https://hub.docker.com/r/iodeveloper/prom_apcupsdexporter/)
================

Command `apcupsd_exporter` provides a Prometheus exporter for the
[apcupsd](http://www.apcupsd.org/) Network Information Server (NIS).
MIT Licensed.


Docker
------
[![Docker Hub repository](http://dockeri.co/image/iodeveloper/prom_apcupsdexporter)](https://registry.hub.docker.com/u/iodeveloper/prom_apcupsdexporter/)

`iodeveloper/prom_apcupsdexporter`


Usage
-----

Available flags for `apcupsd_exporter` include:

```
$ ./apcupsd_exporter -h
Usage of ./apcupsd_exporter:
  -apcupsd.addr string
        address of apcupsd Network Information Server (NIS) (default ":3551")
  -apcupsd.network string
        network of apcupsd Network Information Server (NIS): typically "tcp", "tcp4", or "tcp6" (default "tcp")
  -telemetry.addr string
        address for apcupsd exporter (default ":9162")
  -telemetry.path string
        URL path for surfacing collected metrics (default "/metrics")
```

Docker
------
https://hub.docker.com/r/iodeveloper/prom_apcupsdexporter
```
iodeveloper/prom_apcupsdexporter:latest
```

**docker-compose.yml** example:
```yaml
version: '3.4'

services:

  apcupsd-home-main:
    image: iodeveloper/prom_apcupsdexporter:latest
    user: root
    privileged: true
    restart: unless-stopped
    network_mode: "host"
    command: /apcupsd_exporter -apcupsd.addr="0.0.0.0:3551" -telemetry.addr="172.101.0.1:10001"
    
  apcupsd-home-server:
    image: iodeveloper/prom_apcupsdexporter:latest
    user: root
    privileged: true
    restart: unless-stopped
    network_mode: "host"
    command: /apcupsd_exporter -apcupsd.addr="0.0.0.0:3552" -telemetry.addr="172.101.0.1:10002"
```

Sample
------

Here is a screenshot of an example grafana dashboard (/examples/grafana-dashboard.json) using metrics from
`apcupsd_exporter`.

![Grafana 1](/examples/grafana-1.png?raw=true)

![Grafana 2](/examples/grafana-2.png?raw=true)
