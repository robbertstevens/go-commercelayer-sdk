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

// checks if the POSTAdyenPayments201ResponseDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTAdyenPayments201ResponseDataRelationships{}

// POSTAdyenPayments201ResponseDataRelationships struct for POSTAdyenPayments201ResponseDataRelationships
type POSTAdyenPayments201ResponseDataRelationships struct {
	Order          *POSTAdyenPayments201ResponseDataRelationshipsOrder          `json:"order,omitempty"`
	PaymentGateway *POSTAdyenPayments201ResponseDataRelationshipsPaymentGateway `json:"payment_gateway,omitempty"`
	Versions       *POSTAddresses201ResponseDataRelationshipsVersions           `json:"versions,omitempty"`
}

// NewPOSTAdyenPayments201ResponseDataRelationships instantiates a new POSTAdyenPayments201ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTAdyenPayments201ResponseDataRelationships() *POSTAdyenPayments201ResponseDataRelationships {
	this := POSTAdyenPayments201ResponseDataRelationships{}
	return &this
}

// NewPOSTAdyenPayments201ResponseDataRelationshipsWithDefaults instantiates a new POSTAdyenPayments201ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTAdyenPayments201ResponseDataRelationshipsWithDefaults() *POSTAdyenPayments201ResponseDataRelationships {
	this := POSTAdyenPayments201ResponseDataRelationships{}
	return &this
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *POSTAdyenPayments201ResponseDataRelationships) GetOrder() POSTAdyenPayments201ResponseDataRelationshipsOrder {
	if o == nil || IsNil(o.Order) {
		var ret POSTAdyenPayments201ResponseDataRelationshipsOrder
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTAdyenPayments201ResponseDataRelationships) GetOrderOk() (*POSTAdyenPayments201ResponseDataRelationshipsOrder, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *POSTAdyenPayments201ResponseDataRelationships) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given POSTAdyenPayments201ResponseDataRelationshipsOrder and assigns it to the Order field.
func (o *POSTAdyenPayments201ResponseDataRelationships) SetOrder(v POSTAdyenPayments201ResponseDataRelationshipsOrder) {
	o.Order = &v
}

// GetPaymentGateway returns the PaymentGateway field value if set, zero value otherwise.
func (o *POSTAdyenPayments201ResponseDataRelationships) GetPaymentGateway() POSTAdyenPayments201ResponseDataRelationshipsPaymentGateway {
	if o == nil || IsNil(o.PaymentGateway) {
		var ret POSTAdyenPayments201ResponseDataRelationshipsPaymentGateway
		return ret
	}
	return *o.PaymentGateway
}

// GetPaymentGatewayOk returns a tuple with the PaymentGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTAdyenPayments201ResponseDataRelationships) GetPaymentGatewayOk() (*POSTAdyenPayments201ResponseDataRelationshipsPaymentGateway, bool) {
	if o == nil || IsNil(o.PaymentGateway) {
		return nil, false
	}
	return o.PaymentGateway, true
}

// HasPaymentGateway returns a boolean if a field has been set.
func (o *POSTAdyenPayments201ResponseDataRelationships) HasPaymentGateway() bool {
	if o != nil && !IsNil(o.PaymentGateway) {
		return true
	}

	return false
}

// SetPaymentGateway gets a reference to the given POSTAdyenPayments201ResponseDataRelationshipsPaymentGateway and assigns it to the PaymentGateway field.
func (o *POSTAdyenPayments201ResponseDataRelationships) SetPaymentGateway(v POSTAdyenPayments201ResponseDataRelationshipsPaymentGateway) {
	o.PaymentGateway = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *POSTAdyenPayments201ResponseDataRelationships) GetVersions() POSTAddresses201ResponseDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret POSTAddresses201ResponseDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTAdyenPayments201ResponseDataRelationships) GetVersionsOk() (*POSTAddresses201ResponseDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *POSTAdyenPayments201ResponseDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given POSTAddresses201ResponseDataRelationshipsVersions and assigns it to the Versions field.
func (o *POSTAdyenPayments201ResponseDataRelationships) SetVersions(v POSTAddresses201ResponseDataRelationshipsVersions) {
	o.Versions = &v
}

func (o POSTAdyenPayments201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTAdyenPayments201ResponseDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.PaymentGateway) {
		toSerialize["payment_gateway"] = o.PaymentGateway
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	return toSerialize, nil
}

type NullablePOSTAdyenPayments201ResponseDataRelationships struct {
	value *POSTAdyenPayments201ResponseDataRelationships
	isSet bool
}

func (v NullablePOSTAdyenPayments201ResponseDataRelationships) Get() *POSTAdyenPayments201ResponseDataRelationships {
	return v.value
}

func (v *NullablePOSTAdyenPayments201ResponseDataRelationships) Set(val *POSTAdyenPayments201ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTAdyenPayments201ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTAdyenPayments201ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTAdyenPayments201ResponseDataRelationships(val *POSTAdyenPayments201ResponseDataRelationships) *NullablePOSTAdyenPayments201ResponseDataRelationships {
	return &NullablePOSTAdyenPayments201ResponseDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTAdyenPayments201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTAdyenPayments201ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
