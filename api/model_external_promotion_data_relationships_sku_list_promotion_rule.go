/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 4.1.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ExternalPromotionDataRelationshipsSkuListPromotionRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalPromotionDataRelationshipsSkuListPromotionRule{}

// ExternalPromotionDataRelationshipsSkuListPromotionRule struct for ExternalPromotionDataRelationshipsSkuListPromotionRule
type ExternalPromotionDataRelationshipsSkuListPromotionRule struct {
	Data *POSTExternalPromotionsRequestDataRelationshipsSkuListPromotionRuleData `json:"data,omitempty"`
}

// NewExternalPromotionDataRelationshipsSkuListPromotionRule instantiates a new ExternalPromotionDataRelationshipsSkuListPromotionRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalPromotionDataRelationshipsSkuListPromotionRule() *ExternalPromotionDataRelationshipsSkuListPromotionRule {
	this := ExternalPromotionDataRelationshipsSkuListPromotionRule{}
	return &this
}

// NewExternalPromotionDataRelationshipsSkuListPromotionRuleWithDefaults instantiates a new ExternalPromotionDataRelationshipsSkuListPromotionRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalPromotionDataRelationshipsSkuListPromotionRuleWithDefaults() *ExternalPromotionDataRelationshipsSkuListPromotionRule {
	this := ExternalPromotionDataRelationshipsSkuListPromotionRule{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ExternalPromotionDataRelationshipsSkuListPromotionRule) GetData() POSTExternalPromotionsRequestDataRelationshipsSkuListPromotionRuleData {
	if o == nil || IsNil(o.Data) {
		var ret POSTExternalPromotionsRequestDataRelationshipsSkuListPromotionRuleData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalPromotionDataRelationshipsSkuListPromotionRule) GetDataOk() (*POSTExternalPromotionsRequestDataRelationshipsSkuListPromotionRuleData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ExternalPromotionDataRelationshipsSkuListPromotionRule) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTExternalPromotionsRequestDataRelationshipsSkuListPromotionRuleData and assigns it to the Data field.
func (o *ExternalPromotionDataRelationshipsSkuListPromotionRule) SetData(v POSTExternalPromotionsRequestDataRelationshipsSkuListPromotionRuleData) {
	o.Data = &v
}

func (o ExternalPromotionDataRelationshipsSkuListPromotionRule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalPromotionDataRelationshipsSkuListPromotionRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableExternalPromotionDataRelationshipsSkuListPromotionRule struct {
	value *ExternalPromotionDataRelationshipsSkuListPromotionRule
	isSet bool
}

func (v NullableExternalPromotionDataRelationshipsSkuListPromotionRule) Get() *ExternalPromotionDataRelationshipsSkuListPromotionRule {
	return v.value
}

func (v *NullableExternalPromotionDataRelationshipsSkuListPromotionRule) Set(val *ExternalPromotionDataRelationshipsSkuListPromotionRule) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalPromotionDataRelationshipsSkuListPromotionRule) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalPromotionDataRelationshipsSkuListPromotionRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalPromotionDataRelationshipsSkuListPromotionRule(val *ExternalPromotionDataRelationshipsSkuListPromotionRule) *NullableExternalPromotionDataRelationshipsSkuListPromotionRule {
	return &NullableExternalPromotionDataRelationshipsSkuListPromotionRule{value: val, isSet: true}
}

func (v NullableExternalPromotionDataRelationshipsSkuListPromotionRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalPromotionDataRelationshipsSkuListPromotionRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
