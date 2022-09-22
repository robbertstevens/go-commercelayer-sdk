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

// KlarnaGatewayCreateDataRelationships struct for KlarnaGatewayCreateDataRelationships
type KlarnaGatewayCreateDataRelationships struct {
	KlarnaPayments *KlarnaGatewayDataRelationshipsKlarnaPayments `json:"klarna_payments,omitempty"`
}

// NewKlarnaGatewayCreateDataRelationships instantiates a new KlarnaGatewayCreateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKlarnaGatewayCreateDataRelationships() *KlarnaGatewayCreateDataRelationships {
	this := KlarnaGatewayCreateDataRelationships{}
	return &this
}

// NewKlarnaGatewayCreateDataRelationshipsWithDefaults instantiates a new KlarnaGatewayCreateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKlarnaGatewayCreateDataRelationshipsWithDefaults() *KlarnaGatewayCreateDataRelationships {
	this := KlarnaGatewayCreateDataRelationships{}
	return &this
}

// GetKlarnaPayments returns the KlarnaPayments field value if set, zero value otherwise.
func (o *KlarnaGatewayCreateDataRelationships) GetKlarnaPayments() KlarnaGatewayDataRelationshipsKlarnaPayments {
	if o == nil || o.KlarnaPayments == nil {
		var ret KlarnaGatewayDataRelationshipsKlarnaPayments
		return ret
	}
	return *o.KlarnaPayments
}

// GetKlarnaPaymentsOk returns a tuple with the KlarnaPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlarnaGatewayCreateDataRelationships) GetKlarnaPaymentsOk() (*KlarnaGatewayDataRelationshipsKlarnaPayments, bool) {
	if o == nil || o.KlarnaPayments == nil {
		return nil, false
	}
	return o.KlarnaPayments, true
}

// HasKlarnaPayments returns a boolean if a field has been set.
func (o *KlarnaGatewayCreateDataRelationships) HasKlarnaPayments() bool {
	if o != nil && o.KlarnaPayments != nil {
		return true
	}

	return false
}

// SetKlarnaPayments gets a reference to the given KlarnaGatewayDataRelationshipsKlarnaPayments and assigns it to the KlarnaPayments field.
func (o *KlarnaGatewayCreateDataRelationships) SetKlarnaPayments(v KlarnaGatewayDataRelationshipsKlarnaPayments) {
	o.KlarnaPayments = &v
}

func (o KlarnaGatewayCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.KlarnaPayments != nil {
		toSerialize["klarna_payments"] = o.KlarnaPayments
	}
	return json.Marshal(toSerialize)
}

type NullableKlarnaGatewayCreateDataRelationships struct {
	value *KlarnaGatewayCreateDataRelationships
	isSet bool
}

func (v NullableKlarnaGatewayCreateDataRelationships) Get() *KlarnaGatewayCreateDataRelationships {
	return v.value
}

func (v *NullableKlarnaGatewayCreateDataRelationships) Set(val *KlarnaGatewayCreateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableKlarnaGatewayCreateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableKlarnaGatewayCreateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKlarnaGatewayCreateDataRelationships(val *KlarnaGatewayCreateDataRelationships) *NullableKlarnaGatewayCreateDataRelationships {
	return &NullableKlarnaGatewayCreateDataRelationships{value: val, isSet: true}
}

func (v NullableKlarnaGatewayCreateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKlarnaGatewayCreateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
