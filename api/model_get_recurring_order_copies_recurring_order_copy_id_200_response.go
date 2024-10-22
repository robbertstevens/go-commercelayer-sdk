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

// checks if the GETRecurringOrderCopiesRecurringOrderCopyId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETRecurringOrderCopiesRecurringOrderCopyId200Response{}

// GETRecurringOrderCopiesRecurringOrderCopyId200Response struct for GETRecurringOrderCopiesRecurringOrderCopyId200Response
type GETRecurringOrderCopiesRecurringOrderCopyId200Response struct {
	Data *GETRecurringOrderCopiesRecurringOrderCopyId200ResponseData `json:"data,omitempty"`
}

// NewGETRecurringOrderCopiesRecurringOrderCopyId200Response instantiates a new GETRecurringOrderCopiesRecurringOrderCopyId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETRecurringOrderCopiesRecurringOrderCopyId200Response() *GETRecurringOrderCopiesRecurringOrderCopyId200Response {
	this := GETRecurringOrderCopiesRecurringOrderCopyId200Response{}
	return &this
}

// NewGETRecurringOrderCopiesRecurringOrderCopyId200ResponseWithDefaults instantiates a new GETRecurringOrderCopiesRecurringOrderCopyId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETRecurringOrderCopiesRecurringOrderCopyId200ResponseWithDefaults() *GETRecurringOrderCopiesRecurringOrderCopyId200Response {
	this := GETRecurringOrderCopiesRecurringOrderCopyId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETRecurringOrderCopiesRecurringOrderCopyId200Response) GetData() GETRecurringOrderCopiesRecurringOrderCopyId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret GETRecurringOrderCopiesRecurringOrderCopyId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETRecurringOrderCopiesRecurringOrderCopyId200Response) GetDataOk() (*GETRecurringOrderCopiesRecurringOrderCopyId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETRecurringOrderCopiesRecurringOrderCopyId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GETRecurringOrderCopiesRecurringOrderCopyId200ResponseData and assigns it to the Data field.
func (o *GETRecurringOrderCopiesRecurringOrderCopyId200Response) SetData(v GETRecurringOrderCopiesRecurringOrderCopyId200ResponseData) {
	o.Data = &v
}

func (o GETRecurringOrderCopiesRecurringOrderCopyId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETRecurringOrderCopiesRecurringOrderCopyId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGETRecurringOrderCopiesRecurringOrderCopyId200Response struct {
	value *GETRecurringOrderCopiesRecurringOrderCopyId200Response
	isSet bool
}

func (v NullableGETRecurringOrderCopiesRecurringOrderCopyId200Response) Get() *GETRecurringOrderCopiesRecurringOrderCopyId200Response {
	return v.value
}

func (v *NullableGETRecurringOrderCopiesRecurringOrderCopyId200Response) Set(val *GETRecurringOrderCopiesRecurringOrderCopyId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETRecurringOrderCopiesRecurringOrderCopyId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETRecurringOrderCopiesRecurringOrderCopyId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETRecurringOrderCopiesRecurringOrderCopyId200Response(val *GETRecurringOrderCopiesRecurringOrderCopyId200Response) *NullableGETRecurringOrderCopiesRecurringOrderCopyId200Response {
	return &NullableGETRecurringOrderCopiesRecurringOrderCopyId200Response{value: val, isSet: true}
}

func (v NullableGETRecurringOrderCopiesRecurringOrderCopyId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETRecurringOrderCopiesRecurringOrderCopyId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
