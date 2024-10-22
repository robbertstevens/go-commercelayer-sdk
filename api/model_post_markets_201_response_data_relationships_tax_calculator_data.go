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

// checks if the POSTMarkets201ResponseDataRelationshipsTaxCalculatorData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTMarkets201ResponseDataRelationshipsTaxCalculatorData{}

// POSTMarkets201ResponseDataRelationshipsTaxCalculatorData struct for POSTMarkets201ResponseDataRelationshipsTaxCalculatorData
type POSTMarkets201ResponseDataRelationshipsTaxCalculatorData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewPOSTMarkets201ResponseDataRelationshipsTaxCalculatorData instantiates a new POSTMarkets201ResponseDataRelationshipsTaxCalculatorData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTMarkets201ResponseDataRelationshipsTaxCalculatorData() *POSTMarkets201ResponseDataRelationshipsTaxCalculatorData {
	this := POSTMarkets201ResponseDataRelationshipsTaxCalculatorData{}
	return &this
}

// NewPOSTMarkets201ResponseDataRelationshipsTaxCalculatorDataWithDefaults instantiates a new POSTMarkets201ResponseDataRelationshipsTaxCalculatorData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTMarkets201ResponseDataRelationshipsTaxCalculatorDataWithDefaults() *POSTMarkets201ResponseDataRelationshipsTaxCalculatorData {
	this := POSTMarkets201ResponseDataRelationshipsTaxCalculatorData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTMarkets201ResponseDataRelationshipsTaxCalculatorData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTMarkets201ResponseDataRelationshipsTaxCalculatorData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *POSTMarkets201ResponseDataRelationshipsTaxCalculatorData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *POSTMarkets201ResponseDataRelationshipsTaxCalculatorData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTMarkets201ResponseDataRelationshipsTaxCalculatorData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTMarkets201ResponseDataRelationshipsTaxCalculatorData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *POSTMarkets201ResponseDataRelationshipsTaxCalculatorData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *POSTMarkets201ResponseDataRelationshipsTaxCalculatorData) SetId(v interface{}) {
	o.Id = v
}

func (o POSTMarkets201ResponseDataRelationshipsTaxCalculatorData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTMarkets201ResponseDataRelationshipsTaxCalculatorData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePOSTMarkets201ResponseDataRelationshipsTaxCalculatorData struct {
	value *POSTMarkets201ResponseDataRelationshipsTaxCalculatorData
	isSet bool
}

func (v NullablePOSTMarkets201ResponseDataRelationshipsTaxCalculatorData) Get() *POSTMarkets201ResponseDataRelationshipsTaxCalculatorData {
	return v.value
}

func (v *NullablePOSTMarkets201ResponseDataRelationshipsTaxCalculatorData) Set(val *POSTMarkets201ResponseDataRelationshipsTaxCalculatorData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTMarkets201ResponseDataRelationshipsTaxCalculatorData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTMarkets201ResponseDataRelationshipsTaxCalculatorData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTMarkets201ResponseDataRelationshipsTaxCalculatorData(val *POSTMarkets201ResponseDataRelationshipsTaxCalculatorData) *NullablePOSTMarkets201ResponseDataRelationshipsTaxCalculatorData {
	return &NullablePOSTMarkets201ResponseDataRelationshipsTaxCalculatorData{value: val, isSet: true}
}

func (v NullablePOSTMarkets201ResponseDataRelationshipsTaxCalculatorData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTMarkets201ResponseDataRelationshipsTaxCalculatorData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
