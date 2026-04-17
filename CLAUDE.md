# Faro — Claude Notes

## What this repo is

OpenAPI specification and generated HTTP model packages for the [Grafana Faro](https://github.com/grafana/faro-web-sdk) observability API. Spec-first: the OpenAPI spec is the source of truth, and Go types are generated from it.

## Build system

```sh
make build-all-docker   # recommended — no local deps needed
make build-all          # local build (requires yq + oapi-codegen)
make serve              # Swagger UI for the generated spec
```

**Full pipeline:**

1. `yaml_compose.sh` merges `spec/components/**/*.yaml` + `spec/info/`, `spec/paths/`, `spec/servers/` into `spec/gen/faro.gen.yaml`
2. `oapi-codegen` generates `pkg/go/faro.gen.go` from the composed spec
3. `go_post_process.sh` patches the generated file (see gotcha below)
4. `go mod tidy` in `pkg/go/`

## Critical gotchas

**Python yq, not mikefarah/yq.** `yaml_compose.sh` uses `yq --yaml-output --slurp`, which is Python yq (kislyuk) syntax. The Go-based mikefarah/yq (common on Linux via package managers) has entirely different flags and causes `Error: unknown flag: --yaml-output` for every file. If both are installed, the Python yq must come first in `PATH` (it installs to `~/.local/bin/yq`). Verify with `yq --version` — it should say `yq 3.x.x`, not `yq (https://github.com/mikefarah/yq/)`.

**Post-processing is load-bearing.** `go_post_process.sh` converts the generated `type Traces = ptrace.Traces` (type alias) into an embedded struct. Without this patch, the custom `MarshalJSON`/`UnmarshalJSON` in `traces.go` cannot be defined (Go doesn't allow methods on type aliases of external types).

**`spec/gen/faro.gen.yaml` and `spec/entrypoint.yaml` must be excluded from the `find` loop in `yaml_compose.sh`.** Including them causes null values to overwrite the composed spec (`servers: null`, `components: null`). The `find` uses `-not -path` exclusions for both.

**Schema ordering in the composed spec is determined by `find | sort` output** (alphabetical by file path). This is intentional for determinism.

## oapi-codegen config (`configs/go-config.yaml`)

- Generates models only (`generate.models: true`) — no server stubs or client code
- Package `faro`, output `pkg/go/faro.gen.go`
- `skip-prune: true` — keeps types that aren't directly referenced in paths
- `name-normalizer: ToCamelCaseWithInitialisms` — so `span_id` → `SpanID`, `asn_org` → `ASNOrg`, `build_id` → `BuildID`
- `always-prefix-enum-values: true` — enum values get the type name as prefix

## Key schema conventions

- `x-go-type-skip-optional-pointer: true` is used extensively to generate value types instead of pointers
- `x-omitempty: false` on `TraceContext` fields ensures zero-value span/trace IDs are always serialized
- `Browser.brands` uses `oneOf: [BrandsString, BrandsArray]` — oapi-codegen generates a union type with `As*/From*/Merge*` helpers

**Mixed property-name casing.** The repo has both styles: `browser.yaml` / `app.yaml` use camelCase (`bundleId`, `userAgent`, `viewportWidth`), while `geo.yaml` / `tracecontext.yaml` use snake_case (`asn_org`, `span_id`, `trace_id`). **Prefer snake_case for new schemas that map to OTel semconv** — it mirrors OTel attribute keys on the wire and `ToCamelCaseWithInitialisms` correctly renders them in Go (`SpanID`, `ASNOrg`, `BuildID`). Match the surrounding style when extending an existing schema.

**`x-go-name` is silently ignored next to `$ref`.** Per OpenAPI 3.0, siblings of `$ref` are ignored by resolvers, so `x-go-name: Foo` on a `$ref` property does nothing — the Go field name is derived from the property key via the normalizer. The only clean fix is an `allOf: [$ref: ...]` wrapper (siblings survive composition), but that can produce a wrapper type. `x-go-name` works fine on non-ref properties (see `Browser.os` string, which uses `x-go-name: OS` successfully).

## Tooling versions

- oapi-codegen: **v2.6.0** (pinned in both `Makefile` and `Dockerfile.build`)
- Go: **1.24**
- Python yq: latest (no pin)

## Docker build image (`Dockerfile.build`)

Based on `golang:1.24` (Debian bookworm). Installs `python3`, `python3-pip`, `jq` (required by Python yq), and oapi-codegen. The `--break-system-packages` flag is needed for pip inside this image due to PEP 668 enforcement in Debian 12.

## TypeScript generation

Commented out in the Makefile (`build-ts`). Config exists at `configs/ts-config.yaml` but the feature is incomplete.

## Pull requests

Do not include a test plan section when opening PRs.

## What is NOT in this repo

- The Faro Web SDK → https://github.com/grafana/faro-web-sdk
- The Faro ↔ OTel translator → contributed upstream to `opentelemetry-collector-contrib/pkg/translator/faro`
- `pkg/translator/faro/` is a placeholder README only
