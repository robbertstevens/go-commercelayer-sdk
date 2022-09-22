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

// OrderSubscriptionDataRelationshipsOrderCopies struct for OrderSubscriptionDataRelationshipsOrderCopies
type OrderSubscriptionDataRelationshipsOrderCopies struct {
	Data OrderSubscriptionDataRelationshipsOrderCopiesData `json:"data"`
}

// NewOrderSubscriptionDataRelationshipsOrderCopies instantiates a new OrderSubscriptionDataRelationshipsOrderCopies object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderSubscriptionDataRelationshipsOrderCopies(data OrderSubscriptionDataRelationshipsOrderCopiesData) *OrderSubscriptionDataRelationshipsOrderCopies {
	this := OrderSubscriptionDataRelationshipsOrderCopies{}
	this.Data = data
	return &this
}

// NewOrderSubscriptionDataRelationshipsOrderCopiesWithDefaults instantiates a new OrderSubscriptionDataRelationshipsOrderCopies object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderSubscriptionDataRelationshipsOrderCopiesWithDefaults() *OrderSubscriptionDataRelationshipsOrderCopies {
	this := OrderSubscriptionDataRelationshipsOrderCopies{}
	return &this
}

// GetData returns the Data field value
func (o *OrderSubscriptionDataRelationshipsOrderCopies) GetData() OrderSubscriptionDataRelationshipsOrderCopiesData {
	if o == nil {
		var ret OrderSubscriptionDataRelationshipsOrderCopiesData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *OrderSubscriptionDataRelationshipsOrderCopies) GetDataOk() (*OrderSubscriptionDataRelationshipsOrderCopiesData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *OrderSubscriptionDataRelationshipsOrderCopies) SetData(v OrderSubscriptionDataRelationshipsOrderCopiesData) {
	o.Data = v
}

func (o OrderSubscriptionDataRelationshipsOrderCopies) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableOrderSubscriptionDataRelationshipsOrderCopies struct {
	value *OrderSubscriptionDataRelationshipsOrderCopies
	isSet bool
}

func (v NullableOrderSubscriptionDataRelationshipsOrderCopies) Get() *OrderSubscriptionDataRelationshipsOrderCopies {
	return v.value
}

func (v *NullableOrderSubscriptionDataRelationshipsOrderCopies) Set(val *OrderSubscriptionDataRelationshipsOrderCopies) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderSubscriptionDataRelationshipsOrderCopies) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderSubscriptionDataRelationshipsOrderCopies) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderSubscriptionDataRelationshipsOrderCopies(val *OrderSubscriptionDataRelationshipsOrderCopies) *NullableOrderSubscriptionDataRelationshipsOrderCopies {
	return &NullableOrderSubscriptionDataRelationshipsOrderCopies{value: val, isSet: true}
}

func (v NullableOrderSubscriptionDataRelationshipsOrderCopies) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderSubscriptionDataRelationshipsOrderCopies) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
