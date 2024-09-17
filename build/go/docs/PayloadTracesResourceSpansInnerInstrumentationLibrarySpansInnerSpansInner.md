# PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInnerAttributes**](PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInnerAttributes.md) |  | [optional] 
**DroppedAttributesCount** | Pointer to **int32** |  | [optional] 
**DroppedEventsCount** | Pointer to **int32** |  | [optional] 
**DroppedLinksCount** | Pointer to **int32** |  | [optional] 
**EndTimeUnixNano** | **int32** |  | 
**Events** | Pointer to [**[]SpanEvent**](SpanEvent.md) |  | [optional] 
**Kind** | **int32** |  | 
**Links** | Pointer to **[]map[string]string** |  | [optional] 
**Name** | **string** |  | 
**ParentSpanId** | Pointer to **string** |  | [optional] 
**SpanId** | **string** |  | 
**StartTimeUnixNano** | **int32** |  | 
**Status** | [**PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInnerStatus**](PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInnerStatus.md) |  | 
**TraceId** | **string** |  | 

## Methods

### NewPayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner

`func NewPayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner(endTimeUnixNano int32, kind int32, name string, spanId string, startTimeUnixNano int32, status PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInnerStatus, traceId string, ) *PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner`

NewPayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner instantiates a new PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInnerWithDefaults

`func NewPayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInnerWithDefaults() *PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner`

NewPayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInnerWithDefaults instantiates a new PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner) GetAttributes() PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInnerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner) GetAttributesOk() (*PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInnerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner) SetAttributes(v PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInnerAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetDroppedAttributesCount

`func (o *PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner) GetDroppedAttributesCount() int32`

GetDroppedAttributesCount returns the DroppedAttributesCount field if non-nil, zero value otherwise.

### GetDroppedAttributesCountOk

`func (o *PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner) GetDroppedAttributesCountOk() (*int32, bool)`

GetDroppedAttributesCountOk returns a tuple with the DroppedAttributesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDroppedAttributesCount

`func (o *PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner) SetDroppedAttributesCount(v int32)`

SetDroppedAttributesCount sets DroppedAttributesCount field to given value.

### HasDroppedAttributesCount

`func (o *PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner) HasDroppedAttributesCount() bool`

HasDroppedAttributesCount returns a boolean if a field has been set.

### GetDroppedEventsCount

`func (o *PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner) GetDroppedEventsCount() int32`

GetDroppedEventsCount returns the DroppedEventsCount field if non-nil, zero value otherwise.

### GetDroppedEventsCountOk

`func (o *PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner) GetDroppedEventsCountOk() (*int32, bool)`

GetDroppedEventsCountOk returns a tuple with the DroppedEventsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDroppedEventsCount

`func (o *PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner) SetDroppedEventsCount(v int32)`

SetDroppedEventsCount sets DroppedEventsCount field to given value.

### HasDroppedEventsCount

`func (o *PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner) HasDroppedEventsCount() bool`

HasDroppedEventsCount returns a boolean if a field has been set.

### GetDroppedLinksCount

`func (o *PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner) GetDroppedLinksCount() int32`

GetDroppedLinksCount returns the DroppedLinksCount field if non-nil, zero value otherwise.

### GetDroppedLinksCountOk

`func (o *PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner) GetDroppedLinksCountOk() (*int32, bool)`

GetDroppedLinksCountOk returns a tuple with the DroppedLinksCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDroppedLinksCount

`func (o *PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner) SetDroppedLinksCount(v int32)`

SetDroppedLinksCount sets DroppedLinksCount field to given value.

### HasDroppedLinksCount

`func (o *PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner) HasDroppedLinksCount() bool`

HasDroppedLinksCount returns a boolean if a field has been set.

### GetEndTimeUnixNano

`func (o *PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner) GetEndTimeUnixNano() int32`

GetEndTimeUnixNano returns the EndTimeUnixNano field if non-nil, zero value otherwise.

### GetEndTimeUnixNanoOk

`func (o *PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner) GetEndTimeUnixNanoOk() (*int32, bool)`

GetEndTimeUnixNanoOk returns a tuple with the EndTimeUnixNano field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUnixNano

`func (o *PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner) SetEndTimeUnixNano(v int32)`

SetEndTimeUnixNano sets EndTimeUnixNano field to given value.


### GetEvents

`func (o *PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner) GetEvents() []SpanEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner) GetEventsOk() (*[]SpanEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner) SetEvents(v []SpanEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetKind

`func (o *PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner) GetKind() int32`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner) GetKindOk() (*int32, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner) SetKind(v int32)`

SetKind sets Kind field to given value.


### GetLinks

`func (o *PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner) GetLinks() []map[string]string`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner) GetLinksOk() (*[]map[string]string, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner) SetLinks(v []map[string]string)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetName

`func (o *PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner) SetName(v string)`

SetName sets Name field to given value.


### GetParentSpanId

`func (o *PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner) GetParentSpanId() string`

GetParentSpanId returns the ParentSpanId field if non-nil, zero value otherwise.

### GetParentSpanIdOk

`func (o *PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner) GetParentSpanIdOk() (*string, bool)`

GetParentSpanIdOk returns a tuple with the ParentSpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSpanId

`func (o *PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner) SetParentSpanId(v string)`

SetParentSpanId sets ParentSpanId field to given value.

### HasParentSpanId

`func (o *PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner) HasParentSpanId() bool`

HasParentSpanId returns a boolean if a field has been set.

### GetSpanId

`func (o *PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner) GetSpanId() string`

GetSpanId returns the SpanId field if non-nil, zero value otherwise.

### GetSpanIdOk

`func (o *PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner) GetSpanIdOk() (*string, bool)`

GetSpanIdOk returns a tuple with the SpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanId

`func (o *PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner) SetSpanId(v string)`

SetSpanId sets SpanId field to given value.


### GetStartTimeUnixNano

`func (o *PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner) GetStartTimeUnixNano() int32`

GetStartTimeUnixNano returns the StartTimeUnixNano field if non-nil, zero value otherwise.

### GetStartTimeUnixNanoOk

`func (o *PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner) GetStartTimeUnixNanoOk() (*int32, bool)`

GetStartTimeUnixNanoOk returns a tuple with the StartTimeUnixNano field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUnixNano

`func (o *PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner) SetStartTimeUnixNano(v int32)`

SetStartTimeUnixNano sets StartTimeUnixNano field to given value.


### GetStatus

`func (o *PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner) GetStatus() PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInnerStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner) GetStatusOk() (*PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInnerStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner) SetStatus(v PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInnerStatus)`

SetStatus sets Status field to given value.


### GetTraceId

`func (o *PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInnerSpansInner) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


