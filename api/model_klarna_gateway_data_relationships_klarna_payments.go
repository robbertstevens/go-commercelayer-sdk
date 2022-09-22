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

// KlarnaGatewayDataRelationshipsKlarnaPayments struct for KlarnaGatewayDataRelationshipsKlarnaPayments
type KlarnaGatewayDataRelationshipsKlarnaPayments struct {
	Data KlarnaGatewayDataRelationshipsKlarnaPaymentsData `json:"data"`
}

// NewKlarnaGatewayDataRelationshipsKlarnaPayments instantiates a new KlarnaGatewayDataRelationshipsKlarnaPayments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKlarnaGatewayDataRelationshipsKlarnaPayments(data KlarnaGatewayDataRelationshipsKlarnaPaymentsData) *KlarnaGatewayDataRelationshipsKlarnaPayments {
	this := KlarnaGatewayDataRelationshipsKlarnaPayments{}
	this.Data = data
	return &this
}

// NewKlarnaGatewayDataRelationshipsKlarnaPaymentsWithDefaults instantiates a new KlarnaGatewayDataRelationshipsKlarnaPayments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKlarnaGatewayDataRelationshipsKlarnaPaymentsWithDefaults() *KlarnaGatewayDataRelationshipsKlarnaPayments {
	this := KlarnaGatewayDataRelationshipsKlarnaPayments{}
	return &this
}

// GetData returns the Data field value
func (o *KlarnaGatewayDataRelationshipsKlarnaPayments) GetData() KlarnaGatewayDataRelationshipsKlarnaPaymentsData {
	if o == nil {
		var ret KlarnaGatewayDataRelationshipsKlarnaPaymentsData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *KlarnaGatewayDataRelationshipsKlarnaPayments) GetDataOk() (*KlarnaGatewayDataRelationshipsKlarnaPaymentsData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *KlarnaGatewayDataRelationshipsKlarnaPayments) SetData(v KlarnaGatewayDataRelationshipsKlarnaPaymentsData) {
	o.Data = v
}

func (o KlarnaGatewayDataRelationshipsKlarnaPayments) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableKlarnaGatewayDataRelationshipsKlarnaPayments struct {
	value *KlarnaGatewayDataRelationshipsKlarnaPayments
	isSet bool
}

func (v NullableKlarnaGatewayDataRelationshipsKlarnaPayments) Get() *KlarnaGatewayDataRelationshipsKlarnaPayments {
	return v.value
}

func (v *NullableKlarnaGatewayDataRelationshipsKlarnaPayments) Set(val *KlarnaGatewayDataRelationshipsKlarnaPayments) {
	v.value = val
	v.isSet = true
}

func (v NullableKlarnaGatewayDataRelationshipsKlarnaPayments) IsSet() bool {
	return v.isSet
}

func (v *NullableKlarnaGatewayDataRelationshipsKlarnaPayments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKlarnaGatewayDataRelationshipsKlarnaPayments(val *KlarnaGatewayDataRelationshipsKlarnaPayments) *NullableKlarnaGatewayDataRelationshipsKlarnaPayments {
	return &NullableKlarnaGatewayDataRelationshipsKlarnaPayments{value: val, isSet: true}
}

func (v NullableKlarnaGatewayDataRelationshipsKlarnaPayments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKlarnaGatewayDataRelationshipsKlarnaPayments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
