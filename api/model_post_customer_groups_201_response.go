/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.4.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// POSTCustomerGroups201Response struct for POSTCustomerGroups201Response
type POSTCustomerGroups201Response struct {
	Data *POSTCustomerGroups201ResponseData `json:"data,omitempty"`
}

// NewPOSTCustomerGroups201Response instantiates a new POSTCustomerGroups201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTCustomerGroups201Response() *POSTCustomerGroups201Response {
	this := POSTCustomerGroups201Response{}
	return &this
}

// NewPOSTCustomerGroups201ResponseWithDefaults instantiates a new POSTCustomerGroups201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTCustomerGroups201ResponseWithDefaults() *POSTCustomerGroups201Response {
	this := POSTCustomerGroups201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTCustomerGroups201Response) GetData() POSTCustomerGroups201ResponseData {
	if o == nil || o.Data == nil {
		var ret POSTCustomerGroups201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCustomerGroups201Response) GetDataOk() (*POSTCustomerGroups201ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTCustomerGroups201Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTCustomerGroups201ResponseData and assigns it to the Data field.
func (o *POSTCustomerGroups201Response) SetData(v POSTCustomerGroups201ResponseData) {
	o.Data = &v
}

func (o POSTCustomerGroups201Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePOSTCustomerGroups201Response struct {
	value *POSTCustomerGroups201Response
	isSet bool
}

func (v NullablePOSTCustomerGroups201Response) Get() *POSTCustomerGroups201Response {
	return v.value
}

func (v *NullablePOSTCustomerGroups201Response) Set(val *POSTCustomerGroups201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTCustomerGroups201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTCustomerGroups201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTCustomerGroups201Response(val *POSTCustomerGroups201Response) *NullablePOSTCustomerGroups201Response {
	return &NullablePOSTCustomerGroups201Response{value: val, isSet: true}
}

func (v NullablePOSTCustomerGroups201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTCustomerGroups201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
