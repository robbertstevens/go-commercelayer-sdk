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

// GETEvents200ResponseDataInnerRelationshipsLastEventCallbacksData struct for GETEvents200ResponseDataInnerRelationshipsLastEventCallbacksData
type GETEvents200ResponseDataInnerRelationshipsLastEventCallbacksData struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource ID
	Id *string `json:"id,omitempty"`
}

// NewGETEvents200ResponseDataInnerRelationshipsLastEventCallbacksData instantiates a new GETEvents200ResponseDataInnerRelationshipsLastEventCallbacksData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETEvents200ResponseDataInnerRelationshipsLastEventCallbacksData() *GETEvents200ResponseDataInnerRelationshipsLastEventCallbacksData {
	this := GETEvents200ResponseDataInnerRelationshipsLastEventCallbacksData{}
	return &this
}

// NewGETEvents200ResponseDataInnerRelationshipsLastEventCallbacksDataWithDefaults instantiates a new GETEvents200ResponseDataInnerRelationshipsLastEventCallbacksData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETEvents200ResponseDataInnerRelationshipsLastEventCallbacksDataWithDefaults() *GETEvents200ResponseDataInnerRelationshipsLastEventCallbacksData {
	this := GETEvents200ResponseDataInnerRelationshipsLastEventCallbacksData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GETEvents200ResponseDataInnerRelationshipsLastEventCallbacksData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETEvents200ResponseDataInnerRelationshipsLastEventCallbacksData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETEvents200ResponseDataInnerRelationshipsLastEventCallbacksData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GETEvents200ResponseDataInnerRelationshipsLastEventCallbacksData) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETEvents200ResponseDataInnerRelationshipsLastEventCallbacksData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETEvents200ResponseDataInnerRelationshipsLastEventCallbacksData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETEvents200ResponseDataInnerRelationshipsLastEventCallbacksData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETEvents200ResponseDataInnerRelationshipsLastEventCallbacksData) SetId(v string) {
	o.Id = &v
}

func (o GETEvents200ResponseDataInnerRelationshipsLastEventCallbacksData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETEvents200ResponseDataInnerRelationshipsLastEventCallbacksData struct {
	value *GETEvents200ResponseDataInnerRelationshipsLastEventCallbacksData
	isSet bool
}

func (v NullableGETEvents200ResponseDataInnerRelationshipsLastEventCallbacksData) Get() *GETEvents200ResponseDataInnerRelationshipsLastEventCallbacksData {
	return v.value
}

func (v *NullableGETEvents200ResponseDataInnerRelationshipsLastEventCallbacksData) Set(val *GETEvents200ResponseDataInnerRelationshipsLastEventCallbacksData) {
	v.value = val
	v.isSet = true
}

func (v NullableGETEvents200ResponseDataInnerRelationshipsLastEventCallbacksData) IsSet() bool {
	return v.isSet
}

func (v *NullableGETEvents200ResponseDataInnerRelationshipsLastEventCallbacksData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETEvents200ResponseDataInnerRelationshipsLastEventCallbacksData(val *GETEvents200ResponseDataInnerRelationshipsLastEventCallbacksData) *NullableGETEvents200ResponseDataInnerRelationshipsLastEventCallbacksData {
	return &NullableGETEvents200ResponseDataInnerRelationshipsLastEventCallbacksData{value: val, isSet: true}
}

func (v NullableGETEvents200ResponseDataInnerRelationshipsLastEventCallbacksData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETEvents200ResponseDataInnerRelationshipsLastEventCallbacksData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
