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

// GETShippingMethodsShippingMethodId200Response struct for GETShippingMethodsShippingMethodId200Response
type GETShippingMethodsShippingMethodId200Response struct {
	Data *GETShippingMethods200ResponseDataInner `json:"data,omitempty"`
}

// NewGETShippingMethodsShippingMethodId200Response instantiates a new GETShippingMethodsShippingMethodId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETShippingMethodsShippingMethodId200Response() *GETShippingMethodsShippingMethodId200Response {
	this := GETShippingMethodsShippingMethodId200Response{}
	return &this
}

// NewGETShippingMethodsShippingMethodId200ResponseWithDefaults instantiates a new GETShippingMethodsShippingMethodId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETShippingMethodsShippingMethodId200ResponseWithDefaults() *GETShippingMethodsShippingMethodId200Response {
	this := GETShippingMethodsShippingMethodId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETShippingMethodsShippingMethodId200Response) GetData() GETShippingMethods200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret GETShippingMethods200ResponseDataInner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShippingMethodsShippingMethodId200Response) GetDataOk() (*GETShippingMethods200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETShippingMethodsShippingMethodId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETShippingMethods200ResponseDataInner and assigns it to the Data field.
func (o *GETShippingMethodsShippingMethodId200Response) SetData(v GETShippingMethods200ResponseDataInner) {
	o.Data = &v
}

func (o GETShippingMethodsShippingMethodId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETShippingMethodsShippingMethodId200Response struct {
	value *GETShippingMethodsShippingMethodId200Response
	isSet bool
}

func (v NullableGETShippingMethodsShippingMethodId200Response) Get() *GETShippingMethodsShippingMethodId200Response {
	return v.value
}

func (v *NullableGETShippingMethodsShippingMethodId200Response) Set(val *GETShippingMethodsShippingMethodId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETShippingMethodsShippingMethodId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETShippingMethodsShippingMethodId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETShippingMethodsShippingMethodId200Response(val *GETShippingMethodsShippingMethodId200Response) *NullableGETShippingMethodsShippingMethodId200Response {
	return &NullableGETShippingMethodsShippingMethodId200Response{value: val, isSet: true}
}

func (v NullableGETShippingMethodsShippingMethodId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETShippingMethodsShippingMethodId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
