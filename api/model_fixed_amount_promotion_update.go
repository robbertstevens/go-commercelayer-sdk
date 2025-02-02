/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.3.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the FixedAmountPromotionUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FixedAmountPromotionUpdate{}

// FixedAmountPromotionUpdate struct for FixedAmountPromotionUpdate
type FixedAmountPromotionUpdate struct {
	Data FixedAmountPromotionUpdateData `json:"data"`
}

// NewFixedAmountPromotionUpdate instantiates a new FixedAmountPromotionUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFixedAmountPromotionUpdate(data FixedAmountPromotionUpdateData) *FixedAmountPromotionUpdate {
	this := FixedAmountPromotionUpdate{}
	this.Data = data
	return &this
}

// NewFixedAmountPromotionUpdateWithDefaults instantiates a new FixedAmountPromotionUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFixedAmountPromotionUpdateWithDefaults() *FixedAmountPromotionUpdate {
	this := FixedAmountPromotionUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *FixedAmountPromotionUpdate) GetData() FixedAmountPromotionUpdateData {
	if o == nil {
		var ret FixedAmountPromotionUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *FixedAmountPromotionUpdate) GetDataOk() (*FixedAmountPromotionUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *FixedAmountPromotionUpdate) SetData(v FixedAmountPromotionUpdateData) {
	o.Data = v
}

func (o FixedAmountPromotionUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FixedAmountPromotionUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableFixedAmountPromotionUpdate struct {
	value *FixedAmountPromotionUpdate
	isSet bool
}

func (v NullableFixedAmountPromotionUpdate) Get() *FixedAmountPromotionUpdate {
	return v.value
}

func (v *NullableFixedAmountPromotionUpdate) Set(val *FixedAmountPromotionUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableFixedAmountPromotionUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableFixedAmountPromotionUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFixedAmountPromotionUpdate(val *FixedAmountPromotionUpdate) *NullableFixedAmountPromotionUpdate {
	return &NullableFixedAmountPromotionUpdate{value: val, isSet: true}
}

func (v NullableFixedAmountPromotionUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFixedAmountPromotionUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
