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

// checks if the GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes{}

// GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes struct for GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes
type GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes struct {
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
	// The gateway webhook endpoint ID, generated automatically.
	WebhookEndpointId interface{} `json:"webhook_endpoint_id,omitempty"`
	// The gateway webhook endpoint secret, generated automatically.
	WebhookEndpointSecret interface{} `json:"webhook_endpoint_secret,omitempty"`
	// The gateway webhook URL, generated automatically.
	WebhookEndpointUrl interface{} `json:"webhook_endpoint_url,omitempty"`
}

// NewGETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes instantiates a new GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes() *GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes {
	this := GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes{}
	return &this
}

// NewGETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributesWithDefaults instantiates a new GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributesWithDefaults() *GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes {
	this := GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) HasName() bool {
	if o != nil && IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given interface{} and assigns it to the Name field.
func (o *GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) GetCreatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return &o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) HasCreatedAt() bool {
	if o != nil && IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given interface{} and assigns it to the CreatedAt field.
func (o *GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) SetCreatedAt(v interface{}) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) GetUpdatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) HasUpdatedAt() bool {
	if o != nil && IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given interface{} and assigns it to the UpdatedAt field.
func (o *GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) SetUpdatedAt(v interface{}) {
	o.UpdatedAt = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

// GetWebhookEndpointId returns the WebhookEndpointId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) GetWebhookEndpointId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.WebhookEndpointId
}

// GetWebhookEndpointIdOk returns a tuple with the WebhookEndpointId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) GetWebhookEndpointIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.WebhookEndpointId) {
		return nil, false
	}
	return &o.WebhookEndpointId, true
}

// HasWebhookEndpointId returns a boolean if a field has been set.
func (o *GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) HasWebhookEndpointId() bool {
	if o != nil && IsNil(o.WebhookEndpointId) {
		return true
	}

	return false
}

// SetWebhookEndpointId gets a reference to the given interface{} and assigns it to the WebhookEndpointId field.
func (o *GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) SetWebhookEndpointId(v interface{}) {
	o.WebhookEndpointId = v
}

// GetWebhookEndpointSecret returns the WebhookEndpointSecret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) GetWebhookEndpointSecret() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.WebhookEndpointSecret
}

// GetWebhookEndpointSecretOk returns a tuple with the WebhookEndpointSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) GetWebhookEndpointSecretOk() (*interface{}, bool) {
	if o == nil || IsNil(o.WebhookEndpointSecret) {
		return nil, false
	}
	return &o.WebhookEndpointSecret, true
}

// HasWebhookEndpointSecret returns a boolean if a field has been set.
func (o *GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) HasWebhookEndpointSecret() bool {
	if o != nil && IsNil(o.WebhookEndpointSecret) {
		return true
	}

	return false
}

// SetWebhookEndpointSecret gets a reference to the given interface{} and assigns it to the WebhookEndpointSecret field.
func (o *GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) SetWebhookEndpointSecret(v interface{}) {
	o.WebhookEndpointSecret = v
}

// GetWebhookEndpointUrl returns the WebhookEndpointUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) GetWebhookEndpointUrl() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.WebhookEndpointUrl
}

// GetWebhookEndpointUrlOk returns a tuple with the WebhookEndpointUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) GetWebhookEndpointUrlOk() (*interface{}, bool) {
	if o == nil || IsNil(o.WebhookEndpointUrl) {
		return nil, false
	}
	return &o.WebhookEndpointUrl, true
}

// HasWebhookEndpointUrl returns a boolean if a field has been set.
func (o *GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) HasWebhookEndpointUrl() bool {
	if o != nil && IsNil(o.WebhookEndpointUrl) {
		return true
	}

	return false
}

// SetWebhookEndpointUrl gets a reference to the given interface{} and assigns it to the WebhookEndpointUrl field.
func (o *GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) SetWebhookEndpointUrl(v interface{}) {
	o.WebhookEndpointUrl = v
}

func (o GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
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
	if o.WebhookEndpointId != nil {
		toSerialize["webhook_endpoint_id"] = o.WebhookEndpointId
	}
	if o.WebhookEndpointSecret != nil {
		toSerialize["webhook_endpoint_secret"] = o.WebhookEndpointSecret
	}
	if o.WebhookEndpointUrl != nil {
		toSerialize["webhook_endpoint_url"] = o.WebhookEndpointUrl
	}
	return toSerialize, nil
}

type NullableGETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes struct {
	value *GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes
	isSet bool
}

func (v NullableGETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) Get() *GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes {
	return v.value
}

func (v *NullableGETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) Set(val *GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes(val *GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) *NullableGETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes {
	return &NullableGETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullableGETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
