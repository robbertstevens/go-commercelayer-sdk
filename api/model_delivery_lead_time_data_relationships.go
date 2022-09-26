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

// DeliveryLeadTimeDataRelationships struct for DeliveryLeadTimeDataRelationships
type DeliveryLeadTimeDataRelationships struct {
	StockLocation  *DeliveryLeadTimeDataRelationshipsStockLocation  `json:"stock_location,omitempty"`
	ShippingMethod *DeliveryLeadTimeDataRelationshipsShippingMethod `json:"shipping_method,omitempty"`
	Attachments    *AvalaraAccountDataRelationshipsAttachments      `json:"attachments,omitempty"`
}

// NewDeliveryLeadTimeDataRelationships instantiates a new DeliveryLeadTimeDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryLeadTimeDataRelationships() *DeliveryLeadTimeDataRelationships {
	this := DeliveryLeadTimeDataRelationships{}
	return &this
}

// NewDeliveryLeadTimeDataRelationshipsWithDefaults instantiates a new DeliveryLeadTimeDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryLeadTimeDataRelationshipsWithDefaults() *DeliveryLeadTimeDataRelationships {
	this := DeliveryLeadTimeDataRelationships{}
	return &this
}

// GetStockLocation returns the StockLocation field value if set, zero value otherwise.
func (o *DeliveryLeadTimeDataRelationships) GetStockLocation() DeliveryLeadTimeDataRelationshipsStockLocation {
	if o == nil || o.StockLocation == nil {
		var ret DeliveryLeadTimeDataRelationshipsStockLocation
		return ret
	}
	return *o.StockLocation
}

// GetStockLocationOk returns a tuple with the StockLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryLeadTimeDataRelationships) GetStockLocationOk() (*DeliveryLeadTimeDataRelationshipsStockLocation, bool) {
	if o == nil || o.StockLocation == nil {
		return nil, false
	}
	return o.StockLocation, true
}

// HasStockLocation returns a boolean if a field has been set.
func (o *DeliveryLeadTimeDataRelationships) HasStockLocation() bool {
	if o != nil && o.StockLocation != nil {
		return true
	}

	return false
}

// SetStockLocation gets a reference to the given DeliveryLeadTimeDataRelationshipsStockLocation and assigns it to the StockLocation field.
func (o *DeliveryLeadTimeDataRelationships) SetStockLocation(v DeliveryLeadTimeDataRelationshipsStockLocation) {
	o.StockLocation = &v
}

// GetShippingMethod returns the ShippingMethod field value if set, zero value otherwise.
func (o *DeliveryLeadTimeDataRelationships) GetShippingMethod() DeliveryLeadTimeDataRelationshipsShippingMethod {
	if o == nil || o.ShippingMethod == nil {
		var ret DeliveryLeadTimeDataRelationshipsShippingMethod
		return ret
	}
	return *o.ShippingMethod
}

// GetShippingMethodOk returns a tuple with the ShippingMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryLeadTimeDataRelationships) GetShippingMethodOk() (*DeliveryLeadTimeDataRelationshipsShippingMethod, bool) {
	if o == nil || o.ShippingMethod == nil {
		return nil, false
	}
	return o.ShippingMethod, true
}

// HasShippingMethod returns a boolean if a field has been set.
func (o *DeliveryLeadTimeDataRelationships) HasShippingMethod() bool {
	if o != nil && o.ShippingMethod != nil {
		return true
	}

	return false
}

// SetShippingMethod gets a reference to the given DeliveryLeadTimeDataRelationshipsShippingMethod and assigns it to the ShippingMethod field.
func (o *DeliveryLeadTimeDataRelationships) SetShippingMethod(v DeliveryLeadTimeDataRelationshipsShippingMethod) {
	o.ShippingMethod = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *DeliveryLeadTimeDataRelationships) GetAttachments() AvalaraAccountDataRelationshipsAttachments {
	if o == nil || o.Attachments == nil {
		var ret AvalaraAccountDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryLeadTimeDataRelationships) GetAttachmentsOk() (*AvalaraAccountDataRelationshipsAttachments, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *DeliveryLeadTimeDataRelationships) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given AvalaraAccountDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *DeliveryLeadTimeDataRelationships) SetAttachments(v AvalaraAccountDataRelationshipsAttachments) {
	o.Attachments = &v
}

func (o DeliveryLeadTimeDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StockLocation != nil {
		toSerialize["stock_location"] = o.StockLocation
	}
	if o.ShippingMethod != nil {
		toSerialize["shipping_method"] = o.ShippingMethod
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	return json.Marshal(toSerialize)
}

type NullableDeliveryLeadTimeDataRelationships struct {
	value *DeliveryLeadTimeDataRelationships
	isSet bool
}

func (v NullableDeliveryLeadTimeDataRelationships) Get() *DeliveryLeadTimeDataRelationships {
	return v.value
}

func (v *NullableDeliveryLeadTimeDataRelationships) Set(val *DeliveryLeadTimeDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryLeadTimeDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryLeadTimeDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryLeadTimeDataRelationships(val *DeliveryLeadTimeDataRelationships) *NullableDeliveryLeadTimeDataRelationships {
	return &NullableDeliveryLeadTimeDataRelationships{value: val, isSet: true}
}

func (v NullableDeliveryLeadTimeDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryLeadTimeDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
