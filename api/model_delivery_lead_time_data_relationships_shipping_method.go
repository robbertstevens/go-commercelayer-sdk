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

// DeliveryLeadTimeDataRelationshipsShippingMethod struct for DeliveryLeadTimeDataRelationshipsShippingMethod
type DeliveryLeadTimeDataRelationshipsShippingMethod struct {
	Data DeliveryLeadTimeDataRelationshipsShippingMethodData `json:"data"`
}

// NewDeliveryLeadTimeDataRelationshipsShippingMethod instantiates a new DeliveryLeadTimeDataRelationshipsShippingMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryLeadTimeDataRelationshipsShippingMethod(data DeliveryLeadTimeDataRelationshipsShippingMethodData) *DeliveryLeadTimeDataRelationshipsShippingMethod {
	this := DeliveryLeadTimeDataRelationshipsShippingMethod{}
	this.Data = data
	return &this
}

// NewDeliveryLeadTimeDataRelationshipsShippingMethodWithDefaults instantiates a new DeliveryLeadTimeDataRelationshipsShippingMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryLeadTimeDataRelationshipsShippingMethodWithDefaults() *DeliveryLeadTimeDataRelationshipsShippingMethod {
	this := DeliveryLeadTimeDataRelationshipsShippingMethod{}
	return &this
}

// GetData returns the Data field value
func (o *DeliveryLeadTimeDataRelationshipsShippingMethod) GetData() DeliveryLeadTimeDataRelationshipsShippingMethodData {
	if o == nil {
		var ret DeliveryLeadTimeDataRelationshipsShippingMethodData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *DeliveryLeadTimeDataRelationshipsShippingMethod) GetDataOk() (*DeliveryLeadTimeDataRelationshipsShippingMethodData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *DeliveryLeadTimeDataRelationshipsShippingMethod) SetData(v DeliveryLeadTimeDataRelationshipsShippingMethodData) {
	o.Data = v
}

func (o DeliveryLeadTimeDataRelationshipsShippingMethod) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableDeliveryLeadTimeDataRelationshipsShippingMethod struct {
	value *DeliveryLeadTimeDataRelationshipsShippingMethod
	isSet bool
}

func (v NullableDeliveryLeadTimeDataRelationshipsShippingMethod) Get() *DeliveryLeadTimeDataRelationshipsShippingMethod {
	return v.value
}

func (v *NullableDeliveryLeadTimeDataRelationshipsShippingMethod) Set(val *DeliveryLeadTimeDataRelationshipsShippingMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryLeadTimeDataRelationshipsShippingMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryLeadTimeDataRelationshipsShippingMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryLeadTimeDataRelationshipsShippingMethod(val *DeliveryLeadTimeDataRelationshipsShippingMethod) *NullableDeliveryLeadTimeDataRelationshipsShippingMethod {
	return &NullableDeliveryLeadTimeDataRelationshipsShippingMethod{value: val, isSet: true}
}

func (v NullableDeliveryLeadTimeDataRelationshipsShippingMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryLeadTimeDataRelationshipsShippingMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


