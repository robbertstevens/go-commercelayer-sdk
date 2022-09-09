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

// GETTaxRules200ResponseDataInnerAttributes struct for GETTaxRules200ResponseDataInnerAttributes
type GETTaxRules200ResponseDataInnerAttributes struct {
	// The tax rule internal name.
	Name *string `json:"name,omitempty"`
	// The tax rate for this ruke.
	TaxRate *float32 `json:"tax_rate,omitempty"`
	// The regex that will be evaluated to match the shipping address country code.
	CountryCodeRegex *string `json:"country_code_regex,omitempty"`
	// The regex that will be evaluated as negative match for the shipping address country code.
	NotCountryCodeRegex *string `json:"not_country_code_regex,omitempty"`
	// The regex that will be evaluated to match the shipping address state code.
	StateCodeRegex *string `json:"state_code_regex,omitempty"`
	// The regex that will be evaluated as negative match for the shipping address state code.
	NotStateCodeRegex *string `json:"not_state_code_regex,omitempty"`
	// The regex that will be evaluated to match the shipping address zip code.
	ZipCodeRegex *string `json:"zip_code_regex,omitempty"`
	// The regex that will be evaluated as negative match for the shipping zip country code.
	NotZipCodeRegex *string `json:"not_zip_code_regex,omitempty"`
	// Indicates if the freight is taxable.
	FreightTaxable *bool `json:"freight_taxable,omitempty"`
	// Indicates if the payment method is taxable.
	PaymentMethodTaxable *bool `json:"payment_method_taxable,omitempty"`
	// Indicates if gift cards are taxable.
	GiftCardTaxable *bool `json:"gift_card_taxable,omitempty"`
	// Indicates if adjustemnts are taxable.
	AdjustmentTaxable *bool `json:"adjustment_taxable,omitempty"`
	// The breakdown for this tax rule (if calculated).
	Breakdown map[string]interface{} `json:"breakdown,omitempty"`
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

// NewGETTaxRules200ResponseDataInnerAttributes instantiates a new GETTaxRules200ResponseDataInnerAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETTaxRules200ResponseDataInnerAttributes() *GETTaxRules200ResponseDataInnerAttributes {
	this := GETTaxRules200ResponseDataInnerAttributes{}
	return &this
}

// NewGETTaxRules200ResponseDataInnerAttributesWithDefaults instantiates a new GETTaxRules200ResponseDataInnerAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETTaxRules200ResponseDataInnerAttributesWithDefaults() *GETTaxRules200ResponseDataInnerAttributes {
	this := GETTaxRules200ResponseDataInnerAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GETTaxRules200ResponseDataInnerAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETTaxRules200ResponseDataInnerAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GETTaxRules200ResponseDataInnerAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GETTaxRules200ResponseDataInnerAttributes) SetName(v string) {
	o.Name = &v
}

// GetTaxRate returns the TaxRate field value if set, zero value otherwise.
func (o *GETTaxRules200ResponseDataInnerAttributes) GetTaxRate() float32 {
	if o == nil || o.TaxRate == nil {
		var ret float32
		return ret
	}
	return *o.TaxRate
}

// GetTaxRateOk returns a tuple with the TaxRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETTaxRules200ResponseDataInnerAttributes) GetTaxRateOk() (*float32, bool) {
	if o == nil || o.TaxRate == nil {
		return nil, false
	}
	return o.TaxRate, true
}

// HasTaxRate returns a boolean if a field has been set.
func (o *GETTaxRules200ResponseDataInnerAttributes) HasTaxRate() bool {
	if o != nil && o.TaxRate != nil {
		return true
	}

	return false
}

// SetTaxRate gets a reference to the given float32 and assigns it to the TaxRate field.
func (o *GETTaxRules200ResponseDataInnerAttributes) SetTaxRate(v float32) {
	o.TaxRate = &v
}

// GetCountryCodeRegex returns the CountryCodeRegex field value if set, zero value otherwise.
func (o *GETTaxRules200ResponseDataInnerAttributes) GetCountryCodeRegex() string {
	if o == nil || o.CountryCodeRegex == nil {
		var ret string
		return ret
	}
	return *o.CountryCodeRegex
}

