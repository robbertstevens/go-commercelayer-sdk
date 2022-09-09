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

// BraintreePaymentData struct for BraintreePaymentData
type BraintreePaymentData struct {
	// The resource's type
	Type string `json:"type"`
	Attributes GETBraintreePayments200ResponseDataInnerAttributes `json:"attributes"`
	Relationships *AdyenPaymentDataRelationships `json:"relationships,omitempty"`
}

// NewBraintreePaymentData instantiates a new BraintreePaymentData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBraintreePaymentData(type_ string, attributes GETBraintreePayments200ResponseDataInnerAttributes) *BraintreePaymentData {
	this := BraintreePaymentData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewBraintreePaymentDataWithDefaults instantiates a new BraintreePaymentData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBraintreePaymentDataWithDefaults() *BraintreePaymentData {
	this := BraintreePaymentData{}
	var type_ string = "braintree_payments"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *BraintreePaymentData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BraintreePaymentData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BraintreePaymentData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *BraintreePaymentData) GetAttributes() GETBraintreePayments200ResponseDataInnerAttributes {
	if o == nil {
		var ret GETBraintreePayments200ResponseDataInnerAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *BraintreePaymentData) GetAttributesOk() (*GETBraintreePayments200ResponseDataInnerAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *BraintreePaymentData) SetAttributes(v GETBraintreePayments200ResponseDataInnerAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *BraintreePaymentData) GetRelationships() AdyenPaymentDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret AdyenPaymentDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BraintreePaymentData) GetRelationshipsOk() (*AdyenPaymentDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *BraintreePaymentData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AdyenPaymentDataRelationships and assigns it to the Relationships field.
func (o *BraintreePaymentData) SetRelationships(v AdyenPaymentDataRelationships) {
	o.Relationships = &v
}

func (o BraintreePaymentData) MarshalJSON() ([]byte, error) {
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

type NullableBraintreePaymentData struct {
	value *BraintreePaymentData
	isSet bool
}

func (v NullableBraintreePaymentData) Get() *BraintreePaymentData {
	return v.value
}

func (v *NullableBraintreePaymentData) Set(val *BraintreePaymentData) {
	v.value = val
	v.isSet = true
}

func (v NullableBraintreePaymentData) IsSet() bool {
	return v.isSet
}

func (v *NullableBraintreePaymentData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBraintreePaymentData(val *BraintreePaymentData) *NullableBraintreePaymentData {
	return &NullableBraintreePaymentData{value: val, isSet: true}
}

func (v NullableBraintreePaymentData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBraintreePaymentData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


