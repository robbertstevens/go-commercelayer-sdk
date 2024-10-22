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

// checks if the GETPricesPriceId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETPricesPriceId200ResponseDataAttributes{}

// GETPricesPriceId200ResponseDataAttributes struct for GETPricesPriceId200ResponseDataAttributes
type GETPricesPriceId200ResponseDataAttributes struct {
	// The international 3-letter currency code as defined by the ISO 4217 standard, inherited from the associated price list.
	CurrencyCode interface{} `json:"currency_code,omitempty"`
	// The code of the associated SKU. When creating a price, either a valid sku_code or a SKU relationship must be present.
	SkuCode interface{} `json:"sku_code,omitempty"`
	// The SKU price amount for the associated price list, in cents.
	AmountCents interface{} `json:"amount_cents,omitempty"`
	// The SKU price amount for the associated price list, float.
	AmountFloat interface{} `json:"amount_float,omitempty"`
	// The SKU price amount for the associated price list, formatted.
	FormattedAmount interface{} `json:"formatted_amount,omitempty"`
	// The SKU price amount for the associated price list, in cents before any applied rule.
	OriginalAmountCents interface{} `json:"original_amount_cents,omitempty"`
	// The SKU price amount for the associated price list, in cents before any applied rule, formatted.
	FormattedOriginalAmount interface{} `json:"formatted_original_amount,omitempty"`
	// The compared price amount, in cents. Useful to display a percentage discount.
	CompareAtAmountCents interface{} `json:"compare_at_amount_cents,omitempty"`
	// The compared price amount, float.
	CompareAtAmountFloat interface{} `json:"compare_at_amount_float,omitempty"`
	// The compared price amount, formatted.
	FormattedCompareAtAmount interface{} `json:"formatted_compare_at_amount,omitempty"`
	// Time at which the resource was created.
	CreatedAt interface{} `json:"created_at,omitempty"`
	// Time at which the resource was last updated.
	UpdatedAt interface{} `json:"updated_at,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// The rules (using Rules Engine) to be applied.
	Rules interface{} `json:"rules,omitempty"`
	// The rule outcomes.
	RuleOutcomes interface{} `json:"rule_outcomes,omitempty"`
	// The custom_claim attached to the current JWT (if any).
	JwtCustomClaim interface{} `json:"jwt_custom_claim,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewGETPricesPriceId200ResponseDataAttributes instantiates a new GETPricesPriceId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETPricesPriceId200ResponseDataAttributes() *GETPricesPriceId200ResponseDataAttributes {
	this := GETPricesPriceId200ResponseDataAttributes{}
	return &this
}

// NewGETPricesPriceId200ResponseDataAttributesWithDefaults instantiates a new GETPricesPriceId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETPricesPriceId200ResponseDataAttributesWithDefaults() *GETPricesPriceId200ResponseDataAttributes {
	this := GETPricesPriceId200ResponseDataAttributes{}
	return &this
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPricesPriceId200ResponseDataAttributes) GetCurrencyCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPricesPriceId200ResponseDataAttributes) GetCurrencyCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *GETPricesPriceId200ResponseDataAttributes) HasCurrencyCode() bool {
	if o != nil && IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given interface{} and assigns it to the CurrencyCode field.
func (o *GETPricesPriceId200ResponseDataAttributes) SetCurrencyCode(v interface{}) {
	o.CurrencyCode = v
}

// GetSkuCode returns the SkuCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPricesPriceId200ResponseDataAttributes) GetSkuCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.SkuCode
}

// GetSkuCodeOk returns a tuple with the SkuCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPricesPriceId200ResponseDataAttributes) GetSkuCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.SkuCode) {
		return nil, false
	}
	return &o.SkuCode, true
}

// HasSkuCode returns a boolean if a field has been set.
func (o *GETPricesPriceId200ResponseDataAttributes) HasSkuCode() bool {
	if o != nil && IsNil(o.SkuCode) {
		return true
	}

	return false
}

// SetSkuCode gets a reference to the given interface{} and assigns it to the SkuCode field.
func (o *GETPricesPriceId200ResponseDataAttributes) SetSkuCode(v interface{}) {
	o.SkuCode = v
}

// GetAmountCents returns the AmountCents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPricesPriceId200ResponseDataAttributes) GetAmountCents() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.AmountCents
}

