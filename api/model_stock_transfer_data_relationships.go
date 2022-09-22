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

// StockTransferDataRelationships struct for StockTransferDataRelationships
type StockTransferDataRelationships struct {
	Sku                      *BundleDataRelationshipsSkus                    `json:"sku,omitempty"`
	OriginStockLocation      *DeliveryLeadTimeDataRelationshipsStockLocation `json:"origin_stock_location,omitempty"`
	DestinationStockLocation *DeliveryLeadTimeDataRelationshipsStockLocation `json:"destination_stock_location,omitempty"`
	Shipment                 *OrderDataRelationshipsShipments                `json:"shipment,omitempty"`
	LineItem                 *LineItemOptionDataRelationshipsLineItem        `json:"line_item,omitempty"`
	Events                   *CustomerAddressDataRelationshipsEvents         `json:"events,omitempty"`
}

// NewStockTransferDataRelationships instantiates a new StockTransferDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStockTransferDataRelationships() *StockTransferDataRelationships {
	this := StockTransferDataRelationships{}
	return &this
}

// NewStockTransferDataRelationshipsWithDefaults instantiates a new StockTransferDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStockTransferDataRelationshipsWithDefaults() *StockTransferDataRelationships {
	this := StockTransferDataRelationships{}
	return &this
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *StockTransferDataRelationships) GetSku() BundleDataRelationshipsSkus {
	if o == nil || o.Sku == nil {
		var ret BundleDataRelationshipsSkus
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockTransferDataRelationships) GetSkuOk() (*BundleDataRelationshipsSkus, bool) {
	if o == nil || o.Sku == nil {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *StockTransferDataRelationships) HasSku() bool {
	if o != nil && o.Sku != nil {
		return true
	}

	return false
}

// SetSku gets a reference to the given BundleDataRelationshipsSkus and assigns it to the Sku field.
func (o *StockTransferDataRelationships) SetSku(v BundleDataRelationshipsSkus) {
	o.Sku = &v
}

// GetOriginStockLocation returns the OriginStockLocation field value if set, zero value otherwise.
func (o *StockTransferDataRelationships) GetOriginStockLocation() DeliveryLeadTimeDataRelationshipsStockLocation {
	if o == nil || o.OriginStockLocation == nil {
		var ret DeliveryLeadTimeDataRelationshipsStockLocation
		return ret
	}
	return *o.OriginStockLocation
}

// GetOriginStockLocationOk returns a tuple with the OriginStockLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockTransferDataRelationships) GetOriginStockLocationOk() (*DeliveryLeadTimeDataRelationshipsStockLocation, bool) {
	if o == nil || o.OriginStockLocation == nil {
		return nil, false
	}
	return o.OriginStockLocation, true
}

// HasOriginStockLocation returns a boolean if a field has been set.
func (o *StockTransferDataRelationships) HasOriginStockLocation() bool {
	if o != nil && o.OriginStockLocation != nil {
		return true
	}

	return false
}

// SetOriginStockLocation gets a reference to the given DeliveryLeadTimeDataRelationshipsStockLocation and assigns it to the OriginStockLocation field.
func (o *StockTransferDataRelationships) SetOriginStockLocation(v DeliveryLeadTimeDataRelationshipsStockLocation) {
	o.OriginStockLocation = &v
}

// GetDestinationStockLocation returns the DestinationStockLocation field value if set, zero value otherwise.
func (o *StockTransferDataRelationships) GetDestinationStockLocation() DeliveryLeadTimeDataRelationshipsStockLocation {
	if o == nil || o.DestinationStockLocation == nil {
		var ret DeliveryLeadTimeDataRelationshipsStockLocation
		return ret
	}
	return *o.DestinationStockLocation
}

// GetDestinationStockLocationOk returns a tuple with the DestinationStockLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockTransferDataRelationships) GetDestinationStockLocationOk() (*DeliveryLeadTimeDataRelationshipsStockLocation, bool) {
	if o == nil || o.DestinationStockLocation == nil {
		return nil, false
	}
	return o.DestinationStockLocation, true
}

// HasDestinationStockLocation returns a boolean if a field has been set.
func (o *StockTransferDataRelationships) HasDestinationStockLocation() bool {
	if o != nil && o.DestinationStockLocation != nil {
		return true
	}

	return false
}

