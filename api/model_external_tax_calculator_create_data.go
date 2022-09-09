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

// ExternalTaxCalculatorCreateData struct for ExternalTaxCalculatorCreateData
type ExternalTaxCalculatorCreateData struct {
	// The resource's type
	Type string `json:"type"`
	Attributes POSTExternalTaxCalculators201ResponseDataAttributes `json:"attributes"`
	Relationships *AvalaraAccountCreateDataRelationships `json:"relationships,omitempty"`
}

// NewExternalTaxCalculatorCreateData instantiates a new ExternalTaxCalculatorCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalTaxCalculatorCreateData(type_ string, attributes POSTExternalTaxCalculators201ResponseDataAttributes) *ExternalTaxCalculatorCreateData {
	this := ExternalTaxCalculatorCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewExternalTaxCalculatorCreateDataWithDefaults instantiates a new ExternalTaxCalculatorCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalTaxCalculatorCreateDataWithDefaults() *ExternalTaxCalculatorCreateData {
	this := ExternalTaxCalculatorCreateData{}
	var type_ string = "external_tax_calculators"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *ExternalTaxCalculatorCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ExternalTaxCalculatorCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ExternalTaxCalculatorCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *ExternalTaxCalculatorCreateData) GetAttributes() POSTExternalTaxCalculators201ResponseDataAttributes {
	if o == nil {
		var ret POSTExternalTaxCalculators201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ExternalTaxCalculatorCreateData) GetAttributesOk() (*POSTExternalTaxCalculators201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *ExternalTaxCalculatorCreateData) SetAttributes(v POSTExternalTaxCalculators201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ExternalTaxCalculatorCreateData) GetRelationships() AvalaraAccountCreateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret AvalaraAccountCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalTaxCalculatorCreateData) GetRelationshipsOk() (*AvalaraAccountCreateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ExternalTaxCalculatorCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AvalaraAccountCreateDataRelationships and assigns it to the Relationships field.
func (o *ExternalTaxCalculatorCreateData) SetRelationships(v AvalaraAccountCreateDataRelationships) {
	o.Relationships = &v
}

func (o ExternalTaxCalculatorCreateData) MarshalJSON() ([]byte, error) {
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

type NullableExternalTaxCalculatorCreateData struct {
	value *ExternalTaxCalculatorCreateData
	isSet bool
}

func (v NullableExternalTaxCalculatorCreateData) Get() *ExternalTaxCalculatorCreateData {
	return v.value
}

func (v *NullableExternalTaxCalculatorCreateData) Set(val *ExternalTaxCalculatorCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalTaxCalculatorCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalTaxCalculatorCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalTaxCalculatorCreateData(val *ExternalTaxCalculatorCreateData) *NullableExternalTaxCalculatorCreateData {
	return &NullableExternalTaxCalculatorCreateData{value: val, isSet: true}
}

func (v NullableExternalTaxCalculatorCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalTaxCalculatorCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


