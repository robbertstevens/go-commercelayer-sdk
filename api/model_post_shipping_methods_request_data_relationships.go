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

// checks if the POSTShippingMethodsRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTShippingMethodsRequestDataRelationships{}

// POSTShippingMethodsRequestDataRelationships struct for POSTShippingMethodsRequestDataRelationships
type POSTShippingMethodsRequestDataRelationships struct {
	Market              *POSTBillingInfoValidationRulesRequestDataRelationshipsMarket   `json:"market,omitempty"`
	ShippingZone        *POSTShippingMethodsRequestDataRelationshipsShippingZone        `json:"shipping_zone,omitempty"`
	ShippingCategory    *POSTShippingMethodsRequestDataRelationshipsShippingCategory    `json:"shipping_category,omitempty"`
	StockLocation       *POSTDeliveryLeadTimesRequestDataRelationshipsStockLocation     `json:"stock_location,omitempty"`
	ShippingMethodTiers *POSTShippingMethodsRequestDataRelationshipsShippingMethodTiers `json:"shipping_method_tiers,omitempty"`
}

// NewPOSTShippingMethodsRequestDataRelationships instantiates a new POSTShippingMethodsRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTShippingMethodsRequestDataRelationships() *POSTShippingMethodsRequestDataRelationships {
	this := POSTShippingMethodsRequestDataRelationships{}
	return &this
}

// NewPOSTShippingMethodsRequestDataRelationshipsWithDefaults instantiates a new POSTShippingMethodsRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTShippingMethodsRequestDataRelationshipsWithDefaults() *POSTShippingMethodsRequestDataRelationships {
	this := POSTShippingMethodsRequestDataRelationships{}
	return &this
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *POSTShippingMethodsRequestDataRelationships) GetMarket() POSTBillingInfoValidationRulesRequestDataRelationshipsMarket {
	if o == nil || IsNil(o.Market) {
		var ret POSTBillingInfoValidationRulesRequestDataRelationshipsMarket
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTShippingMethodsRequestDataRelationships) GetMarketOk() (*POSTBillingInfoValidationRulesRequestDataRelationshipsMarket, bool) {
	if o == nil || IsNil(o.Market) {
		return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *POSTShippingMethodsRequestDataRelationships) HasMarket() bool {
	if o != nil && !IsNil(o.Market) {
		return true
	}

	return false
}

// SetMarket gets a reference to the given POSTBillingInfoValidationRulesRequestDataRelationshipsMarket and assigns it to the Market field.
func (o *POSTShippingMethodsRequestDataRelationships) SetMarket(v POSTBillingInfoValidationRulesRequestDataRelationshipsMarket) {
	o.Market = &v
}

// GetShippingZone returns the ShippingZone field value if set, zero value otherwise.
func (o *POSTShippingMethodsRequestDataRelationships) GetShippingZone() POSTShippingMethodsRequestDataRelationshipsShippingZone {
	if o == nil || IsNil(o.ShippingZone) {
		var ret POSTShippingMethodsRequestDataRelationshipsShippingZone
		return ret
	}
	return *o.ShippingZone
}

// GetShippingZoneOk returns a tuple with the ShippingZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTShippingMethodsRequestDataRelationships) GetShippingZoneOk() (*POSTShippingMethodsRequestDataRelationshipsShippingZone, bool) {
	if o == nil || IsNil(o.ShippingZone) {
		return nil, false
	}
	return o.ShippingZone, true
}

// HasShippingZone returns a boolean if a field has been set.
func (o *POSTShippingMethodsRequestDataRelationships) HasShippingZone() bool {
	if o != nil && !IsNil(o.ShippingZone) {
		return true
	}

	return false
}

// SetShippingZone gets a reference to the given POSTShippingMethodsRequestDataRelationshipsShippingZone and assigns it to the ShippingZone field.
func (o *POSTShippingMethodsRequestDataRelationships) SetShippingZone(v POSTShippingMethodsRequestDataRelationshipsShippingZone) {
	o.ShippingZone = &v
}

// GetShippingCategory returns the ShippingCategory field value if set, zero value otherwise.
func (o *POSTShippingMethodsRequestDataRelationships) GetShippingCategory() POSTShippingMethodsRequestDataRelationshipsShippingCategory {
	if o == nil || IsNil(o.ShippingCategory) {
		var ret POSTShippingMethodsRequestDataRelationshipsShippingCategory
		return ret
	}
	return *o.ShippingCategory
}

// GetShippingCategoryOk returns a tuple with the ShippingCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTShippingMethodsRequestDataRelationships) GetShippingCategoryOk() (*POSTShippingMethodsRequestDataRelationshipsShippingCategory, bool) {
	if o == nil || IsNil(o.ShippingCategory) {
		return nil, false
	}
	return o.ShippingCategory, true
}

