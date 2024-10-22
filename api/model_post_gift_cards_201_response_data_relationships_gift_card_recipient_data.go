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

// checks if the POSTGiftCards201ResponseDataRelationshipsGiftCardRecipientData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTGiftCards201ResponseDataRelationshipsGiftCardRecipientData{}

// POSTGiftCards201ResponseDataRelationshipsGiftCardRecipientData struct for POSTGiftCards201ResponseDataRelationshipsGiftCardRecipientData
type POSTGiftCards201ResponseDataRelationshipsGiftCardRecipientData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewPOSTGiftCards201ResponseDataRelationshipsGiftCardRecipientData instantiates a new POSTGiftCards201ResponseDataRelationshipsGiftCardRecipientData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTGiftCards201ResponseDataRelationshipsGiftCardRecipientData() *POSTGiftCards201ResponseDataRelationshipsGiftCardRecipientData {
	this := POSTGiftCards201ResponseDataRelationshipsGiftCardRecipientData{}
	return &this
}

// NewPOSTGiftCards201ResponseDataRelationshipsGiftCardRecipientDataWithDefaults instantiates a new POSTGiftCards201ResponseDataRelationshipsGiftCardRecipientData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTGiftCards201ResponseDataRelationshipsGiftCardRecipientDataWithDefaults() *POSTGiftCards201ResponseDataRelationshipsGiftCardRecipientData {
	this := POSTGiftCards201ResponseDataRelationshipsGiftCardRecipientData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTGiftCards201ResponseDataRelationshipsGiftCardRecipientData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTGiftCards201ResponseDataRelationshipsGiftCardRecipientData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *POSTGiftCards201ResponseDataRelationshipsGiftCardRecipientData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *POSTGiftCards201ResponseDataRelationshipsGiftCardRecipientData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTGiftCards201ResponseDataRelationshipsGiftCardRecipientData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTGiftCards201ResponseDataRelationshipsGiftCardRecipientData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *POSTGiftCards201ResponseDataRelationshipsGiftCardRecipientData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *POSTGiftCards201ResponseDataRelationshipsGiftCardRecipientData) SetId(v interface{}) {
	o.Id = v
}

func (o POSTGiftCards201ResponseDataRelationshipsGiftCardRecipientData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTGiftCards201ResponseDataRelationshipsGiftCardRecipientData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePOSTGiftCards201ResponseDataRelationshipsGiftCardRecipientData struct {
	value *POSTGiftCards201ResponseDataRelationshipsGiftCardRecipientData
	isSet bool
}

func (v NullablePOSTGiftCards201ResponseDataRelationshipsGiftCardRecipientData) Get() *POSTGiftCards201ResponseDataRelationshipsGiftCardRecipientData {
	return v.value
}

func (v *NullablePOSTGiftCards201ResponseDataRelationshipsGiftCardRecipientData) Set(val *POSTGiftCards201ResponseDataRelationshipsGiftCardRecipientData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTGiftCards201ResponseDataRelationshipsGiftCardRecipientData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTGiftCards201ResponseDataRelationshipsGiftCardRecipientData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTGiftCards201ResponseDataRelationshipsGiftCardRecipientData(val *POSTGiftCards201ResponseDataRelationshipsGiftCardRecipientData) *NullablePOSTGiftCards201ResponseDataRelationshipsGiftCardRecipientData {
	return &NullablePOSTGiftCards201ResponseDataRelationshipsGiftCardRecipientData{value: val, isSet: true}
}

func (v NullablePOSTGiftCards201ResponseDataRelationshipsGiftCardRecipientData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTGiftCards201ResponseDataRelationshipsGiftCardRecipientData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
