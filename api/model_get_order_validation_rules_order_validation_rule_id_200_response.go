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

// GETOrderValidationRulesOrderValidationRuleId200Response struct for GETOrderValidationRulesOrderValidationRuleId200Response
type GETOrderValidationRulesOrderValidationRuleId200Response struct {
	Data *GETOrderValidationRules200ResponseDataInner `json:"data,omitempty"`
}

// NewGETOrderValidationRulesOrderValidationRuleId200Response instantiates a new GETOrderValidationRulesOrderValidationRuleId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETOrderValidationRulesOrderValidationRuleId200Response() *GETOrderValidationRulesOrderValidationRuleId200Response {
	this := GETOrderValidationRulesOrderValidationRuleId200Response{}
	return &this
}

// NewGETOrderValidationRulesOrderValidationRuleId200ResponseWithDefaults instantiates a new GETOrderValidationRulesOrderValidationRuleId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETOrderValidationRulesOrderValidationRuleId200ResponseWithDefaults() *GETOrderValidationRulesOrderValidationRuleId200Response {
	this := GETOrderValidationRulesOrderValidationRuleId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETOrderValidationRulesOrderValidationRuleId200Response) GetData() GETOrderValidationRules200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret GETOrderValidationRules200ResponseDataInner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrderValidationRulesOrderValidationRuleId200Response) GetDataOk() (*GETOrderValidationRules200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETOrderValidationRulesOrderValidationRuleId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETOrderValidationRules200ResponseDataInner and assigns it to the Data field.
func (o *GETOrderValidationRulesOrderValidationRuleId200Response) SetData(v GETOrderValidationRules200ResponseDataInner) {
	o.Data = &v
}

func (o GETOrderValidationRulesOrderValidationRuleId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETOrderValidationRulesOrderValidationRuleId200Response struct {
	value *GETOrderValidationRulesOrderValidationRuleId200Response
	isSet bool
}

func (v NullableGETOrderValidationRulesOrderValidationRuleId200Response) Get() *GETOrderValidationRulesOrderValidationRuleId200Response {
	return v.value
}

func (v *NullableGETOrderValidationRulesOrderValidationRuleId200Response) Set(val *GETOrderValidationRulesOrderValidationRuleId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETOrderValidationRulesOrderValidationRuleId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETOrderValidationRulesOrderValidationRuleId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETOrderValidationRulesOrderValidationRuleId200Response(val *GETOrderValidationRulesOrderValidationRuleId200Response) *NullableGETOrderValidationRulesOrderValidationRuleId200Response {
	return &NullableGETOrderValidationRulesOrderValidationRuleId200Response{value: val, isSet: true}
}

func (v NullableGETOrderValidationRulesOrderValidationRuleId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETOrderValidationRulesOrderValidationRuleId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
