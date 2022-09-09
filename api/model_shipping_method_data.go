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

// ShippingMethodData struct for ShippingMethodData
type ShippingMethodData struct {
	// The resource's type
	Type string `json:"type"`
	Attributes GETShippingMethods200ResponseDataInnerAttributes `json:"attributes"`
	Relationships *ShippingMethodDataRelationships `json:"relationships,omitempty"`
}

// NewShippingMethodData instantiates a new ShippingMethodData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingMethodData(type_ string, attributes GETShippingMethods200ResponseDataInnerAttributes) *ShippingMethodData {
	this := ShippingMethodData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewShippingMethodDataWithDefaults instantiates a new ShippingMethodData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingMethodDataWithDefaults() *ShippingMethodData {
	this := ShippingMethodData{}
	var type_ string = "shipping_methods"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *ShippingMethodData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ShippingMethodData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ShippingMethodData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *ShippingMethodData) GetAttributes() GETShippingMethods200ResponseDataInnerAttributes {
	if o == nil {
		var ret GETShippingMethods200ResponseDataInnerAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ShippingMethodData) GetAttributesOk() (*GETShippingMethods200ResponseDataInnerAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *ShippingMethodData) SetAttributes(v GETShippingMethods200ResponseDataInnerAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ShippingMethodData) GetRelationships() ShippingMethodDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret ShippingMethodDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodData) GetRelationshipsOk() (*ShippingMethodDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ShippingMethodData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given ShippingMethodDataRelationships and assigns it to the Relationships field.
func (o *ShippingMethodData) SetRelationships(v ShippingMethodDataRelationships) {
	o.Relationships = &v
}

func (o ShippingMethodData) MarshalJSON() ([]byte, error) {
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

type NullableShippingMethodData struct {
	value *ShippingMethodData
	isSet bool
}

func (v NullableShippingMethodData) Get() *ShippingMethodData {
	return v.value
}

func (v *NullableShippingMethodData) Set(val *ShippingMethodData) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingMethodData) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingMethodData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingMethodData(val *ShippingMethodData) *NullableShippingMethodData {
	return &NullableShippingMethodData{value: val, isSet: true}
}

func (v NullableShippingMethodData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingMethodData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


