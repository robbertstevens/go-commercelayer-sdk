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

// checks if the GETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids{}

// GETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids struct for GETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids
type GETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks                `json:"links,omitempty"`
	Data  *GETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoidsData `json:"data,omitempty"`
}

// NewGETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids instantiates a new GETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids() *GETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids {
	this := GETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids{}
	return &this
}

// NewGETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoidsWithDefaults instantiates a new GETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoidsWithDefaults() *GETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids {
	this := GETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids) GetData() GETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoidsData {
	if o == nil || IsNil(o.Data) {
		var ret GETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoidsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids) GetDataOk() (*GETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoidsData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoidsData and assigns it to the Data field.
func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids) SetData(v GETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoidsData) {
	o.Data = &v
}

func (o GETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids struct {
	value *GETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids
	isSet bool
}

func (v NullableGETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids) Get() *GETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids {
	return v.value
}

func (v *NullableGETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids) Set(val *GETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids) {
	v.value = val
	v.isSet = true
}

func (v NullableGETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids) IsSet() bool {
	return v.isSet
}

func (v *NullableGETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids(val *GETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids) *NullableGETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids {
	return &NullableGETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids{value: val, isSet: true}
}

func (v NullableGETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
