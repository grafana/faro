# Frame

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Colno** | Pointer to **int32** |  | [optional] 
**Filename** | Pointer to **string** |  | [optional] 
**Function** | Pointer to **string** |  | [optional] 
**Lineno** | Pointer to **int32** |  | [optional] 
**Module** | Pointer to **string** |  | [optional] 

## Methods

### NewFrame

`func NewFrame() *Frame`

NewFrame instantiates a new Frame object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFrameWithDefaults

`func NewFrameWithDefaults() *Frame`

NewFrameWithDefaults instantiates a new Frame object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColno

`func (o *Frame) GetColno() int32`

GetColno returns the Colno field if non-nil, zero value otherwise.

### GetColnoOk

`func (o *Frame) GetColnoOk() (*int32, bool)`

GetColnoOk returns a tuple with the Colno field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColno

`func (o *Frame) SetColno(v int32)`

SetColno sets Colno field to given value.

### HasColno

`func (o *Frame) HasColno() bool`

HasColno returns a boolean if a field has been set.

### GetFilename

`func (o *Frame) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *Frame) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *Frame) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *Frame) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetFunction

`func (o *Frame) GetFunction() string`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *Frame) GetFunctionOk() (*string, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *Frame) SetFunction(v string)`

SetFunction sets Function field to given value.

### HasFunction

`func (o *Frame) HasFunction() bool`

HasFunction returns a boolean if a field has been set.

### GetLineno

`func (o *Frame) GetLineno() int32`

GetLineno returns the Lineno field if non-nil, zero value otherwise.

### GetLinenoOk

`func (o *Frame) GetLinenoOk() (*int32, bool)`

GetLinenoOk returns a tuple with the Lineno field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineno

`func (o *Frame) SetLineno(v int32)`

SetLineno sets Lineno field to given value.

### HasLineno

`func (o *Frame) HasLineno() bool`

HasLineno returns a boolean if a field has been set.

### GetModule

`func (o *Frame) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *Frame) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *Frame) SetModule(v string)`

SetModule sets Module field to given value.

### HasModule

`func (o *Frame) HasModule() bool`

HasModule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


