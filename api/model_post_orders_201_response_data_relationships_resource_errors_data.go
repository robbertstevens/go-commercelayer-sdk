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

// checks if the POSTOrders201ResponseDataRelationshipsResourceErrorsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTOrders201ResponseDataRelationshipsResourceErrorsData{}

// POSTOrders201ResponseDataRelationshipsResourceErrorsData struct for POSTOrders201ResponseDataRelationshipsResourceErrorsData
type POSTOrders201ResponseDataRelationshipsResourceErrorsData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewPOSTOrders201ResponseDataRelationshipsResourceErrorsData instantiates a new POSTOrders201ResponseDataRelationshipsResourceErrorsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTOrders201ResponseDataRelationshipsResourceErrorsData() *POSTOrders201ResponseDataRelationshipsResourceErrorsData {
	this := POSTOrders201ResponseDataRelationshipsResourceErrorsData{}
	return &this
}

// NewPOSTOrders201ResponseDataRelationshipsResourceErrorsDataWithDefaults instantiates a new POSTOrders201ResponseDataRelationshipsResourceErrorsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTOrders201ResponseDataRelationshipsResourceErrorsDataWithDefaults() *POSTOrders201ResponseDataRelationshipsResourceErrorsData {
	this := POSTOrders201ResponseDataRelationshipsResourceErrorsData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrders201ResponseDataRelationshipsResourceErrorsData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrders201ResponseDataRelationshipsResourceErrorsData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *POSTOrders201ResponseDataRelationshipsResourceErrorsData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *POSTOrders201ResponseDataRelationshipsResourceErrorsData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrders201ResponseDataRelationshipsResourceErrorsData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrders201ResponseDataRelationshipsResourceErrorsData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *POSTOrders201ResponseDataRelationshipsResourceErrorsData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *POSTOrders201ResponseDataRelationshipsResourceErrorsData) SetId(v interface{}) {
	o.Id = v
}

func (o POSTOrders201ResponseDataRelationshipsResourceErrorsData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTOrders201ResponseDataRelationshipsResourceErrorsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePOSTOrders201ResponseDataRelationshipsResourceErrorsData struct {
	value *POSTOrders201ResponseDataRelationshipsResourceErrorsData
	isSet bool
}

func (v NullablePOSTOrders201ResponseDataRelationshipsResourceErrorsData) Get() *POSTOrders201ResponseDataRelationshipsResourceErrorsData {
	return v.value
}

func (v *NullablePOSTOrders201ResponseDataRelationshipsResourceErrorsData) Set(val *POSTOrders201ResponseDataRelationshipsResourceErrorsData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTOrders201ResponseDataRelationshipsResourceErrorsData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTOrders201ResponseDataRelationshipsResourceErrorsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTOrders201ResponseDataRelationshipsResourceErrorsData(val *POSTOrders201ResponseDataRelationshipsResourceErrorsData) *NullablePOSTOrders201ResponseDataRelationshipsResourceErrorsData {
	return &NullablePOSTOrders201ResponseDataRelationshipsResourceErrorsData{value: val, isSet: true}
}

func (v NullablePOSTOrders201ResponseDataRelationshipsResourceErrorsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTOrders201ResponseDataRelationshipsResourceErrorsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
