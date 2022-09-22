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

// GETPaymentGateways200Response struct for GETPaymentGateways200Response
type GETPaymentGateways200Response struct {
	Data []GETPaymentGateways200ResponseDataInner `json:"data,omitempty"`
}

// NewGETPaymentGateways200Response instantiates a new GETPaymentGateways200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETPaymentGateways200Response() *GETPaymentGateways200Response {
	this := GETPaymentGateways200Response{}
	return &this
}

// NewGETPaymentGateways200ResponseWithDefaults instantiates a new GETPaymentGateways200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETPaymentGateways200ResponseWithDefaults() *GETPaymentGateways200Response {
	this := GETPaymentGateways200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETPaymentGateways200Response) GetData() []GETPaymentGateways200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret []GETPaymentGateways200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETPaymentGateways200Response) GetDataOk() ([]GETPaymentGateways200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETPaymentGateways200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETPaymentGateways200ResponseDataInner and assigns it to the Data field.
func (o *GETPaymentGateways200Response) SetData(v []GETPaymentGateways200ResponseDataInner) {
	o.Data = v
}

func (o GETPaymentGateways200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETPaymentGateways200Response struct {
	value *GETPaymentGateways200Response
	isSet bool
}

func (v NullableGETPaymentGateways200Response) Get() *GETPaymentGateways200Response {
	return v.value
}

func (v *NullableGETPaymentGateways200Response) Set(val *GETPaymentGateways200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETPaymentGateways200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETPaymentGateways200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETPaymentGateways200Response(val *GETPaymentGateways200Response) *NullableGETPaymentGateways200Response {
	return &NullableGETPaymentGateways200Response{value: val, isSet: true}
}

func (v NullableGETPaymentGateways200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETPaymentGateways200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
