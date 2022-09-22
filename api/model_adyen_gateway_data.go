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

// AdyenGatewayData struct for AdyenGatewayData
type AdyenGatewayData struct {
	// The resource's type
	Type          string                                         `json:"type"`
	Attributes    GETAdyenGateways200ResponseDataInnerAttributes `json:"attributes"`
	Relationships *AdyenGatewayDataRelationships                 `json:"relationships,omitempty"`
}

// NewAdyenGatewayData instantiates a new AdyenGatewayData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdyenGatewayData(type_ string, attributes GETAdyenGateways200ResponseDataInnerAttributes) *AdyenGatewayData {
	this := AdyenGatewayData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewAdyenGatewayDataWithDefaults instantiates a new AdyenGatewayData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdyenGatewayDataWithDefaults() *AdyenGatewayData {
	this := AdyenGatewayData{}
	return &this
}

// GetType returns the Type field value
func (o *AdyenGatewayData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AdyenGatewayData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AdyenGatewayData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *AdyenGatewayData) GetAttributes() GETAdyenGateways200ResponseDataInnerAttributes {
	if o == nil {
		var ret GETAdyenGateways200ResponseDataInnerAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AdyenGatewayData) GetAttributesOk() (*GETAdyenGateways200ResponseDataInnerAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *AdyenGatewayData) SetAttributes(v GETAdyenGateways200ResponseDataInnerAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AdyenGatewayData) GetRelationships() AdyenGatewayDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret AdyenGatewayDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdyenGatewayData) GetRelationshipsOk() (*AdyenGatewayDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AdyenGatewayData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AdyenGatewayDataRelationships and assigns it to the Relationships field.
func (o *AdyenGatewayData) SetRelationships(v AdyenGatewayDataRelationships) {
	o.Relationships = &v
}

func (o AdyenGatewayData) MarshalJSON() ([]byte, error) {
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

type NullableAdyenGatewayData struct {
	value *AdyenGatewayData
	isSet bool
}

func (v NullableAdyenGatewayData) Get() *AdyenGatewayData {
	return v.value
}

func (v *NullableAdyenGatewayData) Set(val *AdyenGatewayData) {
	v.value = val
	v.isSet = true
}

func (v NullableAdyenGatewayData) IsSet() bool {
	return v.isSet
}

func (v *NullableAdyenGatewayData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdyenGatewayData(val *AdyenGatewayData) *NullableAdyenGatewayData {
	return &NullableAdyenGatewayData{value: val, isSet: true}
}

func (v NullableAdyenGatewayData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdyenGatewayData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
