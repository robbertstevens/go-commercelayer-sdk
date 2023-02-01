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

// PATCHManualTaxCalculatorsManualTaxCalculatorId200Response struct for PATCHManualTaxCalculatorsManualTaxCalculatorId200Response
type PATCHManualTaxCalculatorsManualTaxCalculatorId200Response struct {
	Data *PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData `json:"data,omitempty"`
}

// NewPATCHManualTaxCalculatorsManualTaxCalculatorId200Response instantiates a new PATCHManualTaxCalculatorsManualTaxCalculatorId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHManualTaxCalculatorsManualTaxCalculatorId200Response() *PATCHManualTaxCalculatorsManualTaxCalculatorId200Response {
	this := PATCHManualTaxCalculatorsManualTaxCalculatorId200Response{}
	return &this
}

// NewPATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseWithDefaults instantiates a new PATCHManualTaxCalculatorsManualTaxCalculatorId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseWithDefaults() *PATCHManualTaxCalculatorsManualTaxCalculatorId200Response {
	this := PATCHManualTaxCalculatorsManualTaxCalculatorId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorId200Response) GetData() PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData {
	if o == nil || o.Data == nil {
		var ret PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorId200Response) GetDataOk() (*PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData and assigns it to the Data field.
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorId200Response) SetData(v PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData) {
	o.Data = &v
}

func (o PATCHManualTaxCalculatorsManualTaxCalculatorId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHManualTaxCalculatorsManualTaxCalculatorId200Response struct {
	value *PATCHManualTaxCalculatorsManualTaxCalculatorId200Response
	isSet bool
}

func (v NullablePATCHManualTaxCalculatorsManualTaxCalculatorId200Response) Get() *PATCHManualTaxCalculatorsManualTaxCalculatorId200Response {
	return v.value
}

func (v *NullablePATCHManualTaxCalculatorsManualTaxCalculatorId200Response) Set(val *PATCHManualTaxCalculatorsManualTaxCalculatorId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHManualTaxCalculatorsManualTaxCalculatorId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHManualTaxCalculatorsManualTaxCalculatorId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHManualTaxCalculatorsManualTaxCalculatorId200Response(val *PATCHManualTaxCalculatorsManualTaxCalculatorId200Response) *NullablePATCHManualTaxCalculatorsManualTaxCalculatorId200Response {
	return &NullablePATCHManualTaxCalculatorsManualTaxCalculatorId200Response{value: val, isSet: true}
}

func (v NullablePATCHManualTaxCalculatorsManualTaxCalculatorId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHManualTaxCalculatorsManualTaxCalculatorId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
