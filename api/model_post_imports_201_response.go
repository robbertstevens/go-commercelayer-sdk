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

// POSTImports201Response struct for POSTImports201Response
type POSTImports201Response struct {
	Data *POSTImports201ResponseData `json:"data,omitempty"`
}

// NewPOSTImports201Response instantiates a new POSTImports201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTImports201Response() *POSTImports201Response {
	this := POSTImports201Response{}
	return &this
}

// NewPOSTImports201ResponseWithDefaults instantiates a new POSTImports201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTImports201ResponseWithDefaults() *POSTImports201Response {
	this := POSTImports201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTImports201Response) GetData() POSTImports201ResponseData {
	if o == nil || o.Data == nil {
		var ret POSTImports201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTImports201Response) GetDataOk() (*POSTImports201ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTImports201Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTImports201ResponseData and assigns it to the Data field.
func (o *POSTImports201Response) SetData(v POSTImports201ResponseData) {
	o.Data = &v
}

func (o POSTImports201Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePOSTImports201Response struct {
	value *POSTImports201Response
	isSet bool
}

func (v NullablePOSTImports201Response) Get() *POSTImports201Response {
	return v.value
}

func (v *NullablePOSTImports201Response) Set(val *POSTImports201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTImports201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTImports201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTImports201Response(val *POSTImports201Response) *NullablePOSTImports201Response {
	return &NullablePOSTImports201Response{value: val, isSet: true}
}

func (v NullablePOSTImports201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTImports201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
