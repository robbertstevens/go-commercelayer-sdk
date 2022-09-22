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

// StockLineItemDataRelationships struct for StockLineItemDataRelationships
type StockLineItemDataRelationships struct {
	Shipment  *OrderDataRelationshipsShipments         `json:"shipment,omitempty"`
	LineItem  *LineItemOptionDataRelationshipsLineItem `json:"line_item,omitempty"`
	StockItem *SkuDataRelationshipsStockItems          `json:"stock_item,omitempty"`
}

// NewStockLineItemDataRelationships instantiates a new StockLineItemDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStockLineItemDataRelationships() *StockLineItemDataRelationships {
	this := StockLineItemDataRelationships{}
	return &this
}

// NewStockLineItemDataRelationshipsWithDefaults instantiates a new StockLineItemDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStockLineItemDataRelationshipsWithDefaults() *StockLineItemDataRelationships {
	this := StockLineItemDataRelationships{}
	return &this
}

// GetShipment returns the Shipment field value if set, zero value otherwise.
func (o *StockLineItemDataRelationships) GetShipment() OrderDataRelationshipsShipments {
	if o == nil || o.Shipment == nil {
		var ret OrderDataRelationshipsShipments
		return ret
	}
	return *o.Shipment
}

// GetShipmentOk returns a tuple with the Shipment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockLineItemDataRelationships) GetShipmentOk() (*OrderDataRelationshipsShipments, bool) {
	if o == nil || o.Shipment == nil {
		return nil, false
	}
	return o.Shipment, true
}

// HasShipment returns a boolean if a field has been set.
func (o *StockLineItemDataRelationships) HasShipment() bool {
	if o != nil && o.Shipment != nil {
		return true
	}

	return false
}

// SetShipment gets a reference to the given OrderDataRelationshipsShipments and assigns it to the Shipment field.
func (o *StockLineItemDataRelationships) SetShipment(v OrderDataRelationshipsShipments) {
	o.Shipment = &v
}

// GetLineItem returns the LineItem field value if set, zero value otherwise.
func (o *StockLineItemDataRelationships) GetLineItem() LineItemOptionDataRelationshipsLineItem {
	if o == nil || o.LineItem == nil {
		var ret LineItemOptionDataRelationshipsLineItem
		return ret
	}
	return *o.LineItem
}

// GetLineItemOk returns a tuple with the LineItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockLineItemDataRelationships) GetLineItemOk() (*LineItemOptionDataRelationshipsLineItem, bool) {
	if o == nil || o.LineItem == nil {
		return nil, false
	}
	return o.LineItem, true
}

// HasLineItem returns a boolean if a field has been set.
func (o *StockLineItemDataRelationships) HasLineItem() bool {
	if o != nil && o.LineItem != nil {
		return true
	}

	return false
}

// SetLineItem gets a reference to the given LineItemOptionDataRelationshipsLineItem and assigns it to the LineItem field.
func (o *StockLineItemDataRelationships) SetLineItem(v LineItemOptionDataRelationshipsLineItem) {
	o.LineItem = &v
}

// GetStockItem returns the StockItem field value if set, zero value otherwise.
func (o *StockLineItemDataRelationships) GetStockItem() SkuDataRelationshipsStockItems {
	if o == nil || o.StockItem == nil {
		var ret SkuDataRelationshipsStockItems
		return ret
	}
	return *o.StockItem
}

// GetStockItemOk returns a tuple with the StockItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockLineItemDataRelationships) GetStockItemOk() (*SkuDataRelationshipsStockItems, bool) {
	if o == nil || o.StockItem == nil {
		return nil, false
	}
	return o.StockItem, true
}

// HasStockItem returns a boolean if a field has been set.
func (o *StockLineItemDataRelationships) HasStockItem() bool {
	if o != nil && o.StockItem != nil {
		return true
	}

	return false
}

// SetStockItem gets a reference to the given SkuDataRelationshipsStockItems and assigns it to the StockItem field.
func (o *StockLineItemDataRelationships) SetStockItem(v SkuDataRelationshipsStockItems) {
	o.StockItem = &v
}

func (o StockLineItemDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Shipment != nil {
		toSerialize["shipment"] = o.Shipment
	}
	if o.LineItem != nil {
		toSerialize["line_item"] = o.LineItem
	}
	if o.StockItem != nil {
		toSerialize["stock_item"] = o.StockItem
	}
	return json.Marshal(toSerialize)
}

type NullableStockLineItemDataRelationships struct {
	value *StockLineItemDataRelationships
	isSet bool
}

func (v NullableStockLineItemDataRelationships) Get() *StockLineItemDataRelationships {
	return v.value
}

func (v *NullableStockLineItemDataRelationships) Set(val *StockLineItemDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableStockLineItemDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableStockLineItemDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStockLineItemDataRelationships(val *StockLineItemDataRelationships) *NullableStockLineItemDataRelationships {
	return &NullableStockLineItemDataRelationships{value: val, isSet: true}
}

func (v NullableStockLineItemDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStockLineItemDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
