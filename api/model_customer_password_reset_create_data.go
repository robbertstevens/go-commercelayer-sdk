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

// CustomerPasswordResetCreateData struct for CustomerPasswordResetCreateData
type CustomerPasswordResetCreateData struct {
	// The resource's type
	Type string `json:"type"`
	Attributes POSTCustomerPasswordResets201ResponseDataAttributes `json:"attributes"`
	Relationships map[string]interface{} `json:"relationships,omitempty"`
}

// NewCustomerPasswordResetCreateData instantiates a new CustomerPasswordResetCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerPasswordResetCreateData(type_ string, attributes POSTCustomerPasswordResets201ResponseDataAttributes) *CustomerPasswordResetCreateData {
	this := CustomerPasswordResetCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewCustomerPasswordResetCreateDataWithDefaults instantiates a new CustomerPasswordResetCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerPasswordResetCreateDataWithDefaults() *CustomerPasswordResetCreateData {
	this := CustomerPasswordResetCreateData{}
	var type_ string = "customer_password_resets"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *CustomerPasswordResetCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CustomerPasswordResetCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CustomerPasswordResetCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *CustomerPasswordResetCreateData) GetAttributes() POSTCustomerPasswordResets201ResponseDataAttributes {
	if o == nil {
		var ret POSTCustomerPasswordResets201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CustomerPasswordResetCreateData) GetAttributesOk() (*POSTCustomerPasswordResets201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CustomerPasswordResetCreateData) SetAttributes(v POSTCustomerPasswordResets201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *CustomerPasswordResetCreateData) GetRelationships() map[string]interface{} {
	if o == nil || o.Relationships == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerPasswordResetCreateData) GetRelationshipsOk() (map[string]interface{}, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CustomerPasswordResetCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]interface{} and assigns it to the Relationships field.
func (o *CustomerPasswordResetCreateData) SetRelationships(v map[string]interface{}) {
	o.Relationships = v
}

func (o CustomerPasswordResetCreateData) MarshalJSON() ([]byte, error) {
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

type NullableCustomerPasswordResetCreateData struct {
	value *CustomerPasswordResetCreateData
	isSet bool
}

func (v NullableCustomerPasswordResetCreateData) Get() *CustomerPasswordResetCreateData {
	return v.value
}

func (v *NullableCustomerPasswordResetCreateData) Set(val *CustomerPasswordResetCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerPasswordResetCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerPasswordResetCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerPasswordResetCreateData(val *CustomerPasswordResetCreateData) *NullableCustomerPasswordResetCreateData {
	return &NullableCustomerPasswordResetCreateData{value: val, isSet: true}
}

func (v NullableCustomerPasswordResetCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerPasswordResetCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


