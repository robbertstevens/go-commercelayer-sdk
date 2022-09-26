/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// MarketCreateDataRelationships struct for MarketCreateDataRelationships
type MarketCreateDataRelationships struct {
	Merchant       MarketDataRelationshipsMerchant                        `json:"merchant"`
	PriceList      MarketDataRelationshipsPriceList                       `json:"price_list"`
	InventoryModel InventoryReturnLocationDataRelationshipsInventoryModel `json:"inventory_model"`
	TaxCalculator  *MarketDataRelationshipsTaxCalculator                  `json:"tax_calculator,omitempty"`
	CustomerGroup  *CustomerDataRelationshipsCustomerGroup                `json:"customer_group,omitempty"`
}

// NewMarketCreateDataRelationships instantiates a new MarketCreateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketCreateDataRelationships(merchant MarketDataRelationshipsMerchant, priceList MarketDataRelationshipsPriceList, inventoryModel InventoryReturnLocationDataRelationshipsInventoryModel) *MarketCreateDataRelationships {
	this := MarketCreateDataRelationships{}
	this.Merchant = merchant
	this.PriceList = priceList
	this.InventoryModel = inventoryModel
	return &this
}

// NewMarketCreateDataRelationshipsWithDefaults instantiates a new MarketCreateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketCreateDataRelationshipsWithDefaults() *MarketCreateDataRelationships {
	this := MarketCreateDataRelationships{}
	return &this
}

// GetMerchant returns the Merchant field value
func (o *MarketCreateDataRelationships) GetMerchant() MarketDataRelationshipsMerchant {
	if o == nil {
		var ret MarketDataRelationshipsMerchant
		return ret
	}

	return o.Merchant
}

// GetMerchantOk returns a tuple with the Merchant field value
// and a boolean to check if the value has been set.
func (o *MarketCreateDataRelationships) GetMerchantOk() (*MarketDataRelationshipsMerchant, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Merchant, true
}

// SetMerchant sets field value
func (o *MarketCreateDataRelationships) SetMerchant(v MarketDataRelationshipsMerchant) {
	o.Merchant = v
}

// GetPriceList returns the PriceList field value
func (o *MarketCreateDataRelationships) GetPriceList() MarketDataRelationshipsPriceList {
	if o == nil {
		var ret MarketDataRelationshipsPriceList
		return ret
	}

	return o.PriceList
}

// GetPriceListOk returns a tuple with the PriceList field value
// and a boolean to check if the value has been set.
func (o *MarketCreateDataRelationships) GetPriceListOk() (*MarketDataRelationshipsPriceList, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PriceList, true
}

// SetPriceList sets field value
func (o *MarketCreateDataRelationships) SetPriceList(v MarketDataRelationshipsPriceList) {
	o.PriceList = v
}

// GetInventoryModel returns the InventoryModel field value
func (o *MarketCreateDataRelationships) GetInventoryModel() InventoryReturnLocationDataRelationshipsInventoryModel {
	if o == nil {
		var ret InventoryReturnLocationDataRelationshipsInventoryModel
		return ret
	}

	return o.InventoryModel
}

// GetInventoryModelOk returns a tuple with the InventoryModel field value
// and a boolean to check if the value has been set.
func (o *MarketCreateDataRelationships) GetInventoryModelOk() (*InventoryReturnLocationDataRelationshipsInventoryModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InventoryModel, true
}

// SetInventoryModel sets field value
func (o *MarketCreateDataRelationships) SetInventoryModel(v InventoryReturnLocationDataRelationshipsInventoryModel) {
	o.InventoryModel = v
}

// GetTaxCalculator returns the TaxCalculator field value if set, zero value otherwise.
func (o *MarketCreateDataRelationships) GetTaxCalculator() MarketDataRelationshipsTaxCalculator {
	if o == nil || o.TaxCalculator == nil {
		var ret MarketDataRelationshipsTaxCalculator
		return ret
	}
	return *o.TaxCalculator
}

// GetTaxCalculatorOk returns a tuple with the TaxCalculator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketCreateDataRelationships) GetTaxCalculatorOk() (*MarketDataRelationshipsTaxCalculator, bool) {
	if o == nil || o.TaxCalculator == nil {
		return nil, false
	}
	return o.TaxCalculator, true
}

// HasTaxCalculator returns a boolean if a field has been set.
func (o *MarketCreateDataRelationships) HasTaxCalculator() bool {
	if o != nil && o.TaxCalculator != nil {
		return true
	}

	return false
}

// SetTaxCalculator gets a reference to the given MarketDataRelationshipsTaxCalculator and assigns it to the TaxCalculator field.
func (o *MarketCreateDataRelationships) SetTaxCalculator(v MarketDataRelationshipsTaxCalculator) {
	o.TaxCalculator = &v
}

// GetCustomerGroup returns the CustomerGroup field value if set, zero value otherwise.
func (o *MarketCreateDataRelationships) GetCustomerGroup() CustomerDataRelationshipsCustomerGroup {
	if o == nil || o.CustomerGroup == nil {
		var ret CustomerDataRelationshipsCustomerGroup
		return ret
	}
	return *o.CustomerGroup
}

// GetCustomerGroupOk returns a tuple with the CustomerGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketCreateDataRelationships) GetCustomerGroupOk() (*CustomerDataRelationshipsCustomerGroup, bool) {
	if o == nil || o.CustomerGroup == nil {
		return nil, false
	}
	return o.CustomerGroup, true
}

// HasCustomerGroup returns a boolean if a field has been set.
func (o *MarketCreateDataRelationships) HasCustomerGroup() bool {
	if o != nil && o.CustomerGroup != nil {
		return true
	}

	return false
}

// SetCustomerGroup gets a reference to the given CustomerDataRelationshipsCustomerGroup and assigns it to the CustomerGroup field.
func (o *MarketCreateDataRelationships) SetCustomerGroup(v CustomerDataRelationshipsCustomerGroup) {
	o.CustomerGroup = &v
}

func (o MarketCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["merchant"] = o.Merchant
	}
	if true {
		toSerialize["price_list"] = o.PriceList
	}
	if true {
		toSerialize["inventory_model"] = o.InventoryModel
	}
	if o.TaxCalculator != nil {
		toSerialize["tax_calculator"] = o.TaxCalculator
	}
	if o.CustomerGroup != nil {
		toSerialize["customer_group"] = o.CustomerGroup
	}
	return json.Marshal(toSerialize)
}

type NullableMarketCreateDataRelationships struct {
	value *MarketCreateDataRelationships
	isSet bool
}

func (v NullableMarketCreateDataRelationships) Get() *MarketCreateDataRelationships {
	return v.value
}

func (v *NullableMarketCreateDataRelationships) Set(val *MarketCreateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketCreateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketCreateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketCreateDataRelationships(val *MarketCreateDataRelationships) *NullableMarketCreateDataRelationships {
	return &NullableMarketCreateDataRelationships{value: val, isSet: true}
}

func (v NullableMarketCreateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketCreateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
