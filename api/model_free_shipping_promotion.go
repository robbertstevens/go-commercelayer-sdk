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

// FreeShippingPromotion struct for FreeShippingPromotion
type FreeShippingPromotion struct {
	Data FreeShippingPromotionData `json:"data"`
}

// NewFreeShippingPromotion instantiates a new FreeShippingPromotion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFreeShippingPromotion(data FreeShippingPromotionData) *FreeShippingPromotion {
	this := FreeShippingPromotion{}
	this.Data = data
	return &this
}

// NewFreeShippingPromotionWithDefaults instantiates a new FreeShippingPromotion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFreeShippingPromotionWithDefaults() *FreeShippingPromotion {
	this := FreeShippingPromotion{}
	return &this
}

// GetData returns the Data field value
func (o *FreeShippingPromotion) GetData() FreeShippingPromotionData {
	if o == nil {
		var ret FreeShippingPromotionData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *FreeShippingPromotion) GetDataOk() (*FreeShippingPromotionData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *FreeShippingPromotion) SetData(v FreeShippingPromotionData) {
	o.Data = v
}

func (o FreeShippingPromotion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableFreeShippingPromotion struct {
	value *FreeShippingPromotion
	isSet bool
}

func (v NullableFreeShippingPromotion) Get() *FreeShippingPromotion {
	return v.value
}

func (v *NullableFreeShippingPromotion) Set(val *FreeShippingPromotion) {
	v.value = val
	v.isSet = true
}

func (v NullableFreeShippingPromotion) IsSet() bool {
	return v.isSet
}

func (v *NullableFreeShippingPromotion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFreeShippingPromotion(val *FreeShippingPromotion) *NullableFreeShippingPromotion {
	return &NullableFreeShippingPromotion{value: val, isSet: true}
}

func (v NullableFreeShippingPromotion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFreeShippingPromotion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
