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

// checks if the PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships{}

// PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships struct for PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships
type PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships struct {
	Market                   *POSTBillingInfoValidationRulesRequestDataRelationshipsMarket           `json:"market,omitempty"`
	PromotionRules           *POSTExternalPromotionsRequestDataRelationshipsPromotionRules           `json:"promotion_rules,omitempty"`
	OrderAmountPromotionRule *POSTExternalPromotionsRequestDataRelationshipsOrderAmountPromotionRule `json:"order_amount_promotion_rule,omitempty"`
	SkuListPromotionRule     *POSTExternalPromotionsRequestDataRelationshipsSkuListPromotionRule     `json:"sku_list_promotion_rule,omitempty"`
	CouponCodesPromotionRule *POSTCouponsRequestDataRelationshipsPromotionRule                       `json:"coupon_codes_promotion_rule,omitempty"`
	SkuList                  *POSTBundlesRequestDataRelationshipsSkuList                             `json:"sku_list,omitempty"`
}

// NewPATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships instantiates a new PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships() *PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships {
	this := PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships{}
	return &this
}

// NewPATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationshipsWithDefaults instantiates a new PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationshipsWithDefaults() *PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships {
	this := PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships{}
	return &this
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships) GetMarket() POSTBillingInfoValidationRulesRequestDataRelationshipsMarket {
	if o == nil || IsNil(o.Market) {
		var ret POSTBillingInfoValidationRulesRequestDataRelationshipsMarket
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships) GetMarketOk() (*POSTBillingInfoValidationRulesRequestDataRelationshipsMarket, bool) {
	if o == nil || IsNil(o.Market) {
		return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships) HasMarket() bool {
	if o != nil && !IsNil(o.Market) {
		return true
	}

	return false
}

// SetMarket gets a reference to the given POSTBillingInfoValidationRulesRequestDataRelationshipsMarket and assigns it to the Market field.
func (o *PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships) SetMarket(v POSTBillingInfoValidationRulesRequestDataRelationshipsMarket) {
	o.Market = &v
}

// GetPromotionRules returns the PromotionRules field value if set, zero value otherwise.
func (o *PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships) GetPromotionRules() POSTExternalPromotionsRequestDataRelationshipsPromotionRules {
	if o == nil || IsNil(o.PromotionRules) {
		var ret POSTExternalPromotionsRequestDataRelationshipsPromotionRules
		return ret
	}
	return *o.PromotionRules
}

// GetPromotionRulesOk returns a tuple with the PromotionRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships) GetPromotionRulesOk() (*POSTExternalPromotionsRequestDataRelationshipsPromotionRules, bool) {
	if o == nil || IsNil(o.PromotionRules) {
		return nil, false
	}
	return o.PromotionRules, true
}

// HasPromotionRules returns a boolean if a field has been set.
func (o *PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships) HasPromotionRules() bool {
	if o != nil && !IsNil(o.PromotionRules) {
		return true
	}

	return false
}

// SetPromotionRules gets a reference to the given POSTExternalPromotionsRequestDataRelationshipsPromotionRules and assigns it to the PromotionRules field.
func (o *PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships) SetPromotionRules(v POSTExternalPromotionsRequestDataRelationshipsPromotionRules) {
	o.PromotionRules = &v
}

// GetOrderAmountPromotionRule returns the OrderAmountPromotionRule field value if set, zero value otherwise.
func (o *PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships) GetOrderAmountPromotionRule() POSTExternalPromotionsRequestDataRelationshipsOrderAmountPromotionRule {
	if o == nil || IsNil(o.OrderAmountPromotionRule) {
		var ret POSTExternalPromotionsRequestDataRelationshipsOrderAmountPromotionRule
		return ret
	}
	return *o.OrderAmountPromotionRule
}

// GetOrderAmountPromotionRuleOk returns a tuple with the OrderAmountPromotionRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships) GetOrderAmountPromotionRuleOk() (*POSTExternalPromotionsRequestDataRelationshipsOrderAmountPromotionRule, bool) {
	if o == nil || IsNil(o.OrderAmountPromotionRule) {
		return nil, false
	}
	return o.OrderAmountPromotionRule, true
}

// HasOrderAmountPromotionRule returns a boolean if a field has been set.
func (o *PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships) HasOrderAmountPromotionRule() bool {
	if o != nil && !IsNil(o.OrderAmountPromotionRule) {
		return true
	}

	return false
}

// SetOrderAmountPromotionRule gets a reference to the given POSTExternalPromotionsRequestDataRelationshipsOrderAmountPromotionRule and assigns it to the OrderAmountPromotionRule field.
func (o *PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships) SetOrderAmountPromotionRule(v POSTExternalPromotionsRequestDataRelationshipsOrderAmountPromotionRule) {
	o.OrderAmountPromotionRule = &v
}

