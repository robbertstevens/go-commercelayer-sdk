/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.3.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the PriceCreateDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriceCreateDataRelationships{}

// PriceCreateDataRelationships struct for PriceCreateDataRelationships
type PriceCreateDataRelationships struct {
	PriceList  MarketCreateDataRelationshipsPriceList        `json:"price_list"`
	Sku        InStockSubscriptionCreateDataRelationshipsSku `json:"sku"`
	PriceTiers *PriceCreateDataRelationshipsPriceTiers       `json:"price_tiers,omitempty"`
}

// NewPriceCreateDataRelationships instantiates a new PriceCreateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceCreateDataRelationships(priceList MarketCreateDataRelationshipsPriceList, sku InStockSubscriptionCreateDataRelationshipsSku) *PriceCreateDataRelationships {
	this := PriceCreateDataRelationships{}
	this.PriceList = priceList
	this.Sku = sku
	return &this
}

// NewPriceCreateDataRelationshipsWithDefaults instantiates a new PriceCreateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceCreateDataRelationshipsWithDefaults() *PriceCreateDataRelationships {
	this := PriceCreateDataRelationships{}
	return &this
}

// GetPriceList returns the PriceList field value
func (o *PriceCreateDataRelationships) GetPriceList() MarketCreateDataRelationshipsPriceList {
	if o == nil {
		var ret MarketCreateDataRelationshipsPriceList
		return ret
	}

	return o.PriceList
}

// GetPriceListOk returns a tuple with the PriceList field value
// and a boolean to check if the value has been set.
func (o *PriceCreateDataRelationships) GetPriceListOk() (*MarketCreateDataRelationshipsPriceList, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PriceList, true
}

// SetPriceList sets field value
func (o *PriceCreateDataRelationships) SetPriceList(v MarketCreateDataRelationshipsPriceList) {
	o.PriceList = v
}

// GetSku returns the Sku field value
func (o *PriceCreateDataRelationships) GetSku() InStockSubscriptionCreateDataRelationshipsSku {
	if o == nil {
		var ret InStockSubscriptionCreateDataRelationshipsSku
		return ret
	}

	return o.Sku
}

// GetSkuOk returns a tuple with the Sku field value
// and a boolean to check if the value has been set.
func (o *PriceCreateDataRelationships) GetSkuOk() (*InStockSubscriptionCreateDataRelationshipsSku, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sku, true
}

// SetSku sets field value
func (o *PriceCreateDataRelationships) SetSku(v InStockSubscriptionCreateDataRelationshipsSku) {
	o.Sku = v
}

// GetPriceTiers returns the PriceTiers field value if set, zero value otherwise.
func (o *PriceCreateDataRelationships) GetPriceTiers() PriceCreateDataRelationshipsPriceTiers {
	if o == nil || IsNil(o.PriceTiers) {
		var ret PriceCreateDataRelationshipsPriceTiers
		return ret
	}
	return *o.PriceTiers
}

// GetPriceTiersOk returns a tuple with the PriceTiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceCreateDataRelationships) GetPriceTiersOk() (*PriceCreateDataRelationshipsPriceTiers, bool) {
	if o == nil || IsNil(o.PriceTiers) {
		return nil, false
	}
	return o.PriceTiers, true
}

// HasPriceTiers returns a boolean if a field has been set.
func (o *PriceCreateDataRelationships) HasPriceTiers() bool {
	if o != nil && !IsNil(o.PriceTiers) {
		return true
	}

	return false
}

// SetPriceTiers gets a reference to the given PriceCreateDataRelationshipsPriceTiers and assigns it to the PriceTiers field.
func (o *PriceCreateDataRelationships) SetPriceTiers(v PriceCreateDataRelationshipsPriceTiers) {
	o.PriceTiers = &v
}

func (o PriceCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PriceCreateDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["price_list"] = o.PriceList
	toSerialize["sku"] = o.Sku
	if !IsNil(o.PriceTiers) {
		toSerialize["price_tiers"] = o.PriceTiers
	}
	return toSerialize, nil
}

type NullablePriceCreateDataRelationships struct {
	value *PriceCreateDataRelationships
	isSet bool
}

func (v NullablePriceCreateDataRelationships) Get() *PriceCreateDataRelationships {
	return v.value
}

func (v *NullablePriceCreateDataRelationships) Set(val *PriceCreateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceCreateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceCreateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceCreateDataRelationships(val *PriceCreateDataRelationships) *NullablePriceCreateDataRelationships {
	return &NullablePriceCreateDataRelationships{value: val, isSet: true}
}

func (v NullablePriceCreateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceCreateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
