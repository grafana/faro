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

// checks if the Stacktrace type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Stacktrace{}

// Stacktrace is a collection of Frames.
type Stacktrace struct {
	Frames []Frame `json:"frames,omitempty"`
}

// NewStacktrace instantiates a new Stacktrace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStacktrace() *Stacktrace {
	this := Stacktrace{}
	return &this
}

// NewStacktraceWithDefaults instantiates a new Stacktrace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStacktraceWithDefaults() *Stacktrace {
	this := Stacktrace{}
	return &this
}

// GetFrames returns the Frames field value if set, zero value otherwise.
func (o *Stacktrace) GetFrames() []Frame {
	if o == nil || IsNil(o.Frames) {
		var ret []Frame
		return ret
	}
	return o.Frames
}

// GetFramesOk returns a tuple with the Frames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Stacktrace) GetFramesOk() ([]Frame, bool) {
	if o == nil || IsNil(o.Frames) {
		return nil, false
	}
	return o.Frames, true
}

// HasFrames returns a boolean if a field has been set.
func (o *Stacktrace) HasFrames() bool {
	if o != nil && !IsNil(o.Frames) {
		return true
	}

	return false
}

// SetFrames gets a reference to the given []Frame and assigns it to the Frames field.
func (o *Stacktrace) SetFrames(v []Frame) {
	o.Frames = v
}

func (o Stacktrace) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Stacktrace) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Frames) {
		toSerialize["frames"] = o.Frames
	}
	return toSerialize, nil
}

type NullableStacktrace struct {
	value *Stacktrace
	isSet bool
}

func (v NullableStacktrace) Get() *Stacktrace {
	return v.value
}

func (v *NullableStacktrace) Set(val *Stacktrace) {
	v.value = val
	v.isSet = true
}

func (v NullableStacktrace) IsSet() bool {
	return v.isSet
}

func (v *NullableStacktrace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStacktrace(val *Stacktrace) *NullableStacktrace {
	return &NullableStacktrace{value: val, isSet: true}
}

func (v NullableStacktrace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStacktrace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

