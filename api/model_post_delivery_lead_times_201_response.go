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

// POSTDeliveryLeadTimes201Response struct for POSTDeliveryLeadTimes201Response
type POSTDeliveryLeadTimes201Response struct {
	Data *POSTDeliveryLeadTimes201ResponseData `json:"data,omitempty"`
}

// NewPOSTDeliveryLeadTimes201Response instantiates a new POSTDeliveryLeadTimes201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTDeliveryLeadTimes201Response() *POSTDeliveryLeadTimes201Response {
	this := POSTDeliveryLeadTimes201Response{}
	return &this
}

// NewPOSTDeliveryLeadTimes201ResponseWithDefaults instantiates a new POSTDeliveryLeadTimes201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTDeliveryLeadTimes201ResponseWithDefaults() *POSTDeliveryLeadTimes201Response {
	this := POSTDeliveryLeadTimes201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTDeliveryLeadTimes201Response) GetData() POSTDeliveryLeadTimes201ResponseData {
	if o == nil || o.Data == nil {
		var ret POSTDeliveryLeadTimes201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTDeliveryLeadTimes201Response) GetDataOk() (*POSTDeliveryLeadTimes201ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTDeliveryLeadTimes201Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTDeliveryLeadTimes201ResponseData and assigns it to the Data field.
func (o *POSTDeliveryLeadTimes201Response) SetData(v POSTDeliveryLeadTimes201ResponseData) {
	o.Data = &v
}

func (o POSTDeliveryLeadTimes201Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePOSTDeliveryLeadTimes201Response struct {
	value *POSTDeliveryLeadTimes201Response
	isSet bool
}

func (v NullablePOSTDeliveryLeadTimes201Response) Get() *POSTDeliveryLeadTimes201Response {
	return v.value
}

func (v *NullablePOSTDeliveryLeadTimes201Response) Set(val *POSTDeliveryLeadTimes201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTDeliveryLeadTimes201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTDeliveryLeadTimes201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTDeliveryLeadTimes201Response(val *POSTDeliveryLeadTimes201Response) *NullablePOSTDeliveryLeadTimes201Response {
	return &NullablePOSTDeliveryLeadTimes201Response{value: val, isSet: true}
}

func (v NullablePOSTDeliveryLeadTimes201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTDeliveryLeadTimes201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
