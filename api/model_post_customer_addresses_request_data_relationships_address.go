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

// checks if the POSTCustomerAddressesRequestDataRelationshipsAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTCustomerAddressesRequestDataRelationshipsAddress{}

// POSTCustomerAddressesRequestDataRelationshipsAddress struct for POSTCustomerAddressesRequestDataRelationshipsAddress
type POSTCustomerAddressesRequestDataRelationshipsAddress struct {
	Data POSTCustomerAddressesRequestDataRelationshipsAddressData `json:"data"`
}

// NewPOSTCustomerAddressesRequestDataRelationshipsAddress instantiates a new POSTCustomerAddressesRequestDataRelationshipsAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTCustomerAddressesRequestDataRelationshipsAddress(data POSTCustomerAddressesRequestDataRelationshipsAddressData) *POSTCustomerAddressesRequestDataRelationshipsAddress {
	this := POSTCustomerAddressesRequestDataRelationshipsAddress{}
	this.Data = data
	return &this
}

// NewPOSTCustomerAddressesRequestDataRelationshipsAddressWithDefaults instantiates a new POSTCustomerAddressesRequestDataRelationshipsAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTCustomerAddressesRequestDataRelationshipsAddressWithDefaults() *POSTCustomerAddressesRequestDataRelationshipsAddress {
	this := POSTCustomerAddressesRequestDataRelationshipsAddress{}
	return &this
}

// GetData returns the Data field value
func (o *POSTCustomerAddressesRequestDataRelationshipsAddress) GetData() POSTCustomerAddressesRequestDataRelationshipsAddressData {
	if o == nil {
		var ret POSTCustomerAddressesRequestDataRelationshipsAddressData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *POSTCustomerAddressesRequestDataRelationshipsAddress) GetDataOk() (*POSTCustomerAddressesRequestDataRelationshipsAddressData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *POSTCustomerAddressesRequestDataRelationshipsAddress) SetData(v POSTCustomerAddressesRequestDataRelationshipsAddressData) {
	o.Data = v
}

func (o POSTCustomerAddressesRequestDataRelationshipsAddress) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTCustomerAddressesRequestDataRelationshipsAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullablePOSTCustomerAddressesRequestDataRelationshipsAddress struct {
	value *POSTCustomerAddressesRequestDataRelationshipsAddress
	isSet bool
}

func (v NullablePOSTCustomerAddressesRequestDataRelationshipsAddress) Get() *POSTCustomerAddressesRequestDataRelationshipsAddress {
	return v.value
}

func (v *NullablePOSTCustomerAddressesRequestDataRelationshipsAddress) Set(val *POSTCustomerAddressesRequestDataRelationshipsAddress) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTCustomerAddressesRequestDataRelationshipsAddress) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTCustomerAddressesRequestDataRelationshipsAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTCustomerAddressesRequestDataRelationshipsAddress(val *POSTCustomerAddressesRequestDataRelationshipsAddress) *NullablePOSTCustomerAddressesRequestDataRelationshipsAddress {
	return &NullablePOSTCustomerAddressesRequestDataRelationshipsAddress{value: val, isSet: true}
}

func (v NullablePOSTCustomerAddressesRequestDataRelationshipsAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTCustomerAddressesRequestDataRelationshipsAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
