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

// FixedPricePromotionDataRelationships struct for FixedPricePromotionDataRelationships
type FixedPricePromotionDataRelationships struct {
	Market *AvalaraAccountDataRelationshipsMarkets `json:"market,omitempty"`
	PromotionRules *ExternalPromotionDataRelationshipsPromotionRules `json:"promotion_rules,omitempty"`
	OrderAmountPromotionRule *ExternalPromotionDataRelationshipsOrderAmountPromotionRule `json:"order_amount_promotion_rule,omitempty"`
	SkuListPromotionRule *ExternalPromotionDataRelationshipsSkuListPromotionRule `json:"sku_list_promotion_rule,omitempty"`
	CouponCodesPromotionRule *CouponDataRelationshipsPromotionRule `json:"coupon_codes_promotion_rule,omitempty"`
	Attachments *AvalaraAccountDataRelationshipsAttachments `json:"attachments,omitempty"`
	SkuList *BundleDataRelationshipsSkuList `json:"sku_list,omitempty"`
	Skus *BundleDataRelationshipsSkus `json:"skus,omitempty"`
}

// NewFixedPricePromotionDataRelationships instantiates a new FixedPricePromotionDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFixedPricePromotionDataRelationships() *FixedPricePromotionDataRelationships {
	this := FixedPricePromotionDataRelationships{}
	return &this
}

// NewFixedPricePromotionDataRelationshipsWithDefaults instantiates a new FixedPricePromotionDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFixedPricePromotionDataRelationshipsWithDefaults() *FixedPricePromotionDataRelationships {
	this := FixedPricePromotionDataRelationships{}
	return &this
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *FixedPricePromotionDataRelationships) GetMarket() AvalaraAccountDataRelationshipsMarkets {
	if o == nil || o.Market == nil {
		var ret AvalaraAccountDataRelationshipsMarkets
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FixedPricePromotionDataRelationships) GetMarketOk() (*AvalaraAccountDataRelationshipsMarkets, bool) {
	if o == nil || o.Market == nil {
		return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *FixedPricePromotionDataRelationships) HasMarket() bool {
	if o != nil && o.Market != nil {
		return true
	}

	return false
}

// SetMarket gets a reference to the given AvalaraAccountDataRelationshipsMarkets and assigns it to the Market field.
func (o *FixedPricePromotionDataRelationships) SetMarket(v AvalaraAccountDataRelationshipsMarkets) {
	o.Market = &v
}

// GetPromotionRules returns the PromotionRules field value if set, zero value otherwise.
func (o *FixedPricePromotionDataRelationships) GetPromotionRules() ExternalPromotionDataRelationshipsPromotionRules {
	if o == nil || o.PromotionRules == nil {
		var ret ExternalPromotionDataRelationshipsPromotionRules
		return ret
	}
	return *o.PromotionRules
}

// GetPromotionRulesOk returns a tuple with the PromotionRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FixedPricePromotionDataRelationships) GetPromotionRulesOk() (*ExternalPromotionDataRelationshipsPromotionRules, bool) {
	if o == nil || o.PromotionRules == nil {
		return nil, false
	}
	return o.PromotionRules, true
}

// HasPromotionRules returns a boolean if a field has been set.
func (o *FixedPricePromotionDataRelationships) HasPromotionRules() bool {
	if o != nil && o.PromotionRules != nil {
		return true
	}

	return false
}

// SetPromotionRules gets a reference to the given ExternalPromotionDataRelationshipsPromotionRules and assigns it to the PromotionRules field.
func (o *FixedPricePromotionDataRelationships) SetPromotionRules(v ExternalPromotionDataRelationshipsPromotionRules) {
	o.PromotionRules = &v
}

// GetOrderAmountPromotionRule returns the OrderAmountPromotionRule field value if set, zero value otherwise.
func (o *FixedPricePromotionDataRelationships) GetOrderAmountPromotionRule() ExternalPromotionDataRelationshipsOrderAmountPromotionRule {
	if o == nil || o.OrderAmountPromotionRule == nil {
		var ret ExternalPromotionDataRelationshipsOrderAmountPromotionRule
		return ret
	}
	return *o.OrderAmountPromotionRule
}

// GetOrderAmountPromotionRuleOk returns a tuple with the OrderAmountPromotionRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FixedPricePromotionDataRelationships) GetOrderAmountPromotionRuleOk() (*ExternalPromotionDataRelationshipsOrderAmountPromotionRule, bool) {
	if o == nil || o.OrderAmountPromotionRule == nil {
		return nil, false
	}
	return o.OrderAmountPromotionRule, true
}

