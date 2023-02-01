/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.4.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PATCHAuthorizationsAuthorizationId200ResponseDataAttributes struct for PATCHAuthorizationsAuthorizationId200ResponseDataAttributes
type PATCHAuthorizationsAuthorizationId200ResponseDataAttributes struct {
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Send this attribute if you want to create a capture for this authorization.
	Capture *bool `json:"_capture,omitempty"`
	// The associated capture amount, in cents.
	CaptureAmountCents *int32 `json:"_capture_amount_cents,omitempty"`
	// Send this attribute if you want to create a void for this authorization.
	Void *bool `json:"_void,omitempty"`
}

// NewPATCHAuthorizationsAuthorizationId200ResponseDataAttributes instantiates a new PATCHAuthorizationsAuthorizationId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHAuthorizationsAuthorizationId200ResponseDataAttributes() *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes {
	this := PATCHAuthorizationsAuthorizationId200ResponseDataAttributes{}
	return &this
}

// NewPATCHAuthorizationsAuthorizationId200ResponseDataAttributesWithDefaults instantiates a new PATCHAuthorizationsAuthorizationId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHAuthorizationsAuthorizationId200ResponseDataAttributesWithDefaults() *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes {
	this := PATCHAuthorizationsAuthorizationId200ResponseDataAttributes{}
	return &this
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetCapture returns the Capture field value if set, zero value otherwise.
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) GetCapture() bool {
	if o == nil || o.Capture == nil {
		var ret bool
		return ret
	}
	return *o.Capture
}

// GetCaptureOk returns a tuple with the Capture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) GetCaptureOk() (*bool, bool) {
	if o == nil || o.Capture == nil {
		return nil, false
	}
	return o.Capture, true
}

// HasCapture returns a boolean if a field has been set.
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) HasCapture() bool {
	if o != nil && o.Capture != nil {
		return true
	}

	return false
}

// SetCapture gets a reference to the given bool and assigns it to the Capture field.
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) SetCapture(v bool) {
	o.Capture = &v
}

// GetCaptureAmountCents returns the CaptureAmountCents field value if set, zero value otherwise.
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) GetCaptureAmountCents() int32 {
	if o == nil || o.CaptureAmountCents == nil {
		var ret int32
		return ret
	}
	return *o.CaptureAmountCents
}

// GetCaptureAmountCentsOk returns a tuple with the CaptureAmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) GetCaptureAmountCentsOk() (*int32, bool) {
	if o == nil || o.CaptureAmountCents == nil {
		return nil, false
	}
	return o.CaptureAmountCents, true
}

// HasCaptureAmountCents returns a boolean if a field has been set.
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) HasCaptureAmountCents() bool {
	if o != nil && o.CaptureAmountCents != nil {
		return true
	}

	return false
}

// SetCaptureAmountCents gets a reference to the given int32 and assigns it to the CaptureAmountCents field.
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) SetCaptureAmountCents(v int32) {
	o.CaptureAmountCents = &v
}

// GetVoid returns the Void field value if set, zero value otherwise.
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) GetVoid() bool {
	if o == nil || o.Void == nil {
		var ret bool
		return ret
	}
	return *o.Void
}

// GetVoidOk returns a tuple with the Void field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) GetVoidOk() (*bool, bool) {
	if o == nil || o.Void == nil {
		return nil, false
	}
	return o.Void, true
}

// HasVoid returns a boolean if a field has been set.
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) HasVoid() bool {
	if o != nil && o.Void != nil {
		return true
	}

	return false
}

// SetVoid gets a reference to the given bool and assigns it to the Void field.
func (o *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) SetVoid(v bool) {
	o.Void = &v
}

func (o PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
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
	if o.Capture != nil {
		toSerialize["_capture"] = o.Capture
	}
	if o.CaptureAmountCents != nil {
		toSerialize["_capture_amount_cents"] = o.CaptureAmountCents
	}
	if o.Void != nil {
		toSerialize["_void"] = o.Void
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHAuthorizationsAuthorizationId200ResponseDataAttributes struct {
	value *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes
	isSet bool
}

func (v NullablePATCHAuthorizationsAuthorizationId200ResponseDataAttributes) Get() *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes {
	return v.value
}

func (v *NullablePATCHAuthorizationsAuthorizationId200ResponseDataAttributes) Set(val *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHAuthorizationsAuthorizationId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHAuthorizationsAuthorizationId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHAuthorizationsAuthorizationId200ResponseDataAttributes(val *PATCHAuthorizationsAuthorizationId200ResponseDataAttributes) *NullablePATCHAuthorizationsAuthorizationId200ResponseDataAttributes {
	return &NullablePATCHAuthorizationsAuthorizationId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHAuthorizationsAuthorizationId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHAuthorizationsAuthorizationId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
