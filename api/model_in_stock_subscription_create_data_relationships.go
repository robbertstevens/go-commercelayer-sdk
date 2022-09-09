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

// InStockSubscriptionCreateDataRelationships struct for InStockSubscriptionCreateDataRelationships
type InStockSubscriptionCreateDataRelationships struct {
	Market AvalaraAccountDataRelationshipsMarkets `json:"market"`
	Customer CouponRecipientDataRelationshipsCustomer `json:"customer"`
	Sku BundleDataRelationshipsSkus `json:"sku"`
}

// NewInStockSubscriptionCreateDataRelationships instantiates a new InStockSubscriptionCreateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInStockSubscriptionCreateDataRelationships(market AvalaraAccountDataRelationshipsMarkets, customer CouponRecipientDataRelationshipsCustomer, sku BundleDataRelationshipsSkus) *InStockSubscriptionCreateDataRelationships {
	this := InStockSubscriptionCreateDataRelationships{}
	this.Market = market
	this.Customer = customer
	this.Sku = sku
	return &this
}

// NewInStockSubscriptionCreateDataRelationshipsWithDefaults instantiates a new InStockSubscriptionCreateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInStockSubscriptionCreateDataRelationshipsWithDefaults() *InStockSubscriptionCreateDataRelationships {
	this := InStockSubscriptionCreateDataRelationships{}
	return &this
}

// GetMarket returns the Market field value
func (o *InStockSubscriptionCreateDataRelationships) GetMarket() AvalaraAccountDataRelationshipsMarkets {
	if o == nil {
		var ret AvalaraAccountDataRelationshipsMarkets
		return ret
	}

	return o.Market
}

// GetMarketOk returns a tuple with the Market field value
// and a boolean to check if the value has been set.
func (o *InStockSubscriptionCreateDataRelationships) GetMarketOk() (*AvalaraAccountDataRelationshipsMarkets, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Market, true
}

// SetMarket sets field value
func (o *InStockSubscriptionCreateDataRelationships) SetMarket(v AvalaraAccountDataRelationshipsMarkets) {
	o.Market = v
}

// GetCustomer returns the Customer field value
func (o *InStockSubscriptionCreateDataRelationships) GetCustomer() CouponRecipientDataRelationshipsCustomer {
	if o == nil {
		var ret CouponRecipientDataRelationshipsCustomer
		return ret
	}

	return o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value
// and a boolean to check if the value has been set.
func (o *InStockSubscriptionCreateDataRelationships) GetCustomerOk() (*CouponRecipientDataRelationshipsCustomer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Customer, true
}

// SetCustomer sets field value
func (o *InStockSubscriptionCreateDataRelationships) SetCustomer(v CouponRecipientDataRelationshipsCustomer) {
	o.Customer = v
}

// GetSku returns the Sku field value
func (o *InStockSubscriptionCreateDataRelationships) GetSku() BundleDataRelationshipsSkus {
	if o == nil {
		var ret BundleDataRelationshipsSkus
		return ret
	}

	return o.Sku
}

// GetSkuOk returns a tuple with the Sku field value
// and a boolean to check if the value has been set.
func (o *InStockSubscriptionCreateDataRelationships) GetSkuOk() (*BundleDataRelationshipsSkus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sku, true
}

// SetSku sets field value
func (o *InStockSubscriptionCreateDataRelationships) SetSku(v BundleDataRelationshipsSkus) {
	o.Sku = v
}

func (o InStockSubscriptionCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["market"] = o.Market
	}
	if true {
		toSerialize["customer"] = o.Customer
	}
	if true {
		toSerialize["sku"] = o.Sku
	}
	return json.Marshal(toSerialize)
}

type NullableInStockSubscriptionCreateDataRelationships struct {
	value *InStockSubscriptionCreateDataRelationships
	isSet bool
}

func (v NullableInStockSubscriptionCreateDataRelationships) Get() *InStockSubscriptionCreateDataRelationships {
	return v.value
}

func (v *NullableInStockSubscriptionCreateDataRelationships) Set(val *InStockSubscriptionCreateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableInStockSubscriptionCreateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableInStockSubscriptionCreateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInStockSubscriptionCreateDataRelationships(val *InStockSubscriptionCreateDataRelationships) *NullableInStockSubscriptionCreateDataRelationships {
	return &NullableInStockSubscriptionCreateDataRelationships{value: val, isSet: true}
}

func (v NullableInStockSubscriptionCreateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInStockSubscriptionCreateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


