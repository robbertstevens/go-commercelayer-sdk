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

// checks if the POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation{}

// POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation struct for POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation
type POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks             `json:"links,omitempty"`
	Data  *POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocationData `json:"data,omitempty"`
}

// NewPOSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation instantiates a new POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation() *POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation {
	this := POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation{}
	return &this
}

// NewPOSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocationWithDefaults instantiates a new POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocationWithDefaults() *POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation {
	this := POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation) GetData() POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocationData {
	if o == nil || IsNil(o.Data) {
		var ret POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocationData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation) GetDataOk() (*POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocationData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocationData and assigns it to the Data field.
func (o *POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation) SetData(v POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocationData) {
	o.Data = &v
}

func (o POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation struct {
	value *POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation
	isSet bool
}

func (v NullablePOSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation) Get() *POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation {
	return v.value
}

func (v *NullablePOSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation) Set(val *POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation(val *POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation) *NullablePOSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation {
	return &NullablePOSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation{value: val, isSet: true}
}

func (v NullablePOSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
