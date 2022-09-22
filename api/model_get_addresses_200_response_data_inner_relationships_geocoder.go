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

// GETAddresses200ResponseDataInnerRelationshipsGeocoder struct for GETAddresses200ResponseDataInnerRelationshipsGeocoder
type GETAddresses200ResponseDataInnerRelationshipsGeocoder struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks `json:"links,omitempty"`
	Data  *GETAddresses200ResponseDataInnerRelationshipsGeocoderData  `json:"data,omitempty"`
}

// NewGETAddresses200ResponseDataInnerRelationshipsGeocoder instantiates a new GETAddresses200ResponseDataInnerRelationshipsGeocoder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETAddresses200ResponseDataInnerRelationshipsGeocoder() *GETAddresses200ResponseDataInnerRelationshipsGeocoder {
	this := GETAddresses200ResponseDataInnerRelationshipsGeocoder{}
	return &this
}

// NewGETAddresses200ResponseDataInnerRelationshipsGeocoderWithDefaults instantiates a new GETAddresses200ResponseDataInnerRelationshipsGeocoder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETAddresses200ResponseDataInnerRelationshipsGeocoderWithDefaults() *GETAddresses200ResponseDataInnerRelationshipsGeocoder {
	this := GETAddresses200ResponseDataInnerRelationshipsGeocoder{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETAddresses200ResponseDataInnerRelationshipsGeocoder) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAddresses200ResponseDataInnerRelationshipsGeocoder) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETAddresses200ResponseDataInnerRelationshipsGeocoder) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETAddresses200ResponseDataInnerRelationshipsGeocoder) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETAddresses200ResponseDataInnerRelationshipsGeocoder) GetData() GETAddresses200ResponseDataInnerRelationshipsGeocoderData {
	if o == nil || o.Data == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAddresses200ResponseDataInnerRelationshipsGeocoder) GetDataOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETAddresses200ResponseDataInnerRelationshipsGeocoder) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderData and assigns it to the Data field.
func (o *GETAddresses200ResponseDataInnerRelationshipsGeocoder) SetData(v GETAddresses200ResponseDataInnerRelationshipsGeocoderData) {
	o.Data = &v
}

func (o GETAddresses200ResponseDataInnerRelationshipsGeocoder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETAddresses200ResponseDataInnerRelationshipsGeocoder struct {
	value *GETAddresses200ResponseDataInnerRelationshipsGeocoder
	isSet bool
}

func (v NullableGETAddresses200ResponseDataInnerRelationshipsGeocoder) Get() *GETAddresses200ResponseDataInnerRelationshipsGeocoder {
	return v.value
}

func (v *NullableGETAddresses200ResponseDataInnerRelationshipsGeocoder) Set(val *GETAddresses200ResponseDataInnerRelationshipsGeocoder) {
	v.value = val
	v.isSet = true
}

func (v NullableGETAddresses200ResponseDataInnerRelationshipsGeocoder) IsSet() bool {
	return v.isSet
}

func (v *NullableGETAddresses200ResponseDataInnerRelationshipsGeocoder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETAddresses200ResponseDataInnerRelationshipsGeocoder(val *GETAddresses200ResponseDataInnerRelationshipsGeocoder) *NullableGETAddresses200ResponseDataInnerRelationshipsGeocoder {
	return &NullableGETAddresses200ResponseDataInnerRelationshipsGeocoder{value: val, isSet: true}
}

func (v NullableGETAddresses200ResponseDataInnerRelationshipsGeocoder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETAddresses200ResponseDataInnerRelationshipsGeocoder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
