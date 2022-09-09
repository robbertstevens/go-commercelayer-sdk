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

// InStockSubscriptionUpdate struct for InStockSubscriptionUpdate
type InStockSubscriptionUpdate struct {
	Data InStockSubscriptionUpdateData `json:"data"`
}

// NewInStockSubscriptionUpdate instantiates a new InStockSubscriptionUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInStockSubscriptionUpdate(data InStockSubscriptionUpdateData) *InStockSubscriptionUpdate {
	this := InStockSubscriptionUpdate{}
	this.Data = data
	return &this
}

// NewInStockSubscriptionUpdateWithDefaults instantiates a new InStockSubscriptionUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInStockSubscriptionUpdateWithDefaults() *InStockSubscriptionUpdate {
	this := InStockSubscriptionUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *InStockSubscriptionUpdate) GetData() InStockSubscriptionUpdateData {
	if o == nil {
		var ret InStockSubscriptionUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *InStockSubscriptionUpdate) GetDataOk() (*InStockSubscriptionUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *InStockSubscriptionUpdate) SetData(v InStockSubscriptionUpdateData) {
	o.Data = v
}

func (o InStockSubscriptionUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableInStockSubscriptionUpdate struct {
	value *InStockSubscriptionUpdate
	isSet bool
}

func (v NullableInStockSubscriptionUpdate) Get() *InStockSubscriptionUpdate {
	return v.value
}

func (v *NullableInStockSubscriptionUpdate) Set(val *InStockSubscriptionUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableInStockSubscriptionUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableInStockSubscriptionUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInStockSubscriptionUpdate(val *InStockSubscriptionUpdate) *NullableInStockSubscriptionUpdate {
	return &NullableInStockSubscriptionUpdate{value: val, isSet: true}
}

func (v NullableInStockSubscriptionUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInStockSubscriptionUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


