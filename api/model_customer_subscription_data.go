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

// CustomerSubscriptionData struct for CustomerSubscriptionData
type CustomerSubscriptionData struct {
	// The resource's type
	Type          string                                                 `json:"type"`
	Attributes    GETCustomerSubscriptions200ResponseDataInnerAttributes `json:"attributes"`
	Relationships *CustomerPasswordResetDataRelationships                `json:"relationships,omitempty"`
}

// NewCustomerSubscriptionData instantiates a new CustomerSubscriptionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerSubscriptionData(type_ string, attributes GETCustomerSubscriptions200ResponseDataInnerAttributes) *CustomerSubscriptionData {
	this := CustomerSubscriptionData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewCustomerSubscriptionDataWithDefaults instantiates a new CustomerSubscriptionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerSubscriptionDataWithDefaults() *CustomerSubscriptionData {
	this := CustomerSubscriptionData{}
	return &this
}

// GetType returns the Type field value
func (o *CustomerSubscriptionData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CustomerSubscriptionData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CustomerSubscriptionData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *CustomerSubscriptionData) GetAttributes() GETCustomerSubscriptions200ResponseDataInnerAttributes {
	if o == nil {
		var ret GETCustomerSubscriptions200ResponseDataInnerAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CustomerSubscriptionData) GetAttributesOk() (*GETCustomerSubscriptions200ResponseDataInnerAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CustomerSubscriptionData) SetAttributes(v GETCustomerSubscriptions200ResponseDataInnerAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *CustomerSubscriptionData) GetRelationships() CustomerPasswordResetDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret CustomerPasswordResetDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSubscriptionData) GetRelationshipsOk() (*CustomerPasswordResetDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CustomerSubscriptionData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given CustomerPasswordResetDataRelationships and assigns it to the Relationships field.
func (o *CustomerSubscriptionData) SetRelationships(v CustomerPasswordResetDataRelationships) {
	o.Relationships = &v
}

func (o CustomerSubscriptionData) MarshalJSON() ([]byte, error) {
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

type NullableCustomerSubscriptionData struct {
	value *CustomerSubscriptionData
	isSet bool
}

func (v NullableCustomerSubscriptionData) Get() *CustomerSubscriptionData {
	return v.value
}

func (v *NullableCustomerSubscriptionData) Set(val *CustomerSubscriptionData) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerSubscriptionData) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerSubscriptionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerSubscriptionData(val *CustomerSubscriptionData) *NullableCustomerSubscriptionData {
	return &NullableCustomerSubscriptionData{value: val, isSet: true}
}

func (v NullableCustomerSubscriptionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerSubscriptionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
