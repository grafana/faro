# Faro Exporter

## Schema supported

This exporter targets to support the Faro OpenAPI schema in JSON format, as defined in github.com/grafana/faro

## Exporter Configuration

The following exporter configuration parameters are supported.

| Name       | Description                                | Default     |
|:-----------|:-------------------------------------------|-------------|
| `endpoint` | Target endpoint that accepts Faro payloads |             |


# Example Configuration

Following example configuration defines to send Faro payload to endpoint https://faro.example.com/collect

```yaml
exporters:
  faro:
    endpoint: 'https://faro.example.com/collect'
```
