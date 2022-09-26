/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// TaxjarAccountCreate struct for TaxjarAccountCreate
type TaxjarAccountCreate struct {
	Data TaxjarAccountCreateData `json:"data"`
}

// NewTaxjarAccountCreate instantiates a new TaxjarAccountCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxjarAccountCreate(data TaxjarAccountCreateData) *TaxjarAccountCreate {
	this := TaxjarAccountCreate{}
	this.Data = data
	return &this
}

// NewTaxjarAccountCreateWithDefaults instantiates a new TaxjarAccountCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxjarAccountCreateWithDefaults() *TaxjarAccountCreate {
	this := TaxjarAccountCreate{}
	return &this
}

// GetData returns the Data field value
func (o *TaxjarAccountCreate) GetData() TaxjarAccountCreateData {
	if o == nil {
		var ret TaxjarAccountCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *TaxjarAccountCreate) GetDataOk() (*TaxjarAccountCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *TaxjarAccountCreate) SetData(v TaxjarAccountCreateData) {
	o.Data = v
}

func (o TaxjarAccountCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableTaxjarAccountCreate struct {
	value *TaxjarAccountCreate
	isSet bool
}

func (v NullableTaxjarAccountCreate) Get() *TaxjarAccountCreate {
	return v.value
}

func (v *NullableTaxjarAccountCreate) Set(val *TaxjarAccountCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxjarAccountCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxjarAccountCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxjarAccountCreate(val *TaxjarAccountCreate) *NullableTaxjarAccountCreate {
	return &NullableTaxjarAccountCreate{value: val, isSet: true}
}

func (v NullableTaxjarAccountCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxjarAccountCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
