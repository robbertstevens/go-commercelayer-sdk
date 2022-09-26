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

// GETSkuOptions200ResponseDataInnerAttributes struct for GETSkuOptions200ResponseDataInnerAttributes
type GETSkuOptions200ResponseDataInnerAttributes struct {
	// The SKU option's internal name
	Name *string `json:"name,omitempty"`
	// The international 3-letter currency code as defined by the ISO 4217 standard.
	CurrencyCode *string `json:"currency_code,omitempty"`
	// An internal description of the SKU option.
	Description *string `json:"description,omitempty"`
	// The price of this shipping method, in cents.
	PriceAmountCents *int32 `json:"price_amount_cents,omitempty"`
	// The price of this shipping method, float.
	PriceAmountFloat *float32 `json:"price_amount_float,omitempty"`
	// The price of this shipping method, formatted.
	FormattedPriceAmount *string `json:"formatted_price_amount,omitempty"`
	// The delay time (in hours) that should be added to the delivery lead time when this option is purchased.
	DelayHours *int32 `json:"delay_hours,omitempty"`
	// The delay time, in days (rounded)
	DelayDays *int32 `json:"delay_days,omitempty"`
	// The regex that will be evaluated to match the SKU codes.
	SkuCodeRegex *string `json:"sku_code_regex,omitempty"`
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

// NewGETSkuOptions200ResponseDataInnerAttributes instantiates a new GETSkuOptions200ResponseDataInnerAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETSkuOptions200ResponseDataInnerAttributes() *GETSkuOptions200ResponseDataInnerAttributes {
	this := GETSkuOptions200ResponseDataInnerAttributes{}
	return &this
}

// NewGETSkuOptions200ResponseDataInnerAttributesWithDefaults instantiates a new GETSkuOptions200ResponseDataInnerAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETSkuOptions200ResponseDataInnerAttributesWithDefaults() *GETSkuOptions200ResponseDataInnerAttributes {
	this := GETSkuOptions200ResponseDataInnerAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GETSkuOptions200ResponseDataInnerAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkuOptions200ResponseDataInnerAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GETSkuOptions200ResponseDataInnerAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GETSkuOptions200ResponseDataInnerAttributes) SetName(v string) {
	o.Name = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *GETSkuOptions200ResponseDataInnerAttributes) GetCurrencyCode() string {
	if o == nil || o.CurrencyCode == nil {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkuOptions200ResponseDataInnerAttributes) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || o.CurrencyCode == nil {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *GETSkuOptions200ResponseDataInnerAttributes) HasCurrencyCode() bool {
	if o != nil && o.CurrencyCode != nil {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *GETSkuOptions200ResponseDataInnerAttributes) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GETSkuOptions200ResponseDataInnerAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkuOptions200ResponseDataInnerAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GETSkuOptions200ResponseDataInnerAttributes) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GETSkuOptions200ResponseDataInnerAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetPriceAmountCents returns the PriceAmountCents field value if set, zero value otherwise.
func (o *GETSkuOptions200ResponseDataInnerAttributes) GetPriceAmountCents() int32 {
	if o == nil || o.PriceAmountCents == nil {
		var ret int32
		return ret
	}
	return *o.PriceAmountCents
}

// GetPriceAmountCentsOk returns a tuple with the PriceAmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkuOptions200ResponseDataInnerAttributes) GetPriceAmountCentsOk() (*int32, bool) {
	if o == nil || o.PriceAmountCents == nil {
		return nil, false
	}
	return o.PriceAmountCents, true
}

// HasPriceAmountCents returns a boolean if a field has been set.
func (o *GETSkuOptions200ResponseDataInnerAttributes) HasPriceAmountCents() bool {
	if o != nil && o.PriceAmountCents != nil {
		return true
	}

	return false
}

// SetPriceAmountCents gets a reference to the given int32 and assigns it to the PriceAmountCents field.
func (o *GETSkuOptions200ResponseDataInnerAttributes) SetPriceAmountCents(v int32) {
	o.PriceAmountCents = &v
}

