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

// checks if the GETReturnsReturnId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETReturnsReturnId200Response{}

// GETReturnsReturnId200Response struct for GETReturnsReturnId200Response
type GETReturnsReturnId200Response struct {
	Data *GETReturnsReturnId200ResponseData `json:"data,omitempty"`
}

// NewGETReturnsReturnId200Response instantiates a new GETReturnsReturnId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETReturnsReturnId200Response() *GETReturnsReturnId200Response {
	this := GETReturnsReturnId200Response{}
	return &this
}

// NewGETReturnsReturnId200ResponseWithDefaults instantiates a new GETReturnsReturnId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETReturnsReturnId200ResponseWithDefaults() *GETReturnsReturnId200Response {
	this := GETReturnsReturnId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETReturnsReturnId200Response) GetData() GETReturnsReturnId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret GETReturnsReturnId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETReturnsReturnId200Response) GetDataOk() (*GETReturnsReturnId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETReturnsReturnId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GETReturnsReturnId200ResponseData and assigns it to the Data field.
func (o *GETReturnsReturnId200Response) SetData(v GETReturnsReturnId200ResponseData) {
	o.Data = &v
}

func (o GETReturnsReturnId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETReturnsReturnId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGETReturnsReturnId200Response struct {
	value *GETReturnsReturnId200Response
	isSet bool
}

func (v NullableGETReturnsReturnId200Response) Get() *GETReturnsReturnId200Response {
	return v.value
}

func (v *NullableGETReturnsReturnId200Response) Set(val *GETReturnsReturnId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETReturnsReturnId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETReturnsReturnId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETReturnsReturnId200Response(val *GETReturnsReturnId200Response) *NullableGETReturnsReturnId200Response {
	return &NullableGETReturnsReturnId200Response{value: val, isSet: true}
}

func (v NullableGETReturnsReturnId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETReturnsReturnId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
