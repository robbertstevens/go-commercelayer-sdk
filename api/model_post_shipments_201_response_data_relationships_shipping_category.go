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

// checks if the POSTShipments201ResponseDataRelationshipsShippingCategory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTShipments201ResponseDataRelationshipsShippingCategory{}

// POSTShipments201ResponseDataRelationshipsShippingCategory struct for POSTShipments201ResponseDataRelationshipsShippingCategory
type POSTShipments201ResponseDataRelationshipsShippingCategory struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks        `json:"links,omitempty"`
	Data  *POSTShipments201ResponseDataRelationshipsShippingCategoryData `json:"data,omitempty"`
}

// NewPOSTShipments201ResponseDataRelationshipsShippingCategory instantiates a new POSTShipments201ResponseDataRelationshipsShippingCategory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTShipments201ResponseDataRelationshipsShippingCategory() *POSTShipments201ResponseDataRelationshipsShippingCategory {
	this := POSTShipments201ResponseDataRelationshipsShippingCategory{}
	return &this
}

// NewPOSTShipments201ResponseDataRelationshipsShippingCategoryWithDefaults instantiates a new POSTShipments201ResponseDataRelationshipsShippingCategory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTShipments201ResponseDataRelationshipsShippingCategoryWithDefaults() *POSTShipments201ResponseDataRelationshipsShippingCategory {
	this := POSTShipments201ResponseDataRelationshipsShippingCategory{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTShipments201ResponseDataRelationshipsShippingCategory) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTShipments201ResponseDataRelationshipsShippingCategory) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTShipments201ResponseDataRelationshipsShippingCategory) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTShipments201ResponseDataRelationshipsShippingCategory) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTShipments201ResponseDataRelationshipsShippingCategory) GetData() POSTShipments201ResponseDataRelationshipsShippingCategoryData {
	if o == nil || IsNil(o.Data) {
		var ret POSTShipments201ResponseDataRelationshipsShippingCategoryData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTShipments201ResponseDataRelationshipsShippingCategory) GetDataOk() (*POSTShipments201ResponseDataRelationshipsShippingCategoryData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTShipments201ResponseDataRelationshipsShippingCategory) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTShipments201ResponseDataRelationshipsShippingCategoryData and assigns it to the Data field.
func (o *POSTShipments201ResponseDataRelationshipsShippingCategory) SetData(v POSTShipments201ResponseDataRelationshipsShippingCategoryData) {
	o.Data = &v
}

func (o POSTShipments201ResponseDataRelationshipsShippingCategory) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTShipments201ResponseDataRelationshipsShippingCategory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTShipments201ResponseDataRelationshipsShippingCategory struct {
	value *POSTShipments201ResponseDataRelationshipsShippingCategory
	isSet bool
}

func (v NullablePOSTShipments201ResponseDataRelationshipsShippingCategory) Get() *POSTShipments201ResponseDataRelationshipsShippingCategory {
	return v.value
}

func (v *NullablePOSTShipments201ResponseDataRelationshipsShippingCategory) Set(val *POSTShipments201ResponseDataRelationshipsShippingCategory) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTShipments201ResponseDataRelationshipsShippingCategory) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTShipments201ResponseDataRelationshipsShippingCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTShipments201ResponseDataRelationshipsShippingCategory(val *POSTShipments201ResponseDataRelationshipsShippingCategory) *NullablePOSTShipments201ResponseDataRelationshipsShippingCategory {
	return &NullablePOSTShipments201ResponseDataRelationshipsShippingCategory{value: val, isSet: true}
}

func (v NullablePOSTShipments201ResponseDataRelationshipsShippingCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTShipments201ResponseDataRelationshipsShippingCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
