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

// GETFixedPricePromotions200ResponseDataInnerRelationships struct for GETFixedPricePromotions200ResponseDataInnerRelationships
type GETFixedPricePromotions200ResponseDataInnerRelationships struct {
	Market                   *GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket           `json:"market,omitempty"`
	PromotionRules           *GETExternalPromotions200ResponseDataInnerRelationshipsPromotionRules           `json:"promotion_rules,omitempty"`
	OrderAmountPromotionRule *GETExternalPromotions200ResponseDataInnerRelationshipsOrderAmountPromotionRule `json:"order_amount_promotion_rule,omitempty"`
	SkuListPromotionRule     *GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule     `json:"sku_list_promotion_rule,omitempty"`
	CouponCodesPromotionRule *GETExternalPromotions200ResponseDataInnerRelationshipsCouponCodesPromotionRule `json:"coupon_codes_promotion_rule,omitempty"`
	Attachments              *GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments                 `json:"attachments,omitempty"`
	SkuList                  *GETBundles200ResponseDataInnerRelationshipsSkuList                             `json:"sku_list,omitempty"`
	Skus                     *GETBundles200ResponseDataInnerRelationshipsSkus                                `json:"skus,omitempty"`
}

// NewGETFixedPricePromotions200ResponseDataInnerRelationships instantiates a new GETFixedPricePromotions200ResponseDataInnerRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETFixedPricePromotions200ResponseDataInnerRelationships() *GETFixedPricePromotions200ResponseDataInnerRelationships {
	this := GETFixedPricePromotions200ResponseDataInnerRelationships{}
	return &this
}

// NewGETFixedPricePromotions200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETFixedPricePromotions200ResponseDataInnerRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETFixedPricePromotions200ResponseDataInnerRelationshipsWithDefaults() *GETFixedPricePromotions200ResponseDataInnerRelationships {
	this := GETFixedPricePromotions200ResponseDataInnerRelationships{}
	return &this
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *GETFixedPricePromotions200ResponseDataInnerRelationships) GetMarket() GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket {
	if o == nil || o.Market == nil {
		var ret GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETFixedPricePromotions200ResponseDataInnerRelationships) GetMarketOk() (*GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket, bool) {
	if o == nil || o.Market == nil {
		return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *GETFixedPricePromotions200ResponseDataInnerRelationships) HasMarket() bool {
	if o != nil && o.Market != nil {
		return true
	}

	return false
}

// SetMarket gets a reference to the given GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket and assigns it to the Market field.
func (o *GETFixedPricePromotions200ResponseDataInnerRelationships) SetMarket(v GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket) {
	o.Market = &v
}

// GetPromotionRules returns the PromotionRules field value if set, zero value otherwise.
func (o *GETFixedPricePromotions200ResponseDataInnerRelationships) GetPromotionRules() GETExternalPromotions200ResponseDataInnerRelationshipsPromotionRules {
	if o == nil || o.PromotionRules == nil {
		var ret GETExternalPromotions200ResponseDataInnerRelationshipsPromotionRules
		return ret
	}
	return *o.PromotionRules
}

// GetPromotionRulesOk returns a tuple with the PromotionRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETFixedPricePromotions200ResponseDataInnerRelationships) GetPromotionRulesOk() (*GETExternalPromotions200ResponseDataInnerRelationshipsPromotionRules, bool) {
	if o == nil || o.PromotionRules == nil {
		return nil, false
	}
	return o.PromotionRules, true
}

// HasPromotionRules returns a boolean if a field has been set.
func (o *GETFixedPricePromotions200ResponseDataInnerRelationships) HasPromotionRules() bool {
	if o != nil && o.PromotionRules != nil {
		return true
	}

	return false
}

// SetPromotionRules gets a reference to the given GETExternalPromotions200ResponseDataInnerRelationshipsPromotionRules and assigns it to the PromotionRules field.
func (o *GETFixedPricePromotions200ResponseDataInnerRelationships) SetPromotionRules(v GETExternalPromotions200ResponseDataInnerRelationshipsPromotionRules) {
	o.PromotionRules = &v
}

// GetOrderAmountPromotionRule returns the OrderAmountPromotionRule field value if set, zero value otherwise.
func (o *GETFixedPricePromotions200ResponseDataInnerRelationships) GetOrderAmountPromotionRule() GETExternalPromotions200ResponseDataInnerRelationshipsOrderAmountPromotionRule {
	if o == nil || o.OrderAmountPromotionRule == nil {
		var ret GETExternalPromotions200ResponseDataInnerRelationshipsOrderAmountPromotionRule
		return ret
	}
	return *o.OrderAmountPromotionRule
}

// GetOrderAmountPromotionRuleOk returns a tuple with the OrderAmountPromotionRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETFixedPricePromotions200ResponseDataInnerRelationships) GetOrderAmountPromotionRuleOk() (*GETExternalPromotions200ResponseDataInnerRelationshipsOrderAmountPromotionRule, bool) {
	if o == nil || o.OrderAmountPromotionRule == nil {
		return nil, false
	}
	return o.OrderAmountPromotionRule, true
}

