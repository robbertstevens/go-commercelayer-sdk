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

// GETSkuLists200ResponseDataInnerRelationshipsSkuListItemsData struct for GETSkuLists200ResponseDataInnerRelationshipsSkuListItemsData
type GETSkuLists200ResponseDataInnerRelationshipsSkuListItemsData struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource ID
	Id *string `json:"id,omitempty"`
}

// NewGETSkuLists200ResponseDataInnerRelationshipsSkuListItemsData instantiates a new GETSkuLists200ResponseDataInnerRelationshipsSkuListItemsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETSkuLists200ResponseDataInnerRelationshipsSkuListItemsData() *GETSkuLists200ResponseDataInnerRelationshipsSkuListItemsData {
	this := GETSkuLists200ResponseDataInnerRelationshipsSkuListItemsData{}
	return &this
}

// NewGETSkuLists200ResponseDataInnerRelationshipsSkuListItemsDataWithDefaults instantiates a new GETSkuLists200ResponseDataInnerRelationshipsSkuListItemsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETSkuLists200ResponseDataInnerRelationshipsSkuListItemsDataWithDefaults() *GETSkuLists200ResponseDataInnerRelationshipsSkuListItemsData {
	this := GETSkuLists200ResponseDataInnerRelationshipsSkuListItemsData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GETSkuLists200ResponseDataInnerRelationshipsSkuListItemsData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkuLists200ResponseDataInnerRelationshipsSkuListItemsData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETSkuLists200ResponseDataInnerRelationshipsSkuListItemsData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GETSkuLists200ResponseDataInnerRelationshipsSkuListItemsData) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETSkuLists200ResponseDataInnerRelationshipsSkuListItemsData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkuLists200ResponseDataInnerRelationshipsSkuListItemsData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETSkuLists200ResponseDataInnerRelationshipsSkuListItemsData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETSkuLists200ResponseDataInnerRelationshipsSkuListItemsData) SetId(v string) {
	o.Id = &v
}

func (o GETSkuLists200ResponseDataInnerRelationshipsSkuListItemsData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETSkuLists200ResponseDataInnerRelationshipsSkuListItemsData struct {
	value *GETSkuLists200ResponseDataInnerRelationshipsSkuListItemsData
	isSet bool
}

func (v NullableGETSkuLists200ResponseDataInnerRelationshipsSkuListItemsData) Get() *GETSkuLists200ResponseDataInnerRelationshipsSkuListItemsData {
	return v.value
}

func (v *NullableGETSkuLists200ResponseDataInnerRelationshipsSkuListItemsData) Set(val *GETSkuLists200ResponseDataInnerRelationshipsSkuListItemsData) {
	v.value = val
	v.isSet = true
}

func (v NullableGETSkuLists200ResponseDataInnerRelationshipsSkuListItemsData) IsSet() bool {
	return v.isSet
}

func (v *NullableGETSkuLists200ResponseDataInnerRelationshipsSkuListItemsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETSkuLists200ResponseDataInnerRelationshipsSkuListItemsData(val *GETSkuLists200ResponseDataInnerRelationshipsSkuListItemsData) *NullableGETSkuLists200ResponseDataInnerRelationshipsSkuListItemsData {
	return &NullableGETSkuLists200ResponseDataInnerRelationshipsSkuListItemsData{value: val, isSet: true}
}

func (v NullableGETSkuLists200ResponseDataInnerRelationshipsSkuListItemsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETSkuLists200ResponseDataInnerRelationshipsSkuListItemsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
