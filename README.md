# Grafana Faro

<p align="left"><img src="docs/faro_logo.png" alt="Grafana Faro logo" width="400"></p>

## About
A project for frontend application observability, Grafana Faro includes a highly configurable web SDK for real user monitoring (RUM) that instruments browser frontend applications to capture observability signals. The frontend telemetry can then be correlated with backend and infrastructure data for seamless, full-stack observability.

For the Grafana Faro Web SDK, please go to [https://github.com/grafana/faro-web-sdk](https://github.com/grafana/faro-web-sdk).

The repository consists of 
- [OpenAPI specification](./spec/gen/faro.gen.yaml)
- Packages with HTTP Models generated from the OpenAPI specification

## Requirements

- [OpenAPI Codegen](https://github.com/oapi-codegen/oapi-codegen)
- [Python YQ](https://pypi.org/project/yq/)

## Packages

### Go
[/pkg/go](./pkg/go) contains HTTP Models in Go generated from the [OpenAPI specification](./spec/gen/faro.gen.yaml) using [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen).    
