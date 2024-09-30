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

// checks if the SessionMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SessionMeta{}

// SessionMeta struct for SessionMeta
type SessionMeta struct {
	Id         *string           `json:"id,omitempty"`
	Attributes map[string]string `json:"attributes,omitempty"`
}

// NewSessionMeta instantiates a new SessionMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionMeta() *SessionMeta {
	this := SessionMeta{}
	return &this
}

// NewSessionMetaWithDefaults instantiates a new SessionMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionMetaWithDefaults() *SessionMeta {
	this := SessionMeta{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SessionMeta) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionMeta) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SessionMeta) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SessionMeta) SetId(v string) {
	o.Id = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *SessionMeta) GetAttributes() map[string]string {
	if o == nil || IsNil(o.Attributes) {
		var ret map[string]string
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionMeta) GetAttributesOk() (map[string]string, bool) {
	if o == nil || IsNil(o.Attributes) {
		return map[string]string{}, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *SessionMeta) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]string and assigns it to the Attributes field.
func (o *SessionMeta) SetAttributes(v map[string]string) {
	o.Attributes = v
}

func (o SessionMeta) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SessionMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableSessionMeta struct {
	value *SessionMeta
	isSet bool
}

func (v NullableSessionMeta) Get() *SessionMeta {
	return v.value
}

func (v *NullableSessionMeta) Set(val *SessionMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionMeta(val *SessionMeta) *NullableSessionMeta {
	return &NullableSessionMeta{value: val, isSet: true}
}

func (v NullableSessionMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}