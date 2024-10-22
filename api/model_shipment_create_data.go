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

// checks if the ShipmentCreateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipmentCreateData{}

// ShipmentCreateData struct for ShipmentCreateData
type ShipmentCreateData struct {
	// The resource's type
	Type          interface{}                                `json:"type"`
	Attributes    POSTAdyenPayments201ResponseDataAttributes `json:"attributes"`
	Relationships *ShipmentCreateDataRelationships           `json:"relationships,omitempty"`
}

// NewShipmentCreateData instantiates a new ShipmentCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipmentCreateData(type_ interface{}, attributes POSTAdyenPayments201ResponseDataAttributes) *ShipmentCreateData {
	this := ShipmentCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewShipmentCreateDataWithDefaults instantiates a new ShipmentCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipmentCreateDataWithDefaults() *ShipmentCreateData {
	this := ShipmentCreateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ShipmentCreateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ShipmentCreateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ShipmentCreateData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *ShipmentCreateData) GetAttributes() POSTAdyenPayments201ResponseDataAttributes {
	if o == nil {
		var ret POSTAdyenPayments201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ShipmentCreateData) GetAttributesOk() (*POSTAdyenPayments201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *ShipmentCreateData) SetAttributes(v POSTAdyenPayments201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ShipmentCreateData) GetRelationships() ShipmentCreateDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret ShipmentCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentCreateData) GetRelationshipsOk() (*ShipmentCreateDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ShipmentCreateData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given ShipmentCreateDataRelationships and assigns it to the Relationships field.
func (o *ShipmentCreateData) SetRelationships(v ShipmentCreateDataRelationships) {
	o.Relationships = &v
}

func (o ShipmentCreateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShipmentCreateData) ToMap() (map[string]interface{}, error) {
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

type NullableShipmentCreateData struct {
	value *ShipmentCreateData
	isSet bool
}

func (v NullableShipmentCreateData) Get() *ShipmentCreateData {
	return v.value
}

func (v *NullableShipmentCreateData) Set(val *ShipmentCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentCreateData(val *ShipmentCreateData) *NullableShipmentCreateData {
	return &NullableShipmentCreateData{value: val, isSet: true}
}

func (v NullableShipmentCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipmentCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