// GetAmountCentsOk returns a tuple with the AmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPricesPriceId200ResponseDataAttributes) GetAmountCentsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.AmountCents) {
		return nil, false
	}
	return &o.AmountCents, true
}

// HasAmountCents returns a boolean if a field has been set.
func (o *GETPricesPriceId200ResponseDataAttributes) HasAmountCents() bool {
	if o != nil && IsNil(o.AmountCents) {
		return true
	}

	return false
}

// SetAmountCents gets a reference to the given interface{} and assigns it to the AmountCents field.
func (o *GETPricesPriceId200ResponseDataAttributes) SetAmountCents(v interface{}) {
	o.AmountCents = v
}

// GetAmountFloat returns the AmountFloat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPricesPriceId200ResponseDataAttributes) GetAmountFloat() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.AmountFloat
}

// GetAmountFloatOk returns a tuple with the AmountFloat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPricesPriceId200ResponseDataAttributes) GetAmountFloatOk() (*interface{}, bool) {
	if o == nil || IsNil(o.AmountFloat) {
		return nil, false
	}
	return &o.AmountFloat, true
}

// HasAmountFloat returns a boolean if a field has been set.
func (o *GETPricesPriceId200ResponseDataAttributes) HasAmountFloat() bool {
	if o != nil && IsNil(o.AmountFloat) {
		return true
	}

	return false
}

// SetAmountFloat gets a reference to the given interface{} and assigns it to the AmountFloat field.
func (o *GETPricesPriceId200ResponseDataAttributes) SetAmountFloat(v interface{}) {
	o.AmountFloat = v
}

// GetFormattedAmount returns the FormattedAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPricesPriceId200ResponseDataAttributes) GetFormattedAmount() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.FormattedAmount
}

// GetFormattedAmountOk returns a tuple with the FormattedAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPricesPriceId200ResponseDataAttributes) GetFormattedAmountOk() (*interface{}, bool) {
	if o == nil || IsNil(o.FormattedAmount) {
		return nil, false
	}
	return &o.FormattedAmount, true
}

// HasFormattedAmount returns a boolean if a field has been set.
func (o *GETPricesPriceId200ResponseDataAttributes) HasFormattedAmount() bool {
	if o != nil && IsNil(o.FormattedAmount) {
		return true
	}

	return false
}

// SetFormattedAmount gets a reference to the given interface{} and assigns it to the FormattedAmount field.
func (o *GETPricesPriceId200ResponseDataAttributes) SetFormattedAmount(v interface{}) {
	o.FormattedAmount = v
}

// GetOriginalAmountCents returns the OriginalAmountCents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPricesPriceId200ResponseDataAttributes) GetOriginalAmountCents() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.OriginalAmountCents
}

// GetOriginalAmountCentsOk returns a tuple with the OriginalAmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPricesPriceId200ResponseDataAttributes) GetOriginalAmountCentsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.OriginalAmountCents) {
		return nil, false
	}
	return &o.OriginalAmountCents, true
}

// HasOriginalAmountCents returns a boolean if a field has been set.
func (o *GETPricesPriceId200ResponseDataAttributes) HasOriginalAmountCents() bool {
	if o != nil && IsNil(o.OriginalAmountCents) {
		return true
	}

	return false
}

// SetOriginalAmountCents gets a reference to the given interface{} and assigns it to the OriginalAmountCents field.
func (o *GETPricesPriceId200ResponseDataAttributes) SetOriginalAmountCents(v interface{}) {
	o.OriginalAmountCents = v
}

// GetFormattedOriginalAmount returns the FormattedOriginalAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPricesPriceId200ResponseDataAttributes) GetFormattedOriginalAmount() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.FormattedOriginalAmount
}

// GetFormattedOriginalAmountOk returns a tuple with the FormattedOriginalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPricesPriceId200ResponseDataAttributes) GetFormattedOriginalAmountOk() (*interface{}, bool) {
	if o == nil || IsNil(o.FormattedOriginalAmount) {
		return nil, false
	}
	return &o.FormattedOriginalAmount, true
}

// HasFormattedOriginalAmount returns a boolean if a field has been set.
func (o *GETPricesPriceId200ResponseDataAttributes) HasFormattedOriginalAmount() bool {
	if o != nil && IsNil(o.FormattedOriginalAmount) {
		return true
	}

	return false
}

