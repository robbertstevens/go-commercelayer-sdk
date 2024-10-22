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

// checks if the POSTMarkets201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTMarkets201Response{}

// POSTMarkets201Response struct for POSTMarkets201Response
type POSTMarkets201Response struct {
	Data *POSTMarkets201ResponseData `json:"data,omitempty"`
}

// NewPOSTMarkets201Response instantiates a new POSTMarkets201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTMarkets201Response() *POSTMarkets201Response {
	this := POSTMarkets201Response{}
	return &this
}

// NewPOSTMarkets201ResponseWithDefaults instantiates a new POSTMarkets201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTMarkets201ResponseWithDefaults() *POSTMarkets201Response {
	this := POSTMarkets201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTMarkets201Response) GetData() POSTMarkets201ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret POSTMarkets201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTMarkets201Response) GetDataOk() (*POSTMarkets201ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTMarkets201Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTMarkets201ResponseData and assigns it to the Data field.
func (o *POSTMarkets201Response) SetData(v POSTMarkets201ResponseData) {
	o.Data = &v
}

func (o POSTMarkets201Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTMarkets201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTMarkets201Response struct {
	value *POSTMarkets201Response
	isSet bool
}

func (v NullablePOSTMarkets201Response) Get() *POSTMarkets201Response {
	return v.value
}

func (v *NullablePOSTMarkets201Response) Set(val *POSTMarkets201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTMarkets201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTMarkets201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTMarkets201Response(val *POSTMarkets201Response) *NullablePOSTMarkets201Response {
	return &NullablePOSTMarkets201Response{value: val, isSet: true}
}

func (v NullablePOSTMarkets201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTMarkets201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
