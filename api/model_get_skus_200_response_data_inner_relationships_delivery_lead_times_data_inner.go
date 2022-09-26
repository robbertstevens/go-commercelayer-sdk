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

// GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesDataInner struct for GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesDataInner
type GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesDataInner struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource ID
	Id *string `json:"id,omitempty"`
}

// NewGETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesDataInner instantiates a new GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesDataInner() *GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesDataInner {
	this := GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesDataInner{}
	return &this
}

// NewGETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesDataInnerWithDefaults instantiates a new GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesDataInnerWithDefaults() *GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesDataInner {
	this := GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesDataInner{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesDataInner) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesDataInner) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesDataInner) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesDataInner) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesDataInner) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesDataInner) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesDataInner) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesDataInner) SetId(v string) {
	o.Id = &v
}

func (o GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesDataInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesDataInner struct {
	value *GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesDataInner
	isSet bool
}

func (v NullableGETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesDataInner) Get() *GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesDataInner {
	return v.value
}

func (v *NullableGETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesDataInner) Set(val *GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesDataInner(val *GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesDataInner) *NullableGETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesDataInner {
	return &NullableGETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesDataInner{value: val, isSet: true}
}

func (v NullableGETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
