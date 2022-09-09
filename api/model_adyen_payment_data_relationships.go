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

// AdyenPaymentDataRelationships struct for AdyenPaymentDataRelationships
type AdyenPaymentDataRelationships struct {
	Order *AdyenPaymentDataRelationshipsOrder `json:"order,omitempty"`
	PaymentGateway *AdyenPaymentDataRelationshipsPaymentGateway `json:"payment_gateway,omitempty"`
}

// NewAdyenPaymentDataRelationships instantiates a new AdyenPaymentDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdyenPaymentDataRelationships() *AdyenPaymentDataRelationships {
	this := AdyenPaymentDataRelationships{}
	return &this
}

// NewAdyenPaymentDataRelationshipsWithDefaults instantiates a new AdyenPaymentDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdyenPaymentDataRelationshipsWithDefaults() *AdyenPaymentDataRelationships {
	this := AdyenPaymentDataRelationships{}
	return &this
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *AdyenPaymentDataRelationships) GetOrder() AdyenPaymentDataRelationshipsOrder {
	if o == nil || o.Order == nil {
		var ret AdyenPaymentDataRelationshipsOrder
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdyenPaymentDataRelationships) GetOrderOk() (*AdyenPaymentDataRelationshipsOrder, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *AdyenPaymentDataRelationships) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given AdyenPaymentDataRelationshipsOrder and assigns it to the Order field.
func (o *AdyenPaymentDataRelationships) SetOrder(v AdyenPaymentDataRelationshipsOrder) {
	o.Order = &v
}

// GetPaymentGateway returns the PaymentGateway field value if set, zero value otherwise.
func (o *AdyenPaymentDataRelationships) GetPaymentGateway() AdyenPaymentDataRelationshipsPaymentGateway {
	if o == nil || o.PaymentGateway == nil {
		var ret AdyenPaymentDataRelationshipsPaymentGateway
		return ret
	}
	return *o.PaymentGateway
}

// GetPaymentGatewayOk returns a tuple with the PaymentGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdyenPaymentDataRelationships) GetPaymentGatewayOk() (*AdyenPaymentDataRelationshipsPaymentGateway, bool) {
	if o == nil || o.PaymentGateway == nil {
		return nil, false
	}
	return o.PaymentGateway, true
}

// HasPaymentGateway returns a boolean if a field has been set.
func (o *AdyenPaymentDataRelationships) HasPaymentGateway() bool {
	if o != nil && o.PaymentGateway != nil {
		return true
	}

	return false
}

// SetPaymentGateway gets a reference to the given AdyenPaymentDataRelationshipsPaymentGateway and assigns it to the PaymentGateway field.
func (o *AdyenPaymentDataRelationships) SetPaymentGateway(v AdyenPaymentDataRelationshipsPaymentGateway) {
	o.PaymentGateway = &v
}

func (o AdyenPaymentDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Order != nil {
		toSerialize["order"] = o.Order
	}
	if o.PaymentGateway != nil {
		toSerialize["payment_gateway"] = o.PaymentGateway
	}
	return json.Marshal(toSerialize)
}

type NullableAdyenPaymentDataRelationships struct {
	value *AdyenPaymentDataRelationships
	isSet bool
}

func (v NullableAdyenPaymentDataRelationships) Get() *AdyenPaymentDataRelationships {
	return v.value
}

func (v *NullableAdyenPaymentDataRelationships) Set(val *AdyenPaymentDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAdyenPaymentDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAdyenPaymentDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdyenPaymentDataRelationships(val *AdyenPaymentDataRelationships) *NullableAdyenPaymentDataRelationships {
	return &NullableAdyenPaymentDataRelationships{value: val, isSet: true}
}

func (v NullableAdyenPaymentDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdyenPaymentDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


