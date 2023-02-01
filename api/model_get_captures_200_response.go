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

// GETCaptures200Response struct for GETCaptures200Response
type GETCaptures200Response struct {
	Data []GETCaptures200ResponseDataInner `json:"data,omitempty"`
}

// NewGETCaptures200Response instantiates a new GETCaptures200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCaptures200Response() *GETCaptures200Response {
	this := GETCaptures200Response{}
	return &this
}

// NewGETCaptures200ResponseWithDefaults instantiates a new GETCaptures200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCaptures200ResponseWithDefaults() *GETCaptures200Response {
	this := GETCaptures200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETCaptures200Response) GetData() []GETCaptures200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret []GETCaptures200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCaptures200Response) GetDataOk() ([]GETCaptures200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETCaptures200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETCaptures200ResponseDataInner and assigns it to the Data field.
func (o *GETCaptures200Response) SetData(v []GETCaptures200ResponseDataInner) {
	o.Data = v
}

func (o GETCaptures200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETCaptures200Response struct {
	value *GETCaptures200Response
	isSet bool
}

func (v NullableGETCaptures200Response) Get() *GETCaptures200Response {
	return v.value
}

func (v *NullableGETCaptures200Response) Set(val *GETCaptures200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCaptures200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCaptures200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCaptures200Response(val *GETCaptures200Response) *NullableGETCaptures200Response {
	return &NullableGETCaptures200Response{value: val, isSet: true}
}

func (v NullableGETCaptures200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCaptures200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
