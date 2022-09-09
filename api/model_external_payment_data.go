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

// ExternalPaymentData struct for ExternalPaymentData
type ExternalPaymentData struct {
	// The resource's type
	Type string `json:"type"`
	Attributes GETExternalPayments200ResponseDataInnerAttributes `json:"attributes"`
	Relationships *ExternalPaymentDataRelationships `json:"relationships,omitempty"`
}

// NewExternalPaymentData instantiates a new ExternalPaymentData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalPaymentData(type_ string, attributes GETExternalPayments200ResponseDataInnerAttributes) *ExternalPaymentData {
	this := ExternalPaymentData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewExternalPaymentDataWithDefaults instantiates a new ExternalPaymentData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalPaymentDataWithDefaults() *ExternalPaymentData {
	this := ExternalPaymentData{}
	var type_ string = "external_payments"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *ExternalPaymentData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ExternalPaymentData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ExternalPaymentData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *ExternalPaymentData) GetAttributes() GETExternalPayments200ResponseDataInnerAttributes {
	if o == nil {
		var ret GETExternalPayments200ResponseDataInnerAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ExternalPaymentData) GetAttributesOk() (*GETExternalPayments200ResponseDataInnerAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *ExternalPaymentData) SetAttributes(v GETExternalPayments200ResponseDataInnerAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ExternalPaymentData) GetRelationships() ExternalPaymentDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret ExternalPaymentDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalPaymentData) GetRelationshipsOk() (*ExternalPaymentDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ExternalPaymentData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given ExternalPaymentDataRelationships and assigns it to the Relationships field.
func (o *ExternalPaymentData) SetRelationships(v ExternalPaymentDataRelationships) {
	o.Relationships = &v
}

func (o ExternalPaymentData) MarshalJSON() ([]byte, error) {
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

type NullableExternalPaymentData struct {
	value *ExternalPaymentData
	isSet bool
}

func (v NullableExternalPaymentData) Get() *ExternalPaymentData {
	return v.value
}

func (v *NullableExternalPaymentData) Set(val *ExternalPaymentData) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalPaymentData) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalPaymentData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalPaymentData(val *ExternalPaymentData) *NullableExternalPaymentData {
	return &NullableExternalPaymentData{value: val, isSet: true}
}

func (v NullableExternalPaymentData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalPaymentData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


