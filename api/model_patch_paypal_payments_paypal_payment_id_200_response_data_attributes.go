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

// PATCHPaypalPaymentsPaypalPaymentId200ResponseDataAttributes struct for PATCHPaypalPaymentsPaypalPaymentId200ResponseDataAttributes
type PATCHPaypalPaymentsPaypalPaymentId200ResponseDataAttributes struct {
	// The id of the payer that PayPal passes in the return_url.
	PaypalPayerId *string `json:"paypal_payer_id,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewPATCHPaypalPaymentsPaypalPaymentId200ResponseDataAttributes instantiates a new PATCHPaypalPaymentsPaypalPaymentId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHPaypalPaymentsPaypalPaymentId200ResponseDataAttributes() *PATCHPaypalPaymentsPaypalPaymentId200ResponseDataAttributes {
	this := PATCHPaypalPaymentsPaypalPaymentId200ResponseDataAttributes{}
	return &this
}

// NewPATCHPaypalPaymentsPaypalPaymentId200ResponseDataAttributesWithDefaults instantiates a new PATCHPaypalPaymentsPaypalPaymentId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHPaypalPaymentsPaypalPaymentId200ResponseDataAttributesWithDefaults() *PATCHPaypalPaymentsPaypalPaymentId200ResponseDataAttributes {
	this := PATCHPaypalPaymentsPaypalPaymentId200ResponseDataAttributes{}
	return &this
}

// GetPaypalPayerId returns the PaypalPayerId field value if set, zero value otherwise.
func (o *PATCHPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) GetPaypalPayerId() string {
	if o == nil || o.PaypalPayerId == nil {
		var ret string
		return ret
	}
	return *o.PaypalPayerId
}

// GetPaypalPayerIdOk returns a tuple with the PaypalPayerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) GetPaypalPayerIdOk() (*string, bool) {
	if o == nil || o.PaypalPayerId == nil {
		return nil, false
	}
	return o.PaypalPayerId, true
}

// HasPaypalPayerId returns a boolean if a field has been set.
func (o *PATCHPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) HasPaypalPayerId() bool {
	if o != nil && o.PaypalPayerId != nil {
		return true
	}

	return false
}

// SetPaypalPayerId gets a reference to the given string and assigns it to the PaypalPayerId field.
func (o *PATCHPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) SetPaypalPayerId(v string) {
	o.PaypalPayerId = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *PATCHPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *PATCHPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *PATCHPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *PATCHPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PATCHPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PATCHPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o PATCHPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PaypalPayerId != nil {
		toSerialize["paypal_payer_id"] = o.PaypalPayerId
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

type NullablePATCHPaypalPaymentsPaypalPaymentId200ResponseDataAttributes struct {
	value *PATCHPaypalPaymentsPaypalPaymentId200ResponseDataAttributes
	isSet bool
}

func (v NullablePATCHPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) Get() *PATCHPaypalPaymentsPaypalPaymentId200ResponseDataAttributes {
	return v.value
}

func (v *NullablePATCHPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) Set(val *PATCHPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHPaypalPaymentsPaypalPaymentId200ResponseDataAttributes(val *PATCHPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) *NullablePATCHPaypalPaymentsPaypalPaymentId200ResponseDataAttributes {
	return &NullablePATCHPaypalPaymentsPaypalPaymentId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHPaypalPaymentsPaypalPaymentId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
