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

// checks if the POSTStockTransfers201ResponseDataRelationshipsOriginStockLocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTStockTransfers201ResponseDataRelationshipsOriginStockLocation{}

// POSTStockTransfers201ResponseDataRelationshipsOriginStockLocation struct for POSTStockTransfers201ResponseDataRelationshipsOriginStockLocation
type POSTStockTransfers201ResponseDataRelationshipsOriginStockLocation struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks                `json:"links,omitempty"`
	Data  *POSTStockTransfers201ResponseDataRelationshipsOriginStockLocationData `json:"data,omitempty"`
}

// NewPOSTStockTransfers201ResponseDataRelationshipsOriginStockLocation instantiates a new POSTStockTransfers201ResponseDataRelationshipsOriginStockLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTStockTransfers201ResponseDataRelationshipsOriginStockLocation() *POSTStockTransfers201ResponseDataRelationshipsOriginStockLocation {
	this := POSTStockTransfers201ResponseDataRelationshipsOriginStockLocation{}
	return &this
}

// NewPOSTStockTransfers201ResponseDataRelationshipsOriginStockLocationWithDefaults instantiates a new POSTStockTransfers201ResponseDataRelationshipsOriginStockLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTStockTransfers201ResponseDataRelationshipsOriginStockLocationWithDefaults() *POSTStockTransfers201ResponseDataRelationshipsOriginStockLocation {
	this := POSTStockTransfers201ResponseDataRelationshipsOriginStockLocation{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTStockTransfers201ResponseDataRelationshipsOriginStockLocation) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTStockTransfers201ResponseDataRelationshipsOriginStockLocation) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTStockTransfers201ResponseDataRelationshipsOriginStockLocation) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTStockTransfers201ResponseDataRelationshipsOriginStockLocation) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTStockTransfers201ResponseDataRelationshipsOriginStockLocation) GetData() POSTStockTransfers201ResponseDataRelationshipsOriginStockLocationData {
	if o == nil || IsNil(o.Data) {
		var ret POSTStockTransfers201ResponseDataRelationshipsOriginStockLocationData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTStockTransfers201ResponseDataRelationshipsOriginStockLocation) GetDataOk() (*POSTStockTransfers201ResponseDataRelationshipsOriginStockLocationData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTStockTransfers201ResponseDataRelationshipsOriginStockLocation) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTStockTransfers201ResponseDataRelationshipsOriginStockLocationData and assigns it to the Data field.
func (o *POSTStockTransfers201ResponseDataRelationshipsOriginStockLocation) SetData(v POSTStockTransfers201ResponseDataRelationshipsOriginStockLocationData) {
	o.Data = &v
}

func (o POSTStockTransfers201ResponseDataRelationshipsOriginStockLocation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTStockTransfers201ResponseDataRelationshipsOriginStockLocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTStockTransfers201ResponseDataRelationshipsOriginStockLocation struct {
	value *POSTStockTransfers201ResponseDataRelationshipsOriginStockLocation
	isSet bool
}

func (v NullablePOSTStockTransfers201ResponseDataRelationshipsOriginStockLocation) Get() *POSTStockTransfers201ResponseDataRelationshipsOriginStockLocation {
	return v.value
}

func (v *NullablePOSTStockTransfers201ResponseDataRelationshipsOriginStockLocation) Set(val *POSTStockTransfers201ResponseDataRelationshipsOriginStockLocation) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTStockTransfers201ResponseDataRelationshipsOriginStockLocation) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTStockTransfers201ResponseDataRelationshipsOriginStockLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTStockTransfers201ResponseDataRelationshipsOriginStockLocation(val *POSTStockTransfers201ResponseDataRelationshipsOriginStockLocation) *NullablePOSTStockTransfers201ResponseDataRelationshipsOriginStockLocation {
	return &NullablePOSTStockTransfers201ResponseDataRelationshipsOriginStockLocation{value: val, isSet: true}
}

func (v NullablePOSTStockTransfers201ResponseDataRelationshipsOriginStockLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTStockTransfers201ResponseDataRelationshipsOriginStockLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
