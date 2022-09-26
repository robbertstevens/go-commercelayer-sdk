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

// CustomerAddressCreateData struct for CustomerAddressCreateData
type CustomerAddressCreateData struct {
	// The resource's type
	Type          string                                     `json:"type"`
	Attributes    POSTAdyenPayments201ResponseDataAttributes `json:"attributes"`
	Relationships *CustomerAddressCreateDataRelationships    `json:"relationships,omitempty"`
}

// NewCustomerAddressCreateData instantiates a new CustomerAddressCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerAddressCreateData(type_ string, attributes POSTAdyenPayments201ResponseDataAttributes) *CustomerAddressCreateData {
	this := CustomerAddressCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewCustomerAddressCreateDataWithDefaults instantiates a new CustomerAddressCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerAddressCreateDataWithDefaults() *CustomerAddressCreateData {
	this := CustomerAddressCreateData{}
	return &this
}

// GetType returns the Type field value
func (o *CustomerAddressCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CustomerAddressCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CustomerAddressCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *CustomerAddressCreateData) GetAttributes() POSTAdyenPayments201ResponseDataAttributes {
	if o == nil {
		var ret POSTAdyenPayments201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CustomerAddressCreateData) GetAttributesOk() (*POSTAdyenPayments201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CustomerAddressCreateData) SetAttributes(v POSTAdyenPayments201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *CustomerAddressCreateData) GetRelationships() CustomerAddressCreateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret CustomerAddressCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressCreateData) GetRelationshipsOk() (*CustomerAddressCreateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CustomerAddressCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given CustomerAddressCreateDataRelationships and assigns it to the Relationships field.
func (o *CustomerAddressCreateData) SetRelationships(v CustomerAddressCreateDataRelationships) {
	o.Relationships = &v
}

func (o CustomerAddressCreateData) MarshalJSON() ([]byte, error) {
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

type NullableCustomerAddressCreateData struct {
	value *CustomerAddressCreateData
	isSet bool
}

func (v NullableCustomerAddressCreateData) Get() *CustomerAddressCreateData {
	return v.value
}

func (v *NullableCustomerAddressCreateData) Set(val *CustomerAddressCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerAddressCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerAddressCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerAddressCreateData(val *CustomerAddressCreateData) *NullableCustomerAddressCreateData {
	return &NullableCustomerAddressCreateData{value: val, isSet: true}
}

func (v NullableCustomerAddressCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerAddressCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
