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

// checks if the POSTStockItems201ResponseDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTStockItems201ResponseDataRelationships{}

// POSTStockItems201ResponseDataRelationships struct for POSTStockItems201ResponseDataRelationships
type POSTStockItems201ResponseDataRelationships struct {
	StockLocation *POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation `json:"stock_location,omitempty"`
	Sku           *POSTInStockSubscriptions201ResponseDataRelationshipsSku        `json:"sku,omitempty"`
	Attachments   *POSTAvalaraAccounts201ResponseDataRelationshipsAttachments     `json:"attachments,omitempty"`
}

// NewPOSTStockItems201ResponseDataRelationships instantiates a new POSTStockItems201ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTStockItems201ResponseDataRelationships() *POSTStockItems201ResponseDataRelationships {
	this := POSTStockItems201ResponseDataRelationships{}
	return &this
}

// NewPOSTStockItems201ResponseDataRelationshipsWithDefaults instantiates a new POSTStockItems201ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTStockItems201ResponseDataRelationshipsWithDefaults() *POSTStockItems201ResponseDataRelationships {
	this := POSTStockItems201ResponseDataRelationships{}
	return &this
}

// GetStockLocation returns the StockLocation field value if set, zero value otherwise.
func (o *POSTStockItems201ResponseDataRelationships) GetStockLocation() POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation {
	if o == nil || IsNil(o.StockLocation) {
		var ret POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation
		return ret
	}
	return *o.StockLocation
}

// GetStockLocationOk returns a tuple with the StockLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTStockItems201ResponseDataRelationships) GetStockLocationOk() (*POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation, bool) {
	if o == nil || IsNil(o.StockLocation) {
		return nil, false
	}
	return o.StockLocation, true
}

// HasStockLocation returns a boolean if a field has been set.
func (o *POSTStockItems201ResponseDataRelationships) HasStockLocation() bool {
	if o != nil && !IsNil(o.StockLocation) {
		return true
	}

	return false
}

// SetStockLocation gets a reference to the given POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation and assigns it to the StockLocation field.
func (o *POSTStockItems201ResponseDataRelationships) SetStockLocation(v POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation) {
	o.StockLocation = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *POSTStockItems201ResponseDataRelationships) GetSku() POSTInStockSubscriptions201ResponseDataRelationshipsSku {
	if o == nil || IsNil(o.Sku) {
		var ret POSTInStockSubscriptions201ResponseDataRelationshipsSku
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTStockItems201ResponseDataRelationships) GetSkuOk() (*POSTInStockSubscriptions201ResponseDataRelationshipsSku, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *POSTStockItems201ResponseDataRelationships) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given POSTInStockSubscriptions201ResponseDataRelationshipsSku and assigns it to the Sku field.
func (o *POSTStockItems201ResponseDataRelationships) SetSku(v POSTInStockSubscriptions201ResponseDataRelationshipsSku) {
	o.Sku = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *POSTStockItems201ResponseDataRelationships) GetAttachments() POSTAvalaraAccounts201ResponseDataRelationshipsAttachments {
	if o == nil || IsNil(o.Attachments) {
		var ret POSTAvalaraAccounts201ResponseDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTStockItems201ResponseDataRelationships) GetAttachmentsOk() (*POSTAvalaraAccounts201ResponseDataRelationshipsAttachments, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *POSTStockItems201ResponseDataRelationships) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given POSTAvalaraAccounts201ResponseDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *POSTStockItems201ResponseDataRelationships) SetAttachments(v POSTAvalaraAccounts201ResponseDataRelationshipsAttachments) {
	o.Attachments = &v
}

func (o POSTStockItems201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTStockItems201ResponseDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StockLocation) {
		toSerialize["stock_location"] = o.StockLocation
	}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	return toSerialize, nil
}

type NullablePOSTStockItems201ResponseDataRelationships struct {
	value *POSTStockItems201ResponseDataRelationships
	isSet bool
}

func (v NullablePOSTStockItems201ResponseDataRelationships) Get() *POSTStockItems201ResponseDataRelationships {
	return v.value
}

func (v *NullablePOSTStockItems201ResponseDataRelationships) Set(val *POSTStockItems201ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTStockItems201ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTStockItems201ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTStockItems201ResponseDataRelationships(val *POSTStockItems201ResponseDataRelationships) *NullablePOSTStockItems201ResponseDataRelationships {
	return &NullablePOSTStockItems201ResponseDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTStockItems201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTStockItems201ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
