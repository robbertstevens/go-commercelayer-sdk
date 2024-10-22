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

// checks if the StockItemUpdateDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StockItemUpdateDataRelationships{}

// StockItemUpdateDataRelationships struct for StockItemUpdateDataRelationships
type StockItemUpdateDataRelationships struct {
	StockLocation *DeliveryLeadTimeCreateDataRelationshipsStockLocation `json:"stock_location,omitempty"`
	Sku           *InStockSubscriptionCreateDataRelationshipsSku        `json:"sku,omitempty"`
}

// NewStockItemUpdateDataRelationships instantiates a new StockItemUpdateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStockItemUpdateDataRelationships() *StockItemUpdateDataRelationships {
	this := StockItemUpdateDataRelationships{}
	return &this
}

// NewStockItemUpdateDataRelationshipsWithDefaults instantiates a new StockItemUpdateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStockItemUpdateDataRelationshipsWithDefaults() *StockItemUpdateDataRelationships {
	this := StockItemUpdateDataRelationships{}
	return &this
}

// GetStockLocation returns the StockLocation field value if set, zero value otherwise.
func (o *StockItemUpdateDataRelationships) GetStockLocation() DeliveryLeadTimeCreateDataRelationshipsStockLocation {
	if o == nil || IsNil(o.StockLocation) {
		var ret DeliveryLeadTimeCreateDataRelationshipsStockLocation
		return ret
	}
	return *o.StockLocation
}

// GetStockLocationOk returns a tuple with the StockLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockItemUpdateDataRelationships) GetStockLocationOk() (*DeliveryLeadTimeCreateDataRelationshipsStockLocation, bool) {
	if o == nil || IsNil(o.StockLocation) {
		return nil, false
	}
	return o.StockLocation, true
}

// HasStockLocation returns a boolean if a field has been set.
func (o *StockItemUpdateDataRelationships) HasStockLocation() bool {
	if o != nil && !IsNil(o.StockLocation) {
		return true
	}

	return false
}

// SetStockLocation gets a reference to the given DeliveryLeadTimeCreateDataRelationshipsStockLocation and assigns it to the StockLocation field.
func (o *StockItemUpdateDataRelationships) SetStockLocation(v DeliveryLeadTimeCreateDataRelationshipsStockLocation) {
	o.StockLocation = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *StockItemUpdateDataRelationships) GetSku() InStockSubscriptionCreateDataRelationshipsSku {
	if o == nil || IsNil(o.Sku) {
		var ret InStockSubscriptionCreateDataRelationshipsSku
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockItemUpdateDataRelationships) GetSkuOk() (*InStockSubscriptionCreateDataRelationshipsSku, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *StockItemUpdateDataRelationships) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given InStockSubscriptionCreateDataRelationshipsSku and assigns it to the Sku field.
func (o *StockItemUpdateDataRelationships) SetSku(v InStockSubscriptionCreateDataRelationshipsSku) {
	o.Sku = &v
}

func (o StockItemUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StockItemUpdateDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StockLocation) {
		toSerialize["stock_location"] = o.StockLocation
	}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	return toSerialize, nil
}

type NullableStockItemUpdateDataRelationships struct {
	value *StockItemUpdateDataRelationships
	isSet bool
}

func (v NullableStockItemUpdateDataRelationships) Get() *StockItemUpdateDataRelationships {
	return v.value
}

func (v *NullableStockItemUpdateDataRelationships) Set(val *StockItemUpdateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableStockItemUpdateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableStockItemUpdateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStockItemUpdateDataRelationships(val *StockItemUpdateDataRelationships) *NullableStockItemUpdateDataRelationships {
	return &NullableStockItemUpdateDataRelationships{value: val, isSet: true}
}

func (v NullableStockItemUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStockItemUpdateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
