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

// GETTaxCategoriesTaxCategoryId200Response struct for GETTaxCategoriesTaxCategoryId200Response
type GETTaxCategoriesTaxCategoryId200Response struct {
	Data *GETTaxCategories200ResponseDataInner `json:"data,omitempty"`
}

// NewGETTaxCategoriesTaxCategoryId200Response instantiates a new GETTaxCategoriesTaxCategoryId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETTaxCategoriesTaxCategoryId200Response() *GETTaxCategoriesTaxCategoryId200Response {
	this := GETTaxCategoriesTaxCategoryId200Response{}
	return &this
}

// NewGETTaxCategoriesTaxCategoryId200ResponseWithDefaults instantiates a new GETTaxCategoriesTaxCategoryId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETTaxCategoriesTaxCategoryId200ResponseWithDefaults() *GETTaxCategoriesTaxCategoryId200Response {
	this := GETTaxCategoriesTaxCategoryId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETTaxCategoriesTaxCategoryId200Response) GetData() GETTaxCategories200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret GETTaxCategories200ResponseDataInner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETTaxCategoriesTaxCategoryId200Response) GetDataOk() (*GETTaxCategories200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETTaxCategoriesTaxCategoryId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETTaxCategories200ResponseDataInner and assigns it to the Data field.
func (o *GETTaxCategoriesTaxCategoryId200Response) SetData(v GETTaxCategories200ResponseDataInner) {
	o.Data = &v
}

func (o GETTaxCategoriesTaxCategoryId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETTaxCategoriesTaxCategoryId200Response struct {
	value *GETTaxCategoriesTaxCategoryId200Response
	isSet bool
}

func (v NullableGETTaxCategoriesTaxCategoryId200Response) Get() *GETTaxCategoriesTaxCategoryId200Response {
	return v.value
}

func (v *NullableGETTaxCategoriesTaxCategoryId200Response) Set(val *GETTaxCategoriesTaxCategoryId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETTaxCategoriesTaxCategoryId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETTaxCategoriesTaxCategoryId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETTaxCategoriesTaxCategoryId200Response(val *GETTaxCategoriesTaxCategoryId200Response) *NullableGETTaxCategoriesTaxCategoryId200Response {
	return &NullableGETTaxCategoriesTaxCategoryId200Response{value: val, isSet: true}
}

func (v NullableGETTaxCategoriesTaxCategoryId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETTaxCategoriesTaxCategoryId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