// SetDestinationStockLocation gets a reference to the given DeliveryLeadTimeDataRelationshipsStockLocation and assigns it to the DestinationStockLocation field.
func (o *StockTransferDataRelationships) SetDestinationStockLocation(v DeliveryLeadTimeDataRelationshipsStockLocation) {
	o.DestinationStockLocation = &v
}

// GetShipment returns the Shipment field value if set, zero value otherwise.
func (o *StockTransferDataRelationships) GetShipment() OrderDataRelationshipsShipments {
	if o == nil || o.Shipment == nil {
		var ret OrderDataRelationshipsShipments
		return ret
	}
	return *o.Shipment
}

// GetShipmentOk returns a tuple with the Shipment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockTransferDataRelationships) GetShipmentOk() (*OrderDataRelationshipsShipments, bool) {
	if o == nil || o.Shipment == nil {
		return nil, false
	}
	return o.Shipment, true
}

// HasShipment returns a boolean if a field has been set.
func (o *StockTransferDataRelationships) HasShipment() bool {
	if o != nil && o.Shipment != nil {
		return true
	}

	return false
}

// SetShipment gets a reference to the given OrderDataRelationshipsShipments and assigns it to the Shipment field.
func (o *StockTransferDataRelationships) SetShipment(v OrderDataRelationshipsShipments) {
	o.Shipment = &v
}

// GetLineItem returns the LineItem field value if set, zero value otherwise.
func (o *StockTransferDataRelationships) GetLineItem() LineItemOptionDataRelationshipsLineItem {
	if o == nil || o.LineItem == nil {
		var ret LineItemOptionDataRelationshipsLineItem
		return ret
	}
	return *o.LineItem
}

// GetLineItemOk returns a tuple with the LineItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockTransferDataRelationships) GetLineItemOk() (*LineItemOptionDataRelationshipsLineItem, bool) {
	if o == nil || o.LineItem == nil {
		return nil, false
	}
	return o.LineItem, true
}

// HasLineItem returns a boolean if a field has been set.
func (o *StockTransferDataRelationships) HasLineItem() bool {
	if o != nil && o.LineItem != nil {
		return true
	}

	return false
}

// SetLineItem gets a reference to the given LineItemOptionDataRelationshipsLineItem and assigns it to the LineItem field.
func (o *StockTransferDataRelationships) SetLineItem(v LineItemOptionDataRelationshipsLineItem) {
	o.LineItem = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *StockTransferDataRelationships) GetEvents() CustomerAddressDataRelationshipsEvents {
	if o == nil || o.Events == nil {
		var ret CustomerAddressDataRelationshipsEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockTransferDataRelationships) GetEventsOk() (*CustomerAddressDataRelationshipsEvents, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *StockTransferDataRelationships) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given CustomerAddressDataRelationshipsEvents and assigns it to the Events field.
func (o *StockTransferDataRelationships) SetEvents(v CustomerAddressDataRelationshipsEvents) {
	o.Events = &v
}

func (o StockTransferDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Sku != nil {
		toSerialize["sku"] = o.Sku
	}
	if o.OriginStockLocation != nil {
		toSerialize["origin_stock_location"] = o.OriginStockLocation
	}
	if o.DestinationStockLocation != nil {
		toSerialize["destination_stock_location"] = o.DestinationStockLocation
	}
	if o.Shipment != nil {
		toSerialize["shipment"] = o.Shipment
	}
	if o.LineItem != nil {
		toSerialize["line_item"] = o.LineItem
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableStockTransferDataRelationships struct {
	value *StockTransferDataRelationships
	isSet bool
}

func (v NullableStockTransferDataRelationships) Get() *StockTransferDataRelationships {
	return v.value
}

func (v *NullableStockTransferDataRelationships) Set(val *StockTransferDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableStockTransferDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableStockTransferDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStockTransferDataRelationships(val *StockTransferDataRelationships) *NullableStockTransferDataRelationships {
	return &NullableStockTransferDataRelationships{value: val, isSet: true}
}

func (v NullableStockTransferDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStockTransferDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
