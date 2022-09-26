/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ShipmentDataRelationshipsCarrierAccounts struct for ShipmentDataRelationshipsCarrierAccounts
type ShipmentDataRelationshipsCarrierAccounts struct {
	Data ShipmentDataRelationshipsCarrierAccountsData `json:"data"`
}

// NewShipmentDataRelationshipsCarrierAccounts instantiates a new ShipmentDataRelationshipsCarrierAccounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipmentDataRelationshipsCarrierAccounts(data ShipmentDataRelationshipsCarrierAccountsData) *ShipmentDataRelationshipsCarrierAccounts {
	this := ShipmentDataRelationshipsCarrierAccounts{}
	this.Data = data
	return &this
}

// NewShipmentDataRelationshipsCarrierAccountsWithDefaults instantiates a new ShipmentDataRelationshipsCarrierAccounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipmentDataRelationshipsCarrierAccountsWithDefaults() *ShipmentDataRelationshipsCarrierAccounts {
	this := ShipmentDataRelationshipsCarrierAccounts{}
	return &this
}

// GetData returns the Data field value
func (o *ShipmentDataRelationshipsCarrierAccounts) GetData() ShipmentDataRelationshipsCarrierAccountsData {
	if o == nil {
		var ret ShipmentDataRelationshipsCarrierAccountsData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ShipmentDataRelationshipsCarrierAccounts) GetDataOk() (*ShipmentDataRelationshipsCarrierAccountsData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ShipmentDataRelationshipsCarrierAccounts) SetData(v ShipmentDataRelationshipsCarrierAccountsData) {
	o.Data = v
}

func (o ShipmentDataRelationshipsCarrierAccounts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableShipmentDataRelationshipsCarrierAccounts struct {
	value *ShipmentDataRelationshipsCarrierAccounts
	isSet bool
}

func (v NullableShipmentDataRelationshipsCarrierAccounts) Get() *ShipmentDataRelationshipsCarrierAccounts {
	return v.value
}

func (v *NullableShipmentDataRelationshipsCarrierAccounts) Set(val *ShipmentDataRelationshipsCarrierAccounts) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentDataRelationshipsCarrierAccounts) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentDataRelationshipsCarrierAccounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentDataRelationshipsCarrierAccounts(val *ShipmentDataRelationshipsCarrierAccounts) *NullableShipmentDataRelationshipsCarrierAccounts {
	return &NullableShipmentDataRelationshipsCarrierAccounts{value: val, isSet: true}
}

func (v NullableShipmentDataRelationshipsCarrierAccounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipmentDataRelationshipsCarrierAccounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
