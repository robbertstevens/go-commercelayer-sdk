/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.4.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETAuthorizations200ResponseDataInnerRelationshipsEvents struct for GETAuthorizations200ResponseDataInnerRelationshipsEvents
type GETAuthorizations200ResponseDataInnerRelationshipsEvents struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks   `json:"links,omitempty"`
	Data  *GETAuthorizations200ResponseDataInnerRelationshipsEventsData `json:"data,omitempty"`
}

// NewGETAuthorizations200ResponseDataInnerRelationshipsEvents instantiates a new GETAuthorizations200ResponseDataInnerRelationshipsEvents object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETAuthorizations200ResponseDataInnerRelationshipsEvents() *GETAuthorizations200ResponseDataInnerRelationshipsEvents {
	this := GETAuthorizations200ResponseDataInnerRelationshipsEvents{}
	return &this
}

// NewGETAuthorizations200ResponseDataInnerRelationshipsEventsWithDefaults instantiates a new GETAuthorizations200ResponseDataInnerRelationshipsEvents object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETAuthorizations200ResponseDataInnerRelationshipsEventsWithDefaults() *GETAuthorizations200ResponseDataInnerRelationshipsEvents {
	this := GETAuthorizations200ResponseDataInnerRelationshipsEvents{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETAuthorizations200ResponseDataInnerRelationshipsEvents) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAuthorizations200ResponseDataInnerRelationshipsEvents) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETAuthorizations200ResponseDataInnerRelationshipsEvents) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETAuthorizations200ResponseDataInnerRelationshipsEvents) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETAuthorizations200ResponseDataInnerRelationshipsEvents) GetData() GETAuthorizations200ResponseDataInnerRelationshipsEventsData {
	if o == nil || o.Data == nil {
		var ret GETAuthorizations200ResponseDataInnerRelationshipsEventsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAuthorizations200ResponseDataInnerRelationshipsEvents) GetDataOk() (*GETAuthorizations200ResponseDataInnerRelationshipsEventsData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETAuthorizations200ResponseDataInnerRelationshipsEvents) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETAuthorizations200ResponseDataInnerRelationshipsEventsData and assigns it to the Data field.
func (o *GETAuthorizations200ResponseDataInnerRelationshipsEvents) SetData(v GETAuthorizations200ResponseDataInnerRelationshipsEventsData) {
	o.Data = &v
}

func (o GETAuthorizations200ResponseDataInnerRelationshipsEvents) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETAuthorizations200ResponseDataInnerRelationshipsEvents struct {
	value *GETAuthorizations200ResponseDataInnerRelationshipsEvents
	isSet bool
}

func (v NullableGETAuthorizations200ResponseDataInnerRelationshipsEvents) Get() *GETAuthorizations200ResponseDataInnerRelationshipsEvents {
	return v.value
}

func (v *NullableGETAuthorizations200ResponseDataInnerRelationshipsEvents) Set(val *GETAuthorizations200ResponseDataInnerRelationshipsEvents) {
	v.value = val
	v.isSet = true
}

func (v NullableGETAuthorizations200ResponseDataInnerRelationshipsEvents) IsSet() bool {
	return v.isSet
}

func (v *NullableGETAuthorizations200ResponseDataInnerRelationshipsEvents) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETAuthorizations200ResponseDataInnerRelationshipsEvents(val *GETAuthorizations200ResponseDataInnerRelationshipsEvents) *NullableGETAuthorizations200ResponseDataInnerRelationshipsEvents {
	return &NullableGETAuthorizations200ResponseDataInnerRelationshipsEvents{value: val, isSet: true}
}

func (v NullableGETAuthorizations200ResponseDataInnerRelationshipsEvents) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETAuthorizations200ResponseDataInnerRelationshipsEvents) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
