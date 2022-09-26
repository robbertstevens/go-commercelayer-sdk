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

// GETSkuListsSkuListId200Response struct for GETSkuListsSkuListId200Response
type GETSkuListsSkuListId200Response struct {
	Data *GETSkuLists200ResponseDataInner `json:"data,omitempty"`
}

// NewGETSkuListsSkuListId200Response instantiates a new GETSkuListsSkuListId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETSkuListsSkuListId200Response() *GETSkuListsSkuListId200Response {
	this := GETSkuListsSkuListId200Response{}
	return &this
}

// NewGETSkuListsSkuListId200ResponseWithDefaults instantiates a new GETSkuListsSkuListId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETSkuListsSkuListId200ResponseWithDefaults() *GETSkuListsSkuListId200Response {
	this := GETSkuListsSkuListId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETSkuListsSkuListId200Response) GetData() GETSkuLists200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret GETSkuLists200ResponseDataInner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkuListsSkuListId200Response) GetDataOk() (*GETSkuLists200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETSkuListsSkuListId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETSkuLists200ResponseDataInner and assigns it to the Data field.
func (o *GETSkuListsSkuListId200Response) SetData(v GETSkuLists200ResponseDataInner) {
	o.Data = &v
}

func (o GETSkuListsSkuListId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETSkuListsSkuListId200Response struct {
	value *GETSkuListsSkuListId200Response
	isSet bool
}

func (v NullableGETSkuListsSkuListId200Response) Get() *GETSkuListsSkuListId200Response {
	return v.value
}

func (v *NullableGETSkuListsSkuListId200Response) Set(val *GETSkuListsSkuListId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETSkuListsSkuListId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETSkuListsSkuListId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETSkuListsSkuListId200Response(val *GETSkuListsSkuListId200Response) *NullableGETSkuListsSkuListId200Response {
	return &NullableGETSkuListsSkuListId200Response{value: val, isSet: true}
}

func (v NullableGETSkuListsSkuListId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETSkuListsSkuListId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
