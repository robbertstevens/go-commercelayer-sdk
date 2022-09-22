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

// POSTCheckoutComPayments201ResponseDataAttributes struct for POSTCheckoutComPayments201ResponseDataAttributes
type POSTCheckoutComPayments201ResponseDataAttributes struct {
	// The payment source type.
	PaymentType string `json:"payment_type"`
	// The Checkout.com card or digital wallet token.
	Token string `json:"token"`
	// A payment session ID used to obtain the details.
	SessionId *string `json:"session_id,omitempty"`
	// The URL to redirect your customer upon 3DS succeeded authentication.
	SuccessUrl *string `json:"success_url,omitempty"`
	// The URL to redirect your customer upon 3DS failed authentication.
	FailureUrl *string `json:"failure_url,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewPOSTCheckoutComPayments201ResponseDataAttributes instantiates a new POSTCheckoutComPayments201ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTCheckoutComPayments201ResponseDataAttributes(paymentType string, token string) *POSTCheckoutComPayments201ResponseDataAttributes {
	this := POSTCheckoutComPayments201ResponseDataAttributes{}
	this.PaymentType = paymentType
	this.Token = token
	return &this
}

// NewPOSTCheckoutComPayments201ResponseDataAttributesWithDefaults instantiates a new POSTCheckoutComPayments201ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTCheckoutComPayments201ResponseDataAttributesWithDefaults() *POSTCheckoutComPayments201ResponseDataAttributes {
	this := POSTCheckoutComPayments201ResponseDataAttributes{}
	return &this
}

// GetPaymentType returns the PaymentType field value
func (o *POSTCheckoutComPayments201ResponseDataAttributes) GetPaymentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentType
}

// GetPaymentTypeOk returns a tuple with the PaymentType field value
// and a boolean to check if the value has been set.
func (o *POSTCheckoutComPayments201ResponseDataAttributes) GetPaymentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentType, true
}

// SetPaymentType sets field value
func (o *POSTCheckoutComPayments201ResponseDataAttributes) SetPaymentType(v string) {
	o.PaymentType = v
}

// GetToken returns the Token field value
func (o *POSTCheckoutComPayments201ResponseDataAttributes) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *POSTCheckoutComPayments201ResponseDataAttributes) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *POSTCheckoutComPayments201ResponseDataAttributes) SetToken(v string) {
	o.Token = v
}

// GetSessionId returns the SessionId field value if set, zero value otherwise.
func (o *POSTCheckoutComPayments201ResponseDataAttributes) GetSessionId() string {
	if o == nil || o.SessionId == nil {
		var ret string
		return ret
	}
	return *o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCheckoutComPayments201ResponseDataAttributes) GetSessionIdOk() (*string, bool) {
	if o == nil || o.SessionId == nil {
		return nil, false
	}
	return o.SessionId, true
}

// HasSessionId returns a boolean if a field has been set.
func (o *POSTCheckoutComPayments201ResponseDataAttributes) HasSessionId() bool {
	if o != nil && o.SessionId != nil {
		return true
	}

	return false
}

// SetSessionId gets a reference to the given string and assigns it to the SessionId field.
func (o *POSTCheckoutComPayments201ResponseDataAttributes) SetSessionId(v string) {
	o.SessionId = &v
}

// GetSuccessUrl returns the SuccessUrl field value if set, zero value otherwise.
func (o *POSTCheckoutComPayments201ResponseDataAttributes) GetSuccessUrl() string {
	if o == nil || o.SuccessUrl == nil {
		var ret string
		return ret
	}
	return *o.SuccessUrl
}

// GetSuccessUrlOk returns a tuple with the SuccessUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCheckoutComPayments201ResponseDataAttributes) GetSuccessUrlOk() (*string, bool) {
	if o == nil || o.SuccessUrl == nil {
		return nil, false
	}
	return o.SuccessUrl, true
}

// HasSuccessUrl returns a boolean if a field has been set.
func (o *POSTCheckoutComPayments201ResponseDataAttributes) HasSuccessUrl() bool {
	if o != nil && o.SuccessUrl != nil {
		return true
	}

	return false
}

// SetSuccessUrl gets a reference to the given string and assigns it to the SuccessUrl field.
func (o *POSTCheckoutComPayments201ResponseDataAttributes) SetSuccessUrl(v string) {
	o.SuccessUrl = &v
}

// GetFailureUrl returns the FailureUrl field value if set, zero value otherwise.
func (o *POSTCheckoutComPayments201ResponseDataAttributes) GetFailureUrl() string {
	if o == nil || o.FailureUrl == nil {
		var ret string
		return ret
	}
	return *o.FailureUrl
}

// GetFailureUrlOk returns a tuple with the FailureUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCheckoutComPayments201ResponseDataAttributes) GetFailureUrlOk() (*string, bool) {
	if o == nil || o.FailureUrl == nil {
		return nil, false
	}
	return o.FailureUrl, true
}

// HasFailureUrl returns a boolean if a field has been set.
func (o *POSTCheckoutComPayments201ResponseDataAttributes) HasFailureUrl() bool {
	if o != nil && o.FailureUrl != nil {
		return true
	}

	return false
}

// SetFailureUrl gets a reference to the given string and assigns it to the FailureUrl field.
func (o *POSTCheckoutComPayments201ResponseDataAttributes) SetFailureUrl(v string) {
	o.FailureUrl = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *POSTCheckoutComPayments201ResponseDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCheckoutComPayments201ResponseDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTCheckoutComPayments201ResponseDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *POSTCheckoutComPayments201ResponseDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *POSTCheckoutComPayments201ResponseDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCheckoutComPayments201ResponseDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTCheckoutComPayments201ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *POSTCheckoutComPayments201ResponseDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *POSTCheckoutComPayments201ResponseDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCheckoutComPayments201ResponseDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTCheckoutComPayments201ResponseDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *POSTCheckoutComPayments201ResponseDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o POSTCheckoutComPayments201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["payment_type"] = o.PaymentType
	}
	if true {
		toSerialize["token"] = o.Token
	}
	if o.SessionId != nil {
		toSerialize["session_id"] = o.SessionId
	}
	if o.SuccessUrl != nil {
		toSerialize["success_url"] = o.SuccessUrl
	}
	if o.FailureUrl != nil {
		toSerialize["failure_url"] = o.FailureUrl
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

type NullablePOSTCheckoutComPayments201ResponseDataAttributes struct {
	value *POSTCheckoutComPayments201ResponseDataAttributes
	isSet bool
}

func (v NullablePOSTCheckoutComPayments201ResponseDataAttributes) Get() *POSTCheckoutComPayments201ResponseDataAttributes {
	return v.value
}

func (v *NullablePOSTCheckoutComPayments201ResponseDataAttributes) Set(val *POSTCheckoutComPayments201ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTCheckoutComPayments201ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTCheckoutComPayments201ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTCheckoutComPayments201ResponseDataAttributes(val *POSTCheckoutComPayments201ResponseDataAttributes) *NullablePOSTCheckoutComPayments201ResponseDataAttributes {
	return &NullablePOSTCheckoutComPayments201ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTCheckoutComPayments201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTCheckoutComPayments201ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
