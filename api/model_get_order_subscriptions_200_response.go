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

// GETOrderSubscriptions200Response struct for GETOrderSubscriptions200Response
type GETOrderSubscriptions200Response struct {
	Data []GETOrderSubscriptions200ResponseDataInner `json:"data,omitempty"`
}

// NewGETOrderSubscriptions200Response instantiates a new GETOrderSubscriptions200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETOrderSubscriptions200Response() *GETOrderSubscriptions200Response {
	this := GETOrderSubscriptions200Response{}
	return &this
}

// NewGETOrderSubscriptions200ResponseWithDefaults instantiates a new GETOrderSubscriptions200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETOrderSubscriptions200ResponseWithDefaults() *GETOrderSubscriptions200Response {
	this := GETOrderSubscriptions200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETOrderSubscriptions200Response) GetData() []GETOrderSubscriptions200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret []GETOrderSubscriptions200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrderSubscriptions200Response) GetDataOk() ([]GETOrderSubscriptions200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETOrderSubscriptions200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETOrderSubscriptions200ResponseDataInner and assigns it to the Data field.
func (o *GETOrderSubscriptions200Response) SetData(v []GETOrderSubscriptions200ResponseDataInner) {
	o.Data = v
}

func (o GETOrderSubscriptions200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETOrderSubscriptions200Response struct {
	value *GETOrderSubscriptions200Response
	isSet bool
}

func (v NullableGETOrderSubscriptions200Response) Get() *GETOrderSubscriptions200Response {
	return v.value
}

func (v *NullableGETOrderSubscriptions200Response) Set(val *GETOrderSubscriptions200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETOrderSubscriptions200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETOrderSubscriptions200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETOrderSubscriptions200Response(val *GETOrderSubscriptions200Response) *NullableGETOrderSubscriptions200Response {
	return &NullableGETOrderSubscriptions200Response{value: val, isSet: true}
}

func (v NullableGETOrderSubscriptions200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETOrderSubscriptions200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