// GetCountryCodeRegexOk returns a tuple with the CountryCodeRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETTaxRules200ResponseDataInnerAttributes) GetCountryCodeRegexOk() (*string, bool) {
	if o == nil || o.CountryCodeRegex == nil {
		return nil, false
	}
	return o.CountryCodeRegex, true
}

// HasCountryCodeRegex returns a boolean if a field has been set.
func (o *GETTaxRules200ResponseDataInnerAttributes) HasCountryCodeRegex() bool {
	if o != nil && o.CountryCodeRegex != nil {
		return true
	}

	return false
}

// SetCountryCodeRegex gets a reference to the given string and assigns it to the CountryCodeRegex field.
func (o *GETTaxRules200ResponseDataInnerAttributes) SetCountryCodeRegex(v string) {
	o.CountryCodeRegex = &v
}

// GetNotCountryCodeRegex returns the NotCountryCodeRegex field value if set, zero value otherwise.
func (o *GETTaxRules200ResponseDataInnerAttributes) GetNotCountryCodeRegex() string {
	if o == nil || o.NotCountryCodeRegex == nil {
		var ret string
		return ret
	}
	return *o.NotCountryCodeRegex
}

// GetNotCountryCodeRegexOk returns a tuple with the NotCountryCodeRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETTaxRules200ResponseDataInnerAttributes) GetNotCountryCodeRegexOk() (*string, bool) {
	if o == nil || o.NotCountryCodeRegex == nil {
		return nil, false
	}
	return o.NotCountryCodeRegex, true
}

// HasNotCountryCodeRegex returns a boolean if a field has been set.
func (o *GETTaxRules200ResponseDataInnerAttributes) HasNotCountryCodeRegex() bool {
	if o != nil && o.NotCountryCodeRegex != nil {
		return true
	}

	return false
}

// SetNotCountryCodeRegex gets a reference to the given string and assigns it to the NotCountryCodeRegex field.
func (o *GETTaxRules200ResponseDataInnerAttributes) SetNotCountryCodeRegex(v string) {
	o.NotCountryCodeRegex = &v
}

// GetStateCodeRegex returns the StateCodeRegex field value if set, zero value otherwise.
func (o *GETTaxRules200ResponseDataInnerAttributes) GetStateCodeRegex() string {
	if o == nil || o.StateCodeRegex == nil {
		var ret string
		return ret
	}
	return *o.StateCodeRegex
}

// GetStateCodeRegexOk returns a tuple with the StateCodeRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETTaxRules200ResponseDataInnerAttributes) GetStateCodeRegexOk() (*string, bool) {
	if o == nil || o.StateCodeRegex == nil {
		return nil, false
	}
	return o.StateCodeRegex, true
}

// HasStateCodeRegex returns a boolean if a field has been set.
func (o *GETTaxRules200ResponseDataInnerAttributes) HasStateCodeRegex() bool {
	if o != nil && o.StateCodeRegex != nil {
		return true
	}

	return false
}

// SetStateCodeRegex gets a reference to the given string and assigns it to the StateCodeRegex field.
func (o *GETTaxRules200ResponseDataInnerAttributes) SetStateCodeRegex(v string) {
	o.StateCodeRegex = &v
}

// GetNotStateCodeRegex returns the NotStateCodeRegex field value if set, zero value otherwise.
func (o *GETTaxRules200ResponseDataInnerAttributes) GetNotStateCodeRegex() string {
	if o == nil || o.NotStateCodeRegex == nil {
		var ret string
		return ret
	}
	return *o.NotStateCodeRegex
}

// GetNotStateCodeRegexOk returns a tuple with the NotStateCodeRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETTaxRules200ResponseDataInnerAttributes) GetNotStateCodeRegexOk() (*string, bool) {
	if o == nil || o.NotStateCodeRegex == nil {
		return nil, false
	}
	return o.NotStateCodeRegex, true
}

