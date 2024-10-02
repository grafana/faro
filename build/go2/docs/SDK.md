# SDK

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Integrations** | Pointer to [**[]SDKIntegration**](SDKIntegration.md) |  | [optional] 

## Methods

### NewSDK

`func NewSDK() *SDK`

NewSDK instantiates a new SDK object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSDKWithDefaults

`func NewSDKWithDefaults() *SDK`

NewSDKWithDefaults instantiates a new SDK object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SDK) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SDK) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SDK) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SDK) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *SDK) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SDK) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SDK) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SDK) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetIntegrations

`func (o *SDK) GetIntegrations() []SDKIntegration`

GetIntegrations returns the Integrations field if non-nil, zero value otherwise.

### GetIntegrationsOk

`func (o *SDK) GetIntegrationsOk() (*[]SDKIntegration, bool)`

GetIntegrationsOk returns a tuple with the Integrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrations

`func (o *SDK) SetIntegrations(v []SDKIntegration)`

SetIntegrations sets Integrations field to given value.

### HasIntegrations

`func (o *SDK) HasIntegrations() bool`

HasIntegrations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


