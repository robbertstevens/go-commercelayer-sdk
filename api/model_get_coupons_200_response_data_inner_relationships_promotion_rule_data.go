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

// GETCoupons200ResponseDataInnerRelationshipsPromotionRuleData struct for GETCoupons200ResponseDataInnerRelationshipsPromotionRuleData
type GETCoupons200ResponseDataInnerRelationshipsPromotionRuleData struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource ID
	Id *string `json:"id,omitempty"`
}

// NewGETCoupons200ResponseDataInnerRelationshipsPromotionRuleData instantiates a new GETCoupons200ResponseDataInnerRelationshipsPromotionRuleData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCoupons200ResponseDataInnerRelationshipsPromotionRuleData() *GETCoupons200ResponseDataInnerRelationshipsPromotionRuleData {
	this := GETCoupons200ResponseDataInnerRelationshipsPromotionRuleData{}
	return &this
}

// NewGETCoupons200ResponseDataInnerRelationshipsPromotionRuleDataWithDefaults instantiates a new GETCoupons200ResponseDataInnerRelationshipsPromotionRuleData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCoupons200ResponseDataInnerRelationshipsPromotionRuleDataWithDefaults() *GETCoupons200ResponseDataInnerRelationshipsPromotionRuleData {
	this := GETCoupons200ResponseDataInnerRelationshipsPromotionRuleData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GETCoupons200ResponseDataInnerRelationshipsPromotionRuleData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCoupons200ResponseDataInnerRelationshipsPromotionRuleData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETCoupons200ResponseDataInnerRelationshipsPromotionRuleData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GETCoupons200ResponseDataInnerRelationshipsPromotionRuleData) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETCoupons200ResponseDataInnerRelationshipsPromotionRuleData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCoupons200ResponseDataInnerRelationshipsPromotionRuleData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETCoupons200ResponseDataInnerRelationshipsPromotionRuleData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETCoupons200ResponseDataInnerRelationshipsPromotionRuleData) SetId(v string) {
	o.Id = &v
}

func (o GETCoupons200ResponseDataInnerRelationshipsPromotionRuleData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETCoupons200ResponseDataInnerRelationshipsPromotionRuleData struct {
	value *GETCoupons200ResponseDataInnerRelationshipsPromotionRuleData
	isSet bool
}

func (v NullableGETCoupons200ResponseDataInnerRelationshipsPromotionRuleData) Get() *GETCoupons200ResponseDataInnerRelationshipsPromotionRuleData {
	return v.value
}

func (v *NullableGETCoupons200ResponseDataInnerRelationshipsPromotionRuleData) Set(val *GETCoupons200ResponseDataInnerRelationshipsPromotionRuleData) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCoupons200ResponseDataInnerRelationshipsPromotionRuleData) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCoupons200ResponseDataInnerRelationshipsPromotionRuleData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCoupons200ResponseDataInnerRelationshipsPromotionRuleData(val *GETCoupons200ResponseDataInnerRelationshipsPromotionRuleData) *NullableGETCoupons200ResponseDataInnerRelationshipsPromotionRuleData {
	return &NullableGETCoupons200ResponseDataInnerRelationshipsPromotionRuleData{value: val, isSet: true}
}

func (v NullableGETCoupons200ResponseDataInnerRelationshipsPromotionRuleData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCoupons200ResponseDataInnerRelationshipsPromotionRuleData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
