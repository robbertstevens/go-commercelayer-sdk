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

// checks if the PATCHStripePaymentsStripePaymentId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHStripePaymentsStripePaymentId200ResponseDataAttributes{}

// PATCHStripePaymentsStripePaymentId200ResponseDataAttributes struct for PATCHStripePaymentsStripePaymentId200ResponseDataAttributes
type PATCHStripePaymentsStripePaymentId200ResponseDataAttributes struct {
	// Stripe payment options: 'customer', 'payment_method', 'return_url', etc. Check Stripe payment intent API for more details.
	Options interface{} `json:"options,omitempty"`
	// The URL to return to when a redirect payment is completed.
	ReturnUrl interface{} `json:"return_url,omitempty"`
	// Send this attribute if you want to update the created payment intent with fresh order data.
	Update interface{} `json:"_update,omitempty"`
	// Send this attribute if you want to refresh the payment status, can be used as webhooks fallback logic.
	Refresh interface{} `json:"_refresh,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPATCHStripePaymentsStripePaymentId200ResponseDataAttributes instantiates a new PATCHStripePaymentsStripePaymentId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHStripePaymentsStripePaymentId200ResponseDataAttributes() *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes {
	this := PATCHStripePaymentsStripePaymentId200ResponseDataAttributes{}
	return &this
}

// NewPATCHStripePaymentsStripePaymentId200ResponseDataAttributesWithDefaults instantiates a new PATCHStripePaymentsStripePaymentId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHStripePaymentsStripePaymentId200ResponseDataAttributesWithDefaults() *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes {
	this := PATCHStripePaymentsStripePaymentId200ResponseDataAttributes{}
	return &this
}

// GetOptions returns the Options field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) GetOptions() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) GetOptionsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return &o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) HasOptions() bool {
	if o != nil && IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given interface{} and assigns it to the Options field.
func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) SetOptions(v interface{}) {
	o.Options = v
}

// GetReturnUrl returns the ReturnUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) GetReturnUrl() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReturnUrl
}

// GetReturnUrlOk returns a tuple with the ReturnUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) GetReturnUrlOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReturnUrl) {
		return nil, false
	}
	return &o.ReturnUrl, true
}

// HasReturnUrl returns a boolean if a field has been set.
func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) HasReturnUrl() bool {
	if o != nil && IsNil(o.ReturnUrl) {
		return true
	}

	return false
}

// SetReturnUrl gets a reference to the given interface{} and assigns it to the ReturnUrl field.
func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) SetReturnUrl(v interface{}) {
	o.ReturnUrl = v
}

// GetUpdate returns the Update field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) GetUpdate() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) GetUpdateOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return &o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) HasUpdate() bool {
	if o != nil && IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given interface{} and assigns it to the Update field.
func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) SetUpdate(v interface{}) {
	o.Update = v
}

// GetRefresh returns the Refresh field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) GetRefresh() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Refresh
}

// GetRefreshOk returns a tuple with the Refresh field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) GetRefreshOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Refresh) {
		return nil, false
	}
	return &o.Refresh, true
}

// HasRefresh returns a boolean if a field has been set.
func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) HasRefresh() bool {
	if o != nil && IsNil(o.Refresh) {
		return true
	}

	return false
}

// SetRefresh gets a reference to the given interface{} and assigns it to the Refresh field.
func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) SetRefresh(v interface{}) {
	o.Refresh = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.ReturnUrl != nil {
		toSerialize["return_url"] = o.ReturnUrl
	}
	if o.Update != nil {
		toSerialize["_update"] = o.Update
	}
	if o.Refresh != nil {
		toSerialize["_refresh"] = o.Refresh
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

type NullablePATCHStripePaymentsStripePaymentId200ResponseDataAttributes struct {
	value *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes
	isSet bool
}

func (v NullablePATCHStripePaymentsStripePaymentId200ResponseDataAttributes) Get() *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes {
	return v.value
}

func (v *NullablePATCHStripePaymentsStripePaymentId200ResponseDataAttributes) Set(val *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHStripePaymentsStripePaymentId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHStripePaymentsStripePaymentId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHStripePaymentsStripePaymentId200ResponseDataAttributes(val *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) *NullablePATCHStripePaymentsStripePaymentId200ResponseDataAttributes {
	return &NullablePATCHStripePaymentsStripePaymentId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHStripePaymentsStripePaymentId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHStripePaymentsStripePaymentId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
