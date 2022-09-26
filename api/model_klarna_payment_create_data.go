/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// KlarnaPaymentCreateData struct for KlarnaPaymentCreateData
type KlarnaPaymentCreateData struct {
	// The resource's type
	Type          string                                     `json:"type"`
	Attributes    POSTAdyenPayments201ResponseDataAttributes `json:"attributes"`
	Relationships *AdyenPaymentCreateDataRelationships       `json:"relationships,omitempty"`
}

// NewKlarnaPaymentCreateData instantiates a new KlarnaPaymentCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKlarnaPaymentCreateData(type_ string, attributes POSTAdyenPayments201ResponseDataAttributes) *KlarnaPaymentCreateData {
	this := KlarnaPaymentCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewKlarnaPaymentCreateDataWithDefaults instantiates a new KlarnaPaymentCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKlarnaPaymentCreateDataWithDefaults() *KlarnaPaymentCreateData {
	this := KlarnaPaymentCreateData{}
	return &this
}

// GetType returns the Type field value
func (o *KlarnaPaymentCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *KlarnaPaymentCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *KlarnaPaymentCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *KlarnaPaymentCreateData) GetAttributes() POSTAdyenPayments201ResponseDataAttributes {
	if o == nil {
		var ret POSTAdyenPayments201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *KlarnaPaymentCreateData) GetAttributesOk() (*POSTAdyenPayments201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *KlarnaPaymentCreateData) SetAttributes(v POSTAdyenPayments201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *KlarnaPaymentCreateData) GetRelationships() AdyenPaymentCreateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret AdyenPaymentCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlarnaPaymentCreateData) GetRelationshipsOk() (*AdyenPaymentCreateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *KlarnaPaymentCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AdyenPaymentCreateDataRelationships and assigns it to the Relationships field.
func (o *KlarnaPaymentCreateData) SetRelationships(v AdyenPaymentCreateDataRelationships) {
	o.Relationships = &v
}

func (o KlarnaPaymentCreateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableKlarnaPaymentCreateData struct {
	value *KlarnaPaymentCreateData
	isSet bool
}

func (v NullableKlarnaPaymentCreateData) Get() *KlarnaPaymentCreateData {
	return v.value
}

func (v *NullableKlarnaPaymentCreateData) Set(val *KlarnaPaymentCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableKlarnaPaymentCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableKlarnaPaymentCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKlarnaPaymentCreateData(val *KlarnaPaymentCreateData) *NullableKlarnaPaymentCreateData {
	return &NullableKlarnaPaymentCreateData{value: val, isSet: true}
}

func (v NullableKlarnaPaymentCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKlarnaPaymentCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