// GetPriceAmountFloat returns the PriceAmountFloat field value if set, zero value otherwise.
func (o *GETSkuOptions200ResponseDataInnerAttributes) GetPriceAmountFloat() float32 {
	if o == nil || o.PriceAmountFloat == nil {
		var ret float32
		return ret
	}
	return *o.PriceAmountFloat
}

// GetPriceAmountFloatOk returns a tuple with the PriceAmountFloat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkuOptions200ResponseDataInnerAttributes) GetPriceAmountFloatOk() (*float32, bool) {
	if o == nil || o.PriceAmountFloat == nil {
		return nil, false
	}
	return o.PriceAmountFloat, true
}

// HasPriceAmountFloat returns a boolean if a field has been set.
func (o *GETSkuOptions200ResponseDataInnerAttributes) HasPriceAmountFloat() bool {
	if o != nil && o.PriceAmountFloat != nil {
		return true
	}

	return false
}

// SetPriceAmountFloat gets a reference to the given float32 and assigns it to the PriceAmountFloat field.
func (o *GETSkuOptions200ResponseDataInnerAttributes) SetPriceAmountFloat(v float32) {
	o.PriceAmountFloat = &v
}

// GetFormattedPriceAmount returns the FormattedPriceAmount field value if set, zero value otherwise.
func (o *GETSkuOptions200ResponseDataInnerAttributes) GetFormattedPriceAmount() string {
	if o == nil || o.FormattedPriceAmount == nil {
		var ret string
		return ret
	}
	return *o.FormattedPriceAmount
}

// GetFormattedPriceAmountOk returns a tuple with the FormattedPriceAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkuOptions200ResponseDataInnerAttributes) GetFormattedPriceAmountOk() (*string, bool) {
	if o == nil || o.FormattedPriceAmount == nil {
		return nil, false
	}
	return o.FormattedPriceAmount, true
}

// HasFormattedPriceAmount returns a boolean if a field has been set.
func (o *GETSkuOptions200ResponseDataInnerAttributes) HasFormattedPriceAmount() bool {
	if o != nil && o.FormattedPriceAmount != nil {
		return true
	}

	return false
}

// SetFormattedPriceAmount gets a reference to the given string and assigns it to the FormattedPriceAmount field.
func (o *GETSkuOptions200ResponseDataInnerAttributes) SetFormattedPriceAmount(v string) {
	o.FormattedPriceAmount = &v
}

// GetDelayHours returns the DelayHours field value if set, zero value otherwise.
func (o *GETSkuOptions200ResponseDataInnerAttributes) GetDelayHours() int32 {
	if o == nil || o.DelayHours == nil {
		var ret int32
		return ret
	}
	return *o.DelayHours
}

// GetDelayHoursOk returns a tuple with the DelayHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkuOptions200ResponseDataInnerAttributes) GetDelayHoursOk() (*int32, bool) {
	if o == nil || o.DelayHours == nil {
		return nil, false
	}
	return o.DelayHours, true
}

// HasDelayHours returns a boolean if a field has been set.
func (o *GETSkuOptions200ResponseDataInnerAttributes) HasDelayHours() bool {
	if o != nil && o.DelayHours != nil {
		return true
	}

	return false
}

// SetDelayHours gets a reference to the given int32 and assigns it to the DelayHours field.
func (o *GETSkuOptions200ResponseDataInnerAttributes) SetDelayHours(v int32) {
	o.DelayHours = &v
}

// GetDelayDays returns the DelayDays field value if set, zero value otherwise.
func (o *GETSkuOptions200ResponseDataInnerAttributes) GetDelayDays() int32 {
	if o == nil || o.DelayDays == nil {
		var ret int32
		return ret
	}
	return *o.DelayDays
}

// GetDelayDaysOk returns a tuple with the DelayDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkuOptions200ResponseDataInnerAttributes) GetDelayDaysOk() (*int32, bool) {
	if o == nil || o.DelayDays == nil {
		return nil, false
	}
	return o.DelayDays, true
}

// HasDelayDays returns a boolean if a field has been set.
func (o *GETSkuOptions200ResponseDataInnerAttributes) HasDelayDays() bool {
	if o != nil && o.DelayDays != nil {
		return true
	}

	return false
}

// SetDelayDays gets a reference to the given int32 and assigns it to the DelayDays field.
func (o *GETSkuOptions200ResponseDataInnerAttributes) SetDelayDays(v int32) {
	o.DelayDays = &v
}

