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

// GETMerchants200ResponseDataInnerRelationships struct for GETMerchants200ResponseDataInnerRelationships
type GETMerchants200ResponseDataInnerRelationships struct {
	Address *GETAddresses200ResponseDataInnerRelationshipsGeocoder `json:"address,omitempty"`
	Attachments *GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods `json:"attachments,omitempty"`
}

// NewGETMerchants200ResponseDataInnerRelationships instantiates a new GETMerchants200ResponseDataInnerRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETMerchants200ResponseDataInnerRelationships() *GETMerchants200ResponseDataInnerRelationships {
	this := GETMerchants200ResponseDataInnerRelationships{}
	return &this
}

// NewGETMerchants200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETMerchants200ResponseDataInnerRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETMerchants200ResponseDataInnerRelationshipsWithDefaults() *GETMerchants200ResponseDataInnerRelationships {
	this := GETMerchants200ResponseDataInnerRelationships{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *GETMerchants200ResponseDataInnerRelationships) GetAddress() GETAddresses200ResponseDataInnerRelationshipsGeocoder {
	if o == nil || o.Address == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoder
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETMerchants200ResponseDataInnerRelationships) GetAddressOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoder, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *GETMerchants200ResponseDataInnerRelationships) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoder and assigns it to the Address field.
func (o *GETMerchants200ResponseDataInnerRelationships) SetAddress(v GETAddresses200ResponseDataInnerRelationshipsGeocoder) {
	o.Address = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *GETMerchants200ResponseDataInnerRelationships) GetAttachments() GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods {
	if o == nil || o.Attachments == nil {
		var ret GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETMerchants200ResponseDataInnerRelationships) GetAttachmentsOk() (*GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *GETMerchants200ResponseDataInnerRelationships) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods and assigns it to the Attachments field.
func (o *GETMerchants200ResponseDataInnerRelationships) SetAttachments(v GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods) {
	o.Attachments = &v
}

func (o GETMerchants200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	return json.Marshal(toSerialize)
}

type NullableGETMerchants200ResponseDataInnerRelationships struct {
	value *GETMerchants200ResponseDataInnerRelationships
	isSet bool
}

func (v NullableGETMerchants200ResponseDataInnerRelationships) Get() *GETMerchants200ResponseDataInnerRelationships {
	return v.value
}

func (v *NullableGETMerchants200ResponseDataInnerRelationships) Set(val *GETMerchants200ResponseDataInnerRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETMerchants200ResponseDataInnerRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETMerchants200ResponseDataInnerRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETMerchants200ResponseDataInnerRelationships(val *GETMerchants200ResponseDataInnerRelationships) *NullableGETMerchants200ResponseDataInnerRelationships {
	return &NullableGETMerchants200ResponseDataInnerRelationships{value: val, isSet: true}
}

func (v NullableGETMerchants200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETMerchants200ResponseDataInnerRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


