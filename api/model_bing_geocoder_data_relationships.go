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

// BingGeocoderDataRelationships struct for BingGeocoderDataRelationships
type BingGeocoderDataRelationships struct {
	Addresses *BingGeocoderDataRelationshipsAddresses `json:"addresses,omitempty"`
	Attachments *AvalaraAccountDataRelationshipsAttachments `json:"attachments,omitempty"`
}

// NewBingGeocoderDataRelationships instantiates a new BingGeocoderDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBingGeocoderDataRelationships() *BingGeocoderDataRelationships {
	this := BingGeocoderDataRelationships{}
	return &this
}

// NewBingGeocoderDataRelationshipsWithDefaults instantiates a new BingGeocoderDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBingGeocoderDataRelationshipsWithDefaults() *BingGeocoderDataRelationships {
	this := BingGeocoderDataRelationships{}
	return &this
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *BingGeocoderDataRelationships) GetAddresses() BingGeocoderDataRelationshipsAddresses {
	if o == nil || o.Addresses == nil {
		var ret BingGeocoderDataRelationshipsAddresses
		return ret
	}
	return *o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BingGeocoderDataRelationships) GetAddressesOk() (*BingGeocoderDataRelationshipsAddresses, bool) {
	if o == nil || o.Addresses == nil {
		return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *BingGeocoderDataRelationships) HasAddresses() bool {
	if o != nil && o.Addresses != nil {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given BingGeocoderDataRelationshipsAddresses and assigns it to the Addresses field.
func (o *BingGeocoderDataRelationships) SetAddresses(v BingGeocoderDataRelationshipsAddresses) {
	o.Addresses = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *BingGeocoderDataRelationships) GetAttachments() AvalaraAccountDataRelationshipsAttachments {
	if o == nil || o.Attachments == nil {
		var ret AvalaraAccountDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BingGeocoderDataRelationships) GetAttachmentsOk() (*AvalaraAccountDataRelationshipsAttachments, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *BingGeocoderDataRelationships) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given AvalaraAccountDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *BingGeocoderDataRelationships) SetAttachments(v AvalaraAccountDataRelationshipsAttachments) {
	o.Attachments = &v
}

func (o BingGeocoderDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Addresses != nil {
		toSerialize["addresses"] = o.Addresses
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	return json.Marshal(toSerialize)
}

type NullableBingGeocoderDataRelationships struct {
	value *BingGeocoderDataRelationships
	isSet bool
}

func (v NullableBingGeocoderDataRelationships) Get() *BingGeocoderDataRelationships {
	return v.value
}

func (v *NullableBingGeocoderDataRelationships) Set(val *BingGeocoderDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableBingGeocoderDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableBingGeocoderDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBingGeocoderDataRelationships(val *BingGeocoderDataRelationships) *NullableBingGeocoderDataRelationships {
	return &NullableBingGeocoderDataRelationships{value: val, isSet: true}
}

func (v NullableBingGeocoderDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBingGeocoderDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


