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

// POSTTaxjarAccounts201Response struct for POSTTaxjarAccounts201Response
type POSTTaxjarAccounts201Response struct {
	Data *POSTTaxjarAccounts201ResponseData `json:"data,omitempty"`
}

// NewPOSTTaxjarAccounts201Response instantiates a new POSTTaxjarAccounts201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTTaxjarAccounts201Response() *POSTTaxjarAccounts201Response {
	this := POSTTaxjarAccounts201Response{}
	return &this
}

// NewPOSTTaxjarAccounts201ResponseWithDefaults instantiates a new POSTTaxjarAccounts201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTTaxjarAccounts201ResponseWithDefaults() *POSTTaxjarAccounts201Response {
	this := POSTTaxjarAccounts201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTTaxjarAccounts201Response) GetData() POSTTaxjarAccounts201ResponseData {
	if o == nil || o.Data == nil {
		var ret POSTTaxjarAccounts201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTTaxjarAccounts201Response) GetDataOk() (*POSTTaxjarAccounts201ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTTaxjarAccounts201Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTTaxjarAccounts201ResponseData and assigns it to the Data field.
func (o *POSTTaxjarAccounts201Response) SetData(v POSTTaxjarAccounts201ResponseData) {
	o.Data = &v
}

func (o POSTTaxjarAccounts201Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePOSTTaxjarAccounts201Response struct {
	value *POSTTaxjarAccounts201Response
	isSet bool
}

func (v NullablePOSTTaxjarAccounts201Response) Get() *POSTTaxjarAccounts201Response {
	return v.value
}

func (v *NullablePOSTTaxjarAccounts201Response) Set(val *POSTTaxjarAccounts201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTTaxjarAccounts201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTTaxjarAccounts201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTTaxjarAccounts201Response(val *POSTTaxjarAccounts201Response) *NullablePOSTTaxjarAccounts201Response {
	return &NullablePOSTTaxjarAccounts201Response{value: val, isSet: true}
}

func (v NullablePOSTTaxjarAccounts201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTTaxjarAccounts201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
