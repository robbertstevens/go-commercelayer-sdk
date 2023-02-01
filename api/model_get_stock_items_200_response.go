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

// GETStockItems200Response struct for GETStockItems200Response
type GETStockItems200Response struct {
	Data []GETStockItems200ResponseDataInner `json:"data,omitempty"`
}

// NewGETStockItems200Response instantiates a new GETStockItems200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETStockItems200Response() *GETStockItems200Response {
	this := GETStockItems200Response{}
	return &this
}

// NewGETStockItems200ResponseWithDefaults instantiates a new GETStockItems200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETStockItems200ResponseWithDefaults() *GETStockItems200Response {
	this := GETStockItems200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETStockItems200Response) GetData() []GETStockItems200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret []GETStockItems200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStockItems200Response) GetDataOk() ([]GETStockItems200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETStockItems200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETStockItems200ResponseDataInner and assigns it to the Data field.
func (o *GETStockItems200Response) SetData(v []GETStockItems200ResponseDataInner) {
	o.Data = v
}

func (o GETStockItems200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETStockItems200Response struct {
	value *GETStockItems200Response
	isSet bool
}

func (v NullableGETStockItems200Response) Get() *GETStockItems200Response {
	return v.value
}

func (v *NullableGETStockItems200Response) Set(val *GETStockItems200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETStockItems200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETStockItems200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETStockItems200Response(val *GETStockItems200Response) *NullableGETStockItems200Response {
	return &NullableGETStockItems200Response{value: val, isSet: true}
}

func (v NullableGETStockItems200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETStockItems200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
