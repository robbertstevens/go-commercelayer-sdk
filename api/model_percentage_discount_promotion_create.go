/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 4.1.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the PercentageDiscountPromotionCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PercentageDiscountPromotionCreate{}

// PercentageDiscountPromotionCreate struct for PercentageDiscountPromotionCreate
type PercentageDiscountPromotionCreate struct {
	Data POSTPercentageDiscountPromotionsRequestData `json:"data"`
}

// NewPercentageDiscountPromotionCreate instantiates a new PercentageDiscountPromotionCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPercentageDiscountPromotionCreate(data POSTPercentageDiscountPromotionsRequestData) *PercentageDiscountPromotionCreate {
	this := PercentageDiscountPromotionCreate{}
	this.Data = data
	return &this
}

// NewPercentageDiscountPromotionCreateWithDefaults instantiates a new PercentageDiscountPromotionCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPercentageDiscountPromotionCreateWithDefaults() *PercentageDiscountPromotionCreate {
	this := PercentageDiscountPromotionCreate{}
	return &this
}

// GetData returns the Data field value
func (o *PercentageDiscountPromotionCreate) GetData() POSTPercentageDiscountPromotionsRequestData {
	if o == nil {
		var ret POSTPercentageDiscountPromotionsRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PercentageDiscountPromotionCreate) GetDataOk() (*POSTPercentageDiscountPromotionsRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PercentageDiscountPromotionCreate) SetData(v POSTPercentageDiscountPromotionsRequestData) {
	o.Data = v
}

func (o PercentageDiscountPromotionCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PercentageDiscountPromotionCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullablePercentageDiscountPromotionCreate struct {
	value *PercentageDiscountPromotionCreate
	isSet bool
}

func (v NullablePercentageDiscountPromotionCreate) Get() *PercentageDiscountPromotionCreate {
	return v.value
}

func (v *NullablePercentageDiscountPromotionCreate) Set(val *PercentageDiscountPromotionCreate) {
	v.value = val
	v.isSet = true
}

func (v NullablePercentageDiscountPromotionCreate) IsSet() bool {
	return v.isSet
}

func (v *NullablePercentageDiscountPromotionCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePercentageDiscountPromotionCreate(val *PercentageDiscountPromotionCreate) *NullablePercentageDiscountPromotionCreate {
	return &NullablePercentageDiscountPromotionCreate{value: val, isSet: true}
}

func (v NullablePercentageDiscountPromotionCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePercentageDiscountPromotionCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
