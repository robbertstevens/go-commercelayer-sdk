/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETMarkets200Response struct for GETMarkets200Response
type GETMarkets200Response struct {
	Data []GETMarkets200ResponseDataInner `json:"data,omitempty"`
}

// NewGETMarkets200Response instantiates a new GETMarkets200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETMarkets200Response() *GETMarkets200Response {
	this := GETMarkets200Response{}
	return &this
}

// NewGETMarkets200ResponseWithDefaults instantiates a new GETMarkets200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETMarkets200ResponseWithDefaults() *GETMarkets200Response {
	this := GETMarkets200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETMarkets200Response) GetData() []GETMarkets200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret []GETMarkets200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETMarkets200Response) GetDataOk() ([]GETMarkets200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETMarkets200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETMarkets200ResponseDataInner and assigns it to the Data field.
func (o *GETMarkets200Response) SetData(v []GETMarkets200ResponseDataInner) {
	o.Data = v
}

func (o GETMarkets200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETMarkets200Response struct {
	value *GETMarkets200Response
	isSet bool
}

func (v NullableGETMarkets200Response) Get() *GETMarkets200Response {
	return v.value
}

func (v *NullableGETMarkets200Response) Set(val *GETMarkets200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETMarkets200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETMarkets200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETMarkets200Response(val *GETMarkets200Response) *NullableGETMarkets200Response {
	return &NullableGETMarkets200Response{value: val, isSet: true}
}

func (v NullableGETMarkets200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETMarkets200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
