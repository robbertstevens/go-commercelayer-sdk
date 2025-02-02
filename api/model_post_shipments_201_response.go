/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.3.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the POSTShipments201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTShipments201Response{}

// POSTShipments201Response struct for POSTShipments201Response
type POSTShipments201Response struct {
	Data *POSTShipments201ResponseData `json:"data,omitempty"`
}

// NewPOSTShipments201Response instantiates a new POSTShipments201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTShipments201Response() *POSTShipments201Response {
	this := POSTShipments201Response{}
	return &this
}

// NewPOSTShipments201ResponseWithDefaults instantiates a new POSTShipments201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTShipments201ResponseWithDefaults() *POSTShipments201Response {
	this := POSTShipments201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTShipments201Response) GetData() POSTShipments201ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret POSTShipments201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTShipments201Response) GetDataOk() (*POSTShipments201ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTShipments201Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTShipments201ResponseData and assigns it to the Data field.
func (o *POSTShipments201Response) SetData(v POSTShipments201ResponseData) {
	o.Data = &v
}

func (o POSTShipments201Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTShipments201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTShipments201Response struct {
	value *POSTShipments201Response
	isSet bool
}

func (v NullablePOSTShipments201Response) Get() *POSTShipments201Response {
	return v.value
}

func (v *NullablePOSTShipments201Response) Set(val *POSTShipments201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTShipments201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTShipments201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTShipments201Response(val *POSTShipments201Response) *NullablePOSTShipments201Response {
	return &NullablePOSTShipments201Response{value: val, isSet: true}
}

func (v NullablePOSTShipments201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTShipments201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
