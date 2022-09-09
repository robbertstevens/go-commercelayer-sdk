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

// POSTCheckoutComPayments201Response struct for POSTCheckoutComPayments201Response
type POSTCheckoutComPayments201Response struct {
	Data *POSTCheckoutComPayments201ResponseData `json:"data,omitempty"`
}

// NewPOSTCheckoutComPayments201Response instantiates a new POSTCheckoutComPayments201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTCheckoutComPayments201Response() *POSTCheckoutComPayments201Response {
	this := POSTCheckoutComPayments201Response{}
	return &this
}

// NewPOSTCheckoutComPayments201ResponseWithDefaults instantiates a new POSTCheckoutComPayments201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTCheckoutComPayments201ResponseWithDefaults() *POSTCheckoutComPayments201Response {
	this := POSTCheckoutComPayments201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTCheckoutComPayments201Response) GetData() POSTCheckoutComPayments201ResponseData {
	if o == nil || o.Data == nil {
		var ret POSTCheckoutComPayments201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCheckoutComPayments201Response) GetDataOk() (*POSTCheckoutComPayments201ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTCheckoutComPayments201Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTCheckoutComPayments201ResponseData and assigns it to the Data field.
func (o *POSTCheckoutComPayments201Response) SetData(v POSTCheckoutComPayments201ResponseData) {
	o.Data = &v
}

func (o POSTCheckoutComPayments201Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePOSTCheckoutComPayments201Response struct {
	value *POSTCheckoutComPayments201Response
	isSet bool
}

func (v NullablePOSTCheckoutComPayments201Response) Get() *POSTCheckoutComPayments201Response {
	return v.value
}

func (v *NullablePOSTCheckoutComPayments201Response) Set(val *POSTCheckoutComPayments201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTCheckoutComPayments201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTCheckoutComPayments201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTCheckoutComPayments201Response(val *POSTCheckoutComPayments201Response) *NullablePOSTCheckoutComPayments201Response {
	return &NullablePOSTCheckoutComPayments201Response{value: val, isSet: true}
}

func (v NullablePOSTCheckoutComPayments201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTCheckoutComPayments201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


