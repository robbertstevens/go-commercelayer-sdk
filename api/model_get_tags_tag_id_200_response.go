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

// checks if the GETTagsTagId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETTagsTagId200Response{}

// GETTagsTagId200Response struct for GETTagsTagId200Response
type GETTagsTagId200Response struct {
	Data *GETTagsTagId200ResponseData `json:"data,omitempty"`
}

// NewGETTagsTagId200Response instantiates a new GETTagsTagId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETTagsTagId200Response() *GETTagsTagId200Response {
	this := GETTagsTagId200Response{}
	return &this
}

// NewGETTagsTagId200ResponseWithDefaults instantiates a new GETTagsTagId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETTagsTagId200ResponseWithDefaults() *GETTagsTagId200Response {
	this := GETTagsTagId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETTagsTagId200Response) GetData() GETTagsTagId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret GETTagsTagId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETTagsTagId200Response) GetDataOk() (*GETTagsTagId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETTagsTagId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GETTagsTagId200ResponseData and assigns it to the Data field.
func (o *GETTagsTagId200Response) SetData(v GETTagsTagId200ResponseData) {
	o.Data = &v
}

func (o GETTagsTagId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETTagsTagId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGETTagsTagId200Response struct {
	value *GETTagsTagId200Response
	isSet bool
}

func (v NullableGETTagsTagId200Response) Get() *GETTagsTagId200Response {
	return v.value
}

func (v *NullableGETTagsTagId200Response) Set(val *GETTagsTagId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETTagsTagId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETTagsTagId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETTagsTagId200Response(val *GETTagsTagId200Response) *NullableGETTagsTagId200Response {
	return &NullableGETTagsTagId200Response{value: val, isSet: true}
}

func (v NullableGETTagsTagId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETTagsTagId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
