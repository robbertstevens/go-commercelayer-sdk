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

// GETExternalPromotions200Response struct for GETExternalPromotions200Response
type GETExternalPromotions200Response struct {
	Data []GETExternalPromotions200ResponseDataInner `json:"data,omitempty"`
}

// NewGETExternalPromotions200Response instantiates a new GETExternalPromotions200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETExternalPromotions200Response() *GETExternalPromotions200Response {
	this := GETExternalPromotions200Response{}
	return &this
}

// NewGETExternalPromotions200ResponseWithDefaults instantiates a new GETExternalPromotions200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETExternalPromotions200ResponseWithDefaults() *GETExternalPromotions200Response {
	this := GETExternalPromotions200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETExternalPromotions200Response) GetData() []GETExternalPromotions200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret []GETExternalPromotions200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETExternalPromotions200Response) GetDataOk() ([]GETExternalPromotions200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETExternalPromotions200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETExternalPromotions200ResponseDataInner and assigns it to the Data field.
func (o *GETExternalPromotions200Response) SetData(v []GETExternalPromotions200ResponseDataInner) {
	o.Data = v
}

func (o GETExternalPromotions200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETExternalPromotions200Response struct {
	value *GETExternalPromotions200Response
	isSet bool
}

func (v NullableGETExternalPromotions200Response) Get() *GETExternalPromotions200Response {
	return v.value
}

func (v *NullableGETExternalPromotions200Response) Set(val *GETExternalPromotions200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETExternalPromotions200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETExternalPromotions200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETExternalPromotions200Response(val *GETExternalPromotions200Response) *NullableGETExternalPromotions200Response {
	return &NullableGETExternalPromotions200Response{value: val, isSet: true}
}

func (v NullableGETExternalPromotions200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETExternalPromotions200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
