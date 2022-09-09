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

// InventoryModelDataRelationshipsInventoryStockLocations struct for InventoryModelDataRelationshipsInventoryStockLocations
type InventoryModelDataRelationshipsInventoryStockLocations struct {
	Data InventoryModelDataRelationshipsInventoryStockLocationsData `json:"data"`
}

// NewInventoryModelDataRelationshipsInventoryStockLocations instantiates a new InventoryModelDataRelationshipsInventoryStockLocations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryModelDataRelationshipsInventoryStockLocations(data InventoryModelDataRelationshipsInventoryStockLocationsData) *InventoryModelDataRelationshipsInventoryStockLocations {
	this := InventoryModelDataRelationshipsInventoryStockLocations{}
	this.Data = data
	return &this
}

// NewInventoryModelDataRelationshipsInventoryStockLocationsWithDefaults instantiates a new InventoryModelDataRelationshipsInventoryStockLocations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryModelDataRelationshipsInventoryStockLocationsWithDefaults() *InventoryModelDataRelationshipsInventoryStockLocations {
	this := InventoryModelDataRelationshipsInventoryStockLocations{}
	return &this
}

// GetData returns the Data field value
func (o *InventoryModelDataRelationshipsInventoryStockLocations) GetData() InventoryModelDataRelationshipsInventoryStockLocationsData {
	if o == nil {
		var ret InventoryModelDataRelationshipsInventoryStockLocationsData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *InventoryModelDataRelationshipsInventoryStockLocations) GetDataOk() (*InventoryModelDataRelationshipsInventoryStockLocationsData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *InventoryModelDataRelationshipsInventoryStockLocations) SetData(v InventoryModelDataRelationshipsInventoryStockLocationsData) {
	o.Data = v
}

func (o InventoryModelDataRelationshipsInventoryStockLocations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableInventoryModelDataRelationshipsInventoryStockLocations struct {
	value *InventoryModelDataRelationshipsInventoryStockLocations
	isSet bool
}

func (v NullableInventoryModelDataRelationshipsInventoryStockLocations) Get() *InventoryModelDataRelationshipsInventoryStockLocations {
	return v.value
}

func (v *NullableInventoryModelDataRelationshipsInventoryStockLocations) Set(val *InventoryModelDataRelationshipsInventoryStockLocations) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryModelDataRelationshipsInventoryStockLocations) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryModelDataRelationshipsInventoryStockLocations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryModelDataRelationshipsInventoryStockLocations(val *InventoryModelDataRelationshipsInventoryStockLocations) *NullableInventoryModelDataRelationshipsInventoryStockLocations {
	return &NullableInventoryModelDataRelationshipsInventoryStockLocations{value: val, isSet: true}
}

func (v NullableInventoryModelDataRelationshipsInventoryStockLocations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryModelDataRelationshipsInventoryStockLocations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


