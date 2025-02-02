/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.3.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the AdjustmentUpdateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdjustmentUpdateData{}

// AdjustmentUpdateData struct for AdjustmentUpdateData
type AdjustmentUpdateData struct {
	// The resource's type
	Type interface{} `json:"type"`
	// The resource's id
	Id            interface{}                                           `json:"id"`
	Attributes    PATCHAdjustmentsAdjustmentId200ResponseDataAttributes `json:"attributes"`
	Relationships interface{}                                           `json:"relationships,omitempty"`
}

// NewAdjustmentUpdateData instantiates a new AdjustmentUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdjustmentUpdateData(type_ interface{}, id interface{}, attributes PATCHAdjustmentsAdjustmentId200ResponseDataAttributes) *AdjustmentUpdateData {
	this := AdjustmentUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewAdjustmentUpdateDataWithDefaults instantiates a new AdjustmentUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdjustmentUpdateDataWithDefaults() *AdjustmentUpdateData {
	this := AdjustmentUpdateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *AdjustmentUpdateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdjustmentUpdateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AdjustmentUpdateData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *AdjustmentUpdateData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdjustmentUpdateData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AdjustmentUpdateData) SetId(v interface{}) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *AdjustmentUpdateData) GetAttributes() PATCHAdjustmentsAdjustmentId200ResponseDataAttributes {
	if o == nil {
		var ret PATCHAdjustmentsAdjustmentId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AdjustmentUpdateData) GetAttributesOk() (*PATCHAdjustmentsAdjustmentId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *AdjustmentUpdateData) SetAttributes(v PATCHAdjustmentsAdjustmentId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdjustmentUpdateData) GetRelationships() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdjustmentUpdateData) GetRelationshipsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return &o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AdjustmentUpdateData) HasRelationships() bool {
	if o != nil && IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given interface{} and assigns it to the Relationships field.
func (o *AdjustmentUpdateData) SetRelationships(v interface{}) {
	o.Relationships = v
}

func (o AdjustmentUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdjustmentUpdateData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["attributes"] = o.Attributes
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullableAdjustmentUpdateData struct {
	value *AdjustmentUpdateData
	isSet bool
}

func (v NullableAdjustmentUpdateData) Get() *AdjustmentUpdateData {
	return v.value
}

func (v *NullableAdjustmentUpdateData) Set(val *AdjustmentUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableAdjustmentUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableAdjustmentUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdjustmentUpdateData(val *AdjustmentUpdateData) *NullableAdjustmentUpdateData {
	return &NullableAdjustmentUpdateData{value: val, isSet: true}
}

func (v NullableAdjustmentUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdjustmentUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
