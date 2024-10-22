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

// checks if the GETMarketsMarketId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETMarketsMarketId200ResponseDataAttributes{}

// GETMarketsMarketId200ResponseDataAttributes struct for GETMarketsMarketId200ResponseDataAttributes
type GETMarketsMarketId200ResponseDataAttributes struct {
	// Unique identifier for the market (numeric).
	Number interface{} `json:"number,omitempty"`
	// The market's internal name.
	Name interface{} `json:"name,omitempty"`
	// A string that you can use to identify the market (must be unique within the environment).
	Code interface{} `json:"code,omitempty"`
	// The Facebook Pixed ID.
	FacebookPixelId interface{} `json:"facebook_pixel_id,omitempty"`
	// The checkout URL for this market.
	CheckoutUrl interface{} `json:"checkout_url,omitempty"`
	// The URL used to overwrite prices by an external source.
	ExternalPricesUrl interface{} `json:"external_prices_url,omitempty"`
	// The URL used to validate orders by an external source.
	ExternalOrderValidationUrl interface{} `json:"external_order_validation_url,omitempty"`
	// Indicates if market belongs to a customer_group.
	Private interface{} `json:"private,omitempty"`
	// When specified indicates the maximum number of shipping line items with cost that will be added to an order.
	ShippingCostCutoff interface{} `json:"shipping_cost_cutoff,omitempty"`
	// Time at which this resource was disabled.
	DisabledAt interface{} `json:"disabled_at,omitempty"`
	// The shared secret used to sign the external request payload.
	SharedSecret interface{} `json:"shared_secret,omitempty"`
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
}

// NewGETMarketsMarketId200ResponseDataAttributes instantiates a new GETMarketsMarketId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETMarketsMarketId200ResponseDataAttributes() *GETMarketsMarketId200ResponseDataAttributes {
	this := GETMarketsMarketId200ResponseDataAttributes{}
	return &this
}

// NewGETMarketsMarketId200ResponseDataAttributesWithDefaults instantiates a new GETMarketsMarketId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETMarketsMarketId200ResponseDataAttributesWithDefaults() *GETMarketsMarketId200ResponseDataAttributes {
	this := GETMarketsMarketId200ResponseDataAttributes{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETMarketsMarketId200ResponseDataAttributes) GetNumber() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETMarketsMarketId200ResponseDataAttributes) GetNumberOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return &o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *GETMarketsMarketId200ResponseDataAttributes) HasNumber() bool {
	if o != nil && IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given interface{} and assigns it to the Number field.
func (o *GETMarketsMarketId200ResponseDataAttributes) SetNumber(v interface{}) {
	o.Number = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETMarketsMarketId200ResponseDataAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETMarketsMarketId200ResponseDataAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GETMarketsMarketId200ResponseDataAttributes) HasName() bool {
	if o != nil && IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given interface{} and assigns it to the Name field.
func (o *GETMarketsMarketId200ResponseDataAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETMarketsMarketId200ResponseDataAttributes) GetCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETMarketsMarketId200ResponseDataAttributes) GetCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return &o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *GETMarketsMarketId200ResponseDataAttributes) HasCode() bool {
	if o != nil && IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given interface{} and assigns it to the Code field.
func (o *GETMarketsMarketId200ResponseDataAttributes) SetCode(v interface{}) {
	o.Code = v
}

// GetFacebookPixelId returns the FacebookPixelId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETMarketsMarketId200ResponseDataAttributes) GetFacebookPixelId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.FacebookPixelId
}

// GetFacebookPixelIdOk returns a tuple with the FacebookPixelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETMarketsMarketId200ResponseDataAttributes) GetFacebookPixelIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.FacebookPixelId) {
		return nil, false
	}
	return &o.FacebookPixelId, true
}

// HasFacebookPixelId returns a boolean if a field has been set.
func (o *GETMarketsMarketId200ResponseDataAttributes) HasFacebookPixelId() bool {
	if o != nil && IsNil(o.FacebookPixelId) {
		return true
	}

	return false
}

// SetFacebookPixelId gets a reference to the given interface{} and assigns it to the FacebookPixelId field.
func (o *GETMarketsMarketId200ResponseDataAttributes) SetFacebookPixelId(v interface{}) {
	o.FacebookPixelId = v
}

// GetCheckoutUrl returns the CheckoutUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETMarketsMarketId200ResponseDataAttributes) GetCheckoutUrl() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CheckoutUrl
}

// GetCheckoutUrlOk returns a tuple with the CheckoutUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETMarketsMarketId200ResponseDataAttributes) GetCheckoutUrlOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CheckoutUrl) {
		return nil, false
	}
	return &o.CheckoutUrl, true
}

