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

// GETCaptures200ResponseDataInnerRelationshipsRefundsData struct for GETCaptures200ResponseDataInnerRelationshipsRefundsData
type GETCaptures200ResponseDataInnerRelationshipsRefundsData struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource ID
	Id *string `json:"id,omitempty"`
}

// NewGETCaptures200ResponseDataInnerRelationshipsRefundsData instantiates a new GETCaptures200ResponseDataInnerRelationshipsRefundsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCaptures200ResponseDataInnerRelationshipsRefundsData() *GETCaptures200ResponseDataInnerRelationshipsRefundsData {
	this := GETCaptures200ResponseDataInnerRelationshipsRefundsData{}
	return &this
}

// NewGETCaptures200ResponseDataInnerRelationshipsRefundsDataWithDefaults instantiates a new GETCaptures200ResponseDataInnerRelationshipsRefundsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCaptures200ResponseDataInnerRelationshipsRefundsDataWithDefaults() *GETCaptures200ResponseDataInnerRelationshipsRefundsData {
	this := GETCaptures200ResponseDataInnerRelationshipsRefundsData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GETCaptures200ResponseDataInnerRelationshipsRefundsData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCaptures200ResponseDataInnerRelationshipsRefundsData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETCaptures200ResponseDataInnerRelationshipsRefundsData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GETCaptures200ResponseDataInnerRelationshipsRefundsData) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETCaptures200ResponseDataInnerRelationshipsRefundsData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCaptures200ResponseDataInnerRelationshipsRefundsData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETCaptures200ResponseDataInnerRelationshipsRefundsData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETCaptures200ResponseDataInnerRelationshipsRefundsData) SetId(v string) {
	o.Id = &v
}

func (o GETCaptures200ResponseDataInnerRelationshipsRefundsData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETCaptures200ResponseDataInnerRelationshipsRefundsData struct {
	value *GETCaptures200ResponseDataInnerRelationshipsRefundsData
	isSet bool
}

func (v NullableGETCaptures200ResponseDataInnerRelationshipsRefundsData) Get() *GETCaptures200ResponseDataInnerRelationshipsRefundsData {
	return v.value
}

func (v *NullableGETCaptures200ResponseDataInnerRelationshipsRefundsData) Set(val *GETCaptures200ResponseDataInnerRelationshipsRefundsData) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCaptures200ResponseDataInnerRelationshipsRefundsData) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCaptures200ResponseDataInnerRelationshipsRefundsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCaptures200ResponseDataInnerRelationshipsRefundsData(val *GETCaptures200ResponseDataInnerRelationshipsRefundsData) *NullableGETCaptures200ResponseDataInnerRelationshipsRefundsData {
	return &NullableGETCaptures200ResponseDataInnerRelationshipsRefundsData{value: val, isSet: true}
}

func (v NullableGETCaptures200ResponseDataInnerRelationshipsRefundsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCaptures200ResponseDataInnerRelationshipsRefundsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
