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

// checks if the AdjustmentCreateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdjustmentCreateData{}

// AdjustmentCreateData struct for AdjustmentCreateData
type AdjustmentCreateData struct {
	// The resource's type
	Type          interface{}                              `json:"type"`
	Attributes    POSTAdjustments201ResponseDataAttributes `json:"attributes"`
	Relationships interface{}                              `json:"relationships,omitempty"`
}

// NewAdjustmentCreateData instantiates a new AdjustmentCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdjustmentCreateData(type_ interface{}, attributes POSTAdjustments201ResponseDataAttributes) *AdjustmentCreateData {
	this := AdjustmentCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewAdjustmentCreateDataWithDefaults instantiates a new AdjustmentCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdjustmentCreateDataWithDefaults() *AdjustmentCreateData {
	this := AdjustmentCreateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *AdjustmentCreateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdjustmentCreateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AdjustmentCreateData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *AdjustmentCreateData) GetAttributes() POSTAdjustments201ResponseDataAttributes {
	if o == nil {
		var ret POSTAdjustments201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AdjustmentCreateData) GetAttributesOk() (*POSTAdjustments201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *AdjustmentCreateData) SetAttributes(v POSTAdjustments201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdjustmentCreateData) GetRelationships() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdjustmentCreateData) GetRelationshipsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return &o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AdjustmentCreateData) HasRelationships() bool {
	if o != nil && IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given interface{} and assigns it to the Relationships field.
func (o *AdjustmentCreateData) SetRelationships(v interface{}) {
	o.Relationships = v
}

func (o AdjustmentCreateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdjustmentCreateData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	toSerialize["attributes"] = o.Attributes
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullableAdjustmentCreateData struct {
	value *AdjustmentCreateData
	isSet bool
}

func (v NullableAdjustmentCreateData) Get() *AdjustmentCreateData {
	return v.value
}

func (v *NullableAdjustmentCreateData) Set(val *AdjustmentCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableAdjustmentCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableAdjustmentCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdjustmentCreateData(val *AdjustmentCreateData) *NullableAdjustmentCreateData {
	return &NullableAdjustmentCreateData{value: val, isSet: true}
}

func (v NullableAdjustmentCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdjustmentCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
