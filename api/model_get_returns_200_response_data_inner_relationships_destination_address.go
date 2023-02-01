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

// GETReturns200ResponseDataInnerRelationshipsDestinationAddress struct for GETReturns200ResponseDataInnerRelationshipsDestinationAddress
type GETReturns200ResponseDataInnerRelationshipsDestinationAddress struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks        `json:"links,omitempty"`
	Data  *GETReturns200ResponseDataInnerRelationshipsDestinationAddressData `json:"data,omitempty"`
}

// NewGETReturns200ResponseDataInnerRelationshipsDestinationAddress instantiates a new GETReturns200ResponseDataInnerRelationshipsDestinationAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETReturns200ResponseDataInnerRelationshipsDestinationAddress() *GETReturns200ResponseDataInnerRelationshipsDestinationAddress {
	this := GETReturns200ResponseDataInnerRelationshipsDestinationAddress{}
	return &this
}

// NewGETReturns200ResponseDataInnerRelationshipsDestinationAddressWithDefaults instantiates a new GETReturns200ResponseDataInnerRelationshipsDestinationAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETReturns200ResponseDataInnerRelationshipsDestinationAddressWithDefaults() *GETReturns200ResponseDataInnerRelationshipsDestinationAddress {
	this := GETReturns200ResponseDataInnerRelationshipsDestinationAddress{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETReturns200ResponseDataInnerRelationshipsDestinationAddress) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETReturns200ResponseDataInnerRelationshipsDestinationAddress) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETReturns200ResponseDataInnerRelationshipsDestinationAddress) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETReturns200ResponseDataInnerRelationshipsDestinationAddress) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETReturns200ResponseDataInnerRelationshipsDestinationAddress) GetData() GETReturns200ResponseDataInnerRelationshipsDestinationAddressData {
	if o == nil || o.Data == nil {
		var ret GETReturns200ResponseDataInnerRelationshipsDestinationAddressData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETReturns200ResponseDataInnerRelationshipsDestinationAddress) GetDataOk() (*GETReturns200ResponseDataInnerRelationshipsDestinationAddressData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETReturns200ResponseDataInnerRelationshipsDestinationAddress) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETReturns200ResponseDataInnerRelationshipsDestinationAddressData and assigns it to the Data field.
func (o *GETReturns200ResponseDataInnerRelationshipsDestinationAddress) SetData(v GETReturns200ResponseDataInnerRelationshipsDestinationAddressData) {
	o.Data = &v
}

func (o GETReturns200ResponseDataInnerRelationshipsDestinationAddress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETReturns200ResponseDataInnerRelationshipsDestinationAddress struct {
	value *GETReturns200ResponseDataInnerRelationshipsDestinationAddress
	isSet bool
}

func (v NullableGETReturns200ResponseDataInnerRelationshipsDestinationAddress) Get() *GETReturns200ResponseDataInnerRelationshipsDestinationAddress {
	return v.value
}

func (v *NullableGETReturns200ResponseDataInnerRelationshipsDestinationAddress) Set(val *GETReturns200ResponseDataInnerRelationshipsDestinationAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableGETReturns200ResponseDataInnerRelationshipsDestinationAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableGETReturns200ResponseDataInnerRelationshipsDestinationAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETReturns200ResponseDataInnerRelationshipsDestinationAddress(val *GETReturns200ResponseDataInnerRelationshipsDestinationAddress) *NullableGETReturns200ResponseDataInnerRelationshipsDestinationAddress {
	return &NullableGETReturns200ResponseDataInnerRelationshipsDestinationAddress{value: val, isSet: true}
}

func (v NullableGETReturns200ResponseDataInnerRelationshipsDestinationAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETReturns200ResponseDataInnerRelationshipsDestinationAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
