# \DefaultAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CollectAppKeyPost**](DefaultAPI.md#CollectAppKeyPost) | **Post** /collect/{appKey} | Submit a Faro payload



## CollectAppKeyPost

> CollectAppKeyPost(ctx, appKey).Payload(payload).Execute()

Submit a Faro payload

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/grafana/faro"
)

func main() {
	appKey := "appKey_example" // string | The application key is provided by your endpoint provider. 
	payload := *openapiclient.NewPayload(*openapiclient.NewPayloadMeta()) // Payload | Optional description in *Markdown*

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.CollectAppKeyPost(context.Background(), appKey).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CollectAppKeyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** | The application key is provided by your endpoint provider.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCollectAppKeyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**Payload**](Payload.md) | Optional description in *Markdown* | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

