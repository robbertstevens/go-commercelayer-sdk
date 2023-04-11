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

// checks if the PackageCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PackageCreate{}

// PackageCreate struct for PackageCreate
type PackageCreate struct {
	Data POSTPackagesRequestData `json:"data"`
}

// NewPackageCreate instantiates a new PackageCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackageCreate(data POSTPackagesRequestData) *PackageCreate {
	this := PackageCreate{}
	this.Data = data
	return &this
}

// NewPackageCreateWithDefaults instantiates a new PackageCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageCreateWithDefaults() *PackageCreate {
	this := PackageCreate{}
	return &this
}

// GetData returns the Data field value
func (o *PackageCreate) GetData() POSTPackagesRequestData {
	if o == nil {
		var ret POSTPackagesRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PackageCreate) GetDataOk() (*POSTPackagesRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PackageCreate) SetData(v POSTPackagesRequestData) {
	o.Data = v
}

func (o PackageCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PackageCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullablePackageCreate struct {
	value *PackageCreate
	isSet bool
}

func (v NullablePackageCreate) Get() *PackageCreate {
	return v.value
}

func (v *NullablePackageCreate) Set(val *PackageCreate) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageCreate) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageCreate(val *PackageCreate) *NullablePackageCreate {
	return &NullablePackageCreate{value: val, isSet: true}
}

func (v NullablePackageCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackageCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
