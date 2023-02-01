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

// GETShippingCategoriesShippingCategoryId200Response struct for GETShippingCategoriesShippingCategoryId200Response
type GETShippingCategoriesShippingCategoryId200Response struct {
	Data *GETShippingCategories200ResponseDataInner `json:"data,omitempty"`
}

// NewGETShippingCategoriesShippingCategoryId200Response instantiates a new GETShippingCategoriesShippingCategoryId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETShippingCategoriesShippingCategoryId200Response() *GETShippingCategoriesShippingCategoryId200Response {
	this := GETShippingCategoriesShippingCategoryId200Response{}
	return &this
}

// NewGETShippingCategoriesShippingCategoryId200ResponseWithDefaults instantiates a new GETShippingCategoriesShippingCategoryId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETShippingCategoriesShippingCategoryId200ResponseWithDefaults() *GETShippingCategoriesShippingCategoryId200Response {
	this := GETShippingCategoriesShippingCategoryId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETShippingCategoriesShippingCategoryId200Response) GetData() GETShippingCategories200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret GETShippingCategories200ResponseDataInner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShippingCategoriesShippingCategoryId200Response) GetDataOk() (*GETShippingCategories200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETShippingCategoriesShippingCategoryId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETShippingCategories200ResponseDataInner and assigns it to the Data field.
func (o *GETShippingCategoriesShippingCategoryId200Response) SetData(v GETShippingCategories200ResponseDataInner) {
	o.Data = &v
}

func (o GETShippingCategoriesShippingCategoryId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETShippingCategoriesShippingCategoryId200Response struct {
	value *GETShippingCategoriesShippingCategoryId200Response
	isSet bool
}

func (v NullableGETShippingCategoriesShippingCategoryId200Response) Get() *GETShippingCategoriesShippingCategoryId200Response {
	return v.value
}

func (v *NullableGETShippingCategoriesShippingCategoryId200Response) Set(val *GETShippingCategoriesShippingCategoryId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETShippingCategoriesShippingCategoryId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETShippingCategoriesShippingCategoryId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETShippingCategoriesShippingCategoryId200Response(val *GETShippingCategoriesShippingCategoryId200Response) *NullableGETShippingCategoriesShippingCategoryId200Response {
	return &NullableGETShippingCategoriesShippingCategoryId200Response{value: val, isSet: true}
}

func (v NullableGETShippingCategoriesShippingCategoryId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETShippingCategoriesShippingCategoryId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
