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

// checks if the POSTSkus201ResponseDataRelationshipsStockItemsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTSkus201ResponseDataRelationshipsStockItemsData{}

// POSTSkus201ResponseDataRelationshipsStockItemsData struct for POSTSkus201ResponseDataRelationshipsStockItemsData
type POSTSkus201ResponseDataRelationshipsStockItemsData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewPOSTSkus201ResponseDataRelationshipsStockItemsData instantiates a new POSTSkus201ResponseDataRelationshipsStockItemsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTSkus201ResponseDataRelationshipsStockItemsData() *POSTSkus201ResponseDataRelationshipsStockItemsData {
	this := POSTSkus201ResponseDataRelationshipsStockItemsData{}
	return &this
}

// NewPOSTSkus201ResponseDataRelationshipsStockItemsDataWithDefaults instantiates a new POSTSkus201ResponseDataRelationshipsStockItemsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTSkus201ResponseDataRelationshipsStockItemsDataWithDefaults() *POSTSkus201ResponseDataRelationshipsStockItemsData {
	this := POSTSkus201ResponseDataRelationshipsStockItemsData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTSkus201ResponseDataRelationshipsStockItemsData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTSkus201ResponseDataRelationshipsStockItemsData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *POSTSkus201ResponseDataRelationshipsStockItemsData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *POSTSkus201ResponseDataRelationshipsStockItemsData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTSkus201ResponseDataRelationshipsStockItemsData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTSkus201ResponseDataRelationshipsStockItemsData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *POSTSkus201ResponseDataRelationshipsStockItemsData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *POSTSkus201ResponseDataRelationshipsStockItemsData) SetId(v interface{}) {
	o.Id = v
}

func (o POSTSkus201ResponseDataRelationshipsStockItemsData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTSkus201ResponseDataRelationshipsStockItemsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePOSTSkus201ResponseDataRelationshipsStockItemsData struct {
	value *POSTSkus201ResponseDataRelationshipsStockItemsData
	isSet bool
}

func (v NullablePOSTSkus201ResponseDataRelationshipsStockItemsData) Get() *POSTSkus201ResponseDataRelationshipsStockItemsData {
	return v.value
}

func (v *NullablePOSTSkus201ResponseDataRelationshipsStockItemsData) Set(val *POSTSkus201ResponseDataRelationshipsStockItemsData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTSkus201ResponseDataRelationshipsStockItemsData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTSkus201ResponseDataRelationshipsStockItemsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTSkus201ResponseDataRelationshipsStockItemsData(val *POSTSkus201ResponseDataRelationshipsStockItemsData) *NullablePOSTSkus201ResponseDataRelationshipsStockItemsData {
	return &NullablePOSTSkus201ResponseDataRelationshipsStockItemsData{value: val, isSet: true}
}

func (v NullablePOSTSkus201ResponseDataRelationshipsStockItemsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTSkus201ResponseDataRelationshipsStockItemsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
