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

// POSTExternalTaxCalculators201Response struct for POSTExternalTaxCalculators201Response
type POSTExternalTaxCalculators201Response struct {
	Data *POSTExternalTaxCalculators201ResponseData `json:"data,omitempty"`
}

// NewPOSTExternalTaxCalculators201Response instantiates a new POSTExternalTaxCalculators201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTExternalTaxCalculators201Response() *POSTExternalTaxCalculators201Response {
	this := POSTExternalTaxCalculators201Response{}
	return &this
}

// NewPOSTExternalTaxCalculators201ResponseWithDefaults instantiates a new POSTExternalTaxCalculators201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTExternalTaxCalculators201ResponseWithDefaults() *POSTExternalTaxCalculators201Response {
	this := POSTExternalTaxCalculators201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTExternalTaxCalculators201Response) GetData() POSTExternalTaxCalculators201ResponseData {
	if o == nil || o.Data == nil {
		var ret POSTExternalTaxCalculators201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTExternalTaxCalculators201Response) GetDataOk() (*POSTExternalTaxCalculators201ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTExternalTaxCalculators201Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTExternalTaxCalculators201ResponseData and assigns it to the Data field.
func (o *POSTExternalTaxCalculators201Response) SetData(v POSTExternalTaxCalculators201ResponseData) {
	o.Data = &v
}

func (o POSTExternalTaxCalculators201Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePOSTExternalTaxCalculators201Response struct {
	value *POSTExternalTaxCalculators201Response
	isSet bool
}

func (v NullablePOSTExternalTaxCalculators201Response) Get() *POSTExternalTaxCalculators201Response {
	return v.value
}

func (v *NullablePOSTExternalTaxCalculators201Response) Set(val *POSTExternalTaxCalculators201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTExternalTaxCalculators201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTExternalTaxCalculators201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTExternalTaxCalculators201Response(val *POSTExternalTaxCalculators201Response) *NullablePOSTExternalTaxCalculators201Response {
	return &NullablePOSTExternalTaxCalculators201Response{value: val, isSet: true}
}

func (v NullablePOSTExternalTaxCalculators201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTExternalTaxCalculators201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