// GetSkuListPromotionRule returns the SkuListPromotionRule field value if set, zero value otherwise.
func (o *PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships) GetSkuListPromotionRule() POSTExternalPromotionsRequestDataRelationshipsSkuListPromotionRule {
	if o == nil || IsNil(o.SkuListPromotionRule) {
		var ret POSTExternalPromotionsRequestDataRelationshipsSkuListPromotionRule
		return ret
	}
	return *o.SkuListPromotionRule
}

// GetSkuListPromotionRuleOk returns a tuple with the SkuListPromotionRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships) GetSkuListPromotionRuleOk() (*POSTExternalPromotionsRequestDataRelationshipsSkuListPromotionRule, bool) {
	if o == nil || IsNil(o.SkuListPromotionRule) {
		return nil, false
	}
	return o.SkuListPromotionRule, true
}

// HasSkuListPromotionRule returns a boolean if a field has been set.
func (o *PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships) HasSkuListPromotionRule() bool {
	if o != nil && !IsNil(o.SkuListPromotionRule) {
		return true
	}

	return false
}

// SetSkuListPromotionRule gets a reference to the given POSTExternalPromotionsRequestDataRelationshipsSkuListPromotionRule and assigns it to the SkuListPromotionRule field.
func (o *PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships) SetSkuListPromotionRule(v POSTExternalPromotionsRequestDataRelationshipsSkuListPromotionRule) {
	o.SkuListPromotionRule = &v
}

// GetCouponCodesPromotionRule returns the CouponCodesPromotionRule field value if set, zero value otherwise.
func (o *PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships) GetCouponCodesPromotionRule() POSTCouponsRequestDataRelationshipsPromotionRule {
	if o == nil || IsNil(o.CouponCodesPromotionRule) {
		var ret POSTCouponsRequestDataRelationshipsPromotionRule
		return ret
	}
	return *o.CouponCodesPromotionRule
}

// GetCouponCodesPromotionRuleOk returns a tuple with the CouponCodesPromotionRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships) GetCouponCodesPromotionRuleOk() (*POSTCouponsRequestDataRelationshipsPromotionRule, bool) {
	if o == nil || IsNil(o.CouponCodesPromotionRule) {
		return nil, false
	}
	return o.CouponCodesPromotionRule, true
}

// HasCouponCodesPromotionRule returns a boolean if a field has been set.
func (o *PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships) HasCouponCodesPromotionRule() bool {
	if o != nil && !IsNil(o.CouponCodesPromotionRule) {
		return true
	}

	return false
}

// SetCouponCodesPromotionRule gets a reference to the given POSTCouponsRequestDataRelationshipsPromotionRule and assigns it to the CouponCodesPromotionRule field.
func (o *PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships) SetCouponCodesPromotionRule(v POSTCouponsRequestDataRelationshipsPromotionRule) {
	o.CouponCodesPromotionRule = &v
}

// GetSkuList returns the SkuList field value if set, zero value otherwise.
func (o *PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships) GetSkuList() POSTBundlesRequestDataRelationshipsSkuList {
	if o == nil || IsNil(o.SkuList) {
		var ret POSTBundlesRequestDataRelationshipsSkuList
		return ret
	}
	return *o.SkuList
}

// GetSkuListOk returns a tuple with the SkuList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships) GetSkuListOk() (*POSTBundlesRequestDataRelationshipsSkuList, bool) {
	if o == nil || IsNil(o.SkuList) {
		return nil, false
	}
	return o.SkuList, true
}

// HasSkuList returns a boolean if a field has been set.
func (o *PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships) HasSkuList() bool {
	if o != nil && !IsNil(o.SkuList) {
		return true
	}

	return false
}

// SetSkuList gets a reference to the given POSTBundlesRequestDataRelationshipsSkuList and assigns it to the SkuList field.
func (o *PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships) SetSkuList(v POSTBundlesRequestDataRelationshipsSkuList) {
	o.SkuList = &v
}

func (o PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Market) {
		toSerialize["market"] = o.Market
	}
	if !IsNil(o.PromotionRules) {
		toSerialize["promotion_rules"] = o.PromotionRules
	}
	if !IsNil(o.OrderAmountPromotionRule) {
		toSerialize["order_amount_promotion_rule"] = o.OrderAmountPromotionRule
	}
	if !IsNil(o.SkuListPromotionRule) {
		toSerialize["sku_list_promotion_rule"] = o.SkuListPromotionRule
	}
	if !IsNil(o.CouponCodesPromotionRule) {
		toSerialize["coupon_codes_promotion_rule"] = o.CouponCodesPromotionRule
	}
	if !IsNil(o.SkuList) {
		toSerialize["sku_list"] = o.SkuList
	}
	return toSerialize, nil
}

type NullablePATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships struct {
	value *PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships
	isSet bool
}

func (v NullablePATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships) Get() *PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships {
	return v.value
}

func (v *NullablePATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships) Set(val *PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships(val *PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships) *NullablePATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships {
	return &NullablePATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships{value: val, isSet: true}
}

func (v NullablePATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
