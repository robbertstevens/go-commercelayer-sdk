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

// checks if the ShippingCategoryData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShippingCategoryData{}

// ShippingCategoryData struct for ShippingCategoryData
type ShippingCategoryData struct {
	// The resource's type
	Type          interface{}                                                      `json:"type"`
	Attributes    GETShippingCategoriesShippingCategoryId200ResponseDataAttributes `json:"attributes"`
	Relationships *ShippingCategoryDataRelationships                               `json:"relationships,omitempty"`
}

// NewShippingCategoryData instantiates a new ShippingCategoryData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingCategoryData(type_ interface{}, attributes GETShippingCategoriesShippingCategoryId200ResponseDataAttributes) *ShippingCategoryData {
	this := ShippingCategoryData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewShippingCategoryDataWithDefaults instantiates a new ShippingCategoryData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingCategoryDataWithDefaults() *ShippingCategoryData {
	this := ShippingCategoryData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ShippingCategoryData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ShippingCategoryData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ShippingCategoryData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *ShippingCategoryData) GetAttributes() GETShippingCategoriesShippingCategoryId200ResponseDataAttributes {
	if o == nil {
		var ret GETShippingCategoriesShippingCategoryId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ShippingCategoryData) GetAttributesOk() (*GETShippingCategoriesShippingCategoryId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *ShippingCategoryData) SetAttributes(v GETShippingCategoriesShippingCategoryId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ShippingCategoryData) GetRelationships() ShippingCategoryDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret ShippingCategoryDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingCategoryData) GetRelationshipsOk() (*ShippingCategoryDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ShippingCategoryData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given ShippingCategoryDataRelationships and assigns it to the Relationships field.
func (o *ShippingCategoryData) SetRelationships(v ShippingCategoryDataRelationships) {
	o.Relationships = &v
}

func (o ShippingCategoryData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShippingCategoryData) ToMap() (map[string]interface{}, error) {
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

type NullableShippingCategoryData struct {
	value *ShippingCategoryData
	isSet bool
}

func (v NullableShippingCategoryData) Get() *ShippingCategoryData {
	return v.value
}

func (v *NullableShippingCategoryData) Set(val *ShippingCategoryData) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingCategoryData) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingCategoryData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingCategoryData(val *ShippingCategoryData) *NullableShippingCategoryData {
	return &NullableShippingCategoryData{value: val, isSet: true}
}

func (v NullableShippingCategoryData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingCategoryData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
