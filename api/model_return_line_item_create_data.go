/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.3.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ReturnLineItemCreateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReturnLineItemCreateData{}

// ReturnLineItemCreateData struct for ReturnLineItemCreateData
type ReturnLineItemCreateData struct {
	// The resource's type
	Type          interface{}                                  `json:"type"`
	Attributes    POSTReturnLineItems201ResponseDataAttributes `json:"attributes"`
	Relationships *ReturnLineItemCreateDataRelationships       `json:"relationships,omitempty"`
}

// NewReturnLineItemCreateData instantiates a new ReturnLineItemCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReturnLineItemCreateData(type_ interface{}, attributes POSTReturnLineItems201ResponseDataAttributes) *ReturnLineItemCreateData {
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
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ReturnLineItemCreateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReturnLineItemCreateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ReturnLineItemCreateData) SetType(v interface{}) {
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
	if o == nil || IsNil(o.Relationships) {
		var ret ReturnLineItemCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnLineItemCreateData) GetRelationshipsOk() (*ReturnLineItemCreateDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ReturnLineItemCreateData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given ReturnLineItemCreateDataRelationships and assigns it to the Relationships field.
func (o *ReturnLineItemCreateData) SetRelationships(v ReturnLineItemCreateDataRelationships) {
	o.Relationships = &v
}

func (o ReturnLineItemCreateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReturnLineItemCreateData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	toSerialize["attributes"] = o.Attributes
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
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
