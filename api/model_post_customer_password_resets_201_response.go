/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.3.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the POSTCustomerPasswordResets201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTCustomerPasswordResets201Response{}

// POSTCustomerPasswordResets201Response struct for POSTCustomerPasswordResets201Response
type POSTCustomerPasswordResets201Response struct {
	Data *POSTCustomerPasswordResets201ResponseData `json:"data,omitempty"`
}

// NewPOSTCustomerPasswordResets201Response instantiates a new POSTCustomerPasswordResets201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTCustomerPasswordResets201Response() *POSTCustomerPasswordResets201Response {
	this := POSTCustomerPasswordResets201Response{}
	return &this
}

// NewPOSTCustomerPasswordResets201ResponseWithDefaults instantiates a new POSTCustomerPasswordResets201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTCustomerPasswordResets201ResponseWithDefaults() *POSTCustomerPasswordResets201Response {
	this := POSTCustomerPasswordResets201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTCustomerPasswordResets201Response) GetData() POSTCustomerPasswordResets201ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret POSTCustomerPasswordResets201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCustomerPasswordResets201Response) GetDataOk() (*POSTCustomerPasswordResets201ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTCustomerPasswordResets201Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTCustomerPasswordResets201ResponseData and assigns it to the Data field.
func (o *POSTCustomerPasswordResets201Response) SetData(v POSTCustomerPasswordResets201ResponseData) {
	o.Data = &v
}

func (o POSTCustomerPasswordResets201Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTCustomerPasswordResets201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTCustomerPasswordResets201Response struct {
	value *POSTCustomerPasswordResets201Response
	isSet bool
}

func (v NullablePOSTCustomerPasswordResets201Response) Get() *POSTCustomerPasswordResets201Response {
	return v.value
}

func (v *NullablePOSTCustomerPasswordResets201Response) Set(val *POSTCustomerPasswordResets201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTCustomerPasswordResets201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTCustomerPasswordResets201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTCustomerPasswordResets201Response(val *POSTCustomerPasswordResets201Response) *NullablePOSTCustomerPasswordResets201Response {
	return &NullablePOSTCustomerPasswordResets201Response{value: val, isSet: true}
}

func (v NullablePOSTCustomerPasswordResets201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTCustomerPasswordResets201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
