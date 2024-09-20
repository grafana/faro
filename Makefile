spec = spec/faro.yaml
buildDir = build

serve:
	docker run -p 80:8080 -e SWAGGER_JSON=/src/openapi.yaml -v ./:/src swaggerapi/swagger-ui
build-all: build-go build-ts
clean-go:
	rm -rf $(buildDir)/go/*
build-go:  clean-go
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
    -i /local/$(spec) \
    -c /local/spec/go-config.yaml \
    -g go \
    -o /local/$(buildDir)/go
	cd build/go && go mod tidy && cd ../..
build-ts:  clean-ts
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
    -i /local/$(spec) \
    -c /local/spec/ts-config.yaml \
    -g typescript \
    -o /local/$(buildDir)/ts
clean-ts:
	rm -rf $(buildDir)/ts/*