// HasOrderAmountPromotionRule returns a boolean if a field has been set.
func (o *FixedPricePromotionDataRelationships) HasOrderAmountPromotionRule() bool {
	if o != nil && o.OrderAmountPromotionRule != nil {
		return true
	}

	return false
}

// SetOrderAmountPromotionRule gets a reference to the given ExternalPromotionDataRelationshipsOrderAmountPromotionRule and assigns it to the OrderAmountPromotionRule field.
func (o *FixedPricePromotionDataRelationships) SetOrderAmountPromotionRule(v ExternalPromotionDataRelationshipsOrderAmountPromotionRule) {
	o.OrderAmountPromotionRule = &v
}

// GetSkuListPromotionRule returns the SkuListPromotionRule field value if set, zero value otherwise.
func (o *FixedPricePromotionDataRelationships) GetSkuListPromotionRule() ExternalPromotionDataRelationshipsSkuListPromotionRule {
	if o == nil || o.SkuListPromotionRule == nil {
		var ret ExternalPromotionDataRelationshipsSkuListPromotionRule
		return ret
	}
	return *o.SkuListPromotionRule
}

// GetSkuListPromotionRuleOk returns a tuple with the SkuListPromotionRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FixedPricePromotionDataRelationships) GetSkuListPromotionRuleOk() (*ExternalPromotionDataRelationshipsSkuListPromotionRule, bool) {
	if o == nil || o.SkuListPromotionRule == nil {
		return nil, false
	}
	return o.SkuListPromotionRule, true
}

// HasSkuListPromotionRule returns a boolean if a field has been set.
func (o *FixedPricePromotionDataRelationships) HasSkuListPromotionRule() bool {
	if o != nil && o.SkuListPromotionRule != nil {
		return true
	}

	return false
}

// SetSkuListPromotionRule gets a reference to the given ExternalPromotionDataRelationshipsSkuListPromotionRule and assigns it to the SkuListPromotionRule field.
func (o *FixedPricePromotionDataRelationships) SetSkuListPromotionRule(v ExternalPromotionDataRelationshipsSkuListPromotionRule) {
	o.SkuListPromotionRule = &v
}

// GetCouponCodesPromotionRule returns the CouponCodesPromotionRule field value if set, zero value otherwise.
func (o *FixedPricePromotionDataRelationships) GetCouponCodesPromotionRule() CouponDataRelationshipsPromotionRule {
	if o == nil || o.CouponCodesPromotionRule == nil {
		var ret CouponDataRelationshipsPromotionRule
		return ret
	}
	return *o.CouponCodesPromotionRule
}

// GetCouponCodesPromotionRuleOk returns a tuple with the CouponCodesPromotionRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FixedPricePromotionDataRelationships) GetCouponCodesPromotionRuleOk() (*CouponDataRelationshipsPromotionRule, bool) {
	if o == nil || o.CouponCodesPromotionRule == nil {
		return nil, false
	}
	return o.CouponCodesPromotionRule, true
}

// HasCouponCodesPromotionRule returns a boolean if a field has been set.
func (o *FixedPricePromotionDataRelationships) HasCouponCodesPromotionRule() bool {
	if o != nil && o.CouponCodesPromotionRule != nil {
		return true
	}

	return false
}

