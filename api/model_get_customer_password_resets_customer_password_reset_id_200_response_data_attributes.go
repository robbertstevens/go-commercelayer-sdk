/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 4.1.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes{}

// GETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes struct for GETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes
type GETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes struct {
	// The email of the customer that requires a password reset
	CustomerEmail interface{} `json:"customer_email,omitempty"`
	// Automatically generated on create. Send its value as the '_reset_password_token' argument when updating the customer password.
	ResetPasswordToken interface{} `json:"reset_password_token,omitempty"`
	// Time at which the password was reset.
	ResetPasswordAt interface{} `json:"reset_password_at,omitempty"`
	// Time at which the resource was created.
	CreatedAt interface{} `json:"created_at,omitempty"`
	// Time at which the resource was last updated.
	UpdatedAt interface{} `json:"updated_at,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewGETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes instantiates a new GETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes() *GETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes {
	this := GETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes{}
	return &this
}

// NewGETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributesWithDefaults instantiates a new GETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributesWithDefaults() *GETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes {
	this := GETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes{}
	return &this
}

// GetCustomerEmail returns the CustomerEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) GetCustomerEmail() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CustomerEmail
}

// GetCustomerEmailOk returns a tuple with the CustomerEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) GetCustomerEmailOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CustomerEmail) {
		return nil, false
	}
	return &o.CustomerEmail, true
}

// HasCustomerEmail returns a boolean if a field has been set.
func (o *GETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) HasCustomerEmail() bool {
	if o != nil && IsNil(o.CustomerEmail) {
		return true
	}

	return false
}

// SetCustomerEmail gets a reference to the given interface{} and assigns it to the CustomerEmail field.
func (o *GETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) SetCustomerEmail(v interface{}) {
	o.CustomerEmail = v
}

// GetResetPasswordToken returns the ResetPasswordToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) GetResetPasswordToken() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ResetPasswordToken
}

// GetResetPasswordTokenOk returns a tuple with the ResetPasswordToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) GetResetPasswordTokenOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ResetPasswordToken) {
		return nil, false
	}
	return &o.ResetPasswordToken, true
}

// HasResetPasswordToken returns a boolean if a field has been set.
func (o *GETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) HasResetPasswordToken() bool {
	if o != nil && IsNil(o.ResetPasswordToken) {
		return true
	}

	return false
}

// SetResetPasswordToken gets a reference to the given interface{} and assigns it to the ResetPasswordToken field.
func (o *GETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) SetResetPasswordToken(v interface{}) {
	o.ResetPasswordToken = v
}

// GetResetPasswordAt returns the ResetPasswordAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) GetResetPasswordAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ResetPasswordAt
}

// GetResetPasswordAtOk returns a tuple with the ResetPasswordAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) GetResetPasswordAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ResetPasswordAt) {
		return nil, false
	}
	return &o.ResetPasswordAt, true
}

// HasResetPasswordAt returns a boolean if a field has been set.
func (o *GETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) HasResetPasswordAt() bool {
	if o != nil && IsNil(o.ResetPasswordAt) {
		return true
	}

	return false
}

// SetResetPasswordAt gets a reference to the given interface{} and assigns it to the ResetPasswordAt field.
func (o *GETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) SetResetPasswordAt(v interface{}) {
	o.ResetPasswordAt = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) GetCreatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return &o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) HasCreatedAt() bool {
	if o != nil && IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given interface{} and assigns it to the CreatedAt field.
func (o *GETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) SetCreatedAt(v interface{}) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) GetUpdatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) HasUpdatedAt() bool {
	if o != nil && IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given interface{} and assigns it to the UpdatedAt field.
func (o *GETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) SetUpdatedAt(v interface{}) {
	o.UpdatedAt = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *GETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *GETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *GETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o GETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CustomerEmail != nil {
		toSerialize["customer_email"] = o.CustomerEmail
	}
	if o.ResetPasswordToken != nil {
		toSerialize["reset_password_token"] = o.ResetPasswordToken
	}
	if o.ResetPasswordAt != nil {
		toSerialize["reset_password_at"] = o.ResetPasswordAt
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
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

type NullableGETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes struct {
	value *GETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes
	isSet bool
}

func (v NullableGETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) Get() *GETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes {
	return v.value
}

func (v *NullableGETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) Set(val *GETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes(val *GETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) *NullableGETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes {
	return &NullableGETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullableGETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
