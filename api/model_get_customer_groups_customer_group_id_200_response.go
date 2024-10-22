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

// checks if the GETCustomerGroupsCustomerGroupId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETCustomerGroupsCustomerGroupId200Response{}

// GETCustomerGroupsCustomerGroupId200Response struct for GETCustomerGroupsCustomerGroupId200Response
type GETCustomerGroupsCustomerGroupId200Response struct {
	Data *GETCustomerGroupsCustomerGroupId200ResponseData `json:"data,omitempty"`
}

// NewGETCustomerGroupsCustomerGroupId200Response instantiates a new GETCustomerGroupsCustomerGroupId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCustomerGroupsCustomerGroupId200Response() *GETCustomerGroupsCustomerGroupId200Response {
	this := GETCustomerGroupsCustomerGroupId200Response{}
	return &this
}

// NewGETCustomerGroupsCustomerGroupId200ResponseWithDefaults instantiates a new GETCustomerGroupsCustomerGroupId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCustomerGroupsCustomerGroupId200ResponseWithDefaults() *GETCustomerGroupsCustomerGroupId200Response {
	this := GETCustomerGroupsCustomerGroupId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETCustomerGroupsCustomerGroupId200Response) GetData() GETCustomerGroupsCustomerGroupId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret GETCustomerGroupsCustomerGroupId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCustomerGroupsCustomerGroupId200Response) GetDataOk() (*GETCustomerGroupsCustomerGroupId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETCustomerGroupsCustomerGroupId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GETCustomerGroupsCustomerGroupId200ResponseData and assigns it to the Data field.
func (o *GETCustomerGroupsCustomerGroupId200Response) SetData(v GETCustomerGroupsCustomerGroupId200ResponseData) {
	o.Data = &v
}

func (o GETCustomerGroupsCustomerGroupId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETCustomerGroupsCustomerGroupId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGETCustomerGroupsCustomerGroupId200Response struct {
	value *GETCustomerGroupsCustomerGroupId200Response
	isSet bool
}

func (v NullableGETCustomerGroupsCustomerGroupId200Response) Get() *GETCustomerGroupsCustomerGroupId200Response {
	return v.value
}

func (v *NullableGETCustomerGroupsCustomerGroupId200Response) Set(val *GETCustomerGroupsCustomerGroupId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCustomerGroupsCustomerGroupId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCustomerGroupsCustomerGroupId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCustomerGroupsCustomerGroupId200Response(val *GETCustomerGroupsCustomerGroupId200Response) *NullableGETCustomerGroupsCustomerGroupId200Response {
	return &NullableGETCustomerGroupsCustomerGroupId200Response{value: val, isSet: true}
}

func (v NullableGETCustomerGroupsCustomerGroupId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCustomerGroupsCustomerGroupId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
