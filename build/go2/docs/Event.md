# Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to **map[string]string** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Trace** | Pointer to [**TraceContext**](TraceContext.md) |  | [optional] 

## Methods

### NewEvent

`func NewEvent(name string, ) *Event`

NewEvent instantiates a new Event object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventWithDefaults

`func NewEventWithDefaults() *Event`

NewEventWithDefaults instantiates a new Event object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *Event) GetAttributes() map[string]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Event) GetAttributesOk() (*map[string]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Event) SetAttributes(v map[string]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Event) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetDomain

`func (o *Event) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Event) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Event) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *Event) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetName

`func (o *Event) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Event) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Event) SetName(v string)`

SetName sets Name field to given value.


### GetTimestamp

`func (o *Event) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Event) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Event) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *Event) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTrace

`func (o *Event) GetTrace() TraceContext`

GetTrace returns the Trace field if non-nil, zero value otherwise.

### GetTraceOk

`func (o *Event) GetTraceOk() (*TraceContext, bool)`

GetTraceOk returns a tuple with the Trace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrace

`func (o *Event) SetTrace(v TraceContext)`

SetTrace sets Trace field to given value.

### HasTrace

`func (o *Event) HasTrace() bool`

HasTrace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


