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

// GETCoupons200Response struct for GETCoupons200Response
type GETCoupons200Response struct {
	Data []GETCoupons200ResponseDataInner `json:"data,omitempty"`
}

// NewGETCoupons200Response instantiates a new GETCoupons200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCoupons200Response() *GETCoupons200Response {
	this := GETCoupons200Response{}
	return &this
}

// NewGETCoupons200ResponseWithDefaults instantiates a new GETCoupons200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCoupons200ResponseWithDefaults() *GETCoupons200Response {
	this := GETCoupons200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETCoupons200Response) GetData() []GETCoupons200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret []GETCoupons200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCoupons200Response) GetDataOk() ([]GETCoupons200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETCoupons200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETCoupons200ResponseDataInner and assigns it to the Data field.
func (o *GETCoupons200Response) SetData(v []GETCoupons200ResponseDataInner) {
	o.Data = v
}

func (o GETCoupons200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETCoupons200Response struct {
	value *GETCoupons200Response
	isSet bool
}

func (v NullableGETCoupons200Response) Get() *GETCoupons200Response {
	return v.value
}

func (v *NullableGETCoupons200Response) Set(val *GETCoupons200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCoupons200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCoupons200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCoupons200Response(val *GETCoupons200Response) *NullableGETCoupons200Response {
	return &NullableGETCoupons200Response{value: val, isSet: true}
}

func (v NullableGETCoupons200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCoupons200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
