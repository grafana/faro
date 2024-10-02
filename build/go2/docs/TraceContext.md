# TraceContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SpanId** | **string** |  | 
**TraceId** | **string** |  | 

## Methods

### NewTraceContext

`func NewTraceContext(spanId string, traceId string, ) *TraceContext`

NewTraceContext instantiates a new TraceContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraceContextWithDefaults

`func NewTraceContextWithDefaults() *TraceContext`

NewTraceContextWithDefaults instantiates a new TraceContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpanId

`func (o *TraceContext) GetSpanId() string`

GetSpanId returns the SpanId field if non-nil, zero value otherwise.

### GetSpanIdOk

`func (o *TraceContext) GetSpanIdOk() (*string, bool)`

GetSpanIdOk returns a tuple with the SpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanId

`func (o *TraceContext) SetSpanId(v string)`

SetSpanId sets SpanId field to given value.


### GetTraceId

`func (o *TraceContext) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *TraceContext) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *TraceContext) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


