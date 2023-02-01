/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.4.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETCustomerPasswordResetsCustomerPasswordResetId200Response struct for GETCustomerPasswordResetsCustomerPasswordResetId200Response
type GETCustomerPasswordResetsCustomerPasswordResetId200Response struct {
	Data *GETCustomerPasswordResets200ResponseDataInner `json:"data,omitempty"`
}

// NewGETCustomerPasswordResetsCustomerPasswordResetId200Response instantiates a new GETCustomerPasswordResetsCustomerPasswordResetId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCustomerPasswordResetsCustomerPasswordResetId200Response() *GETCustomerPasswordResetsCustomerPasswordResetId200Response {
	this := GETCustomerPasswordResetsCustomerPasswordResetId200Response{}
	return &this
}

// NewGETCustomerPasswordResetsCustomerPasswordResetId200ResponseWithDefaults instantiates a new GETCustomerPasswordResetsCustomerPasswordResetId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCustomerPasswordResetsCustomerPasswordResetId200ResponseWithDefaults() *GETCustomerPasswordResetsCustomerPasswordResetId200Response {
	this := GETCustomerPasswordResetsCustomerPasswordResetId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETCustomerPasswordResetsCustomerPasswordResetId200Response) GetData() GETCustomerPasswordResets200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret GETCustomerPasswordResets200ResponseDataInner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCustomerPasswordResetsCustomerPasswordResetId200Response) GetDataOk() (*GETCustomerPasswordResets200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETCustomerPasswordResetsCustomerPasswordResetId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETCustomerPasswordResets200ResponseDataInner and assigns it to the Data field.
func (o *GETCustomerPasswordResetsCustomerPasswordResetId200Response) SetData(v GETCustomerPasswordResets200ResponseDataInner) {
	o.Data = &v
}

func (o GETCustomerPasswordResetsCustomerPasswordResetId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETCustomerPasswordResetsCustomerPasswordResetId200Response struct {
	value *GETCustomerPasswordResetsCustomerPasswordResetId200Response
	isSet bool
}

func (v NullableGETCustomerPasswordResetsCustomerPasswordResetId200Response) Get() *GETCustomerPasswordResetsCustomerPasswordResetId200Response {
	return v.value
}

func (v *NullableGETCustomerPasswordResetsCustomerPasswordResetId200Response) Set(val *GETCustomerPasswordResetsCustomerPasswordResetId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCustomerPasswordResetsCustomerPasswordResetId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCustomerPasswordResetsCustomerPasswordResetId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCustomerPasswordResetsCustomerPasswordResetId200Response(val *GETCustomerPasswordResetsCustomerPasswordResetId200Response) *NullableGETCustomerPasswordResetsCustomerPasswordResetId200Response {
	return &NullableGETCustomerPasswordResetsCustomerPasswordResetId200Response{value: val, isSet: true}
}

func (v NullableGETCustomerPasswordResetsCustomerPasswordResetId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCustomerPasswordResetsCustomerPasswordResetId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
