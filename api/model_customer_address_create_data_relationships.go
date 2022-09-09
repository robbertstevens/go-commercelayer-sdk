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

// CustomerAddressCreateDataRelationships struct for CustomerAddressCreateDataRelationships
type CustomerAddressCreateDataRelationships struct {
	Customer CouponRecipientDataRelationshipsCustomer `json:"customer"`
	Address BingGeocoderDataRelationshipsAddresses `json:"address"`
}

// NewCustomerAddressCreateDataRelationships instantiates a new CustomerAddressCreateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerAddressCreateDataRelationships(customer CouponRecipientDataRelationshipsCustomer, address BingGeocoderDataRelationshipsAddresses) *CustomerAddressCreateDataRelationships {
	this := CustomerAddressCreateDataRelationships{}
	this.Customer = customer
	this.Address = address
	return &this
}

// NewCustomerAddressCreateDataRelationshipsWithDefaults instantiates a new CustomerAddressCreateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerAddressCreateDataRelationshipsWithDefaults() *CustomerAddressCreateDataRelationships {
	this := CustomerAddressCreateDataRelationships{}
	return &this
}

// GetCustomer returns the Customer field value
func (o *CustomerAddressCreateDataRelationships) GetCustomer() CouponRecipientDataRelationshipsCustomer {
	if o == nil {
		var ret CouponRecipientDataRelationshipsCustomer
		return ret
	}

	return o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value
// and a boolean to check if the value has been set.
func (o *CustomerAddressCreateDataRelationships) GetCustomerOk() (*CouponRecipientDataRelationshipsCustomer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Customer, true
}

// SetCustomer sets field value
func (o *CustomerAddressCreateDataRelationships) SetCustomer(v CouponRecipientDataRelationshipsCustomer) {
	o.Customer = v
}

// GetAddress returns the Address field value
func (o *CustomerAddressCreateDataRelationships) GetAddress() BingGeocoderDataRelationshipsAddresses {
	if o == nil {
		var ret BingGeocoderDataRelationshipsAddresses
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *CustomerAddressCreateDataRelationships) GetAddressOk() (*BingGeocoderDataRelationshipsAddresses, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *CustomerAddressCreateDataRelationships) SetAddress(v BingGeocoderDataRelationshipsAddresses) {
	o.Address = v
}

func (o CustomerAddressCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["customer"] = o.Customer
	}
	if true {
		toSerialize["address"] = o.Address
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerAddressCreateDataRelationships struct {
	value *CustomerAddressCreateDataRelationships
	isSet bool
}

func (v NullableCustomerAddressCreateDataRelationships) Get() *CustomerAddressCreateDataRelationships {
	return v.value
}

func (v *NullableCustomerAddressCreateDataRelationships) Set(val *CustomerAddressCreateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerAddressCreateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerAddressCreateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerAddressCreateDataRelationships(val *CustomerAddressCreateDataRelationships) *NullableCustomerAddressCreateDataRelationships {
	return &NullableCustomerAddressCreateDataRelationships{value: val, isSet: true}
}

func (v NullableCustomerAddressCreateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerAddressCreateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


