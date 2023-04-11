/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 4.1.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the POSTShippingMethods201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTShippingMethods201Response{}

// POSTShippingMethods201Response struct for POSTShippingMethods201Response
type POSTShippingMethods201Response struct {
	Data *POSTShippingMethods201ResponseData `json:"data,omitempty"`
}

// NewPOSTShippingMethods201Response instantiates a new POSTShippingMethods201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTShippingMethods201Response() *POSTShippingMethods201Response {
	this := POSTShippingMethods201Response{}
	return &this
}

// NewPOSTShippingMethods201ResponseWithDefaults instantiates a new POSTShippingMethods201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTShippingMethods201ResponseWithDefaults() *POSTShippingMethods201Response {
	this := POSTShippingMethods201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTShippingMethods201Response) GetData() POSTShippingMethods201ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret POSTShippingMethods201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTShippingMethods201Response) GetDataOk() (*POSTShippingMethods201ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTShippingMethods201Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTShippingMethods201ResponseData and assigns it to the Data field.
func (o *POSTShippingMethods201Response) SetData(v POSTShippingMethods201ResponseData) {
	o.Data = &v
}

func (o POSTShippingMethods201Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTShippingMethods201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTShippingMethods201Response struct {
	value *POSTShippingMethods201Response
	isSet bool
}

func (v NullablePOSTShippingMethods201Response) Get() *POSTShippingMethods201Response {
	return v.value
}

func (v *NullablePOSTShippingMethods201Response) Set(val *POSTShippingMethods201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTShippingMethods201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTShippingMethods201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTShippingMethods201Response(val *POSTShippingMethods201Response) *NullablePOSTShippingMethods201Response {
	return &NullablePOSTShippingMethods201Response{value: val, isSet: true}
}

func (v NullablePOSTShippingMethods201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTShippingMethods201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
