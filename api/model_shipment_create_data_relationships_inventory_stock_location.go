/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.3.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ShipmentCreateDataRelationshipsInventoryStockLocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipmentCreateDataRelationshipsInventoryStockLocation{}

// ShipmentCreateDataRelationshipsInventoryStockLocation struct for ShipmentCreateDataRelationshipsInventoryStockLocation
type ShipmentCreateDataRelationshipsInventoryStockLocation struct {
	Data InventoryModelDataRelationshipsInventoryStockLocationsData `json:"data"`
}

// NewShipmentCreateDataRelationshipsInventoryStockLocation instantiates a new ShipmentCreateDataRelationshipsInventoryStockLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipmentCreateDataRelationshipsInventoryStockLocation(data InventoryModelDataRelationshipsInventoryStockLocationsData) *ShipmentCreateDataRelationshipsInventoryStockLocation {
	this := ShipmentCreateDataRelationshipsInventoryStockLocation{}
	this.Data = data
	return &this
}

// NewShipmentCreateDataRelationshipsInventoryStockLocationWithDefaults instantiates a new ShipmentCreateDataRelationshipsInventoryStockLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipmentCreateDataRelationshipsInventoryStockLocationWithDefaults() *ShipmentCreateDataRelationshipsInventoryStockLocation {
	this := ShipmentCreateDataRelationshipsInventoryStockLocation{}
	return &this
}

// GetData returns the Data field value
func (o *ShipmentCreateDataRelationshipsInventoryStockLocation) GetData() InventoryModelDataRelationshipsInventoryStockLocationsData {
	if o == nil {
		var ret InventoryModelDataRelationshipsInventoryStockLocationsData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ShipmentCreateDataRelationshipsInventoryStockLocation) GetDataOk() (*InventoryModelDataRelationshipsInventoryStockLocationsData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ShipmentCreateDataRelationshipsInventoryStockLocation) SetData(v InventoryModelDataRelationshipsInventoryStockLocationsData) {
	o.Data = v
}

func (o ShipmentCreateDataRelationshipsInventoryStockLocation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShipmentCreateDataRelationshipsInventoryStockLocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableShipmentCreateDataRelationshipsInventoryStockLocation struct {
	value *ShipmentCreateDataRelationshipsInventoryStockLocation
	isSet bool
}

func (v NullableShipmentCreateDataRelationshipsInventoryStockLocation) Get() *ShipmentCreateDataRelationshipsInventoryStockLocation {
	return v.value
}

func (v *NullableShipmentCreateDataRelationshipsInventoryStockLocation) Set(val *ShipmentCreateDataRelationshipsInventoryStockLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentCreateDataRelationshipsInventoryStockLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentCreateDataRelationshipsInventoryStockLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentCreateDataRelationshipsInventoryStockLocation(val *ShipmentCreateDataRelationshipsInventoryStockLocation) *NullableShipmentCreateDataRelationshipsInventoryStockLocation {
	return &NullableShipmentCreateDataRelationshipsInventoryStockLocation{value: val, isSet: true}
}

func (v NullableShipmentCreateDataRelationshipsInventoryStockLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipmentCreateDataRelationshipsInventoryStockLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
