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

// POSTBraintreePayments201Response struct for POSTBraintreePayments201Response
type POSTBraintreePayments201Response struct {
	Data *POSTBraintreePayments201ResponseData `json:"data,omitempty"`
}

// NewPOSTBraintreePayments201Response instantiates a new POSTBraintreePayments201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTBraintreePayments201Response() *POSTBraintreePayments201Response {
	this := POSTBraintreePayments201Response{}
	return &this
}

// NewPOSTBraintreePayments201ResponseWithDefaults instantiates a new POSTBraintreePayments201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTBraintreePayments201ResponseWithDefaults() *POSTBraintreePayments201Response {
	this := POSTBraintreePayments201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTBraintreePayments201Response) GetData() POSTBraintreePayments201ResponseData {
	if o == nil || o.Data == nil {
		var ret POSTBraintreePayments201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTBraintreePayments201Response) GetDataOk() (*POSTBraintreePayments201ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTBraintreePayments201Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTBraintreePayments201ResponseData and assigns it to the Data field.
func (o *POSTBraintreePayments201Response) SetData(v POSTBraintreePayments201ResponseData) {
	o.Data = &v
}

func (o POSTBraintreePayments201Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePOSTBraintreePayments201Response struct {
	value *POSTBraintreePayments201Response
	isSet bool
}

func (v NullablePOSTBraintreePayments201Response) Get() *POSTBraintreePayments201Response {
	return v.value
}

func (v *NullablePOSTBraintreePayments201Response) Set(val *POSTBraintreePayments201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTBraintreePayments201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTBraintreePayments201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTBraintreePayments201Response(val *POSTBraintreePayments201Response) *NullablePOSTBraintreePayments201Response {
	return &NullablePOSTBraintreePayments201Response{value: val, isSet: true}
}

func (v NullablePOSTBraintreePayments201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTBraintreePayments201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


