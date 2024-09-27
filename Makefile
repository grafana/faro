buildDir = build
specDir = spec

goConfig = configs/go-config.yaml
tsConfig = configs/ts-config.yaml

genSpec = $(specDir)/gen/faro.gen.yaml
genGo = $(buildDir)/go/gen/faro.gen.go


serve:
	docker run -p 80:8080 -e SWAGGER_JSON=/src/openapi.yaml -v ./\:/src swaggerapi/swagger-ui

install-dependencies: install-go-dependencies

merge-specs:
	@./scripts/yaml_compose.sh $(specDir)

clean-go:
	@echo "Cleaning go generated file: $(genGo)"
	@rm -f $(genGo)

install-go-dependencies:
	go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest

build-go: merge-specs clean-go
	@echo "Building go generated file: $(genGo)"
	@oapi-codegen --config $(goConfig) $(genSpec)
	@echo "Post processing go generated file: $(genGo)"
	@./scripts/go_post_process.sh $(genGo)
	@echo "Running go mod tidy in build/go"
	@cd build/go && go mod tidy && cd ../..
	@echo "Done"

build-all: build-go # build-ts

# build-ts: clean-ts
# 	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
# 	-i /local/$(spec) \
# 	-c /local/$(tsConfig) \
# 	-g typescript \
# 	-o /local/$(buildDir)/ts

# clean-ts:
# 	rm -rf $(buildDir)/ts/*