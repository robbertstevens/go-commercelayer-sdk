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

// POSTExternalGateways201ResponseDataAttributes struct for POSTExternalGateways201ResponseDataAttributes
type POSTExternalGateways201ResponseDataAttributes struct {
	// The payment gateway's internal name.
	Name string `json:"name"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// The endpoint used by the external gateway to authorize payments.
	AuthorizeUrl *string `json:"authorize_url,omitempty"`
	// The endpoint used by the external gateway to capture payments.
	CaptureUrl *string `json:"capture_url,omitempty"`
	// The endpoint used by the external gateway to void payments.
	VoidUrl *string `json:"void_url,omitempty"`
	// The endpoint used by the external gateway to refund payments.
	RefundUrl *string `json:"refund_url,omitempty"`
	// The endpoint used by the external gateway to create a customer payment token.
	TokenUrl *string `json:"token_url,omitempty"`
}

// NewPOSTExternalGateways201ResponseDataAttributes instantiates a new POSTExternalGateways201ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTExternalGateways201ResponseDataAttributes(name string) *POSTExternalGateways201ResponseDataAttributes {
	this := POSTExternalGateways201ResponseDataAttributes{}
	this.Name = name
	return &this
}

// NewPOSTExternalGateways201ResponseDataAttributesWithDefaults instantiates a new POSTExternalGateways201ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTExternalGateways201ResponseDataAttributesWithDefaults() *POSTExternalGateways201ResponseDataAttributes {
	this := POSTExternalGateways201ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value
func (o *POSTExternalGateways201ResponseDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *POSTExternalGateways201ResponseDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *POSTExternalGateways201ResponseDataAttributes) SetName(v string) {
	o.Name = v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *POSTExternalGateways201ResponseDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTExternalGateways201ResponseDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTExternalGateways201ResponseDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *POSTExternalGateways201ResponseDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *POSTExternalGateways201ResponseDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTExternalGateways201ResponseDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTExternalGateways201ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *POSTExternalGateways201ResponseDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *POSTExternalGateways201ResponseDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTExternalGateways201ResponseDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTExternalGateways201ResponseDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *POSTExternalGateways201ResponseDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetAuthorizeUrl returns the AuthorizeUrl field value if set, zero value otherwise.
func (o *POSTExternalGateways201ResponseDataAttributes) GetAuthorizeUrl() string {
	if o == nil || o.AuthorizeUrl == nil {
		var ret string
		return ret
	}
	return *o.AuthorizeUrl
}

// GetAuthorizeUrlOk returns a tuple with the AuthorizeUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTExternalGateways201ResponseDataAttributes) GetAuthorizeUrlOk() (*string, bool) {
	if o == nil || o.AuthorizeUrl == nil {
		return nil, false
	}
	return o.AuthorizeUrl, true
}

// HasAuthorizeUrl returns a boolean if a field has been set.
func (o *POSTExternalGateways201ResponseDataAttributes) HasAuthorizeUrl() bool {
	if o != nil && o.AuthorizeUrl != nil {
		return true
	}

	return false
}

// SetAuthorizeUrl gets a reference to the given string and assigns it to the AuthorizeUrl field.
func (o *POSTExternalGateways201ResponseDataAttributes) SetAuthorizeUrl(v string) {
	o.AuthorizeUrl = &v
}

// GetCaptureUrl returns the CaptureUrl field value if set, zero value otherwise.
func (o *POSTExternalGateways201ResponseDataAttributes) GetCaptureUrl() string {
	if o == nil || o.CaptureUrl == nil {
		var ret string
		return ret
	}
	return *o.CaptureUrl
}

// GetCaptureUrlOk returns a tuple with the CaptureUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTExternalGateways201ResponseDataAttributes) GetCaptureUrlOk() (*string, bool) {
	if o == nil || o.CaptureUrl == nil {
		return nil, false
	}
	return o.CaptureUrl, true
}

// HasCaptureUrl returns a boolean if a field has been set.
func (o *POSTExternalGateways201ResponseDataAttributes) HasCaptureUrl() bool {
	if o != nil && o.CaptureUrl != nil {
		return true
	}

	return false
}

// SetCaptureUrl gets a reference to the given string and assigns it to the CaptureUrl field.
func (o *POSTExternalGateways201ResponseDataAttributes) SetCaptureUrl(v string) {
	o.CaptureUrl = &v
}

// GetVoidUrl returns the VoidUrl field value if set, zero value otherwise.
func (o *POSTExternalGateways201ResponseDataAttributes) GetVoidUrl() string {
	if o == nil || o.VoidUrl == nil {
		var ret string
		return ret
	}
	return *o.VoidUrl
}

// GetVoidUrlOk returns a tuple with the VoidUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTExternalGateways201ResponseDataAttributes) GetVoidUrlOk() (*string, bool) {
	if o == nil || o.VoidUrl == nil {
		return nil, false
	}
	return o.VoidUrl, true
}

// HasVoidUrl returns a boolean if a field has been set.
func (o *POSTExternalGateways201ResponseDataAttributes) HasVoidUrl() bool {
	if o != nil && o.VoidUrl != nil {
		return true
	}

	return false
}

// SetVoidUrl gets a reference to the given string and assigns it to the VoidUrl field.
func (o *POSTExternalGateways201ResponseDataAttributes) SetVoidUrl(v string) {
	o.VoidUrl = &v
}

// GetRefundUrl returns the RefundUrl field value if set, zero value otherwise.
func (o *POSTExternalGateways201ResponseDataAttributes) GetRefundUrl() string {
	if o == nil || o.RefundUrl == nil {
		var ret string
		return ret
	}
	return *o.RefundUrl
}

// GetRefundUrlOk returns a tuple with the RefundUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTExternalGateways201ResponseDataAttributes) GetRefundUrlOk() (*string, bool) {
	if o == nil || o.RefundUrl == nil {
		return nil, false
	}
	return o.RefundUrl, true
}

// HasRefundUrl returns a boolean if a field has been set.
func (o *POSTExternalGateways201ResponseDataAttributes) HasRefundUrl() bool {
	if o != nil && o.RefundUrl != nil {
		return true
	}

	return false
}

// SetRefundUrl gets a reference to the given string and assigns it to the RefundUrl field.
func (o *POSTExternalGateways201ResponseDataAttributes) SetRefundUrl(v string) {
	o.RefundUrl = &v
}

// GetTokenUrl returns the TokenUrl field value if set, zero value otherwise.
func (o *POSTExternalGateways201ResponseDataAttributes) GetTokenUrl() string {
	if o == nil || o.TokenUrl == nil {
		var ret string
		return ret
	}
	return *o.TokenUrl
}

// GetTokenUrlOk returns a tuple with the TokenUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTExternalGateways201ResponseDataAttributes) GetTokenUrlOk() (*string, bool) {
	if o == nil || o.TokenUrl == nil {
		return nil, false
	}
	return o.TokenUrl, true
}

// HasTokenUrl returns a boolean if a field has been set.
func (o *POSTExternalGateways201ResponseDataAttributes) HasTokenUrl() bool {
	if o != nil && o.TokenUrl != nil {
		return true
	}

	return false
}

// SetTokenUrl gets a reference to the given string and assigns it to the TokenUrl field.
func (o *POSTExternalGateways201ResponseDataAttributes) SetTokenUrl(v string) {
	o.TokenUrl = &v
}

func (o POSTExternalGateways201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
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
	if o.AuthorizeUrl != nil {
		toSerialize["authorize_url"] = o.AuthorizeUrl
	}
	if o.CaptureUrl != nil {
		toSerialize["capture_url"] = o.CaptureUrl
	}
	if o.VoidUrl != nil {
		toSerialize["void_url"] = o.VoidUrl
	}
	if o.RefundUrl != nil {
		toSerialize["refund_url"] = o.RefundUrl
	}
	if o.TokenUrl != nil {
		toSerialize["token_url"] = o.TokenUrl
	}
	return json.Marshal(toSerialize)
}

type NullablePOSTExternalGateways201ResponseDataAttributes struct {
	value *POSTExternalGateways201ResponseDataAttributes
	isSet bool
}

func (v NullablePOSTExternalGateways201ResponseDataAttributes) Get() *POSTExternalGateways201ResponseDataAttributes {
	return v.value
}

func (v *NullablePOSTExternalGateways201ResponseDataAttributes) Set(val *POSTExternalGateways201ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTExternalGateways201ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTExternalGateways201ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTExternalGateways201ResponseDataAttributes(val *POSTExternalGateways201ResponseDataAttributes) *NullablePOSTExternalGateways201ResponseDataAttributes {
	return &NullablePOSTExternalGateways201ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTExternalGateways201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTExternalGateways201ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
