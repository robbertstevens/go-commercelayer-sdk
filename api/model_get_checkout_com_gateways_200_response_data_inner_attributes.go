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

// GETCheckoutComGateways200ResponseDataInnerAttributes struct for GETCheckoutComGateways200ResponseDataInnerAttributes
type GETCheckoutComGateways200ResponseDataInnerAttributes struct {
	// The payment gateway's internal name.
	Name *string `json:"name,omitempty"`
	// Unique identifier for the resource (hash).
	Id *string `json:"id,omitempty"`
	// Time at which the resource was created.
	CreatedAt *string `json:"created_at,omitempty"`
	// Time at which the resource was last updated.
	UpdatedAt *string `json:"updated_at,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// The gateway webhook endpoint ID, generated automatically.
	WebhookEndpointId *string `json:"webhook_endpoint_id,omitempty"`
	// The gateway webhook endpoint secret, generated automatically.
	WebhookEndpointSecret *string `json:"webhook_endpoint_secret,omitempty"`
	// The gateway webhook URL, generated automatically.
	WebhookEndpointUrl *string `json:"webhook_endpoint_url,omitempty"`
}

// NewGETCheckoutComGateways200ResponseDataInnerAttributes instantiates a new GETCheckoutComGateways200ResponseDataInnerAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCheckoutComGateways200ResponseDataInnerAttributes() *GETCheckoutComGateways200ResponseDataInnerAttributes {
	this := GETCheckoutComGateways200ResponseDataInnerAttributes{}
	return &this
}

// NewGETCheckoutComGateways200ResponseDataInnerAttributesWithDefaults instantiates a new GETCheckoutComGateways200ResponseDataInnerAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCheckoutComGateways200ResponseDataInnerAttributesWithDefaults() *GETCheckoutComGateways200ResponseDataInnerAttributes {
	this := GETCheckoutComGateways200ResponseDataInnerAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GETCheckoutComGateways200ResponseDataInnerAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCheckoutComGateways200ResponseDataInnerAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GETCheckoutComGateways200ResponseDataInnerAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GETCheckoutComGateways200ResponseDataInnerAttributes) SetName(v string) {
	o.Name = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETCheckoutComGateways200ResponseDataInnerAttributes) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCheckoutComGateways200ResponseDataInnerAttributes) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETCheckoutComGateways200ResponseDataInnerAttributes) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETCheckoutComGateways200ResponseDataInnerAttributes) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GETCheckoutComGateways200ResponseDataInnerAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCheckoutComGateways200ResponseDataInnerAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETCheckoutComGateways200ResponseDataInnerAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *GETCheckoutComGateways200ResponseDataInnerAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GETCheckoutComGateways200ResponseDataInnerAttributes) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCheckoutComGateways200ResponseDataInnerAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETCheckoutComGateways200ResponseDataInnerAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *GETCheckoutComGateways200ResponseDataInnerAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *GETCheckoutComGateways200ResponseDataInnerAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCheckoutComGateways200ResponseDataInnerAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETCheckoutComGateways200ResponseDataInnerAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *GETCheckoutComGateways200ResponseDataInnerAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *GETCheckoutComGateways200ResponseDataInnerAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCheckoutComGateways200ResponseDataInnerAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETCheckoutComGateways200ResponseDataInnerAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *GETCheckoutComGateways200ResponseDataInnerAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GETCheckoutComGateways200ResponseDataInnerAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCheckoutComGateways200ResponseDataInnerAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETCheckoutComGateways200ResponseDataInnerAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *GETCheckoutComGateways200ResponseDataInnerAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetWebhookEndpointId returns the WebhookEndpointId field value if set, zero value otherwise.
func (o *GETCheckoutComGateways200ResponseDataInnerAttributes) GetWebhookEndpointId() string {
	if o == nil || o.WebhookEndpointId == nil {
		var ret string
		return ret
	}
	return *o.WebhookEndpointId
}

// GetWebhookEndpointIdOk returns a tuple with the WebhookEndpointId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCheckoutComGateways200ResponseDataInnerAttributes) GetWebhookEndpointIdOk() (*string, bool) {
	if o == nil || o.WebhookEndpointId == nil {
		return nil, false
	}
	return o.WebhookEndpointId, true
}

// HasWebhookEndpointId returns a boolean if a field has been set.
func (o *GETCheckoutComGateways200ResponseDataInnerAttributes) HasWebhookEndpointId() bool {
	if o != nil && o.WebhookEndpointId != nil {
		return true
	}

	return false
}

// SetWebhookEndpointId gets a reference to the given string and assigns it to the WebhookEndpointId field.
func (o *GETCheckoutComGateways200ResponseDataInnerAttributes) SetWebhookEndpointId(v string) {
	o.WebhookEndpointId = &v
}

// GetWebhookEndpointSecret returns the WebhookEndpointSecret field value if set, zero value otherwise.
func (o *GETCheckoutComGateways200ResponseDataInnerAttributes) GetWebhookEndpointSecret() string {
	if o == nil || o.WebhookEndpointSecret == nil {
		var ret string
		return ret
	}
	return *o.WebhookEndpointSecret
}

// GetWebhookEndpointSecretOk returns a tuple with the WebhookEndpointSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCheckoutComGateways200ResponseDataInnerAttributes) GetWebhookEndpointSecretOk() (*string, bool) {
	if o == nil || o.WebhookEndpointSecret == nil {
		return nil, false
	}
	return o.WebhookEndpointSecret, true
}

// HasWebhookEndpointSecret returns a boolean if a field has been set.
func (o *GETCheckoutComGateways200ResponseDataInnerAttributes) HasWebhookEndpointSecret() bool {
	if o != nil && o.WebhookEndpointSecret != nil {
		return true
	}

	return false
}

// SetWebhookEndpointSecret gets a reference to the given string and assigns it to the WebhookEndpointSecret field.
func (o *GETCheckoutComGateways200ResponseDataInnerAttributes) SetWebhookEndpointSecret(v string) {
	o.WebhookEndpointSecret = &v
}

// GetWebhookEndpointUrl returns the WebhookEndpointUrl field value if set, zero value otherwise.
func (o *GETCheckoutComGateways200ResponseDataInnerAttributes) GetWebhookEndpointUrl() string {
	if o == nil || o.WebhookEndpointUrl == nil {
		var ret string
		return ret
	}
	return *o.WebhookEndpointUrl
}

// GetWebhookEndpointUrlOk returns a tuple with the WebhookEndpointUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCheckoutComGateways200ResponseDataInnerAttributes) GetWebhookEndpointUrlOk() (*string, bool) {
	if o == nil || o.WebhookEndpointUrl == nil {
		return nil, false
	}
	return o.WebhookEndpointUrl, true
}

// HasWebhookEndpointUrl returns a boolean if a field has been set.
func (o *GETCheckoutComGateways200ResponseDataInnerAttributes) HasWebhookEndpointUrl() bool {
	if o != nil && o.WebhookEndpointUrl != nil {
		return true
	}

	return false
}

// SetWebhookEndpointUrl gets a reference to the given string and assigns it to the WebhookEndpointUrl field.
func (o *GETCheckoutComGateways200ResponseDataInnerAttributes) SetWebhookEndpointUrl(v string) {
	o.WebhookEndpointUrl = &v
}

func (o GETCheckoutComGateways200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
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
	return json.Marshal(toSerialize)
}

type NullableGETCheckoutComGateways200ResponseDataInnerAttributes struct {
	value *GETCheckoutComGateways200ResponseDataInnerAttributes
	isSet bool
}

func (v NullableGETCheckoutComGateways200ResponseDataInnerAttributes) Get() *GETCheckoutComGateways200ResponseDataInnerAttributes {
	return v.value
}

func (v *NullableGETCheckoutComGateways200ResponseDataInnerAttributes) Set(val *GETCheckoutComGateways200ResponseDataInnerAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCheckoutComGateways200ResponseDataInnerAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCheckoutComGateways200ResponseDataInnerAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCheckoutComGateways200ResponseDataInnerAttributes(val *GETCheckoutComGateways200ResponseDataInnerAttributes) *NullableGETCheckoutComGateways200ResponseDataInnerAttributes {
	return &NullableGETCheckoutComGateways200ResponseDataInnerAttributes{value: val, isSet: true}
}

func (v NullableGETCheckoutComGateways200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCheckoutComGateways200ResponseDataInnerAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


