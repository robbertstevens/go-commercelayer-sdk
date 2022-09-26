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

// GETCustomerAddresses200ResponseDataInnerRelationshipsAddressData struct for GETCustomerAddresses200ResponseDataInnerRelationshipsAddressData
type GETCustomerAddresses200ResponseDataInnerRelationshipsAddressData struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource ID
	Id *string `json:"id,omitempty"`
}

// NewGETCustomerAddresses200ResponseDataInnerRelationshipsAddressData instantiates a new GETCustomerAddresses200ResponseDataInnerRelationshipsAddressData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCustomerAddresses200ResponseDataInnerRelationshipsAddressData() *GETCustomerAddresses200ResponseDataInnerRelationshipsAddressData {
	this := GETCustomerAddresses200ResponseDataInnerRelationshipsAddressData{}
	return &this
}

// NewGETCustomerAddresses200ResponseDataInnerRelationshipsAddressDataWithDefaults instantiates a new GETCustomerAddresses200ResponseDataInnerRelationshipsAddressData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCustomerAddresses200ResponseDataInnerRelationshipsAddressDataWithDefaults() *GETCustomerAddresses200ResponseDataInnerRelationshipsAddressData {
	this := GETCustomerAddresses200ResponseDataInnerRelationshipsAddressData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GETCustomerAddresses200ResponseDataInnerRelationshipsAddressData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCustomerAddresses200ResponseDataInnerRelationshipsAddressData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETCustomerAddresses200ResponseDataInnerRelationshipsAddressData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GETCustomerAddresses200ResponseDataInnerRelationshipsAddressData) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETCustomerAddresses200ResponseDataInnerRelationshipsAddressData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCustomerAddresses200ResponseDataInnerRelationshipsAddressData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETCustomerAddresses200ResponseDataInnerRelationshipsAddressData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETCustomerAddresses200ResponseDataInnerRelationshipsAddressData) SetId(v string) {
	o.Id = &v
}

func (o GETCustomerAddresses200ResponseDataInnerRelationshipsAddressData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETCustomerAddresses200ResponseDataInnerRelationshipsAddressData struct {
	value *GETCustomerAddresses200ResponseDataInnerRelationshipsAddressData
	isSet bool
}

func (v NullableGETCustomerAddresses200ResponseDataInnerRelationshipsAddressData) Get() *GETCustomerAddresses200ResponseDataInnerRelationshipsAddressData {
	return v.value
}

func (v *NullableGETCustomerAddresses200ResponseDataInnerRelationshipsAddressData) Set(val *GETCustomerAddresses200ResponseDataInnerRelationshipsAddressData) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCustomerAddresses200ResponseDataInnerRelationshipsAddressData) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCustomerAddresses200ResponseDataInnerRelationshipsAddressData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCustomerAddresses200ResponseDataInnerRelationshipsAddressData(val *GETCustomerAddresses200ResponseDataInnerRelationshipsAddressData) *NullableGETCustomerAddresses200ResponseDataInnerRelationshipsAddressData {
	return &NullableGETCustomerAddresses200ResponseDataInnerRelationshipsAddressData{value: val, isSet: true}
}

func (v NullableGETCustomerAddresses200ResponseDataInnerRelationshipsAddressData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCustomerAddresses200ResponseDataInnerRelationshipsAddressData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
