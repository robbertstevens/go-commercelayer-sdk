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

// checks if the POSTBillingInfoValidationRulesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTBillingInfoValidationRulesRequest{}

// POSTBillingInfoValidationRulesRequest struct for POSTBillingInfoValidationRulesRequest
type POSTBillingInfoValidationRulesRequest struct {
	Data POSTBillingInfoValidationRulesRequestData `json:"data"`
}

// NewPOSTBillingInfoValidationRulesRequest instantiates a new POSTBillingInfoValidationRulesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTBillingInfoValidationRulesRequest(data POSTBillingInfoValidationRulesRequestData) *POSTBillingInfoValidationRulesRequest {
	this := POSTBillingInfoValidationRulesRequest{}
	this.Data = data
	return &this
}

// NewPOSTBillingInfoValidationRulesRequestWithDefaults instantiates a new POSTBillingInfoValidationRulesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTBillingInfoValidationRulesRequestWithDefaults() *POSTBillingInfoValidationRulesRequest {
	this := POSTBillingInfoValidationRulesRequest{}
	return &this
}

// GetData returns the Data field value
func (o *POSTBillingInfoValidationRulesRequest) GetData() POSTBillingInfoValidationRulesRequestData {
	if o == nil {
		var ret POSTBillingInfoValidationRulesRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *POSTBillingInfoValidationRulesRequest) GetDataOk() (*POSTBillingInfoValidationRulesRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *POSTBillingInfoValidationRulesRequest) SetData(v POSTBillingInfoValidationRulesRequestData) {
	o.Data = v
}

func (o POSTBillingInfoValidationRulesRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTBillingInfoValidationRulesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullablePOSTBillingInfoValidationRulesRequest struct {
	value *POSTBillingInfoValidationRulesRequest
	isSet bool
}

func (v NullablePOSTBillingInfoValidationRulesRequest) Get() *POSTBillingInfoValidationRulesRequest {
	return v.value
}

func (v *NullablePOSTBillingInfoValidationRulesRequest) Set(val *POSTBillingInfoValidationRulesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTBillingInfoValidationRulesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTBillingInfoValidationRulesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTBillingInfoValidationRulesRequest(val *POSTBillingInfoValidationRulesRequest) *NullablePOSTBillingInfoValidationRulesRequest {
	return &NullablePOSTBillingInfoValidationRulesRequest{value: val, isSet: true}
}

func (v NullablePOSTBillingInfoValidationRulesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTBillingInfoValidationRulesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
