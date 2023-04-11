/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 4.1.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the POSTCouponsRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTCouponsRequestData{}

// POSTCouponsRequestData struct for POSTCouponsRequestData
type POSTCouponsRequestData struct {
	// The resource's type
	Type          interface{}                          `json:"type"`
	Attributes    POSTCouponsRequestDataAttributes     `json:"attributes"`
	Relationships *POSTCouponsRequestDataRelationships `json:"relationships,omitempty"`
}

// NewPOSTCouponsRequestData instantiates a new POSTCouponsRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTCouponsRequestData(type_ interface{}, attributes POSTCouponsRequestDataAttributes) *POSTCouponsRequestData {
	this := POSTCouponsRequestData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewPOSTCouponsRequestDataWithDefaults instantiates a new POSTCouponsRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTCouponsRequestDataWithDefaults() *POSTCouponsRequestData {
	this := POSTCouponsRequestData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTCouponsRequestData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTCouponsRequestData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *POSTCouponsRequestData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *POSTCouponsRequestData) GetAttributes() POSTCouponsRequestDataAttributes {
	if o == nil {
		var ret POSTCouponsRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *POSTCouponsRequestData) GetAttributesOk() (*POSTCouponsRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *POSTCouponsRequestData) SetAttributes(v POSTCouponsRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *POSTCouponsRequestData) GetRelationships() POSTCouponsRequestDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret POSTCouponsRequestDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCouponsRequestData) GetRelationshipsOk() (*POSTCouponsRequestDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *POSTCouponsRequestData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given POSTCouponsRequestDataRelationships and assigns it to the Relationships field.
func (o *POSTCouponsRequestData) SetRelationships(v POSTCouponsRequestDataRelationships) {
	o.Relationships = &v
}

func (o POSTCouponsRequestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTCouponsRequestData) ToMap() (map[string]interface{}, error) {
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

type NullablePOSTCouponsRequestData struct {
	value *POSTCouponsRequestData
	isSet bool
}

func (v NullablePOSTCouponsRequestData) Get() *POSTCouponsRequestData {
	return v.value
}

func (v *NullablePOSTCouponsRequestData) Set(val *POSTCouponsRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTCouponsRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTCouponsRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTCouponsRequestData(val *POSTCouponsRequestData) *NullablePOSTCouponsRequestData {
	return &NullablePOSTCouponsRequestData{value: val, isSet: true}
}

func (v NullablePOSTCouponsRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTCouponsRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
