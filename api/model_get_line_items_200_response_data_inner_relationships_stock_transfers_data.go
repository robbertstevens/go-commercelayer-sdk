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

// GETLineItems200ResponseDataInnerRelationshipsStockTransfersData struct for GETLineItems200ResponseDataInnerRelationshipsStockTransfersData
type GETLineItems200ResponseDataInnerRelationshipsStockTransfersData struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource ID
	Id *string `json:"id,omitempty"`
}

// NewGETLineItems200ResponseDataInnerRelationshipsStockTransfersData instantiates a new GETLineItems200ResponseDataInnerRelationshipsStockTransfersData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETLineItems200ResponseDataInnerRelationshipsStockTransfersData() *GETLineItems200ResponseDataInnerRelationshipsStockTransfersData {
	this := GETLineItems200ResponseDataInnerRelationshipsStockTransfersData{}
	return &this
}

// NewGETLineItems200ResponseDataInnerRelationshipsStockTransfersDataWithDefaults instantiates a new GETLineItems200ResponseDataInnerRelationshipsStockTransfersData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETLineItems200ResponseDataInnerRelationshipsStockTransfersDataWithDefaults() *GETLineItems200ResponseDataInnerRelationshipsStockTransfersData {
	this := GETLineItems200ResponseDataInnerRelationshipsStockTransfersData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GETLineItems200ResponseDataInnerRelationshipsStockTransfersData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETLineItems200ResponseDataInnerRelationshipsStockTransfersData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETLineItems200ResponseDataInnerRelationshipsStockTransfersData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GETLineItems200ResponseDataInnerRelationshipsStockTransfersData) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETLineItems200ResponseDataInnerRelationshipsStockTransfersData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETLineItems200ResponseDataInnerRelationshipsStockTransfersData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETLineItems200ResponseDataInnerRelationshipsStockTransfersData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETLineItems200ResponseDataInnerRelationshipsStockTransfersData) SetId(v string) {
	o.Id = &v
}

func (o GETLineItems200ResponseDataInnerRelationshipsStockTransfersData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETLineItems200ResponseDataInnerRelationshipsStockTransfersData struct {
	value *GETLineItems200ResponseDataInnerRelationshipsStockTransfersData
	isSet bool
}

func (v NullableGETLineItems200ResponseDataInnerRelationshipsStockTransfersData) Get() *GETLineItems200ResponseDataInnerRelationshipsStockTransfersData {
	return v.value
}

func (v *NullableGETLineItems200ResponseDataInnerRelationshipsStockTransfersData) Set(val *GETLineItems200ResponseDataInnerRelationshipsStockTransfersData) {
	v.value = val
	v.isSet = true
}

func (v NullableGETLineItems200ResponseDataInnerRelationshipsStockTransfersData) IsSet() bool {
	return v.isSet
}

func (v *NullableGETLineItems200ResponseDataInnerRelationshipsStockTransfersData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETLineItems200ResponseDataInnerRelationshipsStockTransfersData(val *GETLineItems200ResponseDataInnerRelationshipsStockTransfersData) *NullableGETLineItems200ResponseDataInnerRelationshipsStockTransfersData {
	return &NullableGETLineItems200ResponseDataInnerRelationshipsStockTransfersData{value: val, isSet: true}
}

func (v NullableGETLineItems200ResponseDataInnerRelationshipsStockTransfersData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETLineItems200ResponseDataInnerRelationshipsStockTransfersData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
