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

// POSTLineItemOptions201Response struct for POSTLineItemOptions201Response
type POSTLineItemOptions201Response struct {
	Data *POSTLineItemOptions201ResponseData `json:"data,omitempty"`
}

// NewPOSTLineItemOptions201Response instantiates a new POSTLineItemOptions201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTLineItemOptions201Response() *POSTLineItemOptions201Response {
	this := POSTLineItemOptions201Response{}
	return &this
}

// NewPOSTLineItemOptions201ResponseWithDefaults instantiates a new POSTLineItemOptions201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTLineItemOptions201ResponseWithDefaults() *POSTLineItemOptions201Response {
	this := POSTLineItemOptions201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTLineItemOptions201Response) GetData() POSTLineItemOptions201ResponseData {
	if o == nil || o.Data == nil {
		var ret POSTLineItemOptions201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTLineItemOptions201Response) GetDataOk() (*POSTLineItemOptions201ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTLineItemOptions201Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTLineItemOptions201ResponseData and assigns it to the Data field.
func (o *POSTLineItemOptions201Response) SetData(v POSTLineItemOptions201ResponseData) {
	o.Data = &v
}

func (o POSTLineItemOptions201Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePOSTLineItemOptions201Response struct {
	value *POSTLineItemOptions201Response
	isSet bool
}

func (v NullablePOSTLineItemOptions201Response) Get() *POSTLineItemOptions201Response {
	return v.value
}

func (v *NullablePOSTLineItemOptions201Response) Set(val *POSTLineItemOptions201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTLineItemOptions201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTLineItemOptions201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTLineItemOptions201Response(val *POSTLineItemOptions201Response) *NullablePOSTLineItemOptions201Response {
	return &NullablePOSTLineItemOptions201Response{value: val, isSet: true}
}

func (v NullablePOSTLineItemOptions201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTLineItemOptions201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
