/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.4.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// SkuDataRelationships struct for SkuDataRelationships
type SkuDataRelationships struct {
	ShippingCategory  *ShipmentDataRelationshipsShippingCategory  `json:"shipping_category,omitempty"`
	Prices            *PriceListDataRelationshipsPrices           `json:"prices,omitempty"`
	StockItems        *SkuDataRelationshipsStockItems             `json:"stock_items,omitempty"`
	DeliveryLeadTimes *ShipmentDataRelationshipsDeliveryLeadTime  `json:"delivery_lead_times,omitempty"`
	SkuOptions        *LineItemOptionDataRelationshipsSkuOption   `json:"sku_options,omitempty"`
	Attachments       *AvalaraAccountDataRelationshipsAttachments `json:"attachments,omitempty"`
}

// NewSkuDataRelationships instantiates a new SkuDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuDataRelationships() *SkuDataRelationships {
	this := SkuDataRelationships{}
	return &this
}

// NewSkuDataRelationshipsWithDefaults instantiates a new SkuDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuDataRelationshipsWithDefaults() *SkuDataRelationships {
	this := SkuDataRelationships{}
	return &this
}

// GetShippingCategory returns the ShippingCategory field value if set, zero value otherwise.
func (o *SkuDataRelationships) GetShippingCategory() ShipmentDataRelationshipsShippingCategory {
	if o == nil || o.ShippingCategory == nil {
		var ret ShipmentDataRelationshipsShippingCategory
		return ret
	}
	return *o.ShippingCategory
}

// GetShippingCategoryOk returns a tuple with the ShippingCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuDataRelationships) GetShippingCategoryOk() (*ShipmentDataRelationshipsShippingCategory, bool) {
	if o == nil || o.ShippingCategory == nil {
		return nil, false
	}
	return o.ShippingCategory, true
}

// HasShippingCategory returns a boolean if a field has been set.
func (o *SkuDataRelationships) HasShippingCategory() bool {
	if o != nil && o.ShippingCategory != nil {
		return true
	}

	return false
}

// SetShippingCategory gets a reference to the given ShipmentDataRelationshipsShippingCategory and assigns it to the ShippingCategory field.
func (o *SkuDataRelationships) SetShippingCategory(v ShipmentDataRelationshipsShippingCategory) {
	o.ShippingCategory = &v
}

// GetPrices returns the Prices field value if set, zero value otherwise.
func (o *SkuDataRelationships) GetPrices() PriceListDataRelationshipsPrices {
	if o == nil || o.Prices == nil {
		var ret PriceListDataRelationshipsPrices
		return ret
	}
	return *o.Prices
}

// GetPricesOk returns a tuple with the Prices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuDataRelationships) GetPricesOk() (*PriceListDataRelationshipsPrices, bool) {
	if o == nil || o.Prices == nil {
		return nil, false
	}
	return o.Prices, true
}

// HasPrices returns a boolean if a field has been set.
func (o *SkuDataRelationships) HasPrices() bool {
	if o != nil && o.Prices != nil {
		return true
	}

	return false
}

// SetPrices gets a reference to the given PriceListDataRelationshipsPrices and assigns it to the Prices field.
func (o *SkuDataRelationships) SetPrices(v PriceListDataRelationshipsPrices) {
	o.Prices = &v
}

// GetStockItems returns the StockItems field value if set, zero value otherwise.
func (o *SkuDataRelationships) GetStockItems() SkuDataRelationshipsStockItems {
	if o == nil || o.StockItems == nil {
		var ret SkuDataRelationshipsStockItems
		return ret
	}
	return *o.StockItems
}

// GetStockItemsOk returns a tuple with the StockItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuDataRelationships) GetStockItemsOk() (*SkuDataRelationshipsStockItems, bool) {
	if o == nil || o.StockItems == nil {
		return nil, false
	}
	return o.StockItems, true
}

// HasStockItems returns a boolean if a field has been set.
func (o *SkuDataRelationships) HasStockItems() bool {
	if o != nil && o.StockItems != nil {
		return true
	}

	return false
}

// SetStockItems gets a reference to the given SkuDataRelationshipsStockItems and assigns it to the StockItems field.
func (o *SkuDataRelationships) SetStockItems(v SkuDataRelationshipsStockItems) {
	o.StockItems = &v
}

