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

// GETPaymentMethods200Response struct for GETPaymentMethods200Response
type GETPaymentMethods200Response struct {
	Data []GETPaymentMethods200ResponseDataInner `json:"data,omitempty"`
}

// NewGETPaymentMethods200Response instantiates a new GETPaymentMethods200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETPaymentMethods200Response() *GETPaymentMethods200Response {
	this := GETPaymentMethods200Response{}
	return &this
}

// NewGETPaymentMethods200ResponseWithDefaults instantiates a new GETPaymentMethods200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETPaymentMethods200ResponseWithDefaults() *GETPaymentMethods200Response {
	this := GETPaymentMethods200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETPaymentMethods200Response) GetData() []GETPaymentMethods200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret []GETPaymentMethods200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETPaymentMethods200Response) GetDataOk() ([]GETPaymentMethods200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETPaymentMethods200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETPaymentMethods200ResponseDataInner and assigns it to the Data field.
func (o *GETPaymentMethods200Response) SetData(v []GETPaymentMethods200ResponseDataInner) {
	o.Data = v
}

func (o GETPaymentMethods200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETPaymentMethods200Response struct {
	value *GETPaymentMethods200Response
	isSet bool
}

func (v NullableGETPaymentMethods200Response) Get() *GETPaymentMethods200Response {
	return v.value
}

func (v *NullableGETPaymentMethods200Response) Set(val *GETPaymentMethods200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETPaymentMethods200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETPaymentMethods200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETPaymentMethods200Response(val *GETPaymentMethods200Response) *NullableGETPaymentMethods200Response {
	return &NullableGETPaymentMethods200Response{value: val, isSet: true}
}

func (v NullableGETPaymentMethods200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETPaymentMethods200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
