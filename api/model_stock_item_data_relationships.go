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

// StockItemDataRelationships struct for StockItemDataRelationships
type StockItemDataRelationships struct {
	StockLocation *DeliveryLeadTimeDataRelationshipsStockLocation `json:"stock_location,omitempty"`
	Sku           *BundleDataRelationshipsSkus                    `json:"sku,omitempty"`
	Attachments   *AvalaraAccountDataRelationshipsAttachments     `json:"attachments,omitempty"`
}

// NewStockItemDataRelationships instantiates a new StockItemDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStockItemDataRelationships() *StockItemDataRelationships {
	this := StockItemDataRelationships{}
	return &this
}

// NewStockItemDataRelationshipsWithDefaults instantiates a new StockItemDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStockItemDataRelationshipsWithDefaults() *StockItemDataRelationships {
	this := StockItemDataRelationships{}
	return &this
}

// GetStockLocation returns the StockLocation field value if set, zero value otherwise.
func (o *StockItemDataRelationships) GetStockLocation() DeliveryLeadTimeDataRelationshipsStockLocation {
	if o == nil || o.StockLocation == nil {
		var ret DeliveryLeadTimeDataRelationshipsStockLocation
		return ret
	}
	return *o.StockLocation
}

// GetStockLocationOk returns a tuple with the StockLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockItemDataRelationships) GetStockLocationOk() (*DeliveryLeadTimeDataRelationshipsStockLocation, bool) {
	if o == nil || o.StockLocation == nil {
		return nil, false
	}
	return o.StockLocation, true
}

// HasStockLocation returns a boolean if a field has been set.
func (o *StockItemDataRelationships) HasStockLocation() bool {
	if o != nil && o.StockLocation != nil {
		return true
	}

	return false
}

// SetStockLocation gets a reference to the given DeliveryLeadTimeDataRelationshipsStockLocation and assigns it to the StockLocation field.
func (o *StockItemDataRelationships) SetStockLocation(v DeliveryLeadTimeDataRelationshipsStockLocation) {
	o.StockLocation = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *StockItemDataRelationships) GetSku() BundleDataRelationshipsSkus {
	if o == nil || o.Sku == nil {
		var ret BundleDataRelationshipsSkus
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockItemDataRelationships) GetSkuOk() (*BundleDataRelationshipsSkus, bool) {
	if o == nil || o.Sku == nil {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *StockItemDataRelationships) HasSku() bool {
	if o != nil && o.Sku != nil {
		return true
	}

	return false
}

// SetSku gets a reference to the given BundleDataRelationshipsSkus and assigns it to the Sku field.
func (o *StockItemDataRelationships) SetSku(v BundleDataRelationshipsSkus) {
	o.Sku = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *StockItemDataRelationships) GetAttachments() AvalaraAccountDataRelationshipsAttachments {
	if o == nil || o.Attachments == nil {
		var ret AvalaraAccountDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockItemDataRelationships) GetAttachmentsOk() (*AvalaraAccountDataRelationshipsAttachments, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *StockItemDataRelationships) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given AvalaraAccountDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *StockItemDataRelationships) SetAttachments(v AvalaraAccountDataRelationshipsAttachments) {
	o.Attachments = &v
}

func (o StockItemDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StockLocation != nil {
		toSerialize["stock_location"] = o.StockLocation
	}
	if o.Sku != nil {
		toSerialize["sku"] = o.Sku
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	return json.Marshal(toSerialize)
}

type NullableStockItemDataRelationships struct {
	value *StockItemDataRelationships
	isSet bool
}

func (v NullableStockItemDataRelationships) Get() *StockItemDataRelationships {
	return v.value
}

func (v *NullableStockItemDataRelationships) Set(val *StockItemDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableStockItemDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableStockItemDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStockItemDataRelationships(val *StockItemDataRelationships) *NullableStockItemDataRelationships {
	return &NullableStockItemDataRelationships{value: val, isSet: true}
}

func (v NullableStockItemDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStockItemDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
