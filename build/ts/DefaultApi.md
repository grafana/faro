# .DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**collectAppKeyPost**](DefaultApi.md#collectAppKeyPost) | **POST** /collect/{appKey} | Submit a Faro payload


# **collectAppKeyPost**
> void collectAppKeyPost(payload)


### Example


```typescript
import {  } from '';
import * as fs from 'fs';

const configuration = .createConfiguration();
const apiInstance = new .DefaultApi(configuration);

let body:.DefaultApiCollectAppKeyPostRequest = {
  // string | The application key is provided by your endpoint provider. 
  appKey: "appKey_example",
  // Payload | Optional description in *Markdown*
  payload: {
    events: [
      {
        name: "name_example",
        domain: "domain_example",
        attributes: {
          "key": "key_example",
        },
        timestamp: new Date('1970-01-01T00:00:00.00Z'),
      },
    ],
    logs: [
      {
        message: "message_example",
        level: "level_example",
        context: {
          "key": "key_example",
        },
        timestamp: new Date('1970-01-01T00:00:00.00Z'),
        trace: {
          traceId: "traceId_example",
          spanId: "spanId_example",
        },
      },
    ],
    exceptions: [
      {
        type: "type_example",
        value: "value_example",
        stacktrace: [
          {
            colno: 1,
            filename: "filename_example",
            _function: "_function_example",
            inApp: true,
            lineno: 1,
          },
        ],
        timestamp: new Date('1970-01-01T00:00:00.00Z'),
        trace: {
          traceId: "traceId_example",
          spanId: "spanId_example",
        },
        context: {
          "key": "key_example",
        },
      },
    ],
    measurements: [
      {
        values: {
          "key": 3.14,
        },
        type: "type_example",
        timestamp: new Date('1970-01-01T00:00:00.00Z'),
        trace: {
          traceId: "traceId_example",
          spanId: "spanId_example",
        },
        context: {
          "key": "key_example",
        },
      },
    ],
    meta: {
      sdk: {
        name: "name_example",
        version: "version_example",
      },
      app: {
        name: "name_example",
        release: "release_example",
        version: "version_example",
        environment: "environment_example",
      },
      user: {
        username: "username_example",
        id: "id_example",
        email: "email_example",
        attributes: {
          "key": "key_example",
        },
      },
      session: {
        id: "id_example",
        attributes: {
          "key": "key_example",
        },
      },
      page: {
        url: "url_example",
      },
      browser: {
        name: "name_example",
        version: "version_example",
        os: "os_example",
        mobile: true,
      },
      view: {
        name: "name_example",
      },
    },
    traces: {
      resourceSpans: [
        {
          resource: {
            attributes: {
              key: "key_example",
              value: {
                "key": "key_example",
              },
            },
          },
          instrumentationLibrarySpans: [
            {
              instrumentationLibrary: {
                name: "name_example",
                version: "version_example",
              },
              spans: [
                {
                  attributes: {
                    key: "key_example",
                    value: {},
                  },
                  droppedAttributesCount: 1,
                  droppedEventsCount: 1,
                  droppedLinksCount: 1,
                  endTimeUnixNano: 1,
                  events: [
                    {
                      attributes: [
                        "attributes_example",
                      ],
                      droppedAttributesCount: 1,
                      name: "name_example",
                      timeUnixNano: "timeUnixNano_example",
                    },
                  ],
                  kind: 1,
                  links: [
                    {
                      "key": "key_example",
                    },
                  ],
                  name: "name_example",
                  parentSpanId: "parentSpanId_example",
                  spanId: "spanId_example",
                  startTimeUnixNano: 1,
                  status: {
                    code: 1,
                    message: "message_example",
                  },
                  traceId: "traceId_example",
                },
              ],
            },
          ],
        },
      ],
    },
  },
};

apiInstance.collectAppKeyPost(body).then((data:any) => {
  console.log('API called successfully. Returned data: ' + data);
}).catch((error:any) => console.error(error));
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | **Payload**| Optional description in *Markdown* |
 **appKey** | [**string**] | The application key is provided by your endpoint provider.  | defaults to undefined


### Return type

**void**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**202** | OK |  -  |
**0** | Unexpected error |  -  |

[[Back to top]](#) [[Back to API list]](README.md#documentation-for-api-endpoints) [[Back to Model list]](README.md#documentation-for-models) [[Back to README]](README.md)


