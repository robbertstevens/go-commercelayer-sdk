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

// checks if the AuthorizationData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizationData{}

// AuthorizationData struct for AuthorizationData
type AuthorizationData struct {
	// The resource's type
	Type          interface{}                                               `json:"type"`
	Attributes    GETAuthorizationsAuthorizationId200ResponseDataAttributes `json:"attributes"`
	Relationships *AuthorizationDataRelationships                           `json:"relationships,omitempty"`
}

// NewAuthorizationData instantiates a new AuthorizationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationData(type_ interface{}, attributes GETAuthorizationsAuthorizationId200ResponseDataAttributes) *AuthorizationData {
	this := AuthorizationData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewAuthorizationDataWithDefaults instantiates a new AuthorizationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationDataWithDefaults() *AuthorizationData {
	this := AuthorizationData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *AuthorizationData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthorizationData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AuthorizationData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *AuthorizationData) GetAttributes() GETAuthorizationsAuthorizationId200ResponseDataAttributes {
	if o == nil {
		var ret GETAuthorizationsAuthorizationId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AuthorizationData) GetAttributesOk() (*GETAuthorizationsAuthorizationId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *AuthorizationData) SetAttributes(v GETAuthorizationsAuthorizationId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AuthorizationData) GetRelationships() AuthorizationDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret AuthorizationDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationData) GetRelationshipsOk() (*AuthorizationDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AuthorizationData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AuthorizationDataRelationships and assigns it to the Relationships field.
func (o *AuthorizationData) SetRelationships(v AuthorizationDataRelationships) {
	o.Relationships = &v
}

func (o AuthorizationData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizationData) ToMap() (map[string]interface{}, error) {
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

type NullableAuthorizationData struct {
	value *AuthorizationData
	isSet bool
}

func (v NullableAuthorizationData) Get() *AuthorizationData {
	return v.value
}

func (v *NullableAuthorizationData) Set(val *AuthorizationData) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationData) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationData(val *AuthorizationData) *NullableAuthorizationData {
	return &NullableAuthorizationData{value: val, isSet: true}
}

func (v NullableAuthorizationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
