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

// PATCHCapturesCaptureId200ResponseDataAttributes struct for PATCHCapturesCaptureId200ResponseDataAttributes
type PATCHCapturesCaptureId200ResponseDataAttributes struct {
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Send this attribute if you want to create a refund for this capture.
	Refund *bool `json:"_refund,omitempty"`
	// The associated refund amount, in cents.
	RefundAmountCents *int32 `json:"_refund_amount_cents,omitempty"`
}

// NewPATCHCapturesCaptureId200ResponseDataAttributes instantiates a new PATCHCapturesCaptureId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHCapturesCaptureId200ResponseDataAttributes() *PATCHCapturesCaptureId200ResponseDataAttributes {
	this := PATCHCapturesCaptureId200ResponseDataAttributes{}
	return &this
}

// NewPATCHCapturesCaptureId200ResponseDataAttributesWithDefaults instantiates a new PATCHCapturesCaptureId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHCapturesCaptureId200ResponseDataAttributesWithDefaults() *PATCHCapturesCaptureId200ResponseDataAttributes {
	this := PATCHCapturesCaptureId200ResponseDataAttributes{}
	return &this
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *PATCHCapturesCaptureId200ResponseDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHCapturesCaptureId200ResponseDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHCapturesCaptureId200ResponseDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *PATCHCapturesCaptureId200ResponseDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *PATCHCapturesCaptureId200ResponseDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHCapturesCaptureId200ResponseDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHCapturesCaptureId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *PATCHCapturesCaptureId200ResponseDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PATCHCapturesCaptureId200ResponseDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHCapturesCaptureId200ResponseDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHCapturesCaptureId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PATCHCapturesCaptureId200ResponseDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetRefund returns the Refund field value if set, zero value otherwise.
func (o *PATCHCapturesCaptureId200ResponseDataAttributes) GetRefund() bool {
	if o == nil || o.Refund == nil {
		var ret bool
		return ret
	}
	return *o.Refund
}

// GetRefundOk returns a tuple with the Refund field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHCapturesCaptureId200ResponseDataAttributes) GetRefundOk() (*bool, bool) {
	if o == nil || o.Refund == nil {
		return nil, false
	}
	return o.Refund, true
}

// HasRefund returns a boolean if a field has been set.
func (o *PATCHCapturesCaptureId200ResponseDataAttributes) HasRefund() bool {
	if o != nil && o.Refund != nil {
		return true
	}

	return false
}

// SetRefund gets a reference to the given bool and assigns it to the Refund field.
func (o *PATCHCapturesCaptureId200ResponseDataAttributes) SetRefund(v bool) {
	o.Refund = &v
}

// GetRefundAmountCents returns the RefundAmountCents field value if set, zero value otherwise.
func (o *PATCHCapturesCaptureId200ResponseDataAttributes) GetRefundAmountCents() int32 {
	if o == nil || o.RefundAmountCents == nil {
		var ret int32
		return ret
	}
	return *o.RefundAmountCents
}

// GetRefundAmountCentsOk returns a tuple with the RefundAmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHCapturesCaptureId200ResponseDataAttributes) GetRefundAmountCentsOk() (*int32, bool) {
	if o == nil || o.RefundAmountCents == nil {
		return nil, false
	}
	return o.RefundAmountCents, true
}

// HasRefundAmountCents returns a boolean if a field has been set.
func (o *PATCHCapturesCaptureId200ResponseDataAttributes) HasRefundAmountCents() bool {
	if o != nil && o.RefundAmountCents != nil {
		return true
	}

	return false
}

// SetRefundAmountCents gets a reference to the given int32 and assigns it to the RefundAmountCents field.
func (o *PATCHCapturesCaptureId200ResponseDataAttributes) SetRefundAmountCents(v int32) {
	o.RefundAmountCents = &v
}

func (o PATCHCapturesCaptureId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
	}
	if o.ReferenceOrigin != nil {
		toSerialize["reference_origin"] = o.ReferenceOrigin
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Refund != nil {
		toSerialize["_refund"] = o.Refund
	}
	if o.RefundAmountCents != nil {
		toSerialize["_refund_amount_cents"] = o.RefundAmountCents
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHCapturesCaptureId200ResponseDataAttributes struct {
	value *PATCHCapturesCaptureId200ResponseDataAttributes
	isSet bool
}

func (v NullablePATCHCapturesCaptureId200ResponseDataAttributes) Get() *PATCHCapturesCaptureId200ResponseDataAttributes {
	return v.value
}

func (v *NullablePATCHCapturesCaptureId200ResponseDataAttributes) Set(val *PATCHCapturesCaptureId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHCapturesCaptureId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHCapturesCaptureId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHCapturesCaptureId200ResponseDataAttributes(val *PATCHCapturesCaptureId200ResponseDataAttributes) *NullablePATCHCapturesCaptureId200ResponseDataAttributes {
	return &NullablePATCHCapturesCaptureId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHCapturesCaptureId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHCapturesCaptureId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
