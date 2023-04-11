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

// checks if the POSTCustomerSubscriptionsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTCustomerSubscriptionsRequest{}

// POSTCustomerSubscriptionsRequest struct for POSTCustomerSubscriptionsRequest
type POSTCustomerSubscriptionsRequest struct {
	Data POSTCustomerSubscriptionsRequestData `json:"data"`
}

// NewPOSTCustomerSubscriptionsRequest instantiates a new POSTCustomerSubscriptionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTCustomerSubscriptionsRequest(data POSTCustomerSubscriptionsRequestData) *POSTCustomerSubscriptionsRequest {
	this := POSTCustomerSubscriptionsRequest{}
	this.Data = data
	return &this
}

// NewPOSTCustomerSubscriptionsRequestWithDefaults instantiates a new POSTCustomerSubscriptionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTCustomerSubscriptionsRequestWithDefaults() *POSTCustomerSubscriptionsRequest {
	this := POSTCustomerSubscriptionsRequest{}
	return &this
}

// GetData returns the Data field value
func (o *POSTCustomerSubscriptionsRequest) GetData() POSTCustomerSubscriptionsRequestData {
	if o == nil {
		var ret POSTCustomerSubscriptionsRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *POSTCustomerSubscriptionsRequest) GetDataOk() (*POSTCustomerSubscriptionsRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *POSTCustomerSubscriptionsRequest) SetData(v POSTCustomerSubscriptionsRequestData) {
	o.Data = v
}

func (o POSTCustomerSubscriptionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTCustomerSubscriptionsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullablePOSTCustomerSubscriptionsRequest struct {
	value *POSTCustomerSubscriptionsRequest
	isSet bool
}

func (v NullablePOSTCustomerSubscriptionsRequest) Get() *POSTCustomerSubscriptionsRequest {
	return v.value
}

func (v *NullablePOSTCustomerSubscriptionsRequest) Set(val *POSTCustomerSubscriptionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTCustomerSubscriptionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTCustomerSubscriptionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTCustomerSubscriptionsRequest(val *POSTCustomerSubscriptionsRequest) *NullablePOSTCustomerSubscriptionsRequest {
	return &NullablePOSTCustomerSubscriptionsRequest{value: val, isSet: true}
}

func (v NullablePOSTCustomerSubscriptionsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTCustomerSubscriptionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