// HasCheckoutUrl returns a boolean if a field has been set.
func (o *GETMarketsMarketId200ResponseDataAttributes) HasCheckoutUrl() bool {
	if o != nil && IsNil(o.CheckoutUrl) {
		return true
	}

	return false
}

// SetCheckoutUrl gets a reference to the given interface{} and assigns it to the CheckoutUrl field.
func (o *GETMarketsMarketId200ResponseDataAttributes) SetCheckoutUrl(v interface{}) {
	o.CheckoutUrl = v
}

// GetExternalPricesUrl returns the ExternalPricesUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETMarketsMarketId200ResponseDataAttributes) GetExternalPricesUrl() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ExternalPricesUrl
}

// GetExternalPricesUrlOk returns a tuple with the ExternalPricesUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETMarketsMarketId200ResponseDataAttributes) GetExternalPricesUrlOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ExternalPricesUrl) {
		return nil, false
	}
	return &o.ExternalPricesUrl, true
}

// HasExternalPricesUrl returns a boolean if a field has been set.
func (o *GETMarketsMarketId200ResponseDataAttributes) HasExternalPricesUrl() bool {
	if o != nil && IsNil(o.ExternalPricesUrl) {
		return true
	}

	return false
}

// SetExternalPricesUrl gets a reference to the given interface{} and assigns it to the ExternalPricesUrl field.
func (o *GETMarketsMarketId200ResponseDataAttributes) SetExternalPricesUrl(v interface{}) {
	o.ExternalPricesUrl = v
}

// GetExternalOrderValidationUrl returns the ExternalOrderValidationUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETMarketsMarketId200ResponseDataAttributes) GetExternalOrderValidationUrl() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ExternalOrderValidationUrl
}

// GetExternalOrderValidationUrlOk returns a tuple with the ExternalOrderValidationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETMarketsMarketId200ResponseDataAttributes) GetExternalOrderValidationUrlOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ExternalOrderValidationUrl) {
		return nil, false
	}
	return &o.ExternalOrderValidationUrl, true
}

// HasExternalOrderValidationUrl returns a boolean if a field has been set.
func (o *GETMarketsMarketId200ResponseDataAttributes) HasExternalOrderValidationUrl() bool {
	if o != nil && IsNil(o.ExternalOrderValidationUrl) {
		return true
	}

	return false
}

// SetExternalOrderValidationUrl gets a reference to the given interface{} and assigns it to the ExternalOrderValidationUrl field.
func (o *GETMarketsMarketId200ResponseDataAttributes) SetExternalOrderValidationUrl(v interface{}) {
	o.ExternalOrderValidationUrl = v
}

// GetPrivate returns the Private field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETMarketsMarketId200ResponseDataAttributes) GetPrivate() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Private
}

// GetPrivateOk returns a tuple with the Private field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETMarketsMarketId200ResponseDataAttributes) GetPrivateOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Private) {
		return nil, false
	}
	return &o.Private, true
}

// HasPrivate returns a boolean if a field has been set.
func (o *GETMarketsMarketId200ResponseDataAttributes) HasPrivate() bool {
	if o != nil && IsNil(o.Private) {
		return true
	}

	return false
}

// SetPrivate gets a reference to the given interface{} and assigns it to the Private field.
func (o *GETMarketsMarketId200ResponseDataAttributes) SetPrivate(v interface{}) {
	o.Private = v
}

// GetShippingCostCutoff returns the ShippingCostCutoff field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETMarketsMarketId200ResponseDataAttributes) GetShippingCostCutoff() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ShippingCostCutoff
}

// GetShippingCostCutoffOk returns a tuple with the ShippingCostCutoff field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETMarketsMarketId200ResponseDataAttributes) GetShippingCostCutoffOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ShippingCostCutoff) {
		return nil, false
	}
	return &o.ShippingCostCutoff, true
}

// HasShippingCostCutoff returns a boolean if a field has been set.
func (o *GETMarketsMarketId200ResponseDataAttributes) HasShippingCostCutoff() bool {
	if o != nil && IsNil(o.ShippingCostCutoff) {
		return true
	}

	return false
}

// SetShippingCostCutoff gets a reference to the given interface{} and assigns it to the ShippingCostCutoff field.
func (o *GETMarketsMarketId200ResponseDataAttributes) SetShippingCostCutoff(v interface{}) {
	o.ShippingCostCutoff = v
}

// GetDisabledAt returns the DisabledAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETMarketsMarketId200ResponseDataAttributes) GetDisabledAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.DisabledAt
}

