/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.4.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// SkuListPromotionRuleUpdateData struct for SkuListPromotionRuleUpdateData
type SkuListPromotionRuleUpdateData struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id            string                                             `json:"id"`
	Attributes    POSTSkuListPromotionRules201ResponseDataAttributes `json:"attributes"`
	Relationships *SkuListPromotionRuleUpdateDataRelationships       `json:"relationships,omitempty"`
}

// NewSkuListPromotionRuleUpdateData instantiates a new SkuListPromotionRuleUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuListPromotionRuleUpdateData(type_ string, id string, attributes POSTSkuListPromotionRules201ResponseDataAttributes) *SkuListPromotionRuleUpdateData {
	this := SkuListPromotionRuleUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewSkuListPromotionRuleUpdateDataWithDefaults instantiates a new SkuListPromotionRuleUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuListPromotionRuleUpdateDataWithDefaults() *SkuListPromotionRuleUpdateData {
	this := SkuListPromotionRuleUpdateData{}
	return &this
}

// GetType returns the Type field value
func (o *SkuListPromotionRuleUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SkuListPromotionRuleUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SkuListPromotionRuleUpdateData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *SkuListPromotionRuleUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SkuListPromotionRuleUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SkuListPromotionRuleUpdateData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *SkuListPromotionRuleUpdateData) GetAttributes() POSTSkuListPromotionRules201ResponseDataAttributes {
	if o == nil {
		var ret POSTSkuListPromotionRules201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SkuListPromotionRuleUpdateData) GetAttributesOk() (*POSTSkuListPromotionRules201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *SkuListPromotionRuleUpdateData) SetAttributes(v POSTSkuListPromotionRules201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *SkuListPromotionRuleUpdateData) GetRelationships() SkuListPromotionRuleUpdateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret SkuListPromotionRuleUpdateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuListPromotionRuleUpdateData) GetRelationshipsOk() (*SkuListPromotionRuleUpdateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *SkuListPromotionRuleUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given SkuListPromotionRuleUpdateDataRelationships and assigns it to the Relationships field.
func (o *SkuListPromotionRuleUpdateData) SetRelationships(v SkuListPromotionRuleUpdateDataRelationships) {
	o.Relationships = &v
}

func (o SkuListPromotionRuleUpdateData) MarshalJSON() ([]byte, error) {
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

type NullableSkuListPromotionRuleUpdateData struct {
	value *SkuListPromotionRuleUpdateData
	isSet bool
}

func (v NullableSkuListPromotionRuleUpdateData) Get() *SkuListPromotionRuleUpdateData {
	return v.value
}

func (v *NullableSkuListPromotionRuleUpdateData) Set(val *SkuListPromotionRuleUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuListPromotionRuleUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuListPromotionRuleUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuListPromotionRuleUpdateData(val *SkuListPromotionRuleUpdateData) *NullableSkuListPromotionRuleUpdateData {
	return &NullableSkuListPromotionRuleUpdateData{value: val, isSet: true}
}

func (v NullableSkuListPromotionRuleUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkuListPromotionRuleUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
