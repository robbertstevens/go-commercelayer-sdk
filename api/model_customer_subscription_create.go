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

// CustomerSubscriptionCreate struct for CustomerSubscriptionCreate
type CustomerSubscriptionCreate struct {
	Data CustomerSubscriptionCreateData `json:"data"`
}

// NewCustomerSubscriptionCreate instantiates a new CustomerSubscriptionCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerSubscriptionCreate(data CustomerSubscriptionCreateData) *CustomerSubscriptionCreate {
	this := CustomerSubscriptionCreate{}
	this.Data = data
	return &this
}

// NewCustomerSubscriptionCreateWithDefaults instantiates a new CustomerSubscriptionCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerSubscriptionCreateWithDefaults() *CustomerSubscriptionCreate {
	this := CustomerSubscriptionCreate{}
	return &this
}

// GetData returns the Data field value
func (o *CustomerSubscriptionCreate) GetData() CustomerSubscriptionCreateData {
	if o == nil {
		var ret CustomerSubscriptionCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CustomerSubscriptionCreate) GetDataOk() (*CustomerSubscriptionCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CustomerSubscriptionCreate) SetData(v CustomerSubscriptionCreateData) {
	o.Data = v
}

func (o CustomerSubscriptionCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerSubscriptionCreate struct {
	value *CustomerSubscriptionCreate
	isSet bool
}

func (v NullableCustomerSubscriptionCreate) Get() *CustomerSubscriptionCreate {
	return v.value
}

func (v *NullableCustomerSubscriptionCreate) Set(val *CustomerSubscriptionCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerSubscriptionCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerSubscriptionCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerSubscriptionCreate(val *CustomerSubscriptionCreate) *NullableCustomerSubscriptionCreate {
	return &NullableCustomerSubscriptionCreate{value: val, isSet: true}
}

func (v NullableCustomerSubscriptionCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerSubscriptionCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


