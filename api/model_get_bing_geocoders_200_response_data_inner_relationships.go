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

// GETBingGeocoders200ResponseDataInnerRelationships struct for GETBingGeocoders200ResponseDataInnerRelationships
type GETBingGeocoders200ResponseDataInnerRelationships struct {
	Addresses *GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods `json:"addresses,omitempty"`
	Attachments *GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods `json:"attachments,omitempty"`
}

// NewGETBingGeocoders200ResponseDataInnerRelationships instantiates a new GETBingGeocoders200ResponseDataInnerRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETBingGeocoders200ResponseDataInnerRelationships() *GETBingGeocoders200ResponseDataInnerRelationships {
	this := GETBingGeocoders200ResponseDataInnerRelationships{}
	return &this
}

// NewGETBingGeocoders200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETBingGeocoders200ResponseDataInnerRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETBingGeocoders200ResponseDataInnerRelationshipsWithDefaults() *GETBingGeocoders200ResponseDataInnerRelationships {
	this := GETBingGeocoders200ResponseDataInnerRelationships{}
	return &this
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *GETBingGeocoders200ResponseDataInnerRelationships) GetAddresses() GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods {
	if o == nil || o.Addresses == nil {
		var ret GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods
		return ret
	}
	return *o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETBingGeocoders200ResponseDataInnerRelationships) GetAddressesOk() (*GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods, bool) {
	if o == nil || o.Addresses == nil {
		return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *GETBingGeocoders200ResponseDataInnerRelationships) HasAddresses() bool {
	if o != nil && o.Addresses != nil {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods and assigns it to the Addresses field.
func (o *GETBingGeocoders200ResponseDataInnerRelationships) SetAddresses(v GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods) {
	o.Addresses = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *GETBingGeocoders200ResponseDataInnerRelationships) GetAttachments() GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods {
	if o == nil || o.Attachments == nil {
		var ret GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETBingGeocoders200ResponseDataInnerRelationships) GetAttachmentsOk() (*GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *GETBingGeocoders200ResponseDataInnerRelationships) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods and assigns it to the Attachments field.
func (o *GETBingGeocoders200ResponseDataInnerRelationships) SetAttachments(v GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods) {
	o.Attachments = &v
}

func (o GETBingGeocoders200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Addresses != nil {
		toSerialize["addresses"] = o.Addresses
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	return json.Marshal(toSerialize)
}

type NullableGETBingGeocoders200ResponseDataInnerRelationships struct {
	value *GETBingGeocoders200ResponseDataInnerRelationships
	isSet bool
}

func (v NullableGETBingGeocoders200ResponseDataInnerRelationships) Get() *GETBingGeocoders200ResponseDataInnerRelationships {
	return v.value
}

func (v *NullableGETBingGeocoders200ResponseDataInnerRelationships) Set(val *GETBingGeocoders200ResponseDataInnerRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETBingGeocoders200ResponseDataInnerRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETBingGeocoders200ResponseDataInnerRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETBingGeocoders200ResponseDataInnerRelationships(val *GETBingGeocoders200ResponseDataInnerRelationships) *NullableGETBingGeocoders200ResponseDataInnerRelationships {
	return &NullableGETBingGeocoders200ResponseDataInnerRelationships{value: val, isSet: true}
}

func (v NullableGETBingGeocoders200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETBingGeocoders200ResponseDataInnerRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


