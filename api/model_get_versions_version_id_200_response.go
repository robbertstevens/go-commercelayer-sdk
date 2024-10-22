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

// checks if the GETVersionsVersionId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETVersionsVersionId200Response{}

// GETVersionsVersionId200Response struct for GETVersionsVersionId200Response
type GETVersionsVersionId200Response struct {
	Data *GETVersionsVersionId200ResponseData `json:"data,omitempty"`
}

// NewGETVersionsVersionId200Response instantiates a new GETVersionsVersionId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETVersionsVersionId200Response() *GETVersionsVersionId200Response {
	this := GETVersionsVersionId200Response{}
	return &this
}

// NewGETVersionsVersionId200ResponseWithDefaults instantiates a new GETVersionsVersionId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETVersionsVersionId200ResponseWithDefaults() *GETVersionsVersionId200Response {
	this := GETVersionsVersionId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETVersionsVersionId200Response) GetData() GETVersionsVersionId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret GETVersionsVersionId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETVersionsVersionId200Response) GetDataOk() (*GETVersionsVersionId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETVersionsVersionId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GETVersionsVersionId200ResponseData and assigns it to the Data field.
func (o *GETVersionsVersionId200Response) SetData(v GETVersionsVersionId200ResponseData) {
	o.Data = &v
}

func (o GETVersionsVersionId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETVersionsVersionId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGETVersionsVersionId200Response struct {
	value *GETVersionsVersionId200Response
	isSet bool
}

func (v NullableGETVersionsVersionId200Response) Get() *GETVersionsVersionId200Response {
	return v.value
}

func (v *NullableGETVersionsVersionId200Response) Set(val *GETVersionsVersionId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETVersionsVersionId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETVersionsVersionId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETVersionsVersionId200Response(val *GETVersionsVersionId200Response) *NullableGETVersionsVersionId200Response {
	return &NullableGETVersionsVersionId200Response{value: val, isSet: true}
}

func (v NullableGETVersionsVersionId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETVersionsVersionId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
