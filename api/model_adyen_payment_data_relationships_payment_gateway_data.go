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

// AdyenPaymentDataRelationshipsPaymentGatewayData struct for AdyenPaymentDataRelationshipsPaymentGatewayData
type AdyenPaymentDataRelationshipsPaymentGatewayData struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource's id
	Id *string `json:"id,omitempty"`
}

// NewAdyenPaymentDataRelationshipsPaymentGatewayData instantiates a new AdyenPaymentDataRelationshipsPaymentGatewayData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdyenPaymentDataRelationshipsPaymentGatewayData() *AdyenPaymentDataRelationshipsPaymentGatewayData {
	this := AdyenPaymentDataRelationshipsPaymentGatewayData{}
	var type_ string = "payment_gateways"
	this.Type = &type_
	return &this
}

// NewAdyenPaymentDataRelationshipsPaymentGatewayDataWithDefaults instantiates a new AdyenPaymentDataRelationshipsPaymentGatewayData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdyenPaymentDataRelationshipsPaymentGatewayDataWithDefaults() *AdyenPaymentDataRelationshipsPaymentGatewayData {
	this := AdyenPaymentDataRelationshipsPaymentGatewayData{}
	var type_ string = "payment_gateways"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AdyenPaymentDataRelationshipsPaymentGatewayData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdyenPaymentDataRelationshipsPaymentGatewayData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AdyenPaymentDataRelationshipsPaymentGatewayData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AdyenPaymentDataRelationshipsPaymentGatewayData) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AdyenPaymentDataRelationshipsPaymentGatewayData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdyenPaymentDataRelationshipsPaymentGatewayData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AdyenPaymentDataRelationshipsPaymentGatewayData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AdyenPaymentDataRelationshipsPaymentGatewayData) SetId(v string) {
	o.Id = &v
}

func (o AdyenPaymentDataRelationshipsPaymentGatewayData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableAdyenPaymentDataRelationshipsPaymentGatewayData struct {
	value *AdyenPaymentDataRelationshipsPaymentGatewayData
	isSet bool
}

func (v NullableAdyenPaymentDataRelationshipsPaymentGatewayData) Get() *AdyenPaymentDataRelationshipsPaymentGatewayData {
	return v.value
}

func (v *NullableAdyenPaymentDataRelationshipsPaymentGatewayData) Set(val *AdyenPaymentDataRelationshipsPaymentGatewayData) {
	v.value = val
	v.isSet = true
}

func (v NullableAdyenPaymentDataRelationshipsPaymentGatewayData) IsSet() bool {
	return v.isSet
}

func (v *NullableAdyenPaymentDataRelationshipsPaymentGatewayData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdyenPaymentDataRelationshipsPaymentGatewayData(val *AdyenPaymentDataRelationshipsPaymentGatewayData) *NullableAdyenPaymentDataRelationshipsPaymentGatewayData {
	return &NullableAdyenPaymentDataRelationshipsPaymentGatewayData{value: val, isSet: true}
}

func (v NullableAdyenPaymentDataRelationshipsPaymentGatewayData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdyenPaymentDataRelationshipsPaymentGatewayData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


