/*
Faro API

The Faro API accepts payloads from the Faro SDKs.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package faro

import (
	"encoding/json"
)

// checks if the BrowserMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BrowserMeta{}

// BrowserMeta struct for BrowserMeta
type BrowserMeta struct {
	Name    *string `json:"name,omitempty"`
	Version *string `json:"version,omitempty"`
	Os      *string `json:"os,omitempty"`
	Mobile  *bool   `json:"mobile,omitempty"`
}

// NewBrowserMeta instantiates a new BrowserMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrowserMeta() *BrowserMeta {
	this := BrowserMeta{}
	return &this
}

// NewBrowserMetaWithDefaults instantiates a new BrowserMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrowserMetaWithDefaults() *BrowserMeta {
	this := BrowserMeta{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BrowserMeta) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrowserMeta) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BrowserMeta) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BrowserMeta) SetName(v string) {
	o.Name = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *BrowserMeta) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrowserMeta) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *BrowserMeta) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *BrowserMeta) SetVersion(v string) {
	o.Version = &v
}

// GetOs returns the Os field value if set, zero value otherwise.
func (o *BrowserMeta) GetOs() string {
	if o == nil || IsNil(o.Os) {
		var ret string
		return ret
	}
	return *o.Os
}

// GetOsOk returns a tuple with the Os field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrowserMeta) GetOsOk() (*string, bool) {
	if o == nil || IsNil(o.Os) {
		return nil, false
	}
	return o.Os, true
}

// HasOs returns a boolean if a field has been set.
func (o *BrowserMeta) HasOs() bool {
	if o != nil && !IsNil(o.Os) {
		return true
	}

	return false
}

// SetOs gets a reference to the given string and assigns it to the Os field.
func (o *BrowserMeta) SetOs(v string) {
	o.Os = &v
}

// GetMobile returns the Mobile field value if set, zero value otherwise.
func (o *BrowserMeta) GetMobile() bool {
	if o == nil || IsNil(o.Mobile) {
		var ret bool
		return ret
	}
	return *o.Mobile
}

// GetMobileOk returns a tuple with the Mobile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrowserMeta) GetMobileOk() (*bool, bool) {
	if o == nil || IsNil(o.Mobile) {
		return nil, false
	}
	return o.Mobile, true
}

// HasMobile returns a boolean if a field has been set.
func (o *BrowserMeta) HasMobile() bool {
	if o != nil && !IsNil(o.Mobile) {
		return true
	}

	return false
}

// SetMobile gets a reference to the given bool and assigns it to the Mobile field.
func (o *BrowserMeta) SetMobile(v bool) {
	o.Mobile = &v
}

func (o BrowserMeta) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BrowserMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.Os) {
		toSerialize["os"] = o.Os
	}
	if !IsNil(o.Mobile) {
		toSerialize["mobile"] = o.Mobile
	}
	return toSerialize, nil
}

type NullableBrowserMeta struct {
	value *BrowserMeta
	isSet bool
}

func (v NullableBrowserMeta) Get() *BrowserMeta {
	return v.value
}

func (v *NullableBrowserMeta) Set(val *BrowserMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableBrowserMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableBrowserMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrowserMeta(val *BrowserMeta) *NullableBrowserMeta {
	return &NullableBrowserMeta{value: val, isSet: true}
}

func (v NullableBrowserMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrowserMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}