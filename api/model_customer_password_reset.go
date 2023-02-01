/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.4.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// CustomerPasswordReset struct for CustomerPasswordReset
type CustomerPasswordReset struct {
	Data *CustomerPasswordResetData `json:"data,omitempty"`
}

// NewCustomerPasswordReset instantiates a new CustomerPasswordReset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerPasswordReset() *CustomerPasswordReset {
	this := CustomerPasswordReset{}
	return &this
}

// NewCustomerPasswordResetWithDefaults instantiates a new CustomerPasswordReset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerPasswordResetWithDefaults() *CustomerPasswordReset {
	this := CustomerPasswordReset{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CustomerPasswordReset) GetData() CustomerPasswordResetData {
	if o == nil || o.Data == nil {
		var ret CustomerPasswordResetData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerPasswordReset) GetDataOk() (*CustomerPasswordResetData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CustomerPasswordReset) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given CustomerPasswordResetData and assigns it to the Data field.
func (o *CustomerPasswordReset) SetData(v CustomerPasswordResetData) {
	o.Data = &v
}

func (o CustomerPasswordReset) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerPasswordReset struct {
	value *CustomerPasswordReset
	isSet bool
}

func (v NullableCustomerPasswordReset) Get() *CustomerPasswordReset {
	return v.value
}

func (v *NullableCustomerPasswordReset) Set(val *CustomerPasswordReset) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerPasswordReset) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerPasswordReset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerPasswordReset(val *CustomerPasswordReset) *NullableCustomerPasswordReset {
	return &NullableCustomerPasswordReset{value: val, isSet: true}
}

func (v NullableCustomerPasswordReset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerPasswordReset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