// SetFormattedOriginalAmount gets a reference to the given interface{} and assigns it to the FormattedOriginalAmount field.
func (o *GETPricesPriceId200ResponseDataAttributes) SetFormattedOriginalAmount(v interface{}) {
	o.FormattedOriginalAmount = v
}

// GetCompareAtAmountCents returns the CompareAtAmountCents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPricesPriceId200ResponseDataAttributes) GetCompareAtAmountCents() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CompareAtAmountCents
}

// GetCompareAtAmountCentsOk returns a tuple with the CompareAtAmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPricesPriceId200ResponseDataAttributes) GetCompareAtAmountCentsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CompareAtAmountCents) {
		return nil, false
	}
	return &o.CompareAtAmountCents, true
}

// HasCompareAtAmountCents returns a boolean if a field has been set.
func (o *GETPricesPriceId200ResponseDataAttributes) HasCompareAtAmountCents() bool {
	if o != nil && IsNil(o.CompareAtAmountCents) {
		return true
	}

	return false
}

// SetCompareAtAmountCents gets a reference to the given interface{} and assigns it to the CompareAtAmountCents field.
func (o *GETPricesPriceId200ResponseDataAttributes) SetCompareAtAmountCents(v interface{}) {
	o.CompareAtAmountCents = v
}

// GetCompareAtAmountFloat returns the CompareAtAmountFloat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPricesPriceId200ResponseDataAttributes) GetCompareAtAmountFloat() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CompareAtAmountFloat
}

// GetCompareAtAmountFloatOk returns a tuple with the CompareAtAmountFloat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPricesPriceId200ResponseDataAttributes) GetCompareAtAmountFloatOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CompareAtAmountFloat) {
		return nil, false
	}
	return &o.CompareAtAmountFloat, true
}

// HasCompareAtAmountFloat returns a boolean if a field has been set.
func (o *GETPricesPriceId200ResponseDataAttributes) HasCompareAtAmountFloat() bool {
	if o != nil && IsNil(o.CompareAtAmountFloat) {
		return true
	}

	return false
}

// SetCompareAtAmountFloat gets a reference to the given interface{} and assigns it to the CompareAtAmountFloat field.
func (o *GETPricesPriceId200ResponseDataAttributes) SetCompareAtAmountFloat(v interface{}) {
	o.CompareAtAmountFloat = v
}

// GetFormattedCompareAtAmount returns the FormattedCompareAtAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPricesPriceId200ResponseDataAttributes) GetFormattedCompareAtAmount() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.FormattedCompareAtAmount
}

// GetFormattedCompareAtAmountOk returns a tuple with the FormattedCompareAtAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPricesPriceId200ResponseDataAttributes) GetFormattedCompareAtAmountOk() (*interface{}, bool) {
	if o == nil || IsNil(o.FormattedCompareAtAmount) {
		return nil, false
	}
	return &o.FormattedCompareAtAmount, true
}

// HasFormattedCompareAtAmount returns a boolean if a field has been set.
func (o *GETPricesPriceId200ResponseDataAttributes) HasFormattedCompareAtAmount() bool {
	if o != nil && IsNil(o.FormattedCompareAtAmount) {
		return true
	}

	return false
}

// SetFormattedCompareAtAmount gets a reference to the given interface{} and assigns it to the FormattedCompareAtAmount field.
func (o *GETPricesPriceId200ResponseDataAttributes) SetFormattedCompareAtAmount(v interface{}) {
	o.FormattedCompareAtAmount = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPricesPriceId200ResponseDataAttributes) GetCreatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPricesPriceId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return &o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETPricesPriceId200ResponseDataAttributes) HasCreatedAt() bool {
	if o != nil && IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given interface{} and assigns it to the CreatedAt field.
func (o *GETPricesPriceId200ResponseDataAttributes) SetCreatedAt(v interface{}) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPricesPriceId200ResponseDataAttributes) GetUpdatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPricesPriceId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETPricesPriceId200ResponseDataAttributes) HasUpdatedAt() bool {
	if o != nil && IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given interface{} and assigns it to the UpdatedAt field.
func (o *GETPricesPriceId200ResponseDataAttributes) SetUpdatedAt(v interface{}) {
	o.UpdatedAt = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPricesPriceId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPricesPriceId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETPricesPriceId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *GETPricesPriceId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPricesPriceId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPricesPriceId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETPricesPriceId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *GETPricesPriceId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetRules returns the Rules field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPricesPriceId200ResponseDataAttributes) GetRules() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPricesPriceId200ResponseDataAttributes) GetRulesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Rules) {
		return nil, false
	}
	return &o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *GETPricesPriceId200ResponseDataAttributes) HasRules() bool {
	if o != nil && IsNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given interface{} and assigns it to the Rules field.
func (o *GETPricesPriceId200ResponseDataAttributes) SetRules(v interface{}) {
	o.Rules = v
}

// GetRuleOutcomes returns the RuleOutcomes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPricesPriceId200ResponseDataAttributes) GetRuleOutcomes() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.RuleOutcomes
}

