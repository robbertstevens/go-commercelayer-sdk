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

// checks if the AxervePaymentData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AxervePaymentData{}

// AxervePaymentData struct for AxervePaymentData
type AxervePaymentData struct {
	// The resource's type
	Type          interface{}                                               `json:"type"`
	Attributes    GETAxervePaymentsAxervePaymentId200ResponseDataAttributes `json:"attributes"`
	Relationships *AdyenPaymentDataRelationships                            `json:"relationships,omitempty"`
}

// NewAxervePaymentData instantiates a new AxervePaymentData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAxervePaymentData(type_ interface{}, attributes GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) *AxervePaymentData {
	this := AxervePaymentData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewAxervePaymentDataWithDefaults instantiates a new AxervePaymentData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAxervePaymentDataWithDefaults() *AxervePaymentData {
	this := AxervePaymentData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *AxervePaymentData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AxervePaymentData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AxervePaymentData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *AxervePaymentData) GetAttributes() GETAxervePaymentsAxervePaymentId200ResponseDataAttributes {
	if o == nil {
		var ret GETAxervePaymentsAxervePaymentId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AxervePaymentData) GetAttributesOk() (*GETAxervePaymentsAxervePaymentId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *AxervePaymentData) SetAttributes(v GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AxervePaymentData) GetRelationships() AdyenPaymentDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret AdyenPaymentDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AxervePaymentData) GetRelationshipsOk() (*AdyenPaymentDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AxervePaymentData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AdyenPaymentDataRelationships and assigns it to the Relationships field.
func (o *AxervePaymentData) SetRelationships(v AdyenPaymentDataRelationships) {
	o.Relationships = &v
}

func (o AxervePaymentData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AxervePaymentData) ToMap() (map[string]interface{}, error) {
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

type NullableAxervePaymentData struct {
	value *AxervePaymentData
	isSet bool
}

func (v NullableAxervePaymentData) Get() *AxervePaymentData {
	return v.value
}

func (v *NullableAxervePaymentData) Set(val *AxervePaymentData) {
	v.value = val
	v.isSet = true
}

func (v NullableAxervePaymentData) IsSet() bool {
	return v.isSet
}

func (v *NullableAxervePaymentData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAxervePaymentData(val *AxervePaymentData) *NullableAxervePaymentData {
	return &NullableAxervePaymentData{value: val, isSet: true}
}

func (v NullableAxervePaymentData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAxervePaymentData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
