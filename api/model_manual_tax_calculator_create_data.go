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

// ManualTaxCalculatorCreateData struct for ManualTaxCalculatorCreateData
type ManualTaxCalculatorCreateData struct {
	// The resource's type
	Type          string                                            `json:"type"`
	Attributes    POSTManualTaxCalculators201ResponseDataAttributes `json:"attributes"`
	Relationships *ManualTaxCalculatorCreateDataRelationships       `json:"relationships,omitempty"`
}

// NewManualTaxCalculatorCreateData instantiates a new ManualTaxCalculatorCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManualTaxCalculatorCreateData(type_ string, attributes POSTManualTaxCalculators201ResponseDataAttributes) *ManualTaxCalculatorCreateData {
	this := ManualTaxCalculatorCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewManualTaxCalculatorCreateDataWithDefaults instantiates a new ManualTaxCalculatorCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManualTaxCalculatorCreateDataWithDefaults() *ManualTaxCalculatorCreateData {
	this := ManualTaxCalculatorCreateData{}
	return &this
}

// GetType returns the Type field value
func (o *ManualTaxCalculatorCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ManualTaxCalculatorCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ManualTaxCalculatorCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *ManualTaxCalculatorCreateData) GetAttributes() POSTManualTaxCalculators201ResponseDataAttributes {
	if o == nil {
		var ret POSTManualTaxCalculators201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ManualTaxCalculatorCreateData) GetAttributesOk() (*POSTManualTaxCalculators201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *ManualTaxCalculatorCreateData) SetAttributes(v POSTManualTaxCalculators201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ManualTaxCalculatorCreateData) GetRelationships() ManualTaxCalculatorCreateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret ManualTaxCalculatorCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualTaxCalculatorCreateData) GetRelationshipsOk() (*ManualTaxCalculatorCreateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ManualTaxCalculatorCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given ManualTaxCalculatorCreateDataRelationships and assigns it to the Relationships field.
func (o *ManualTaxCalculatorCreateData) SetRelationships(v ManualTaxCalculatorCreateDataRelationships) {
	o.Relationships = &v
}

func (o ManualTaxCalculatorCreateData) MarshalJSON() ([]byte, error) {
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

type NullableManualTaxCalculatorCreateData struct {
	value *ManualTaxCalculatorCreateData
	isSet bool
}

func (v NullableManualTaxCalculatorCreateData) Get() *ManualTaxCalculatorCreateData {
	return v.value
}

func (v *NullableManualTaxCalculatorCreateData) Set(val *ManualTaxCalculatorCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableManualTaxCalculatorCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableManualTaxCalculatorCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManualTaxCalculatorCreateData(val *ManualTaxCalculatorCreateData) *NullableManualTaxCalculatorCreateData {
	return &NullableManualTaxCalculatorCreateData{value: val, isSet: true}
}

func (v NullableManualTaxCalculatorCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManualTaxCalculatorCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
