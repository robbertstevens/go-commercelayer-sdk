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

// checks if the AddressCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddressCreate{}

// AddressCreate struct for AddressCreate
type AddressCreate struct {
	Data POSTAddressesRequestData `json:"data"`
}

// NewAddressCreate instantiates a new AddressCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressCreate(data POSTAddressesRequestData) *AddressCreate {
	this := AddressCreate{}
	this.Data = data
	return &this
}

// NewAddressCreateWithDefaults instantiates a new AddressCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressCreateWithDefaults() *AddressCreate {
	this := AddressCreate{}
	return &this
}

// GetData returns the Data field value
func (o *AddressCreate) GetData() POSTAddressesRequestData {
	if o == nil {
		var ret POSTAddressesRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AddressCreate) GetDataOk() (*POSTAddressesRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AddressCreate) SetData(v POSTAddressesRequestData) {
	o.Data = v
}

func (o AddressCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableAddressCreate struct {
	value *AddressCreate
	isSet bool
}

func (v NullableAddressCreate) Get() *AddressCreate {
	return v.value
}

func (v *NullableAddressCreate) Set(val *AddressCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressCreate(val *AddressCreate) *NullableAddressCreate {
	return &NullableAddressCreate{value: val, isSet: true}
}

func (v NullableAddressCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
