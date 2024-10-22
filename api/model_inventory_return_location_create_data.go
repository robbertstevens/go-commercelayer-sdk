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

// checks if the InventoryReturnLocationCreateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InventoryReturnLocationCreateData{}

// InventoryReturnLocationCreateData struct for InventoryReturnLocationCreateData
type InventoryReturnLocationCreateData struct {
	// The resource's type
	Type          interface{}                                           `json:"type"`
	Attributes    POSTInventoryReturnLocations201ResponseDataAttributes `json:"attributes"`
	Relationships *InventoryReturnLocationCreateDataRelationships       `json:"relationships,omitempty"`
}

// NewInventoryReturnLocationCreateData instantiates a new InventoryReturnLocationCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryReturnLocationCreateData(type_ interface{}, attributes POSTInventoryReturnLocations201ResponseDataAttributes) *InventoryReturnLocationCreateData {
	this := InventoryReturnLocationCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewInventoryReturnLocationCreateDataWithDefaults instantiates a new InventoryReturnLocationCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryReturnLocationCreateDataWithDefaults() *InventoryReturnLocationCreateData {
	this := InventoryReturnLocationCreateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *InventoryReturnLocationCreateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryReturnLocationCreateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InventoryReturnLocationCreateData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *InventoryReturnLocationCreateData) GetAttributes() POSTInventoryReturnLocations201ResponseDataAttributes {
	if o == nil {
		var ret POSTInventoryReturnLocations201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *InventoryReturnLocationCreateData) GetAttributesOk() (*POSTInventoryReturnLocations201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *InventoryReturnLocationCreateData) SetAttributes(v POSTInventoryReturnLocations201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *InventoryReturnLocationCreateData) GetRelationships() InventoryReturnLocationCreateDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret InventoryReturnLocationCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryReturnLocationCreateData) GetRelationshipsOk() (*InventoryReturnLocationCreateDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *InventoryReturnLocationCreateData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given InventoryReturnLocationCreateDataRelationships and assigns it to the Relationships field.
func (o *InventoryReturnLocationCreateData) SetRelationships(v InventoryReturnLocationCreateDataRelationships) {
	o.Relationships = &v
}

func (o InventoryReturnLocationCreateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InventoryReturnLocationCreateData) ToMap() (map[string]interface{}, error) {
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

type NullableInventoryReturnLocationCreateData struct {
	value *InventoryReturnLocationCreateData
	isSet bool
}

func (v NullableInventoryReturnLocationCreateData) Get() *InventoryReturnLocationCreateData {
	return v.value
}

func (v *NullableInventoryReturnLocationCreateData) Set(val *InventoryReturnLocationCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryReturnLocationCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryReturnLocationCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryReturnLocationCreateData(val *InventoryReturnLocationCreateData) *NullableInventoryReturnLocationCreateData {
	return &NullableInventoryReturnLocationCreateData{value: val, isSet: true}
}

func (v NullableInventoryReturnLocationCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryReturnLocationCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
