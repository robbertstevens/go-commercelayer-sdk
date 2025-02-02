/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.3.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the POSTMarkets201ResponseDataRelationshipsBasePriceList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTMarkets201ResponseDataRelationshipsBasePriceList{}

// POSTMarkets201ResponseDataRelationshipsBasePriceList struct for POSTMarkets201ResponseDataRelationshipsBasePriceList
type POSTMarkets201ResponseDataRelationshipsBasePriceList struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks   `json:"links,omitempty"`
	Data  *POSTMarkets201ResponseDataRelationshipsBasePriceListData `json:"data,omitempty"`
}

// NewPOSTMarkets201ResponseDataRelationshipsBasePriceList instantiates a new POSTMarkets201ResponseDataRelationshipsBasePriceList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTMarkets201ResponseDataRelationshipsBasePriceList() *POSTMarkets201ResponseDataRelationshipsBasePriceList {
	this := POSTMarkets201ResponseDataRelationshipsBasePriceList{}
	return &this
}

// NewPOSTMarkets201ResponseDataRelationshipsBasePriceListWithDefaults instantiates a new POSTMarkets201ResponseDataRelationshipsBasePriceList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTMarkets201ResponseDataRelationshipsBasePriceListWithDefaults() *POSTMarkets201ResponseDataRelationshipsBasePriceList {
	this := POSTMarkets201ResponseDataRelationshipsBasePriceList{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTMarkets201ResponseDataRelationshipsBasePriceList) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTMarkets201ResponseDataRelationshipsBasePriceList) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTMarkets201ResponseDataRelationshipsBasePriceList) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTMarkets201ResponseDataRelationshipsBasePriceList) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTMarkets201ResponseDataRelationshipsBasePriceList) GetData() POSTMarkets201ResponseDataRelationshipsBasePriceListData {
	if o == nil || IsNil(o.Data) {
		var ret POSTMarkets201ResponseDataRelationshipsBasePriceListData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTMarkets201ResponseDataRelationshipsBasePriceList) GetDataOk() (*POSTMarkets201ResponseDataRelationshipsBasePriceListData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTMarkets201ResponseDataRelationshipsBasePriceList) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTMarkets201ResponseDataRelationshipsBasePriceListData and assigns it to the Data field.
func (o *POSTMarkets201ResponseDataRelationshipsBasePriceList) SetData(v POSTMarkets201ResponseDataRelationshipsBasePriceListData) {
	o.Data = &v
}

func (o POSTMarkets201ResponseDataRelationshipsBasePriceList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTMarkets201ResponseDataRelationshipsBasePriceList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTMarkets201ResponseDataRelationshipsBasePriceList struct {
	value *POSTMarkets201ResponseDataRelationshipsBasePriceList
	isSet bool
}

func (v NullablePOSTMarkets201ResponseDataRelationshipsBasePriceList) Get() *POSTMarkets201ResponseDataRelationshipsBasePriceList {
	return v.value
}

func (v *NullablePOSTMarkets201ResponseDataRelationshipsBasePriceList) Set(val *POSTMarkets201ResponseDataRelationshipsBasePriceList) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTMarkets201ResponseDataRelationshipsBasePriceList) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTMarkets201ResponseDataRelationshipsBasePriceList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTMarkets201ResponseDataRelationshipsBasePriceList(val *POSTMarkets201ResponseDataRelationshipsBasePriceList) *NullablePOSTMarkets201ResponseDataRelationshipsBasePriceList {
	return &NullablePOSTMarkets201ResponseDataRelationshipsBasePriceList{value: val, isSet: true}
}

func (v NullablePOSTMarkets201ResponseDataRelationshipsBasePriceList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTMarkets201ResponseDataRelationshipsBasePriceList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
