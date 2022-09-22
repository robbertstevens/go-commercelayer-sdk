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

// GETAdjustments200ResponseDataInnerAttributes struct for GETAdjustments200ResponseDataInnerAttributes
type GETAdjustments200ResponseDataInnerAttributes struct {
	// The adjustment name
	Name *string `json:"name,omitempty"`
	// The international 3-letter currency code as defined by the ISO 4217 standard.
	CurrencyCode *string `json:"currency_code,omitempty"`
	// The adjustment amount, in cents.
	AmountCents *int32 `json:"amount_cents,omitempty"`
	// The adjustment amount, float.
	AmountFloat *float32 `json:"amount_float,omitempty"`
	// The adjustment amount, formatted.
	FormattedAmount *string `json:"formatted_amount,omitempty"`
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
}

// NewGETAdjustments200ResponseDataInnerAttributes instantiates a new GETAdjustments200ResponseDataInnerAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETAdjustments200ResponseDataInnerAttributes() *GETAdjustments200ResponseDataInnerAttributes {
	this := GETAdjustments200ResponseDataInnerAttributes{}
	return &this
}

// NewGETAdjustments200ResponseDataInnerAttributesWithDefaults instantiates a new GETAdjustments200ResponseDataInnerAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETAdjustments200ResponseDataInnerAttributesWithDefaults() *GETAdjustments200ResponseDataInnerAttributes {
	this := GETAdjustments200ResponseDataInnerAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GETAdjustments200ResponseDataInnerAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAdjustments200ResponseDataInnerAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GETAdjustments200ResponseDataInnerAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GETAdjustments200ResponseDataInnerAttributes) SetName(v string) {
	o.Name = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *GETAdjustments200ResponseDataInnerAttributes) GetCurrencyCode() string {
	if o == nil || o.CurrencyCode == nil {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAdjustments200ResponseDataInnerAttributes) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || o.CurrencyCode == nil {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *GETAdjustments200ResponseDataInnerAttributes) HasCurrencyCode() bool {
	if o != nil && o.CurrencyCode != nil {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *GETAdjustments200ResponseDataInnerAttributes) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetAmountCents returns the AmountCents field value if set, zero value otherwise.
func (o *GETAdjustments200ResponseDataInnerAttributes) GetAmountCents() int32 {
	if o == nil || o.AmountCents == nil {
		var ret int32
		return ret
	}
	return *o.AmountCents
}

// GetAmountCentsOk returns a tuple with the AmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAdjustments200ResponseDataInnerAttributes) GetAmountCentsOk() (*int32, bool) {
	if o == nil || o.AmountCents == nil {
		return nil, false
	}
	return o.AmountCents, true
}

// HasAmountCents returns a boolean if a field has been set.
func (o *GETAdjustments200ResponseDataInnerAttributes) HasAmountCents() bool {
	if o != nil && o.AmountCents != nil {
		return true
	}

	return false
}

// SetAmountCents gets a reference to the given int32 and assigns it to the AmountCents field.
func (o *GETAdjustments200ResponseDataInnerAttributes) SetAmountCents(v int32) {
	o.AmountCents = &v
}

// GetAmountFloat returns the AmountFloat field value if set, zero value otherwise.
func (o *GETAdjustments200ResponseDataInnerAttributes) GetAmountFloat() float32 {
	if o == nil || o.AmountFloat == nil {
		var ret float32
		return ret
	}
	return *o.AmountFloat
}

// GetAmountFloatOk returns a tuple with the AmountFloat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAdjustments200ResponseDataInnerAttributes) GetAmountFloatOk() (*float32, bool) {
	if o == nil || o.AmountFloat == nil {
		return nil, false
	}
	return o.AmountFloat, true
}

// HasAmountFloat returns a boolean if a field has been set.
func (o *GETAdjustments200ResponseDataInnerAttributes) HasAmountFloat() bool {
	if o != nil && o.AmountFloat != nil {
		return true
	}

	return false
}

