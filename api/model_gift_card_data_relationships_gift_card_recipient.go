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

// GiftCardDataRelationshipsGiftCardRecipient struct for GiftCardDataRelationshipsGiftCardRecipient
type GiftCardDataRelationshipsGiftCardRecipient struct {
	Data GiftCardDataRelationshipsGiftCardRecipientData `json:"data"`
}

// NewGiftCardDataRelationshipsGiftCardRecipient instantiates a new GiftCardDataRelationshipsGiftCardRecipient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGiftCardDataRelationshipsGiftCardRecipient(data GiftCardDataRelationshipsGiftCardRecipientData) *GiftCardDataRelationshipsGiftCardRecipient {
	this := GiftCardDataRelationshipsGiftCardRecipient{}
	this.Data = data
	return &this
}

// NewGiftCardDataRelationshipsGiftCardRecipientWithDefaults instantiates a new GiftCardDataRelationshipsGiftCardRecipient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGiftCardDataRelationshipsGiftCardRecipientWithDefaults() *GiftCardDataRelationshipsGiftCardRecipient {
	this := GiftCardDataRelationshipsGiftCardRecipient{}
	return &this
}

// GetData returns the Data field value
func (o *GiftCardDataRelationshipsGiftCardRecipient) GetData() GiftCardDataRelationshipsGiftCardRecipientData {
	if o == nil {
		var ret GiftCardDataRelationshipsGiftCardRecipientData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GiftCardDataRelationshipsGiftCardRecipient) GetDataOk() (*GiftCardDataRelationshipsGiftCardRecipientData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GiftCardDataRelationshipsGiftCardRecipient) SetData(v GiftCardDataRelationshipsGiftCardRecipientData) {
	o.Data = v
}

func (o GiftCardDataRelationshipsGiftCardRecipient) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGiftCardDataRelationshipsGiftCardRecipient struct {
	value *GiftCardDataRelationshipsGiftCardRecipient
	isSet bool
}

func (v NullableGiftCardDataRelationshipsGiftCardRecipient) Get() *GiftCardDataRelationshipsGiftCardRecipient {
	return v.value
}

func (v *NullableGiftCardDataRelationshipsGiftCardRecipient) Set(val *GiftCardDataRelationshipsGiftCardRecipient) {
	v.value = val
	v.isSet = true
}

func (v NullableGiftCardDataRelationshipsGiftCardRecipient) IsSet() bool {
	return v.isSet
}

func (v *NullableGiftCardDataRelationshipsGiftCardRecipient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGiftCardDataRelationshipsGiftCardRecipient(val *GiftCardDataRelationshipsGiftCardRecipient) *NullableGiftCardDataRelationshipsGiftCardRecipient {
	return &NullableGiftCardDataRelationshipsGiftCardRecipient{value: val, isSet: true}
}

func (v NullableGiftCardDataRelationshipsGiftCardRecipient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGiftCardDataRelationshipsGiftCardRecipient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


