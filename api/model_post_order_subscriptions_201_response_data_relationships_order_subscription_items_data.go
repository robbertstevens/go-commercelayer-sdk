/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 4.1.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItemsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItemsData{}

// POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItemsData struct for POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItemsData
type POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItemsData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewPOSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItemsData instantiates a new POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItemsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItemsData() *POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItemsData {
	this := POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItemsData{}
	return &this
}

// NewPOSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItemsDataWithDefaults instantiates a new POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItemsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItemsDataWithDefaults() *POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItemsData {
	this := POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItemsData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItemsData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItemsData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItemsData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItemsData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItemsData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItemsData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItemsData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItemsData) SetId(v interface{}) {
	o.Id = v
}

func (o POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItemsData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItemsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePOSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItemsData struct {
	value *POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItemsData
	isSet bool
}

func (v NullablePOSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItemsData) Get() *POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItemsData {
	return v.value
}

func (v *NullablePOSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItemsData) Set(val *POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItemsData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItemsData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItemsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItemsData(val *POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItemsData) *NullablePOSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItemsData {
	return &NullablePOSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItemsData{value: val, isSet: true}
}

func (v NullablePOSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItemsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItemsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
