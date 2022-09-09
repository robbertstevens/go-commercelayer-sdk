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

// POSTCustomerPasswordResets201ResponseDataAttributes struct for POSTCustomerPasswordResets201ResponseDataAttributes
type POSTCustomerPasswordResets201ResponseDataAttributes struct {
	// The email of the customer that requires a password reset
	CustomerEmail string `json:"customer_email"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewPOSTCustomerPasswordResets201ResponseDataAttributes instantiates a new POSTCustomerPasswordResets201ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTCustomerPasswordResets201ResponseDataAttributes(customerEmail string) *POSTCustomerPasswordResets201ResponseDataAttributes {
	this := POSTCustomerPasswordResets201ResponseDataAttributes{}
	this.CustomerEmail = customerEmail
	return &this
}

// NewPOSTCustomerPasswordResets201ResponseDataAttributesWithDefaults instantiates a new POSTCustomerPasswordResets201ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTCustomerPasswordResets201ResponseDataAttributesWithDefaults() *POSTCustomerPasswordResets201ResponseDataAttributes {
	this := POSTCustomerPasswordResets201ResponseDataAttributes{}
	return &this
}

// GetCustomerEmail returns the CustomerEmail field value
func (o *POSTCustomerPasswordResets201ResponseDataAttributes) GetCustomerEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerEmail
}

// GetCustomerEmailOk returns a tuple with the CustomerEmail field value
// and a boolean to check if the value has been set.
func (o *POSTCustomerPasswordResets201ResponseDataAttributes) GetCustomerEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerEmail, true
}

// SetCustomerEmail sets field value
func (o *POSTCustomerPasswordResets201ResponseDataAttributes) SetCustomerEmail(v string) {
	o.CustomerEmail = v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *POSTCustomerPasswordResets201ResponseDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCustomerPasswordResets201ResponseDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTCustomerPasswordResets201ResponseDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *POSTCustomerPasswordResets201ResponseDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *POSTCustomerPasswordResets201ResponseDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCustomerPasswordResets201ResponseDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTCustomerPasswordResets201ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *POSTCustomerPasswordResets201ResponseDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *POSTCustomerPasswordResets201ResponseDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCustomerPasswordResets201ResponseDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTCustomerPasswordResets201ResponseDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *POSTCustomerPasswordResets201ResponseDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o POSTCustomerPasswordResets201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["customer_email"] = o.CustomerEmail
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

type NullablePOSTCustomerPasswordResets201ResponseDataAttributes struct {
	value *POSTCustomerPasswordResets201ResponseDataAttributes
	isSet bool
}

func (v NullablePOSTCustomerPasswordResets201ResponseDataAttributes) Get() *POSTCustomerPasswordResets201ResponseDataAttributes {
	return v.value
}

func (v *NullablePOSTCustomerPasswordResets201ResponseDataAttributes) Set(val *POSTCustomerPasswordResets201ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTCustomerPasswordResets201ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTCustomerPasswordResets201ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTCustomerPasswordResets201ResponseDataAttributes(val *POSTCustomerPasswordResets201ResponseDataAttributes) *NullablePOSTCustomerPasswordResets201ResponseDataAttributes {
	return &NullablePOSTCustomerPasswordResets201ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTCustomerPasswordResets201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTCustomerPasswordResets201ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


