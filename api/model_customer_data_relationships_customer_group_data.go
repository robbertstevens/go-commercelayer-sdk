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

// checks if the CustomerDataRelationshipsCustomerGroupData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerDataRelationshipsCustomerGroupData{}

// CustomerDataRelationshipsCustomerGroupData struct for CustomerDataRelationshipsCustomerGroupData
type CustomerDataRelationshipsCustomerGroupData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource's id
	Id interface{} `json:"id,omitempty"`
}

// NewCustomerDataRelationshipsCustomerGroupData instantiates a new CustomerDataRelationshipsCustomerGroupData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerDataRelationshipsCustomerGroupData() *CustomerDataRelationshipsCustomerGroupData {
	this := CustomerDataRelationshipsCustomerGroupData{}
	return &this
}

// NewCustomerDataRelationshipsCustomerGroupDataWithDefaults instantiates a new CustomerDataRelationshipsCustomerGroupData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerDataRelationshipsCustomerGroupDataWithDefaults() *CustomerDataRelationshipsCustomerGroupData {
	this := CustomerDataRelationshipsCustomerGroupData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerDataRelationshipsCustomerGroupData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerDataRelationshipsCustomerGroupData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CustomerDataRelationshipsCustomerGroupData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *CustomerDataRelationshipsCustomerGroupData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerDataRelationshipsCustomerGroupData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerDataRelationshipsCustomerGroupData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustomerDataRelationshipsCustomerGroupData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *CustomerDataRelationshipsCustomerGroupData) SetId(v interface{}) {
	o.Id = v
}

func (o CustomerDataRelationshipsCustomerGroupData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerDataRelationshipsCustomerGroupData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableCustomerDataRelationshipsCustomerGroupData struct {
	value *CustomerDataRelationshipsCustomerGroupData
	isSet bool
}

func (v NullableCustomerDataRelationshipsCustomerGroupData) Get() *CustomerDataRelationshipsCustomerGroupData {
	return v.value
}

func (v *NullableCustomerDataRelationshipsCustomerGroupData) Set(val *CustomerDataRelationshipsCustomerGroupData) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerDataRelationshipsCustomerGroupData) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerDataRelationshipsCustomerGroupData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerDataRelationshipsCustomerGroupData(val *CustomerDataRelationshipsCustomerGroupData) *NullableCustomerDataRelationshipsCustomerGroupData {
	return &NullableCustomerDataRelationshipsCustomerGroupData{value: val, isSet: true}
}

func (v NullableCustomerDataRelationshipsCustomerGroupData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerDataRelationshipsCustomerGroupData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
