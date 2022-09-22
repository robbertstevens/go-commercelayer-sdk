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

// GETAddresses200ResponseDataInnerRelationships struct for GETAddresses200ResponseDataInnerRelationships
type GETAddresses200ResponseDataInnerRelationships struct {
	Geocoder *GETAddresses200ResponseDataInnerRelationshipsGeocoder `json:"geocoder,omitempty"`
}

// NewGETAddresses200ResponseDataInnerRelationships instantiates a new GETAddresses200ResponseDataInnerRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETAddresses200ResponseDataInnerRelationships() *GETAddresses200ResponseDataInnerRelationships {
	this := GETAddresses200ResponseDataInnerRelationships{}
	return &this
}

// NewGETAddresses200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETAddresses200ResponseDataInnerRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETAddresses200ResponseDataInnerRelationshipsWithDefaults() *GETAddresses200ResponseDataInnerRelationships {
	this := GETAddresses200ResponseDataInnerRelationships{}
	return &this
}

// GetGeocoder returns the Geocoder field value if set, zero value otherwise.
func (o *GETAddresses200ResponseDataInnerRelationships) GetGeocoder() GETAddresses200ResponseDataInnerRelationshipsGeocoder {
	if o == nil || o.Geocoder == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoder
		return ret
	}
	return *o.Geocoder
}

// GetGeocoderOk returns a tuple with the Geocoder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAddresses200ResponseDataInnerRelationships) GetGeocoderOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoder, bool) {
	if o == nil || o.Geocoder == nil {
		return nil, false
	}
	return o.Geocoder, true
}

// HasGeocoder returns a boolean if a field has been set.
func (o *GETAddresses200ResponseDataInnerRelationships) HasGeocoder() bool {
	if o != nil && o.Geocoder != nil {
		return true
	}

	return false
}

// SetGeocoder gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoder and assigns it to the Geocoder field.
func (o *GETAddresses200ResponseDataInnerRelationships) SetGeocoder(v GETAddresses200ResponseDataInnerRelationshipsGeocoder) {
	o.Geocoder = &v
}

func (o GETAddresses200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Geocoder != nil {
		toSerialize["geocoder"] = o.Geocoder
	}
	return json.Marshal(toSerialize)
}

type NullableGETAddresses200ResponseDataInnerRelationships struct {
	value *GETAddresses200ResponseDataInnerRelationships
	isSet bool
}

func (v NullableGETAddresses200ResponseDataInnerRelationships) Get() *GETAddresses200ResponseDataInnerRelationships {
	return v.value
}

func (v *NullableGETAddresses200ResponseDataInnerRelationships) Set(val *GETAddresses200ResponseDataInnerRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETAddresses200ResponseDataInnerRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETAddresses200ResponseDataInnerRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETAddresses200ResponseDataInnerRelationships(val *GETAddresses200ResponseDataInnerRelationships) *NullableGETAddresses200ResponseDataInnerRelationships {
	return &NullableGETAddresses200ResponseDataInnerRelationships{value: val, isSet: true}
}

func (v NullableGETAddresses200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETAddresses200ResponseDataInnerRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
