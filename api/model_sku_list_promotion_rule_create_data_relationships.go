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

// SkuListPromotionRuleCreateDataRelationships struct for SkuListPromotionRuleCreateDataRelationships
type SkuListPromotionRuleCreateDataRelationships struct {
	Promotion CouponCodesPromotionRuleDataRelationshipsPromotion `json:"promotion"`
	SkuList   *BundleDataRelationshipsSkuList                    `json:"sku_list,omitempty"`
}

// NewSkuListPromotionRuleCreateDataRelationships instantiates a new SkuListPromotionRuleCreateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuListPromotionRuleCreateDataRelationships(promotion CouponCodesPromotionRuleDataRelationshipsPromotion) *SkuListPromotionRuleCreateDataRelationships {
	this := SkuListPromotionRuleCreateDataRelationships{}
	this.Promotion = promotion
	return &this
}

// NewSkuListPromotionRuleCreateDataRelationshipsWithDefaults instantiates a new SkuListPromotionRuleCreateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuListPromotionRuleCreateDataRelationshipsWithDefaults() *SkuListPromotionRuleCreateDataRelationships {
	this := SkuListPromotionRuleCreateDataRelationships{}
	return &this
}

// GetPromotion returns the Promotion field value
func (o *SkuListPromotionRuleCreateDataRelationships) GetPromotion() CouponCodesPromotionRuleDataRelationshipsPromotion {
	if o == nil {
		var ret CouponCodesPromotionRuleDataRelationshipsPromotion
		return ret
	}

	return o.Promotion
}

// GetPromotionOk returns a tuple with the Promotion field value
// and a boolean to check if the value has been set.
func (o *SkuListPromotionRuleCreateDataRelationships) GetPromotionOk() (*CouponCodesPromotionRuleDataRelationshipsPromotion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Promotion, true
}

// SetPromotion sets field value
func (o *SkuListPromotionRuleCreateDataRelationships) SetPromotion(v CouponCodesPromotionRuleDataRelationshipsPromotion) {
	o.Promotion = v
}

// GetSkuList returns the SkuList field value if set, zero value otherwise.
func (o *SkuListPromotionRuleCreateDataRelationships) GetSkuList() BundleDataRelationshipsSkuList {
	if o == nil || o.SkuList == nil {
		var ret BundleDataRelationshipsSkuList
		return ret
	}
	return *o.SkuList
}

// GetSkuListOk returns a tuple with the SkuList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuListPromotionRuleCreateDataRelationships) GetSkuListOk() (*BundleDataRelationshipsSkuList, bool) {
	if o == nil || o.SkuList == nil {
		return nil, false
	}
	return o.SkuList, true
}

// HasSkuList returns a boolean if a field has been set.
func (o *SkuListPromotionRuleCreateDataRelationships) HasSkuList() bool {
	if o != nil && o.SkuList != nil {
		return true
	}

	return false
}

// SetSkuList gets a reference to the given BundleDataRelationshipsSkuList and assigns it to the SkuList field.
func (o *SkuListPromotionRuleCreateDataRelationships) SetSkuList(v BundleDataRelationshipsSkuList) {
	o.SkuList = &v
}

func (o SkuListPromotionRuleCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["promotion"] = o.Promotion
	}
	if o.SkuList != nil {
		toSerialize["sku_list"] = o.SkuList
	}
	return json.Marshal(toSerialize)
}

type NullableSkuListPromotionRuleCreateDataRelationships struct {
	value *SkuListPromotionRuleCreateDataRelationships
	isSet bool
}

func (v NullableSkuListPromotionRuleCreateDataRelationships) Get() *SkuListPromotionRuleCreateDataRelationships {
	return v.value
}

func (v *NullableSkuListPromotionRuleCreateDataRelationships) Set(val *SkuListPromotionRuleCreateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuListPromotionRuleCreateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuListPromotionRuleCreateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuListPromotionRuleCreateDataRelationships(val *SkuListPromotionRuleCreateDataRelationships) *NullableSkuListPromotionRuleCreateDataRelationships {
	return &NullableSkuListPromotionRuleCreateDataRelationships{value: val, isSet: true}
}

func (v NullableSkuListPromotionRuleCreateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkuListPromotionRuleCreateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
