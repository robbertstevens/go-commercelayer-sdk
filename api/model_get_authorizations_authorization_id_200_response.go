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

// GETAuthorizationsAuthorizationId200Response struct for GETAuthorizationsAuthorizationId200Response
type GETAuthorizationsAuthorizationId200Response struct {
	Data *GETAuthorizations200ResponseDataInner `json:"data,omitempty"`
}

// NewGETAuthorizationsAuthorizationId200Response instantiates a new GETAuthorizationsAuthorizationId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETAuthorizationsAuthorizationId200Response() *GETAuthorizationsAuthorizationId200Response {
	this := GETAuthorizationsAuthorizationId200Response{}
	return &this
}

// NewGETAuthorizationsAuthorizationId200ResponseWithDefaults instantiates a new GETAuthorizationsAuthorizationId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETAuthorizationsAuthorizationId200ResponseWithDefaults() *GETAuthorizationsAuthorizationId200Response {
	this := GETAuthorizationsAuthorizationId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETAuthorizationsAuthorizationId200Response) GetData() GETAuthorizations200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret GETAuthorizations200ResponseDataInner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAuthorizationsAuthorizationId200Response) GetDataOk() (*GETAuthorizations200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETAuthorizationsAuthorizationId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETAuthorizations200ResponseDataInner and assigns it to the Data field.
func (o *GETAuthorizationsAuthorizationId200Response) SetData(v GETAuthorizations200ResponseDataInner) {
	o.Data = &v
}

func (o GETAuthorizationsAuthorizationId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETAuthorizationsAuthorizationId200Response struct {
	value *GETAuthorizationsAuthorizationId200Response
	isSet bool
}

func (v NullableGETAuthorizationsAuthorizationId200Response) Get() *GETAuthorizationsAuthorizationId200Response {
	return v.value
}

func (v *NullableGETAuthorizationsAuthorizationId200Response) Set(val *GETAuthorizationsAuthorizationId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETAuthorizationsAuthorizationId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETAuthorizationsAuthorizationId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETAuthorizationsAuthorizationId200Response(val *GETAuthorizationsAuthorizationId200Response) *NullableGETAuthorizationsAuthorizationId200Response {
	return &NullableGETAuthorizationsAuthorizationId200Response{value: val, isSet: true}
}

func (v NullableGETAuthorizationsAuthorizationId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETAuthorizationsAuthorizationId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