// HasNotStateCodeRegex returns a boolean if a field has been set.
func (o *GETTaxRules200ResponseDataInnerAttributes) HasNotStateCodeRegex() bool {
	if o != nil && o.NotStateCodeRegex != nil {
		return true
	}

	return false
}

// SetNotStateCodeRegex gets a reference to the given string and assigns it to the NotStateCodeRegex field.
func (o *GETTaxRules200ResponseDataInnerAttributes) SetNotStateCodeRegex(v string) {
	o.NotStateCodeRegex = &v
}

// GetZipCodeRegex returns the ZipCodeRegex field value if set, zero value otherwise.
func (o *GETTaxRules200ResponseDataInnerAttributes) GetZipCodeRegex() string {
	if o == nil || o.ZipCodeRegex == nil {
		var ret string
		return ret
	}
	return *o.ZipCodeRegex
}

// GetZipCodeRegexOk returns a tuple with the ZipCodeRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETTaxRules200ResponseDataInnerAttributes) GetZipCodeRegexOk() (*string, bool) {
	if o == nil || o.ZipCodeRegex == nil {
		return nil, false
	}
	return o.ZipCodeRegex, true
}

// HasZipCodeRegex returns a boolean if a field has been set.
func (o *GETTaxRules200ResponseDataInnerAttributes) HasZipCodeRegex() bool {
	if o != nil && o.ZipCodeRegex != nil {
		return true
	}

	return false
}

// SetZipCodeRegex gets a reference to the given string and assigns it to the ZipCodeRegex field.
func (o *GETTaxRules200ResponseDataInnerAttributes) SetZipCodeRegex(v string) {
	o.ZipCodeRegex = &v
}

// GetNotZipCodeRegex returns the NotZipCodeRegex field value if set, zero value otherwise.
func (o *GETTaxRules200ResponseDataInnerAttributes) GetNotZipCodeRegex() string {
	if o == nil || o.NotZipCodeRegex == nil {
		var ret string
		return ret
	}
	return *o.NotZipCodeRegex
}

// GetNotZipCodeRegexOk returns a tuple with the NotZipCodeRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETTaxRules200ResponseDataInnerAttributes) GetNotZipCodeRegexOk() (*string, bool) {
	if o == nil || o.NotZipCodeRegex == nil {
		return nil, false
	}
	return o.NotZipCodeRegex, true
}

// HasNotZipCodeRegex returns a boolean if a field has been set.
func (o *GETTaxRules200ResponseDataInnerAttributes) HasNotZipCodeRegex() bool {
	if o != nil && o.NotZipCodeRegex != nil {
		return true
	}

	return false
}

// SetNotZipCodeRegex gets a reference to the given string and assigns it to the NotZipCodeRegex field.
func (o *GETTaxRules200ResponseDataInnerAttributes) SetNotZipCodeRegex(v string) {
	o.NotZipCodeRegex = &v
}

// GetFreightTaxable returns the FreightTaxable field value if set, zero value otherwise.
func (o *GETTaxRules200ResponseDataInnerAttributes) GetFreightTaxable() bool {
	if o == nil || o.FreightTaxable == nil {
		var ret bool
		return ret
	}
	return *o.FreightTaxable
}

// GetFreightTaxableOk returns a tuple with the FreightTaxable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETTaxRules200ResponseDataInnerAttributes) GetFreightTaxableOk() (*bool, bool) {
	if o == nil || o.FreightTaxable == nil {
		return nil, false
	}
	return o.FreightTaxable, true
}

// HasFreightTaxable returns a boolean if a field has been set.
func (o *GETTaxRules200ResponseDataInnerAttributes) HasFreightTaxable() bool {
	if o != nil && o.FreightTaxable != nil {
		return true
	}

	return false
}

// SetFreightTaxable gets a reference to the given bool and assigns it to the FreightTaxable field.
func (o *GETTaxRules200ResponseDataInnerAttributes) SetFreightTaxable(v bool) {
	o.FreightTaxable = &v
}

