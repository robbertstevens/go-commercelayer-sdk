/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AttachmentUpdate struct for AttachmentUpdate
type AttachmentUpdate struct {
	Data AttachmentUpdateData `json:"data"`
}

// NewAttachmentUpdate instantiates a new AttachmentUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttachmentUpdate(data AttachmentUpdateData) *AttachmentUpdate {
	this := AttachmentUpdate{}
	this.Data = data
	return &this
}

// NewAttachmentUpdateWithDefaults instantiates a new AttachmentUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttachmentUpdateWithDefaults() *AttachmentUpdate {
	this := AttachmentUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *AttachmentUpdate) GetData() AttachmentUpdateData {
	if o == nil {
		var ret AttachmentUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AttachmentUpdate) GetDataOk() (*AttachmentUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AttachmentUpdate) SetData(v AttachmentUpdateData) {
	o.Data = v
}

func (o AttachmentUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableAttachmentUpdate struct {
	value *AttachmentUpdate
	isSet bool
}

func (v NullableAttachmentUpdate) Get() *AttachmentUpdate {
	return v.value
}

func (v *NullableAttachmentUpdate) Set(val *AttachmentUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableAttachmentUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableAttachmentUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttachmentUpdate(val *AttachmentUpdate) *NullableAttachmentUpdate {
	return &NullableAttachmentUpdate{value: val, isSet: true}
}

func (v NullableAttachmentUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttachmentUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
