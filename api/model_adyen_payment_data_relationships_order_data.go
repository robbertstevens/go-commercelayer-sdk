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

// checks if the AdyenPaymentDataRelationshipsOrderData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdyenPaymentDataRelationshipsOrderData{}

// AdyenPaymentDataRelationshipsOrderData struct for AdyenPaymentDataRelationshipsOrderData
type AdyenPaymentDataRelationshipsOrderData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource's id
	Id interface{} `json:"id,omitempty"`
}

// NewAdyenPaymentDataRelationshipsOrderData instantiates a new AdyenPaymentDataRelationshipsOrderData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdyenPaymentDataRelationshipsOrderData() *AdyenPaymentDataRelationshipsOrderData {
	this := AdyenPaymentDataRelationshipsOrderData{}
	return &this
}

// NewAdyenPaymentDataRelationshipsOrderDataWithDefaults instantiates a new AdyenPaymentDataRelationshipsOrderData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdyenPaymentDataRelationshipsOrderDataWithDefaults() *AdyenPaymentDataRelationshipsOrderData {
	this := AdyenPaymentDataRelationshipsOrderData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdyenPaymentDataRelationshipsOrderData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdyenPaymentDataRelationshipsOrderData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AdyenPaymentDataRelationshipsOrderData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *AdyenPaymentDataRelationshipsOrderData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdyenPaymentDataRelationshipsOrderData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdyenPaymentDataRelationshipsOrderData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AdyenPaymentDataRelationshipsOrderData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *AdyenPaymentDataRelationshipsOrderData) SetId(v interface{}) {
	o.Id = v
}

func (o AdyenPaymentDataRelationshipsOrderData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdyenPaymentDataRelationshipsOrderData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableAdyenPaymentDataRelationshipsOrderData struct {
	value *AdyenPaymentDataRelationshipsOrderData
	isSet bool
}

func (v NullableAdyenPaymentDataRelationshipsOrderData) Get() *AdyenPaymentDataRelationshipsOrderData {
	return v.value
}

func (v *NullableAdyenPaymentDataRelationshipsOrderData) Set(val *AdyenPaymentDataRelationshipsOrderData) {
	v.value = val
	v.isSet = true
}

func (v NullableAdyenPaymentDataRelationshipsOrderData) IsSet() bool {
	return v.isSet
}

func (v *NullableAdyenPaymentDataRelationshipsOrderData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdyenPaymentDataRelationshipsOrderData(val *AdyenPaymentDataRelationshipsOrderData) *NullableAdyenPaymentDataRelationshipsOrderData {
	return &NullableAdyenPaymentDataRelationshipsOrderData{value: val, isSet: true}
}

func (v NullableAdyenPaymentDataRelationshipsOrderData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdyenPaymentDataRelationshipsOrderData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
