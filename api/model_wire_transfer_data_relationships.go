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

// checks if the WireTransferDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WireTransferDataRelationships{}

// WireTransferDataRelationships struct for WireTransferDataRelationships
type WireTransferDataRelationships struct {
	Order    *AdyenPaymentDataRelationshipsOrder `json:"order,omitempty"`
	Versions *AddressDataRelationshipsVersions   `json:"versions,omitempty"`
}

// NewWireTransferDataRelationships instantiates a new WireTransferDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWireTransferDataRelationships() *WireTransferDataRelationships {
	this := WireTransferDataRelationships{}
	return &this
}

// NewWireTransferDataRelationshipsWithDefaults instantiates a new WireTransferDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWireTransferDataRelationshipsWithDefaults() *WireTransferDataRelationships {
	this := WireTransferDataRelationships{}
	return &this
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *WireTransferDataRelationships) GetOrder() AdyenPaymentDataRelationshipsOrder {
	if o == nil || IsNil(o.Order) {
		var ret AdyenPaymentDataRelationshipsOrder
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WireTransferDataRelationships) GetOrderOk() (*AdyenPaymentDataRelationshipsOrder, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *WireTransferDataRelationships) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given AdyenPaymentDataRelationshipsOrder and assigns it to the Order field.
func (o *WireTransferDataRelationships) SetOrder(v AdyenPaymentDataRelationshipsOrder) {
	o.Order = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *WireTransferDataRelationships) GetVersions() AddressDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret AddressDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WireTransferDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *WireTransferDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given AddressDataRelationshipsVersions and assigns it to the Versions field.
func (o *WireTransferDataRelationships) SetVersions(v AddressDataRelationshipsVersions) {
	o.Versions = &v
}

func (o WireTransferDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WireTransferDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	return toSerialize, nil
}

type NullableWireTransferDataRelationships struct {
	value *WireTransferDataRelationships
	isSet bool
}

func (v NullableWireTransferDataRelationships) Get() *WireTransferDataRelationships {
	return v.value
}

func (v *NullableWireTransferDataRelationships) Set(val *WireTransferDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableWireTransferDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableWireTransferDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWireTransferDataRelationships(val *WireTransferDataRelationships) *NullableWireTransferDataRelationships {
	return &NullableWireTransferDataRelationships{value: val, isSet: true}
}

func (v NullableWireTransferDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWireTransferDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
