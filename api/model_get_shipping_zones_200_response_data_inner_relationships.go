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

// GETShippingZones200ResponseDataInnerRelationships struct for GETShippingZones200ResponseDataInnerRelationships
type GETShippingZones200ResponseDataInnerRelationships struct {
	Attachments *GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments `json:"attachments,omitempty"`
}

// NewGETShippingZones200ResponseDataInnerRelationships instantiates a new GETShippingZones200ResponseDataInnerRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETShippingZones200ResponseDataInnerRelationships() *GETShippingZones200ResponseDataInnerRelationships {
	this := GETShippingZones200ResponseDataInnerRelationships{}
	return &this
}

// NewGETShippingZones200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETShippingZones200ResponseDataInnerRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETShippingZones200ResponseDataInnerRelationshipsWithDefaults() *GETShippingZones200ResponseDataInnerRelationships {
	this := GETShippingZones200ResponseDataInnerRelationships{}
	return &this
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *GETShippingZones200ResponseDataInnerRelationships) GetAttachments() GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments {
	if o == nil || o.Attachments == nil {
		var ret GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShippingZones200ResponseDataInnerRelationships) GetAttachmentsOk() (*GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *GETShippingZones200ResponseDataInnerRelationships) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments and assigns it to the Attachments field.
func (o *GETShippingZones200ResponseDataInnerRelationships) SetAttachments(v GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments) {
	o.Attachments = &v
}

func (o GETShippingZones200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	return json.Marshal(toSerialize)
}

type NullableGETShippingZones200ResponseDataInnerRelationships struct {
	value *GETShippingZones200ResponseDataInnerRelationships
	isSet bool
}

func (v NullableGETShippingZones200ResponseDataInnerRelationships) Get() *GETShippingZones200ResponseDataInnerRelationships {
	return v.value
}

func (v *NullableGETShippingZones200ResponseDataInnerRelationships) Set(val *GETShippingZones200ResponseDataInnerRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETShippingZones200ResponseDataInnerRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETShippingZones200ResponseDataInnerRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETShippingZones200ResponseDataInnerRelationships(val *GETShippingZones200ResponseDataInnerRelationships) *NullableGETShippingZones200ResponseDataInnerRelationships {
	return &NullableGETShippingZones200ResponseDataInnerRelationships{value: val, isSet: true}
}

func (v NullableGETShippingZones200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETShippingZones200ResponseDataInnerRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
