# HTTP middleman for app-to-InfluxDB

From InfluxDB documentation on the v1.8 HTTP API:

    curl -i -XPOST 'http://localhost:8086/write?db=mydb' --data-binary 'cpu_load_short,host=server01,region=us-west value=0.64 1434055562000000000'

## Installation

```sh
go get github.com/asmundstavdahl/iridium-to-influxdb
go install github.com/asmundstavdahl/iridium-to-influxdb
iridium-to-influxdb -help
```

## Usage

Scenario 1:

- middleman is to be run on the same host as InfluxDB (IP 5.6.7.8)
- InfluxDB HTTP API is available on port 8086
- middleman is to listen for POST request on port 1234
- data is temperature measurements which are to be stored in the "temp" database

On your InfluxDB server:

```sh
iridium-to-influxdb \
    -database-host localhost \
    -database-port 8086 \
    -database-name temp \
    -server-port   1234
```

On your Iridium-connected data aquisition device:

```sh
curl -i -XPOST "http://5.6.7.8:1234/" --data-raw "temp,sensor=livingroom1 value=23 1592398708000000000"
```
