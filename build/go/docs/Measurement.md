# Measurement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Values** | **map[string]float32** |  | 
**Type** | **string** |  | 
**Timestamp** | **time.Time** |  | 
**Trace** | Pointer to [**ExceptionTrace**](ExceptionTrace.md) |  | [optional] 
**Context** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewMeasurement

`func NewMeasurement(values map[string]float32, type_ string, timestamp time.Time, ) *Measurement`

NewMeasurement instantiates a new Measurement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeasurementWithDefaults

`func NewMeasurementWithDefaults() *Measurement`

NewMeasurementWithDefaults instantiates a new Measurement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValues

`func (o *Measurement) GetValues() map[string]float32`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *Measurement) GetValuesOk() (*map[string]float32, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *Measurement) SetValues(v map[string]float32)`

SetValues sets Values field to given value.


### GetType

`func (o *Measurement) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Measurement) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Measurement) SetType(v string)`

SetType sets Type field to given value.


### GetTimestamp

`func (o *Measurement) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Measurement) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Measurement) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetTrace

`func (o *Measurement) GetTrace() ExceptionTrace`

GetTrace returns the Trace field if non-nil, zero value otherwise.

### GetTraceOk

`func (o *Measurement) GetTraceOk() (*ExceptionTrace, bool)`

GetTraceOk returns a tuple with the Trace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrace

`func (o *Measurement) SetTrace(v ExceptionTrace)`

SetTrace sets Trace field to given value.

### HasTrace

`func (o *Measurement) HasTrace() bool`

HasTrace returns a boolean if a field has been set.

### GetContext

`func (o *Measurement) GetContext() map[string]string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *Measurement) GetContextOk() (*map[string]string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *Measurement) SetContext(v map[string]string)`

SetContext sets Context field to given value.

### HasContext

`func (o *Measurement) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


