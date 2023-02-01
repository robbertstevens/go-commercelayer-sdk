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

// PaymentMethodCreate struct for PaymentMethodCreate
type PaymentMethodCreate struct {
	Data PaymentMethodCreateData `json:"data"`
}

// NewPaymentMethodCreate instantiates a new PaymentMethodCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodCreate(data PaymentMethodCreateData) *PaymentMethodCreate {
	this := PaymentMethodCreate{}
	this.Data = data
	return &this
}

// NewPaymentMethodCreateWithDefaults instantiates a new PaymentMethodCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodCreateWithDefaults() *PaymentMethodCreate {
	this := PaymentMethodCreate{}
	return &this
}

// GetData returns the Data field value
func (o *PaymentMethodCreate) GetData() PaymentMethodCreateData {
	if o == nil {
		var ret PaymentMethodCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodCreate) GetDataOk() (*PaymentMethodCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PaymentMethodCreate) SetData(v PaymentMethodCreateData) {
	o.Data = v
}

func (o PaymentMethodCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentMethodCreate struct {
	value *PaymentMethodCreate
	isSet bool
}

func (v NullablePaymentMethodCreate) Get() *PaymentMethodCreate {
	return v.value
}

func (v *NullablePaymentMethodCreate) Set(val *PaymentMethodCreate) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodCreate) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodCreate(val *PaymentMethodCreate) *NullablePaymentMethodCreate {
	return &NullablePaymentMethodCreate{value: val, isSet: true}
}

func (v NullablePaymentMethodCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
