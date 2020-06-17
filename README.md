# HTTP middleman for app-to-InfluxDB

From InfluxDB documentation on the v1.8 HTTP API:

    curl -i -XPOST 'http://localhost:8086/write?db=mydb' --data-binary 'cpu_load_short,host=server01,region=us-west value=0.64 1434055562000000000'

## Usage

```sh
go get github.com/asmundstavdahl/iridium-to-influxdb.git
go install github.com/asmundstavdahl/iridium-to-influxdb.git
iridium-to-influxdb -help
```