// GetRuleOutcomesOk returns a tuple with the RuleOutcomes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPricesPriceId200ResponseDataAttributes) GetRuleOutcomesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.RuleOutcomes) {
		return nil, false
	}
	return &o.RuleOutcomes, true
}

// HasRuleOutcomes returns a boolean if a field has been set.
func (o *GETPricesPriceId200ResponseDataAttributes) HasRuleOutcomes() bool {
	if o != nil && IsNil(o.RuleOutcomes) {
		return true
	}

	return false
}

// SetRuleOutcomes gets a reference to the given interface{} and assigns it to the RuleOutcomes field.
func (o *GETPricesPriceId200ResponseDataAttributes) SetRuleOutcomes(v interface{}) {
	o.RuleOutcomes = v
}

// GetJwtCustomClaim returns the JwtCustomClaim field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPricesPriceId200ResponseDataAttributes) GetJwtCustomClaim() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.JwtCustomClaim
}

// GetJwtCustomClaimOk returns a tuple with the JwtCustomClaim field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPricesPriceId200ResponseDataAttributes) GetJwtCustomClaimOk() (*interface{}, bool) {
	if o == nil || IsNil(o.JwtCustomClaim) {
		return nil, false
	}
	return &o.JwtCustomClaim, true
}

// HasJwtCustomClaim returns a boolean if a field has been set.
func (o *GETPricesPriceId200ResponseDataAttributes) HasJwtCustomClaim() bool {
	if o != nil && IsNil(o.JwtCustomClaim) {
		return true
	}

	return false
}

// SetJwtCustomClaim gets a reference to the given interface{} and assigns it to the JwtCustomClaim field.
func (o *GETPricesPriceId200ResponseDataAttributes) SetJwtCustomClaim(v interface{}) {
	o.JwtCustomClaim = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPricesPriceId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPricesPriceId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETPricesPriceId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *GETPricesPriceId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o GETPricesPriceId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETPricesPriceId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CurrencyCode != nil {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	if o.SkuCode != nil {
		toSerialize["sku_code"] = o.SkuCode
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
	if o.OriginalAmountCents != nil {
		toSerialize["original_amount_cents"] = o.OriginalAmountCents
	}
	if o.FormattedOriginalAmount != nil {
		toSerialize["formatted_original_amount"] = o.FormattedOriginalAmount
	}
	if o.CompareAtAmountCents != nil {
		toSerialize["compare_at_amount_cents"] = o.CompareAtAmountCents
	}
	if o.CompareAtAmountFloat != nil {
		toSerialize["compare_at_amount_float"] = o.CompareAtAmountFloat
	}
	if o.FormattedCompareAtAmount != nil {
		toSerialize["formatted_compare_at_amount"] = o.FormattedCompareAtAmount
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
	if o.Rules != nil {
		toSerialize["rules"] = o.Rules
	}
	if o.RuleOutcomes != nil {
		toSerialize["rule_outcomes"] = o.RuleOutcomes
	}
	if o.JwtCustomClaim != nil {
		toSerialize["jwt_custom_claim"] = o.JwtCustomClaim
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableGETPricesPriceId200ResponseDataAttributes struct {
	value *GETPricesPriceId200ResponseDataAttributes
	isSet bool
}

func (v NullableGETPricesPriceId200ResponseDataAttributes) Get() *GETPricesPriceId200ResponseDataAttributes {
	return v.value
}

func (v *NullableGETPricesPriceId200ResponseDataAttributes) Set(val *GETPricesPriceId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETPricesPriceId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETPricesPriceId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETPricesPriceId200ResponseDataAttributes(val *GETPricesPriceId200ResponseDataAttributes) *NullableGETPricesPriceId200ResponseDataAttributes {
	return &NullableGETPricesPriceId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullableGETPricesPriceId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETPricesPriceId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
