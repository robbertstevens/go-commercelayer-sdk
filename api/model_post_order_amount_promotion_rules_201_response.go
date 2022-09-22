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

// POSTOrderAmountPromotionRules201Response struct for POSTOrderAmountPromotionRules201Response
type POSTOrderAmountPromotionRules201Response struct {
	Data *POSTOrderAmountPromotionRules201ResponseData `json:"data,omitempty"`
}

// NewPOSTOrderAmountPromotionRules201Response instantiates a new POSTOrderAmountPromotionRules201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTOrderAmountPromotionRules201Response() *POSTOrderAmountPromotionRules201Response {
	this := POSTOrderAmountPromotionRules201Response{}
	return &this
}

// NewPOSTOrderAmountPromotionRules201ResponseWithDefaults instantiates a new POSTOrderAmountPromotionRules201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTOrderAmountPromotionRules201ResponseWithDefaults() *POSTOrderAmountPromotionRules201Response {
	this := POSTOrderAmountPromotionRules201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTOrderAmountPromotionRules201Response) GetData() POSTOrderAmountPromotionRules201ResponseData {
	if o == nil || o.Data == nil {
		var ret POSTOrderAmountPromotionRules201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrderAmountPromotionRules201Response) GetDataOk() (*POSTOrderAmountPromotionRules201ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTOrderAmountPromotionRules201Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTOrderAmountPromotionRules201ResponseData and assigns it to the Data field.
func (o *POSTOrderAmountPromotionRules201Response) SetData(v POSTOrderAmountPromotionRules201ResponseData) {
	o.Data = &v
}

func (o POSTOrderAmountPromotionRules201Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePOSTOrderAmountPromotionRules201Response struct {
	value *POSTOrderAmountPromotionRules201Response
	isSet bool
}

func (v NullablePOSTOrderAmountPromotionRules201Response) Get() *POSTOrderAmountPromotionRules201Response {
	return v.value
}

func (v *NullablePOSTOrderAmountPromotionRules201Response) Set(val *POSTOrderAmountPromotionRules201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTOrderAmountPromotionRules201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTOrderAmountPromotionRules201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTOrderAmountPromotionRules201Response(val *POSTOrderAmountPromotionRules201Response) *NullablePOSTOrderAmountPromotionRules201Response {
	return &NullablePOSTOrderAmountPromotionRules201Response{value: val, isSet: true}
}

func (v NullablePOSTOrderAmountPromotionRules201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTOrderAmountPromotionRules201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
