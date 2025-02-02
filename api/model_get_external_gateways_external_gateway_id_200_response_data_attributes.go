/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.3.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GETExternalGatewaysExternalGatewayId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETExternalGatewaysExternalGatewayId200ResponseDataAttributes{}

// GETExternalGatewaysExternalGatewayId200ResponseDataAttributes struct for GETExternalGatewaysExternalGatewayId200ResponseDataAttributes
type GETExternalGatewaysExternalGatewayId200ResponseDataAttributes struct {
	// The payment gateway's internal name.
	Name interface{} `json:"name,omitempty"`
	// Time at which the resource was created.
	CreatedAt interface{} `json:"created_at,omitempty"`
	// Time at which the resource was last updated.
	UpdatedAt interface{} `json:"updated_at,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
	// The endpoint used by the external gateway to authorize payments.
	AuthorizeUrl interface{} `json:"authorize_url,omitempty"`
	// The endpoint used by the external gateway to capture payments.
	CaptureUrl interface{} `json:"capture_url,omitempty"`
	// The endpoint used by the external gateway to void payments.
	VoidUrl interface{} `json:"void_url,omitempty"`
	// The endpoint used by the external gateway to refund payments.
	RefundUrl interface{} `json:"refund_url,omitempty"`
	// The endpoint used by the external gateway to create a customer payment token.
	TokenUrl interface{} `json:"token_url,omitempty"`
	// The circuit breaker state, by default it is 'closed'. It can become 'open' once the number of consecutive failures overlaps the specified threshold, in such case no further calls to the failing callback are made.
	CircuitState interface{} `json:"circuit_state,omitempty"`
	// The number of consecutive failures recorded by the circuit breaker associated to this resource, will be reset on first successful call to callback.
	CircuitFailureCount interface{} `json:"circuit_failure_count,omitempty"`
	// The shared secret used to sign the external request payload.
	SharedSecret interface{} `json:"shared_secret,omitempty"`
}

// NewGETExternalGatewaysExternalGatewayId200ResponseDataAttributes instantiates a new GETExternalGatewaysExternalGatewayId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETExternalGatewaysExternalGatewayId200ResponseDataAttributes() *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes {
	this := GETExternalGatewaysExternalGatewayId200ResponseDataAttributes{}
	return &this
}

// NewGETExternalGatewaysExternalGatewayId200ResponseDataAttributesWithDefaults instantiates a new GETExternalGatewaysExternalGatewayId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETExternalGatewaysExternalGatewayId200ResponseDataAttributesWithDefaults() *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes {
	this := GETExternalGatewaysExternalGatewayId200ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) HasName() bool {
	if o != nil && IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given interface{} and assigns it to the Name field.
func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetCreatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return &o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) HasCreatedAt() bool {
	if o != nil && IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given interface{} and assigns it to the CreatedAt field.
func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetCreatedAt(v interface{}) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetUpdatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) HasUpdatedAt() bool {
	if o != nil && IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given interface{} and assigns it to the UpdatedAt field.
func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetUpdatedAt(v interface{}) {
	o.UpdatedAt = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

// GetAuthorizeUrl returns the AuthorizeUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetAuthorizeUrl() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.AuthorizeUrl
}

// GetAuthorizeUrlOk returns a tuple with the AuthorizeUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetAuthorizeUrlOk() (*interface{}, bool) {
	if o == nil || IsNil(o.AuthorizeUrl) {
		return nil, false
	}
	return &o.AuthorizeUrl, true
}

// HasAuthorizeUrl returns a boolean if a field has been set.
func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) HasAuthorizeUrl() bool {
	if o != nil && IsNil(o.AuthorizeUrl) {
		return true
	}

	return false
}

// SetAuthorizeUrl gets a reference to the given interface{} and assigns it to the AuthorizeUrl field.
func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetAuthorizeUrl(v interface{}) {
	o.AuthorizeUrl = v
}

// GetCaptureUrl returns the CaptureUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetCaptureUrl() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CaptureUrl
}

// GetCaptureUrlOk returns a tuple with the CaptureUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetCaptureUrlOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CaptureUrl) {
		return nil, false
	}
	return &o.CaptureUrl, true
}

// HasCaptureUrl returns a boolean if a field has been set.
func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) HasCaptureUrl() bool {
	if o != nil && IsNil(o.CaptureUrl) {
		return true
	}

	return false
}

// SetCaptureUrl gets a reference to the given interface{} and assigns it to the CaptureUrl field.
func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetCaptureUrl(v interface{}) {
	o.CaptureUrl = v
}

// GetVoidUrl returns the VoidUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetVoidUrl() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.VoidUrl
}

// GetVoidUrlOk returns a tuple with the VoidUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetVoidUrlOk() (*interface{}, bool) {
	if o == nil || IsNil(o.VoidUrl) {
		return nil, false
	}
	return &o.VoidUrl, true
}

// HasVoidUrl returns a boolean if a field has been set.
func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) HasVoidUrl() bool {
	if o != nil && IsNil(o.VoidUrl) {
		return true
	}

	return false
}

// SetVoidUrl gets a reference to the given interface{} and assigns it to the VoidUrl field.
func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetVoidUrl(v interface{}) {
	o.VoidUrl = v
}

// GetRefundUrl returns the RefundUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetRefundUrl() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.RefundUrl
}

// GetRefundUrlOk returns a tuple with the RefundUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetRefundUrlOk() (*interface{}, bool) {
	if o == nil || IsNil(o.RefundUrl) {
		return nil, false
	}
	return &o.RefundUrl, true
}

// HasRefundUrl returns a boolean if a field has been set.
func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) HasRefundUrl() bool {
	if o != nil && IsNil(o.RefundUrl) {
		return true
	}

	return false
}

// SetRefundUrl gets a reference to the given interface{} and assigns it to the RefundUrl field.
func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetRefundUrl(v interface{}) {
	o.RefundUrl = v
}

// GetTokenUrl returns the TokenUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetTokenUrl() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.TokenUrl
}

// GetTokenUrlOk returns a tuple with the TokenUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetTokenUrlOk() (*interface{}, bool) {
	if o == nil || IsNil(o.TokenUrl) {
		return nil, false
	}
	return &o.TokenUrl, true
}

// HasTokenUrl returns a boolean if a field has been set.
func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) HasTokenUrl() bool {
	if o != nil && IsNil(o.TokenUrl) {
		return true
	}

	return false
}

// SetTokenUrl gets a reference to the given interface{} and assigns it to the TokenUrl field.
func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetTokenUrl(v interface{}) {
	o.TokenUrl = v
}

// GetCircuitState returns the CircuitState field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetCircuitState() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CircuitState
}

// GetCircuitStateOk returns a tuple with the CircuitState field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetCircuitStateOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CircuitState) {
		return nil, false
	}
	return &o.CircuitState, true
}

// HasCircuitState returns a boolean if a field has been set.
func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) HasCircuitState() bool {
	if o != nil && IsNil(o.CircuitState) {
		return true
	}

	return false
}

// SetCircuitState gets a reference to the given interface{} and assigns it to the CircuitState field.
func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetCircuitState(v interface{}) {
	o.CircuitState = v
}

// GetCircuitFailureCount returns the CircuitFailureCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetCircuitFailureCount() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CircuitFailureCount
}

// GetCircuitFailureCountOk returns a tuple with the CircuitFailureCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetCircuitFailureCountOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CircuitFailureCount) {
		return nil, false
	}
	return &o.CircuitFailureCount, true
}

// HasCircuitFailureCount returns a boolean if a field has been set.
func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) HasCircuitFailureCount() bool {
	if o != nil && IsNil(o.CircuitFailureCount) {
		return true
	}

	return false
}

// SetCircuitFailureCount gets a reference to the given interface{} and assigns it to the CircuitFailureCount field.
func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetCircuitFailureCount(v interface{}) {
	o.CircuitFailureCount = v
}

// GetSharedSecret returns the SharedSecret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetSharedSecret() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.SharedSecret
}

// GetSharedSecretOk returns a tuple with the SharedSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetSharedSecretOk() (*interface{}, bool) {
	if o == nil || IsNil(o.SharedSecret) {
		return nil, false
	}
	return &o.SharedSecret, true
}

// HasSharedSecret returns a boolean if a field has been set.
func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) HasSharedSecret() bool {
	if o != nil && IsNil(o.SharedSecret) {
		return true
	}

	return false
}

// SetSharedSecret gets a reference to the given interface{} and assigns it to the SharedSecret field.
func (o *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetSharedSecret(v interface{}) {
	o.SharedSecret = v
}

func (o GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
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
	if o.CircuitState != nil {
		toSerialize["circuit_state"] = o.CircuitState
	}
	if o.CircuitFailureCount != nil {
		toSerialize["circuit_failure_count"] = o.CircuitFailureCount
	}
	if o.SharedSecret != nil {
		toSerialize["shared_secret"] = o.SharedSecret
	}
	return toSerialize, nil
}

type NullableGETExternalGatewaysExternalGatewayId200ResponseDataAttributes struct {
	value *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes
	isSet bool
}

func (v NullableGETExternalGatewaysExternalGatewayId200ResponseDataAttributes) Get() *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes {
	return v.value
}

func (v *NullableGETExternalGatewaysExternalGatewayId200ResponseDataAttributes) Set(val *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETExternalGatewaysExternalGatewayId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETExternalGatewaysExternalGatewayId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETExternalGatewaysExternalGatewayId200ResponseDataAttributes(val *GETExternalGatewaysExternalGatewayId200ResponseDataAttributes) *NullableGETExternalGatewaysExternalGatewayId200ResponseDataAttributes {
	return &NullableGETExternalGatewaysExternalGatewayId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullableGETExternalGatewaysExternalGatewayId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETExternalGatewaysExternalGatewayId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
