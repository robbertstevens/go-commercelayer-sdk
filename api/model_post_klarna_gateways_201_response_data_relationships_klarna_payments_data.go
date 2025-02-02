/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.3.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPaymentsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPaymentsData{}

// POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPaymentsData struct for POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPaymentsData
type POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPaymentsData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewPOSTKlarnaGateways201ResponseDataRelationshipsKlarnaPaymentsData instantiates a new POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPaymentsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTKlarnaGateways201ResponseDataRelationshipsKlarnaPaymentsData() *POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPaymentsData {
	this := POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPaymentsData{}
	return &this
}

// NewPOSTKlarnaGateways201ResponseDataRelationshipsKlarnaPaymentsDataWithDefaults instantiates a new POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPaymentsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTKlarnaGateways201ResponseDataRelationshipsKlarnaPaymentsDataWithDefaults() *POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPaymentsData {
	this := POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPaymentsData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPaymentsData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPaymentsData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPaymentsData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPaymentsData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPaymentsData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPaymentsData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPaymentsData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPaymentsData) SetId(v interface{}) {
	o.Id = v
}

func (o POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPaymentsData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPaymentsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePOSTKlarnaGateways201ResponseDataRelationshipsKlarnaPaymentsData struct {
	value *POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPaymentsData
	isSet bool
}

func (v NullablePOSTKlarnaGateways201ResponseDataRelationshipsKlarnaPaymentsData) Get() *POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPaymentsData {
	return v.value
}

func (v *NullablePOSTKlarnaGateways201ResponseDataRelationshipsKlarnaPaymentsData) Set(val *POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPaymentsData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTKlarnaGateways201ResponseDataRelationshipsKlarnaPaymentsData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTKlarnaGateways201ResponseDataRelationshipsKlarnaPaymentsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTKlarnaGateways201ResponseDataRelationshipsKlarnaPaymentsData(val *POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPaymentsData) *NullablePOSTKlarnaGateways201ResponseDataRelationshipsKlarnaPaymentsData {
	return &NullablePOSTKlarnaGateways201ResponseDataRelationshipsKlarnaPaymentsData{value: val, isSet: true}
}

func (v NullablePOSTKlarnaGateways201ResponseDataRelationshipsKlarnaPaymentsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTKlarnaGateways201ResponseDataRelationshipsKlarnaPaymentsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