// GetDeliveryLeadTimes returns the DeliveryLeadTimes field value if set, zero value otherwise.
func (o *SkuDataRelationships) GetDeliveryLeadTimes() ShipmentDataRelationshipsDeliveryLeadTime {
	if o == nil || o.DeliveryLeadTimes == nil {
		var ret ShipmentDataRelationshipsDeliveryLeadTime
		return ret
	}
	return *o.DeliveryLeadTimes
}

// GetDeliveryLeadTimesOk returns a tuple with the DeliveryLeadTimes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuDataRelationships) GetDeliveryLeadTimesOk() (*ShipmentDataRelationshipsDeliveryLeadTime, bool) {
	if o == nil || o.DeliveryLeadTimes == nil {
		return nil, false
	}
	return o.DeliveryLeadTimes, true
}

// HasDeliveryLeadTimes returns a boolean if a field has been set.
func (o *SkuDataRelationships) HasDeliveryLeadTimes() bool {
	if o != nil && o.DeliveryLeadTimes != nil {
		return true
	}

	return false
}

// SetDeliveryLeadTimes gets a reference to the given ShipmentDataRelationshipsDeliveryLeadTime and assigns it to the DeliveryLeadTimes field.
func (o *SkuDataRelationships) SetDeliveryLeadTimes(v ShipmentDataRelationshipsDeliveryLeadTime) {
	o.DeliveryLeadTimes = &v
}

// GetSkuOptions returns the SkuOptions field value if set, zero value otherwise.
func (o *SkuDataRelationships) GetSkuOptions() LineItemOptionDataRelationshipsSkuOption {
	if o == nil || o.SkuOptions == nil {
		var ret LineItemOptionDataRelationshipsSkuOption
		return ret
	}
	return *o.SkuOptions
}

// GetSkuOptionsOk returns a tuple with the SkuOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuDataRelationships) GetSkuOptionsOk() (*LineItemOptionDataRelationshipsSkuOption, bool) {
	if o == nil || o.SkuOptions == nil {
		return nil, false
	}
	return o.SkuOptions, true
}

// HasSkuOptions returns a boolean if a field has been set.
func (o *SkuDataRelationships) HasSkuOptions() bool {
	if o != nil && o.SkuOptions != nil {
		return true
	}

	return false
}

// SetSkuOptions gets a reference to the given LineItemOptionDataRelationshipsSkuOption and assigns it to the SkuOptions field.
func (o *SkuDataRelationships) SetSkuOptions(v LineItemOptionDataRelationshipsSkuOption) {
	o.SkuOptions = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *SkuDataRelationships) GetAttachments() AvalaraAccountDataRelationshipsAttachments {
	if o == nil || o.Attachments == nil {
		var ret AvalaraAccountDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuDataRelationships) GetAttachmentsOk() (*AvalaraAccountDataRelationshipsAttachments, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *SkuDataRelationships) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given AvalaraAccountDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *SkuDataRelationships) SetAttachments(v AvalaraAccountDataRelationshipsAttachments) {
	o.Attachments = &v
}

func (o SkuDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ShippingCategory != nil {
		toSerialize["shipping_category"] = o.ShippingCategory
	}
	if o.Prices != nil {
		toSerialize["prices"] = o.Prices
	}
	if o.StockItems != nil {
		toSerialize["stock_items"] = o.StockItems
	}
	if o.DeliveryLeadTimes != nil {
		toSerialize["delivery_lead_times"] = o.DeliveryLeadTimes
	}
	if o.SkuOptions != nil {
		toSerialize["sku_options"] = o.SkuOptions
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	return json.Marshal(toSerialize)
}

type NullableSkuDataRelationships struct {
	value *SkuDataRelationships
	isSet bool
}

func (v NullableSkuDataRelationships) Get() *SkuDataRelationships {
	return v.value
}

func (v *NullableSkuDataRelationships) Set(val *SkuDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuDataRelationships(val *SkuDataRelationships) *NullableSkuDataRelationships {
	return &NullableSkuDataRelationships{value: val, isSet: true}
}

func (v NullableSkuDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkuDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
