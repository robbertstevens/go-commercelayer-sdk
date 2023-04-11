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

// checks if the POSTCustomers201ResponseDataRelationshipsOrders type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTCustomers201ResponseDataRelationshipsOrders{}

// POSTCustomers201ResponseDataRelationshipsOrders struct for POSTCustomers201ResponseDataRelationshipsOrders
type POSTCustomers201ResponseDataRelationshipsOrders struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks `json:"links,omitempty"`
	Data  *POSTCustomers201ResponseDataRelationshipsOrdersData    `json:"data,omitempty"`
}

// NewPOSTCustomers201ResponseDataRelationshipsOrders instantiates a new POSTCustomers201ResponseDataRelationshipsOrders object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTCustomers201ResponseDataRelationshipsOrders() *POSTCustomers201ResponseDataRelationshipsOrders {
	this := POSTCustomers201ResponseDataRelationshipsOrders{}
	return &this
}

// NewPOSTCustomers201ResponseDataRelationshipsOrdersWithDefaults instantiates a new POSTCustomers201ResponseDataRelationshipsOrders object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTCustomers201ResponseDataRelationshipsOrdersWithDefaults() *POSTCustomers201ResponseDataRelationshipsOrders {
	this := POSTCustomers201ResponseDataRelationshipsOrders{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTCustomers201ResponseDataRelationshipsOrders) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCustomers201ResponseDataRelationshipsOrders) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTCustomers201ResponseDataRelationshipsOrders) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTCustomers201ResponseDataRelationshipsOrders) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTCustomers201ResponseDataRelationshipsOrders) GetData() POSTCustomers201ResponseDataRelationshipsOrdersData {
	if o == nil || IsNil(o.Data) {
		var ret POSTCustomers201ResponseDataRelationshipsOrdersData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCustomers201ResponseDataRelationshipsOrders) GetDataOk() (*POSTCustomers201ResponseDataRelationshipsOrdersData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTCustomers201ResponseDataRelationshipsOrders) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTCustomers201ResponseDataRelationshipsOrdersData and assigns it to the Data field.
func (o *POSTCustomers201ResponseDataRelationshipsOrders) SetData(v POSTCustomers201ResponseDataRelationshipsOrdersData) {
	o.Data = &v
}

func (o POSTCustomers201ResponseDataRelationshipsOrders) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTCustomers201ResponseDataRelationshipsOrders) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTCustomers201ResponseDataRelationshipsOrders struct {
	value *POSTCustomers201ResponseDataRelationshipsOrders
	isSet bool
}

func (v NullablePOSTCustomers201ResponseDataRelationshipsOrders) Get() *POSTCustomers201ResponseDataRelationshipsOrders {
	return v.value
}

func (v *NullablePOSTCustomers201ResponseDataRelationshipsOrders) Set(val *POSTCustomers201ResponseDataRelationshipsOrders) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTCustomers201ResponseDataRelationshipsOrders) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTCustomers201ResponseDataRelationshipsOrders) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTCustomers201ResponseDataRelationshipsOrders(val *POSTCustomers201ResponseDataRelationshipsOrders) *NullablePOSTCustomers201ResponseDataRelationshipsOrders {
	return &NullablePOSTCustomers201ResponseDataRelationshipsOrders{value: val, isSet: true}
}

func (v NullablePOSTCustomers201ResponseDataRelationshipsOrders) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTCustomers201ResponseDataRelationshipsOrders) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
