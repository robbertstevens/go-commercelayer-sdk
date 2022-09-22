/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ImportDataRelationships struct for ImportDataRelationships
type ImportDataRelationships struct {
	Events *CustomerAddressDataRelationshipsEvents `json:"events,omitempty"`
}

// NewImportDataRelationships instantiates a new ImportDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportDataRelationships() *ImportDataRelationships {
	this := ImportDataRelationships{}
	return &this
}

// NewImportDataRelationshipsWithDefaults instantiates a new ImportDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportDataRelationshipsWithDefaults() *ImportDataRelationships {
	this := ImportDataRelationships{}
	return &this
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *ImportDataRelationships) GetEvents() CustomerAddressDataRelationshipsEvents {
	if o == nil || o.Events == nil {
		var ret CustomerAddressDataRelationshipsEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportDataRelationships) GetEventsOk() (*CustomerAddressDataRelationshipsEvents, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *ImportDataRelationships) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given CustomerAddressDataRelationshipsEvents and assigns it to the Events field.
func (o *ImportDataRelationships) SetEvents(v CustomerAddressDataRelationshipsEvents) {
	o.Events = &v
}

func (o ImportDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableImportDataRelationships struct {
	value *ImportDataRelationships
	isSet bool
}

func (v NullableImportDataRelationships) Get() *ImportDataRelationships {
	return v.value
}

func (v *NullableImportDataRelationships) Set(val *ImportDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableImportDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableImportDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportDataRelationships(val *ImportDataRelationships) *NullableImportDataRelationships {
	return &NullableImportDataRelationships{value: val, isSet: true}
}

func (v NullableImportDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
