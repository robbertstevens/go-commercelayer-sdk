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

// SkuOptionData struct for SkuOptionData
type SkuOptionData struct {
	// The resource's type
	Type          string                                      `json:"type"`
	Attributes    GETSkuOptions200ResponseDataInnerAttributes `json:"attributes"`
	Relationships *CarrierAccountDataRelationships            `json:"relationships,omitempty"`
}

// NewSkuOptionData instantiates a new SkuOptionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuOptionData(type_ string, attributes GETSkuOptions200ResponseDataInnerAttributes) *SkuOptionData {
	this := SkuOptionData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewSkuOptionDataWithDefaults instantiates a new SkuOptionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuOptionDataWithDefaults() *SkuOptionData {
	this := SkuOptionData{}
	return &this
}

// GetType returns the Type field value
func (o *SkuOptionData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SkuOptionData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SkuOptionData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *SkuOptionData) GetAttributes() GETSkuOptions200ResponseDataInnerAttributes {
	if o == nil {
		var ret GETSkuOptions200ResponseDataInnerAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SkuOptionData) GetAttributesOk() (*GETSkuOptions200ResponseDataInnerAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *SkuOptionData) SetAttributes(v GETSkuOptions200ResponseDataInnerAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *SkuOptionData) GetRelationships() CarrierAccountDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret CarrierAccountDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuOptionData) GetRelationshipsOk() (*CarrierAccountDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *SkuOptionData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given CarrierAccountDataRelationships and assigns it to the Relationships field.
func (o *SkuOptionData) SetRelationships(v CarrierAccountDataRelationships) {
	o.Relationships = &v
}

func (o SkuOptionData) MarshalJSON() ([]byte, error) {
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

type NullableSkuOptionData struct {
	value *SkuOptionData
	isSet bool
}

func (v NullableSkuOptionData) Get() *SkuOptionData {
	return v.value
}

func (v *NullableSkuOptionData) Set(val *SkuOptionData) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuOptionData) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuOptionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuOptionData(val *SkuOptionData) *NullableSkuOptionData {
	return &NullableSkuOptionData{value: val, isSet: true}
}

func (v NullableSkuOptionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkuOptionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
