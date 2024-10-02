# Exception

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to **map[string]string** |  | [optional] 
**Stacktrace** | Pointer to [**Stacktrace**](Stacktrace.md) |  | [optional] 
**Timestamp** | **time.Time** |  | 
**Trace** | Pointer to [**TraceContext**](TraceContext.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 

## Methods

### NewException

`func NewException(timestamp time.Time, ) *Exception`

NewException instantiates a new Exception object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExceptionWithDefaults

`func NewExceptionWithDefaults() *Exception`

NewExceptionWithDefaults instantiates a new Exception object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *Exception) GetContext() map[string]string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *Exception) GetContextOk() (*map[string]string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *Exception) SetContext(v map[string]string)`

SetContext sets Context field to given value.

### HasContext

`func (o *Exception) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetStacktrace

`func (o *Exception) GetStacktrace() Stacktrace`

GetStacktrace returns the Stacktrace field if non-nil, zero value otherwise.

### GetStacktraceOk

`func (o *Exception) GetStacktraceOk() (*Stacktrace, bool)`

GetStacktraceOk returns a tuple with the Stacktrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStacktrace

`func (o *Exception) SetStacktrace(v Stacktrace)`

SetStacktrace sets Stacktrace field to given value.

### HasStacktrace

`func (o *Exception) HasStacktrace() bool`

HasStacktrace returns a boolean if a field has been set.

### GetTimestamp

`func (o *Exception) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Exception) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Exception) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetTrace

`func (o *Exception) GetTrace() TraceContext`

GetTrace returns the Trace field if non-nil, zero value otherwise.

### GetTraceOk

`func (o *Exception) GetTraceOk() (*TraceContext, bool)`

GetTraceOk returns a tuple with the Trace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrace

`func (o *Exception) SetTrace(v TraceContext)`

SetTrace sets Trace field to given value.

### HasTrace

`func (o *Exception) HasTrace() bool`

HasTrace returns a boolean if a field has been set.

### GetType

`func (o *Exception) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Exception) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Exception) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Exception) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *Exception) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Exception) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Exception) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *Exception) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


