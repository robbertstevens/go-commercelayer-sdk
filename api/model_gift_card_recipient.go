/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GiftCardRecipient struct for GiftCardRecipient
type GiftCardRecipient struct {
	Data GiftCardRecipientData `json:"data"`
}

// NewGiftCardRecipient instantiates a new GiftCardRecipient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGiftCardRecipient(data GiftCardRecipientData) *GiftCardRecipient {
	this := GiftCardRecipient{}
	this.Data = data
	return &this
}

// NewGiftCardRecipientWithDefaults instantiates a new GiftCardRecipient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGiftCardRecipientWithDefaults() *GiftCardRecipient {
	this := GiftCardRecipient{}
	return &this
}

// GetData returns the Data field value
func (o *GiftCardRecipient) GetData() GiftCardRecipientData {
	if o == nil {
		var ret GiftCardRecipientData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GiftCardRecipient) GetDataOk() (*GiftCardRecipientData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GiftCardRecipient) SetData(v GiftCardRecipientData) {
	o.Data = v
}

func (o GiftCardRecipient) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGiftCardRecipient struct {
	value *GiftCardRecipient
	isSet bool
}

func (v NullableGiftCardRecipient) Get() *GiftCardRecipient {
	return v.value
}

func (v *NullableGiftCardRecipient) Set(val *GiftCardRecipient) {
	v.value = val
	v.isSet = true
}

func (v NullableGiftCardRecipient) IsSet() bool {
	return v.isSet
}

func (v *NullableGiftCardRecipient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGiftCardRecipient(val *GiftCardRecipient) *NullableGiftCardRecipient {
	return &NullableGiftCardRecipient{value: val, isSet: true}
}

func (v NullableGiftCardRecipient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGiftCardRecipient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


