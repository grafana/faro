# StackFrame

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Colno** | Pointer to **int32** |  | [optional] 
**Filename** | Pointer to **string** |  | [optional] 
**Function** | Pointer to **string** |  | [optional] 
**InApp** | Pointer to **bool** |  | [optional] 
**Lineno** | Pointer to **int32** |  | [optional] 

## Methods

### NewStackFrame

`func NewStackFrame() *StackFrame`

NewStackFrame instantiates a new StackFrame object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackFrameWithDefaults

`func NewStackFrameWithDefaults() *StackFrame`

NewStackFrameWithDefaults instantiates a new StackFrame object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColno

`func (o *StackFrame) GetColno() int32`

GetColno returns the Colno field if non-nil, zero value otherwise.

### GetColnoOk

`func (o *StackFrame) GetColnoOk() (*int32, bool)`

GetColnoOk returns a tuple with the Colno field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColno

`func (o *StackFrame) SetColno(v int32)`

SetColno sets Colno field to given value.

### HasColno

`func (o *StackFrame) HasColno() bool`

HasColno returns a boolean if a field has been set.

### GetFilename

`func (o *StackFrame) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *StackFrame) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *StackFrame) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *StackFrame) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetFunction

`func (o *StackFrame) GetFunction() string`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *StackFrame) GetFunctionOk() (*string, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *StackFrame) SetFunction(v string)`

SetFunction sets Function field to given value.

### HasFunction

`func (o *StackFrame) HasFunction() bool`

HasFunction returns a boolean if a field has been set.

### GetInApp

`func (o *StackFrame) GetInApp() bool`

GetInApp returns the InApp field if non-nil, zero value otherwise.

### GetInAppOk

`func (o *StackFrame) GetInAppOk() (*bool, bool)`

GetInAppOk returns a tuple with the InApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInApp

`func (o *StackFrame) SetInApp(v bool)`

SetInApp sets InApp field to given value.

### HasInApp

`func (o *StackFrame) HasInApp() bool`

HasInApp returns a boolean if a field has been set.

### GetLineno

`func (o *StackFrame) GetLineno() int32`

GetLineno returns the Lineno field if non-nil, zero value otherwise.

### GetLinenoOk

`func (o *StackFrame) GetLinenoOk() (*int32, bool)`

GetLinenoOk returns a tuple with the Lineno field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineno

`func (o *StackFrame) SetLineno(v int32)`

SetLineno sets Lineno field to given value.

### HasLineno

`func (o *StackFrame) HasLineno() bool`

HasLineno returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


