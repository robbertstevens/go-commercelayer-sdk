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

// GETMerchants200Response struct for GETMerchants200Response
type GETMerchants200Response struct {
	Data []GETMerchants200ResponseDataInner `json:"data,omitempty"`
}

// NewGETMerchants200Response instantiates a new GETMerchants200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETMerchants200Response() *GETMerchants200Response {
	this := GETMerchants200Response{}
	return &this
}

// NewGETMerchants200ResponseWithDefaults instantiates a new GETMerchants200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETMerchants200ResponseWithDefaults() *GETMerchants200Response {
	this := GETMerchants200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETMerchants200Response) GetData() []GETMerchants200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret []GETMerchants200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETMerchants200Response) GetDataOk() ([]GETMerchants200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETMerchants200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETMerchants200ResponseDataInner and assigns it to the Data field.
func (o *GETMerchants200Response) SetData(v []GETMerchants200ResponseDataInner) {
	o.Data = v
}

func (o GETMerchants200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETMerchants200Response struct {
	value *GETMerchants200Response
	isSet bool
}

func (v NullableGETMerchants200Response) Get() *GETMerchants200Response {
	return v.value
}

func (v *NullableGETMerchants200Response) Set(val *GETMerchants200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETMerchants200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETMerchants200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETMerchants200Response(val *GETMerchants200Response) *NullableGETMerchants200Response {
	return &NullableGETMerchants200Response{value: val, isSet: true}
}

func (v NullableGETMerchants200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETMerchants200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


