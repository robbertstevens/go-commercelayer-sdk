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

// InventoryModelCreateData struct for InventoryModelCreateData
type InventoryModelCreateData struct {
	// The resource's type
	Type string `json:"type"`
	Attributes POSTInventoryModels201ResponseDataAttributes `json:"attributes"`
	Relationships map[string]interface{} `json:"relationships,omitempty"`
}

// NewInventoryModelCreateData instantiates a new InventoryModelCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryModelCreateData(type_ string, attributes POSTInventoryModels201ResponseDataAttributes) *InventoryModelCreateData {
	this := InventoryModelCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewInventoryModelCreateDataWithDefaults instantiates a new InventoryModelCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryModelCreateDataWithDefaults() *InventoryModelCreateData {
	this := InventoryModelCreateData{}
	var type_ string = "inventory_models"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *InventoryModelCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InventoryModelCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InventoryModelCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *InventoryModelCreateData) GetAttributes() POSTInventoryModels201ResponseDataAttributes {
	if o == nil {
		var ret POSTInventoryModels201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *InventoryModelCreateData) GetAttributesOk() (*POSTInventoryModels201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *InventoryModelCreateData) SetAttributes(v POSTInventoryModels201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *InventoryModelCreateData) GetRelationships() map[string]interface{} {
	if o == nil || o.Relationships == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryModelCreateData) GetRelationshipsOk() (map[string]interface{}, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *InventoryModelCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]interface{} and assigns it to the Relationships field.
func (o *InventoryModelCreateData) SetRelationships(v map[string]interface{}) {
	o.Relationships = v
}

func (o InventoryModelCreateData) MarshalJSON() ([]byte, error) {
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

type NullableInventoryModelCreateData struct {
	value *InventoryModelCreateData
	isSet bool
}

func (v NullableInventoryModelCreateData) Get() *InventoryModelCreateData {
	return v.value
}

func (v *NullableInventoryModelCreateData) Set(val *InventoryModelCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryModelCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryModelCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryModelCreateData(val *InventoryModelCreateData) *NullableInventoryModelCreateData {
	return &NullableInventoryModelCreateData{value: val, isSet: true}
}

func (v NullableInventoryModelCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryModelCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


