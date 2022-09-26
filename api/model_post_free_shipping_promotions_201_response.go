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

// POSTFreeShippingPromotions201Response struct for POSTFreeShippingPromotions201Response
type POSTFreeShippingPromotions201Response struct {
	Data *POSTFreeShippingPromotions201ResponseData `json:"data,omitempty"`
}

// NewPOSTFreeShippingPromotions201Response instantiates a new POSTFreeShippingPromotions201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTFreeShippingPromotions201Response() *POSTFreeShippingPromotions201Response {
	this := POSTFreeShippingPromotions201Response{}
	return &this
}

// NewPOSTFreeShippingPromotions201ResponseWithDefaults instantiates a new POSTFreeShippingPromotions201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTFreeShippingPromotions201ResponseWithDefaults() *POSTFreeShippingPromotions201Response {
	this := POSTFreeShippingPromotions201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTFreeShippingPromotions201Response) GetData() POSTFreeShippingPromotions201ResponseData {
	if o == nil || o.Data == nil {
		var ret POSTFreeShippingPromotions201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTFreeShippingPromotions201Response) GetDataOk() (*POSTFreeShippingPromotions201ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTFreeShippingPromotions201Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTFreeShippingPromotions201ResponseData and assigns it to the Data field.
func (o *POSTFreeShippingPromotions201Response) SetData(v POSTFreeShippingPromotions201ResponseData) {
	o.Data = &v
}

func (o POSTFreeShippingPromotions201Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePOSTFreeShippingPromotions201Response struct {
	value *POSTFreeShippingPromotions201Response
	isSet bool
}

func (v NullablePOSTFreeShippingPromotions201Response) Get() *POSTFreeShippingPromotions201Response {
	return v.value
}

func (v *NullablePOSTFreeShippingPromotions201Response) Set(val *POSTFreeShippingPromotions201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTFreeShippingPromotions201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTFreeShippingPromotions201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTFreeShippingPromotions201Response(val *POSTFreeShippingPromotions201Response) *NullablePOSTFreeShippingPromotions201Response {
	return &NullablePOSTFreeShippingPromotions201Response{value: val, isSet: true}
}

func (v NullablePOSTFreeShippingPromotions201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTFreeShippingPromotions201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
