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

// checks if the GiftCardUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GiftCardUpdate{}

// GiftCardUpdate struct for GiftCardUpdate
type GiftCardUpdate struct {
	Data PATCHGiftCardsGiftCardIdRequestData `json:"data"`
}

// NewGiftCardUpdate instantiates a new GiftCardUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGiftCardUpdate(data PATCHGiftCardsGiftCardIdRequestData) *GiftCardUpdate {
	this := GiftCardUpdate{}
	this.Data = data
	return &this
}

// NewGiftCardUpdateWithDefaults instantiates a new GiftCardUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGiftCardUpdateWithDefaults() *GiftCardUpdate {
	this := GiftCardUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *GiftCardUpdate) GetData() PATCHGiftCardsGiftCardIdRequestData {
	if o == nil {
		var ret PATCHGiftCardsGiftCardIdRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GiftCardUpdate) GetDataOk() (*PATCHGiftCardsGiftCardIdRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GiftCardUpdate) SetData(v PATCHGiftCardsGiftCardIdRequestData) {
	o.Data = v
}

func (o GiftCardUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GiftCardUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableGiftCardUpdate struct {
	value *GiftCardUpdate
	isSet bool
}

func (v NullableGiftCardUpdate) Get() *GiftCardUpdate {
	return v.value
}

func (v *NullableGiftCardUpdate) Set(val *GiftCardUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableGiftCardUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableGiftCardUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGiftCardUpdate(val *GiftCardUpdate) *NullableGiftCardUpdate {
	return &NullableGiftCardUpdate{value: val, isSet: true}
}

func (v NullableGiftCardUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGiftCardUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