// GetPaymentMethodTaxable returns the PaymentMethodTaxable field value if set, zero value otherwise.
func (o *GETTaxRules200ResponseDataInnerAttributes) GetPaymentMethodTaxable() bool {
	if o == nil || o.PaymentMethodTaxable == nil {
		var ret bool
		return ret
	}
	return *o.PaymentMethodTaxable
}

// GetPaymentMethodTaxableOk returns a tuple with the PaymentMethodTaxable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETTaxRules200ResponseDataInnerAttributes) GetPaymentMethodTaxableOk() (*bool, bool) {
	if o == nil || o.PaymentMethodTaxable == nil {
		return nil, false
	}
	return o.PaymentMethodTaxable, true
}

// HasPaymentMethodTaxable returns a boolean if a field has been set.
func (o *GETTaxRules200ResponseDataInnerAttributes) HasPaymentMethodTaxable() bool {
	if o != nil && o.PaymentMethodTaxable != nil {
		return true
	}

	return false
}

// SetPaymentMethodTaxable gets a reference to the given bool and assigns it to the PaymentMethodTaxable field.
func (o *GETTaxRules200ResponseDataInnerAttributes) SetPaymentMethodTaxable(v bool) {
	o.PaymentMethodTaxable = &v
}

// GetGiftCardTaxable returns the GiftCardTaxable field value if set, zero value otherwise.
func (o *GETTaxRules200ResponseDataInnerAttributes) GetGiftCardTaxable() bool {
	if o == nil || o.GiftCardTaxable == nil {
		var ret bool
		return ret
	}
	return *o.GiftCardTaxable
}

// GetGiftCardTaxableOk returns a tuple with the GiftCardTaxable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETTaxRules200ResponseDataInnerAttributes) GetGiftCardTaxableOk() (*bool, bool) {
	if o == nil || o.GiftCardTaxable == nil {
		return nil, false
	}
	return o.GiftCardTaxable, true
}

// HasGiftCardTaxable returns a boolean if a field has been set.
func (o *GETTaxRules200ResponseDataInnerAttributes) HasGiftCardTaxable() bool {
	if o != nil && o.GiftCardTaxable != nil {
		return true
	}

	return false
}

// SetGiftCardTaxable gets a reference to the given bool and assigns it to the GiftCardTaxable field.
func (o *GETTaxRules200ResponseDataInnerAttributes) SetGiftCardTaxable(v bool) {
	o.GiftCardTaxable = &v
}

// GetAdjustmentTaxable returns the AdjustmentTaxable field value if set, zero value otherwise.
func (o *GETTaxRules200ResponseDataInnerAttributes) GetAdjustmentTaxable() bool {
	if o == nil || o.AdjustmentTaxable == nil {
		var ret bool
		return ret
	}
	return *o.AdjustmentTaxable
}

// GetAdjustmentTaxableOk returns a tuple with the AdjustmentTaxable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETTaxRules200ResponseDataInnerAttributes) GetAdjustmentTaxableOk() (*bool, bool) {
	if o == nil || o.AdjustmentTaxable == nil {
		return nil, false
	}
	return o.AdjustmentTaxable, true
}

// HasAdjustmentTaxable returns a boolean if a field has been set.
func (o *GETTaxRules200ResponseDataInnerAttributes) HasAdjustmentTaxable() bool {
	if o != nil && o.AdjustmentTaxable != nil {
		return true
	}

	return false
}

// SetAdjustmentTaxable gets a reference to the given bool and assigns it to the AdjustmentTaxable field.
func (o *GETTaxRules200ResponseDataInnerAttributes) SetAdjustmentTaxable(v bool) {
	o.AdjustmentTaxable = &v
}

// GetBreakdown returns the Breakdown field value if set, zero value otherwise.
func (o *GETTaxRules200ResponseDataInnerAttributes) GetBreakdown() map[string]interface{} {
	if o == nil || o.Breakdown == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Breakdown
}

// GetBreakdownOk returns a tuple with the Breakdown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETTaxRules200ResponseDataInnerAttributes) GetBreakdownOk() (map[string]interface{}, bool) {
	if o == nil || o.Breakdown == nil {
		return nil, false
	}
	return o.Breakdown, true
}

