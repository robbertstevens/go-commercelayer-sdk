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

// POSTExternalGateways201Response struct for POSTExternalGateways201Response
type POSTExternalGateways201Response struct {
	Data *POSTExternalGateways201ResponseData `json:"data,omitempty"`
}

// NewPOSTExternalGateways201Response instantiates a new POSTExternalGateways201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTExternalGateways201Response() *POSTExternalGateways201Response {
	this := POSTExternalGateways201Response{}
	return &this
}

// NewPOSTExternalGateways201ResponseWithDefaults instantiates a new POSTExternalGateways201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTExternalGateways201ResponseWithDefaults() *POSTExternalGateways201Response {
	this := POSTExternalGateways201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTExternalGateways201Response) GetData() POSTExternalGateways201ResponseData {
	if o == nil || o.Data == nil {
		var ret POSTExternalGateways201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTExternalGateways201Response) GetDataOk() (*POSTExternalGateways201ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTExternalGateways201Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTExternalGateways201ResponseData and assigns it to the Data field.
func (o *POSTExternalGateways201Response) SetData(v POSTExternalGateways201ResponseData) {
	o.Data = &v
}

func (o POSTExternalGateways201Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePOSTExternalGateways201Response struct {
	value *POSTExternalGateways201Response
	isSet bool
}

func (v NullablePOSTExternalGateways201Response) Get() *POSTExternalGateways201Response {
	return v.value
}

func (v *NullablePOSTExternalGateways201Response) Set(val *POSTExternalGateways201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTExternalGateways201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTExternalGateways201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTExternalGateways201Response(val *POSTExternalGateways201Response) *NullablePOSTExternalGateways201Response {
	return &NullablePOSTExternalGateways201Response{value: val, isSet: true}
}

func (v NullablePOSTExternalGateways201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTExternalGateways201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
