/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AttachmentDataRelationships struct for AttachmentDataRelationships
type AttachmentDataRelationships struct {
	Attachable *AttachmentDataRelationshipsAttachable `json:"attachable,omitempty"`
}

// NewAttachmentDataRelationships instantiates a new AttachmentDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttachmentDataRelationships() *AttachmentDataRelationships {
	this := AttachmentDataRelationships{}
	return &this
}

// NewAttachmentDataRelationshipsWithDefaults instantiates a new AttachmentDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttachmentDataRelationshipsWithDefaults() *AttachmentDataRelationships {
	this := AttachmentDataRelationships{}
	return &this
}

// GetAttachable returns the Attachable field value if set, zero value otherwise.
func (o *AttachmentDataRelationships) GetAttachable() AttachmentDataRelationshipsAttachable {
	if o == nil || o.Attachable == nil {
		var ret AttachmentDataRelationshipsAttachable
		return ret
	}
	return *o.Attachable
}

// GetAttachableOk returns a tuple with the Attachable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentDataRelationships) GetAttachableOk() (*AttachmentDataRelationshipsAttachable, bool) {
	if o == nil || o.Attachable == nil {
		return nil, false
	}
	return o.Attachable, true
}

// HasAttachable returns a boolean if a field has been set.
func (o *AttachmentDataRelationships) HasAttachable() bool {
	if o != nil && o.Attachable != nil {
		return true
	}

	return false
}

// SetAttachable gets a reference to the given AttachmentDataRelationshipsAttachable and assigns it to the Attachable field.
func (o *AttachmentDataRelationships) SetAttachable(v AttachmentDataRelationshipsAttachable) {
	o.Attachable = &v
}

func (o AttachmentDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Attachable != nil {
		toSerialize["attachable"] = o.Attachable
	}
	return json.Marshal(toSerialize)
}

type NullableAttachmentDataRelationships struct {
	value *AttachmentDataRelationships
	isSet bool
}

func (v NullableAttachmentDataRelationships) Get() *AttachmentDataRelationships {
	return v.value
}

func (v *NullableAttachmentDataRelationships) Set(val *AttachmentDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAttachmentDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAttachmentDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttachmentDataRelationships(val *AttachmentDataRelationships) *NullableAttachmentDataRelationships {
	return &NullableAttachmentDataRelationships{value: val, isSet: true}
}

func (v NullableAttachmentDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttachmentDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
