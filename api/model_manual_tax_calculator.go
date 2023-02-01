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

// ManualTaxCalculator struct for ManualTaxCalculator
type ManualTaxCalculator struct {
	Data *ManualTaxCalculatorData `json:"data,omitempty"`
}

// NewManualTaxCalculator instantiates a new ManualTaxCalculator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManualTaxCalculator() *ManualTaxCalculator {
	this := ManualTaxCalculator{}
	return &this
}

// NewManualTaxCalculatorWithDefaults instantiates a new ManualTaxCalculator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManualTaxCalculatorWithDefaults() *ManualTaxCalculator {
	this := ManualTaxCalculator{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ManualTaxCalculator) GetData() ManualTaxCalculatorData {
	if o == nil || o.Data == nil {
		var ret ManualTaxCalculatorData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualTaxCalculator) GetDataOk() (*ManualTaxCalculatorData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ManualTaxCalculator) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given ManualTaxCalculatorData and assigns it to the Data field.
func (o *ManualTaxCalculator) SetData(v ManualTaxCalculatorData) {
	o.Data = &v
}

func (o ManualTaxCalculator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableManualTaxCalculator struct {
	value *ManualTaxCalculator
	isSet bool
}

func (v NullableManualTaxCalculator) Get() *ManualTaxCalculator {
	return v.value
}

func (v *NullableManualTaxCalculator) Set(val *ManualTaxCalculator) {
	v.value = val
	v.isSet = true
}

func (v NullableManualTaxCalculator) IsSet() bool {
	return v.isSet
}

func (v *NullableManualTaxCalculator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManualTaxCalculator(val *ManualTaxCalculator) *NullableManualTaxCalculator {
	return &NullableManualTaxCalculator{value: val, isSet: true}
}

func (v NullableManualTaxCalculator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManualTaxCalculator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
