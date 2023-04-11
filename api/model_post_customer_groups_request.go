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

// checks if the POSTCustomerGroupsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTCustomerGroupsRequest{}

// POSTCustomerGroupsRequest struct for POSTCustomerGroupsRequest
type POSTCustomerGroupsRequest struct {
	Data POSTCustomerGroupsRequestData `json:"data"`
}

// NewPOSTCustomerGroupsRequest instantiates a new POSTCustomerGroupsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTCustomerGroupsRequest(data POSTCustomerGroupsRequestData) *POSTCustomerGroupsRequest {
	this := POSTCustomerGroupsRequest{}
	this.Data = data
	return &this
}

// NewPOSTCustomerGroupsRequestWithDefaults instantiates a new POSTCustomerGroupsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTCustomerGroupsRequestWithDefaults() *POSTCustomerGroupsRequest {
	this := POSTCustomerGroupsRequest{}
	return &this
}

// GetData returns the Data field value
func (o *POSTCustomerGroupsRequest) GetData() POSTCustomerGroupsRequestData {
	if o == nil {
		var ret POSTCustomerGroupsRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *POSTCustomerGroupsRequest) GetDataOk() (*POSTCustomerGroupsRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *POSTCustomerGroupsRequest) SetData(v POSTCustomerGroupsRequestData) {
	o.Data = v
}

func (o POSTCustomerGroupsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTCustomerGroupsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullablePOSTCustomerGroupsRequest struct {
	value *POSTCustomerGroupsRequest
	isSet bool
}

func (v NullablePOSTCustomerGroupsRequest) Get() *POSTCustomerGroupsRequest {
	return v.value
}

func (v *NullablePOSTCustomerGroupsRequest) Set(val *POSTCustomerGroupsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTCustomerGroupsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTCustomerGroupsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTCustomerGroupsRequest(val *POSTCustomerGroupsRequest) *NullablePOSTCustomerGroupsRequest {
	return &NullablePOSTCustomerGroupsRequest{value: val, isSet: true}
}

func (v NullablePOSTCustomerGroupsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTCustomerGroupsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
