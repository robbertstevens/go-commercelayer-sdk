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

// GETSkuLists200ResponseDataInnerRelationshipsSkuListItems struct for GETSkuLists200ResponseDataInnerRelationshipsSkuListItems
type GETSkuLists200ResponseDataInnerRelationshipsSkuListItems struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks   `json:"links,omitempty"`
	Data  *GETSkuLists200ResponseDataInnerRelationshipsSkuListItemsData `json:"data,omitempty"`
}

// NewGETSkuLists200ResponseDataInnerRelationshipsSkuListItems instantiates a new GETSkuLists200ResponseDataInnerRelationshipsSkuListItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETSkuLists200ResponseDataInnerRelationshipsSkuListItems() *GETSkuLists200ResponseDataInnerRelationshipsSkuListItems {
	this := GETSkuLists200ResponseDataInnerRelationshipsSkuListItems{}
	return &this
}

// NewGETSkuLists200ResponseDataInnerRelationshipsSkuListItemsWithDefaults instantiates a new GETSkuLists200ResponseDataInnerRelationshipsSkuListItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETSkuLists200ResponseDataInnerRelationshipsSkuListItemsWithDefaults() *GETSkuLists200ResponseDataInnerRelationshipsSkuListItems {
	this := GETSkuLists200ResponseDataInnerRelationshipsSkuListItems{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETSkuLists200ResponseDataInnerRelationshipsSkuListItems) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkuLists200ResponseDataInnerRelationshipsSkuListItems) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETSkuLists200ResponseDataInnerRelationshipsSkuListItems) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETSkuLists200ResponseDataInnerRelationshipsSkuListItems) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETSkuLists200ResponseDataInnerRelationshipsSkuListItems) GetData() GETSkuLists200ResponseDataInnerRelationshipsSkuListItemsData {
	if o == nil || o.Data == nil {
		var ret GETSkuLists200ResponseDataInnerRelationshipsSkuListItemsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkuLists200ResponseDataInnerRelationshipsSkuListItems) GetDataOk() (*GETSkuLists200ResponseDataInnerRelationshipsSkuListItemsData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETSkuLists200ResponseDataInnerRelationshipsSkuListItems) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETSkuLists200ResponseDataInnerRelationshipsSkuListItemsData and assigns it to the Data field.
func (o *GETSkuLists200ResponseDataInnerRelationshipsSkuListItems) SetData(v GETSkuLists200ResponseDataInnerRelationshipsSkuListItemsData) {
	o.Data = &v
}

func (o GETSkuLists200ResponseDataInnerRelationshipsSkuListItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETSkuLists200ResponseDataInnerRelationshipsSkuListItems struct {
	value *GETSkuLists200ResponseDataInnerRelationshipsSkuListItems
	isSet bool
}

func (v NullableGETSkuLists200ResponseDataInnerRelationshipsSkuListItems) Get() *GETSkuLists200ResponseDataInnerRelationshipsSkuListItems {
	return v.value
}

func (v *NullableGETSkuLists200ResponseDataInnerRelationshipsSkuListItems) Set(val *GETSkuLists200ResponseDataInnerRelationshipsSkuListItems) {
	v.value = val
	v.isSet = true
}

func (v NullableGETSkuLists200ResponseDataInnerRelationshipsSkuListItems) IsSet() bool {
	return v.isSet
}

func (v *NullableGETSkuLists200ResponseDataInnerRelationshipsSkuListItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETSkuLists200ResponseDataInnerRelationshipsSkuListItems(val *GETSkuLists200ResponseDataInnerRelationshipsSkuListItems) *NullableGETSkuLists200ResponseDataInnerRelationshipsSkuListItems {
	return &NullableGETSkuLists200ResponseDataInnerRelationshipsSkuListItems{value: val, isSet: true}
}

func (v NullableGETSkuLists200ResponseDataInnerRelationshipsSkuListItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETSkuLists200ResponseDataInnerRelationshipsSkuListItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
