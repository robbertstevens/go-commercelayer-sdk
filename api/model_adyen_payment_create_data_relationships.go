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

// AdyenPaymentCreateDataRelationships struct for AdyenPaymentCreateDataRelationships
type AdyenPaymentCreateDataRelationships struct {
	Order AdyenPaymentCreateDataRelationshipsOrder `json:"order"`
}

// NewAdyenPaymentCreateDataRelationships instantiates a new AdyenPaymentCreateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdyenPaymentCreateDataRelationships(order AdyenPaymentCreateDataRelationshipsOrder) *AdyenPaymentCreateDataRelationships {
	this := AdyenPaymentCreateDataRelationships{}
	this.Order = order
	return &this
}

// NewAdyenPaymentCreateDataRelationshipsWithDefaults instantiates a new AdyenPaymentCreateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdyenPaymentCreateDataRelationshipsWithDefaults() *AdyenPaymentCreateDataRelationships {
	this := AdyenPaymentCreateDataRelationships{}
	return &this
}

// GetOrder returns the Order field value
func (o *AdyenPaymentCreateDataRelationships) GetOrder() AdyenPaymentCreateDataRelationshipsOrder {
	if o == nil {
		var ret AdyenPaymentCreateDataRelationshipsOrder
		return ret
	}

	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *AdyenPaymentCreateDataRelationships) GetOrderOk() (*AdyenPaymentCreateDataRelationshipsOrder, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value
func (o *AdyenPaymentCreateDataRelationships) SetOrder(v AdyenPaymentCreateDataRelationshipsOrder) {
	o.Order = v
}

func (o AdyenPaymentCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["order"] = o.Order
	}
	return json.Marshal(toSerialize)
}

type NullableAdyenPaymentCreateDataRelationships struct {
	value *AdyenPaymentCreateDataRelationships
	isSet bool
}

func (v NullableAdyenPaymentCreateDataRelationships) Get() *AdyenPaymentCreateDataRelationships {
	return v.value
}

func (v *NullableAdyenPaymentCreateDataRelationships) Set(val *AdyenPaymentCreateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAdyenPaymentCreateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAdyenPaymentCreateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdyenPaymentCreateDataRelationships(val *AdyenPaymentCreateDataRelationships) *NullableAdyenPaymentCreateDataRelationships {
	return &NullableAdyenPaymentCreateDataRelationships{value: val, isSet: true}
}

func (v NullableAdyenPaymentCreateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdyenPaymentCreateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
