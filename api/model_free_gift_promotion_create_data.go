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

// FreeGiftPromotionCreateData struct for FreeGiftPromotionCreateData
type FreeGiftPromotionCreateData struct {
	// The resource's type
	Type          string                                          `json:"type"`
	Attributes    POSTFreeGiftPromotions201ResponseDataAttributes `json:"attributes"`
	Relationships *FixedPricePromotionCreateDataRelationships     `json:"relationships,omitempty"`
}

// NewFreeGiftPromotionCreateData instantiates a new FreeGiftPromotionCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFreeGiftPromotionCreateData(type_ string, attributes POSTFreeGiftPromotions201ResponseDataAttributes) *FreeGiftPromotionCreateData {
	this := FreeGiftPromotionCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewFreeGiftPromotionCreateDataWithDefaults instantiates a new FreeGiftPromotionCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFreeGiftPromotionCreateDataWithDefaults() *FreeGiftPromotionCreateData {
	this := FreeGiftPromotionCreateData{}
	return &this
}

// GetType returns the Type field value
func (o *FreeGiftPromotionCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FreeGiftPromotionCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FreeGiftPromotionCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *FreeGiftPromotionCreateData) GetAttributes() POSTFreeGiftPromotions201ResponseDataAttributes {
	if o == nil {
		var ret POSTFreeGiftPromotions201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *FreeGiftPromotionCreateData) GetAttributesOk() (*POSTFreeGiftPromotions201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *FreeGiftPromotionCreateData) SetAttributes(v POSTFreeGiftPromotions201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *FreeGiftPromotionCreateData) GetRelationships() FixedPricePromotionCreateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret FixedPricePromotionCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreeGiftPromotionCreateData) GetRelationshipsOk() (*FixedPricePromotionCreateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *FreeGiftPromotionCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given FixedPricePromotionCreateDataRelationships and assigns it to the Relationships field.
func (o *FreeGiftPromotionCreateData) SetRelationships(v FixedPricePromotionCreateDataRelationships) {
	o.Relationships = &v
}

func (o FreeGiftPromotionCreateData) MarshalJSON() ([]byte, error) {
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

type NullableFreeGiftPromotionCreateData struct {
	value *FreeGiftPromotionCreateData
	isSet bool
}

func (v NullableFreeGiftPromotionCreateData) Get() *FreeGiftPromotionCreateData {
	return v.value
}

func (v *NullableFreeGiftPromotionCreateData) Set(val *FreeGiftPromotionCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableFreeGiftPromotionCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableFreeGiftPromotionCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFreeGiftPromotionCreateData(val *FreeGiftPromotionCreateData) *NullableFreeGiftPromotionCreateData {
	return &NullableFreeGiftPromotionCreateData{value: val, isSet: true}
}

func (v NullableFreeGiftPromotionCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFreeGiftPromotionCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
