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

// checks if the OrderSubscriptionItemCreateDataRelationshipsAdjustment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderSubscriptionItemCreateDataRelationshipsAdjustment{}

// OrderSubscriptionItemCreateDataRelationshipsAdjustment struct for OrderSubscriptionItemCreateDataRelationshipsAdjustment
type OrderSubscriptionItemCreateDataRelationshipsAdjustment struct {
	Data LineItemDataRelationshipsAdjustmentData `json:"data"`
}

// NewOrderSubscriptionItemCreateDataRelationshipsAdjustment instantiates a new OrderSubscriptionItemCreateDataRelationshipsAdjustment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderSubscriptionItemCreateDataRelationshipsAdjustment(data LineItemDataRelationshipsAdjustmentData) *OrderSubscriptionItemCreateDataRelationshipsAdjustment {
	this := OrderSubscriptionItemCreateDataRelationshipsAdjustment{}
	this.Data = data
	return &this
}

// NewOrderSubscriptionItemCreateDataRelationshipsAdjustmentWithDefaults instantiates a new OrderSubscriptionItemCreateDataRelationshipsAdjustment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderSubscriptionItemCreateDataRelationshipsAdjustmentWithDefaults() *OrderSubscriptionItemCreateDataRelationshipsAdjustment {
	this := OrderSubscriptionItemCreateDataRelationshipsAdjustment{}
	return &this
}

// GetData returns the Data field value
func (o *OrderSubscriptionItemCreateDataRelationshipsAdjustment) GetData() LineItemDataRelationshipsAdjustmentData {
	if o == nil {
		var ret LineItemDataRelationshipsAdjustmentData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *OrderSubscriptionItemCreateDataRelationshipsAdjustment) GetDataOk() (*LineItemDataRelationshipsAdjustmentData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *OrderSubscriptionItemCreateDataRelationshipsAdjustment) SetData(v LineItemDataRelationshipsAdjustmentData) {
	o.Data = v
}

func (o OrderSubscriptionItemCreateDataRelationshipsAdjustment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderSubscriptionItemCreateDataRelationshipsAdjustment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableOrderSubscriptionItemCreateDataRelationshipsAdjustment struct {
	value *OrderSubscriptionItemCreateDataRelationshipsAdjustment
	isSet bool
}

func (v NullableOrderSubscriptionItemCreateDataRelationshipsAdjustment) Get() *OrderSubscriptionItemCreateDataRelationshipsAdjustment {
	return v.value
}

func (v *NullableOrderSubscriptionItemCreateDataRelationshipsAdjustment) Set(val *OrderSubscriptionItemCreateDataRelationshipsAdjustment) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderSubscriptionItemCreateDataRelationshipsAdjustment) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderSubscriptionItemCreateDataRelationshipsAdjustment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderSubscriptionItemCreateDataRelationshipsAdjustment(val *OrderSubscriptionItemCreateDataRelationshipsAdjustment) *NullableOrderSubscriptionItemCreateDataRelationshipsAdjustment {
	return &NullableOrderSubscriptionItemCreateDataRelationshipsAdjustment{value: val, isSet: true}
}

func (v NullableOrderSubscriptionItemCreateDataRelationshipsAdjustment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderSubscriptionItemCreateDataRelationshipsAdjustment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
