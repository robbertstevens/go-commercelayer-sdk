/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// InventoryStockLocationCreateData struct for InventoryStockLocationCreateData
type InventoryStockLocationCreateData struct {
	// The resource's type
	Type          string                                               `json:"type"`
	Attributes    POSTInventoryStockLocations201ResponseDataAttributes `json:"attributes"`
	Relationships *InventoryReturnLocationCreateDataRelationships      `json:"relationships,omitempty"`
}

// NewInventoryStockLocationCreateData instantiates a new InventoryStockLocationCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryStockLocationCreateData(type_ string, attributes POSTInventoryStockLocations201ResponseDataAttributes) *InventoryStockLocationCreateData {
	this := InventoryStockLocationCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewInventoryStockLocationCreateDataWithDefaults instantiates a new InventoryStockLocationCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryStockLocationCreateDataWithDefaults() *InventoryStockLocationCreateData {
	this := InventoryStockLocationCreateData{}
	return &this
}

// GetType returns the Type field value
func (o *InventoryStockLocationCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InventoryStockLocationCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InventoryStockLocationCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *InventoryStockLocationCreateData) GetAttributes() POSTInventoryStockLocations201ResponseDataAttributes {
	if o == nil {
		var ret POSTInventoryStockLocations201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *InventoryStockLocationCreateData) GetAttributesOk() (*POSTInventoryStockLocations201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *InventoryStockLocationCreateData) SetAttributes(v POSTInventoryStockLocations201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *InventoryStockLocationCreateData) GetRelationships() InventoryReturnLocationCreateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret InventoryReturnLocationCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryStockLocationCreateData) GetRelationshipsOk() (*InventoryReturnLocationCreateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *InventoryStockLocationCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given InventoryReturnLocationCreateDataRelationships and assigns it to the Relationships field.
func (o *InventoryStockLocationCreateData) SetRelationships(v InventoryReturnLocationCreateDataRelationships) {
	o.Relationships = &v
}

func (o InventoryStockLocationCreateData) MarshalJSON() ([]byte, error) {
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

type NullableInventoryStockLocationCreateData struct {
	value *InventoryStockLocationCreateData
	isSet bool
}

func (v NullableInventoryStockLocationCreateData) Get() *InventoryStockLocationCreateData {
	return v.value
}

func (v *NullableInventoryStockLocationCreateData) Set(val *InventoryStockLocationCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryStockLocationCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryStockLocationCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryStockLocationCreateData(val *InventoryStockLocationCreateData) *NullableInventoryStockLocationCreateData {
	return &NullableInventoryStockLocationCreateData{value: val, isSet: true}
}

func (v NullableInventoryStockLocationCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryStockLocationCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
