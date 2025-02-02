/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.3.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the PaymentMethodUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentMethodUpdate{}

// PaymentMethodUpdate struct for PaymentMethodUpdate
type PaymentMethodUpdate struct {
	Data PaymentMethodUpdateData `json:"data"`
}

// NewPaymentMethodUpdate instantiates a new PaymentMethodUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodUpdate(data PaymentMethodUpdateData) *PaymentMethodUpdate {
	this := PaymentMethodUpdate{}
	this.Data = data
	return &this
}

// NewPaymentMethodUpdateWithDefaults instantiates a new PaymentMethodUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodUpdateWithDefaults() *PaymentMethodUpdate {
	this := PaymentMethodUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *PaymentMethodUpdate) GetData() PaymentMethodUpdateData {
	if o == nil {
		var ret PaymentMethodUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodUpdate) GetDataOk() (*PaymentMethodUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PaymentMethodUpdate) SetData(v PaymentMethodUpdateData) {
	o.Data = v
}

func (o PaymentMethodUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullablePaymentMethodUpdate struct {
	value *PaymentMethodUpdate
	isSet bool
}

func (v NullablePaymentMethodUpdate) Get() *PaymentMethodUpdate {
	return v.value
}

func (v *NullablePaymentMethodUpdate) Set(val *PaymentMethodUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodUpdate(val *PaymentMethodUpdate) *NullablePaymentMethodUpdate {
	return &NullablePaymentMethodUpdate{value: val, isSet: true}
}

func (v NullablePaymentMethodUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
