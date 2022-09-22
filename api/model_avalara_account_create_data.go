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

// AvalaraAccountCreateData struct for AvalaraAccountCreateData
type AvalaraAccountCreateData struct {
	// The resource's type
	Type          string                                       `json:"type"`
	Attributes    POSTAvalaraAccounts201ResponseDataAttributes `json:"attributes"`
	Relationships *AvalaraAccountCreateDataRelationships       `json:"relationships,omitempty"`
}

// NewAvalaraAccountCreateData instantiates a new AvalaraAccountCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvalaraAccountCreateData(type_ string, attributes POSTAvalaraAccounts201ResponseDataAttributes) *AvalaraAccountCreateData {
	this := AvalaraAccountCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewAvalaraAccountCreateDataWithDefaults instantiates a new AvalaraAccountCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvalaraAccountCreateDataWithDefaults() *AvalaraAccountCreateData {
	this := AvalaraAccountCreateData{}
	return &this
}

// GetType returns the Type field value
func (o *AvalaraAccountCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AvalaraAccountCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AvalaraAccountCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *AvalaraAccountCreateData) GetAttributes() POSTAvalaraAccounts201ResponseDataAttributes {
	if o == nil {
		var ret POSTAvalaraAccounts201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AvalaraAccountCreateData) GetAttributesOk() (*POSTAvalaraAccounts201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *AvalaraAccountCreateData) SetAttributes(v POSTAvalaraAccounts201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AvalaraAccountCreateData) GetRelationships() AvalaraAccountCreateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret AvalaraAccountCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvalaraAccountCreateData) GetRelationshipsOk() (*AvalaraAccountCreateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AvalaraAccountCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AvalaraAccountCreateDataRelationships and assigns it to the Relationships field.
func (o *AvalaraAccountCreateData) SetRelationships(v AvalaraAccountCreateDataRelationships) {
	o.Relationships = &v
}

func (o AvalaraAccountCreateData) MarshalJSON() ([]byte, error) {
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

type NullableAvalaraAccountCreateData struct {
	value *AvalaraAccountCreateData
	isSet bool
}

func (v NullableAvalaraAccountCreateData) Get() *AvalaraAccountCreateData {
	return v.value
}

func (v *NullableAvalaraAccountCreateData) Set(val *AvalaraAccountCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableAvalaraAccountCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableAvalaraAccountCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvalaraAccountCreateData(val *AvalaraAccountCreateData) *NullableAvalaraAccountCreateData {
	return &NullableAvalaraAccountCreateData{value: val, isSet: true}
}

func (v NullableAvalaraAccountCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvalaraAccountCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
