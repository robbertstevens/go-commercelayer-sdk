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

// checks if the POSTParcelLineItems201ResponseDataRelationshipsStockLineItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTParcelLineItems201ResponseDataRelationshipsStockLineItem{}

// POSTParcelLineItems201ResponseDataRelationshipsStockLineItem struct for POSTParcelLineItems201ResponseDataRelationshipsStockLineItem
type POSTParcelLineItems201ResponseDataRelationshipsStockLineItem struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks           `json:"links,omitempty"`
	Data  *POSTParcelLineItems201ResponseDataRelationshipsStockLineItemData `json:"data,omitempty"`
}

// NewPOSTParcelLineItems201ResponseDataRelationshipsStockLineItem instantiates a new POSTParcelLineItems201ResponseDataRelationshipsStockLineItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTParcelLineItems201ResponseDataRelationshipsStockLineItem() *POSTParcelLineItems201ResponseDataRelationshipsStockLineItem {
	this := POSTParcelLineItems201ResponseDataRelationshipsStockLineItem{}
	return &this
}

// NewPOSTParcelLineItems201ResponseDataRelationshipsStockLineItemWithDefaults instantiates a new POSTParcelLineItems201ResponseDataRelationshipsStockLineItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTParcelLineItems201ResponseDataRelationshipsStockLineItemWithDefaults() *POSTParcelLineItems201ResponseDataRelationshipsStockLineItem {
	this := POSTParcelLineItems201ResponseDataRelationshipsStockLineItem{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTParcelLineItems201ResponseDataRelationshipsStockLineItem) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTParcelLineItems201ResponseDataRelationshipsStockLineItem) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTParcelLineItems201ResponseDataRelationshipsStockLineItem) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTParcelLineItems201ResponseDataRelationshipsStockLineItem) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTParcelLineItems201ResponseDataRelationshipsStockLineItem) GetData() POSTParcelLineItems201ResponseDataRelationshipsStockLineItemData {
	if o == nil || IsNil(o.Data) {
		var ret POSTParcelLineItems201ResponseDataRelationshipsStockLineItemData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTParcelLineItems201ResponseDataRelationshipsStockLineItem) GetDataOk() (*POSTParcelLineItems201ResponseDataRelationshipsStockLineItemData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTParcelLineItems201ResponseDataRelationshipsStockLineItem) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTParcelLineItems201ResponseDataRelationshipsStockLineItemData and assigns it to the Data field.
func (o *POSTParcelLineItems201ResponseDataRelationshipsStockLineItem) SetData(v POSTParcelLineItems201ResponseDataRelationshipsStockLineItemData) {
	o.Data = &v
}

func (o POSTParcelLineItems201ResponseDataRelationshipsStockLineItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTParcelLineItems201ResponseDataRelationshipsStockLineItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTParcelLineItems201ResponseDataRelationshipsStockLineItem struct {
	value *POSTParcelLineItems201ResponseDataRelationshipsStockLineItem
	isSet bool
}

func (v NullablePOSTParcelLineItems201ResponseDataRelationshipsStockLineItem) Get() *POSTParcelLineItems201ResponseDataRelationshipsStockLineItem {
	return v.value
}

func (v *NullablePOSTParcelLineItems201ResponseDataRelationshipsStockLineItem) Set(val *POSTParcelLineItems201ResponseDataRelationshipsStockLineItem) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTParcelLineItems201ResponseDataRelationshipsStockLineItem) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTParcelLineItems201ResponseDataRelationshipsStockLineItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTParcelLineItems201ResponseDataRelationshipsStockLineItem(val *POSTParcelLineItems201ResponseDataRelationshipsStockLineItem) *NullablePOSTParcelLineItems201ResponseDataRelationshipsStockLineItem {
	return &NullablePOSTParcelLineItems201ResponseDataRelationshipsStockLineItem{value: val, isSet: true}
}

func (v NullablePOSTParcelLineItems201ResponseDataRelationshipsStockLineItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTParcelLineItems201ResponseDataRelationshipsStockLineItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
