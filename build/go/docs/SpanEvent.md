# SpanEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to **[]string** |  | [optional] 
**DroppedAttributesCount** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  | 
**TimeUnixNano** | Pointer to **string** |  | [optional] 

## Methods

### NewSpanEvent

`func NewSpanEvent(name string, ) *SpanEvent`

NewSpanEvent instantiates a new SpanEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpanEventWithDefaults

`func NewSpanEventWithDefaults() *SpanEvent`

NewSpanEventWithDefaults instantiates a new SpanEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *SpanEvent) GetAttributes() []string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SpanEvent) GetAttributesOk() (*[]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SpanEvent) SetAttributes(v []string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *SpanEvent) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetDroppedAttributesCount

`func (o *SpanEvent) GetDroppedAttributesCount() int32`

GetDroppedAttributesCount returns the DroppedAttributesCount field if non-nil, zero value otherwise.

### GetDroppedAttributesCountOk

`func (o *SpanEvent) GetDroppedAttributesCountOk() (*int32, bool)`

GetDroppedAttributesCountOk returns a tuple with the DroppedAttributesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDroppedAttributesCount

`func (o *SpanEvent) SetDroppedAttributesCount(v int32)`

SetDroppedAttributesCount sets DroppedAttributesCount field to given value.

### HasDroppedAttributesCount

`func (o *SpanEvent) HasDroppedAttributesCount() bool`

HasDroppedAttributesCount returns a boolean if a field has been set.

### GetName

`func (o *SpanEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SpanEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SpanEvent) SetName(v string)`

SetName sets Name field to given value.


### GetTimeUnixNano

`func (o *SpanEvent) GetTimeUnixNano() string`

GetTimeUnixNano returns the TimeUnixNano field if non-nil, zero value otherwise.

### GetTimeUnixNanoOk

`func (o *SpanEvent) GetTimeUnixNanoOk() (*string, bool)`

GetTimeUnixNanoOk returns a tuple with the TimeUnixNano field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUnixNano

`func (o *SpanEvent) SetTimeUnixNano(v string)`

SetTimeUnixNano sets TimeUnixNano field to given value.

### HasTimeUnixNano

`func (o *SpanEvent) HasTimeUnixNano() bool`

HasTimeUnixNano returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


