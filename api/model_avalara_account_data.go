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

// AvalaraAccountData struct for AvalaraAccountData
type AvalaraAccountData struct {
	// The resource's type
	Type          string                                           `json:"type"`
	Attributes    GETAvalaraAccounts200ResponseDataInnerAttributes `json:"attributes"`
	Relationships *AvalaraAccountDataRelationships                 `json:"relationships,omitempty"`
}

// NewAvalaraAccountData instantiates a new AvalaraAccountData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvalaraAccountData(type_ string, attributes GETAvalaraAccounts200ResponseDataInnerAttributes) *AvalaraAccountData {
	this := AvalaraAccountData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewAvalaraAccountDataWithDefaults instantiates a new AvalaraAccountData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvalaraAccountDataWithDefaults() *AvalaraAccountData {
	this := AvalaraAccountData{}
	return &this
}

// GetType returns the Type field value
func (o *AvalaraAccountData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AvalaraAccountData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AvalaraAccountData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *AvalaraAccountData) GetAttributes() GETAvalaraAccounts200ResponseDataInnerAttributes {
	if o == nil {
		var ret GETAvalaraAccounts200ResponseDataInnerAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AvalaraAccountData) GetAttributesOk() (*GETAvalaraAccounts200ResponseDataInnerAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *AvalaraAccountData) SetAttributes(v GETAvalaraAccounts200ResponseDataInnerAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AvalaraAccountData) GetRelationships() AvalaraAccountDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret AvalaraAccountDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvalaraAccountData) GetRelationshipsOk() (*AvalaraAccountDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AvalaraAccountData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AvalaraAccountDataRelationships and assigns it to the Relationships field.
func (o *AvalaraAccountData) SetRelationships(v AvalaraAccountDataRelationships) {
	o.Relationships = &v
}

func (o AvalaraAccountData) MarshalJSON() ([]byte, error) {
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

type NullableAvalaraAccountData struct {
	value *AvalaraAccountData
	isSet bool
}

func (v NullableAvalaraAccountData) Get() *AvalaraAccountData {
	return v.value
}

func (v *NullableAvalaraAccountData) Set(val *AvalaraAccountData) {
	v.value = val
	v.isSet = true
}

func (v NullableAvalaraAccountData) IsSet() bool {
	return v.isSet
}

func (v *NullableAvalaraAccountData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvalaraAccountData(val *AvalaraAccountData) *NullableAvalaraAccountData {
	return &NullableAvalaraAccountData{value: val, isSet: true}
}

func (v NullableAvalaraAccountData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvalaraAccountData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
