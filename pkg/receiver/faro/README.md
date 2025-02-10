# Faro Receiver

This receiver can receive telemetry data from the [Grafana Faro Web SDK](https://github.com/grafana/faro-web-sdk).
The telemetry data is in JSON format and adheres to the Faro OpenAPI schema, as defined in the [github.com/grafana/faro](https://github.com/grafana/faro)

## Receiver Configuration

The following receiver configuration parameters are supported.

| Name       | Description                       | Default        |
|:-----------|:----------------------------------|----------------|
| `endpoint` | Endpoint exposed by this receiver | localhost:8080 |


# Example Configuration

```yaml
receivers:
  faro:
    endpoint: 'localhost:8081'
```
