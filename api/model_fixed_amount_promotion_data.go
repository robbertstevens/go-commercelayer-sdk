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

// FixedAmountPromotionData struct for FixedAmountPromotionData
type FixedAmountPromotionData struct {
	// The resource's type
	Type          string                                                 `json:"type"`
	Attributes    GETFixedAmountPromotions200ResponseDataInnerAttributes `json:"attributes"`
	Relationships *ExternalPromotionDataRelationships                    `json:"relationships,omitempty"`
}

// NewFixedAmountPromotionData instantiates a new FixedAmountPromotionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFixedAmountPromotionData(type_ string, attributes GETFixedAmountPromotions200ResponseDataInnerAttributes) *FixedAmountPromotionData {
	this := FixedAmountPromotionData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewFixedAmountPromotionDataWithDefaults instantiates a new FixedAmountPromotionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFixedAmountPromotionDataWithDefaults() *FixedAmountPromotionData {
	this := FixedAmountPromotionData{}
	return &this
}

// GetType returns the Type field value
func (o *FixedAmountPromotionData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FixedAmountPromotionData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FixedAmountPromotionData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *FixedAmountPromotionData) GetAttributes() GETFixedAmountPromotions200ResponseDataInnerAttributes {
	if o == nil {
		var ret GETFixedAmountPromotions200ResponseDataInnerAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *FixedAmountPromotionData) GetAttributesOk() (*GETFixedAmountPromotions200ResponseDataInnerAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *FixedAmountPromotionData) SetAttributes(v GETFixedAmountPromotions200ResponseDataInnerAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *FixedAmountPromotionData) GetRelationships() ExternalPromotionDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret ExternalPromotionDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FixedAmountPromotionData) GetRelationshipsOk() (*ExternalPromotionDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *FixedAmountPromotionData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given ExternalPromotionDataRelationships and assigns it to the Relationships field.
func (o *FixedAmountPromotionData) SetRelationships(v ExternalPromotionDataRelationships) {
	o.Relationships = &v
}

func (o FixedAmountPromotionData) MarshalJSON() ([]byte, error) {
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

type NullableFixedAmountPromotionData struct {
	value *FixedAmountPromotionData
	isSet bool
}

func (v NullableFixedAmountPromotionData) Get() *FixedAmountPromotionData {
	return v.value
}

func (v *NullableFixedAmountPromotionData) Set(val *FixedAmountPromotionData) {
	v.value = val
	v.isSet = true
}

func (v NullableFixedAmountPromotionData) IsSet() bool {
	return v.isSet
}

func (v *NullableFixedAmountPromotionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFixedAmountPromotionData(val *FixedAmountPromotionData) *NullableFixedAmountPromotionData {
	return &NullableFixedAmountPromotionData{value: val, isSet: true}
}

func (v NullableFixedAmountPromotionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFixedAmountPromotionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
