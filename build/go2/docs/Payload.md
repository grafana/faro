# Payload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to [**[]Event**](Event.md) |  | [optional] 
**Exceptions** | Pointer to [**[]Exception**](Exception.md) |  | [optional] 
**Logs** | Pointer to [**[]Log**](Log.md) |  | [optional] 
**Measurements** | Pointer to [**[]Measurement**](Measurement.md) |  | [optional] 
**Meta** | Pointer to [**Meta**](Meta.md) |  | [optional] 
**Traces** | Pointer to **interface{}** | otel traces model. | [optional] 

## Methods

### NewPayload

`func NewPayload() *Payload`

NewPayload instantiates a new Payload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayloadWithDefaults

`func NewPayloadWithDefaults() *Payload`

NewPayloadWithDefaults instantiates a new Payload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *Payload) GetEvents() []Event`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *Payload) GetEventsOk() (*[]Event, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *Payload) SetEvents(v []Event)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *Payload) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetExceptions

`func (o *Payload) GetExceptions() []Exception`

GetExceptions returns the Exceptions field if non-nil, zero value otherwise.

### GetExceptionsOk

`func (o *Payload) GetExceptionsOk() (*[]Exception, bool)`

GetExceptionsOk returns a tuple with the Exceptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptions

`func (o *Payload) SetExceptions(v []Exception)`

SetExceptions sets Exceptions field to given value.

### HasExceptions

`func (o *Payload) HasExceptions() bool`

HasExceptions returns a boolean if a field has been set.

### GetLogs

`func (o *Payload) GetLogs() []Log`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *Payload) GetLogsOk() (*[]Log, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *Payload) SetLogs(v []Log)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *Payload) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetMeasurements

`func (o *Payload) GetMeasurements() []Measurement`

GetMeasurements returns the Measurements field if non-nil, zero value otherwise.

### GetMeasurementsOk

`func (o *Payload) GetMeasurementsOk() (*[]Measurement, bool)`

GetMeasurementsOk returns a tuple with the Measurements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurements

`func (o *Payload) SetMeasurements(v []Measurement)`

SetMeasurements sets Measurements field to given value.

### HasMeasurements

`func (o *Payload) HasMeasurements() bool`

HasMeasurements returns a boolean if a field has been set.

### GetMeta

`func (o *Payload) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Payload) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Payload) SetMeta(v Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Payload) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetTraces

`func (o *Payload) GetTraces() interface{}`

GetTraces returns the Traces field if non-nil, zero value otherwise.

### GetTracesOk

`func (o *Payload) GetTracesOk() (*interface{}, bool)`

GetTracesOk returns a tuple with the Traces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraces

`func (o *Payload) SetTraces(v interface{})`

SetTraces sets Traces field to given value.

### HasTraces

`func (o *Payload) HasTraces() bool`

HasTraces returns a boolean if a field has been set.

### SetTracesNil

`func (o *Payload) SetTracesNil(b bool)`

 SetTracesNil sets the value for Traces to be an explicit nil

### UnsetTraces
`func (o *Payload) UnsetTraces()`

UnsetTraces ensures that no value is present for Traces, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

