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

// GETShippingMethods200Response struct for GETShippingMethods200Response
type GETShippingMethods200Response struct {
	Data []GETShippingMethods200ResponseDataInner `json:"data,omitempty"`
}

// NewGETShippingMethods200Response instantiates a new GETShippingMethods200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETShippingMethods200Response() *GETShippingMethods200Response {
	this := GETShippingMethods200Response{}
	return &this
}

// NewGETShippingMethods200ResponseWithDefaults instantiates a new GETShippingMethods200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETShippingMethods200ResponseWithDefaults() *GETShippingMethods200Response {
	this := GETShippingMethods200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETShippingMethods200Response) GetData() []GETShippingMethods200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret []GETShippingMethods200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShippingMethods200Response) GetDataOk() ([]GETShippingMethods200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETShippingMethods200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETShippingMethods200ResponseDataInner and assigns it to the Data field.
func (o *GETShippingMethods200Response) SetData(v []GETShippingMethods200ResponseDataInner) {
	o.Data = v
}

func (o GETShippingMethods200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETShippingMethods200Response struct {
	value *GETShippingMethods200Response
	isSet bool
}

func (v NullableGETShippingMethods200Response) Get() *GETShippingMethods200Response {
	return v.value
}

func (v *NullableGETShippingMethods200Response) Set(val *GETShippingMethods200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETShippingMethods200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETShippingMethods200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETShippingMethods200Response(val *GETShippingMethods200Response) *NullableGETShippingMethods200Response {
	return &NullableGETShippingMethods200Response{value: val, isSet: true}
}

func (v NullableGETShippingMethods200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETShippingMethods200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
