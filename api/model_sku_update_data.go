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

// SkuUpdateData struct for SkuUpdateData
type SkuUpdateData struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id            string                                  `json:"id"`
	Attributes    PATCHSkusSkuId200ResponseDataAttributes `json:"attributes"`
	Relationships *SkuUpdateDataRelationships             `json:"relationships,omitempty"`
}

// NewSkuUpdateData instantiates a new SkuUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuUpdateData(type_ string, id string, attributes PATCHSkusSkuId200ResponseDataAttributes) *SkuUpdateData {
	this := SkuUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewSkuUpdateDataWithDefaults instantiates a new SkuUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuUpdateDataWithDefaults() *SkuUpdateData {
	this := SkuUpdateData{}
	return &this
}

// GetType returns the Type field value
func (o *SkuUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SkuUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SkuUpdateData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *SkuUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SkuUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SkuUpdateData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *SkuUpdateData) GetAttributes() PATCHSkusSkuId200ResponseDataAttributes {
	if o == nil {
		var ret PATCHSkusSkuId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SkuUpdateData) GetAttributesOk() (*PATCHSkusSkuId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *SkuUpdateData) SetAttributes(v PATCHSkusSkuId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *SkuUpdateData) GetRelationships() SkuUpdateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret SkuUpdateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuUpdateData) GetRelationshipsOk() (*SkuUpdateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *SkuUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given SkuUpdateDataRelationships and assigns it to the Relationships field.
func (o *SkuUpdateData) SetRelationships(v SkuUpdateDataRelationships) {
	o.Relationships = &v
}

func (o SkuUpdateData) MarshalJSON() ([]byte, error) {
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

type NullableSkuUpdateData struct {
	value *SkuUpdateData
	isSet bool
}

func (v NullableSkuUpdateData) Get() *SkuUpdateData {
	return v.value
}

func (v *NullableSkuUpdateData) Set(val *SkuUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuUpdateData(val *SkuUpdateData) *NullableSkuUpdateData {
	return &NullableSkuUpdateData{value: val, isSet: true}
}

func (v NullableSkuUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkuUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
