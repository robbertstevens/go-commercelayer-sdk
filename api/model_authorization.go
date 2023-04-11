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

// checks if the Authorization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Authorization{}

// Authorization struct for Authorization
type Authorization struct {
	Data *AuthorizationData `json:"data,omitempty"`
}

// NewAuthorization instantiates a new Authorization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorization() *Authorization {
	this := Authorization{}
	return &this
}

// NewAuthorizationWithDefaults instantiates a new Authorization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationWithDefaults() *Authorization {
	this := Authorization{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Authorization) GetData() AuthorizationData {
	if o == nil || IsNil(o.Data) {
		var ret AuthorizationData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Authorization) GetDataOk() (*AuthorizationData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Authorization) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AuthorizationData and assigns it to the Data field.
func (o *Authorization) SetData(v AuthorizationData) {
	o.Data = &v
}

func (o Authorization) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Authorization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAuthorization struct {
	value *Authorization
	isSet bool
}

func (v NullableAuthorization) Get() *Authorization {
	return v.value
}

func (v *NullableAuthorization) Set(val *Authorization) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorization) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorization(val *Authorization) *NullableAuthorization {
	return &NullableAuthorization{value: val, isSet: true}
}

func (v NullableAuthorization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
