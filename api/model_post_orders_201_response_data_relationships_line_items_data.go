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

// checks if the POSTOrders201ResponseDataRelationshipsLineItemsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTOrders201ResponseDataRelationshipsLineItemsData{}

// POSTOrders201ResponseDataRelationshipsLineItemsData struct for POSTOrders201ResponseDataRelationshipsLineItemsData
type POSTOrders201ResponseDataRelationshipsLineItemsData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewPOSTOrders201ResponseDataRelationshipsLineItemsData instantiates a new POSTOrders201ResponseDataRelationshipsLineItemsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTOrders201ResponseDataRelationshipsLineItemsData() *POSTOrders201ResponseDataRelationshipsLineItemsData {
	this := POSTOrders201ResponseDataRelationshipsLineItemsData{}
	return &this
}

// NewPOSTOrders201ResponseDataRelationshipsLineItemsDataWithDefaults instantiates a new POSTOrders201ResponseDataRelationshipsLineItemsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTOrders201ResponseDataRelationshipsLineItemsDataWithDefaults() *POSTOrders201ResponseDataRelationshipsLineItemsData {
	this := POSTOrders201ResponseDataRelationshipsLineItemsData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrders201ResponseDataRelationshipsLineItemsData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrders201ResponseDataRelationshipsLineItemsData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *POSTOrders201ResponseDataRelationshipsLineItemsData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *POSTOrders201ResponseDataRelationshipsLineItemsData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrders201ResponseDataRelationshipsLineItemsData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrders201ResponseDataRelationshipsLineItemsData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *POSTOrders201ResponseDataRelationshipsLineItemsData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *POSTOrders201ResponseDataRelationshipsLineItemsData) SetId(v interface{}) {
	o.Id = v
}

func (o POSTOrders201ResponseDataRelationshipsLineItemsData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTOrders201ResponseDataRelationshipsLineItemsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePOSTOrders201ResponseDataRelationshipsLineItemsData struct {
	value *POSTOrders201ResponseDataRelationshipsLineItemsData
	isSet bool
}

func (v NullablePOSTOrders201ResponseDataRelationshipsLineItemsData) Get() *POSTOrders201ResponseDataRelationshipsLineItemsData {
	return v.value
}

func (v *NullablePOSTOrders201ResponseDataRelationshipsLineItemsData) Set(val *POSTOrders201ResponseDataRelationshipsLineItemsData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTOrders201ResponseDataRelationshipsLineItemsData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTOrders201ResponseDataRelationshipsLineItemsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTOrders201ResponseDataRelationshipsLineItemsData(val *POSTOrders201ResponseDataRelationshipsLineItemsData) *NullablePOSTOrders201ResponseDataRelationshipsLineItemsData {
	return &NullablePOSTOrders201ResponseDataRelationshipsLineItemsData{value: val, isSet: true}
}

func (v NullablePOSTOrders201ResponseDataRelationshipsLineItemsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTOrders201ResponseDataRelationshipsLineItemsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
