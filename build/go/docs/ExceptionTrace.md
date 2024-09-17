# ExceptionTrace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TraceId** | Pointer to **string** |  | [optional] 
**SpanId** | Pointer to **string** |  | [optional] 

## Methods

### NewExceptionTrace

`func NewExceptionTrace() *ExceptionTrace`

NewExceptionTrace instantiates a new ExceptionTrace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExceptionTraceWithDefaults

`func NewExceptionTraceWithDefaults() *ExceptionTrace`

NewExceptionTraceWithDefaults instantiates a new ExceptionTrace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTraceId

`func (o *ExceptionTrace) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *ExceptionTrace) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *ExceptionTrace) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *ExceptionTrace) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.

### GetSpanId

`func (o *ExceptionTrace) GetSpanId() string`

GetSpanId returns the SpanId field if non-nil, zero value otherwise.

### GetSpanIdOk

`func (o *ExceptionTrace) GetSpanIdOk() (*string, bool)`

GetSpanIdOk returns a tuple with the SpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanId

`func (o *ExceptionTrace) SetSpanId(v string)`

SetSpanId sets SpanId field to given value.

### HasSpanId

`func (o *ExceptionTrace) HasSpanId() bool`

HasSpanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


