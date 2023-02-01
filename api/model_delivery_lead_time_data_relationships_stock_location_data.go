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

// DeliveryLeadTimeDataRelationshipsStockLocationData struct for DeliveryLeadTimeDataRelationshipsStockLocationData
type DeliveryLeadTimeDataRelationshipsStockLocationData struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource's id
	Id *string `json:"id,omitempty"`
}

// NewDeliveryLeadTimeDataRelationshipsStockLocationData instantiates a new DeliveryLeadTimeDataRelationshipsStockLocationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryLeadTimeDataRelationshipsStockLocationData() *DeliveryLeadTimeDataRelationshipsStockLocationData {
	this := DeliveryLeadTimeDataRelationshipsStockLocationData{}
	return &this
}

// NewDeliveryLeadTimeDataRelationshipsStockLocationDataWithDefaults instantiates a new DeliveryLeadTimeDataRelationshipsStockLocationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryLeadTimeDataRelationshipsStockLocationDataWithDefaults() *DeliveryLeadTimeDataRelationshipsStockLocationData {
	this := DeliveryLeadTimeDataRelationshipsStockLocationData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DeliveryLeadTimeDataRelationshipsStockLocationData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryLeadTimeDataRelationshipsStockLocationData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DeliveryLeadTimeDataRelationshipsStockLocationData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DeliveryLeadTimeDataRelationshipsStockLocationData) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeliveryLeadTimeDataRelationshipsStockLocationData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryLeadTimeDataRelationshipsStockLocationData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeliveryLeadTimeDataRelationshipsStockLocationData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DeliveryLeadTimeDataRelationshipsStockLocationData) SetId(v string) {
	o.Id = &v
}

func (o DeliveryLeadTimeDataRelationshipsStockLocationData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableDeliveryLeadTimeDataRelationshipsStockLocationData struct {
	value *DeliveryLeadTimeDataRelationshipsStockLocationData
	isSet bool
}

func (v NullableDeliveryLeadTimeDataRelationshipsStockLocationData) Get() *DeliveryLeadTimeDataRelationshipsStockLocationData {
	return v.value
}

func (v *NullableDeliveryLeadTimeDataRelationshipsStockLocationData) Set(val *DeliveryLeadTimeDataRelationshipsStockLocationData) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryLeadTimeDataRelationshipsStockLocationData) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryLeadTimeDataRelationshipsStockLocationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryLeadTimeDataRelationshipsStockLocationData(val *DeliveryLeadTimeDataRelationshipsStockLocationData) *NullableDeliveryLeadTimeDataRelationshipsStockLocationData {
	return &NullableDeliveryLeadTimeDataRelationshipsStockLocationData{value: val, isSet: true}
}

func (v NullableDeliveryLeadTimeDataRelationshipsStockLocationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryLeadTimeDataRelationshipsStockLocationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