// GetSkuCodeRegex returns the SkuCodeRegex field value if set, zero value otherwise.
func (o *GETSkuOptions200ResponseDataInnerAttributes) GetSkuCodeRegex() string {
	if o == nil || o.SkuCodeRegex == nil {
		var ret string
		return ret
	}
	return *o.SkuCodeRegex
}

// GetSkuCodeRegexOk returns a tuple with the SkuCodeRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkuOptions200ResponseDataInnerAttributes) GetSkuCodeRegexOk() (*string, bool) {
	if o == nil || o.SkuCodeRegex == nil {
		return nil, false
	}
	return o.SkuCodeRegex, true
}

// HasSkuCodeRegex returns a boolean if a field has been set.
func (o *GETSkuOptions200ResponseDataInnerAttributes) HasSkuCodeRegex() bool {
	if o != nil && o.SkuCodeRegex != nil {
		return true
	}

	return false
}

// SetSkuCodeRegex gets a reference to the given string and assigns it to the SkuCodeRegex field.
func (o *GETSkuOptions200ResponseDataInnerAttributes) SetSkuCodeRegex(v string) {
	o.SkuCodeRegex = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETSkuOptions200ResponseDataInnerAttributes) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkuOptions200ResponseDataInnerAttributes) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETSkuOptions200ResponseDataInnerAttributes) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETSkuOptions200ResponseDataInnerAttributes) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GETSkuOptions200ResponseDataInnerAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkuOptions200ResponseDataInnerAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETSkuOptions200ResponseDataInnerAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *GETSkuOptions200ResponseDataInnerAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GETSkuOptions200ResponseDataInnerAttributes) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkuOptions200ResponseDataInnerAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETSkuOptions200ResponseDataInnerAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *GETSkuOptions200ResponseDataInnerAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *GETSkuOptions200ResponseDataInnerAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkuOptions200ResponseDataInnerAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETSkuOptions200ResponseDataInnerAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *GETSkuOptions200ResponseDataInnerAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *GETSkuOptions200ResponseDataInnerAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkuOptions200ResponseDataInnerAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETSkuOptions200ResponseDataInnerAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *GETSkuOptions200ResponseDataInnerAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GETSkuOptions200ResponseDataInnerAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkuOptions200ResponseDataInnerAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETSkuOptions200ResponseDataInnerAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *GETSkuOptions200ResponseDataInnerAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o GETSkuOptions200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.CurrencyCode != nil {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.PriceAmountCents != nil {
		toSerialize["price_amount_cents"] = o.PriceAmountCents
	}
	if o.PriceAmountFloat != nil {
		toSerialize["price_amount_float"] = o.PriceAmountFloat
	}
	if o.FormattedPriceAmount != nil {
		toSerialize["formatted_price_amount"] = o.FormattedPriceAmount
	}
	if o.DelayHours != nil {
		toSerialize["delay_hours"] = o.DelayHours
	}
	if o.DelayDays != nil {
		toSerialize["delay_days"] = o.DelayDays
	}
	if o.SkuCodeRegex != nil {
		toSerialize["sku_code_regex"] = o.SkuCodeRegex
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

type NullableGETSkuOptions200ResponseDataInnerAttributes struct {
	value *GETSkuOptions200ResponseDataInnerAttributes
	isSet bool
}

func (v NullableGETSkuOptions200ResponseDataInnerAttributes) Get() *GETSkuOptions200ResponseDataInnerAttributes {
	return v.value
}

func (v *NullableGETSkuOptions200ResponseDataInnerAttributes) Set(val *GETSkuOptions200ResponseDataInnerAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETSkuOptions200ResponseDataInnerAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETSkuOptions200ResponseDataInnerAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETSkuOptions200ResponseDataInnerAttributes(val *GETSkuOptions200ResponseDataInnerAttributes) *NullableGETSkuOptions200ResponseDataInnerAttributes {
	return &NullableGETSkuOptions200ResponseDataInnerAttributes{value: val, isSet: true}
}

func (v NullableGETSkuOptions200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETSkuOptions200ResponseDataInnerAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