// HasOrderAmountPromotionRule returns a boolean if a field has been set.
func (o *GETFixedPricePromotions200ResponseDataInnerRelationships) HasOrderAmountPromotionRule() bool {
	if o != nil && o.OrderAmountPromotionRule != nil {
		return true
	}

	return false
}

// SetOrderAmountPromotionRule gets a reference to the given GETExternalPromotions200ResponseDataInnerRelationshipsOrderAmountPromotionRule and assigns it to the OrderAmountPromotionRule field.
func (o *GETFixedPricePromotions200ResponseDataInnerRelationships) SetOrderAmountPromotionRule(v GETExternalPromotions200ResponseDataInnerRelationshipsOrderAmountPromotionRule) {
	o.OrderAmountPromotionRule = &v
}

// GetSkuListPromotionRule returns the SkuListPromotionRule field value if set, zero value otherwise.
func (o *GETFixedPricePromotions200ResponseDataInnerRelationships) GetSkuListPromotionRule() GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule {
	if o == nil || o.SkuListPromotionRule == nil {
		var ret GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule
		return ret
	}
	return *o.SkuListPromotionRule
}

// GetSkuListPromotionRuleOk returns a tuple with the SkuListPromotionRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETFixedPricePromotions200ResponseDataInnerRelationships) GetSkuListPromotionRuleOk() (*GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule, bool) {
	if o == nil || o.SkuListPromotionRule == nil {
		return nil, false
	}
	return o.SkuListPromotionRule, true
}

// HasSkuListPromotionRule returns a boolean if a field has been set.
func (o *GETFixedPricePromotions200ResponseDataInnerRelationships) HasSkuListPromotionRule() bool {
	if o != nil && o.SkuListPromotionRule != nil {
		return true
	}

	return false
}

// SetSkuListPromotionRule gets a reference to the given GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule and assigns it to the SkuListPromotionRule field.
func (o *GETFixedPricePromotions200ResponseDataInnerRelationships) SetSkuListPromotionRule(v GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule) {
	o.SkuListPromotionRule = &v
}

// GetCouponCodesPromotionRule returns the CouponCodesPromotionRule field value if set, zero value otherwise.
func (o *GETFixedPricePromotions200ResponseDataInnerRelationships) GetCouponCodesPromotionRule() GETExternalPromotions200ResponseDataInnerRelationshipsCouponCodesPromotionRule {
	if o == nil || o.CouponCodesPromotionRule == nil {
		var ret GETExternalPromotions200ResponseDataInnerRelationshipsCouponCodesPromotionRule
		return ret
	}
	return *o.CouponCodesPromotionRule
}

// GetCouponCodesPromotionRuleOk returns a tuple with the CouponCodesPromotionRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETFixedPricePromotions200ResponseDataInnerRelationships) GetCouponCodesPromotionRuleOk() (*GETExternalPromotions200ResponseDataInnerRelationshipsCouponCodesPromotionRule, bool) {
	if o == nil || o.CouponCodesPromotionRule == nil {
		return nil, false
	}
	return o.CouponCodesPromotionRule, true
}

// HasCouponCodesPromotionRule returns a boolean if a field has been set.
func (o *GETFixedPricePromotions200ResponseDataInnerRelationships) HasCouponCodesPromotionRule() bool {
	if o != nil && o.CouponCodesPromotionRule != nil {
		return true
	}

	return false
}