// SetAmountFloat gets a reference to the given float32 and assigns it to the AmountFloat field.
func (o *GETAdjustments200ResponseDataInnerAttributes) SetAmountFloat(v float32) {
	o.AmountFloat = &v
}

// GetFormattedAmount returns the FormattedAmount field value if set, zero value otherwise.
func (o *GETAdjustments200ResponseDataInnerAttributes) GetFormattedAmount() string {
	if o == nil || o.FormattedAmount == nil {
		var ret string
		return ret
	}
	return *o.FormattedAmount
}

// GetFormattedAmountOk returns a tuple with the FormattedAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAdjustments200ResponseDataInnerAttributes) GetFormattedAmountOk() (*string, bool) {
	if o == nil || o.FormattedAmount == nil {
		return nil, false
	}
	return o.FormattedAmount, true
}

// HasFormattedAmount returns a boolean if a field has been set.
func (o *GETAdjustments200ResponseDataInnerAttributes) HasFormattedAmount() bool {
	if o != nil && o.FormattedAmount != nil {
		return true
	}

	return false
}

// SetFormattedAmount gets a reference to the given string and assigns it to the FormattedAmount field.
func (o *GETAdjustments200ResponseDataInnerAttributes) SetFormattedAmount(v string) {
	o.FormattedAmount = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETAdjustments200ResponseDataInnerAttributes) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAdjustments200ResponseDataInnerAttributes) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETAdjustments200ResponseDataInnerAttributes) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETAdjustments200ResponseDataInnerAttributes) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GETAdjustments200ResponseDataInnerAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAdjustments200ResponseDataInnerAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETAdjustments200ResponseDataInnerAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *GETAdjustments200ResponseDataInnerAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GETAdjustments200ResponseDataInnerAttributes) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAdjustments200ResponseDataInnerAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETAdjustments200ResponseDataInnerAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *GETAdjustments200ResponseDataInnerAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *GETAdjustments200ResponseDataInnerAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAdjustments200ResponseDataInnerAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETAdjustments200ResponseDataInnerAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *GETAdjustments200ResponseDataInnerAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *GETAdjustments200ResponseDataInnerAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAdjustments200ResponseDataInnerAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETAdjustments200ResponseDataInnerAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *GETAdjustments200ResponseDataInnerAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GETAdjustments200ResponseDataInnerAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAdjustments200ResponseDataInnerAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETAdjustments200ResponseDataInnerAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *GETAdjustments200ResponseDataInnerAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o GETAdjustments200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.CurrencyCode != nil {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	if o.AmountCents != nil {
		toSerialize["amount_cents"] = o.AmountCents
	}
	if o.AmountFloat != nil {
		toSerialize["amount_float"] = o.AmountFloat
	}
	if o.FormattedAmount != nil {
		toSerialize["formatted_amount"] = o.FormattedAmount
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
	return json.Marshal(toSerialize)
}

type NullableGETAdjustments200ResponseDataInnerAttributes struct {
	value *GETAdjustments200ResponseDataInnerAttributes
	isSet bool
}

func (v NullableGETAdjustments200ResponseDataInnerAttributes) Get() *GETAdjustments200ResponseDataInnerAttributes {
	return v.value
}

func (v *NullableGETAdjustments200ResponseDataInnerAttributes) Set(val *GETAdjustments200ResponseDataInnerAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETAdjustments200ResponseDataInnerAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETAdjustments200ResponseDataInnerAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETAdjustments200ResponseDataInnerAttributes(val *GETAdjustments200ResponseDataInnerAttributes) *NullableGETAdjustments200ResponseDataInnerAttributes {
	return &NullableGETAdjustments200ResponseDataInnerAttributes{value: val, isSet: true}
}

func (v NullableGETAdjustments200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETAdjustments200ResponseDataInnerAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
