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

// CouponCodesPromotionRule struct for CouponCodesPromotionRule
type CouponCodesPromotionRule struct {
	Data CouponCodesPromotionRuleData `json:"data"`
}

// NewCouponCodesPromotionRule instantiates a new CouponCodesPromotionRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCouponCodesPromotionRule(data CouponCodesPromotionRuleData) *CouponCodesPromotionRule {
	this := CouponCodesPromotionRule{}
	this.Data = data
	return &this
}

// NewCouponCodesPromotionRuleWithDefaults instantiates a new CouponCodesPromotionRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCouponCodesPromotionRuleWithDefaults() *CouponCodesPromotionRule {
	this := CouponCodesPromotionRule{}
	return &this
}

// GetData returns the Data field value
func (o *CouponCodesPromotionRule) GetData() CouponCodesPromotionRuleData {
	if o == nil {
		var ret CouponCodesPromotionRuleData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CouponCodesPromotionRule) GetDataOk() (*CouponCodesPromotionRuleData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CouponCodesPromotionRule) SetData(v CouponCodesPromotionRuleData) {
	o.Data = v
}

func (o CouponCodesPromotionRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCouponCodesPromotionRule struct {
	value *CouponCodesPromotionRule
	isSet bool
}

func (v NullableCouponCodesPromotionRule) Get() *CouponCodesPromotionRule {
	return v.value
}

func (v *NullableCouponCodesPromotionRule) Set(val *CouponCodesPromotionRule) {
	v.value = val
	v.isSet = true
}

func (v NullableCouponCodesPromotionRule) IsSet() bool {
	return v.isSet
}

func (v *NullableCouponCodesPromotionRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCouponCodesPromotionRule(val *CouponCodesPromotionRule) *NullableCouponCodesPromotionRule {
	return &NullableCouponCodesPromotionRule{value: val, isSet: true}
}

func (v NullableCouponCodesPromotionRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCouponCodesPromotionRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


