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

// LineItemUpdateData struct for LineItemUpdateData
type LineItemUpdateData struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id            string                                            `json:"id"`
	Attributes    PATCHLineItemsLineItemId200ResponseDataAttributes `json:"attributes"`
	Relationships map[string]interface{}                            `json:"relationships,omitempty"`
}

// NewLineItemUpdateData instantiates a new LineItemUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLineItemUpdateData(type_ string, id string, attributes PATCHLineItemsLineItemId200ResponseDataAttributes) *LineItemUpdateData {
	this := LineItemUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewLineItemUpdateDataWithDefaults instantiates a new LineItemUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLineItemUpdateDataWithDefaults() *LineItemUpdateData {
	this := LineItemUpdateData{}
	return &this
}

// GetType returns the Type field value
func (o *LineItemUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LineItemUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LineItemUpdateData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *LineItemUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LineItemUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LineItemUpdateData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *LineItemUpdateData) GetAttributes() PATCHLineItemsLineItemId200ResponseDataAttributes {
	if o == nil {
		var ret PATCHLineItemsLineItemId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *LineItemUpdateData) GetAttributesOk() (*PATCHLineItemsLineItemId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *LineItemUpdateData) SetAttributes(v PATCHLineItemsLineItemId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *LineItemUpdateData) GetRelationships() map[string]interface{} {
	if o == nil || o.Relationships == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemUpdateData) GetRelationshipsOk() (map[string]interface{}, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *LineItemUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]interface{} and assigns it to the Relationships field.
func (o *LineItemUpdateData) SetRelationships(v map[string]interface{}) {
	o.Relationships = v
}

func (o LineItemUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableLineItemUpdateData struct {
	value *LineItemUpdateData
	isSet bool
}

func (v NullableLineItemUpdateData) Get() *LineItemUpdateData {
	return v.value
}

func (v *NullableLineItemUpdateData) Set(val *LineItemUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableLineItemUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableLineItemUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLineItemUpdateData(val *LineItemUpdateData) *NullableLineItemUpdateData {
	return &NullableLineItemUpdateData{value: val, isSet: true}
}

func (v NullableLineItemUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLineItemUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
