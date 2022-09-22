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

// CustomerPaymentSourceCreateData struct for CustomerPaymentSourceCreateData
type CustomerPaymentSourceCreateData struct {
	// The resource's type
	Type          string                                        `json:"type"`
	Attributes    POSTAdyenPayments201ResponseDataAttributes    `json:"attributes"`
	Relationships *CustomerPaymentSourceCreateDataRelationships `json:"relationships,omitempty"`
}

// NewCustomerPaymentSourceCreateData instantiates a new CustomerPaymentSourceCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerPaymentSourceCreateData(type_ string, attributes POSTAdyenPayments201ResponseDataAttributes) *CustomerPaymentSourceCreateData {
	this := CustomerPaymentSourceCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewCustomerPaymentSourceCreateDataWithDefaults instantiates a new CustomerPaymentSourceCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerPaymentSourceCreateDataWithDefaults() *CustomerPaymentSourceCreateData {
	this := CustomerPaymentSourceCreateData{}
	return &this
}

// GetType returns the Type field value
func (o *CustomerPaymentSourceCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CustomerPaymentSourceCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CustomerPaymentSourceCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *CustomerPaymentSourceCreateData) GetAttributes() POSTAdyenPayments201ResponseDataAttributes {
	if o == nil {
		var ret POSTAdyenPayments201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CustomerPaymentSourceCreateData) GetAttributesOk() (*POSTAdyenPayments201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CustomerPaymentSourceCreateData) SetAttributes(v POSTAdyenPayments201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *CustomerPaymentSourceCreateData) GetRelationships() CustomerPaymentSourceCreateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret CustomerPaymentSourceCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerPaymentSourceCreateData) GetRelationshipsOk() (*CustomerPaymentSourceCreateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CustomerPaymentSourceCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given CustomerPaymentSourceCreateDataRelationships and assigns it to the Relationships field.
func (o *CustomerPaymentSourceCreateData) SetRelationships(v CustomerPaymentSourceCreateDataRelationships) {
	o.Relationships = &v
}

func (o CustomerPaymentSourceCreateData) MarshalJSON() ([]byte, error) {
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

type NullableCustomerPaymentSourceCreateData struct {
	value *CustomerPaymentSourceCreateData
	isSet bool
}

func (v NullableCustomerPaymentSourceCreateData) Get() *CustomerPaymentSourceCreateData {
	return v.value
}

func (v *NullableCustomerPaymentSourceCreateData) Set(val *CustomerPaymentSourceCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerPaymentSourceCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerPaymentSourceCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerPaymentSourceCreateData(val *CustomerPaymentSourceCreateData) *NullableCustomerPaymentSourceCreateData {
	return &NullableCustomerPaymentSourceCreateData{value: val, isSet: true}
}

func (v NullableCustomerPaymentSourceCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerPaymentSourceCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
