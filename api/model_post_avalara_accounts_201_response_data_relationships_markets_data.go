/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.3.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the POSTAvalaraAccounts201ResponseDataRelationshipsMarketsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTAvalaraAccounts201ResponseDataRelationshipsMarketsData{}

// POSTAvalaraAccounts201ResponseDataRelationshipsMarketsData struct for POSTAvalaraAccounts201ResponseDataRelationshipsMarketsData
type POSTAvalaraAccounts201ResponseDataRelationshipsMarketsData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewPOSTAvalaraAccounts201ResponseDataRelationshipsMarketsData instantiates a new POSTAvalaraAccounts201ResponseDataRelationshipsMarketsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTAvalaraAccounts201ResponseDataRelationshipsMarketsData() *POSTAvalaraAccounts201ResponseDataRelationshipsMarketsData {
	this := POSTAvalaraAccounts201ResponseDataRelationshipsMarketsData{}
	return &this
}

// NewPOSTAvalaraAccounts201ResponseDataRelationshipsMarketsDataWithDefaults instantiates a new POSTAvalaraAccounts201ResponseDataRelationshipsMarketsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTAvalaraAccounts201ResponseDataRelationshipsMarketsDataWithDefaults() *POSTAvalaraAccounts201ResponseDataRelationshipsMarketsData {
	this := POSTAvalaraAccounts201ResponseDataRelationshipsMarketsData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTAvalaraAccounts201ResponseDataRelationshipsMarketsData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTAvalaraAccounts201ResponseDataRelationshipsMarketsData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *POSTAvalaraAccounts201ResponseDataRelationshipsMarketsData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *POSTAvalaraAccounts201ResponseDataRelationshipsMarketsData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTAvalaraAccounts201ResponseDataRelationshipsMarketsData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTAvalaraAccounts201ResponseDataRelationshipsMarketsData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *POSTAvalaraAccounts201ResponseDataRelationshipsMarketsData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *POSTAvalaraAccounts201ResponseDataRelationshipsMarketsData) SetId(v interface{}) {
	o.Id = v
}

func (o POSTAvalaraAccounts201ResponseDataRelationshipsMarketsData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTAvalaraAccounts201ResponseDataRelationshipsMarketsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePOSTAvalaraAccounts201ResponseDataRelationshipsMarketsData struct {
	value *POSTAvalaraAccounts201ResponseDataRelationshipsMarketsData
	isSet bool
}

func (v NullablePOSTAvalaraAccounts201ResponseDataRelationshipsMarketsData) Get() *POSTAvalaraAccounts201ResponseDataRelationshipsMarketsData {
	return v.value
}

func (v *NullablePOSTAvalaraAccounts201ResponseDataRelationshipsMarketsData) Set(val *POSTAvalaraAccounts201ResponseDataRelationshipsMarketsData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTAvalaraAccounts201ResponseDataRelationshipsMarketsData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTAvalaraAccounts201ResponseDataRelationshipsMarketsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTAvalaraAccounts201ResponseDataRelationshipsMarketsData(val *POSTAvalaraAccounts201ResponseDataRelationshipsMarketsData) *NullablePOSTAvalaraAccounts201ResponseDataRelationshipsMarketsData {
	return &NullablePOSTAvalaraAccounts201ResponseDataRelationshipsMarketsData{value: val, isSet: true}
}

func (v NullablePOSTAvalaraAccounts201ResponseDataRelationshipsMarketsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTAvalaraAccounts201ResponseDataRelationshipsMarketsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
