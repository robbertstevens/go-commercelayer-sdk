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

// PATCHCustomersCustomerId200ResponseDataAttributes struct for PATCHCustomersCustomerId200ResponseDataAttributes
type PATCHCustomersCustomerId200ResponseDataAttributes struct {
	// The customer's email address
	Email *string `json:"email,omitempty"`
	// The customer's password. Initiate a customer password reset flow if you need to change it.
	Password *string `json:"password,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewPATCHCustomersCustomerId200ResponseDataAttributes instantiates a new PATCHCustomersCustomerId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHCustomersCustomerId200ResponseDataAttributes() *PATCHCustomersCustomerId200ResponseDataAttributes {
	this := PATCHCustomersCustomerId200ResponseDataAttributes{}
	return &this
}

// NewPATCHCustomersCustomerId200ResponseDataAttributesWithDefaults instantiates a new PATCHCustomersCustomerId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHCustomersCustomerId200ResponseDataAttributesWithDefaults() *PATCHCustomersCustomerId200ResponseDataAttributes {
	this := PATCHCustomersCustomerId200ResponseDataAttributes{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *PATCHCustomersCustomerId200ResponseDataAttributes) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHCustomersCustomerId200ResponseDataAttributes) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *PATCHCustomersCustomerId200ResponseDataAttributes) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *PATCHCustomersCustomerId200ResponseDataAttributes) SetEmail(v string) {
	o.Email = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *PATCHCustomersCustomerId200ResponseDataAttributes) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHCustomersCustomerId200ResponseDataAttributes) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *PATCHCustomersCustomerId200ResponseDataAttributes) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *PATCHCustomersCustomerId200ResponseDataAttributes) SetPassword(v string) {
	o.Password = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *PATCHCustomersCustomerId200ResponseDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHCustomersCustomerId200ResponseDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHCustomersCustomerId200ResponseDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *PATCHCustomersCustomerId200ResponseDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *PATCHCustomersCustomerId200ResponseDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHCustomersCustomerId200ResponseDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHCustomersCustomerId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *PATCHCustomersCustomerId200ResponseDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PATCHCustomersCustomerId200ResponseDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHCustomersCustomerId200ResponseDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHCustomersCustomerId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PATCHCustomersCustomerId200ResponseDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o PATCHCustomersCustomerId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
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

type NullablePATCHCustomersCustomerId200ResponseDataAttributes struct {
	value *PATCHCustomersCustomerId200ResponseDataAttributes
	isSet bool
}

func (v NullablePATCHCustomersCustomerId200ResponseDataAttributes) Get() *PATCHCustomersCustomerId200ResponseDataAttributes {
	return v.value
}

func (v *NullablePATCHCustomersCustomerId200ResponseDataAttributes) Set(val *PATCHCustomersCustomerId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHCustomersCustomerId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHCustomersCustomerId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHCustomersCustomerId200ResponseDataAttributes(val *PATCHCustomersCustomerId200ResponseDataAttributes) *NullablePATCHCustomersCustomerId200ResponseDataAttributes {
	return &NullablePATCHCustomersCustomerId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHCustomersCustomerId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHCustomersCustomerId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


