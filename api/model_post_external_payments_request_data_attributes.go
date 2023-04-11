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

// checks if the POSTExternalPaymentsRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTExternalPaymentsRequestDataAttributes{}

// POSTExternalPaymentsRequestDataAttributes struct for POSTExternalPaymentsRequestDataAttributes
type POSTExternalPaymentsRequestDataAttributes struct {
	// The payment source token, as generated by the external gateway SDK. Credit Card numbers are rejected.
	PaymentSourceToken interface{} `json:"payment_source_token"`
	// External payment options.
	Options interface{} `json:"options,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPOSTExternalPaymentsRequestDataAttributes instantiates a new POSTExternalPaymentsRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTExternalPaymentsRequestDataAttributes(paymentSourceToken interface{}) *POSTExternalPaymentsRequestDataAttributes {
	this := POSTExternalPaymentsRequestDataAttributes{}
	this.PaymentSourceToken = paymentSourceToken
	return &this
}

// NewPOSTExternalPaymentsRequestDataAttributesWithDefaults instantiates a new POSTExternalPaymentsRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTExternalPaymentsRequestDataAttributesWithDefaults() *POSTExternalPaymentsRequestDataAttributes {
	this := POSTExternalPaymentsRequestDataAttributes{}
	return &this
}

// GetPaymentSourceToken returns the PaymentSourceToken field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTExternalPaymentsRequestDataAttributes) GetPaymentSourceToken() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.PaymentSourceToken
}

// GetPaymentSourceTokenOk returns a tuple with the PaymentSourceToken field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTExternalPaymentsRequestDataAttributes) GetPaymentSourceTokenOk() (*interface{}, bool) {
	if o == nil || IsNil(o.PaymentSourceToken) {
		return nil, false
	}
	return &o.PaymentSourceToken, true
}

// SetPaymentSourceToken sets field value
func (o *POSTExternalPaymentsRequestDataAttributes) SetPaymentSourceToken(v interface{}) {
	o.PaymentSourceToken = v
}

// GetOptions returns the Options field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTExternalPaymentsRequestDataAttributes) GetOptions() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTExternalPaymentsRequestDataAttributes) GetOptionsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return &o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *POSTExternalPaymentsRequestDataAttributes) HasOptions() bool {
	if o != nil && IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given interface{} and assigns it to the Options field.
func (o *POSTExternalPaymentsRequestDataAttributes) SetOptions(v interface{}) {
	o.Options = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTExternalPaymentsRequestDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTExternalPaymentsRequestDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTExternalPaymentsRequestDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *POSTExternalPaymentsRequestDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTExternalPaymentsRequestDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTExternalPaymentsRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTExternalPaymentsRequestDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *POSTExternalPaymentsRequestDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTExternalPaymentsRequestDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTExternalPaymentsRequestDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTExternalPaymentsRequestDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *POSTExternalPaymentsRequestDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o POSTExternalPaymentsRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTExternalPaymentsRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.PaymentSourceToken != nil {
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
	return toSerialize, nil
}

type NullablePOSTExternalPaymentsRequestDataAttributes struct {
	value *POSTExternalPaymentsRequestDataAttributes
	isSet bool
}

func (v NullablePOSTExternalPaymentsRequestDataAttributes) Get() *POSTExternalPaymentsRequestDataAttributes {
	return v.value
}

func (v *NullablePOSTExternalPaymentsRequestDataAttributes) Set(val *POSTExternalPaymentsRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTExternalPaymentsRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTExternalPaymentsRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTExternalPaymentsRequestDataAttributes(val *POSTExternalPaymentsRequestDataAttributes) *NullablePOSTExternalPaymentsRequestDataAttributes {
	return &NullablePOSTExternalPaymentsRequestDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTExternalPaymentsRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTExternalPaymentsRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
