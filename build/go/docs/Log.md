# Log

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Level** | **string** |  | 
**Context** | Pointer to **map[string]string** |  | [optional] 
**Timestamp** | **time.Time** |  | 
**Trace** | Pointer to [**LogTrace**](LogTrace.md) |  | [optional] 

## Methods

### NewLog

`func NewLog(message string, level string, timestamp time.Time, ) *Log`

NewLog instantiates a new Log object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogWithDefaults

`func NewLogWithDefaults() *Log`

NewLogWithDefaults instantiates a new Log object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *Log) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Log) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Log) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetLevel

`func (o *Log) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *Log) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *Log) SetLevel(v string)`

SetLevel sets Level field to given value.


### GetContext

`func (o *Log) GetContext() map[string]string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *Log) GetContextOk() (*map[string]string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *Log) SetContext(v map[string]string)`

SetContext sets Context field to given value.

### HasContext

`func (o *Log) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetTimestamp

`func (o *Log) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Log) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Log) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetTrace

`func (o *Log) GetTrace() LogTrace`

GetTrace returns the Trace field if non-nil, zero value otherwise.

### GetTraceOk

`func (o *Log) GetTraceOk() (*LogTrace, bool)`

GetTraceOk returns a tuple with the Trace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrace

`func (o *Log) SetTrace(v LogTrace)`

SetTrace sets Trace field to given value.

### HasTrace

`func (o *Log) HasTrace() bool`

HasTrace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


