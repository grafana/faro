package faro

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"

	faroTypes "github.com/grafana/faro/pkg/go"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/otel"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

// TranslateFromTraces converts a Traces pipeline data into []*faro.Payload
func TranslateFromTraces(ctx context.Context, td ptrace.Traces) ([]*faroTypes.Payload, error) {
	ctx, span := otel.Tracer("").Start(ctx, "TranslateFromTraces")
	defer span.End()

	resourceSpans := td.ResourceSpans()
	resourceSpansLen := resourceSpans.Len()
	if resourceSpansLen == 0 {
		return nil, nil
	}
	metaMap := make(map[string]*faroTypes.Payload, 0)
	payloads := make([]*faroTypes.Payload, 0)

	for i := 0; i < resourceSpansLen; i++ {
		rs := resourceSpans.At(i)
		payload := resourceSpansToFaroPayload(rs)
		if payload == nil {
			continue
		}

		meta := payload.Meta
		// if payload meta already exists in the metaMap merge payload to the existing payload
		w := md5.New()
		if encodeErr := json.NewEncoder(w).Encode(meta); encodeErr != nil {
			return payloads, encodeErr
		}
		metaKey := fmt.Sprintf("%x", w.Sum(nil))
		existingPayload, found := metaMap[metaKey]
		if found {
			// merge payloads with the same meta
			mergePayloads(existingPayload, payload)
		} else {
			// if payload meta doesn't exist in the metaMap add new meta key to the metaMap
			metaMap[metaKey] = payload
		}
	}
	for _, payload := range metaMap {
		payloads = append(payloads, payload)
	}
	return payloads, nil
}

func resourceSpansToFaroPayload(rs ptrace.ResourceSpans) *faroTypes.Payload {
	resource := rs.Resource()
	scopeSpans := rs.ScopeSpans()

	if resource.Attributes().Len() == 0 && scopeSpans.Len() == 0 {
		return nil
	}

	traces := ptrace.NewTraces()
	rs.CopyTo(traces.ResourceSpans().AppendEmpty())

	meta := extractMetaFromResourceAttributes(resource.Attributes())
	payload := &faroTypes.Payload{
		Meta: meta,
		Traces: &faroTypes.Traces{
			Traces: traces,
		},
	}

	return payload
}

func extractMetaFromResourceAttributes(resourceAttributes pcommon.Map) faroTypes.Meta {
	var meta faroTypes.Meta
	var app faroTypes.App

	if appName, ok := resourceAttributes.Get(string(semconv.ServiceNameKey)); ok {
		app.Name = appName.Str()
	}
	if appNamespace, ok := resourceAttributes.Get(string(semconv.ServiceNamespaceKey)); ok {
		app.Namespace = appNamespace.Str()
	}
	if version, ok := resourceAttributes.Get(string(semconv.ServiceVersionKey)); ok {
		app.Version = version.Str()
	}
	if environment, ok := resourceAttributes.Get(string(semconv.DeploymentEnvironmentKey)); ok {
		app.Environment = environment.Str()
	}
	if appBundleID, ok := resourceAttributes.Get("app_bundle_id"); ok {
		app.BundleID = appBundleID.Str()
	}

	meta.App = app
	return meta
}
