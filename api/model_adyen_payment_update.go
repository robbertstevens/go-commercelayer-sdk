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

// checks if the AdyenPaymentUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdyenPaymentUpdate{}

// AdyenPaymentUpdate struct for AdyenPaymentUpdate
type AdyenPaymentUpdate struct {
	Data AdyenPaymentUpdateData `json:"data"`
}

// NewAdyenPaymentUpdate instantiates a new AdyenPaymentUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdyenPaymentUpdate(data AdyenPaymentUpdateData) *AdyenPaymentUpdate {
	this := AdyenPaymentUpdate{}
	this.Data = data
	return &this
}

// NewAdyenPaymentUpdateWithDefaults instantiates a new AdyenPaymentUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdyenPaymentUpdateWithDefaults() *AdyenPaymentUpdate {
	this := AdyenPaymentUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *AdyenPaymentUpdate) GetData() AdyenPaymentUpdateData {
	if o == nil {
		var ret AdyenPaymentUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AdyenPaymentUpdate) GetDataOk() (*AdyenPaymentUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AdyenPaymentUpdate) SetData(v AdyenPaymentUpdateData) {
	o.Data = v
}

func (o AdyenPaymentUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdyenPaymentUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableAdyenPaymentUpdate struct {
	value *AdyenPaymentUpdate
	isSet bool
}

func (v NullableAdyenPaymentUpdate) Get() *AdyenPaymentUpdate {
	return v.value
}

func (v *NullableAdyenPaymentUpdate) Set(val *AdyenPaymentUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableAdyenPaymentUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableAdyenPaymentUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdyenPaymentUpdate(val *AdyenPaymentUpdate) *NullableAdyenPaymentUpdate {
	return &NullableAdyenPaymentUpdate{value: val, isSet: true}
}

func (v NullableAdyenPaymentUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdyenPaymentUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
