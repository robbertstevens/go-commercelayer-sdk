/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.3.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the PATCHAxervePaymentsAxervePaymentId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHAxervePaymentsAxervePaymentId200ResponseDataAttributes{}

// PATCHAxervePaymentsAxervePaymentId200ResponseDataAttributes struct for PATCHAxervePaymentsAxervePaymentId200ResponseDataAttributes
type PATCHAxervePaymentsAxervePaymentId200ResponseDataAttributes struct {
	// The Axerve payment request data, collected by client.
	PaymentRequestData interface{} `json:"payment_request_data,omitempty"`
	// Send this attribute if you want to update the payment with fresh order data.
	Update interface{} `json:"_update,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPATCHAxervePaymentsAxervePaymentId200ResponseDataAttributes instantiates a new PATCHAxervePaymentsAxervePaymentId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHAxervePaymentsAxervePaymentId200ResponseDataAttributes() *PATCHAxervePaymentsAxervePaymentId200ResponseDataAttributes {
	this := PATCHAxervePaymentsAxervePaymentId200ResponseDataAttributes{}
	return &this
}

// NewPATCHAxervePaymentsAxervePaymentId200ResponseDataAttributesWithDefaults instantiates a new PATCHAxervePaymentsAxervePaymentId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHAxervePaymentsAxervePaymentId200ResponseDataAttributesWithDefaults() *PATCHAxervePaymentsAxervePaymentId200ResponseDataAttributes {
	this := PATCHAxervePaymentsAxervePaymentId200ResponseDataAttributes{}
	return &this
}

// GetPaymentRequestData returns the PaymentRequestData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHAxervePaymentsAxervePaymentId200ResponseDataAttributes) GetPaymentRequestData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.PaymentRequestData
}

// GetPaymentRequestDataOk returns a tuple with the PaymentRequestData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHAxervePaymentsAxervePaymentId200ResponseDataAttributes) GetPaymentRequestDataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.PaymentRequestData) {
		return nil, false
	}
	return &o.PaymentRequestData, true
}

// HasPaymentRequestData returns a boolean if a field has been set.
func (o *PATCHAxervePaymentsAxervePaymentId200ResponseDataAttributes) HasPaymentRequestData() bool {
	if o != nil && IsNil(o.PaymentRequestData) {
		return true
	}

	return false
}

// SetPaymentRequestData gets a reference to the given interface{} and assigns it to the PaymentRequestData field.
func (o *PATCHAxervePaymentsAxervePaymentId200ResponseDataAttributes) SetPaymentRequestData(v interface{}) {
	o.PaymentRequestData = v
}

// GetUpdate returns the Update field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHAxervePaymentsAxervePaymentId200ResponseDataAttributes) GetUpdate() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHAxervePaymentsAxervePaymentId200ResponseDataAttributes) GetUpdateOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return &o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *PATCHAxervePaymentsAxervePaymentId200ResponseDataAttributes) HasUpdate() bool {
	if o != nil && IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given interface{} and assigns it to the Update field.
func (o *PATCHAxervePaymentsAxervePaymentId200ResponseDataAttributes) SetUpdate(v interface{}) {
	o.Update = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHAxervePaymentsAxervePaymentId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHAxervePaymentsAxervePaymentId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHAxervePaymentsAxervePaymentId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *PATCHAxervePaymentsAxervePaymentId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHAxervePaymentsAxervePaymentId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHAxervePaymentsAxervePaymentId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHAxervePaymentsAxervePaymentId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *PATCHAxervePaymentsAxervePaymentId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHAxervePaymentsAxervePaymentId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHAxervePaymentsAxervePaymentId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHAxervePaymentsAxervePaymentId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *PATCHAxervePaymentsAxervePaymentId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o PATCHAxervePaymentsAxervePaymentId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHAxervePaymentsAxervePaymentId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.PaymentRequestData != nil {
		toSerialize["payment_request_data"] = o.PaymentRequestData
	}
	if o.Update != nil {
		toSerialize["_update"] = o.Update
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

type NullablePATCHAxervePaymentsAxervePaymentId200ResponseDataAttributes struct {
	value *PATCHAxervePaymentsAxervePaymentId200ResponseDataAttributes
	isSet bool
}

func (v NullablePATCHAxervePaymentsAxervePaymentId200ResponseDataAttributes) Get() *PATCHAxervePaymentsAxervePaymentId200ResponseDataAttributes {
	return v.value
}

func (v *NullablePATCHAxervePaymentsAxervePaymentId200ResponseDataAttributes) Set(val *PATCHAxervePaymentsAxervePaymentId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHAxervePaymentsAxervePaymentId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHAxervePaymentsAxervePaymentId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHAxervePaymentsAxervePaymentId200ResponseDataAttributes(val *PATCHAxervePaymentsAxervePaymentId200ResponseDataAttributes) *NullablePATCHAxervePaymentsAxervePaymentId200ResponseDataAttributes {
	return &NullablePATCHAxervePaymentsAxervePaymentId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHAxervePaymentsAxervePaymentId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHAxervePaymentsAxervePaymentId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
