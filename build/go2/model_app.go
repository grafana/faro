/*
Faro API

The Faro API accepts payloads from the Faro SDKs.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the App type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &App{}

// App holds metadata about the application event originates from.
type App struct {
	BundleId *string `json:"bundleId,omitempty"`
	Environment *string `json:"environment,omitempty"`
	Name *string `json:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
	Release *string `json:"release,omitempty"`
	Version *string `json:"version,omitempty"`
}

// NewApp instantiates a new App object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApp() *App {
	this := App{}
	return &this
}

// NewAppWithDefaults instantiates a new App object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppWithDefaults() *App {
	this := App{}
	return &this
}

// GetBundleId returns the BundleId field value if set, zero value otherwise.
func (o *App) GetBundleId() string {
	if o == nil || IsNil(o.BundleId) {
		var ret string
		return ret
	}
	return *o.BundleId
}

// GetBundleIdOk returns a tuple with the BundleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetBundleIdOk() (*string, bool) {
	if o == nil || IsNil(o.BundleId) {
		return nil, false
	}
	return o.BundleId, true
}

// HasBundleId returns a boolean if a field has been set.
func (o *App) HasBundleId() bool {
	if o != nil && !IsNil(o.BundleId) {
		return true
	}

	return false
}

// SetBundleId gets a reference to the given string and assigns it to the BundleId field.
func (o *App) SetBundleId(v string) {
	o.BundleId = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *App) GetEnvironment() string {
	if o == nil || IsNil(o.Environment) {
		var ret string
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetEnvironmentOk() (*string, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *App) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given string and assigns it to the Environment field.
func (o *App) SetEnvironment(v string) {
	o.Environment = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *App) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *App) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *App) SetName(v string) {
	o.Name = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *App) GetNamespace() string {
	if o == nil || IsNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *App) HasNamespace() bool {
	if o != nil && !IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *App) SetNamespace(v string) {
	o.Namespace = &v
}

// GetRelease returns the Release field value if set, zero value otherwise.
func (o *App) GetRelease() string {
	if o == nil || IsNil(o.Release) {
		var ret string
		return ret
	}
	return *o.Release
}

// GetReleaseOk returns a tuple with the Release field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetReleaseOk() (*string, bool) {
	if o == nil || IsNil(o.Release) {
		return nil, false
	}
	return o.Release, true
}

// HasRelease returns a boolean if a field has been set.
func (o *App) HasRelease() bool {
	if o != nil && !IsNil(o.Release) {
		return true
	}

	return false
}

// SetRelease gets a reference to the given string and assigns it to the Release field.
func (o *App) SetRelease(v string) {
	o.Release = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *App) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *App) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *App) SetVersion(v string) {
	o.Version = &v
}

func (o App) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o App) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BundleId) {
		toSerialize["bundleId"] = o.BundleId
	}
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	if !IsNil(o.Release) {
		toSerialize["release"] = o.Release
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableApp struct {
	value *App
	isSet bool
}

func (v NullableApp) Get() *App {
	return v.value
}

func (v *NullableApp) Set(val *App) {
	v.value = val
	v.isSet = true
}

func (v NullableApp) IsSet() bool {
	return v.isSet
}

func (v *NullableApp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApp(val *App) *NullableApp {
	return &NullableApp{value: val, isSet: true}
}

func (v NullableApp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


