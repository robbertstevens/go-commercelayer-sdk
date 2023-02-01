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

// DeliveryLeadTimeCreateDataRelationshipsStockLocation struct for DeliveryLeadTimeCreateDataRelationshipsStockLocation
type DeliveryLeadTimeCreateDataRelationshipsStockLocation struct {
	Data DeliveryLeadTimeDataRelationshipsStockLocationData `json:"data"`
}

// NewDeliveryLeadTimeCreateDataRelationshipsStockLocation instantiates a new DeliveryLeadTimeCreateDataRelationshipsStockLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryLeadTimeCreateDataRelationshipsStockLocation(data DeliveryLeadTimeDataRelationshipsStockLocationData) *DeliveryLeadTimeCreateDataRelationshipsStockLocation {
	this := DeliveryLeadTimeCreateDataRelationshipsStockLocation{}
	this.Data = data
	return &this
}

// NewDeliveryLeadTimeCreateDataRelationshipsStockLocationWithDefaults instantiates a new DeliveryLeadTimeCreateDataRelationshipsStockLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryLeadTimeCreateDataRelationshipsStockLocationWithDefaults() *DeliveryLeadTimeCreateDataRelationshipsStockLocation {
	this := DeliveryLeadTimeCreateDataRelationshipsStockLocation{}
	return &this
}

// GetData returns the Data field value
func (o *DeliveryLeadTimeCreateDataRelationshipsStockLocation) GetData() DeliveryLeadTimeDataRelationshipsStockLocationData {
	if o == nil {
		var ret DeliveryLeadTimeDataRelationshipsStockLocationData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *DeliveryLeadTimeCreateDataRelationshipsStockLocation) GetDataOk() (*DeliveryLeadTimeDataRelationshipsStockLocationData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *DeliveryLeadTimeCreateDataRelationshipsStockLocation) SetData(v DeliveryLeadTimeDataRelationshipsStockLocationData) {
	o.Data = v
}

func (o DeliveryLeadTimeCreateDataRelationshipsStockLocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableDeliveryLeadTimeCreateDataRelationshipsStockLocation struct {
	value *DeliveryLeadTimeCreateDataRelationshipsStockLocation
	isSet bool
}

func (v NullableDeliveryLeadTimeCreateDataRelationshipsStockLocation) Get() *DeliveryLeadTimeCreateDataRelationshipsStockLocation {
	return v.value
}

func (v *NullableDeliveryLeadTimeCreateDataRelationshipsStockLocation) Set(val *DeliveryLeadTimeCreateDataRelationshipsStockLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryLeadTimeCreateDataRelationshipsStockLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryLeadTimeCreateDataRelationshipsStockLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryLeadTimeCreateDataRelationshipsStockLocation(val *DeliveryLeadTimeCreateDataRelationshipsStockLocation) *NullableDeliveryLeadTimeCreateDataRelationshipsStockLocation {
	return &NullableDeliveryLeadTimeCreateDataRelationshipsStockLocation{value: val, isSet: true}
}

func (v NullableDeliveryLeadTimeCreateDataRelationshipsStockLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryLeadTimeCreateDataRelationshipsStockLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
