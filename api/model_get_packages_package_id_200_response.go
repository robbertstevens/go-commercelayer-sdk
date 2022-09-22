/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETPackagesPackageId200Response struct for GETPackagesPackageId200Response
type GETPackagesPackageId200Response struct {
	Data *GETPackages200ResponseDataInner `json:"data,omitempty"`
}

// NewGETPackagesPackageId200Response instantiates a new GETPackagesPackageId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETPackagesPackageId200Response() *GETPackagesPackageId200Response {
	this := GETPackagesPackageId200Response{}
	return &this
}

// NewGETPackagesPackageId200ResponseWithDefaults instantiates a new GETPackagesPackageId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETPackagesPackageId200ResponseWithDefaults() *GETPackagesPackageId200Response {
	this := GETPackagesPackageId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETPackagesPackageId200Response) GetData() GETPackages200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret GETPackages200ResponseDataInner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETPackagesPackageId200Response) GetDataOk() (*GETPackages200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETPackagesPackageId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETPackages200ResponseDataInner and assigns it to the Data field.
func (o *GETPackagesPackageId200Response) SetData(v GETPackages200ResponseDataInner) {
	o.Data = &v
}

func (o GETPackagesPackageId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETPackagesPackageId200Response struct {
	value *GETPackagesPackageId200Response
	isSet bool
}

func (v NullableGETPackagesPackageId200Response) Get() *GETPackagesPackageId200Response {
	return v.value
}

func (v *NullableGETPackagesPackageId200Response) Set(val *GETPackagesPackageId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETPackagesPackageId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETPackagesPackageId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETPackagesPackageId200Response(val *GETPackagesPackageId200Response) *NullableGETPackagesPackageId200Response {
	return &NullableGETPackagesPackageId200Response{value: val, isSet: true}
}

func (v NullableGETPackagesPackageId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETPackagesPackageId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
