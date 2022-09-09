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

// OrderCopyCreateData struct for OrderCopyCreateData
type OrderCopyCreateData struct {
	// The resource's type
	Type string `json:"type"`
	Attributes POSTOrderCopies201ResponseDataAttributes `json:"attributes"`
	Relationships *OrderCopyCreateDataRelationships `json:"relationships,omitempty"`
}

// NewOrderCopyCreateData instantiates a new OrderCopyCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderCopyCreateData(type_ string, attributes POSTOrderCopies201ResponseDataAttributes) *OrderCopyCreateData {
	this := OrderCopyCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewOrderCopyCreateDataWithDefaults instantiates a new OrderCopyCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderCopyCreateDataWithDefaults() *OrderCopyCreateData {
	this := OrderCopyCreateData{}
	var type_ string = "order_copies"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *OrderCopyCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OrderCopyCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OrderCopyCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *OrderCopyCreateData) GetAttributes() POSTOrderCopies201ResponseDataAttributes {
	if o == nil {
		var ret POSTOrderCopies201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *OrderCopyCreateData) GetAttributesOk() (*POSTOrderCopies201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *OrderCopyCreateData) SetAttributes(v POSTOrderCopies201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *OrderCopyCreateData) GetRelationships() OrderCopyCreateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret OrderCopyCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCopyCreateData) GetRelationshipsOk() (*OrderCopyCreateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *OrderCopyCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given OrderCopyCreateDataRelationships and assigns it to the Relationships field.
func (o *OrderCopyCreateData) SetRelationships(v OrderCopyCreateDataRelationships) {
	o.Relationships = &v
}

func (o OrderCopyCreateData) MarshalJSON() ([]byte, error) {
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

type NullableOrderCopyCreateData struct {
	value *OrderCopyCreateData
	isSet bool
}

func (v NullableOrderCopyCreateData) Get() *OrderCopyCreateData {
	return v.value
}

func (v *NullableOrderCopyCreateData) Set(val *OrderCopyCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCopyCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCopyCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCopyCreateData(val *OrderCopyCreateData) *NullableOrderCopyCreateData {
	return &NullableOrderCopyCreateData{value: val, isSet: true}
}

func (v NullableOrderCopyCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCopyCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


