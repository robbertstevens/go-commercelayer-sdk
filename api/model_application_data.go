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

// ApplicationData struct for ApplicationData
type ApplicationData struct {
	// The resource's type
	Type          string                                               `json:"type"`
	Attributes    GETApplicationApplicationId200ResponseDataAttributes `json:"attributes"`
	Relationships map[string]interface{}                               `json:"relationships,omitempty"`
}

// NewApplicationData instantiates a new ApplicationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationData(type_ string, attributes GETApplicationApplicationId200ResponseDataAttributes) *ApplicationData {
	this := ApplicationData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewApplicationDataWithDefaults instantiates a new ApplicationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationDataWithDefaults() *ApplicationData {
	this := ApplicationData{}
	return &this
}

// GetType returns the Type field value
func (o *ApplicationData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ApplicationData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ApplicationData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *ApplicationData) GetAttributes() GETApplicationApplicationId200ResponseDataAttributes {
	if o == nil {
		var ret GETApplicationApplicationId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ApplicationData) GetAttributesOk() (*GETApplicationApplicationId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *ApplicationData) SetAttributes(v GETApplicationApplicationId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ApplicationData) GetRelationships() map[string]interface{} {
	if o == nil || o.Relationships == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationData) GetRelationshipsOk() (map[string]interface{}, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ApplicationData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]interface{} and assigns it to the Relationships field.
func (o *ApplicationData) SetRelationships(v map[string]interface{}) {
	o.Relationships = v
}

func (o ApplicationData) MarshalJSON() ([]byte, error) {
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

type NullableApplicationData struct {
	value *ApplicationData
	isSet bool
}

func (v NullableApplicationData) Get() *ApplicationData {
	return v.value
}

func (v *NullableApplicationData) Set(val *ApplicationData) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationData) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationData(val *ApplicationData) *NullableApplicationData {
	return &NullableApplicationData{value: val, isSet: true}
}

func (v NullableApplicationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
