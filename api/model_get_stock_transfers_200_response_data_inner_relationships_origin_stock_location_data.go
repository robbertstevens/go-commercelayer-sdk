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

// GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocationData struct for GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocationData
type GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocationData struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource ID
	Id *string `json:"id,omitempty"`
}

// NewGETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocationData instantiates a new GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocationData() *GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocationData {
	this := GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocationData{}
	return &this
}

// NewGETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocationDataWithDefaults instantiates a new GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocationDataWithDefaults() *GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocationData {
	this := GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocationData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocationData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocationData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocationData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocationData) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocationData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocationData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocationData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocationData) SetId(v string) {
	o.Id = &v
}

func (o GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocationData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocationData struct {
	value *GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocationData
	isSet bool
}

func (v NullableGETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocationData) Get() *GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocationData {
	return v.value
}

func (v *NullableGETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocationData) Set(val *GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocationData) {
	v.value = val
	v.isSet = true
}

func (v NullableGETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocationData) IsSet() bool {
	return v.isSet
}

func (v *NullableGETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocationData(val *GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocationData) *NullableGETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocationData {
	return &NullableGETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocationData{value: val, isSet: true}
}

func (v NullableGETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
