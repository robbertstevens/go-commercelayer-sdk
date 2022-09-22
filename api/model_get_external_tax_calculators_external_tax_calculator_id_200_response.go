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

// GETExternalTaxCalculatorsExternalTaxCalculatorId200Response struct for GETExternalTaxCalculatorsExternalTaxCalculatorId200Response
type GETExternalTaxCalculatorsExternalTaxCalculatorId200Response struct {
	Data *GETExternalTaxCalculators200ResponseDataInner `json:"data,omitempty"`
}

// NewGETExternalTaxCalculatorsExternalTaxCalculatorId200Response instantiates a new GETExternalTaxCalculatorsExternalTaxCalculatorId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETExternalTaxCalculatorsExternalTaxCalculatorId200Response() *GETExternalTaxCalculatorsExternalTaxCalculatorId200Response {
	this := GETExternalTaxCalculatorsExternalTaxCalculatorId200Response{}
	return &this
}

// NewGETExternalTaxCalculatorsExternalTaxCalculatorId200ResponseWithDefaults instantiates a new GETExternalTaxCalculatorsExternalTaxCalculatorId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETExternalTaxCalculatorsExternalTaxCalculatorId200ResponseWithDefaults() *GETExternalTaxCalculatorsExternalTaxCalculatorId200Response {
	this := GETExternalTaxCalculatorsExternalTaxCalculatorId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETExternalTaxCalculatorsExternalTaxCalculatorId200Response) GetData() GETExternalTaxCalculators200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret GETExternalTaxCalculators200ResponseDataInner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETExternalTaxCalculatorsExternalTaxCalculatorId200Response) GetDataOk() (*GETExternalTaxCalculators200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETExternalTaxCalculatorsExternalTaxCalculatorId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETExternalTaxCalculators200ResponseDataInner and assigns it to the Data field.
func (o *GETExternalTaxCalculatorsExternalTaxCalculatorId200Response) SetData(v GETExternalTaxCalculators200ResponseDataInner) {
	o.Data = &v
}

func (o GETExternalTaxCalculatorsExternalTaxCalculatorId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETExternalTaxCalculatorsExternalTaxCalculatorId200Response struct {
	value *GETExternalTaxCalculatorsExternalTaxCalculatorId200Response
	isSet bool
}

func (v NullableGETExternalTaxCalculatorsExternalTaxCalculatorId200Response) Get() *GETExternalTaxCalculatorsExternalTaxCalculatorId200Response {
	return v.value
}

func (v *NullableGETExternalTaxCalculatorsExternalTaxCalculatorId200Response) Set(val *GETExternalTaxCalculatorsExternalTaxCalculatorId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETExternalTaxCalculatorsExternalTaxCalculatorId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETExternalTaxCalculatorsExternalTaxCalculatorId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETExternalTaxCalculatorsExternalTaxCalculatorId200Response(val *GETExternalTaxCalculatorsExternalTaxCalculatorId200Response) *NullableGETExternalTaxCalculatorsExternalTaxCalculatorId200Response {
	return &NullableGETExternalTaxCalculatorsExternalTaxCalculatorId200Response{value: val, isSet: true}
}

func (v NullableGETExternalTaxCalculatorsExternalTaxCalculatorId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETExternalTaxCalculatorsExternalTaxCalculatorId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
