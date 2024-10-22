/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.3.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the PaymentOptionCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentOptionCreate{}

// PaymentOptionCreate struct for PaymentOptionCreate
type PaymentOptionCreate struct {
	Data PaymentOptionCreateData `json:"data"`
}

// NewPaymentOptionCreate instantiates a new PaymentOptionCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentOptionCreate(data PaymentOptionCreateData) *PaymentOptionCreate {
	this := PaymentOptionCreate{}
	this.Data = data
	return &this
}

// NewPaymentOptionCreateWithDefaults instantiates a new PaymentOptionCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentOptionCreateWithDefaults() *PaymentOptionCreate {
	this := PaymentOptionCreate{}
	return &this
}

// GetData returns the Data field value
func (o *PaymentOptionCreate) GetData() PaymentOptionCreateData {
	if o == nil {
		var ret PaymentOptionCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PaymentOptionCreate) GetDataOk() (*PaymentOptionCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PaymentOptionCreate) SetData(v PaymentOptionCreateData) {
	o.Data = v
}

func (o PaymentOptionCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentOptionCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullablePaymentOptionCreate struct {
	value *PaymentOptionCreate
	isSet bool
}

func (v NullablePaymentOptionCreate) Get() *PaymentOptionCreate {
	return v.value
}

func (v *NullablePaymentOptionCreate) Set(val *PaymentOptionCreate) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentOptionCreate) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentOptionCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentOptionCreate(val *PaymentOptionCreate) *NullablePaymentOptionCreate {
	return &NullablePaymentOptionCreate{value: val, isSet: true}
}

func (v NullablePaymentOptionCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentOptionCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
