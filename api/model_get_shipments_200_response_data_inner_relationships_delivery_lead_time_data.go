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

// GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTimeData struct for GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTimeData
type GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTimeData struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource ID
	Id *string `json:"id,omitempty"`
}

// NewGETShipments200ResponseDataInnerRelationshipsDeliveryLeadTimeData instantiates a new GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTimeData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETShipments200ResponseDataInnerRelationshipsDeliveryLeadTimeData() *GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTimeData {
	this := GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTimeData{}
	return &this
}

// NewGETShipments200ResponseDataInnerRelationshipsDeliveryLeadTimeDataWithDefaults instantiates a new GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTimeData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETShipments200ResponseDataInnerRelationshipsDeliveryLeadTimeDataWithDefaults() *GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTimeData {
	this := GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTimeData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTimeData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTimeData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTimeData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTimeData) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTimeData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTimeData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTimeData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTimeData) SetId(v string) {
	o.Id = &v
}

func (o GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTimeData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETShipments200ResponseDataInnerRelationshipsDeliveryLeadTimeData struct {
	value *GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTimeData
	isSet bool
}

func (v NullableGETShipments200ResponseDataInnerRelationshipsDeliveryLeadTimeData) Get() *GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTimeData {
	return v.value
}

func (v *NullableGETShipments200ResponseDataInnerRelationshipsDeliveryLeadTimeData) Set(val *GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTimeData) {
	v.value = val
	v.isSet = true
}

func (v NullableGETShipments200ResponseDataInnerRelationshipsDeliveryLeadTimeData) IsSet() bool {
	return v.isSet
}

func (v *NullableGETShipments200ResponseDataInnerRelationshipsDeliveryLeadTimeData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETShipments200ResponseDataInnerRelationshipsDeliveryLeadTimeData(val *GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTimeData) *NullableGETShipments200ResponseDataInnerRelationshipsDeliveryLeadTimeData {
	return &NullableGETShipments200ResponseDataInnerRelationshipsDeliveryLeadTimeData{value: val, isSet: true}
}

func (v NullableGETShipments200ResponseDataInnerRelationshipsDeliveryLeadTimeData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETShipments200ResponseDataInnerRelationshipsDeliveryLeadTimeData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