// SetCouponCodesPromotionRule gets a reference to the given CouponDataRelationshipsPromotionRule and assigns it to the CouponCodesPromotionRule field.
func (o *FixedPricePromotionDataRelationships) SetCouponCodesPromotionRule(v CouponDataRelationshipsPromotionRule) {
	o.CouponCodesPromotionRule = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *FixedPricePromotionDataRelationships) GetAttachments() AvalaraAccountDataRelationshipsAttachments {
	if o == nil || o.Attachments == nil {
		var ret AvalaraAccountDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FixedPricePromotionDataRelationships) GetAttachmentsOk() (*AvalaraAccountDataRelationshipsAttachments, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *FixedPricePromotionDataRelationships) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given AvalaraAccountDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *FixedPricePromotionDataRelationships) SetAttachments(v AvalaraAccountDataRelationshipsAttachments) {
	o.Attachments = &v
}

// GetSkuList returns the SkuList field value if set, zero value otherwise.
func (o *FixedPricePromotionDataRelationships) GetSkuList() BundleDataRelationshipsSkuList {
	if o == nil || o.SkuList == nil {
		var ret BundleDataRelationshipsSkuList
		return ret
	}
	return *o.SkuList
}

// GetSkuListOk returns a tuple with the SkuList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FixedPricePromotionDataRelationships) GetSkuListOk() (*BundleDataRelationshipsSkuList, bool) {
	if o == nil || o.SkuList == nil {
		return nil, false
	}
	return o.SkuList, true
}

// HasSkuList returns a boolean if a field has been set.
func (o *FixedPricePromotionDataRelationships) HasSkuList() bool {
	if o != nil && o.SkuList != nil {
		return true
	}

	return false
}

// SetSkuList gets a reference to the given BundleDataRelationshipsSkuList and assigns it to the SkuList field.
func (o *FixedPricePromotionDataRelationships) SetSkuList(v BundleDataRelationshipsSkuList) {
	o.SkuList = &v
}

// GetSkus returns the Skus field value if set, zero value otherwise.
func (o *FixedPricePromotionDataRelationships) GetSkus() BundleDataRelationshipsSkus {
	if o == nil || o.Skus == nil {
		var ret BundleDataRelationshipsSkus
		return ret
	}
	return *o.Skus
}

// GetSkusOk returns a tuple with the Skus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FixedPricePromotionDataRelationships) GetSkusOk() (*BundleDataRelationshipsSkus, bool) {
	if o == nil || o.Skus == nil {
		return nil, false
	}
	return o.Skus, true
}

// HasSkus returns a boolean if a field has been set.
func (o *FixedPricePromotionDataRelationships) HasSkus() bool {
	if o != nil && o.Skus != nil {
		return true
	}

	return false
}

// SetSkus gets a reference to the given BundleDataRelationshipsSkus and assigns it to the Skus field.
func (o *FixedPricePromotionDataRelationships) SetSkus(v BundleDataRelationshipsSkus) {
	o.Skus = &v
}

func (o FixedPricePromotionDataRelationships) MarshalJSON() ([]byte, error) {
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

type NullableFixedPricePromotionDataRelationships struct {
	value *FixedPricePromotionDataRelationships
	isSet bool
}

func (v NullableFixedPricePromotionDataRelationships) Get() *FixedPricePromotionDataRelationships {
	return v.value
}

func (v *NullableFixedPricePromotionDataRelationships) Set(val *FixedPricePromotionDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableFixedPricePromotionDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableFixedPricePromotionDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFixedPricePromotionDataRelationships(val *FixedPricePromotionDataRelationships) *NullableFixedPricePromotionDataRelationships {
	return &NullableFixedPricePromotionDataRelationships{value: val, isSet: true}
}

func (v NullableFixedPricePromotionDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFixedPricePromotionDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


