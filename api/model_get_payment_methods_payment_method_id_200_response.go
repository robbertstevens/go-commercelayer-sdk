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

// GETPaymentMethodsPaymentMethodId200Response struct for GETPaymentMethodsPaymentMethodId200Response
type GETPaymentMethodsPaymentMethodId200Response struct {
	Data *GETPaymentMethods200ResponseDataInner `json:"data,omitempty"`
}

// NewGETPaymentMethodsPaymentMethodId200Response instantiates a new GETPaymentMethodsPaymentMethodId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETPaymentMethodsPaymentMethodId200Response() *GETPaymentMethodsPaymentMethodId200Response {
	this := GETPaymentMethodsPaymentMethodId200Response{}
	return &this
}

// NewGETPaymentMethodsPaymentMethodId200ResponseWithDefaults instantiates a new GETPaymentMethodsPaymentMethodId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETPaymentMethodsPaymentMethodId200ResponseWithDefaults() *GETPaymentMethodsPaymentMethodId200Response {
	this := GETPaymentMethodsPaymentMethodId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETPaymentMethodsPaymentMethodId200Response) GetData() GETPaymentMethods200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret GETPaymentMethods200ResponseDataInner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETPaymentMethodsPaymentMethodId200Response) GetDataOk() (*GETPaymentMethods200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETPaymentMethodsPaymentMethodId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETPaymentMethods200ResponseDataInner and assigns it to the Data field.
func (o *GETPaymentMethodsPaymentMethodId200Response) SetData(v GETPaymentMethods200ResponseDataInner) {
	o.Data = &v
}

func (o GETPaymentMethodsPaymentMethodId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETPaymentMethodsPaymentMethodId200Response struct {
	value *GETPaymentMethodsPaymentMethodId200Response
	isSet bool
}

func (v NullableGETPaymentMethodsPaymentMethodId200Response) Get() *GETPaymentMethodsPaymentMethodId200Response {
	return v.value
}

func (v *NullableGETPaymentMethodsPaymentMethodId200Response) Set(val *GETPaymentMethodsPaymentMethodId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETPaymentMethodsPaymentMethodId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETPaymentMethodsPaymentMethodId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETPaymentMethodsPaymentMethodId200Response(val *GETPaymentMethodsPaymentMethodId200Response) *NullableGETPaymentMethodsPaymentMethodId200Response {
	return &NullableGETPaymentMethodsPaymentMethodId200Response{value: val, isSet: true}
}

func (v NullableGETPaymentMethodsPaymentMethodId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETPaymentMethodsPaymentMethodId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
