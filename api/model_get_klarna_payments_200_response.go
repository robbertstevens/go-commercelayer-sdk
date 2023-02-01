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

// GETKlarnaPayments200Response struct for GETKlarnaPayments200Response
type GETKlarnaPayments200Response struct {
	Data []GETKlarnaPayments200ResponseDataInner `json:"data,omitempty"`
}

// NewGETKlarnaPayments200Response instantiates a new GETKlarnaPayments200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETKlarnaPayments200Response() *GETKlarnaPayments200Response {
	this := GETKlarnaPayments200Response{}
	return &this
}

// NewGETKlarnaPayments200ResponseWithDefaults instantiates a new GETKlarnaPayments200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETKlarnaPayments200ResponseWithDefaults() *GETKlarnaPayments200Response {
	this := GETKlarnaPayments200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETKlarnaPayments200Response) GetData() []GETKlarnaPayments200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret []GETKlarnaPayments200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETKlarnaPayments200Response) GetDataOk() ([]GETKlarnaPayments200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETKlarnaPayments200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETKlarnaPayments200ResponseDataInner and assigns it to the Data field.
func (o *GETKlarnaPayments200Response) SetData(v []GETKlarnaPayments200ResponseDataInner) {
	o.Data = v
}

func (o GETKlarnaPayments200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETKlarnaPayments200Response struct {
	value *GETKlarnaPayments200Response
	isSet bool
}

func (v NullableGETKlarnaPayments200Response) Get() *GETKlarnaPayments200Response {
	return v.value
}

func (v *NullableGETKlarnaPayments200Response) Set(val *GETKlarnaPayments200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETKlarnaPayments200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETKlarnaPayments200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETKlarnaPayments200Response(val *GETKlarnaPayments200Response) *NullableGETKlarnaPayments200Response {
	return &NullableGETKlarnaPayments200Response{value: val, isSet: true}
}

func (v NullableGETKlarnaPayments200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETKlarnaPayments200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
