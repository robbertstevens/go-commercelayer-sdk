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

// checks if the POSTPrices201ResponseDataRelationshipsJwtCustomer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTPrices201ResponseDataRelationshipsJwtCustomer{}

// POSTPrices201ResponseDataRelationshipsJwtCustomer struct for POSTPrices201ResponseDataRelationshipsJwtCustomer
type POSTPrices201ResponseDataRelationshipsJwtCustomer struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks `json:"links,omitempty"`
	Data  *POSTPrices201ResponseDataRelationshipsJwtCustomerData  `json:"data,omitempty"`
}

// NewPOSTPrices201ResponseDataRelationshipsJwtCustomer instantiates a new POSTPrices201ResponseDataRelationshipsJwtCustomer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTPrices201ResponseDataRelationshipsJwtCustomer() *POSTPrices201ResponseDataRelationshipsJwtCustomer {
	this := POSTPrices201ResponseDataRelationshipsJwtCustomer{}
	return &this
}

// NewPOSTPrices201ResponseDataRelationshipsJwtCustomerWithDefaults instantiates a new POSTPrices201ResponseDataRelationshipsJwtCustomer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTPrices201ResponseDataRelationshipsJwtCustomerWithDefaults() *POSTPrices201ResponseDataRelationshipsJwtCustomer {
	this := POSTPrices201ResponseDataRelationshipsJwtCustomer{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTPrices201ResponseDataRelationshipsJwtCustomer) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTPrices201ResponseDataRelationshipsJwtCustomer) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTPrices201ResponseDataRelationshipsJwtCustomer) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTPrices201ResponseDataRelationshipsJwtCustomer) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTPrices201ResponseDataRelationshipsJwtCustomer) GetData() POSTPrices201ResponseDataRelationshipsJwtCustomerData {
	if o == nil || IsNil(o.Data) {
		var ret POSTPrices201ResponseDataRelationshipsJwtCustomerData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTPrices201ResponseDataRelationshipsJwtCustomer) GetDataOk() (*POSTPrices201ResponseDataRelationshipsJwtCustomerData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTPrices201ResponseDataRelationshipsJwtCustomer) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTPrices201ResponseDataRelationshipsJwtCustomerData and assigns it to the Data field.
func (o *POSTPrices201ResponseDataRelationshipsJwtCustomer) SetData(v POSTPrices201ResponseDataRelationshipsJwtCustomerData) {
	o.Data = &v
}

func (o POSTPrices201ResponseDataRelationshipsJwtCustomer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTPrices201ResponseDataRelationshipsJwtCustomer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTPrices201ResponseDataRelationshipsJwtCustomer struct {
	value *POSTPrices201ResponseDataRelationshipsJwtCustomer
	isSet bool
}

func (v NullablePOSTPrices201ResponseDataRelationshipsJwtCustomer) Get() *POSTPrices201ResponseDataRelationshipsJwtCustomer {
	return v.value
}

func (v *NullablePOSTPrices201ResponseDataRelationshipsJwtCustomer) Set(val *POSTPrices201ResponseDataRelationshipsJwtCustomer) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTPrices201ResponseDataRelationshipsJwtCustomer) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTPrices201ResponseDataRelationshipsJwtCustomer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTPrices201ResponseDataRelationshipsJwtCustomer(val *POSTPrices201ResponseDataRelationshipsJwtCustomer) *NullablePOSTPrices201ResponseDataRelationshipsJwtCustomer {
	return &NullablePOSTPrices201ResponseDataRelationshipsJwtCustomer{value: val, isSet: true}
}

func (v NullablePOSTPrices201ResponseDataRelationshipsJwtCustomer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTPrices201ResponseDataRelationshipsJwtCustomer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
