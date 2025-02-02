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

// checks if the ShippingMethodCreateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShippingMethodCreateData{}

// ShippingMethodCreateData struct for ShippingMethodCreateData
type ShippingMethodCreateData struct {
	// The resource's type
	Type          interface{}                                  `json:"type"`
	Attributes    POSTShippingMethods201ResponseDataAttributes `json:"attributes"`
	Relationships *ShippingMethodCreateDataRelationships       `json:"relationships,omitempty"`
}

// NewShippingMethodCreateData instantiates a new ShippingMethodCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingMethodCreateData(type_ interface{}, attributes POSTShippingMethods201ResponseDataAttributes) *ShippingMethodCreateData {
	this := ShippingMethodCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewShippingMethodCreateDataWithDefaults instantiates a new ShippingMethodCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingMethodCreateDataWithDefaults() *ShippingMethodCreateData {
	this := ShippingMethodCreateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ShippingMethodCreateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ShippingMethodCreateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ShippingMethodCreateData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *ShippingMethodCreateData) GetAttributes() POSTShippingMethods201ResponseDataAttributes {
	if o == nil {
		var ret POSTShippingMethods201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ShippingMethodCreateData) GetAttributesOk() (*POSTShippingMethods201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *ShippingMethodCreateData) SetAttributes(v POSTShippingMethods201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ShippingMethodCreateData) GetRelationships() ShippingMethodCreateDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret ShippingMethodCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodCreateData) GetRelationshipsOk() (*ShippingMethodCreateDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ShippingMethodCreateData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given ShippingMethodCreateDataRelationships and assigns it to the Relationships field.
func (o *ShippingMethodCreateData) SetRelationships(v ShippingMethodCreateDataRelationships) {
	o.Relationships = &v
}

func (o ShippingMethodCreateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShippingMethodCreateData) ToMap() (map[string]interface{}, error) {
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

type NullableShippingMethodCreateData struct {
	value *ShippingMethodCreateData
	isSet bool
}

func (v NullableShippingMethodCreateData) Get() *ShippingMethodCreateData {
	return v.value
}

func (v *NullableShippingMethodCreateData) Set(val *ShippingMethodCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingMethodCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingMethodCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingMethodCreateData(val *ShippingMethodCreateData) *NullableShippingMethodCreateData {
	return &NullableShippingMethodCreateData{value: val, isSet: true}
}

func (v NullableShippingMethodCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingMethodCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
