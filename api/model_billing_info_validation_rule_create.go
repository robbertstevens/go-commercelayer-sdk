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

// checks if the BillingInfoValidationRuleCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingInfoValidationRuleCreate{}

// BillingInfoValidationRuleCreate struct for BillingInfoValidationRuleCreate
type BillingInfoValidationRuleCreate struct {
	Data POSTBillingInfoValidationRulesRequestData `json:"data"`
}

// NewBillingInfoValidationRuleCreate instantiates a new BillingInfoValidationRuleCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingInfoValidationRuleCreate(data POSTBillingInfoValidationRulesRequestData) *BillingInfoValidationRuleCreate {
	this := BillingInfoValidationRuleCreate{}
	this.Data = data
	return &this
}

// NewBillingInfoValidationRuleCreateWithDefaults instantiates a new BillingInfoValidationRuleCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingInfoValidationRuleCreateWithDefaults() *BillingInfoValidationRuleCreate {
	this := BillingInfoValidationRuleCreate{}
	return &this
}

// GetData returns the Data field value
func (o *BillingInfoValidationRuleCreate) GetData() POSTBillingInfoValidationRulesRequestData {
	if o == nil {
		var ret POSTBillingInfoValidationRulesRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BillingInfoValidationRuleCreate) GetDataOk() (*POSTBillingInfoValidationRulesRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BillingInfoValidationRuleCreate) SetData(v POSTBillingInfoValidationRulesRequestData) {
	o.Data = v
}

func (o BillingInfoValidationRuleCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingInfoValidationRuleCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableBillingInfoValidationRuleCreate struct {
	value *BillingInfoValidationRuleCreate
	isSet bool
}

func (v NullableBillingInfoValidationRuleCreate) Get() *BillingInfoValidationRuleCreate {
	return v.value
}

func (v *NullableBillingInfoValidationRuleCreate) Set(val *BillingInfoValidationRuleCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingInfoValidationRuleCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingInfoValidationRuleCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingInfoValidationRuleCreate(val *BillingInfoValidationRuleCreate) *NullableBillingInfoValidationRuleCreate {
	return &NullableBillingInfoValidationRuleCreate{value: val, isSet: true}
}

func (v NullableBillingInfoValidationRuleCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingInfoValidationRuleCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
