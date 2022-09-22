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

// GETSkus200ResponseDataInnerRelationships struct for GETSkus200ResponseDataInnerRelationships
type GETSkus200ResponseDataInnerRelationships struct {
	ShippingCategory  *GETShipments200ResponseDataInnerRelationshipsShippingCategory  `json:"shipping_category,omitempty"`
	Prices            *GETPriceLists200ResponseDataInnerRelationshipsPrices           `json:"prices,omitempty"`
	StockItems        *GETSkus200ResponseDataInnerRelationshipsStockItems             `json:"stock_items,omitempty"`
	DeliveryLeadTimes *GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimes      `json:"delivery_lead_times,omitempty"`
	SkuOptions        *GETSkus200ResponseDataInnerRelationshipsSkuOptions             `json:"sku_options,omitempty"`
	Attachments       *GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments `json:"attachments,omitempty"`
}

// NewGETSkus200ResponseDataInnerRelationships instantiates a new GETSkus200ResponseDataInnerRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETSkus200ResponseDataInnerRelationships() *GETSkus200ResponseDataInnerRelationships {
	this := GETSkus200ResponseDataInnerRelationships{}
	return &this
}

// NewGETSkus200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETSkus200ResponseDataInnerRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETSkus200ResponseDataInnerRelationshipsWithDefaults() *GETSkus200ResponseDataInnerRelationships {
	this := GETSkus200ResponseDataInnerRelationships{}
	return &this
}

// GetShippingCategory returns the ShippingCategory field value if set, zero value otherwise.
func (o *GETSkus200ResponseDataInnerRelationships) GetShippingCategory() GETShipments200ResponseDataInnerRelationshipsShippingCategory {
	if o == nil || o.ShippingCategory == nil {
		var ret GETShipments200ResponseDataInnerRelationshipsShippingCategory
		return ret
	}
	return *o.ShippingCategory
}

// GetShippingCategoryOk returns a tuple with the ShippingCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkus200ResponseDataInnerRelationships) GetShippingCategoryOk() (*GETShipments200ResponseDataInnerRelationshipsShippingCategory, bool) {
	if o == nil || o.ShippingCategory == nil {
		return nil, false
	}
	return o.ShippingCategory, true
}

// HasShippingCategory returns a boolean if a field has been set.
func (o *GETSkus200ResponseDataInnerRelationships) HasShippingCategory() bool {
	if o != nil && o.ShippingCategory != nil {
		return true
	}

	return false
}

// SetShippingCategory gets a reference to the given GETShipments200ResponseDataInnerRelationshipsShippingCategory and assigns it to the ShippingCategory field.
func (o *GETSkus200ResponseDataInnerRelationships) SetShippingCategory(v GETShipments200ResponseDataInnerRelationshipsShippingCategory) {
	o.ShippingCategory = &v
}

// GetPrices returns the Prices field value if set, zero value otherwise.
func (o *GETSkus200ResponseDataInnerRelationships) GetPrices() GETPriceLists200ResponseDataInnerRelationshipsPrices {
	if o == nil || o.Prices == nil {
		var ret GETPriceLists200ResponseDataInnerRelationshipsPrices
		return ret
	}
	return *o.Prices
}

// GetPricesOk returns a tuple with the Prices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkus200ResponseDataInnerRelationships) GetPricesOk() (*GETPriceLists200ResponseDataInnerRelationshipsPrices, bool) {
	if o == nil || o.Prices == nil {
		return nil, false
	}
	return o.Prices, true
}

// HasPrices returns a boolean if a field has been set.
func (o *GETSkus200ResponseDataInnerRelationships) HasPrices() bool {
	if o != nil && o.Prices != nil {
		return true
	}

	return false
}

// SetPrices gets a reference to the given GETPriceLists200ResponseDataInnerRelationshipsPrices and assigns it to the Prices field.
func (o *GETSkus200ResponseDataInnerRelationships) SetPrices(v GETPriceLists200ResponseDataInnerRelationshipsPrices) {
	o.Prices = &v
}

// GetStockItems returns the StockItems field value if set, zero value otherwise.
func (o *GETSkus200ResponseDataInnerRelationships) GetStockItems() GETSkus200ResponseDataInnerRelationshipsStockItems {
	if o == nil || o.StockItems == nil {
		var ret GETSkus200ResponseDataInnerRelationshipsStockItems
		return ret
	}
	return *o.StockItems
}

// GetStockItemsOk returns a tuple with the StockItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkus200ResponseDataInnerRelationships) GetStockItemsOk() (*GETSkus200ResponseDataInnerRelationshipsStockItems, bool) {
	if o == nil || o.StockItems == nil {
		return nil, false
	}
	return o.StockItems, true
}

