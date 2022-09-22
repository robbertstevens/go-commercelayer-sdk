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

// MerchantCreateData struct for MerchantCreateData
type MerchantCreateData struct {
	// The resource's type
	Type          string                                 `json:"type"`
	Attributes    POSTMerchants201ResponseDataAttributes `json:"attributes"`
	Relationships *MerchantCreateDataRelationships       `json:"relationships,omitempty"`
}

// NewMerchantCreateData instantiates a new MerchantCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantCreateData(type_ string, attributes POSTMerchants201ResponseDataAttributes) *MerchantCreateData {
	this := MerchantCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewMerchantCreateDataWithDefaults instantiates a new MerchantCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantCreateDataWithDefaults() *MerchantCreateData {
	this := MerchantCreateData{}
	return &this
}

// GetType returns the Type field value
func (o *MerchantCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MerchantCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MerchantCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *MerchantCreateData) GetAttributes() POSTMerchants201ResponseDataAttributes {
	if o == nil {
		var ret POSTMerchants201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *MerchantCreateData) GetAttributesOk() (*POSTMerchants201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *MerchantCreateData) SetAttributes(v POSTMerchants201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *MerchantCreateData) GetRelationships() MerchantCreateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret MerchantCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantCreateData) GetRelationshipsOk() (*MerchantCreateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *MerchantCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given MerchantCreateDataRelationships and assigns it to the Relationships field.
func (o *MerchantCreateData) SetRelationships(v MerchantCreateDataRelationships) {
	o.Relationships = &v
}

func (o MerchantCreateData) MarshalJSON() ([]byte, error) {
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

type NullableMerchantCreateData struct {
	value *MerchantCreateData
	isSet bool
}

func (v NullableMerchantCreateData) Get() *MerchantCreateData {
	return v.value
}

func (v *NullableMerchantCreateData) Set(val *MerchantCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantCreateData(val *MerchantCreateData) *NullableMerchantCreateData {
	return &NullableMerchantCreateData{value: val, isSet: true}
}

func (v NullableMerchantCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
