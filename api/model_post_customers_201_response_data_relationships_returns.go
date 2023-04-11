/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 4.1.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the POSTCustomers201ResponseDataRelationshipsReturns type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTCustomers201ResponseDataRelationshipsReturns{}

// POSTCustomers201ResponseDataRelationshipsReturns struct for POSTCustomers201ResponseDataRelationshipsReturns
type POSTCustomers201ResponseDataRelationshipsReturns struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks `json:"links,omitempty"`
	Data  *POSTCustomers201ResponseDataRelationshipsReturnsData   `json:"data,omitempty"`
}

// NewPOSTCustomers201ResponseDataRelationshipsReturns instantiates a new POSTCustomers201ResponseDataRelationshipsReturns object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTCustomers201ResponseDataRelationshipsReturns() *POSTCustomers201ResponseDataRelationshipsReturns {
	this := POSTCustomers201ResponseDataRelationshipsReturns{}
	return &this
}

// NewPOSTCustomers201ResponseDataRelationshipsReturnsWithDefaults instantiates a new POSTCustomers201ResponseDataRelationshipsReturns object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTCustomers201ResponseDataRelationshipsReturnsWithDefaults() *POSTCustomers201ResponseDataRelationshipsReturns {
	this := POSTCustomers201ResponseDataRelationshipsReturns{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTCustomers201ResponseDataRelationshipsReturns) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCustomers201ResponseDataRelationshipsReturns) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTCustomers201ResponseDataRelationshipsReturns) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTCustomers201ResponseDataRelationshipsReturns) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTCustomers201ResponseDataRelationshipsReturns) GetData() POSTCustomers201ResponseDataRelationshipsReturnsData {
	if o == nil || IsNil(o.Data) {
		var ret POSTCustomers201ResponseDataRelationshipsReturnsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCustomers201ResponseDataRelationshipsReturns) GetDataOk() (*POSTCustomers201ResponseDataRelationshipsReturnsData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTCustomers201ResponseDataRelationshipsReturns) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTCustomers201ResponseDataRelationshipsReturnsData and assigns it to the Data field.
func (o *POSTCustomers201ResponseDataRelationshipsReturns) SetData(v POSTCustomers201ResponseDataRelationshipsReturnsData) {
	o.Data = &v
}

func (o POSTCustomers201ResponseDataRelationshipsReturns) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTCustomers201ResponseDataRelationshipsReturns) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTCustomers201ResponseDataRelationshipsReturns struct {
	value *POSTCustomers201ResponseDataRelationshipsReturns
	isSet bool
}

func (v NullablePOSTCustomers201ResponseDataRelationshipsReturns) Get() *POSTCustomers201ResponseDataRelationshipsReturns {
	return v.value
}

func (v *NullablePOSTCustomers201ResponseDataRelationshipsReturns) Set(val *POSTCustomers201ResponseDataRelationshipsReturns) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTCustomers201ResponseDataRelationshipsReturns) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTCustomers201ResponseDataRelationshipsReturns) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTCustomers201ResponseDataRelationshipsReturns(val *POSTCustomers201ResponseDataRelationshipsReturns) *NullablePOSTCustomers201ResponseDataRelationshipsReturns {
	return &NullablePOSTCustomers201ResponseDataRelationshipsReturns{value: val, isSet: true}
}

func (v NullablePOSTCustomers201ResponseDataRelationshipsReturns) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTCustomers201ResponseDataRelationshipsReturns) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
