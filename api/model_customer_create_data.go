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

// CustomerCreateData struct for CustomerCreateData
type CustomerCreateData struct {
	// The resource's type
	Type          string                                 `json:"type"`
	Attributes    POSTCustomers201ResponseDataAttributes `json:"attributes"`
	Relationships *CustomerCreateDataRelationships       `json:"relationships,omitempty"`
}

// NewCustomerCreateData instantiates a new CustomerCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerCreateData(type_ string, attributes POSTCustomers201ResponseDataAttributes) *CustomerCreateData {
	this := CustomerCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewCustomerCreateDataWithDefaults instantiates a new CustomerCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerCreateDataWithDefaults() *CustomerCreateData {
	this := CustomerCreateData{}
	return &this
}

// GetType returns the Type field value
func (o *CustomerCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CustomerCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CustomerCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *CustomerCreateData) GetAttributes() POSTCustomers201ResponseDataAttributes {
	if o == nil {
		var ret POSTCustomers201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CustomerCreateData) GetAttributesOk() (*POSTCustomers201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CustomerCreateData) SetAttributes(v POSTCustomers201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *CustomerCreateData) GetRelationships() CustomerCreateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret CustomerCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCreateData) GetRelationshipsOk() (*CustomerCreateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CustomerCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given CustomerCreateDataRelationships and assigns it to the Relationships field.
func (o *CustomerCreateData) SetRelationships(v CustomerCreateDataRelationships) {
	o.Relationships = &v
}

func (o CustomerCreateData) MarshalJSON() ([]byte, error) {
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

type NullableCustomerCreateData struct {
	value *CustomerCreateData
	isSet bool
}

func (v NullableCustomerCreateData) Get() *CustomerCreateData {
	return v.value
}

func (v *NullableCustomerCreateData) Set(val *CustomerCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerCreateData(val *CustomerCreateData) *NullableCustomerCreateData {
	return &NullableCustomerCreateData{value: val, isSet: true}
}

func (v NullableCustomerCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
