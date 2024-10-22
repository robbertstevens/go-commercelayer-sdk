/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.3.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the POSTBraintreeGateways201ResponseDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTBraintreeGateways201ResponseDataRelationships{}

// POSTBraintreeGateways201ResponseDataRelationships struct for POSTBraintreeGateways201ResponseDataRelationships
type POSTBraintreeGateways201ResponseDataRelationships struct {
	PaymentMethods    *POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods        `json:"payment_methods,omitempty"`
	Versions          *POSTAddresses201ResponseDataRelationshipsVersions                  `json:"versions,omitempty"`
	BraintreePayments *POSTBraintreeGateways201ResponseDataRelationshipsBraintreePayments `json:"braintree_payments,omitempty"`
}

// NewPOSTBraintreeGateways201ResponseDataRelationships instantiates a new POSTBraintreeGateways201ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTBraintreeGateways201ResponseDataRelationships() *POSTBraintreeGateways201ResponseDataRelationships {
	this := POSTBraintreeGateways201ResponseDataRelationships{}
	return &this
}

// NewPOSTBraintreeGateways201ResponseDataRelationshipsWithDefaults instantiates a new POSTBraintreeGateways201ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTBraintreeGateways201ResponseDataRelationshipsWithDefaults() *POSTBraintreeGateways201ResponseDataRelationships {
	this := POSTBraintreeGateways201ResponseDataRelationships{}
	return &this
}

// GetPaymentMethods returns the PaymentMethods field value if set, zero value otherwise.
func (o *POSTBraintreeGateways201ResponseDataRelationships) GetPaymentMethods() POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods {
	if o == nil || IsNil(o.PaymentMethods) {
		var ret POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods
		return ret
	}
	return *o.PaymentMethods
}

// GetPaymentMethodsOk returns a tuple with the PaymentMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTBraintreeGateways201ResponseDataRelationships) GetPaymentMethodsOk() (*POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods, bool) {
	if o == nil || IsNil(o.PaymentMethods) {
		return nil, false
	}
	return o.PaymentMethods, true
}

// HasPaymentMethods returns a boolean if a field has been set.
func (o *POSTBraintreeGateways201ResponseDataRelationships) HasPaymentMethods() bool {
	if o != nil && !IsNil(o.PaymentMethods) {
		return true
	}

	return false
}

// SetPaymentMethods gets a reference to the given POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods and assigns it to the PaymentMethods field.
func (o *POSTBraintreeGateways201ResponseDataRelationships) SetPaymentMethods(v POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods) {
	o.PaymentMethods = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *POSTBraintreeGateways201ResponseDataRelationships) GetVersions() POSTAddresses201ResponseDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret POSTAddresses201ResponseDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTBraintreeGateways201ResponseDataRelationships) GetVersionsOk() (*POSTAddresses201ResponseDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *POSTBraintreeGateways201ResponseDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given POSTAddresses201ResponseDataRelationshipsVersions and assigns it to the Versions field.
func (o *POSTBraintreeGateways201ResponseDataRelationships) SetVersions(v POSTAddresses201ResponseDataRelationshipsVersions) {
	o.Versions = &v
}

// GetBraintreePayments returns the BraintreePayments field value if set, zero value otherwise.
func (o *POSTBraintreeGateways201ResponseDataRelationships) GetBraintreePayments() POSTBraintreeGateways201ResponseDataRelationshipsBraintreePayments {
	if o == nil || IsNil(o.BraintreePayments) {
		var ret POSTBraintreeGateways201ResponseDataRelationshipsBraintreePayments
		return ret
	}
	return *o.BraintreePayments
}

// GetBraintreePaymentsOk returns a tuple with the BraintreePayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTBraintreeGateways201ResponseDataRelationships) GetBraintreePaymentsOk() (*POSTBraintreeGateways201ResponseDataRelationshipsBraintreePayments, bool) {
	if o == nil || IsNil(o.BraintreePayments) {
		return nil, false
	}
	return o.BraintreePayments, true
}

// HasBraintreePayments returns a boolean if a field has been set.
func (o *POSTBraintreeGateways201ResponseDataRelationships) HasBraintreePayments() bool {
	if o != nil && !IsNil(o.BraintreePayments) {
		return true
	}

	return false
}

// SetBraintreePayments gets a reference to the given POSTBraintreeGateways201ResponseDataRelationshipsBraintreePayments and assigns it to the BraintreePayments field.
func (o *POSTBraintreeGateways201ResponseDataRelationships) SetBraintreePayments(v POSTBraintreeGateways201ResponseDataRelationshipsBraintreePayments) {
	o.BraintreePayments = &v
}

func (o POSTBraintreeGateways201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTBraintreeGateways201ResponseDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PaymentMethods) {
		toSerialize["payment_methods"] = o.PaymentMethods
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	if !IsNil(o.BraintreePayments) {
		toSerialize["braintree_payments"] = o.BraintreePayments
	}
	return toSerialize, nil
}

type NullablePOSTBraintreeGateways201ResponseDataRelationships struct {
	value *POSTBraintreeGateways201ResponseDataRelationships
	isSet bool
}

func (v NullablePOSTBraintreeGateways201ResponseDataRelationships) Get() *POSTBraintreeGateways201ResponseDataRelationships {
	return v.value
}

func (v *NullablePOSTBraintreeGateways201ResponseDataRelationships) Set(val *POSTBraintreeGateways201ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTBraintreeGateways201ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTBraintreeGateways201ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTBraintreeGateways201ResponseDataRelationships(val *POSTBraintreeGateways201ResponseDataRelationships) *NullablePOSTBraintreeGateways201ResponseDataRelationships {
	return &NullablePOSTBraintreeGateways201ResponseDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTBraintreeGateways201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTBraintreeGateways201ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
