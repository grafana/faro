spec = spec/faro.yaml
buildDir = build

serve:
	docker run -p 80:8080 -e SWAGGER_JSON=/src/openapi.yaml -v ./:/src swaggerapi/swagger-ui
build-all: build-go build-ts
clean-go:
	rm -rf $(buildDir)/go/*
build-go:  clean-go
	docker run --rm \
    -e GO_POST_PROCESS_FILE="/usr/local/bin/gofmt -w" \
    -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
    -i /local/$(spec) \
    -c /local/spec/go-config.yaml \
    -g go \
    -o /local/$(buildDir)/go \
    --import-mappings=ptrace.Traces=go.opentelemetry.io/collector/pdata/ptrace \
    --type-mappings=object+Traces=ptrace.Traces
	cd build/go && go mod tidy && cd ../..
build-ts:  clean-ts
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
    -i /local/$(spec) \
    -c /local/spec/ts-config.yaml \
    -g typescript \
    -o /local/$(buildDir)/ts
clean-ts:
	rm -rf $(buildDir)/ts/*
