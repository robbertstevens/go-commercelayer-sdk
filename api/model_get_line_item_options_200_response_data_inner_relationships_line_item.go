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

// GETLineItemOptions200ResponseDataInnerRelationshipsLineItem struct for GETLineItemOptions200ResponseDataInnerRelationshipsLineItem
type GETLineItemOptions200ResponseDataInnerRelationshipsLineItem struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks      `json:"links,omitempty"`
	Data  *GETLineItemOptions200ResponseDataInnerRelationshipsLineItemData `json:"data,omitempty"`
}

// NewGETLineItemOptions200ResponseDataInnerRelationshipsLineItem instantiates a new GETLineItemOptions200ResponseDataInnerRelationshipsLineItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETLineItemOptions200ResponseDataInnerRelationshipsLineItem() *GETLineItemOptions200ResponseDataInnerRelationshipsLineItem {
	this := GETLineItemOptions200ResponseDataInnerRelationshipsLineItem{}
	return &this
}

// NewGETLineItemOptions200ResponseDataInnerRelationshipsLineItemWithDefaults instantiates a new GETLineItemOptions200ResponseDataInnerRelationshipsLineItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETLineItemOptions200ResponseDataInnerRelationshipsLineItemWithDefaults() *GETLineItemOptions200ResponseDataInnerRelationshipsLineItem {
	this := GETLineItemOptions200ResponseDataInnerRelationshipsLineItem{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETLineItemOptions200ResponseDataInnerRelationshipsLineItem) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETLineItemOptions200ResponseDataInnerRelationshipsLineItem) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETLineItemOptions200ResponseDataInnerRelationshipsLineItem) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETLineItemOptions200ResponseDataInnerRelationshipsLineItem) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETLineItemOptions200ResponseDataInnerRelationshipsLineItem) GetData() GETLineItemOptions200ResponseDataInnerRelationshipsLineItemData {
	if o == nil || o.Data == nil {
		var ret GETLineItemOptions200ResponseDataInnerRelationshipsLineItemData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETLineItemOptions200ResponseDataInnerRelationshipsLineItem) GetDataOk() (*GETLineItemOptions200ResponseDataInnerRelationshipsLineItemData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETLineItemOptions200ResponseDataInnerRelationshipsLineItem) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETLineItemOptions200ResponseDataInnerRelationshipsLineItemData and assigns it to the Data field.
func (o *GETLineItemOptions200ResponseDataInnerRelationshipsLineItem) SetData(v GETLineItemOptions200ResponseDataInnerRelationshipsLineItemData) {
	o.Data = &v
}

func (o GETLineItemOptions200ResponseDataInnerRelationshipsLineItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETLineItemOptions200ResponseDataInnerRelationshipsLineItem struct {
	value *GETLineItemOptions200ResponseDataInnerRelationshipsLineItem
	isSet bool
}

func (v NullableGETLineItemOptions200ResponseDataInnerRelationshipsLineItem) Get() *GETLineItemOptions200ResponseDataInnerRelationshipsLineItem {
	return v.value
}

func (v *NullableGETLineItemOptions200ResponseDataInnerRelationshipsLineItem) Set(val *GETLineItemOptions200ResponseDataInnerRelationshipsLineItem) {
	v.value = val
	v.isSet = true
}

func (v NullableGETLineItemOptions200ResponseDataInnerRelationshipsLineItem) IsSet() bool {
	return v.isSet
}

func (v *NullableGETLineItemOptions200ResponseDataInnerRelationshipsLineItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETLineItemOptions200ResponseDataInnerRelationshipsLineItem(val *GETLineItemOptions200ResponseDataInnerRelationshipsLineItem) *NullableGETLineItemOptions200ResponseDataInnerRelationshipsLineItem {
	return &NullableGETLineItemOptions200ResponseDataInnerRelationshipsLineItem{value: val, isSet: true}
}

func (v NullableGETLineItemOptions200ResponseDataInnerRelationshipsLineItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETLineItemOptions200ResponseDataInnerRelationshipsLineItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
