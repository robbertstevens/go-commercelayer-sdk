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

// PATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes struct for PATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes
type PATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes struct {
	// The expiration date/time of this subscription (must be after starts_at).
	ExpiresAt *string `json:"expires_at,omitempty"`
	// Send this attribute if you want to mark this subscription as active.
	Activate *bool `json:"_activate,omitempty"`
	// Send this attribute if you want to mark this subscription as inactive.
	Deactivate *bool `json:"_deactivate,omitempty"`
	// Send this attribute if you want to mark this subscription as cancelled.
	Cancel *bool `json:"_cancel,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewPATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes instantiates a new PATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes() *PATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes {
	this := PATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes{}
	return &this
}

// NewPATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributesWithDefaults instantiates a new PATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributesWithDefaults() *PATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes {
	this := PATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes{}
	return &this
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *PATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetExpiresAt() string {
	if o == nil || o.ExpiresAt == nil {
		var ret string
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetExpiresAtOk() (*string, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *PATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt != nil {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given string and assigns it to the ExpiresAt field.
func (o *PATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetExpiresAt(v string) {
	o.ExpiresAt = &v
}

// GetActivate returns the Activate field value if set, zero value otherwise.
func (o *PATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetActivate() bool {
	if o == nil || o.Activate == nil {
		var ret bool
		return ret
	}
	return *o.Activate
}

// GetActivateOk returns a tuple with the Activate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetActivateOk() (*bool, bool) {
	if o == nil || o.Activate == nil {
		return nil, false
	}
	return o.Activate, true
}

// HasActivate returns a boolean if a field has been set.
func (o *PATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) HasActivate() bool {
	if o != nil && o.Activate != nil {
		return true
	}

	return false
}

// SetActivate gets a reference to the given bool and assigns it to the Activate field.
func (o *PATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetActivate(v bool) {
	o.Activate = &v
}

// GetDeactivate returns the Deactivate field value if set, zero value otherwise.
func (o *PATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetDeactivate() bool {
	if o == nil || o.Deactivate == nil {
		var ret bool
		return ret
	}
	return *o.Deactivate
}

// GetDeactivateOk returns a tuple with the Deactivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetDeactivateOk() (*bool, bool) {
	if o == nil || o.Deactivate == nil {
		return nil, false
	}
	return o.Deactivate, true
}

// HasDeactivate returns a boolean if a field has been set.
func (o *PATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) HasDeactivate() bool {
	if o != nil && o.Deactivate != nil {
		return true
	}

	return false
}

// SetDeactivate gets a reference to the given bool and assigns it to the Deactivate field.
func (o *PATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetDeactivate(v bool) {
	o.Deactivate = &v
}

// GetCancel returns the Cancel field value if set, zero value otherwise.
func (o *PATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetCancel() bool {
	if o == nil || o.Cancel == nil {
		var ret bool
		return ret
	}
	return *o.Cancel
}

// GetCancelOk returns a tuple with the Cancel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetCancelOk() (*bool, bool) {
	if o == nil || o.Cancel == nil {
		return nil, false
	}
	return o.Cancel, true
}

// HasCancel returns a boolean if a field has been set.
func (o *PATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) HasCancel() bool {
	if o != nil && o.Cancel != nil {
		return true
	}

	return false
}

// SetCancel gets a reference to the given bool and assigns it to the Cancel field.
func (o *PATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetCancel(v bool) {
	o.Cancel = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *PATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *PATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *PATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *PATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o PATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExpiresAt != nil {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if o.Activate != nil {
		toSerialize["_activate"] = o.Activate
	}
	if o.Deactivate != nil {
		toSerialize["_deactivate"] = o.Deactivate
	}
	if o.Cancel != nil {
		toSerialize["_cancel"] = o.Cancel
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

type NullablePATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes struct {
	value *PATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes
	isSet bool
}

func (v NullablePATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) Get() *PATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes {
	return v.value
}

func (v *NullablePATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) Set(val *PATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes(val *PATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) *NullablePATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes {
	return &NullablePATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
