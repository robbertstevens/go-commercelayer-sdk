/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 4.1.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the POSTCouponsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTCouponsRequest{}

// POSTCouponsRequest struct for POSTCouponsRequest
type POSTCouponsRequest struct {
	Data POSTCouponsRequestData `json:"data"`
}

// NewPOSTCouponsRequest instantiates a new POSTCouponsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTCouponsRequest(data POSTCouponsRequestData) *POSTCouponsRequest {
	this := POSTCouponsRequest{}
	this.Data = data
	return &this
}

// NewPOSTCouponsRequestWithDefaults instantiates a new POSTCouponsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTCouponsRequestWithDefaults() *POSTCouponsRequest {
	this := POSTCouponsRequest{}
	return &this
}

// GetData returns the Data field value
func (o *POSTCouponsRequest) GetData() POSTCouponsRequestData {
	if o == nil {
		var ret POSTCouponsRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *POSTCouponsRequest) GetDataOk() (*POSTCouponsRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *POSTCouponsRequest) SetData(v POSTCouponsRequestData) {
	o.Data = v
}

func (o POSTCouponsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTCouponsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullablePOSTCouponsRequest struct {
	value *POSTCouponsRequest
	isSet bool
}

func (v NullablePOSTCouponsRequest) Get() *POSTCouponsRequest {
	return v.value
}

func (v *NullablePOSTCouponsRequest) Set(val *POSTCouponsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTCouponsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTCouponsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTCouponsRequest(val *POSTCouponsRequest) *NullablePOSTCouponsRequest {
	return &NullablePOSTCouponsRequest{value: val, isSet: true}
}

func (v NullablePOSTCouponsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTCouponsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
