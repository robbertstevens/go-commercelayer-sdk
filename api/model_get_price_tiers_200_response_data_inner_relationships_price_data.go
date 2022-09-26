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

// GETPriceTiers200ResponseDataInnerRelationshipsPriceData struct for GETPriceTiers200ResponseDataInnerRelationshipsPriceData
type GETPriceTiers200ResponseDataInnerRelationshipsPriceData struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource ID
	Id *string `json:"id,omitempty"`
}

// NewGETPriceTiers200ResponseDataInnerRelationshipsPriceData instantiates a new GETPriceTiers200ResponseDataInnerRelationshipsPriceData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETPriceTiers200ResponseDataInnerRelationshipsPriceData() *GETPriceTiers200ResponseDataInnerRelationshipsPriceData {
	this := GETPriceTiers200ResponseDataInnerRelationshipsPriceData{}
	return &this
}

// NewGETPriceTiers200ResponseDataInnerRelationshipsPriceDataWithDefaults instantiates a new GETPriceTiers200ResponseDataInnerRelationshipsPriceData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETPriceTiers200ResponseDataInnerRelationshipsPriceDataWithDefaults() *GETPriceTiers200ResponseDataInnerRelationshipsPriceData {
	this := GETPriceTiers200ResponseDataInnerRelationshipsPriceData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GETPriceTiers200ResponseDataInnerRelationshipsPriceData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETPriceTiers200ResponseDataInnerRelationshipsPriceData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETPriceTiers200ResponseDataInnerRelationshipsPriceData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GETPriceTiers200ResponseDataInnerRelationshipsPriceData) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETPriceTiers200ResponseDataInnerRelationshipsPriceData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETPriceTiers200ResponseDataInnerRelationshipsPriceData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETPriceTiers200ResponseDataInnerRelationshipsPriceData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETPriceTiers200ResponseDataInnerRelationshipsPriceData) SetId(v string) {
	o.Id = &v
}

func (o GETPriceTiers200ResponseDataInnerRelationshipsPriceData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETPriceTiers200ResponseDataInnerRelationshipsPriceData struct {
	value *GETPriceTiers200ResponseDataInnerRelationshipsPriceData
	isSet bool
}

func (v NullableGETPriceTiers200ResponseDataInnerRelationshipsPriceData) Get() *GETPriceTiers200ResponseDataInnerRelationshipsPriceData {
	return v.value
}

func (v *NullableGETPriceTiers200ResponseDataInnerRelationshipsPriceData) Set(val *GETPriceTiers200ResponseDataInnerRelationshipsPriceData) {
	v.value = val
	v.isSet = true
}

func (v NullableGETPriceTiers200ResponseDataInnerRelationshipsPriceData) IsSet() bool {
	return v.isSet
}

func (v *NullableGETPriceTiers200ResponseDataInnerRelationshipsPriceData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETPriceTiers200ResponseDataInnerRelationshipsPriceData(val *GETPriceTiers200ResponseDataInnerRelationshipsPriceData) *NullableGETPriceTiers200ResponseDataInnerRelationshipsPriceData {
	return &NullableGETPriceTiers200ResponseDataInnerRelationshipsPriceData{value: val, isSet: true}
}

func (v NullableGETPriceTiers200ResponseDataInnerRelationshipsPriceData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETPriceTiers200ResponseDataInnerRelationshipsPriceData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
