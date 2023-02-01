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

// OrderAmountPromotionRuleCreateDataRelationships struct for OrderAmountPromotionRuleCreateDataRelationships
type OrderAmountPromotionRuleCreateDataRelationships struct {
	Promotion CouponCodesPromotionRuleCreateDataRelationshipsPromotion `json:"promotion"`
}

// NewOrderAmountPromotionRuleCreateDataRelationships instantiates a new OrderAmountPromotionRuleCreateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderAmountPromotionRuleCreateDataRelationships(promotion CouponCodesPromotionRuleCreateDataRelationshipsPromotion) *OrderAmountPromotionRuleCreateDataRelationships {
	this := OrderAmountPromotionRuleCreateDataRelationships{}
	this.Promotion = promotion
	return &this
}

// NewOrderAmountPromotionRuleCreateDataRelationshipsWithDefaults instantiates a new OrderAmountPromotionRuleCreateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderAmountPromotionRuleCreateDataRelationshipsWithDefaults() *OrderAmountPromotionRuleCreateDataRelationships {
	this := OrderAmountPromotionRuleCreateDataRelationships{}
	return &this
}

// GetPromotion returns the Promotion field value
func (o *OrderAmountPromotionRuleCreateDataRelationships) GetPromotion() CouponCodesPromotionRuleCreateDataRelationshipsPromotion {
	if o == nil {
		var ret CouponCodesPromotionRuleCreateDataRelationshipsPromotion
		return ret
	}

	return o.Promotion
}

// GetPromotionOk returns a tuple with the Promotion field value
// and a boolean to check if the value has been set.
func (o *OrderAmountPromotionRuleCreateDataRelationships) GetPromotionOk() (*CouponCodesPromotionRuleCreateDataRelationshipsPromotion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Promotion, true
}

// SetPromotion sets field value
func (o *OrderAmountPromotionRuleCreateDataRelationships) SetPromotion(v CouponCodesPromotionRuleCreateDataRelationshipsPromotion) {
	o.Promotion = v
}

func (o OrderAmountPromotionRuleCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["promotion"] = o.Promotion
	}
	return json.Marshal(toSerialize)
}

type NullableOrderAmountPromotionRuleCreateDataRelationships struct {
	value *OrderAmountPromotionRuleCreateDataRelationships
	isSet bool
}

func (v NullableOrderAmountPromotionRuleCreateDataRelationships) Get() *OrderAmountPromotionRuleCreateDataRelationships {
	return v.value
}

func (v *NullableOrderAmountPromotionRuleCreateDataRelationships) Set(val *OrderAmountPromotionRuleCreateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderAmountPromotionRuleCreateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderAmountPromotionRuleCreateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderAmountPromotionRuleCreateDataRelationships(val *OrderAmountPromotionRuleCreateDataRelationships) *NullableOrderAmountPromotionRuleCreateDataRelationships {
	return &NullableOrderAmountPromotionRuleCreateDataRelationships{value: val, isSet: true}
}

func (v NullableOrderAmountPromotionRuleCreateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderAmountPromotionRuleCreateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
