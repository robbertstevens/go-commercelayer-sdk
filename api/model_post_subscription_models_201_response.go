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

// checks if the POSTSubscriptionModels201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTSubscriptionModels201Response{}

// POSTSubscriptionModels201Response struct for POSTSubscriptionModels201Response
type POSTSubscriptionModels201Response struct {
	Data *POSTSubscriptionModels201ResponseData `json:"data,omitempty"`
}

// NewPOSTSubscriptionModels201Response instantiates a new POSTSubscriptionModels201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTSubscriptionModels201Response() *POSTSubscriptionModels201Response {
	this := POSTSubscriptionModels201Response{}
	return &this
}

// NewPOSTSubscriptionModels201ResponseWithDefaults instantiates a new POSTSubscriptionModels201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTSubscriptionModels201ResponseWithDefaults() *POSTSubscriptionModels201Response {
	this := POSTSubscriptionModels201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTSubscriptionModels201Response) GetData() POSTSubscriptionModels201ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret POSTSubscriptionModels201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTSubscriptionModels201Response) GetDataOk() (*POSTSubscriptionModels201ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTSubscriptionModels201Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTSubscriptionModels201ResponseData and assigns it to the Data field.
func (o *POSTSubscriptionModels201Response) SetData(v POSTSubscriptionModels201ResponseData) {
	o.Data = &v
}

func (o POSTSubscriptionModels201Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTSubscriptionModels201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTSubscriptionModels201Response struct {
	value *POSTSubscriptionModels201Response
	isSet bool
}

func (v NullablePOSTSubscriptionModels201Response) Get() *POSTSubscriptionModels201Response {
	return v.value
}

func (v *NullablePOSTSubscriptionModels201Response) Set(val *POSTSubscriptionModels201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTSubscriptionModels201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTSubscriptionModels201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTSubscriptionModels201Response(val *POSTSubscriptionModels201Response) *NullablePOSTSubscriptionModels201Response {
	return &NullablePOSTSubscriptionModels201Response{value: val, isSet: true}
}

func (v NullablePOSTSubscriptionModels201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTSubscriptionModels201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
