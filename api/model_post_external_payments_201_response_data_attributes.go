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

// POSTExternalPayments201ResponseDataAttributes struct for POSTExternalPayments201ResponseDataAttributes
type POSTExternalPayments201ResponseDataAttributes struct {
	// The payment source token, as generated by the external gateway SDK. Credit Card numbers are rejected.
	PaymentSourceToken string `json:"payment_source_token"`
	// External payment options.
	Options map[string]interface{} `json:"options,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewPOSTExternalPayments201ResponseDataAttributes instantiates a new POSTExternalPayments201ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTExternalPayments201ResponseDataAttributes(paymentSourceToken string) *POSTExternalPayments201ResponseDataAttributes {
	this := POSTExternalPayments201ResponseDataAttributes{}
	this.PaymentSourceToken = paymentSourceToken
	return &this
}

// NewPOSTExternalPayments201ResponseDataAttributesWithDefaults instantiates a new POSTExternalPayments201ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTExternalPayments201ResponseDataAttributesWithDefaults() *POSTExternalPayments201ResponseDataAttributes {
	this := POSTExternalPayments201ResponseDataAttributes{}
	return &this
}

// GetPaymentSourceToken returns the PaymentSourceToken field value
func (o *POSTExternalPayments201ResponseDataAttributes) GetPaymentSourceToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentSourceToken
}

// GetPaymentSourceTokenOk returns a tuple with the PaymentSourceToken field value
// and a boolean to check if the value has been set.
func (o *POSTExternalPayments201ResponseDataAttributes) GetPaymentSourceTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentSourceToken, true
}

// SetPaymentSourceToken sets field value
func (o *POSTExternalPayments201ResponseDataAttributes) SetPaymentSourceToken(v string) {
	o.PaymentSourceToken = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *POSTExternalPayments201ResponseDataAttributes) GetOptions() map[string]interface{} {
	if o == nil || o.Options == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTExternalPayments201ResponseDataAttributes) GetOptionsOk() (map[string]interface{}, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *POSTExternalPayments201ResponseDataAttributes) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given map[string]interface{} and assigns it to the Options field.
func (o *POSTExternalPayments201ResponseDataAttributes) SetOptions(v map[string]interface{}) {
	o.Options = v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *POSTExternalPayments201ResponseDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTExternalPayments201ResponseDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTExternalPayments201ResponseDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *POSTExternalPayments201ResponseDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *POSTExternalPayments201ResponseDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTExternalPayments201ResponseDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTExternalPayments201ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *POSTExternalPayments201ResponseDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *POSTExternalPayments201ResponseDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTExternalPayments201ResponseDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTExternalPayments201ResponseDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *POSTExternalPayments201ResponseDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o POSTExternalPayments201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["payment_source_token"] = o.PaymentSourceToken
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
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

type NullablePOSTExternalPayments201ResponseDataAttributes struct {
	value *POSTExternalPayments201ResponseDataAttributes
	isSet bool
}

func (v NullablePOSTExternalPayments201ResponseDataAttributes) Get() *POSTExternalPayments201ResponseDataAttributes {
	return v.value
}

func (v *NullablePOSTExternalPayments201ResponseDataAttributes) Set(val *POSTExternalPayments201ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTExternalPayments201ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTExternalPayments201ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTExternalPayments201ResponseDataAttributes(val *POSTExternalPayments201ResponseDataAttributes) *NullablePOSTExternalPayments201ResponseDataAttributes {
	return &NullablePOSTExternalPayments201ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTExternalPayments201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTExternalPayments201ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
