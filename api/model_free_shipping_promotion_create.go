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

// checks if the FreeShippingPromotionCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FreeShippingPromotionCreate{}

// FreeShippingPromotionCreate struct for FreeShippingPromotionCreate
type FreeShippingPromotionCreate struct {
	Data FreeShippingPromotionCreateData `json:"data"`
}

// NewFreeShippingPromotionCreate instantiates a new FreeShippingPromotionCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFreeShippingPromotionCreate(data FreeShippingPromotionCreateData) *FreeShippingPromotionCreate {
	this := FreeShippingPromotionCreate{}
	this.Data = data
	return &this
}

// NewFreeShippingPromotionCreateWithDefaults instantiates a new FreeShippingPromotionCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFreeShippingPromotionCreateWithDefaults() *FreeShippingPromotionCreate {
	this := FreeShippingPromotionCreate{}
	return &this
}

// GetData returns the Data field value
func (o *FreeShippingPromotionCreate) GetData() FreeShippingPromotionCreateData {
	if o == nil {
		var ret FreeShippingPromotionCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *FreeShippingPromotionCreate) GetDataOk() (*FreeShippingPromotionCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *FreeShippingPromotionCreate) SetData(v FreeShippingPromotionCreateData) {
	o.Data = v
}

func (o FreeShippingPromotionCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FreeShippingPromotionCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableFreeShippingPromotionCreate struct {
	value *FreeShippingPromotionCreate
	isSet bool
}

func (v NullableFreeShippingPromotionCreate) Get() *FreeShippingPromotionCreate {
	return v.value
}

func (v *NullableFreeShippingPromotionCreate) Set(val *FreeShippingPromotionCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableFreeShippingPromotionCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableFreeShippingPromotionCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFreeShippingPromotionCreate(val *FreeShippingPromotionCreate) *NullableFreeShippingPromotionCreate {
	return &NullableFreeShippingPromotionCreate{value: val, isSet: true}
}

func (v NullableFreeShippingPromotionCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFreeShippingPromotionCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
