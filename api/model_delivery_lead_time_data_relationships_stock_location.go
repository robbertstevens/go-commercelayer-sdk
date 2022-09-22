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

// DeliveryLeadTimeDataRelationshipsStockLocation struct for DeliveryLeadTimeDataRelationshipsStockLocation
type DeliveryLeadTimeDataRelationshipsStockLocation struct {
	Data DeliveryLeadTimeDataRelationshipsStockLocationData `json:"data"`
}

// NewDeliveryLeadTimeDataRelationshipsStockLocation instantiates a new DeliveryLeadTimeDataRelationshipsStockLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryLeadTimeDataRelationshipsStockLocation(data DeliveryLeadTimeDataRelationshipsStockLocationData) *DeliveryLeadTimeDataRelationshipsStockLocation {
	this := DeliveryLeadTimeDataRelationshipsStockLocation{}
	this.Data = data
	return &this
}

// NewDeliveryLeadTimeDataRelationshipsStockLocationWithDefaults instantiates a new DeliveryLeadTimeDataRelationshipsStockLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryLeadTimeDataRelationshipsStockLocationWithDefaults() *DeliveryLeadTimeDataRelationshipsStockLocation {
	this := DeliveryLeadTimeDataRelationshipsStockLocation{}
	return &this
}

// GetData returns the Data field value
func (o *DeliveryLeadTimeDataRelationshipsStockLocation) GetData() DeliveryLeadTimeDataRelationshipsStockLocationData {
	if o == nil {
		var ret DeliveryLeadTimeDataRelationshipsStockLocationData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *DeliveryLeadTimeDataRelationshipsStockLocation) GetDataOk() (*DeliveryLeadTimeDataRelationshipsStockLocationData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *DeliveryLeadTimeDataRelationshipsStockLocation) SetData(v DeliveryLeadTimeDataRelationshipsStockLocationData) {
	o.Data = v
}

func (o DeliveryLeadTimeDataRelationshipsStockLocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableDeliveryLeadTimeDataRelationshipsStockLocation struct {
	value *DeliveryLeadTimeDataRelationshipsStockLocation
	isSet bool
}

func (v NullableDeliveryLeadTimeDataRelationshipsStockLocation) Get() *DeliveryLeadTimeDataRelationshipsStockLocation {
	return v.value
}

func (v *NullableDeliveryLeadTimeDataRelationshipsStockLocation) Set(val *DeliveryLeadTimeDataRelationshipsStockLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryLeadTimeDataRelationshipsStockLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryLeadTimeDataRelationshipsStockLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryLeadTimeDataRelationshipsStockLocation(val *DeliveryLeadTimeDataRelationshipsStockLocation) *NullableDeliveryLeadTimeDataRelationshipsStockLocation {
	return &NullableDeliveryLeadTimeDataRelationshipsStockLocation{value: val, isSet: true}
}

func (v NullableDeliveryLeadTimeDataRelationshipsStockLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryLeadTimeDataRelationshipsStockLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
