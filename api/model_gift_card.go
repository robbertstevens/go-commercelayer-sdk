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

// GiftCard struct for GiftCard
type GiftCard struct {
	Data *GiftCardData `json:"data,omitempty"`
}

// NewGiftCard instantiates a new GiftCard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGiftCard() *GiftCard {
	this := GiftCard{}
	return &this
}

// NewGiftCardWithDefaults instantiates a new GiftCard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGiftCardWithDefaults() *GiftCard {
	this := GiftCard{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GiftCard) GetData() GiftCardData {
	if o == nil || o.Data == nil {
		var ret GiftCardData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCard) GetDataOk() (*GiftCardData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GiftCard) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GiftCardData and assigns it to the Data field.
func (o *GiftCard) SetData(v GiftCardData) {
	o.Data = &v
}

func (o GiftCard) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGiftCard struct {
	value *GiftCard
	isSet bool
}

func (v NullableGiftCard) Get() *GiftCard {
	return v.value
}

func (v *NullableGiftCard) Set(val *GiftCard) {
	v.value = val
	v.isSet = true
}

func (v NullableGiftCard) IsSet() bool {
	return v.isSet
}

func (v *NullableGiftCard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGiftCard(val *GiftCard) *NullableGiftCard {
	return &NullableGiftCard{value: val, isSet: true}
}

func (v NullableGiftCard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGiftCard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
