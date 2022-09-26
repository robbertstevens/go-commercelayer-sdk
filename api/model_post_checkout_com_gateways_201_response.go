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

// POSTCheckoutComGateways201Response struct for POSTCheckoutComGateways201Response
type POSTCheckoutComGateways201Response struct {
	Data *POSTCheckoutComGateways201ResponseData `json:"data,omitempty"`
}

// NewPOSTCheckoutComGateways201Response instantiates a new POSTCheckoutComGateways201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTCheckoutComGateways201Response() *POSTCheckoutComGateways201Response {
	this := POSTCheckoutComGateways201Response{}
	return &this
}

// NewPOSTCheckoutComGateways201ResponseWithDefaults instantiates a new POSTCheckoutComGateways201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTCheckoutComGateways201ResponseWithDefaults() *POSTCheckoutComGateways201Response {
	this := POSTCheckoutComGateways201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTCheckoutComGateways201Response) GetData() POSTCheckoutComGateways201ResponseData {
	if o == nil || o.Data == nil {
		var ret POSTCheckoutComGateways201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCheckoutComGateways201Response) GetDataOk() (*POSTCheckoutComGateways201ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTCheckoutComGateways201Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTCheckoutComGateways201ResponseData and assigns it to the Data field.
func (o *POSTCheckoutComGateways201Response) SetData(v POSTCheckoutComGateways201ResponseData) {
	o.Data = &v
}

func (o POSTCheckoutComGateways201Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePOSTCheckoutComGateways201Response struct {
	value *POSTCheckoutComGateways201Response
	isSet bool
}

func (v NullablePOSTCheckoutComGateways201Response) Get() *POSTCheckoutComGateways201Response {
	return v.value
}

func (v *NullablePOSTCheckoutComGateways201Response) Set(val *POSTCheckoutComGateways201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTCheckoutComGateways201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTCheckoutComGateways201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTCheckoutComGateways201Response(val *POSTCheckoutComGateways201Response) *NullablePOSTCheckoutComGateways201Response {
	return &NullablePOSTCheckoutComGateways201Response{value: val, isSet: true}
}

func (v NullablePOSTCheckoutComGateways201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTCheckoutComGateways201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
