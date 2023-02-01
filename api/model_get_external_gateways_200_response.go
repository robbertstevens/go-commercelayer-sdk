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

// GETExternalGateways200Response struct for GETExternalGateways200Response
type GETExternalGateways200Response struct {
	Data []GETExternalGateways200ResponseDataInner `json:"data,omitempty"`
}

// NewGETExternalGateways200Response instantiates a new GETExternalGateways200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETExternalGateways200Response() *GETExternalGateways200Response {
	this := GETExternalGateways200Response{}
	return &this
}

// NewGETExternalGateways200ResponseWithDefaults instantiates a new GETExternalGateways200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETExternalGateways200ResponseWithDefaults() *GETExternalGateways200Response {
	this := GETExternalGateways200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETExternalGateways200Response) GetData() []GETExternalGateways200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret []GETExternalGateways200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETExternalGateways200Response) GetDataOk() ([]GETExternalGateways200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETExternalGateways200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETExternalGateways200ResponseDataInner and assigns it to the Data field.
func (o *GETExternalGateways200Response) SetData(v []GETExternalGateways200ResponseDataInner) {
	o.Data = v
}

func (o GETExternalGateways200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETExternalGateways200Response struct {
	value *GETExternalGateways200Response
	isSet bool
}

func (v NullableGETExternalGateways200Response) Get() *GETExternalGateways200Response {
	return v.value
}

func (v *NullableGETExternalGateways200Response) Set(val *GETExternalGateways200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETExternalGateways200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETExternalGateways200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETExternalGateways200Response(val *GETExternalGateways200Response) *NullableGETExternalGateways200Response {
	return &NullableGETExternalGateways200Response{value: val, isSet: true}
}

func (v NullableGETExternalGateways200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETExternalGateways200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
