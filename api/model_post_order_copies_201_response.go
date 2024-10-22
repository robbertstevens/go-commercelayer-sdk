/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.3.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the POSTOrderCopies201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTOrderCopies201Response{}

// POSTOrderCopies201Response struct for POSTOrderCopies201Response
type POSTOrderCopies201Response struct {
	Data *POSTOrderCopies201ResponseData `json:"data,omitempty"`
}

// NewPOSTOrderCopies201Response instantiates a new POSTOrderCopies201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTOrderCopies201Response() *POSTOrderCopies201Response {
	this := POSTOrderCopies201Response{}
	return &this
}

// NewPOSTOrderCopies201ResponseWithDefaults instantiates a new POSTOrderCopies201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTOrderCopies201ResponseWithDefaults() *POSTOrderCopies201Response {
	this := POSTOrderCopies201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTOrderCopies201Response) GetData() POSTOrderCopies201ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret POSTOrderCopies201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrderCopies201Response) GetDataOk() (*POSTOrderCopies201ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTOrderCopies201Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTOrderCopies201ResponseData and assigns it to the Data field.
func (o *POSTOrderCopies201Response) SetData(v POSTOrderCopies201ResponseData) {
	o.Data = &v
}

func (o POSTOrderCopies201Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTOrderCopies201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTOrderCopies201Response struct {
	value *POSTOrderCopies201Response
	isSet bool
}

func (v NullablePOSTOrderCopies201Response) Get() *POSTOrderCopies201Response {
	return v.value
}

func (v *NullablePOSTOrderCopies201Response) Set(val *POSTOrderCopies201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTOrderCopies201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTOrderCopies201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTOrderCopies201Response(val *POSTOrderCopies201Response) *NullablePOSTOrderCopies201Response {
	return &NullablePOSTOrderCopies201Response{value: val, isSet: true}
}

func (v NullablePOSTOrderCopies201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTOrderCopies201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
