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

// checks if the GiftCardUpdateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GiftCardUpdateData{}

// GiftCardUpdateData struct for GiftCardUpdateData
type GiftCardUpdateData struct {
	// The resource's type
	Type interface{} `json:"type"`
	// The resource's id
	Id            interface{}                                       `json:"id"`
	Attributes    PATCHGiftCardsGiftCardId200ResponseDataAttributes `json:"attributes"`
	Relationships *GiftCardCreateDataRelationships                  `json:"relationships,omitempty"`
}

// NewGiftCardUpdateData instantiates a new GiftCardUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGiftCardUpdateData(type_ interface{}, id interface{}, attributes PATCHGiftCardsGiftCardId200ResponseDataAttributes) *GiftCardUpdateData {
	this := GiftCardUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewGiftCardUpdateDataWithDefaults instantiates a new GiftCardUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGiftCardUpdateDataWithDefaults() *GiftCardUpdateData {
	this := GiftCardUpdateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *GiftCardUpdateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftCardUpdateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GiftCardUpdateData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *GiftCardUpdateData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftCardUpdateData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GiftCardUpdateData) SetId(v interface{}) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *GiftCardUpdateData) GetAttributes() PATCHGiftCardsGiftCardId200ResponseDataAttributes {
	if o == nil {
		var ret PATCHGiftCardsGiftCardId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *GiftCardUpdateData) GetAttributesOk() (*PATCHGiftCardsGiftCardId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *GiftCardUpdateData) SetAttributes(v PATCHGiftCardsGiftCardId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *GiftCardUpdateData) GetRelationships() GiftCardCreateDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret GiftCardCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardUpdateData) GetRelationshipsOk() (*GiftCardCreateDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *GiftCardUpdateData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given GiftCardCreateDataRelationships and assigns it to the Relationships field.
func (o *GiftCardUpdateData) SetRelationships(v GiftCardCreateDataRelationships) {
	o.Relationships = &v
}

func (o GiftCardUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GiftCardUpdateData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["attributes"] = o.Attributes
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullableGiftCardUpdateData struct {
	value *GiftCardUpdateData
	isSet bool
}

func (v NullableGiftCardUpdateData) Get() *GiftCardUpdateData {
	return v.value
}

func (v *NullableGiftCardUpdateData) Set(val *GiftCardUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableGiftCardUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableGiftCardUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGiftCardUpdateData(val *GiftCardUpdateData) *NullableGiftCardUpdateData {
	return &NullableGiftCardUpdateData{value: val, isSet: true}
}

func (v NullableGiftCardUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGiftCardUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