// GetDisabledAtOk returns a tuple with the DisabledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETMarketsMarketId200ResponseDataAttributes) GetDisabledAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.DisabledAt) {
		return nil, false
	}
	return &o.DisabledAt, true
}

// HasDisabledAt returns a boolean if a field has been set.
func (o *GETMarketsMarketId200ResponseDataAttributes) HasDisabledAt() bool {
	if o != nil && IsNil(o.DisabledAt) {
		return true
	}

	return false
}

// SetDisabledAt gets a reference to the given interface{} and assigns it to the DisabledAt field.
func (o *GETMarketsMarketId200ResponseDataAttributes) SetDisabledAt(v interface{}) {
	o.DisabledAt = v
}

// GetSharedSecret returns the SharedSecret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETMarketsMarketId200ResponseDataAttributes) GetSharedSecret() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.SharedSecret
}

// GetSharedSecretOk returns a tuple with the SharedSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETMarketsMarketId200ResponseDataAttributes) GetSharedSecretOk() (*interface{}, bool) {
	if o == nil || IsNil(o.SharedSecret) {
		return nil, false
	}
	return &o.SharedSecret, true
}

// HasSharedSecret returns a boolean if a field has been set.
func (o *GETMarketsMarketId200ResponseDataAttributes) HasSharedSecret() bool {
	if o != nil && IsNil(o.SharedSecret) {
		return true
	}

	return false
}

// SetSharedSecret gets a reference to the given interface{} and assigns it to the SharedSecret field.
func (o *GETMarketsMarketId200ResponseDataAttributes) SetSharedSecret(v interface{}) {
	o.SharedSecret = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETMarketsMarketId200ResponseDataAttributes) GetCreatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETMarketsMarketId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return &o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETMarketsMarketId200ResponseDataAttributes) HasCreatedAt() bool {
	if o != nil && IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given interface{} and assigns it to the CreatedAt field.
func (o *GETMarketsMarketId200ResponseDataAttributes) SetCreatedAt(v interface{}) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETMarketsMarketId200ResponseDataAttributes) GetUpdatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETMarketsMarketId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETMarketsMarketId200ResponseDataAttributes) HasUpdatedAt() bool {
	if o != nil && IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given interface{} and assigns it to the UpdatedAt field.
func (o *GETMarketsMarketId200ResponseDataAttributes) SetUpdatedAt(v interface{}) {
	o.UpdatedAt = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETMarketsMarketId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETMarketsMarketId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETMarketsMarketId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *GETMarketsMarketId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETMarketsMarketId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETMarketsMarketId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETMarketsMarketId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *GETMarketsMarketId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETMarketsMarketId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETMarketsMarketId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETMarketsMarketId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *GETMarketsMarketId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o GETMarketsMarketId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETMarketsMarketId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Number != nil {
		toSerialize["number"] = o.Number
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.FacebookPixelId != nil {
		toSerialize["facebook_pixel_id"] = o.FacebookPixelId
	}
	if o.CheckoutUrl != nil {
		toSerialize["checkout_url"] = o.CheckoutUrl
	}
	if o.ExternalPricesUrl != nil {
		toSerialize["external_prices_url"] = o.ExternalPricesUrl
	}
	if o.ExternalOrderValidationUrl != nil {
		toSerialize["external_order_validation_url"] = o.ExternalOrderValidationUrl
	}
	if o.Private != nil {
		toSerialize["private"] = o.Private
	}
	if o.ShippingCostCutoff != nil {
		toSerialize["shipping_cost_cutoff"] = o.ShippingCostCutoff
	}
	if o.DisabledAt != nil {
		toSerialize["disabled_at"] = o.DisabledAt
	}
	if o.SharedSecret != nil {
		toSerialize["shared_secret"] = o.SharedSecret
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

type NullableGETMarketsMarketId200ResponseDataAttributes struct {
	value *GETMarketsMarketId200ResponseDataAttributes
	isSet bool
}

func (v NullableGETMarketsMarketId200ResponseDataAttributes) Get() *GETMarketsMarketId200ResponseDataAttributes {
	return v.value
}

func (v *NullableGETMarketsMarketId200ResponseDataAttributes) Set(val *GETMarketsMarketId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETMarketsMarketId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETMarketsMarketId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETMarketsMarketId200ResponseDataAttributes(val *GETMarketsMarketId200ResponseDataAttributes) *NullableGETMarketsMarketId200ResponseDataAttributes {
	return &NullableGETMarketsMarketId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullableGETMarketsMarketId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETMarketsMarketId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
