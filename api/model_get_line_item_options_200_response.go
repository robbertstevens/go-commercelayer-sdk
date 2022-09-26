/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETLineItemOptions200Response struct for GETLineItemOptions200Response
type GETLineItemOptions200Response struct {
	Data []GETLineItemOptions200ResponseDataInner `json:"data,omitempty"`
}

// NewGETLineItemOptions200Response instantiates a new GETLineItemOptions200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETLineItemOptions200Response() *GETLineItemOptions200Response {
	this := GETLineItemOptions200Response{}
	return &this
}

// NewGETLineItemOptions200ResponseWithDefaults instantiates a new GETLineItemOptions200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETLineItemOptions200ResponseWithDefaults() *GETLineItemOptions200Response {
	this := GETLineItemOptions200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETLineItemOptions200Response) GetData() []GETLineItemOptions200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret []GETLineItemOptions200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETLineItemOptions200Response) GetDataOk() ([]GETLineItemOptions200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETLineItemOptions200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETLineItemOptions200ResponseDataInner and assigns it to the Data field.
func (o *GETLineItemOptions200Response) SetData(v []GETLineItemOptions200ResponseDataInner) {
	o.Data = v
}

func (o GETLineItemOptions200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETLineItemOptions200Response struct {
	value *GETLineItemOptions200Response
	isSet bool
}

func (v NullableGETLineItemOptions200Response) Get() *GETLineItemOptions200Response {
	return v.value
}

func (v *NullableGETLineItemOptions200Response) Set(val *GETLineItemOptions200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETLineItemOptions200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETLineItemOptions200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETLineItemOptions200Response(val *GETLineItemOptions200Response) *NullableGETLineItemOptions200Response {
	return &NullableGETLineItemOptions200Response{value: val, isSet: true}
}

func (v NullableGETLineItemOptions200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETLineItemOptions200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
