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

// POSTAdjustments201ResponseDataAttributes struct for POSTAdjustments201ResponseDataAttributes
type POSTAdjustments201ResponseDataAttributes struct {
	// The adjustment name
	Name string `json:"name"`
	// The international 3-letter currency code as defined by the ISO 4217 standard.
	CurrencyCode string `json:"currency_code"`
	// The adjustment amount, in cents.
	AmountCents int32 `json:"amount_cents"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewPOSTAdjustments201ResponseDataAttributes instantiates a new POSTAdjustments201ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTAdjustments201ResponseDataAttributes(name string, currencyCode string, amountCents int32) *POSTAdjustments201ResponseDataAttributes {
	this := POSTAdjustments201ResponseDataAttributes{}
	this.Name = name
	this.CurrencyCode = currencyCode
	this.AmountCents = amountCents
	return &this
}

// NewPOSTAdjustments201ResponseDataAttributesWithDefaults instantiates a new POSTAdjustments201ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTAdjustments201ResponseDataAttributesWithDefaults() *POSTAdjustments201ResponseDataAttributes {
	this := POSTAdjustments201ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value
func (o *POSTAdjustments201ResponseDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *POSTAdjustments201ResponseDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *POSTAdjustments201ResponseDataAttributes) SetName(v string) {
	o.Name = v
}

// GetCurrencyCode returns the CurrencyCode field value
func (o *POSTAdjustments201ResponseDataAttributes) GetCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value
// and a boolean to check if the value has been set.
func (o *POSTAdjustments201ResponseDataAttributes) GetCurrencyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// SetCurrencyCode sets field value
func (o *POSTAdjustments201ResponseDataAttributes) SetCurrencyCode(v string) {
	o.CurrencyCode = v
}

// GetAmountCents returns the AmountCents field value
func (o *POSTAdjustments201ResponseDataAttributes) GetAmountCents() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AmountCents
}

// GetAmountCentsOk returns a tuple with the AmountCents field value
// and a boolean to check if the value has been set.
func (o *POSTAdjustments201ResponseDataAttributes) GetAmountCentsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountCents, true
}

// SetAmountCents sets field value
func (o *POSTAdjustments201ResponseDataAttributes) SetAmountCents(v int32) {
	o.AmountCents = v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *POSTAdjustments201ResponseDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTAdjustments201ResponseDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTAdjustments201ResponseDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *POSTAdjustments201ResponseDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *POSTAdjustments201ResponseDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTAdjustments201ResponseDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTAdjustments201ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *POSTAdjustments201ResponseDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *POSTAdjustments201ResponseDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTAdjustments201ResponseDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTAdjustments201ResponseDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *POSTAdjustments201ResponseDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o POSTAdjustments201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	if true {
		toSerialize["amount_cents"] = o.AmountCents
	}
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
	}
	if o.ReferenceOrigin != nil {
		toSerialize["reference_origin"] = o.ReferenceOrigin
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullablePOSTAdjustments201ResponseDataAttributes struct {
	value *POSTAdjustments201ResponseDataAttributes
	isSet bool
}

func (v NullablePOSTAdjustments201ResponseDataAttributes) Get() *POSTAdjustments201ResponseDataAttributes {
	return v.value
}

func (v *NullablePOSTAdjustments201ResponseDataAttributes) Set(val *POSTAdjustments201ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTAdjustments201ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTAdjustments201ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTAdjustments201ResponseDataAttributes(val *POSTAdjustments201ResponseDataAttributes) *NullablePOSTAdjustments201ResponseDataAttributes {
	return &NullablePOSTAdjustments201ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTAdjustments201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTAdjustments201ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


