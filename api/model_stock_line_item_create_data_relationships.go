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

// checks if the StockLineItemCreateDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StockLineItemCreateDataRelationships{}

// StockLineItemCreateDataRelationships struct for StockLineItemCreateDataRelationships
type StockLineItemCreateDataRelationships struct {
	Shipment  *ParcelCreateDataRelationshipsShipment         `json:"shipment,omitempty"`
	LineItem  *LineItemOptionCreateDataRelationshipsLineItem `json:"line_item,omitempty"`
	StockItem *StockLineItemCreateDataRelationshipsStockItem `json:"stock_item,omitempty"`
	Sku       *InStockSubscriptionCreateDataRelationshipsSku `json:"sku,omitempty"`
}

// NewStockLineItemCreateDataRelationships instantiates a new StockLineItemCreateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStockLineItemCreateDataRelationships() *StockLineItemCreateDataRelationships {
	this := StockLineItemCreateDataRelationships{}
	return &this
}

// NewStockLineItemCreateDataRelationshipsWithDefaults instantiates a new StockLineItemCreateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStockLineItemCreateDataRelationshipsWithDefaults() *StockLineItemCreateDataRelationships {
	this := StockLineItemCreateDataRelationships{}
	return &this
}

// GetShipment returns the Shipment field value if set, zero value otherwise.
func (o *StockLineItemCreateDataRelationships) GetShipment() ParcelCreateDataRelationshipsShipment {
	if o == nil || IsNil(o.Shipment) {
		var ret ParcelCreateDataRelationshipsShipment
		return ret
	}
	return *o.Shipment
}

// GetShipmentOk returns a tuple with the Shipment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockLineItemCreateDataRelationships) GetShipmentOk() (*ParcelCreateDataRelationshipsShipment, bool) {
	if o == nil || IsNil(o.Shipment) {
		return nil, false
	}
	return o.Shipment, true
}

// HasShipment returns a boolean if a field has been set.
func (o *StockLineItemCreateDataRelationships) HasShipment() bool {
	if o != nil && !IsNil(o.Shipment) {
		return true
	}

	return false
}

// SetShipment gets a reference to the given ParcelCreateDataRelationshipsShipment and assigns it to the Shipment field.
func (o *StockLineItemCreateDataRelationships) SetShipment(v ParcelCreateDataRelationshipsShipment) {
	o.Shipment = &v
}

// GetLineItem returns the LineItem field value if set, zero value otherwise.
func (o *StockLineItemCreateDataRelationships) GetLineItem() LineItemOptionCreateDataRelationshipsLineItem {
	if o == nil || IsNil(o.LineItem) {
		var ret LineItemOptionCreateDataRelationshipsLineItem
		return ret
	}
	return *o.LineItem
}

// GetLineItemOk returns a tuple with the LineItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockLineItemCreateDataRelationships) GetLineItemOk() (*LineItemOptionCreateDataRelationshipsLineItem, bool) {
	if o == nil || IsNil(o.LineItem) {
		return nil, false
	}
	return o.LineItem, true
}

// HasLineItem returns a boolean if a field has been set.
func (o *StockLineItemCreateDataRelationships) HasLineItem() bool {
	if o != nil && !IsNil(o.LineItem) {
		return true
	}

	return false
}

// SetLineItem gets a reference to the given LineItemOptionCreateDataRelationshipsLineItem and assigns it to the LineItem field.
func (o *StockLineItemCreateDataRelationships) SetLineItem(v LineItemOptionCreateDataRelationshipsLineItem) {
	o.LineItem = &v
}

// GetStockItem returns the StockItem field value if set, zero value otherwise.
func (o *StockLineItemCreateDataRelationships) GetStockItem() StockLineItemCreateDataRelationshipsStockItem {
	if o == nil || IsNil(o.StockItem) {
		var ret StockLineItemCreateDataRelationshipsStockItem
		return ret
	}
	return *o.StockItem
}

// GetStockItemOk returns a tuple with the StockItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockLineItemCreateDataRelationships) GetStockItemOk() (*StockLineItemCreateDataRelationshipsStockItem, bool) {
	if o == nil || IsNil(o.StockItem) {
		return nil, false
	}
	return o.StockItem, true
}

// HasStockItem returns a boolean if a field has been set.
func (o *StockLineItemCreateDataRelationships) HasStockItem() bool {
	if o != nil && !IsNil(o.StockItem) {
		return true
	}

	return false
}

// SetStockItem gets a reference to the given StockLineItemCreateDataRelationshipsStockItem and assigns it to the StockItem field.
func (o *StockLineItemCreateDataRelationships) SetStockItem(v StockLineItemCreateDataRelationshipsStockItem) {
	o.StockItem = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *StockLineItemCreateDataRelationships) GetSku() InStockSubscriptionCreateDataRelationshipsSku {
	if o == nil || IsNil(o.Sku) {
		var ret InStockSubscriptionCreateDataRelationshipsSku
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockLineItemCreateDataRelationships) GetSkuOk() (*InStockSubscriptionCreateDataRelationshipsSku, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *StockLineItemCreateDataRelationships) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given InStockSubscriptionCreateDataRelationshipsSku and assigns it to the Sku field.
func (o *StockLineItemCreateDataRelationships) SetSku(v InStockSubscriptionCreateDataRelationshipsSku) {
	o.Sku = &v
}

func (o StockLineItemCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StockLineItemCreateDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Shipment) {
		toSerialize["shipment"] = o.Shipment
	}
	if !IsNil(o.LineItem) {
		toSerialize["line_item"] = o.LineItem
	}
	if !IsNil(o.StockItem) {
		toSerialize["stock_item"] = o.StockItem
	}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	return toSerialize, nil
}

type NullableStockLineItemCreateDataRelationships struct {
	value *StockLineItemCreateDataRelationships
	isSet bool
}

func (v NullableStockLineItemCreateDataRelationships) Get() *StockLineItemCreateDataRelationships {
	return v.value
}

func (v *NullableStockLineItemCreateDataRelationships) Set(val *StockLineItemCreateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableStockLineItemCreateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableStockLineItemCreateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStockLineItemCreateDataRelationships(val *StockLineItemCreateDataRelationships) *NullableStockLineItemCreateDataRelationships {
	return &NullableStockLineItemCreateDataRelationships{value: val, isSet: true}
}

func (v NullableStockLineItemCreateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStockLineItemCreateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
