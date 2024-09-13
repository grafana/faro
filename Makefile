serve:
	docker run -p 80:8080 -e SWAGGER_JSON=/src/openapi.yaml -v ./:/src swaggerapi/swagger-ui
clean-go:
	rm -rf build/go/oas*
build-go: clean-go
	docker run --rm --volume ".:/workspace" \
		ghcr.io/ogen-go/ogen:latest \
			--target="workspace/build/go/" \
			--package="faro" \
			--clean \
			workspace/spec/faro.yaml
clean-ts:
	rm -rf build/ts/*
build-ts: clean-ts
	npx openapi-typescript spec/faro.yaml -o build/ts/faro.d.ts
build: build-go build-ts