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

// checks if the BundleData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BundleData{}

// BundleData struct for BundleData
type BundleData struct {
	// The resource's type
	Type          interface{}                                 `json:"type"`
	Attributes    GETBundlesBundleId200ResponseDataAttributes `json:"attributes"`
	Relationships *BundleDataRelationships                    `json:"relationships,omitempty"`
}

// NewBundleData instantiates a new BundleData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBundleData(type_ interface{}, attributes GETBundlesBundleId200ResponseDataAttributes) *BundleData {
	this := BundleData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewBundleDataWithDefaults instantiates a new BundleData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBundleDataWithDefaults() *BundleData {
	this := BundleData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *BundleData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BundleData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BundleData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *BundleData) GetAttributes() GETBundlesBundleId200ResponseDataAttributes {
	if o == nil {
		var ret GETBundlesBundleId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *BundleData) GetAttributesOk() (*GETBundlesBundleId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *BundleData) SetAttributes(v GETBundlesBundleId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *BundleData) GetRelationships() BundleDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret BundleDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleData) GetRelationshipsOk() (*BundleDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *BundleData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given BundleDataRelationships and assigns it to the Relationships field.
func (o *BundleData) SetRelationships(v BundleDataRelationships) {
	o.Relationships = &v
}

func (o BundleData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BundleData) ToMap() (map[string]interface{}, error) {
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

type NullableBundleData struct {
	value *BundleData
	isSet bool
}

func (v NullableBundleData) Get() *BundleData {
	return v.value
}

func (v *NullableBundleData) Set(val *BundleData) {
	v.value = val
	v.isSet = true
}

func (v NullableBundleData) IsSet() bool {
	return v.isSet
}

func (v *NullableBundleData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundleData(val *BundleData) *NullableBundleData {
	return &NullableBundleData{value: val, isSet: true}
}

func (v NullableBundleData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundleData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
