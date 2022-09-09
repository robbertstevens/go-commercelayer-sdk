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

// POSTBraintreePayments201ResponseDataAttributes struct for POSTBraintreePayments201ResponseDataAttributes
type POSTBraintreePayments201ResponseDataAttributes struct {
	// The Braintree payment ID used by local payment and sent by the Braintree JS SDK.
	PaymentId *string `json:"payment_id,omitempty"`
	// Indicates if the payment is local, in such case Braintree will trigger a webhook call passing the \"payment_id\" and \"payment_method_nonce\" in order to complete the transaction.
	Local *bool `json:"local,omitempty"`
	// Braintree payment options: 'customer_id' and 'payment_method_token'
	Options map[string]interface{} `json:"options,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewPOSTBraintreePayments201ResponseDataAttributes instantiates a new POSTBraintreePayments201ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTBraintreePayments201ResponseDataAttributes() *POSTBraintreePayments201ResponseDataAttributes {
	this := POSTBraintreePayments201ResponseDataAttributes{}
	return &this
}

// NewPOSTBraintreePayments201ResponseDataAttributesWithDefaults instantiates a new POSTBraintreePayments201ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTBraintreePayments201ResponseDataAttributesWithDefaults() *POSTBraintreePayments201ResponseDataAttributes {
	this := POSTBraintreePayments201ResponseDataAttributes{}
	return &this
}

// GetPaymentId returns the PaymentId field value if set, zero value otherwise.
func (o *POSTBraintreePayments201ResponseDataAttributes) GetPaymentId() string {
	if o == nil || o.PaymentId == nil {
		var ret string
		return ret
	}
	return *o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTBraintreePayments201ResponseDataAttributes) GetPaymentIdOk() (*string, bool) {
	if o == nil || o.PaymentId == nil {
		return nil, false
	}
	return o.PaymentId, true
}

// HasPaymentId returns a boolean if a field has been set.
func (o *POSTBraintreePayments201ResponseDataAttributes) HasPaymentId() bool {
	if o != nil && o.PaymentId != nil {
		return true
	}

	return false
}

// SetPaymentId gets a reference to the given string and assigns it to the PaymentId field.
func (o *POSTBraintreePayments201ResponseDataAttributes) SetPaymentId(v string) {
	o.PaymentId = &v
}

// GetLocal returns the Local field value if set, zero value otherwise.
func (o *POSTBraintreePayments201ResponseDataAttributes) GetLocal() bool {
	if o == nil || o.Local == nil {
		var ret bool
		return ret
	}
	return *o.Local
}

// GetLocalOk returns a tuple with the Local field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTBraintreePayments201ResponseDataAttributes) GetLocalOk() (*bool, bool) {
	if o == nil || o.Local == nil {
		return nil, false
	}
	return o.Local, true
}

// HasLocal returns a boolean if a field has been set.
func (o *POSTBraintreePayments201ResponseDataAttributes) HasLocal() bool {
	if o != nil && o.Local != nil {
		return true
	}

	return false
}

// SetLocal gets a reference to the given bool and assigns it to the Local field.
func (o *POSTBraintreePayments201ResponseDataAttributes) SetLocal(v bool) {
	o.Local = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *POSTBraintreePayments201ResponseDataAttributes) GetOptions() map[string]interface{} {
	if o == nil || o.Options == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTBraintreePayments201ResponseDataAttributes) GetOptionsOk() (map[string]interface{}, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *POSTBraintreePayments201ResponseDataAttributes) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given map[string]interface{} and assigns it to the Options field.
func (o *POSTBraintreePayments201ResponseDataAttributes) SetOptions(v map[string]interface{}) {
	o.Options = v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *POSTBraintreePayments201ResponseDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTBraintreePayments201ResponseDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTBraintreePayments201ResponseDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *POSTBraintreePayments201ResponseDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *POSTBraintreePayments201ResponseDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTBraintreePayments201ResponseDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTBraintreePayments201ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *POSTBraintreePayments201ResponseDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *POSTBraintreePayments201ResponseDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTBraintreePayments201ResponseDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTBraintreePayments201ResponseDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *POSTBraintreePayments201ResponseDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o POSTBraintreePayments201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PaymentId != nil {
		toSerialize["payment_id"] = o.PaymentId
	}
	if o.Local != nil {
		toSerialize["local"] = o.Local
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

type NullablePOSTBraintreePayments201ResponseDataAttributes struct {
	value *POSTBraintreePayments201ResponseDataAttributes
	isSet bool
}

func (v NullablePOSTBraintreePayments201ResponseDataAttributes) Get() *POSTBraintreePayments201ResponseDataAttributes {
	return v.value
}

func (v *NullablePOSTBraintreePayments201ResponseDataAttributes) Set(val *POSTBraintreePayments201ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTBraintreePayments201ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTBraintreePayments201ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTBraintreePayments201ResponseDataAttributes(val *POSTBraintreePayments201ResponseDataAttributes) *NullablePOSTBraintreePayments201ResponseDataAttributes {
	return &NullablePOSTBraintreePayments201ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTBraintreePayments201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTBraintreePayments201ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


