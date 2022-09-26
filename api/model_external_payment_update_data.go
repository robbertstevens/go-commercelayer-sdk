/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ExternalPaymentUpdateData struct for ExternalPaymentUpdateData
type ExternalPaymentUpdateData struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id            string                                                          `json:"id"`
	Attributes    PATCHExternalPaymentsExternalPaymentId200ResponseDataAttributes `json:"attributes"`
	Relationships *AdyenPaymentUpdateDataRelationships                            `json:"relationships,omitempty"`
}

// NewExternalPaymentUpdateData instantiates a new ExternalPaymentUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalPaymentUpdateData(type_ string, id string, attributes PATCHExternalPaymentsExternalPaymentId200ResponseDataAttributes) *ExternalPaymentUpdateData {
	this := ExternalPaymentUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewExternalPaymentUpdateDataWithDefaults instantiates a new ExternalPaymentUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalPaymentUpdateDataWithDefaults() *ExternalPaymentUpdateData {
	this := ExternalPaymentUpdateData{}
	return &this
}

// GetType returns the Type field value
func (o *ExternalPaymentUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ExternalPaymentUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ExternalPaymentUpdateData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *ExternalPaymentUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ExternalPaymentUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ExternalPaymentUpdateData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *ExternalPaymentUpdateData) GetAttributes() PATCHExternalPaymentsExternalPaymentId200ResponseDataAttributes {
	if o == nil {
		var ret PATCHExternalPaymentsExternalPaymentId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ExternalPaymentUpdateData) GetAttributesOk() (*PATCHExternalPaymentsExternalPaymentId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *ExternalPaymentUpdateData) SetAttributes(v PATCHExternalPaymentsExternalPaymentId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ExternalPaymentUpdateData) GetRelationships() AdyenPaymentUpdateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret AdyenPaymentUpdateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalPaymentUpdateData) GetRelationshipsOk() (*AdyenPaymentUpdateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ExternalPaymentUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AdyenPaymentUpdateDataRelationships and assigns it to the Relationships field.
func (o *ExternalPaymentUpdateData) SetRelationships(v AdyenPaymentUpdateDataRelationships) {
	o.Relationships = &v
}

func (o ExternalPaymentUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableExternalPaymentUpdateData struct {
	value *ExternalPaymentUpdateData
	isSet bool
}

func (v NullableExternalPaymentUpdateData) Get() *ExternalPaymentUpdateData {
	return v.value
}

func (v *NullableExternalPaymentUpdateData) Set(val *ExternalPaymentUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalPaymentUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalPaymentUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalPaymentUpdateData(val *ExternalPaymentUpdateData) *NullableExternalPaymentUpdateData {
	return &NullableExternalPaymentUpdateData{value: val, isSet: true}
}

func (v NullableExternalPaymentUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalPaymentUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
