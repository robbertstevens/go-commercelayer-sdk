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

// KlarnaGatewayData struct for KlarnaGatewayData
type KlarnaGatewayData struct {
	// The resource's type
	Type          string                                          `json:"type"`
	Attributes    GETKlarnaGateways200ResponseDataInnerAttributes `json:"attributes"`
	Relationships *KlarnaGatewayDataRelationships                 `json:"relationships,omitempty"`
}

// NewKlarnaGatewayData instantiates a new KlarnaGatewayData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKlarnaGatewayData(type_ string, attributes GETKlarnaGateways200ResponseDataInnerAttributes) *KlarnaGatewayData {
	this := KlarnaGatewayData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewKlarnaGatewayDataWithDefaults instantiates a new KlarnaGatewayData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKlarnaGatewayDataWithDefaults() *KlarnaGatewayData {
	this := KlarnaGatewayData{}
	return &this
}

// GetType returns the Type field value
func (o *KlarnaGatewayData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *KlarnaGatewayData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *KlarnaGatewayData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *KlarnaGatewayData) GetAttributes() GETKlarnaGateways200ResponseDataInnerAttributes {
	if o == nil {
		var ret GETKlarnaGateways200ResponseDataInnerAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *KlarnaGatewayData) GetAttributesOk() (*GETKlarnaGateways200ResponseDataInnerAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *KlarnaGatewayData) SetAttributes(v GETKlarnaGateways200ResponseDataInnerAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *KlarnaGatewayData) GetRelationships() KlarnaGatewayDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret KlarnaGatewayDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlarnaGatewayData) GetRelationshipsOk() (*KlarnaGatewayDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *KlarnaGatewayData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given KlarnaGatewayDataRelationships and assigns it to the Relationships field.
func (o *KlarnaGatewayData) SetRelationships(v KlarnaGatewayDataRelationships) {
	o.Relationships = &v
}

func (o KlarnaGatewayData) MarshalJSON() ([]byte, error) {
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

type NullableKlarnaGatewayData struct {
	value *KlarnaGatewayData
	isSet bool
}

func (v NullableKlarnaGatewayData) Get() *KlarnaGatewayData {
	return v.value
}

func (v *NullableKlarnaGatewayData) Set(val *KlarnaGatewayData) {
	v.value = val
	v.isSet = true
}

func (v NullableKlarnaGatewayData) IsSet() bool {
	return v.isSet
}

func (v *NullableKlarnaGatewayData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKlarnaGatewayData(val *KlarnaGatewayData) *NullableKlarnaGatewayData {
	return &NullableKlarnaGatewayData{value: val, isSet: true}
}

func (v NullableKlarnaGatewayData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKlarnaGatewayData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
