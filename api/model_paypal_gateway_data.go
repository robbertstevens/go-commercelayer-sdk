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

// PaypalGatewayData struct for PaypalGatewayData
type PaypalGatewayData struct {
	// The resource's type
	Type          string                                          `json:"type"`
	Attributes    GETKlarnaGateways200ResponseDataInnerAttributes `json:"attributes"`
	Relationships *PaypalGatewayDataRelationships                 `json:"relationships,omitempty"`
}

// NewPaypalGatewayData instantiates a new PaypalGatewayData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaypalGatewayData(type_ string, attributes GETKlarnaGateways200ResponseDataInnerAttributes) *PaypalGatewayData {
	this := PaypalGatewayData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewPaypalGatewayDataWithDefaults instantiates a new PaypalGatewayData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaypalGatewayDataWithDefaults() *PaypalGatewayData {
	this := PaypalGatewayData{}
	return &this
}

// GetType returns the Type field value
func (o *PaypalGatewayData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PaypalGatewayData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PaypalGatewayData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *PaypalGatewayData) GetAttributes() GETKlarnaGateways200ResponseDataInnerAttributes {
	if o == nil {
		var ret GETKlarnaGateways200ResponseDataInnerAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *PaypalGatewayData) GetAttributesOk() (*GETKlarnaGateways200ResponseDataInnerAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *PaypalGatewayData) SetAttributes(v GETKlarnaGateways200ResponseDataInnerAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PaypalGatewayData) GetRelationships() PaypalGatewayDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret PaypalGatewayDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaypalGatewayData) GetRelationshipsOk() (*PaypalGatewayDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PaypalGatewayData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given PaypalGatewayDataRelationships and assigns it to the Relationships field.
func (o *PaypalGatewayData) SetRelationships(v PaypalGatewayDataRelationships) {
	o.Relationships = &v
}

func (o PaypalGatewayData) MarshalJSON() ([]byte, error) {
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

type NullablePaypalGatewayData struct {
	value *PaypalGatewayData
	isSet bool
}

func (v NullablePaypalGatewayData) Get() *PaypalGatewayData {
	return v.value
}

func (v *NullablePaypalGatewayData) Set(val *PaypalGatewayData) {
	v.value = val
	v.isSet = true
}

func (v NullablePaypalGatewayData) IsSet() bool {
	return v.isSet
}

func (v *NullablePaypalGatewayData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaypalGatewayData(val *PaypalGatewayData) *NullablePaypalGatewayData {
	return &NullablePaypalGatewayData{value: val, isSet: true}
}

func (v NullablePaypalGatewayData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaypalGatewayData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
