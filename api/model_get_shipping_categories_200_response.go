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

// GETShippingCategories200Response struct for GETShippingCategories200Response
type GETShippingCategories200Response struct {
	Data []GETShippingCategories200ResponseDataInner `json:"data,omitempty"`
}

// NewGETShippingCategories200Response instantiates a new GETShippingCategories200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETShippingCategories200Response() *GETShippingCategories200Response {
	this := GETShippingCategories200Response{}
	return &this
}

// NewGETShippingCategories200ResponseWithDefaults instantiates a new GETShippingCategories200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETShippingCategories200ResponseWithDefaults() *GETShippingCategories200Response {
	this := GETShippingCategories200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETShippingCategories200Response) GetData() []GETShippingCategories200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret []GETShippingCategories200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShippingCategories200Response) GetDataOk() ([]GETShippingCategories200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETShippingCategories200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETShippingCategories200ResponseDataInner and assigns it to the Data field.
func (o *GETShippingCategories200Response) SetData(v []GETShippingCategories200ResponseDataInner) {
	o.Data = v
}

func (o GETShippingCategories200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETShippingCategories200Response struct {
	value *GETShippingCategories200Response
	isSet bool
}

func (v NullableGETShippingCategories200Response) Get() *GETShippingCategories200Response {
	return v.value
}

func (v *NullableGETShippingCategories200Response) Set(val *GETShippingCategories200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETShippingCategories200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETShippingCategories200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETShippingCategories200Response(val *GETShippingCategories200Response) *NullableGETShippingCategories200Response {
	return &NullableGETShippingCategories200Response{value: val, isSet: true}
}

func (v NullableGETShippingCategories200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETShippingCategories200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
