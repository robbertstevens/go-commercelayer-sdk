/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.3.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the MerchantData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantData{}

// MerchantData struct for MerchantData
type MerchantData struct {
	// The resource's type
	Type          interface{}                                     `json:"type"`
	Attributes    GETMerchantsMerchantId200ResponseDataAttributes `json:"attributes"`
	Relationships *MerchantDataRelationships                      `json:"relationships,omitempty"`
}

// NewMerchantData instantiates a new MerchantData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantData(type_ interface{}, attributes GETMerchantsMerchantId200ResponseDataAttributes) *MerchantData {
	this := MerchantData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewMerchantDataWithDefaults instantiates a new MerchantData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantDataWithDefaults() *MerchantData {
	this := MerchantData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *MerchantData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MerchantData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *MerchantData) GetAttributes() GETMerchantsMerchantId200ResponseDataAttributes {
	if o == nil {
		var ret GETMerchantsMerchantId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *MerchantData) GetAttributesOk() (*GETMerchantsMerchantId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *MerchantData) SetAttributes(v GETMerchantsMerchantId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *MerchantData) GetRelationships() MerchantDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret MerchantDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantData) GetRelationshipsOk() (*MerchantDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *MerchantData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given MerchantDataRelationships and assigns it to the Relationships field.
func (o *MerchantData) SetRelationships(v MerchantDataRelationships) {
	o.Relationships = &v
}

func (o MerchantData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	toSerialize["attributes"] = o.Attributes
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullableMerchantData struct {
	value *MerchantData
	isSet bool
}

func (v NullableMerchantData) Get() *MerchantData {
	return v.value
}

func (v *NullableMerchantData) Set(val *MerchantData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantData(val *MerchantData) *NullableMerchantData {
	return &NullableMerchantData{value: val, isSet: true}
}

func (v NullableMerchantData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
