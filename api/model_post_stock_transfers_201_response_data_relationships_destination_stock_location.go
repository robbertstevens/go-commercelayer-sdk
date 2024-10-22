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

// checks if the POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocation{}

// POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocation struct for POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocation
type POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocation struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks                     `json:"links,omitempty"`
	Data  *POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocationData `json:"data,omitempty"`
}

// NewPOSTStockTransfers201ResponseDataRelationshipsDestinationStockLocation instantiates a new POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTStockTransfers201ResponseDataRelationshipsDestinationStockLocation() *POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocation {
	this := POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocation{}
	return &this
}

// NewPOSTStockTransfers201ResponseDataRelationshipsDestinationStockLocationWithDefaults instantiates a new POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTStockTransfers201ResponseDataRelationshipsDestinationStockLocationWithDefaults() *POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocation {
	this := POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocation{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocation) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocation) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocation) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocation) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocation) GetData() POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocationData {
	if o == nil || IsNil(o.Data) {
		var ret POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocationData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocation) GetDataOk() (*POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocationData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocation) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocationData and assigns it to the Data field.
func (o *POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocation) SetData(v POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocationData) {
	o.Data = &v
}

func (o POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTStockTransfers201ResponseDataRelationshipsDestinationStockLocation struct {
	value *POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocation
	isSet bool
}

func (v NullablePOSTStockTransfers201ResponseDataRelationshipsDestinationStockLocation) Get() *POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocation {
	return v.value
}

func (v *NullablePOSTStockTransfers201ResponseDataRelationshipsDestinationStockLocation) Set(val *POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocation) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTStockTransfers201ResponseDataRelationshipsDestinationStockLocation) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTStockTransfers201ResponseDataRelationshipsDestinationStockLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTStockTransfers201ResponseDataRelationshipsDestinationStockLocation(val *POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocation) *NullablePOSTStockTransfers201ResponseDataRelationshipsDestinationStockLocation {
	return &NullablePOSTStockTransfers201ResponseDataRelationshipsDestinationStockLocation{value: val, isSet: true}
}

func (v NullablePOSTStockTransfers201ResponseDataRelationshipsDestinationStockLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTStockTransfers201ResponseDataRelationshipsDestinationStockLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
