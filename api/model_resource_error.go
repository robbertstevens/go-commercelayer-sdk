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

// checks if the ResourceError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceError{}

// ResourceError struct for ResourceError
type ResourceError struct {
	Data *ResourceErrorData `json:"data,omitempty"`
}

// NewResourceError instantiates a new ResourceError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceError() *ResourceError {
	this := ResourceError{}
	return &this
}

// NewResourceErrorWithDefaults instantiates a new ResourceError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceErrorWithDefaults() *ResourceError {
	this := ResourceError{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ResourceError) GetData() ResourceErrorData {
	if o == nil || IsNil(o.Data) {
		var ret ResourceErrorData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceError) GetDataOk() (*ResourceErrorData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ResourceError) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ResourceErrorData and assigns it to the Data field.
func (o *ResourceError) SetData(v ResourceErrorData) {
	o.Data = &v
}

func (o ResourceError) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableResourceError struct {
	value *ResourceError
	isSet bool
}

func (v NullableResourceError) Get() *ResourceError {
	return v.value
}

func (v *NullableResourceError) Set(val *ResourceError) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceError) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceError(val *ResourceError) *NullableResourceError {
	return &NullableResourceError{value: val, isSet: true}
}

func (v NullableResourceError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
