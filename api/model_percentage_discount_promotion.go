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

// PercentageDiscountPromotion struct for PercentageDiscountPromotion
type PercentageDiscountPromotion struct {
	Data PercentageDiscountPromotionData `json:"data"`
}

// NewPercentageDiscountPromotion instantiates a new PercentageDiscountPromotion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPercentageDiscountPromotion(data PercentageDiscountPromotionData) *PercentageDiscountPromotion {
	this := PercentageDiscountPromotion{}
	this.Data = data
	return &this
}

// NewPercentageDiscountPromotionWithDefaults instantiates a new PercentageDiscountPromotion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPercentageDiscountPromotionWithDefaults() *PercentageDiscountPromotion {
	this := PercentageDiscountPromotion{}
	return &this
}

// GetData returns the Data field value
func (o *PercentageDiscountPromotion) GetData() PercentageDiscountPromotionData {
	if o == nil {
		var ret PercentageDiscountPromotionData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PercentageDiscountPromotion) GetDataOk() (*PercentageDiscountPromotionData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PercentageDiscountPromotion) SetData(v PercentageDiscountPromotionData) {
	o.Data = v
}

func (o PercentageDiscountPromotion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePercentageDiscountPromotion struct {
	value *PercentageDiscountPromotion
	isSet bool
}

func (v NullablePercentageDiscountPromotion) Get() *PercentageDiscountPromotion {
	return v.value
}

func (v *NullablePercentageDiscountPromotion) Set(val *PercentageDiscountPromotion) {
	v.value = val
	v.isSet = true
}

func (v NullablePercentageDiscountPromotion) IsSet() bool {
	return v.isSet
}

func (v *NullablePercentageDiscountPromotion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePercentageDiscountPromotion(val *PercentageDiscountPromotion) *NullablePercentageDiscountPromotion {
	return &NullablePercentageDiscountPromotion{value: val, isSet: true}
}

func (v NullablePercentageDiscountPromotion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePercentageDiscountPromotion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


