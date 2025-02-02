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

// checks if the POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSource{}

// POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSource struct for POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSource
type POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSource struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks                  `json:"links,omitempty"`
	Data  *POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSourceData `json:"data,omitempty"`
}

// NewPOSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSource instantiates a new POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSource() *POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSource {
	this := POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSource{}
	return &this
}

// NewPOSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSourceWithDefaults instantiates a new POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSourceWithDefaults() *POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSource {
	this := POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSource{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSource) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSource) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSource) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSource) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSource) GetData() POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSourceData {
	if o == nil || IsNil(o.Data) {
		var ret POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSourceData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSource) GetDataOk() (*POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSourceData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSource) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSourceData and assigns it to the Data field.
func (o *POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSource) SetData(v POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSourceData) {
	o.Data = &v
}

func (o POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSource) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSource struct {
	value *POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSource
	isSet bool
}

func (v NullablePOSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSource) Get() *POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSource {
	return v.value
}

func (v *NullablePOSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSource) Set(val *POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSource) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSource) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSource(val *POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSource) *NullablePOSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSource {
	return &NullablePOSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSource{value: val, isSet: true}
}

func (v NullablePOSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
