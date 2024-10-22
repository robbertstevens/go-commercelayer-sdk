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

// checks if the ApplicationData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationData{}

// ApplicationData struct for ApplicationData
type ApplicationData struct {
	// The resource's type
	Type          interface{}                                          `json:"type"`
	Attributes    GETApplicationApplicationId200ResponseDataAttributes `json:"attributes"`
	Relationships interface{}                                          `json:"relationships,omitempty"`
}

// NewApplicationData instantiates a new ApplicationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationData(type_ interface{}, attributes GETApplicationApplicationId200ResponseDataAttributes) *ApplicationData {
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
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ApplicationData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ApplicationData) SetType(v interface{}) {
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

// GetRelationships returns the Relationships field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationData) GetRelationships() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationData) GetRelationshipsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return &o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ApplicationData) HasRelationships() bool {
	if o != nil && IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given interface{} and assigns it to the Relationships field.
func (o *ApplicationData) SetRelationships(v interface{}) {
	o.Relationships = v
}

func (o ApplicationData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	toSerialize["attributes"] = o.Attributes
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
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
