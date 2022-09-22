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

// GETEvents200Response struct for GETEvents200Response
type GETEvents200Response struct {
	Data []GETEvents200ResponseDataInner `json:"data,omitempty"`
}

// NewGETEvents200Response instantiates a new GETEvents200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETEvents200Response() *GETEvents200Response {
	this := GETEvents200Response{}
	return &this
}

// NewGETEvents200ResponseWithDefaults instantiates a new GETEvents200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETEvents200ResponseWithDefaults() *GETEvents200Response {
	this := GETEvents200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETEvents200Response) GetData() []GETEvents200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret []GETEvents200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETEvents200Response) GetDataOk() ([]GETEvents200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETEvents200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETEvents200ResponseDataInner and assigns it to the Data field.
func (o *GETEvents200Response) SetData(v []GETEvents200ResponseDataInner) {
	o.Data = v
}

func (o GETEvents200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETEvents200Response struct {
	value *GETEvents200Response
	isSet bool
}

func (v NullableGETEvents200Response) Get() *GETEvents200Response {
	return v.value
}

func (v *NullableGETEvents200Response) Set(val *GETEvents200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETEvents200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETEvents200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETEvents200Response(val *GETEvents200Response) *NullableGETEvents200Response {
	return &NullableGETEvents200Response{value: val, isSet: true}
}

func (v NullableGETEvents200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETEvents200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
