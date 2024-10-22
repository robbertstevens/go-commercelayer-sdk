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

// checks if the POSTPaypalGateways201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTPaypalGateways201Response{}

// POSTPaypalGateways201Response struct for POSTPaypalGateways201Response
type POSTPaypalGateways201Response struct {
	Data *POSTPaypalGateways201ResponseData `json:"data,omitempty"`
}

// NewPOSTPaypalGateways201Response instantiates a new POSTPaypalGateways201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTPaypalGateways201Response() *POSTPaypalGateways201Response {
	this := POSTPaypalGateways201Response{}
	return &this
}

// NewPOSTPaypalGateways201ResponseWithDefaults instantiates a new POSTPaypalGateways201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTPaypalGateways201ResponseWithDefaults() *POSTPaypalGateways201Response {
	this := POSTPaypalGateways201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTPaypalGateways201Response) GetData() POSTPaypalGateways201ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret POSTPaypalGateways201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTPaypalGateways201Response) GetDataOk() (*POSTPaypalGateways201ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTPaypalGateways201Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTPaypalGateways201ResponseData and assigns it to the Data field.
func (o *POSTPaypalGateways201Response) SetData(v POSTPaypalGateways201ResponseData) {
	o.Data = &v
}

func (o POSTPaypalGateways201Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTPaypalGateways201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTPaypalGateways201Response struct {
	value *POSTPaypalGateways201Response
	isSet bool
}

func (v NullablePOSTPaypalGateways201Response) Get() *POSTPaypalGateways201Response {
	return v.value
}

func (v *NullablePOSTPaypalGateways201Response) Set(val *POSTPaypalGateways201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTPaypalGateways201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTPaypalGateways201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTPaypalGateways201Response(val *POSTPaypalGateways201Response) *NullablePOSTPaypalGateways201Response {
	return &NullablePOSTPaypalGateways201Response{value: val, isSet: true}
}

func (v NullablePOSTPaypalGateways201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTPaypalGateways201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
