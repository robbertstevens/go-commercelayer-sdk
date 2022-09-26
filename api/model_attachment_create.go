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

// AttachmentCreate struct for AttachmentCreate
type AttachmentCreate struct {
	Data AttachmentCreateData `json:"data"`
}

// NewAttachmentCreate instantiates a new AttachmentCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttachmentCreate(data AttachmentCreateData) *AttachmentCreate {
	this := AttachmentCreate{}
	this.Data = data
	return &this
}

// NewAttachmentCreateWithDefaults instantiates a new AttachmentCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttachmentCreateWithDefaults() *AttachmentCreate {
	this := AttachmentCreate{}
	return &this
}

// GetData returns the Data field value
func (o *AttachmentCreate) GetData() AttachmentCreateData {
	if o == nil {
		var ret AttachmentCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AttachmentCreate) GetDataOk() (*AttachmentCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AttachmentCreate) SetData(v AttachmentCreateData) {
	o.Data = v
}

func (o AttachmentCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableAttachmentCreate struct {
	value *AttachmentCreate
	isSet bool
}

func (v NullableAttachmentCreate) Get() *AttachmentCreate {
	return v.value
}

func (v *NullableAttachmentCreate) Set(val *AttachmentCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableAttachmentCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableAttachmentCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttachmentCreate(val *AttachmentCreate) *NullableAttachmentCreate {
	return &NullableAttachmentCreate{value: val, isSet: true}
}

func (v NullableAttachmentCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttachmentCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