// HasShippingCategory returns a boolean if a field has been set.
func (o *POSTShippingMethodsRequestDataRelationships) HasShippingCategory() bool {
	if o != nil && !IsNil(o.ShippingCategory) {
		return true
	}

	return false
}

// SetShippingCategory gets a reference to the given POSTShippingMethodsRequestDataRelationshipsShippingCategory and assigns it to the ShippingCategory field.
func (o *POSTShippingMethodsRequestDataRelationships) SetShippingCategory(v POSTShippingMethodsRequestDataRelationshipsShippingCategory) {
	o.ShippingCategory = &v
}

// GetStockLocation returns the StockLocation field value if set, zero value otherwise.
func (o *POSTShippingMethodsRequestDataRelationships) GetStockLocation() POSTDeliveryLeadTimesRequestDataRelationshipsStockLocation {
	if o == nil || IsNil(o.StockLocation) {
		var ret POSTDeliveryLeadTimesRequestDataRelationshipsStockLocation
		return ret
	}
	return *o.StockLocation
}

// GetStockLocationOk returns a tuple with the StockLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTShippingMethodsRequestDataRelationships) GetStockLocationOk() (*POSTDeliveryLeadTimesRequestDataRelationshipsStockLocation, bool) {
	if o == nil || IsNil(o.StockLocation) {
		return nil, false
	}
	return o.StockLocation, true
}

// HasStockLocation returns a boolean if a field has been set.
func (o *POSTShippingMethodsRequestDataRelationships) HasStockLocation() bool {
	if o != nil && !IsNil(o.StockLocation) {
		return true
	}

	return false
}

// SetStockLocation gets a reference to the given POSTDeliveryLeadTimesRequestDataRelationshipsStockLocation and assigns it to the StockLocation field.
func (o *POSTShippingMethodsRequestDataRelationships) SetStockLocation(v POSTDeliveryLeadTimesRequestDataRelationshipsStockLocation) {
	o.StockLocation = &v
}

// GetShippingMethodTiers returns the ShippingMethodTiers field value if set, zero value otherwise.
func (o *POSTShippingMethodsRequestDataRelationships) GetShippingMethodTiers() POSTShippingMethodsRequestDataRelationshipsShippingMethodTiers {
	if o == nil || IsNil(o.ShippingMethodTiers) {
		var ret POSTShippingMethodsRequestDataRelationshipsShippingMethodTiers
		return ret
	}
	return *o.ShippingMethodTiers
}

// GetShippingMethodTiersOk returns a tuple with the ShippingMethodTiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTShippingMethodsRequestDataRelationships) GetShippingMethodTiersOk() (*POSTShippingMethodsRequestDataRelationshipsShippingMethodTiers, bool) {
	if o == nil || IsNil(o.ShippingMethodTiers) {
		return nil, false
	}
	return o.ShippingMethodTiers, true
}

// HasShippingMethodTiers returns a boolean if a field has been set.
func (o *POSTShippingMethodsRequestDataRelationships) HasShippingMethodTiers() bool {
	if o != nil && !IsNil(o.ShippingMethodTiers) {
		return true
	}

	return false
}

// SetShippingMethodTiers gets a reference to the given POSTShippingMethodsRequestDataRelationshipsShippingMethodTiers and assigns it to the ShippingMethodTiers field.
func (o *POSTShippingMethodsRequestDataRelationships) SetShippingMethodTiers(v POSTShippingMethodsRequestDataRelationshipsShippingMethodTiers) {
	o.ShippingMethodTiers = &v
}

func (o POSTShippingMethodsRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTShippingMethodsRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Market) {
		toSerialize["market"] = o.Market
	}
	if !IsNil(o.ShippingZone) {
		toSerialize["shipping_zone"] = o.ShippingZone
	}
	if !IsNil(o.ShippingCategory) {
		toSerialize["shipping_category"] = o.ShippingCategory
	}
	if !IsNil(o.StockLocation) {
		toSerialize["stock_location"] = o.StockLocation
	}
	if !IsNil(o.ShippingMethodTiers) {
		toSerialize["shipping_method_tiers"] = o.ShippingMethodTiers
	}
	return toSerialize, nil
}

type NullablePOSTShippingMethodsRequestDataRelationships struct {
	value *POSTShippingMethodsRequestDataRelationships
	isSet bool
}

func (v NullablePOSTShippingMethodsRequestDataRelationships) Get() *POSTShippingMethodsRequestDataRelationships {
	return v.value
}

func (v *NullablePOSTShippingMethodsRequestDataRelationships) Set(val *POSTShippingMethodsRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTShippingMethodsRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTShippingMethodsRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTShippingMethodsRequestDataRelationships(val *POSTShippingMethodsRequestDataRelationships) *NullablePOSTShippingMethodsRequestDataRelationships {
	return &NullablePOSTShippingMethodsRequestDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTShippingMethodsRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTShippingMethodsRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
