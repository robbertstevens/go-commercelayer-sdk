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

// checks if the POSTStripePaymentsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTStripePaymentsRequest{}

// POSTStripePaymentsRequest struct for POSTStripePaymentsRequest
type POSTStripePaymentsRequest struct {
	Data POSTStripePaymentsRequestData `json:"data"`
}

// NewPOSTStripePaymentsRequest instantiates a new POSTStripePaymentsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTStripePaymentsRequest(data POSTStripePaymentsRequestData) *POSTStripePaymentsRequest {
	this := POSTStripePaymentsRequest{}
	this.Data = data
	return &this
}

// NewPOSTStripePaymentsRequestWithDefaults instantiates a new POSTStripePaymentsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTStripePaymentsRequestWithDefaults() *POSTStripePaymentsRequest {
	this := POSTStripePaymentsRequest{}
	return &this
}

// GetData returns the Data field value
func (o *POSTStripePaymentsRequest) GetData() POSTStripePaymentsRequestData {
	if o == nil {
		var ret POSTStripePaymentsRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *POSTStripePaymentsRequest) GetDataOk() (*POSTStripePaymentsRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *POSTStripePaymentsRequest) SetData(v POSTStripePaymentsRequestData) {
	o.Data = v
}

func (o POSTStripePaymentsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTStripePaymentsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullablePOSTStripePaymentsRequest struct {
	value *POSTStripePaymentsRequest
	isSet bool
}

func (v NullablePOSTStripePaymentsRequest) Get() *POSTStripePaymentsRequest {
	return v.value
}

func (v *NullablePOSTStripePaymentsRequest) Set(val *POSTStripePaymentsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTStripePaymentsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTStripePaymentsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTStripePaymentsRequest(val *POSTStripePaymentsRequest) *NullablePOSTStripePaymentsRequest {
	return &NullablePOSTStripePaymentsRequest{value: val, isSet: true}
}

func (v NullablePOSTStripePaymentsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTStripePaymentsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