// SetCouponCodesPromotionRule gets a reference to the given GETExternalPromotions200ResponseDataInnerRelationshipsCouponCodesPromotionRule and assigns it to the CouponCodesPromotionRule field.
func (o *GETFixedPricePromotions200ResponseDataInnerRelationships) SetCouponCodesPromotionRule(v GETExternalPromotions200ResponseDataInnerRelationshipsCouponCodesPromotionRule) {
	o.CouponCodesPromotionRule = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *GETFixedPricePromotions200ResponseDataInnerRelationships) GetAttachments() GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments {
	if o == nil || o.Attachments == nil {
		var ret GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETFixedPricePromotions200ResponseDataInnerRelationships) GetAttachmentsOk() (*GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *GETFixedPricePromotions200ResponseDataInnerRelationships) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments and assigns it to the Attachments field.
func (o *GETFixedPricePromotions200ResponseDataInnerRelationships) SetAttachments(v GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments) {
	o.Attachments = &v
}

// GetSkuList returns the SkuList field value if set, zero value otherwise.
func (o *GETFixedPricePromotions200ResponseDataInnerRelationships) GetSkuList() GETBundles200ResponseDataInnerRelationshipsSkuList {
	if o == nil || o.SkuList == nil {
		var ret GETBundles200ResponseDataInnerRelationshipsSkuList
		return ret
	}
	return *o.SkuList
}

// GetSkuListOk returns a tuple with the SkuList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETFixedPricePromotions200ResponseDataInnerRelationships) GetSkuListOk() (*GETBundles200ResponseDataInnerRelationshipsSkuList, bool) {
	if o == nil || o.SkuList == nil {
		return nil, false
	}
	return o.SkuList, true
}

// HasSkuList returns a boolean if a field has been set.
func (o *GETFixedPricePromotions200ResponseDataInnerRelationships) HasSkuList() bool {
	if o != nil && o.SkuList != nil {
		return true
	}

	return false
}

// SetSkuList gets a reference to the given GETBundles200ResponseDataInnerRelationshipsSkuList and assigns it to the SkuList field.
func (o *GETFixedPricePromotions200ResponseDataInnerRelationships) SetSkuList(v GETBundles200ResponseDataInnerRelationshipsSkuList) {
	o.SkuList = &v
}

// GetSkus returns the Skus field value if set, zero value otherwise.
func (o *GETFixedPricePromotions200ResponseDataInnerRelationships) GetSkus() GETBundles200ResponseDataInnerRelationshipsSkus {
	if o == nil || o.Skus == nil {
		var ret GETBundles200ResponseDataInnerRelationshipsSkus
		return ret
	}
	return *o.Skus
}

// GetSkusOk returns a tuple with the Skus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETFixedPricePromotions200ResponseDataInnerRelationships) GetSkusOk() (*GETBundles200ResponseDataInnerRelationshipsSkus, bool) {
	if o == nil || o.Skus == nil {
		return nil, false
	}
	return o.Skus, true
}

// HasSkus returns a boolean if a field has been set.
func (o *GETFixedPricePromotions200ResponseDataInnerRelationships) HasSkus() bool {
	if o != nil && o.Skus != nil {
		return true
	}

	return false
}

// SetSkus gets a reference to the given GETBundles200ResponseDataInnerRelationshipsSkus and assigns it to the Skus field.
func (o *GETFixedPricePromotions200ResponseDataInnerRelationships) SetSkus(v GETBundles200ResponseDataInnerRelationshipsSkus) {
	o.Skus = &v
}

func (o GETFixedPricePromotions200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Market != nil {
		toSerialize["market"] = o.Market
	}
	if o.PromotionRules != nil {
		toSerialize["promotion_rules"] = o.PromotionRules
	}
	if o.OrderAmountPromotionRule != nil {
		toSerialize["order_amount_promotion_rule"] = o.OrderAmountPromotionRule
	}
	if o.SkuListPromotionRule != nil {
		toSerialize["sku_list_promotion_rule"] = o.SkuListPromotionRule
	}
	if o.CouponCodesPromotionRule != nil {
		toSerialize["coupon_codes_promotion_rule"] = o.CouponCodesPromotionRule
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	if o.SkuList != nil {
		toSerialize["sku_list"] = o.SkuList
	}
	if o.Skus != nil {
		toSerialize["skus"] = o.Skus
	}
	return json.Marshal(toSerialize)
}

type NullableGETFixedPricePromotions200ResponseDataInnerRelationships struct {
	value *GETFixedPricePromotions200ResponseDataInnerRelationships
	isSet bool
}

func (v NullableGETFixedPricePromotions200ResponseDataInnerRelationships) Get() *GETFixedPricePromotions200ResponseDataInnerRelationships {
	return v.value
}

func (v *NullableGETFixedPricePromotions200ResponseDataInnerRelationships) Set(val *GETFixedPricePromotions200ResponseDataInnerRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETFixedPricePromotions200ResponseDataInnerRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETFixedPricePromotions200ResponseDataInnerRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETFixedPricePromotions200ResponseDataInnerRelationships(val *GETFixedPricePromotions200ResponseDataInnerRelationships) *NullableGETFixedPricePromotions200ResponseDataInnerRelationships {
	return &NullableGETFixedPricePromotions200ResponseDataInnerRelationships{value: val, isSet: true}
}

func (v NullableGETFixedPricePromotions200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETFixedPricePromotions200ResponseDataInnerRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
