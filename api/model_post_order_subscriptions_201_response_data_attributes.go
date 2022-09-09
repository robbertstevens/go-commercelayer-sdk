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

// POSTOrderSubscriptions201ResponseDataAttributes struct for POSTOrderSubscriptions201ResponseDataAttributes
type POSTOrderSubscriptions201ResponseDataAttributes struct {
	// The frequency of the subscription. One of 'hourly', 'daily', 'weekly', 'monthly', 'two-month', 'three-month', 'four-month', 'six-month', or 'yearly'.
	Frequency string `json:"frequency"`
	// Indicates if the subscription will be activated considering the placed source order as its first run, default to true.
	ActivateBySourceOrder *bool `json:"activate_by_source_order,omitempty"`
	// The activation date/time of this subscription.
	StartsAt *string `json:"starts_at,omitempty"`
	// The expiration date/time of this subscription (must be after starts_at).
	ExpiresAt *string `json:"expires_at,omitempty"`
	// The subscription options used to create the order copy (check order_copies for more information). For subscriptions the `place_target_order` is enabled by default, specify custom options to overwrite it.
	Options map[string]interface{} `json:"options,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewPOSTOrderSubscriptions201ResponseDataAttributes instantiates a new POSTOrderSubscriptions201ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTOrderSubscriptions201ResponseDataAttributes(frequency string) *POSTOrderSubscriptions201ResponseDataAttributes {
	this := POSTOrderSubscriptions201ResponseDataAttributes{}
	this.Frequency = frequency
	return &this
}

// NewPOSTOrderSubscriptions201ResponseDataAttributesWithDefaults instantiates a new POSTOrderSubscriptions201ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTOrderSubscriptions201ResponseDataAttributesWithDefaults() *POSTOrderSubscriptions201ResponseDataAttributes {
	this := POSTOrderSubscriptions201ResponseDataAttributes{}
	return &this
}

// GetFrequency returns the Frequency field value
func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetFrequency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value
// and a boolean to check if the value has been set.
func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetFrequencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Frequency, true
}

// SetFrequency sets field value
func (o *POSTOrderSubscriptions201ResponseDataAttributes) SetFrequency(v string) {
	o.Frequency = v
}

// GetActivateBySourceOrder returns the ActivateBySourceOrder field value if set, zero value otherwise.
func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetActivateBySourceOrder() bool {
	if o == nil || o.ActivateBySourceOrder == nil {
		var ret bool
		return ret
	}
	return *o.ActivateBySourceOrder
}

// GetActivateBySourceOrderOk returns a tuple with the ActivateBySourceOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetActivateBySourceOrderOk() (*bool, bool) {
	if o == nil || o.ActivateBySourceOrder == nil {
		return nil, false
	}
	return o.ActivateBySourceOrder, true
}

// HasActivateBySourceOrder returns a boolean if a field has been set.
func (o *POSTOrderSubscriptions201ResponseDataAttributes) HasActivateBySourceOrder() bool {
	if o != nil && o.ActivateBySourceOrder != nil {
		return true
	}

	return false
}

// SetActivateBySourceOrder gets a reference to the given bool and assigns it to the ActivateBySourceOrder field.
func (o *POSTOrderSubscriptions201ResponseDataAttributes) SetActivateBySourceOrder(v bool) {
	o.ActivateBySourceOrder = &v
}

// GetStartsAt returns the StartsAt field value if set, zero value otherwise.
func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetStartsAt() string {
	if o == nil || o.StartsAt == nil {
		var ret string
		return ret
	}
	return *o.StartsAt
}

// GetStartsAtOk returns a tuple with the StartsAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetStartsAtOk() (*string, bool) {
	if o == nil || o.StartsAt == nil {
		return nil, false
	}
	return o.StartsAt, true
}

// HasStartsAt returns a boolean if a field has been set.
func (o *POSTOrderSubscriptions201ResponseDataAttributes) HasStartsAt() bool {
	if o != nil && o.StartsAt != nil {
		return true
	}

	return false
}

// SetStartsAt gets a reference to the given string and assigns it to the StartsAt field.
func (o *POSTOrderSubscriptions201ResponseDataAttributes) SetStartsAt(v string) {
	o.StartsAt = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetExpiresAt() string {
	if o == nil || o.ExpiresAt == nil {
		var ret string
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetExpiresAtOk() (*string, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *POSTOrderSubscriptions201ResponseDataAttributes) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt != nil {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given string and assigns it to the ExpiresAt field.
func (o *POSTOrderSubscriptions201ResponseDataAttributes) SetExpiresAt(v string) {
	o.ExpiresAt = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetOptions() map[string]interface{} {
	if o == nil || o.Options == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetOptionsOk() (map[string]interface{}, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *POSTOrderSubscriptions201ResponseDataAttributes) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given map[string]interface{} and assigns it to the Options field.
func (o *POSTOrderSubscriptions201ResponseDataAttributes) SetOptions(v map[string]interface{}) {
	o.Options = v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTOrderSubscriptions201ResponseDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *POSTOrderSubscriptions201ResponseDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTOrderSubscriptions201ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *POSTOrderSubscriptions201ResponseDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTOrderSubscriptions201ResponseDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *POSTOrderSubscriptions201ResponseDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o POSTOrderSubscriptions201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["frequency"] = o.Frequency
	}
	if o.ActivateBySourceOrder != nil {
		toSerialize["activate_by_source_order"] = o.ActivateBySourceOrder
	}
	if o.StartsAt != nil {
		toSerialize["starts_at"] = o.StartsAt
	}
	if o.ExpiresAt != nil {
		toSerialize["expires_at"] = o.ExpiresAt
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

type NullablePOSTOrderSubscriptions201ResponseDataAttributes struct {
	value *POSTOrderSubscriptions201ResponseDataAttributes
	isSet bool
}

func (v NullablePOSTOrderSubscriptions201ResponseDataAttributes) Get() *POSTOrderSubscriptions201ResponseDataAttributes {
	return v.value
}

func (v *NullablePOSTOrderSubscriptions201ResponseDataAttributes) Set(val *POSTOrderSubscriptions201ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTOrderSubscriptions201ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTOrderSubscriptions201ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTOrderSubscriptions201ResponseDataAttributes(val *POSTOrderSubscriptions201ResponseDataAttributes) *NullablePOSTOrderSubscriptions201ResponseDataAttributes {
	return &NullablePOSTOrderSubscriptions201ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTOrderSubscriptions201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTOrderSubscriptions201ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


