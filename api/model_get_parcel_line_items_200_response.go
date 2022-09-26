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

// GETParcelLineItems200Response struct for GETParcelLineItems200Response
type GETParcelLineItems200Response struct {
	Data []GETParcelLineItems200ResponseDataInner `json:"data,omitempty"`
}

// NewGETParcelLineItems200Response instantiates a new GETParcelLineItems200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETParcelLineItems200Response() *GETParcelLineItems200Response {
	this := GETParcelLineItems200Response{}
	return &this
}

// NewGETParcelLineItems200ResponseWithDefaults instantiates a new GETParcelLineItems200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETParcelLineItems200ResponseWithDefaults() *GETParcelLineItems200Response {
	this := GETParcelLineItems200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETParcelLineItems200Response) GetData() []GETParcelLineItems200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret []GETParcelLineItems200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETParcelLineItems200Response) GetDataOk() ([]GETParcelLineItems200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETParcelLineItems200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETParcelLineItems200ResponseDataInner and assigns it to the Data field.
func (o *GETParcelLineItems200Response) SetData(v []GETParcelLineItems200ResponseDataInner) {
	o.Data = v
}

func (o GETParcelLineItems200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETParcelLineItems200Response struct {
	value *GETParcelLineItems200Response
	isSet bool
}

func (v NullableGETParcelLineItems200Response) Get() *GETParcelLineItems200Response {
	return v.value
}

func (v *NullableGETParcelLineItems200Response) Set(val *GETParcelLineItems200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETParcelLineItems200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETParcelLineItems200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETParcelLineItems200Response(val *GETParcelLineItems200Response) *NullableGETParcelLineItems200Response {
	return &NullableGETParcelLineItems200Response{value: val, isSet: true}
}

func (v NullableGETParcelLineItems200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETParcelLineItems200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
