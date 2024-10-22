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

// checks if the BraintreeGatewayData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BraintreeGatewayData{}

// BraintreeGatewayData struct for BraintreeGatewayData
type BraintreeGatewayData struct {
	// The resource's type
	Type          interface{}                                                     `json:"type"`
	Attributes    GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes `json:"attributes"`
	Relationships *BraintreeGatewayDataRelationships                              `json:"relationships,omitempty"`
}

// NewBraintreeGatewayData instantiates a new BraintreeGatewayData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBraintreeGatewayData(type_ interface{}, attributes GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) *BraintreeGatewayData {
	this := BraintreeGatewayData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewBraintreeGatewayDataWithDefaults instantiates a new BraintreeGatewayData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBraintreeGatewayDataWithDefaults() *BraintreeGatewayData {
	this := BraintreeGatewayData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *BraintreeGatewayData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BraintreeGatewayData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BraintreeGatewayData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *BraintreeGatewayData) GetAttributes() GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes {
	if o == nil {
		var ret GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *BraintreeGatewayData) GetAttributesOk() (*GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *BraintreeGatewayData) SetAttributes(v GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *BraintreeGatewayData) GetRelationships() BraintreeGatewayDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret BraintreeGatewayDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BraintreeGatewayData) GetRelationshipsOk() (*BraintreeGatewayDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *BraintreeGatewayData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given BraintreeGatewayDataRelationships and assigns it to the Relationships field.
func (o *BraintreeGatewayData) SetRelationships(v BraintreeGatewayDataRelationships) {
	o.Relationships = &v
}

func (o BraintreeGatewayData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BraintreeGatewayData) ToMap() (map[string]interface{}, error) {
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

type NullableBraintreeGatewayData struct {
	value *BraintreeGatewayData
	isSet bool
}

func (v NullableBraintreeGatewayData) Get() *BraintreeGatewayData {
	return v.value
}

func (v *NullableBraintreeGatewayData) Set(val *BraintreeGatewayData) {
	v.value = val
	v.isSet = true
}

func (v NullableBraintreeGatewayData) IsSet() bool {
	return v.isSet
}

func (v *NullableBraintreeGatewayData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBraintreeGatewayData(val *BraintreeGatewayData) *NullableBraintreeGatewayData {
	return &NullableBraintreeGatewayData{value: val, isSet: true}
}

func (v NullableBraintreeGatewayData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBraintreeGatewayData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
