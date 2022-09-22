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

// ShippingWeightTierData struct for ShippingWeightTierData
type ShippingWeightTierData struct {
	// The resource's type
	Type          string                                               `json:"type"`
	Attributes    GETShippingMethodTiers200ResponseDataInnerAttributes `json:"attributes"`
	Relationships *ShippingMethodTierDataRelationships                 `json:"relationships,omitempty"`
}

// NewShippingWeightTierData instantiates a new ShippingWeightTierData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingWeightTierData(type_ string, attributes GETShippingMethodTiers200ResponseDataInnerAttributes) *ShippingWeightTierData {
	this := ShippingWeightTierData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewShippingWeightTierDataWithDefaults instantiates a new ShippingWeightTierData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingWeightTierDataWithDefaults() *ShippingWeightTierData {
	this := ShippingWeightTierData{}
	return &this
}

// GetType returns the Type field value
func (o *ShippingWeightTierData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ShippingWeightTierData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ShippingWeightTierData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *ShippingWeightTierData) GetAttributes() GETShippingMethodTiers200ResponseDataInnerAttributes {
	if o == nil {
		var ret GETShippingMethodTiers200ResponseDataInnerAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ShippingWeightTierData) GetAttributesOk() (*GETShippingMethodTiers200ResponseDataInnerAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *ShippingWeightTierData) SetAttributes(v GETShippingMethodTiers200ResponseDataInnerAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ShippingWeightTierData) GetRelationships() ShippingMethodTierDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret ShippingMethodTierDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingWeightTierData) GetRelationshipsOk() (*ShippingMethodTierDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ShippingWeightTierData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given ShippingMethodTierDataRelationships and assigns it to the Relationships field.
func (o *ShippingWeightTierData) SetRelationships(v ShippingMethodTierDataRelationships) {
	o.Relationships = &v
}

func (o ShippingWeightTierData) MarshalJSON() ([]byte, error) {
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

type NullableShippingWeightTierData struct {
	value *ShippingWeightTierData
	isSet bool
}

func (v NullableShippingWeightTierData) Get() *ShippingWeightTierData {
	return v.value
}

func (v *NullableShippingWeightTierData) Set(val *ShippingWeightTierData) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingWeightTierData) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingWeightTierData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingWeightTierData(val *ShippingWeightTierData) *NullableShippingWeightTierData {
	return &NullableShippingWeightTierData{value: val, isSet: true}
}

func (v NullableShippingWeightTierData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingWeightTierData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
