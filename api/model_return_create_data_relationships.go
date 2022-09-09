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

// ReturnCreateDataRelationships struct for ReturnCreateDataRelationships
type ReturnCreateDataRelationships struct {
	Order AdyenPaymentDataRelationshipsOrder `json:"order"`
	StockLocation *DeliveryLeadTimeDataRelationshipsStockLocation `json:"stock_location,omitempty"`
}

// NewReturnCreateDataRelationships instantiates a new ReturnCreateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReturnCreateDataRelationships(order AdyenPaymentDataRelationshipsOrder) *ReturnCreateDataRelationships {
	this := ReturnCreateDataRelationships{}
	this.Order = order
	return &this
}

// NewReturnCreateDataRelationshipsWithDefaults instantiates a new ReturnCreateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReturnCreateDataRelationshipsWithDefaults() *ReturnCreateDataRelationships {
	this := ReturnCreateDataRelationships{}
	return &this
}

// GetOrder returns the Order field value
func (o *ReturnCreateDataRelationships) GetOrder() AdyenPaymentDataRelationshipsOrder {
	if o == nil {
		var ret AdyenPaymentDataRelationshipsOrder
		return ret
	}

	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *ReturnCreateDataRelationships) GetOrderOk() (*AdyenPaymentDataRelationshipsOrder, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value
func (o *ReturnCreateDataRelationships) SetOrder(v AdyenPaymentDataRelationshipsOrder) {
	o.Order = v
}

// GetStockLocation returns the StockLocation field value if set, zero value otherwise.
func (o *ReturnCreateDataRelationships) GetStockLocation() DeliveryLeadTimeDataRelationshipsStockLocation {
	if o == nil || o.StockLocation == nil {
		var ret DeliveryLeadTimeDataRelationshipsStockLocation
		return ret
	}
	return *o.StockLocation
}

// GetStockLocationOk returns a tuple with the StockLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnCreateDataRelationships) GetStockLocationOk() (*DeliveryLeadTimeDataRelationshipsStockLocation, bool) {
	if o == nil || o.StockLocation == nil {
		return nil, false
	}
	return o.StockLocation, true
}

// HasStockLocation returns a boolean if a field has been set.
func (o *ReturnCreateDataRelationships) HasStockLocation() bool {
	if o != nil && o.StockLocation != nil {
		return true
	}

	return false
}

// SetStockLocation gets a reference to the given DeliveryLeadTimeDataRelationshipsStockLocation and assigns it to the StockLocation field.
func (o *ReturnCreateDataRelationships) SetStockLocation(v DeliveryLeadTimeDataRelationshipsStockLocation) {
	o.StockLocation = &v
}

func (o ReturnCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["order"] = o.Order
	}
	if o.StockLocation != nil {
		toSerialize["stock_location"] = o.StockLocation
	}
	return json.Marshal(toSerialize)
}

type NullableReturnCreateDataRelationships struct {
	value *ReturnCreateDataRelationships
	isSet bool
}

func (v NullableReturnCreateDataRelationships) Get() *ReturnCreateDataRelationships {
	return v.value
}

func (v *NullableReturnCreateDataRelationships) Set(val *ReturnCreateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnCreateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnCreateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnCreateDataRelationships(val *ReturnCreateDataRelationships) *NullableReturnCreateDataRelationships {
	return &NullableReturnCreateDataRelationships{value: val, isSet: true}
}

func (v NullableReturnCreateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnCreateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


