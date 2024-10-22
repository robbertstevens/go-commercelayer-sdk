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

// checks if the POSTExternalGateways201ResponseDataRelationshipsExternalPaymentsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTExternalGateways201ResponseDataRelationshipsExternalPaymentsData{}

// POSTExternalGateways201ResponseDataRelationshipsExternalPaymentsData struct for POSTExternalGateways201ResponseDataRelationshipsExternalPaymentsData
type POSTExternalGateways201ResponseDataRelationshipsExternalPaymentsData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewPOSTExternalGateways201ResponseDataRelationshipsExternalPaymentsData instantiates a new POSTExternalGateways201ResponseDataRelationshipsExternalPaymentsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTExternalGateways201ResponseDataRelationshipsExternalPaymentsData() *POSTExternalGateways201ResponseDataRelationshipsExternalPaymentsData {
	this := POSTExternalGateways201ResponseDataRelationshipsExternalPaymentsData{}
	return &this
}

// NewPOSTExternalGateways201ResponseDataRelationshipsExternalPaymentsDataWithDefaults instantiates a new POSTExternalGateways201ResponseDataRelationshipsExternalPaymentsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTExternalGateways201ResponseDataRelationshipsExternalPaymentsDataWithDefaults() *POSTExternalGateways201ResponseDataRelationshipsExternalPaymentsData {
	this := POSTExternalGateways201ResponseDataRelationshipsExternalPaymentsData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTExternalGateways201ResponseDataRelationshipsExternalPaymentsData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTExternalGateways201ResponseDataRelationshipsExternalPaymentsData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *POSTExternalGateways201ResponseDataRelationshipsExternalPaymentsData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *POSTExternalGateways201ResponseDataRelationshipsExternalPaymentsData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTExternalGateways201ResponseDataRelationshipsExternalPaymentsData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTExternalGateways201ResponseDataRelationshipsExternalPaymentsData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *POSTExternalGateways201ResponseDataRelationshipsExternalPaymentsData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *POSTExternalGateways201ResponseDataRelationshipsExternalPaymentsData) SetId(v interface{}) {
	o.Id = v
}

func (o POSTExternalGateways201ResponseDataRelationshipsExternalPaymentsData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTExternalGateways201ResponseDataRelationshipsExternalPaymentsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePOSTExternalGateways201ResponseDataRelationshipsExternalPaymentsData struct {
	value *POSTExternalGateways201ResponseDataRelationshipsExternalPaymentsData
	isSet bool
}

func (v NullablePOSTExternalGateways201ResponseDataRelationshipsExternalPaymentsData) Get() *POSTExternalGateways201ResponseDataRelationshipsExternalPaymentsData {
	return v.value
}

func (v *NullablePOSTExternalGateways201ResponseDataRelationshipsExternalPaymentsData) Set(val *POSTExternalGateways201ResponseDataRelationshipsExternalPaymentsData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTExternalGateways201ResponseDataRelationshipsExternalPaymentsData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTExternalGateways201ResponseDataRelationshipsExternalPaymentsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTExternalGateways201ResponseDataRelationshipsExternalPaymentsData(val *POSTExternalGateways201ResponseDataRelationshipsExternalPaymentsData) *NullablePOSTExternalGateways201ResponseDataRelationshipsExternalPaymentsData {
	return &NullablePOSTExternalGateways201ResponseDataRelationshipsExternalPaymentsData{value: val, isSet: true}
}

func (v NullablePOSTExternalGateways201ResponseDataRelationshipsExternalPaymentsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTExternalGateways201ResponseDataRelationshipsExternalPaymentsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
