/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.3.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the POSTBuyXPayYPromotions201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTBuyXPayYPromotions201Response{}

// POSTBuyXPayYPromotions201Response struct for POSTBuyXPayYPromotions201Response
type POSTBuyXPayYPromotions201Response struct {
	Data *POSTBuyXPayYPromotions201ResponseData `json:"data,omitempty"`
}

// NewPOSTBuyXPayYPromotions201Response instantiates a new POSTBuyXPayYPromotions201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTBuyXPayYPromotions201Response() *POSTBuyXPayYPromotions201Response {
	this := POSTBuyXPayYPromotions201Response{}
	return &this
}

// NewPOSTBuyXPayYPromotions201ResponseWithDefaults instantiates a new POSTBuyXPayYPromotions201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTBuyXPayYPromotions201ResponseWithDefaults() *POSTBuyXPayYPromotions201Response {
	this := POSTBuyXPayYPromotions201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTBuyXPayYPromotions201Response) GetData() POSTBuyXPayYPromotions201ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret POSTBuyXPayYPromotions201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTBuyXPayYPromotions201Response) GetDataOk() (*POSTBuyXPayYPromotions201ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTBuyXPayYPromotions201Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTBuyXPayYPromotions201ResponseData and assigns it to the Data field.
func (o *POSTBuyXPayYPromotions201Response) SetData(v POSTBuyXPayYPromotions201ResponseData) {
	o.Data = &v
}

func (o POSTBuyXPayYPromotions201Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTBuyXPayYPromotions201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTBuyXPayYPromotions201Response struct {
	value *POSTBuyXPayYPromotions201Response
	isSet bool
}

func (v NullablePOSTBuyXPayYPromotions201Response) Get() *POSTBuyXPayYPromotions201Response {
	return v.value
}

func (v *NullablePOSTBuyXPayYPromotions201Response) Set(val *POSTBuyXPayYPromotions201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTBuyXPayYPromotions201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTBuyXPayYPromotions201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTBuyXPayYPromotions201Response(val *POSTBuyXPayYPromotions201Response) *NullablePOSTBuyXPayYPromotions201Response {
	return &NullablePOSTBuyXPayYPromotions201Response{value: val, isSet: true}
}

func (v NullablePOSTBuyXPayYPromotions201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTBuyXPayYPromotions201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
