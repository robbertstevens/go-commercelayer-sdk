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

// POSTCouponRecipients201Response struct for POSTCouponRecipients201Response
type POSTCouponRecipients201Response struct {
	Data *POSTCouponRecipients201ResponseData `json:"data,omitempty"`
}

// NewPOSTCouponRecipients201Response instantiates a new POSTCouponRecipients201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTCouponRecipients201Response() *POSTCouponRecipients201Response {
	this := POSTCouponRecipients201Response{}
	return &this
}

// NewPOSTCouponRecipients201ResponseWithDefaults instantiates a new POSTCouponRecipients201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTCouponRecipients201ResponseWithDefaults() *POSTCouponRecipients201Response {
	this := POSTCouponRecipients201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTCouponRecipients201Response) GetData() POSTCouponRecipients201ResponseData {
	if o == nil || o.Data == nil {
		var ret POSTCouponRecipients201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCouponRecipients201Response) GetDataOk() (*POSTCouponRecipients201ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTCouponRecipients201Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTCouponRecipients201ResponseData and assigns it to the Data field.
func (o *POSTCouponRecipients201Response) SetData(v POSTCouponRecipients201ResponseData) {
	o.Data = &v
}

func (o POSTCouponRecipients201Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePOSTCouponRecipients201Response struct {
	value *POSTCouponRecipients201Response
	isSet bool
}

func (v NullablePOSTCouponRecipients201Response) Get() *POSTCouponRecipients201Response {
	return v.value
}

func (v *NullablePOSTCouponRecipients201Response) Set(val *POSTCouponRecipients201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTCouponRecipients201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTCouponRecipients201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTCouponRecipients201Response(val *POSTCouponRecipients201Response) *NullablePOSTCouponRecipients201Response {
	return &NullablePOSTCouponRecipients201Response{value: val, isSet: true}
}

func (v NullablePOSTCouponRecipients201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTCouponRecipients201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
