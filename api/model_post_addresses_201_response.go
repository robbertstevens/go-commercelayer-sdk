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

// POSTAddresses201Response struct for POSTAddresses201Response
type POSTAddresses201Response struct {
	Data *POSTAddresses201ResponseData `json:"data,omitempty"`
}

// NewPOSTAddresses201Response instantiates a new POSTAddresses201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTAddresses201Response() *POSTAddresses201Response {
	this := POSTAddresses201Response{}
	return &this
}

// NewPOSTAddresses201ResponseWithDefaults instantiates a new POSTAddresses201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTAddresses201ResponseWithDefaults() *POSTAddresses201Response {
	this := POSTAddresses201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTAddresses201Response) GetData() POSTAddresses201ResponseData {
	if o == nil || o.Data == nil {
		var ret POSTAddresses201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTAddresses201Response) GetDataOk() (*POSTAddresses201ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTAddresses201Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTAddresses201ResponseData and assigns it to the Data field.
func (o *POSTAddresses201Response) SetData(v POSTAddresses201ResponseData) {
	o.Data = &v
}

func (o POSTAddresses201Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePOSTAddresses201Response struct {
	value *POSTAddresses201Response
	isSet bool
}

func (v NullablePOSTAddresses201Response) Get() *POSTAddresses201Response {
	return v.value
}

func (v *NullablePOSTAddresses201Response) Set(val *POSTAddresses201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTAddresses201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTAddresses201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTAddresses201Response(val *POSTAddresses201Response) *NullablePOSTAddresses201Response {
	return &NullablePOSTAddresses201Response{value: val, isSet: true}
}

func (v NullablePOSTAddresses201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTAddresses201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


