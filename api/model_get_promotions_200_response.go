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

// GETPromotions200Response struct for GETPromotions200Response
type GETPromotions200Response struct {
	Data []GETPromotions200ResponseDataInner `json:"data,omitempty"`
}

// NewGETPromotions200Response instantiates a new GETPromotions200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETPromotions200Response() *GETPromotions200Response {
	this := GETPromotions200Response{}
	return &this
}

// NewGETPromotions200ResponseWithDefaults instantiates a new GETPromotions200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETPromotions200ResponseWithDefaults() *GETPromotions200Response {
	this := GETPromotions200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETPromotions200Response) GetData() []GETPromotions200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret []GETPromotions200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETPromotions200Response) GetDataOk() ([]GETPromotions200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETPromotions200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETPromotions200ResponseDataInner and assigns it to the Data field.
func (o *GETPromotions200Response) SetData(v []GETPromotions200ResponseDataInner) {
	o.Data = v
}

func (o GETPromotions200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETPromotions200Response struct {
	value *GETPromotions200Response
	isSet bool
}

func (v NullableGETPromotions200Response) Get() *GETPromotions200Response {
	return v.value
}

func (v *NullableGETPromotions200Response) Set(val *GETPromotions200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETPromotions200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETPromotions200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETPromotions200Response(val *GETPromotions200Response) *NullableGETPromotions200Response {
	return &NullableGETPromotions200Response{value: val, isSet: true}
}

func (v NullableGETPromotions200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETPromotions200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
