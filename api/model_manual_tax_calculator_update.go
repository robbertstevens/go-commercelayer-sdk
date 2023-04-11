/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 4.1.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ManualTaxCalculatorUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManualTaxCalculatorUpdate{}

// ManualTaxCalculatorUpdate struct for ManualTaxCalculatorUpdate
type ManualTaxCalculatorUpdate struct {
	Data PATCHManualTaxCalculatorsManualTaxCalculatorIdRequestData `json:"data"`
}

// NewManualTaxCalculatorUpdate instantiates a new ManualTaxCalculatorUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManualTaxCalculatorUpdate(data PATCHManualTaxCalculatorsManualTaxCalculatorIdRequestData) *ManualTaxCalculatorUpdate {
	this := ManualTaxCalculatorUpdate{}
	this.Data = data
	return &this
}

// NewManualTaxCalculatorUpdateWithDefaults instantiates a new ManualTaxCalculatorUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManualTaxCalculatorUpdateWithDefaults() *ManualTaxCalculatorUpdate {
	this := ManualTaxCalculatorUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *ManualTaxCalculatorUpdate) GetData() PATCHManualTaxCalculatorsManualTaxCalculatorIdRequestData {
	if o == nil {
		var ret PATCHManualTaxCalculatorsManualTaxCalculatorIdRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ManualTaxCalculatorUpdate) GetDataOk() (*PATCHManualTaxCalculatorsManualTaxCalculatorIdRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ManualTaxCalculatorUpdate) SetData(v PATCHManualTaxCalculatorsManualTaxCalculatorIdRequestData) {
	o.Data = v
}

func (o ManualTaxCalculatorUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManualTaxCalculatorUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableManualTaxCalculatorUpdate struct {
	value *ManualTaxCalculatorUpdate
	isSet bool
}

func (v NullableManualTaxCalculatorUpdate) Get() *ManualTaxCalculatorUpdate {
	return v.value
}

func (v *NullableManualTaxCalculatorUpdate) Set(val *ManualTaxCalculatorUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableManualTaxCalculatorUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableManualTaxCalculatorUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManualTaxCalculatorUpdate(val *ManualTaxCalculatorUpdate) *NullableManualTaxCalculatorUpdate {
	return &NullableManualTaxCalculatorUpdate{value: val, isSet: true}
}

func (v NullableManualTaxCalculatorUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManualTaxCalculatorUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
