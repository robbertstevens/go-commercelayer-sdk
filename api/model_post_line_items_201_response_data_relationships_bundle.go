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

// checks if the POSTLineItems201ResponseDataRelationshipsBundle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTLineItems201ResponseDataRelationshipsBundle{}

// POSTLineItems201ResponseDataRelationshipsBundle struct for POSTLineItems201ResponseDataRelationshipsBundle
type POSTLineItems201ResponseDataRelationshipsBundle struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks `json:"links,omitempty"`
	Data  *POSTLineItems201ResponseDataRelationshipsBundleData    `json:"data,omitempty"`
}

// NewPOSTLineItems201ResponseDataRelationshipsBundle instantiates a new POSTLineItems201ResponseDataRelationshipsBundle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTLineItems201ResponseDataRelationshipsBundle() *POSTLineItems201ResponseDataRelationshipsBundle {
	this := POSTLineItems201ResponseDataRelationshipsBundle{}
	return &this
}

// NewPOSTLineItems201ResponseDataRelationshipsBundleWithDefaults instantiates a new POSTLineItems201ResponseDataRelationshipsBundle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTLineItems201ResponseDataRelationshipsBundleWithDefaults() *POSTLineItems201ResponseDataRelationshipsBundle {
	this := POSTLineItems201ResponseDataRelationshipsBundle{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTLineItems201ResponseDataRelationshipsBundle) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTLineItems201ResponseDataRelationshipsBundle) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTLineItems201ResponseDataRelationshipsBundle) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTLineItems201ResponseDataRelationshipsBundle) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTLineItems201ResponseDataRelationshipsBundle) GetData() POSTLineItems201ResponseDataRelationshipsBundleData {
	if o == nil || IsNil(o.Data) {
		var ret POSTLineItems201ResponseDataRelationshipsBundleData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTLineItems201ResponseDataRelationshipsBundle) GetDataOk() (*POSTLineItems201ResponseDataRelationshipsBundleData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTLineItems201ResponseDataRelationshipsBundle) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTLineItems201ResponseDataRelationshipsBundleData and assigns it to the Data field.
func (o *POSTLineItems201ResponseDataRelationshipsBundle) SetData(v POSTLineItems201ResponseDataRelationshipsBundleData) {
	o.Data = &v
}

func (o POSTLineItems201ResponseDataRelationshipsBundle) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTLineItems201ResponseDataRelationshipsBundle) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTLineItems201ResponseDataRelationshipsBundle struct {
	value *POSTLineItems201ResponseDataRelationshipsBundle
	isSet bool
}

func (v NullablePOSTLineItems201ResponseDataRelationshipsBundle) Get() *POSTLineItems201ResponseDataRelationshipsBundle {
	return v.value
}

func (v *NullablePOSTLineItems201ResponseDataRelationshipsBundle) Set(val *POSTLineItems201ResponseDataRelationshipsBundle) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTLineItems201ResponseDataRelationshipsBundle) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTLineItems201ResponseDataRelationshipsBundle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTLineItems201ResponseDataRelationshipsBundle(val *POSTLineItems201ResponseDataRelationshipsBundle) *NullablePOSTLineItems201ResponseDataRelationshipsBundle {
	return &NullablePOSTLineItems201ResponseDataRelationshipsBundle{value: val, isSet: true}
}

func (v NullablePOSTLineItems201ResponseDataRelationshipsBundle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTLineItems201ResponseDataRelationshipsBundle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
