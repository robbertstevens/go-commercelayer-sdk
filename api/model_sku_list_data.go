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

// SkuListData struct for SkuListData
type SkuListData struct {
	// The resource's type
	Type          string                                    `json:"type"`
	Attributes    GETSkuLists200ResponseDataInnerAttributes `json:"attributes"`
	Relationships *SkuListDataRelationships                 `json:"relationships,omitempty"`
}

// NewSkuListData instantiates a new SkuListData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuListData(type_ string, attributes GETSkuLists200ResponseDataInnerAttributes) *SkuListData {
	this := SkuListData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewSkuListDataWithDefaults instantiates a new SkuListData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuListDataWithDefaults() *SkuListData {
	this := SkuListData{}
	return &this
}

// GetType returns the Type field value
func (o *SkuListData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SkuListData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SkuListData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *SkuListData) GetAttributes() GETSkuLists200ResponseDataInnerAttributes {
	if o == nil {
		var ret GETSkuLists200ResponseDataInnerAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SkuListData) GetAttributesOk() (*GETSkuLists200ResponseDataInnerAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *SkuListData) SetAttributes(v GETSkuLists200ResponseDataInnerAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *SkuListData) GetRelationships() SkuListDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret SkuListDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuListData) GetRelationshipsOk() (*SkuListDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *SkuListData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given SkuListDataRelationships and assigns it to the Relationships field.
func (o *SkuListData) SetRelationships(v SkuListDataRelationships) {
	o.Relationships = &v
}

func (o SkuListData) MarshalJSON() ([]byte, error) {
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

type NullableSkuListData struct {
	value *SkuListData
	isSet bool
}

func (v NullableSkuListData) Get() *SkuListData {
	return v.value
}

func (v *NullableSkuListData) Set(val *SkuListData) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuListData) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuListData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuListData(val *SkuListData) *NullableSkuListData {
	return &NullableSkuListData{value: val, isSet: true}
}

func (v NullableSkuListData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkuListData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
