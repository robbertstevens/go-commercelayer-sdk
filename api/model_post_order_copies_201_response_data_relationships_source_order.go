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

// checks if the POSTOrderCopies201ResponseDataRelationshipsSourceOrder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTOrderCopies201ResponseDataRelationshipsSourceOrder{}

// POSTOrderCopies201ResponseDataRelationshipsSourceOrder struct for POSTOrderCopies201ResponseDataRelationshipsSourceOrder
type POSTOrderCopies201ResponseDataRelationshipsSourceOrder struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks     `json:"links,omitempty"`
	Data  *POSTOrderCopies201ResponseDataRelationshipsSourceOrderData `json:"data,omitempty"`
}

// NewPOSTOrderCopies201ResponseDataRelationshipsSourceOrder instantiates a new POSTOrderCopies201ResponseDataRelationshipsSourceOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTOrderCopies201ResponseDataRelationshipsSourceOrder() *POSTOrderCopies201ResponseDataRelationshipsSourceOrder {
	this := POSTOrderCopies201ResponseDataRelationshipsSourceOrder{}
	return &this
}

// NewPOSTOrderCopies201ResponseDataRelationshipsSourceOrderWithDefaults instantiates a new POSTOrderCopies201ResponseDataRelationshipsSourceOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTOrderCopies201ResponseDataRelationshipsSourceOrderWithDefaults() *POSTOrderCopies201ResponseDataRelationshipsSourceOrder {
	this := POSTOrderCopies201ResponseDataRelationshipsSourceOrder{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTOrderCopies201ResponseDataRelationshipsSourceOrder) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrderCopies201ResponseDataRelationshipsSourceOrder) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTOrderCopies201ResponseDataRelationshipsSourceOrder) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTOrderCopies201ResponseDataRelationshipsSourceOrder) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTOrderCopies201ResponseDataRelationshipsSourceOrder) GetData() POSTOrderCopies201ResponseDataRelationshipsSourceOrderData {
	if o == nil || IsNil(o.Data) {
		var ret POSTOrderCopies201ResponseDataRelationshipsSourceOrderData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrderCopies201ResponseDataRelationshipsSourceOrder) GetDataOk() (*POSTOrderCopies201ResponseDataRelationshipsSourceOrderData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTOrderCopies201ResponseDataRelationshipsSourceOrder) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTOrderCopies201ResponseDataRelationshipsSourceOrderData and assigns it to the Data field.
func (o *POSTOrderCopies201ResponseDataRelationshipsSourceOrder) SetData(v POSTOrderCopies201ResponseDataRelationshipsSourceOrderData) {
	o.Data = &v
}

func (o POSTOrderCopies201ResponseDataRelationshipsSourceOrder) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTOrderCopies201ResponseDataRelationshipsSourceOrder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTOrderCopies201ResponseDataRelationshipsSourceOrder struct {
	value *POSTOrderCopies201ResponseDataRelationshipsSourceOrder
	isSet bool
}

func (v NullablePOSTOrderCopies201ResponseDataRelationshipsSourceOrder) Get() *POSTOrderCopies201ResponseDataRelationshipsSourceOrder {
	return v.value
}

func (v *NullablePOSTOrderCopies201ResponseDataRelationshipsSourceOrder) Set(val *POSTOrderCopies201ResponseDataRelationshipsSourceOrder) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTOrderCopies201ResponseDataRelationshipsSourceOrder) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTOrderCopies201ResponseDataRelationshipsSourceOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTOrderCopies201ResponseDataRelationshipsSourceOrder(val *POSTOrderCopies201ResponseDataRelationshipsSourceOrder) *NullablePOSTOrderCopies201ResponseDataRelationshipsSourceOrder {
	return &NullablePOSTOrderCopies201ResponseDataRelationshipsSourceOrder{value: val, isSet: true}
}

func (v NullablePOSTOrderCopies201ResponseDataRelationshipsSourceOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTOrderCopies201ResponseDataRelationshipsSourceOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
