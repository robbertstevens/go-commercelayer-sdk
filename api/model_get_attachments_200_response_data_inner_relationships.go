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

// GETAttachments200ResponseDataInnerRelationships struct for GETAttachments200ResponseDataInnerRelationships
type GETAttachments200ResponseDataInnerRelationships struct {
	Attachable *GETAttachments200ResponseDataInnerRelationshipsAttachable `json:"attachable,omitempty"`
}

// NewGETAttachments200ResponseDataInnerRelationships instantiates a new GETAttachments200ResponseDataInnerRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETAttachments200ResponseDataInnerRelationships() *GETAttachments200ResponseDataInnerRelationships {
	this := GETAttachments200ResponseDataInnerRelationships{}
	return &this
}

// NewGETAttachments200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETAttachments200ResponseDataInnerRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETAttachments200ResponseDataInnerRelationshipsWithDefaults() *GETAttachments200ResponseDataInnerRelationships {
	this := GETAttachments200ResponseDataInnerRelationships{}
	return &this
}

// GetAttachable returns the Attachable field value if set, zero value otherwise.
func (o *GETAttachments200ResponseDataInnerRelationships) GetAttachable() GETAttachments200ResponseDataInnerRelationshipsAttachable {
	if o == nil || o.Attachable == nil {
		var ret GETAttachments200ResponseDataInnerRelationshipsAttachable
		return ret
	}
	return *o.Attachable
}

// GetAttachableOk returns a tuple with the Attachable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAttachments200ResponseDataInnerRelationships) GetAttachableOk() (*GETAttachments200ResponseDataInnerRelationshipsAttachable, bool) {
	if o == nil || o.Attachable == nil {
		return nil, false
	}
	return o.Attachable, true
}

// HasAttachable returns a boolean if a field has been set.
func (o *GETAttachments200ResponseDataInnerRelationships) HasAttachable() bool {
	if o != nil && o.Attachable != nil {
		return true
	}

	return false
}

// SetAttachable gets a reference to the given GETAttachments200ResponseDataInnerRelationshipsAttachable and assigns it to the Attachable field.
func (o *GETAttachments200ResponseDataInnerRelationships) SetAttachable(v GETAttachments200ResponseDataInnerRelationshipsAttachable) {
	o.Attachable = &v
}

func (o GETAttachments200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Attachable != nil {
		toSerialize["attachable"] = o.Attachable
	}
	return json.Marshal(toSerialize)
}

type NullableGETAttachments200ResponseDataInnerRelationships struct {
	value *GETAttachments200ResponseDataInnerRelationships
	isSet bool
}

func (v NullableGETAttachments200ResponseDataInnerRelationships) Get() *GETAttachments200ResponseDataInnerRelationships {
	return v.value
}

func (v *NullableGETAttachments200ResponseDataInnerRelationships) Set(val *GETAttachments200ResponseDataInnerRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETAttachments200ResponseDataInnerRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETAttachments200ResponseDataInnerRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETAttachments200ResponseDataInnerRelationships(val *GETAttachments200ResponseDataInnerRelationships) *NullableGETAttachments200ResponseDataInnerRelationships {
	return &NullableGETAttachments200ResponseDataInnerRelationships{value: val, isSet: true}
}

func (v NullableGETAttachments200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETAttachments200ResponseDataInnerRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
