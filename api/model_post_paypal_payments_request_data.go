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

// checks if the POSTPaypalPaymentsRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTPaypalPaymentsRequestData{}

// POSTPaypalPaymentsRequestData struct for POSTPaypalPaymentsRequestData
type POSTPaypalPaymentsRequestData struct {
	// The resource's type
	Type          interface{}                                `json:"type"`
	Attributes    POSTPaypalPaymentsRequestDataAttributes    `json:"attributes"`
	Relationships *POSTAdyenPaymentsRequestDataRelationships `json:"relationships,omitempty"`
}

// NewPOSTPaypalPaymentsRequestData instantiates a new POSTPaypalPaymentsRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTPaypalPaymentsRequestData(type_ interface{}, attributes POSTPaypalPaymentsRequestDataAttributes) *POSTPaypalPaymentsRequestData {
	this := POSTPaypalPaymentsRequestData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewPOSTPaypalPaymentsRequestDataWithDefaults instantiates a new POSTPaypalPaymentsRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTPaypalPaymentsRequestDataWithDefaults() *POSTPaypalPaymentsRequestData {
	this := POSTPaypalPaymentsRequestData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTPaypalPaymentsRequestData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPaypalPaymentsRequestData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *POSTPaypalPaymentsRequestData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *POSTPaypalPaymentsRequestData) GetAttributes() POSTPaypalPaymentsRequestDataAttributes {
	if o == nil {
		var ret POSTPaypalPaymentsRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *POSTPaypalPaymentsRequestData) GetAttributesOk() (*POSTPaypalPaymentsRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *POSTPaypalPaymentsRequestData) SetAttributes(v POSTPaypalPaymentsRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *POSTPaypalPaymentsRequestData) GetRelationships() POSTAdyenPaymentsRequestDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret POSTAdyenPaymentsRequestDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTPaypalPaymentsRequestData) GetRelationshipsOk() (*POSTAdyenPaymentsRequestDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *POSTPaypalPaymentsRequestData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given POSTAdyenPaymentsRequestDataRelationships and assigns it to the Relationships field.
func (o *POSTPaypalPaymentsRequestData) SetRelationships(v POSTAdyenPaymentsRequestDataRelationships) {
	o.Relationships = &v
}

func (o POSTPaypalPaymentsRequestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTPaypalPaymentsRequestData) ToMap() (map[string]interface{}, error) {
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

type NullablePOSTPaypalPaymentsRequestData struct {
	value *POSTPaypalPaymentsRequestData
	isSet bool
}

func (v NullablePOSTPaypalPaymentsRequestData) Get() *POSTPaypalPaymentsRequestData {
	return v.value
}

func (v *NullablePOSTPaypalPaymentsRequestData) Set(val *POSTPaypalPaymentsRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTPaypalPaymentsRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTPaypalPaymentsRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTPaypalPaymentsRequestData(val *POSTPaypalPaymentsRequestData) *NullablePOSTPaypalPaymentsRequestData {
	return &NullablePOSTPaypalPaymentsRequestData{value: val, isSet: true}
}

func (v NullablePOSTPaypalPaymentsRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTPaypalPaymentsRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
