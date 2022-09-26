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

// POSTStripeGateways201ResponseDataAttributes struct for POSTStripeGateways201ResponseDataAttributes
type POSTStripeGateways201ResponseDataAttributes struct {
	// The payment gateway's internal name.
	Name string `json:"name"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// The gateway login.
	Login string `json:"login"`
	// The gateway publishable API key.
	PublishableKey *string `json:"publishable_key,omitempty"`
}

// NewPOSTStripeGateways201ResponseDataAttributes instantiates a new POSTStripeGateways201ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTStripeGateways201ResponseDataAttributes(name string, login string) *POSTStripeGateways201ResponseDataAttributes {
	this := POSTStripeGateways201ResponseDataAttributes{}
	this.Name = name
	this.Login = login
	return &this
}

// NewPOSTStripeGateways201ResponseDataAttributesWithDefaults instantiates a new POSTStripeGateways201ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTStripeGateways201ResponseDataAttributesWithDefaults() *POSTStripeGateways201ResponseDataAttributes {
	this := POSTStripeGateways201ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value
func (o *POSTStripeGateways201ResponseDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *POSTStripeGateways201ResponseDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *POSTStripeGateways201ResponseDataAttributes) SetName(v string) {
	o.Name = v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *POSTStripeGateways201ResponseDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTStripeGateways201ResponseDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTStripeGateways201ResponseDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *POSTStripeGateways201ResponseDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *POSTStripeGateways201ResponseDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTStripeGateways201ResponseDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTStripeGateways201ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *POSTStripeGateways201ResponseDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *POSTStripeGateways201ResponseDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTStripeGateways201ResponseDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTStripeGateways201ResponseDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *POSTStripeGateways201ResponseDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetLogin returns the Login field value
func (o *POSTStripeGateways201ResponseDataAttributes) GetLogin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Login
}

// GetLoginOk returns a tuple with the Login field value
// and a boolean to check if the value has been set.
func (o *POSTStripeGateways201ResponseDataAttributes) GetLoginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Login, true
}

// SetLogin sets field value
func (o *POSTStripeGateways201ResponseDataAttributes) SetLogin(v string) {
	o.Login = v
}

// GetPublishableKey returns the PublishableKey field value if set, zero value otherwise.
func (o *POSTStripeGateways201ResponseDataAttributes) GetPublishableKey() string {
	if o == nil || o.PublishableKey == nil {
		var ret string
		return ret
	}
	return *o.PublishableKey
}

// GetPublishableKeyOk returns a tuple with the PublishableKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTStripeGateways201ResponseDataAttributes) GetPublishableKeyOk() (*string, bool) {
	if o == nil || o.PublishableKey == nil {
		return nil, false
	}
	return o.PublishableKey, true
}

// HasPublishableKey returns a boolean if a field has been set.
func (o *POSTStripeGateways201ResponseDataAttributes) HasPublishableKey() bool {
	if o != nil && o.PublishableKey != nil {
		return true
	}

	return false
}

// SetPublishableKey gets a reference to the given string and assigns it to the PublishableKey field.
func (o *POSTStripeGateways201ResponseDataAttributes) SetPublishableKey(v string) {
	o.PublishableKey = &v
}

func (o POSTStripeGateways201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["login"] = o.Login
	}
	if o.PublishableKey != nil {
		toSerialize["publishable_key"] = o.PublishableKey
	}
	return json.Marshal(toSerialize)
}

type NullablePOSTStripeGateways201ResponseDataAttributes struct {
	value *POSTStripeGateways201ResponseDataAttributes
	isSet bool
}

func (v NullablePOSTStripeGateways201ResponseDataAttributes) Get() *POSTStripeGateways201ResponseDataAttributes {
	return v.value
}

func (v *NullablePOSTStripeGateways201ResponseDataAttributes) Set(val *POSTStripeGateways201ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTStripeGateways201ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTStripeGateways201ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTStripeGateways201ResponseDataAttributes(val *POSTStripeGateways201ResponseDataAttributes) *NullablePOSTStripeGateways201ResponseDataAttributes {
	return &NullablePOSTStripeGateways201ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTStripeGateways201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTStripeGateways201ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
