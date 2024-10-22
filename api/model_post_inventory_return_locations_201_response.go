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

// checks if the POSTInventoryReturnLocations201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTInventoryReturnLocations201Response{}

// POSTInventoryReturnLocations201Response struct for POSTInventoryReturnLocations201Response
type POSTInventoryReturnLocations201Response struct {
	Data *POSTInventoryReturnLocations201ResponseData `json:"data,omitempty"`
}

// NewPOSTInventoryReturnLocations201Response instantiates a new POSTInventoryReturnLocations201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTInventoryReturnLocations201Response() *POSTInventoryReturnLocations201Response {
	this := POSTInventoryReturnLocations201Response{}
	return &this
}

// NewPOSTInventoryReturnLocations201ResponseWithDefaults instantiates a new POSTInventoryReturnLocations201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTInventoryReturnLocations201ResponseWithDefaults() *POSTInventoryReturnLocations201Response {
	this := POSTInventoryReturnLocations201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTInventoryReturnLocations201Response) GetData() POSTInventoryReturnLocations201ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret POSTInventoryReturnLocations201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTInventoryReturnLocations201Response) GetDataOk() (*POSTInventoryReturnLocations201ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTInventoryReturnLocations201Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTInventoryReturnLocations201ResponseData and assigns it to the Data field.
func (o *POSTInventoryReturnLocations201Response) SetData(v POSTInventoryReturnLocations201ResponseData) {
	o.Data = &v
}

func (o POSTInventoryReturnLocations201Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTInventoryReturnLocations201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTInventoryReturnLocations201Response struct {
	value *POSTInventoryReturnLocations201Response
	isSet bool
}

func (v NullablePOSTInventoryReturnLocations201Response) Get() *POSTInventoryReturnLocations201Response {
	return v.value
}

func (v *NullablePOSTInventoryReturnLocations201Response) Set(val *POSTInventoryReturnLocations201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTInventoryReturnLocations201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTInventoryReturnLocations201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTInventoryReturnLocations201Response(val *POSTInventoryReturnLocations201Response) *NullablePOSTInventoryReturnLocations201Response {
	return &NullablePOSTInventoryReturnLocations201Response{value: val, isSet: true}
}

func (v NullablePOSTInventoryReturnLocations201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTInventoryReturnLocations201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
