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

// BraintreeGatewayCreateDataRelationships struct for BraintreeGatewayCreateDataRelationships
type BraintreeGatewayCreateDataRelationships struct {
	BraintreePayments *BraintreeGatewayDataRelationshipsBraintreePayments `json:"braintree_payments,omitempty"`
}

// NewBraintreeGatewayCreateDataRelationships instantiates a new BraintreeGatewayCreateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBraintreeGatewayCreateDataRelationships() *BraintreeGatewayCreateDataRelationships {
	this := BraintreeGatewayCreateDataRelationships{}
	return &this
}

// NewBraintreeGatewayCreateDataRelationshipsWithDefaults instantiates a new BraintreeGatewayCreateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBraintreeGatewayCreateDataRelationshipsWithDefaults() *BraintreeGatewayCreateDataRelationships {
	this := BraintreeGatewayCreateDataRelationships{}
	return &this
}

// GetBraintreePayments returns the BraintreePayments field value if set, zero value otherwise.
func (o *BraintreeGatewayCreateDataRelationships) GetBraintreePayments() BraintreeGatewayDataRelationshipsBraintreePayments {
	if o == nil || o.BraintreePayments == nil {
		var ret BraintreeGatewayDataRelationshipsBraintreePayments
		return ret
	}
	return *o.BraintreePayments
}

// GetBraintreePaymentsOk returns a tuple with the BraintreePayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BraintreeGatewayCreateDataRelationships) GetBraintreePaymentsOk() (*BraintreeGatewayDataRelationshipsBraintreePayments, bool) {
	if o == nil || o.BraintreePayments == nil {
		return nil, false
	}
	return o.BraintreePayments, true
}

// HasBraintreePayments returns a boolean if a field has been set.
func (o *BraintreeGatewayCreateDataRelationships) HasBraintreePayments() bool {
	if o != nil && o.BraintreePayments != nil {
		return true
	}

	return false
}

// SetBraintreePayments gets a reference to the given BraintreeGatewayDataRelationshipsBraintreePayments and assigns it to the BraintreePayments field.
func (o *BraintreeGatewayCreateDataRelationships) SetBraintreePayments(v BraintreeGatewayDataRelationshipsBraintreePayments) {
	o.BraintreePayments = &v
}

func (o BraintreeGatewayCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BraintreePayments != nil {
		toSerialize["braintree_payments"] = o.BraintreePayments
	}
	return json.Marshal(toSerialize)
}

type NullableBraintreeGatewayCreateDataRelationships struct {
	value *BraintreeGatewayCreateDataRelationships
	isSet bool
}

func (v NullableBraintreeGatewayCreateDataRelationships) Get() *BraintreeGatewayCreateDataRelationships {
	return v.value
}

func (v *NullableBraintreeGatewayCreateDataRelationships) Set(val *BraintreeGatewayCreateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableBraintreeGatewayCreateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableBraintreeGatewayCreateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBraintreeGatewayCreateDataRelationships(val *BraintreeGatewayCreateDataRelationships) *NullableBraintreeGatewayCreateDataRelationships {
	return &NullableBraintreeGatewayCreateDataRelationships{value: val, isSet: true}
}

func (v NullableBraintreeGatewayCreateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBraintreeGatewayCreateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
