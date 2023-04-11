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

// checks if the POSTSubscriptionModelsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTSubscriptionModelsRequest{}

// POSTSubscriptionModelsRequest struct for POSTSubscriptionModelsRequest
type POSTSubscriptionModelsRequest struct {
	Data POSTSubscriptionModelsRequestData `json:"data"`
}

// NewPOSTSubscriptionModelsRequest instantiates a new POSTSubscriptionModelsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTSubscriptionModelsRequest(data POSTSubscriptionModelsRequestData) *POSTSubscriptionModelsRequest {
	this := POSTSubscriptionModelsRequest{}
	this.Data = data
	return &this
}

// NewPOSTSubscriptionModelsRequestWithDefaults instantiates a new POSTSubscriptionModelsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTSubscriptionModelsRequestWithDefaults() *POSTSubscriptionModelsRequest {
	this := POSTSubscriptionModelsRequest{}
	return &this
}

// GetData returns the Data field value
func (o *POSTSubscriptionModelsRequest) GetData() POSTSubscriptionModelsRequestData {
	if o == nil {
		var ret POSTSubscriptionModelsRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *POSTSubscriptionModelsRequest) GetDataOk() (*POSTSubscriptionModelsRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *POSTSubscriptionModelsRequest) SetData(v POSTSubscriptionModelsRequestData) {
	o.Data = v
}

func (o POSTSubscriptionModelsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTSubscriptionModelsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullablePOSTSubscriptionModelsRequest struct {
	value *POSTSubscriptionModelsRequest
	isSet bool
}

func (v NullablePOSTSubscriptionModelsRequest) Get() *POSTSubscriptionModelsRequest {
	return v.value
}

func (v *NullablePOSTSubscriptionModelsRequest) Set(val *POSTSubscriptionModelsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTSubscriptionModelsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTSubscriptionModelsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTSubscriptionModelsRequest(val *POSTSubscriptionModelsRequest) *NullablePOSTSubscriptionModelsRequest {
	return &NullablePOSTSubscriptionModelsRequest{value: val, isSet: true}
}

func (v NullablePOSTSubscriptionModelsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTSubscriptionModelsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
