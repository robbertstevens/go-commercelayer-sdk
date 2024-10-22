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

// checks if the POSTExports201ResponseDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTExports201ResponseDataRelationships{}

// POSTExports201ResponseDataRelationships struct for POSTExports201ResponseDataRelationships
type POSTExports201ResponseDataRelationships struct {
	Events *POSTAddresses201ResponseDataRelationshipsEvents `json:"events,omitempty"`
}

// NewPOSTExports201ResponseDataRelationships instantiates a new POSTExports201ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTExports201ResponseDataRelationships() *POSTExports201ResponseDataRelationships {
	this := POSTExports201ResponseDataRelationships{}
	return &this
}

// NewPOSTExports201ResponseDataRelationshipsWithDefaults instantiates a new POSTExports201ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTExports201ResponseDataRelationshipsWithDefaults() *POSTExports201ResponseDataRelationships {
	this := POSTExports201ResponseDataRelationships{}
	return &this
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *POSTExports201ResponseDataRelationships) GetEvents() POSTAddresses201ResponseDataRelationshipsEvents {
	if o == nil || IsNil(o.Events) {
		var ret POSTAddresses201ResponseDataRelationshipsEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTExports201ResponseDataRelationships) GetEventsOk() (*POSTAddresses201ResponseDataRelationshipsEvents, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *POSTExports201ResponseDataRelationships) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given POSTAddresses201ResponseDataRelationshipsEvents and assigns it to the Events field.
func (o *POSTExports201ResponseDataRelationships) SetEvents(v POSTAddresses201ResponseDataRelationshipsEvents) {
	o.Events = &v
}

func (o POSTExports201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTExports201ResponseDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	return toSerialize, nil
}

type NullablePOSTExports201ResponseDataRelationships struct {
	value *POSTExports201ResponseDataRelationships
	isSet bool
}

func (v NullablePOSTExports201ResponseDataRelationships) Get() *POSTExports201ResponseDataRelationships {
	return v.value
}

func (v *NullablePOSTExports201ResponseDataRelationships) Set(val *POSTExports201ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTExports201ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTExports201ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTExports201ResponseDataRelationships(val *POSTExports201ResponseDataRelationships) *NullablePOSTExports201ResponseDataRelationships {
	return &NullablePOSTExports201ResponseDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTExports201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTExports201ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
