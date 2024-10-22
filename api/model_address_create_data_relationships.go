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

// checks if the AddressCreateDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddressCreateDataRelationships{}

// AddressCreateDataRelationships struct for AddressCreateDataRelationships
type AddressCreateDataRelationships struct {
	Geocoder *AddressCreateDataRelationshipsGeocoder `json:"geocoder,omitempty"`
	Tags     *AddressCreateDataRelationshipsTags     `json:"tags,omitempty"`
}

// NewAddressCreateDataRelationships instantiates a new AddressCreateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressCreateDataRelationships() *AddressCreateDataRelationships {
	this := AddressCreateDataRelationships{}
	return &this
}

// NewAddressCreateDataRelationshipsWithDefaults instantiates a new AddressCreateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressCreateDataRelationshipsWithDefaults() *AddressCreateDataRelationships {
	this := AddressCreateDataRelationships{}
	return &this
}

// GetGeocoder returns the Geocoder field value if set, zero value otherwise.
func (o *AddressCreateDataRelationships) GetGeocoder() AddressCreateDataRelationshipsGeocoder {
	if o == nil || IsNil(o.Geocoder) {
		var ret AddressCreateDataRelationshipsGeocoder
		return ret
	}
	return *o.Geocoder
}

// GetGeocoderOk returns a tuple with the Geocoder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressCreateDataRelationships) GetGeocoderOk() (*AddressCreateDataRelationshipsGeocoder, bool) {
	if o == nil || IsNil(o.Geocoder) {
		return nil, false
	}
	return o.Geocoder, true
}

// HasGeocoder returns a boolean if a field has been set.
func (o *AddressCreateDataRelationships) HasGeocoder() bool {
	if o != nil && !IsNil(o.Geocoder) {
		return true
	}

	return false
}

// SetGeocoder gets a reference to the given AddressCreateDataRelationshipsGeocoder and assigns it to the Geocoder field.
func (o *AddressCreateDataRelationships) SetGeocoder(v AddressCreateDataRelationshipsGeocoder) {
	o.Geocoder = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *AddressCreateDataRelationships) GetTags() AddressCreateDataRelationshipsTags {
	if o == nil || IsNil(o.Tags) {
		var ret AddressCreateDataRelationshipsTags
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressCreateDataRelationships) GetTagsOk() (*AddressCreateDataRelationshipsTags, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *AddressCreateDataRelationships) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given AddressCreateDataRelationshipsTags and assigns it to the Tags field.
func (o *AddressCreateDataRelationships) SetTags(v AddressCreateDataRelationshipsTags) {
	o.Tags = &v
}

func (o AddressCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressCreateDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Geocoder) {
		toSerialize["geocoder"] = o.Geocoder
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableAddressCreateDataRelationships struct {
	value *AddressCreateDataRelationships
	isSet bool
}

func (v NullableAddressCreateDataRelationships) Get() *AddressCreateDataRelationships {
	return v.value
}

func (v *NullableAddressCreateDataRelationships) Set(val *AddressCreateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressCreateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressCreateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressCreateDataRelationships(val *AddressCreateDataRelationships) *NullableAddressCreateDataRelationships {
	return &NullableAddressCreateDataRelationships{value: val, isSet: true}
}

func (v NullableAddressCreateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressCreateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
