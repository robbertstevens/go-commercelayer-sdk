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

// checks if the POSTPrices201ResponseDataRelationshipsJwtCustomerData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTPrices201ResponseDataRelationshipsJwtCustomerData{}

// POSTPrices201ResponseDataRelationshipsJwtCustomerData struct for POSTPrices201ResponseDataRelationshipsJwtCustomerData
type POSTPrices201ResponseDataRelationshipsJwtCustomerData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewPOSTPrices201ResponseDataRelationshipsJwtCustomerData instantiates a new POSTPrices201ResponseDataRelationshipsJwtCustomerData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTPrices201ResponseDataRelationshipsJwtCustomerData() *POSTPrices201ResponseDataRelationshipsJwtCustomerData {
	this := POSTPrices201ResponseDataRelationshipsJwtCustomerData{}
	return &this
}

// NewPOSTPrices201ResponseDataRelationshipsJwtCustomerDataWithDefaults instantiates a new POSTPrices201ResponseDataRelationshipsJwtCustomerData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTPrices201ResponseDataRelationshipsJwtCustomerDataWithDefaults() *POSTPrices201ResponseDataRelationshipsJwtCustomerData {
	this := POSTPrices201ResponseDataRelationshipsJwtCustomerData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTPrices201ResponseDataRelationshipsJwtCustomerData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPrices201ResponseDataRelationshipsJwtCustomerData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *POSTPrices201ResponseDataRelationshipsJwtCustomerData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *POSTPrices201ResponseDataRelationshipsJwtCustomerData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTPrices201ResponseDataRelationshipsJwtCustomerData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPrices201ResponseDataRelationshipsJwtCustomerData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *POSTPrices201ResponseDataRelationshipsJwtCustomerData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *POSTPrices201ResponseDataRelationshipsJwtCustomerData) SetId(v interface{}) {
	o.Id = v
}

func (o POSTPrices201ResponseDataRelationshipsJwtCustomerData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTPrices201ResponseDataRelationshipsJwtCustomerData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePOSTPrices201ResponseDataRelationshipsJwtCustomerData struct {
	value *POSTPrices201ResponseDataRelationshipsJwtCustomerData
	isSet bool
}

func (v NullablePOSTPrices201ResponseDataRelationshipsJwtCustomerData) Get() *POSTPrices201ResponseDataRelationshipsJwtCustomerData {
	return v.value
}

func (v *NullablePOSTPrices201ResponseDataRelationshipsJwtCustomerData) Set(val *POSTPrices201ResponseDataRelationshipsJwtCustomerData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTPrices201ResponseDataRelationshipsJwtCustomerData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTPrices201ResponseDataRelationshipsJwtCustomerData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTPrices201ResponseDataRelationshipsJwtCustomerData(val *POSTPrices201ResponseDataRelationshipsJwtCustomerData) *NullablePOSTPrices201ResponseDataRelationshipsJwtCustomerData {
	return &NullablePOSTPrices201ResponseDataRelationshipsJwtCustomerData{value: val, isSet: true}
}

func (v NullablePOSTPrices201ResponseDataRelationshipsJwtCustomerData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTPrices201ResponseDataRelationshipsJwtCustomerData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
