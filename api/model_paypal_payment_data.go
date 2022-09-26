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

// PaypalPaymentData struct for PaypalPaymentData
type PaypalPaymentData struct {
	// The resource's type
	Type          string                                          `json:"type"`
	Attributes    GETPaypalPayments200ResponseDataInnerAttributes `json:"attributes"`
	Relationships *AdyenPaymentDataRelationships                  `json:"relationships,omitempty"`
}

// NewPaypalPaymentData instantiates a new PaypalPaymentData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaypalPaymentData(type_ string, attributes GETPaypalPayments200ResponseDataInnerAttributes) *PaypalPaymentData {
	this := PaypalPaymentData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewPaypalPaymentDataWithDefaults instantiates a new PaypalPaymentData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaypalPaymentDataWithDefaults() *PaypalPaymentData {
	this := PaypalPaymentData{}
	return &this
}

// GetType returns the Type field value
func (o *PaypalPaymentData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PaypalPaymentData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PaypalPaymentData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *PaypalPaymentData) GetAttributes() GETPaypalPayments200ResponseDataInnerAttributes {
	if o == nil {
		var ret GETPaypalPayments200ResponseDataInnerAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *PaypalPaymentData) GetAttributesOk() (*GETPaypalPayments200ResponseDataInnerAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *PaypalPaymentData) SetAttributes(v GETPaypalPayments200ResponseDataInnerAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PaypalPaymentData) GetRelationships() AdyenPaymentDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret AdyenPaymentDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaypalPaymentData) GetRelationshipsOk() (*AdyenPaymentDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PaypalPaymentData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AdyenPaymentDataRelationships and assigns it to the Relationships field.
func (o *PaypalPaymentData) SetRelationships(v AdyenPaymentDataRelationships) {
	o.Relationships = &v
}

func (o PaypalPaymentData) MarshalJSON() ([]byte, error) {
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

type NullablePaypalPaymentData struct {
	value *PaypalPaymentData
	isSet bool
}

func (v NullablePaypalPaymentData) Get() *PaypalPaymentData {
	return v.value
}

func (v *NullablePaypalPaymentData) Set(val *PaypalPaymentData) {
	v.value = val
	v.isSet = true
}

func (v NullablePaypalPaymentData) IsSet() bool {
	return v.isSet
}

func (v *NullablePaypalPaymentData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaypalPaymentData(val *PaypalPaymentData) *NullablePaypalPaymentData {
	return &NullablePaypalPaymentData{value: val, isSet: true}
}

func (v NullablePaypalPaymentData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaypalPaymentData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
