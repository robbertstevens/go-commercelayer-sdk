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

// POSTShippingZones201Response struct for POSTShippingZones201Response
type POSTShippingZones201Response struct {
	Data *POSTShippingZones201ResponseData `json:"data,omitempty"`
}

// NewPOSTShippingZones201Response instantiates a new POSTShippingZones201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTShippingZones201Response() *POSTShippingZones201Response {
	this := POSTShippingZones201Response{}
	return &this
}

// NewPOSTShippingZones201ResponseWithDefaults instantiates a new POSTShippingZones201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTShippingZones201ResponseWithDefaults() *POSTShippingZones201Response {
	this := POSTShippingZones201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTShippingZones201Response) GetData() POSTShippingZones201ResponseData {
	if o == nil || o.Data == nil {
		var ret POSTShippingZones201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTShippingZones201Response) GetDataOk() (*POSTShippingZones201ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTShippingZones201Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTShippingZones201ResponseData and assigns it to the Data field.
func (o *POSTShippingZones201Response) SetData(v POSTShippingZones201ResponseData) {
	o.Data = &v
}

func (o POSTShippingZones201Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePOSTShippingZones201Response struct {
	value *POSTShippingZones201Response
	isSet bool
}

func (v NullablePOSTShippingZones201Response) Get() *POSTShippingZones201Response {
	return v.value
}

func (v *NullablePOSTShippingZones201Response) Set(val *POSTShippingZones201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTShippingZones201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTShippingZones201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTShippingZones201Response(val *POSTShippingZones201Response) *NullablePOSTShippingZones201Response {
	return &NullablePOSTShippingZones201Response{value: val, isSet: true}
}

func (v NullablePOSTShippingZones201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTShippingZones201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
