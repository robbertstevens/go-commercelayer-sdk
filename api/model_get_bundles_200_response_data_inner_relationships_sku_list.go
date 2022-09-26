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

// GETBundles200ResponseDataInnerRelationshipsSkuList struct for GETBundles200ResponseDataInnerRelationshipsSkuList
type GETBundles200ResponseDataInnerRelationshipsSkuList struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks `json:"links,omitempty"`
	Data  *GETBundles200ResponseDataInnerRelationshipsSkuListData     `json:"data,omitempty"`
}

// NewGETBundles200ResponseDataInnerRelationshipsSkuList instantiates a new GETBundles200ResponseDataInnerRelationshipsSkuList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETBundles200ResponseDataInnerRelationshipsSkuList() *GETBundles200ResponseDataInnerRelationshipsSkuList {
	this := GETBundles200ResponseDataInnerRelationshipsSkuList{}
	return &this
}

// NewGETBundles200ResponseDataInnerRelationshipsSkuListWithDefaults instantiates a new GETBundles200ResponseDataInnerRelationshipsSkuList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETBundles200ResponseDataInnerRelationshipsSkuListWithDefaults() *GETBundles200ResponseDataInnerRelationshipsSkuList {
	this := GETBundles200ResponseDataInnerRelationshipsSkuList{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETBundles200ResponseDataInnerRelationshipsSkuList) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETBundles200ResponseDataInnerRelationshipsSkuList) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETBundles200ResponseDataInnerRelationshipsSkuList) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETBundles200ResponseDataInnerRelationshipsSkuList) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETBundles200ResponseDataInnerRelationshipsSkuList) GetData() GETBundles200ResponseDataInnerRelationshipsSkuListData {
	if o == nil || o.Data == nil {
		var ret GETBundles200ResponseDataInnerRelationshipsSkuListData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETBundles200ResponseDataInnerRelationshipsSkuList) GetDataOk() (*GETBundles200ResponseDataInnerRelationshipsSkuListData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETBundles200ResponseDataInnerRelationshipsSkuList) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETBundles200ResponseDataInnerRelationshipsSkuListData and assigns it to the Data field.
func (o *GETBundles200ResponseDataInnerRelationshipsSkuList) SetData(v GETBundles200ResponseDataInnerRelationshipsSkuListData) {
	o.Data = &v
}

func (o GETBundles200ResponseDataInnerRelationshipsSkuList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETBundles200ResponseDataInnerRelationshipsSkuList struct {
	value *GETBundles200ResponseDataInnerRelationshipsSkuList
	isSet bool
}

func (v NullableGETBundles200ResponseDataInnerRelationshipsSkuList) Get() *GETBundles200ResponseDataInnerRelationshipsSkuList {
	return v.value
}

func (v *NullableGETBundles200ResponseDataInnerRelationshipsSkuList) Set(val *GETBundles200ResponseDataInnerRelationshipsSkuList) {
	v.value = val
	v.isSet = true
}

func (v NullableGETBundles200ResponseDataInnerRelationshipsSkuList) IsSet() bool {
	return v.isSet
}

func (v *NullableGETBundles200ResponseDataInnerRelationshipsSkuList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETBundles200ResponseDataInnerRelationshipsSkuList(val *GETBundles200ResponseDataInnerRelationshipsSkuList) *NullableGETBundles200ResponseDataInnerRelationshipsSkuList {
	return &NullableGETBundles200ResponseDataInnerRelationshipsSkuList{value: val, isSet: true}
}

func (v NullableGETBundles200ResponseDataInnerRelationshipsSkuList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETBundles200ResponseDataInnerRelationshipsSkuList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
