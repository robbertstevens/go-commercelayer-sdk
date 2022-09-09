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

// GETFixedPricePromotions200Response struct for GETFixedPricePromotions200Response
type GETFixedPricePromotions200Response struct {
	Data []GETFixedPricePromotions200ResponseDataInner `json:"data,omitempty"`
}

// NewGETFixedPricePromotions200Response instantiates a new GETFixedPricePromotions200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETFixedPricePromotions200Response() *GETFixedPricePromotions200Response {
	this := GETFixedPricePromotions200Response{}
	return &this
}

// NewGETFixedPricePromotions200ResponseWithDefaults instantiates a new GETFixedPricePromotions200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETFixedPricePromotions200ResponseWithDefaults() *GETFixedPricePromotions200Response {
	this := GETFixedPricePromotions200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETFixedPricePromotions200Response) GetData() []GETFixedPricePromotions200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret []GETFixedPricePromotions200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETFixedPricePromotions200Response) GetDataOk() ([]GETFixedPricePromotions200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETFixedPricePromotions200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETFixedPricePromotions200ResponseDataInner and assigns it to the Data field.
func (o *GETFixedPricePromotions200Response) SetData(v []GETFixedPricePromotions200ResponseDataInner) {
	o.Data = v
}

func (o GETFixedPricePromotions200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETFixedPricePromotions200Response struct {
	value *GETFixedPricePromotions200Response
	isSet bool
}

func (v NullableGETFixedPricePromotions200Response) Get() *GETFixedPricePromotions200Response {
	return v.value
}

func (v *NullableGETFixedPricePromotions200Response) Set(val *GETFixedPricePromotions200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETFixedPricePromotions200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETFixedPricePromotions200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETFixedPricePromotions200Response(val *GETFixedPricePromotions200Response) *NullableGETFixedPricePromotions200Response {
	return &NullableGETFixedPricePromotions200Response{value: val, isSet: true}
}

func (v NullableGETFixedPricePromotions200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETFixedPricePromotions200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


