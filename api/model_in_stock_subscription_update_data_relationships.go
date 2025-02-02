/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.3.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the InStockSubscriptionUpdateDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InStockSubscriptionUpdateDataRelationships{}

// InStockSubscriptionUpdateDataRelationships struct for InStockSubscriptionUpdateDataRelationships
type InStockSubscriptionUpdateDataRelationships struct {
	Market   *BillingInfoValidationRuleCreateDataRelationshipsMarket `json:"market,omitempty"`
	Customer *CouponRecipientCreateDataRelationshipsCustomer         `json:"customer,omitempty"`
	Sku      *InStockSubscriptionCreateDataRelationshipsSku          `json:"sku,omitempty"`
}

// NewInStockSubscriptionUpdateDataRelationships instantiates a new InStockSubscriptionUpdateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInStockSubscriptionUpdateDataRelationships() *InStockSubscriptionUpdateDataRelationships {
	this := InStockSubscriptionUpdateDataRelationships{}
	return &this
}

// NewInStockSubscriptionUpdateDataRelationshipsWithDefaults instantiates a new InStockSubscriptionUpdateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInStockSubscriptionUpdateDataRelationshipsWithDefaults() *InStockSubscriptionUpdateDataRelationships {
	this := InStockSubscriptionUpdateDataRelationships{}
	return &this
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *InStockSubscriptionUpdateDataRelationships) GetMarket() BillingInfoValidationRuleCreateDataRelationshipsMarket {
	if o == nil || IsNil(o.Market) {
		var ret BillingInfoValidationRuleCreateDataRelationshipsMarket
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InStockSubscriptionUpdateDataRelationships) GetMarketOk() (*BillingInfoValidationRuleCreateDataRelationshipsMarket, bool) {
	if o == nil || IsNil(o.Market) {
		return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *InStockSubscriptionUpdateDataRelationships) HasMarket() bool {
	if o != nil && !IsNil(o.Market) {
		return true
	}

	return false
}

// SetMarket gets a reference to the given BillingInfoValidationRuleCreateDataRelationshipsMarket and assigns it to the Market field.
func (o *InStockSubscriptionUpdateDataRelationships) SetMarket(v BillingInfoValidationRuleCreateDataRelationshipsMarket) {
	o.Market = &v
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *InStockSubscriptionUpdateDataRelationships) GetCustomer() CouponRecipientCreateDataRelationshipsCustomer {
	if o == nil || IsNil(o.Customer) {
		var ret CouponRecipientCreateDataRelationshipsCustomer
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InStockSubscriptionUpdateDataRelationships) GetCustomerOk() (*CouponRecipientCreateDataRelationshipsCustomer, bool) {
	if o == nil || IsNil(o.Customer) {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *InStockSubscriptionUpdateDataRelationships) HasCustomer() bool {
	if o != nil && !IsNil(o.Customer) {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given CouponRecipientCreateDataRelationshipsCustomer and assigns it to the Customer field.
func (o *InStockSubscriptionUpdateDataRelationships) SetCustomer(v CouponRecipientCreateDataRelationshipsCustomer) {
	o.Customer = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *InStockSubscriptionUpdateDataRelationships) GetSku() InStockSubscriptionCreateDataRelationshipsSku {
	if o == nil || IsNil(o.Sku) {
		var ret InStockSubscriptionCreateDataRelationshipsSku
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InStockSubscriptionUpdateDataRelationships) GetSkuOk() (*InStockSubscriptionCreateDataRelationshipsSku, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *InStockSubscriptionUpdateDataRelationships) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given InStockSubscriptionCreateDataRelationshipsSku and assigns it to the Sku field.
func (o *InStockSubscriptionUpdateDataRelationships) SetSku(v InStockSubscriptionCreateDataRelationshipsSku) {
	o.Sku = &v
}

func (o InStockSubscriptionUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InStockSubscriptionUpdateDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Market) {
		toSerialize["market"] = o.Market
	}
	if !IsNil(o.Customer) {
		toSerialize["customer"] = o.Customer
	}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	return toSerialize, nil
}

type NullableInStockSubscriptionUpdateDataRelationships struct {
	value *InStockSubscriptionUpdateDataRelationships
	isSet bool
}

func (v NullableInStockSubscriptionUpdateDataRelationships) Get() *InStockSubscriptionUpdateDataRelationships {
	return v.value
}

func (v *NullableInStockSubscriptionUpdateDataRelationships) Set(val *InStockSubscriptionUpdateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableInStockSubscriptionUpdateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableInStockSubscriptionUpdateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInStockSubscriptionUpdateDataRelationships(val *InStockSubscriptionUpdateDataRelationships) *NullableInStockSubscriptionUpdateDataRelationships {
	return &NullableInStockSubscriptionUpdateDataRelationships{value: val, isSet: true}
}

func (v NullableInStockSubscriptionUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInStockSubscriptionUpdateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
