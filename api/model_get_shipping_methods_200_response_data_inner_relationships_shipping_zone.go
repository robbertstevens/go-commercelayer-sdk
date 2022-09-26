/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETShippingMethods200ResponseDataInnerRelationshipsShippingZone struct for GETShippingMethods200ResponseDataInnerRelationshipsShippingZone
type GETShippingMethods200ResponseDataInnerRelationshipsShippingZone struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks          `json:"links,omitempty"`
	Data  *GETShippingMethods200ResponseDataInnerRelationshipsShippingZoneData `json:"data,omitempty"`
}

// NewGETShippingMethods200ResponseDataInnerRelationshipsShippingZone instantiates a new GETShippingMethods200ResponseDataInnerRelationshipsShippingZone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETShippingMethods200ResponseDataInnerRelationshipsShippingZone() *GETShippingMethods200ResponseDataInnerRelationshipsShippingZone {
	this := GETShippingMethods200ResponseDataInnerRelationshipsShippingZone{}
	return &this
}

// NewGETShippingMethods200ResponseDataInnerRelationshipsShippingZoneWithDefaults instantiates a new GETShippingMethods200ResponseDataInnerRelationshipsShippingZone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETShippingMethods200ResponseDataInnerRelationshipsShippingZoneWithDefaults() *GETShippingMethods200ResponseDataInnerRelationshipsShippingZone {
	this := GETShippingMethods200ResponseDataInnerRelationshipsShippingZone{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETShippingMethods200ResponseDataInnerRelationshipsShippingZone) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShippingMethods200ResponseDataInnerRelationshipsShippingZone) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETShippingMethods200ResponseDataInnerRelationshipsShippingZone) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETShippingMethods200ResponseDataInnerRelationshipsShippingZone) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETShippingMethods200ResponseDataInnerRelationshipsShippingZone) GetData() GETShippingMethods200ResponseDataInnerRelationshipsShippingZoneData {
	if o == nil || o.Data == nil {
		var ret GETShippingMethods200ResponseDataInnerRelationshipsShippingZoneData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShippingMethods200ResponseDataInnerRelationshipsShippingZone) GetDataOk() (*GETShippingMethods200ResponseDataInnerRelationshipsShippingZoneData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETShippingMethods200ResponseDataInnerRelationshipsShippingZone) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETShippingMethods200ResponseDataInnerRelationshipsShippingZoneData and assigns it to the Data field.
func (o *GETShippingMethods200ResponseDataInnerRelationshipsShippingZone) SetData(v GETShippingMethods200ResponseDataInnerRelationshipsShippingZoneData) {
	o.Data = &v
}

func (o GETShippingMethods200ResponseDataInnerRelationshipsShippingZone) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETShippingMethods200ResponseDataInnerRelationshipsShippingZone struct {
	value *GETShippingMethods200ResponseDataInnerRelationshipsShippingZone
	isSet bool
}

func (v NullableGETShippingMethods200ResponseDataInnerRelationshipsShippingZone) Get() *GETShippingMethods200ResponseDataInnerRelationshipsShippingZone {
	return v.value
}

func (v *NullableGETShippingMethods200ResponseDataInnerRelationshipsShippingZone) Set(val *GETShippingMethods200ResponseDataInnerRelationshipsShippingZone) {
	v.value = val
	v.isSet = true
}

func (v NullableGETShippingMethods200ResponseDataInnerRelationshipsShippingZone) IsSet() bool {
	return v.isSet
}

func (v *NullableGETShippingMethods200ResponseDataInnerRelationshipsShippingZone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETShippingMethods200ResponseDataInnerRelationshipsShippingZone(val *GETShippingMethods200ResponseDataInnerRelationshipsShippingZone) *NullableGETShippingMethods200ResponseDataInnerRelationshipsShippingZone {
	return &NullableGETShippingMethods200ResponseDataInnerRelationshipsShippingZone{value: val, isSet: true}
}

func (v NullableGETShippingMethods200ResponseDataInnerRelationshipsShippingZone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETShippingMethods200ResponseDataInnerRelationshipsShippingZone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
