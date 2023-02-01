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

// GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsData struct for GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsData
type GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsData struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource ID
	Id *string `json:"id,omitempty"`
}

// NewGETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsData instantiates a new GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsData() *GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsData {
	this := GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsData{}
	return &this
}

// NewGETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsDataWithDefaults instantiates a new GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsDataWithDefaults() *GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsData {
	this := GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsData) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsData) SetId(v string) {
	o.Id = &v
}

func (o GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsData struct {
	value *GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsData
	isSet bool
}

func (v NullableGETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsData) Get() *GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsData {
	return v.value
}

func (v *NullableGETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsData) Set(val *GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsData) {
	v.value = val
	v.isSet = true
}

func (v NullableGETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsData) IsSet() bool {
	return v.isSet
}

func (v *NullableGETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsData(val *GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsData) *NullableGETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsData {
	return &NullableGETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsData{value: val, isSet: true}
}

func (v NullableGETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethodsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
