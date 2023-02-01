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

// GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRuleData struct for GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRuleData
type GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRuleData struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource ID
	Id *string `json:"id,omitempty"`
}

// NewGETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRuleData instantiates a new GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRuleData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRuleData() *GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRuleData {
	this := GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRuleData{}
	return &this
}

// NewGETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRuleDataWithDefaults instantiates a new GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRuleData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRuleDataWithDefaults() *GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRuleData {
	this := GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRuleData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRuleData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRuleData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRuleData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRuleData) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRuleData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRuleData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRuleData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRuleData) SetId(v string) {
	o.Id = &v
}

func (o GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRuleData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRuleData struct {
	value *GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRuleData
	isSet bool
}

func (v NullableGETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRuleData) Get() *GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRuleData {
	return v.value
}

func (v *NullableGETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRuleData) Set(val *GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRuleData) {
	v.value = val
	v.isSet = true
}

func (v NullableGETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRuleData) IsSet() bool {
	return v.isSet
}

func (v *NullableGETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRuleData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRuleData(val *GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRuleData) *NullableGETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRuleData {
	return &NullableGETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRuleData{value: val, isSet: true}
}

func (v NullableGETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRuleData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRuleData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
