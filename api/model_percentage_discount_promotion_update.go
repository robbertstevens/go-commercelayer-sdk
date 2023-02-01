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

// PercentageDiscountPromotionUpdate struct for PercentageDiscountPromotionUpdate
type PercentageDiscountPromotionUpdate struct {
	Data PercentageDiscountPromotionUpdateData `json:"data"`
}

// NewPercentageDiscountPromotionUpdate instantiates a new PercentageDiscountPromotionUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPercentageDiscountPromotionUpdate(data PercentageDiscountPromotionUpdateData) *PercentageDiscountPromotionUpdate {
	this := PercentageDiscountPromotionUpdate{}
	this.Data = data
	return &this
}

// NewPercentageDiscountPromotionUpdateWithDefaults instantiates a new PercentageDiscountPromotionUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPercentageDiscountPromotionUpdateWithDefaults() *PercentageDiscountPromotionUpdate {
	this := PercentageDiscountPromotionUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *PercentageDiscountPromotionUpdate) GetData() PercentageDiscountPromotionUpdateData {
	if o == nil {
		var ret PercentageDiscountPromotionUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PercentageDiscountPromotionUpdate) GetDataOk() (*PercentageDiscountPromotionUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PercentageDiscountPromotionUpdate) SetData(v PercentageDiscountPromotionUpdateData) {
	o.Data = v
}

func (o PercentageDiscountPromotionUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePercentageDiscountPromotionUpdate struct {
	value *PercentageDiscountPromotionUpdate
	isSet bool
}

func (v NullablePercentageDiscountPromotionUpdate) Get() *PercentageDiscountPromotionUpdate {
	return v.value
}

func (v *NullablePercentageDiscountPromotionUpdate) Set(val *PercentageDiscountPromotionUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullablePercentageDiscountPromotionUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullablePercentageDiscountPromotionUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePercentageDiscountPromotionUpdate(val *PercentageDiscountPromotionUpdate) *NullablePercentageDiscountPromotionUpdate {
	return &NullablePercentageDiscountPromotionUpdate{value: val, isSet: true}
}

func (v NullablePercentageDiscountPromotionUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePercentageDiscountPromotionUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
