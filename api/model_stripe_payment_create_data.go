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

// StripePaymentCreateData struct for StripePaymentCreateData
type StripePaymentCreateData struct {
	// The resource's type
	Type          string                                      `json:"type"`
	Attributes    POSTStripePayments201ResponseDataAttributes `json:"attributes"`
	Relationships *AdyenPaymentCreateDataRelationships        `json:"relationships,omitempty"`
}

// NewStripePaymentCreateData instantiates a new StripePaymentCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStripePaymentCreateData(type_ string, attributes POSTStripePayments201ResponseDataAttributes) *StripePaymentCreateData {
	this := StripePaymentCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewStripePaymentCreateDataWithDefaults instantiates a new StripePaymentCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStripePaymentCreateDataWithDefaults() *StripePaymentCreateData {
	this := StripePaymentCreateData{}
	return &this
}

// GetType returns the Type field value
func (o *StripePaymentCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StripePaymentCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *StripePaymentCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *StripePaymentCreateData) GetAttributes() POSTStripePayments201ResponseDataAttributes {
	if o == nil {
		var ret POSTStripePayments201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *StripePaymentCreateData) GetAttributesOk() (*POSTStripePayments201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *StripePaymentCreateData) SetAttributes(v POSTStripePayments201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *StripePaymentCreateData) GetRelationships() AdyenPaymentCreateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret AdyenPaymentCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripePaymentCreateData) GetRelationshipsOk() (*AdyenPaymentCreateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *StripePaymentCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AdyenPaymentCreateDataRelationships and assigns it to the Relationships field.
func (o *StripePaymentCreateData) SetRelationships(v AdyenPaymentCreateDataRelationships) {
	o.Relationships = &v
}

func (o StripePaymentCreateData) MarshalJSON() ([]byte, error) {
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

type NullableStripePaymentCreateData struct {
	value *StripePaymentCreateData
	isSet bool
}

func (v NullableStripePaymentCreateData) Get() *StripePaymentCreateData {
	return v.value
}

func (v *NullableStripePaymentCreateData) Set(val *StripePaymentCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableStripePaymentCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableStripePaymentCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStripePaymentCreateData(val *StripePaymentCreateData) *NullableStripePaymentCreateData {
	return &NullableStripePaymentCreateData{value: val, isSet: true}
}

func (v NullableStripePaymentCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStripePaymentCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