// HasBreakdown returns a boolean if a field has been set.
func (o *GETTaxRules200ResponseDataInnerAttributes) HasBreakdown() bool {
	if o != nil && o.Breakdown != nil {
		return true
	}

	return false
}

// SetBreakdown gets a reference to the given map[string]interface{} and assigns it to the Breakdown field.
func (o *GETTaxRules200ResponseDataInnerAttributes) SetBreakdown(v map[string]interface{}) {
	o.Breakdown = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETTaxRules200ResponseDataInnerAttributes) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETTaxRules200ResponseDataInnerAttributes) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETTaxRules200ResponseDataInnerAttributes) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETTaxRules200ResponseDataInnerAttributes) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GETTaxRules200ResponseDataInnerAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETTaxRules200ResponseDataInnerAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETTaxRules200ResponseDataInnerAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *GETTaxRules200ResponseDataInnerAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GETTaxRules200ResponseDataInnerAttributes) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETTaxRules200ResponseDataInnerAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETTaxRules200ResponseDataInnerAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *GETTaxRules200ResponseDataInnerAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *GETTaxRules200ResponseDataInnerAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETTaxRules200ResponseDataInnerAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETTaxRules200ResponseDataInnerAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *GETTaxRules200ResponseDataInnerAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *GETTaxRules200ResponseDataInnerAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETTaxRules200ResponseDataInnerAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETTaxRules200ResponseDataInnerAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *GETTaxRules200ResponseDataInnerAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GETTaxRules200ResponseDataInnerAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETTaxRules200ResponseDataInnerAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETTaxRules200ResponseDataInnerAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *GETTaxRules200ResponseDataInnerAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o GETTaxRules200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.TaxRate != nil {
		toSerialize["tax_rate"] = o.TaxRate
	}
	if o.CountryCodeRegex != nil {
		toSerialize["country_code_regex"] = o.CountryCodeRegex
	}
	if o.NotCountryCodeRegex != nil {
		toSerialize["not_country_code_regex"] = o.NotCountryCodeRegex
	}
	if o.StateCodeRegex != nil {
		toSerialize["state_code_regex"] = o.StateCodeRegex
	}
	if o.NotStateCodeRegex != nil {
		toSerialize["not_state_code_regex"] = o.NotStateCodeRegex
	}
	if o.ZipCodeRegex != nil {
		toSerialize["zip_code_regex"] = o.ZipCodeRegex
	}
	if o.NotZipCodeRegex != nil {
		toSerialize["not_zip_code_regex"] = o.NotZipCodeRegex
	}
	if o.FreightTaxable != nil {
		toSerialize["freight_taxable"] = o.FreightTaxable
	}
	if o.PaymentMethodTaxable != nil {
		toSerialize["payment_method_taxable"] = o.PaymentMethodTaxable
	}
	if o.GiftCardTaxable != nil {
		toSerialize["gift_card_taxable"] = o.GiftCardTaxable
	}
	if o.AdjustmentTaxable != nil {
		toSerialize["adjustment_taxable"] = o.AdjustmentTaxable
	}
	if o.Breakdown != nil {
		toSerialize["breakdown"] = o.Breakdown
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

type NullableGETTaxRules200ResponseDataInnerAttributes struct {
	value *GETTaxRules200ResponseDataInnerAttributes
	isSet bool
}

func (v NullableGETTaxRules200ResponseDataInnerAttributes) Get() *GETTaxRules200ResponseDataInnerAttributes {
	return v.value
}

func (v *NullableGETTaxRules200ResponseDataInnerAttributes) Set(val *GETTaxRules200ResponseDataInnerAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETTaxRules200ResponseDataInnerAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETTaxRules200ResponseDataInnerAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETTaxRules200ResponseDataInnerAttributes(val *GETTaxRules200ResponseDataInnerAttributes) *NullableGETTaxRules200ResponseDataInnerAttributes {
	return &NullableGETTaxRules200ResponseDataInnerAttributes{value: val, isSet: true}
}

func (v NullableGETTaxRules200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETTaxRules200ResponseDataInnerAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


