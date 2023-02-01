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

// GETCaptures200ResponseDataInnerRelationshipsRefunds struct for GETCaptures200ResponseDataInnerRelationshipsRefunds
type GETCaptures200ResponseDataInnerRelationshipsRefunds struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks `json:"links,omitempty"`
	Data  *GETCaptures200ResponseDataInnerRelationshipsRefundsData    `json:"data,omitempty"`
}

// NewGETCaptures200ResponseDataInnerRelationshipsRefunds instantiates a new GETCaptures200ResponseDataInnerRelationshipsRefunds object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCaptures200ResponseDataInnerRelationshipsRefunds() *GETCaptures200ResponseDataInnerRelationshipsRefunds {
	this := GETCaptures200ResponseDataInnerRelationshipsRefunds{}
	return &this
}

// NewGETCaptures200ResponseDataInnerRelationshipsRefundsWithDefaults instantiates a new GETCaptures200ResponseDataInnerRelationshipsRefunds object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCaptures200ResponseDataInnerRelationshipsRefundsWithDefaults() *GETCaptures200ResponseDataInnerRelationshipsRefunds {
	this := GETCaptures200ResponseDataInnerRelationshipsRefunds{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETCaptures200ResponseDataInnerRelationshipsRefunds) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCaptures200ResponseDataInnerRelationshipsRefunds) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETCaptures200ResponseDataInnerRelationshipsRefunds) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETCaptures200ResponseDataInnerRelationshipsRefunds) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETCaptures200ResponseDataInnerRelationshipsRefunds) GetData() GETCaptures200ResponseDataInnerRelationshipsRefundsData {
	if o == nil || o.Data == nil {
		var ret GETCaptures200ResponseDataInnerRelationshipsRefundsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCaptures200ResponseDataInnerRelationshipsRefunds) GetDataOk() (*GETCaptures200ResponseDataInnerRelationshipsRefundsData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETCaptures200ResponseDataInnerRelationshipsRefunds) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETCaptures200ResponseDataInnerRelationshipsRefundsData and assigns it to the Data field.
func (o *GETCaptures200ResponseDataInnerRelationshipsRefunds) SetData(v GETCaptures200ResponseDataInnerRelationshipsRefundsData) {
	o.Data = &v
}

func (o GETCaptures200ResponseDataInnerRelationshipsRefunds) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETCaptures200ResponseDataInnerRelationshipsRefunds struct {
	value *GETCaptures200ResponseDataInnerRelationshipsRefunds
	isSet bool
}

func (v NullableGETCaptures200ResponseDataInnerRelationshipsRefunds) Get() *GETCaptures200ResponseDataInnerRelationshipsRefunds {
	return v.value
}

func (v *NullableGETCaptures200ResponseDataInnerRelationshipsRefunds) Set(val *GETCaptures200ResponseDataInnerRelationshipsRefunds) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCaptures200ResponseDataInnerRelationshipsRefunds) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCaptures200ResponseDataInnerRelationshipsRefunds) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCaptures200ResponseDataInnerRelationshipsRefunds(val *GETCaptures200ResponseDataInnerRelationshipsRefunds) *NullableGETCaptures200ResponseDataInnerRelationshipsRefunds {
	return &NullableGETCaptures200ResponseDataInnerRelationshipsRefunds{value: val, isSet: true}
}

func (v NullableGETCaptures200ResponseDataInnerRelationshipsRefunds) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCaptures200ResponseDataInnerRelationshipsRefunds) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
