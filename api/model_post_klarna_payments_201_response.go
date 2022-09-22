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

// POSTKlarnaPayments201Response struct for POSTKlarnaPayments201Response
type POSTKlarnaPayments201Response struct {
	Data *POSTKlarnaPayments201ResponseData `json:"data,omitempty"`
}

// NewPOSTKlarnaPayments201Response instantiates a new POSTKlarnaPayments201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTKlarnaPayments201Response() *POSTKlarnaPayments201Response {
	this := POSTKlarnaPayments201Response{}
	return &this
}

// NewPOSTKlarnaPayments201ResponseWithDefaults instantiates a new POSTKlarnaPayments201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTKlarnaPayments201ResponseWithDefaults() *POSTKlarnaPayments201Response {
	this := POSTKlarnaPayments201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTKlarnaPayments201Response) GetData() POSTKlarnaPayments201ResponseData {
	if o == nil || o.Data == nil {
		var ret POSTKlarnaPayments201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTKlarnaPayments201Response) GetDataOk() (*POSTKlarnaPayments201ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTKlarnaPayments201Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTKlarnaPayments201ResponseData and assigns it to the Data field.
func (o *POSTKlarnaPayments201Response) SetData(v POSTKlarnaPayments201ResponseData) {
	o.Data = &v
}

func (o POSTKlarnaPayments201Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePOSTKlarnaPayments201Response struct {
	value *POSTKlarnaPayments201Response
	isSet bool
}

func (v NullablePOSTKlarnaPayments201Response) Get() *POSTKlarnaPayments201Response {
	return v.value
}

func (v *NullablePOSTKlarnaPayments201Response) Set(val *POSTKlarnaPayments201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTKlarnaPayments201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTKlarnaPayments201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTKlarnaPayments201Response(val *POSTKlarnaPayments201Response) *NullablePOSTKlarnaPayments201Response {
	return &NullablePOSTKlarnaPayments201Response{value: val, isSet: true}
}

func (v NullablePOSTKlarnaPayments201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTKlarnaPayments201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
