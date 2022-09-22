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

// DeliveryLeadTimeCreateDataRelationships struct for DeliveryLeadTimeCreateDataRelationships
type DeliveryLeadTimeCreateDataRelationships struct {
	StockLocation  DeliveryLeadTimeDataRelationshipsStockLocation  `json:"stock_location"`
	ShippingMethod DeliveryLeadTimeDataRelationshipsShippingMethod `json:"shipping_method"`
}

// NewDeliveryLeadTimeCreateDataRelationships instantiates a new DeliveryLeadTimeCreateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryLeadTimeCreateDataRelationships(stockLocation DeliveryLeadTimeDataRelationshipsStockLocation, shippingMethod DeliveryLeadTimeDataRelationshipsShippingMethod) *DeliveryLeadTimeCreateDataRelationships {
	this := DeliveryLeadTimeCreateDataRelationships{}
	this.StockLocation = stockLocation
	this.ShippingMethod = shippingMethod
	return &this
}

// NewDeliveryLeadTimeCreateDataRelationshipsWithDefaults instantiates a new DeliveryLeadTimeCreateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryLeadTimeCreateDataRelationshipsWithDefaults() *DeliveryLeadTimeCreateDataRelationships {
	this := DeliveryLeadTimeCreateDataRelationships{}
	return &this
}

// GetStockLocation returns the StockLocation field value
func (o *DeliveryLeadTimeCreateDataRelationships) GetStockLocation() DeliveryLeadTimeDataRelationshipsStockLocation {
	if o == nil {
		var ret DeliveryLeadTimeDataRelationshipsStockLocation
		return ret
	}

	return o.StockLocation
}

// GetStockLocationOk returns a tuple with the StockLocation field value
// and a boolean to check if the value has been set.
func (o *DeliveryLeadTimeCreateDataRelationships) GetStockLocationOk() (*DeliveryLeadTimeDataRelationshipsStockLocation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StockLocation, true
}

// SetStockLocation sets field value
func (o *DeliveryLeadTimeCreateDataRelationships) SetStockLocation(v DeliveryLeadTimeDataRelationshipsStockLocation) {
	o.StockLocation = v
}

// GetShippingMethod returns the ShippingMethod field value
func (o *DeliveryLeadTimeCreateDataRelationships) GetShippingMethod() DeliveryLeadTimeDataRelationshipsShippingMethod {
	if o == nil {
		var ret DeliveryLeadTimeDataRelationshipsShippingMethod
		return ret
	}

	return o.ShippingMethod
}

// GetShippingMethodOk returns a tuple with the ShippingMethod field value
// and a boolean to check if the value has been set.
func (o *DeliveryLeadTimeCreateDataRelationships) GetShippingMethodOk() (*DeliveryLeadTimeDataRelationshipsShippingMethod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShippingMethod, true
}

// SetShippingMethod sets field value
func (o *DeliveryLeadTimeCreateDataRelationships) SetShippingMethod(v DeliveryLeadTimeDataRelationshipsShippingMethod) {
	o.ShippingMethod = v
}

func (o DeliveryLeadTimeCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["stock_location"] = o.StockLocation
	}
	if true {
		toSerialize["shipping_method"] = o.ShippingMethod
	}
	return json.Marshal(toSerialize)
}

type NullableDeliveryLeadTimeCreateDataRelationships struct {
	value *DeliveryLeadTimeCreateDataRelationships
	isSet bool
}

func (v NullableDeliveryLeadTimeCreateDataRelationships) Get() *DeliveryLeadTimeCreateDataRelationships {
	return v.value
}

func (v *NullableDeliveryLeadTimeCreateDataRelationships) Set(val *DeliveryLeadTimeCreateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryLeadTimeCreateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryLeadTimeCreateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryLeadTimeCreateDataRelationships(val *DeliveryLeadTimeCreateDataRelationships) *NullableDeliveryLeadTimeCreateDataRelationships {
	return &NullableDeliveryLeadTimeCreateDataRelationships{value: val, isSet: true}
}

func (v NullableDeliveryLeadTimeCreateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryLeadTimeCreateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
