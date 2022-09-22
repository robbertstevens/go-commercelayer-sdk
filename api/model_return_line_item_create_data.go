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

// ReturnLineItemCreateData struct for ReturnLineItemCreateData
type ReturnLineItemCreateData struct {
	// The resource's type
	Type          string                                       `json:"type"`
	Attributes    POSTReturnLineItems201ResponseDataAttributes `json:"attributes"`
	Relationships *ReturnLineItemCreateDataRelationships       `json:"relationships,omitempty"`
}

// NewReturnLineItemCreateData instantiates a new ReturnLineItemCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReturnLineItemCreateData(type_ string, attributes POSTReturnLineItems201ResponseDataAttributes) *ReturnLineItemCreateData {
	this := ReturnLineItemCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewReturnLineItemCreateDataWithDefaults instantiates a new ReturnLineItemCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReturnLineItemCreateDataWithDefaults() *ReturnLineItemCreateData {
	this := ReturnLineItemCreateData{}
	return &this
}

// GetType returns the Type field value
func (o *ReturnLineItemCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ReturnLineItemCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ReturnLineItemCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *ReturnLineItemCreateData) GetAttributes() POSTReturnLineItems201ResponseDataAttributes {
	if o == nil {
		var ret POSTReturnLineItems201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ReturnLineItemCreateData) GetAttributesOk() (*POSTReturnLineItems201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *ReturnLineItemCreateData) SetAttributes(v POSTReturnLineItems201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ReturnLineItemCreateData) GetRelationships() ReturnLineItemCreateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret ReturnLineItemCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnLineItemCreateData) GetRelationshipsOk() (*ReturnLineItemCreateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ReturnLineItemCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given ReturnLineItemCreateDataRelationships and assigns it to the Relationships field.
func (o *ReturnLineItemCreateData) SetRelationships(v ReturnLineItemCreateDataRelationships) {
	o.Relationships = &v
}

func (o ReturnLineItemCreateData) MarshalJSON() ([]byte, error) {
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

type NullableReturnLineItemCreateData struct {
	value *ReturnLineItemCreateData
	isSet bool
}

func (v NullableReturnLineItemCreateData) Get() *ReturnLineItemCreateData {
	return v.value
}

func (v *NullableReturnLineItemCreateData) Set(val *ReturnLineItemCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnLineItemCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnLineItemCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnLineItemCreateData(val *ReturnLineItemCreateData) *NullableReturnLineItemCreateData {
	return &NullableReturnLineItemCreateData{value: val, isSet: true}
}

func (v NullableReturnLineItemCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnLineItemCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
