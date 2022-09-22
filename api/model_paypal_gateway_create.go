/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaypalGatewayCreate struct for PaypalGatewayCreate
type PaypalGatewayCreate struct {
	Data PaypalGatewayCreateData `json:"data"`
}

// NewPaypalGatewayCreate instantiates a new PaypalGatewayCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaypalGatewayCreate(data PaypalGatewayCreateData) *PaypalGatewayCreate {
	this := PaypalGatewayCreate{}
	this.Data = data
	return &this
}

// NewPaypalGatewayCreateWithDefaults instantiates a new PaypalGatewayCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaypalGatewayCreateWithDefaults() *PaypalGatewayCreate {
	this := PaypalGatewayCreate{}
	return &this
}

// GetData returns the Data field value
func (o *PaypalGatewayCreate) GetData() PaypalGatewayCreateData {
	if o == nil {
		var ret PaypalGatewayCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PaypalGatewayCreate) GetDataOk() (*PaypalGatewayCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PaypalGatewayCreate) SetData(v PaypalGatewayCreateData) {
	o.Data = v
}

func (o PaypalGatewayCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePaypalGatewayCreate struct {
	value *PaypalGatewayCreate
	isSet bool
}

func (v NullablePaypalGatewayCreate) Get() *PaypalGatewayCreate {
	return v.value
}

func (v *NullablePaypalGatewayCreate) Set(val *PaypalGatewayCreate) {
	v.value = val
	v.isSet = true
}

func (v NullablePaypalGatewayCreate) IsSet() bool {
	return v.isSet
}

func (v *NullablePaypalGatewayCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaypalGatewayCreate(val *PaypalGatewayCreate) *NullablePaypalGatewayCreate {
	return &NullablePaypalGatewayCreate{value: val, isSet: true}
}

func (v NullablePaypalGatewayCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaypalGatewayCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
