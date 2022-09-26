/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsDataInner struct for GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsDataInner
type GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsDataInner struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource ID
	Id *string `json:"id,omitempty"`
}

// NewGETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsDataInner instantiates a new GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsDataInner() *GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsDataInner {
	this := GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsDataInner{}
	return &this
}

// NewGETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsDataInnerWithDefaults instantiates a new GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsDataInnerWithDefaults() *GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsDataInner {
	this := GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsDataInner{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsDataInner) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsDataInner) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsDataInner) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsDataInner) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsDataInner) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsDataInner) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsDataInner) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsDataInner) SetId(v string) {
	o.Id = &v
}

func (o GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsDataInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsDataInner struct {
	value *GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsDataInner
	isSet bool
}

func (v NullableGETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsDataInner) Get() *GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsDataInner {
	return v.value
}

func (v *NullableGETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsDataInner) Set(val *GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsDataInner(val *GETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsDataInner) *NullableGETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsDataInner {
	return &NullableGETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsDataInner{value: val, isSet: true}
}

func (v NullableGETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETOrders200ResponseDataInnerRelationshipsAvailablePaymentMethodsDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