// HasStockItems returns a boolean if a field has been set.
func (o *GETSkus200ResponseDataInnerRelationships) HasStockItems() bool {
	if o != nil && o.StockItems != nil {
		return true
	}

	return false
}

// SetStockItems gets a reference to the given GETSkus200ResponseDataInnerRelationshipsStockItems and assigns it to the StockItems field.
func (o *GETSkus200ResponseDataInnerRelationships) SetStockItems(v GETSkus200ResponseDataInnerRelationshipsStockItems) {
	o.StockItems = &v
}

// GetDeliveryLeadTimes returns the DeliveryLeadTimes field value if set, zero value otherwise.
func (o *GETSkus200ResponseDataInnerRelationships) GetDeliveryLeadTimes() GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimes {
	if o == nil || o.DeliveryLeadTimes == nil {
		var ret GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimes
		return ret
	}
	return *o.DeliveryLeadTimes
}

// GetDeliveryLeadTimesOk returns a tuple with the DeliveryLeadTimes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkus200ResponseDataInnerRelationships) GetDeliveryLeadTimesOk() (*GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimes, bool) {
	if o == nil || o.DeliveryLeadTimes == nil {
		return nil, false
	}
	return o.DeliveryLeadTimes, true
}

// HasDeliveryLeadTimes returns a boolean if a field has been set.
func (o *GETSkus200ResponseDataInnerRelationships) HasDeliveryLeadTimes() bool {
	if o != nil && o.DeliveryLeadTimes != nil {
		return true
	}

	return false
}

// SetDeliveryLeadTimes gets a reference to the given GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimes and assigns it to the DeliveryLeadTimes field.
func (o *GETSkus200ResponseDataInnerRelationships) SetDeliveryLeadTimes(v GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimes) {
	o.DeliveryLeadTimes = &v
}

// GetSkuOptions returns the SkuOptions field value if set, zero value otherwise.
func (o *GETSkus200ResponseDataInnerRelationships) GetSkuOptions() GETSkus200ResponseDataInnerRelationshipsSkuOptions {
	if o == nil || o.SkuOptions == nil {
		var ret GETSkus200ResponseDataInnerRelationshipsSkuOptions
		return ret
	}
	return *o.SkuOptions
}

// GetSkuOptionsOk returns a tuple with the SkuOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkus200ResponseDataInnerRelationships) GetSkuOptionsOk() (*GETSkus200ResponseDataInnerRelationshipsSkuOptions, bool) {
	if o == nil || o.SkuOptions == nil {
		return nil, false
	}
	return o.SkuOptions, true
}

// HasSkuOptions returns a boolean if a field has been set.
func (o *GETSkus200ResponseDataInnerRelationships) HasSkuOptions() bool {
	if o != nil && o.SkuOptions != nil {
		return true
	}

	return false
}

// SetSkuOptions gets a reference to the given GETSkus200ResponseDataInnerRelationshipsSkuOptions and assigns it to the SkuOptions field.
func (o *GETSkus200ResponseDataInnerRelationships) SetSkuOptions(v GETSkus200ResponseDataInnerRelationshipsSkuOptions) {
	o.SkuOptions = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *GETSkus200ResponseDataInnerRelationships) GetAttachments() GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments {
	if o == nil || o.Attachments == nil {
		var ret GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkus200ResponseDataInnerRelationships) GetAttachmentsOk() (*GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *GETSkus200ResponseDataInnerRelationships) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments and assigns it to the Attachments field.
func (o *GETSkus200ResponseDataInnerRelationships) SetAttachments(v GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments) {
	o.Attachments = &v
}

func (o GETSkus200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
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

type NullableGETSkus200ResponseDataInnerRelationships struct {
	value *GETSkus200ResponseDataInnerRelationships
	isSet bool
}

func (v NullableGETSkus200ResponseDataInnerRelationships) Get() *GETSkus200ResponseDataInnerRelationships {
	return v.value
}

func (v *NullableGETSkus200ResponseDataInnerRelationships) Set(val *GETSkus200ResponseDataInnerRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETSkus200ResponseDataInnerRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETSkus200ResponseDataInnerRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETSkus200ResponseDataInnerRelationships(val *GETSkus200ResponseDataInnerRelationships) *NullableGETSkus200ResponseDataInnerRelationships {
	return &NullableGETSkus200ResponseDataInnerRelationships{value: val, isSet: true}
}

func (v NullableGETSkus200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETSkus200ResponseDataInnerRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
