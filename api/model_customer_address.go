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

// CustomerAddress struct for CustomerAddress
type CustomerAddress struct {
	Data CustomerAddressData `json:"data"`
}

// NewCustomerAddress instantiates a new CustomerAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerAddress(data CustomerAddressData) *CustomerAddress {
	this := CustomerAddress{}
	this.Data = data
	return &this
}

// NewCustomerAddressWithDefaults instantiates a new CustomerAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerAddressWithDefaults() *CustomerAddress {
	this := CustomerAddress{}
	return &this
}

// GetData returns the Data field value
func (o *CustomerAddress) GetData() CustomerAddressData {
	if o == nil {
		var ret CustomerAddressData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CustomerAddress) GetDataOk() (*CustomerAddressData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CustomerAddress) SetData(v CustomerAddressData) {
	o.Data = v
}

func (o CustomerAddress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerAddress struct {
	value *CustomerAddress
	isSet bool
}

func (v NullableCustomerAddress) Get() *CustomerAddress {
	return v.value
}

func (v *NullableCustomerAddress) Set(val *CustomerAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerAddress(val *CustomerAddress) *NullableCustomerAddress {
	return &NullableCustomerAddress{value: val, isSet: true}
}

func (v NullableCustomerAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


