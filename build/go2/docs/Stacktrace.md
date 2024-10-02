# Stacktrace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Frames** | Pointer to [**[]Frame**](Frame.md) |  | [optional] 

## Methods

### NewStacktrace

`func NewStacktrace() *Stacktrace`

NewStacktrace instantiates a new Stacktrace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStacktraceWithDefaults

`func NewStacktraceWithDefaults() *Stacktrace`

NewStacktraceWithDefaults instantiates a new Stacktrace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrames

`func (o *Stacktrace) GetFrames() []Frame`

GetFrames returns the Frames field if non-nil, zero value otherwise.

### GetFramesOk

`func (o *Stacktrace) GetFramesOk() (*[]Frame, bool)`

GetFramesOk returns a tuple with the Frames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrames

`func (o *Stacktrace) SetFrames(v []Frame)`

SetFrames sets Frames field to given value.

### HasFrames

`func (o *Stacktrace) HasFrames() bool`

HasFrames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


