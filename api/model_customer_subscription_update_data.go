/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.4.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// CustomerSubscriptionUpdateData struct for CustomerSubscriptionUpdateData
type CustomerSubscriptionUpdateData struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id            string                                     `json:"id"`
	Attributes    POSTAdyenPayments201ResponseDataAttributes `json:"attributes"`
	Relationships map[string]interface{}                     `json:"relationships,omitempty"`
}

// NewCustomerSubscriptionUpdateData instantiates a new CustomerSubscriptionUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerSubscriptionUpdateData(type_ string, id string, attributes POSTAdyenPayments201ResponseDataAttributes) *CustomerSubscriptionUpdateData {
	this := CustomerSubscriptionUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewCustomerSubscriptionUpdateDataWithDefaults instantiates a new CustomerSubscriptionUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerSubscriptionUpdateDataWithDefaults() *CustomerSubscriptionUpdateData {
	this := CustomerSubscriptionUpdateData{}
	return &this
}

// GetType returns the Type field value
func (o *CustomerSubscriptionUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CustomerSubscriptionUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CustomerSubscriptionUpdateData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *CustomerSubscriptionUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CustomerSubscriptionUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CustomerSubscriptionUpdateData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *CustomerSubscriptionUpdateData) GetAttributes() POSTAdyenPayments201ResponseDataAttributes {
	if o == nil {
		var ret POSTAdyenPayments201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CustomerSubscriptionUpdateData) GetAttributesOk() (*POSTAdyenPayments201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CustomerSubscriptionUpdateData) SetAttributes(v POSTAdyenPayments201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *CustomerSubscriptionUpdateData) GetRelationships() map[string]interface{} {
	if o == nil || o.Relationships == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSubscriptionUpdateData) GetRelationshipsOk() (map[string]interface{}, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CustomerSubscriptionUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]interface{} and assigns it to the Relationships field.
func (o *CustomerSubscriptionUpdateData) SetRelationships(v map[string]interface{}) {
	o.Relationships = v
}

func (o CustomerSubscriptionUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerSubscriptionUpdateData struct {
	value *CustomerSubscriptionUpdateData
	isSet bool
}

func (v NullableCustomerSubscriptionUpdateData) Get() *CustomerSubscriptionUpdateData {
	return v.value
}

func (v *NullableCustomerSubscriptionUpdateData) Set(val *CustomerSubscriptionUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerSubscriptionUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerSubscriptionUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerSubscriptionUpdateData(val *CustomerSubscriptionUpdateData) *NullableCustomerSubscriptionUpdateData {
	return &NullableCustomerSubscriptionUpdateData{value: val, isSet: true}
}

func (v NullableCustomerSubscriptionUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerSubscriptionUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
