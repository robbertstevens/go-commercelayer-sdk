/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.4.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PATCHGiftCardsGiftCardId200ResponseDataAttributes struct for PATCHGiftCardsGiftCardId200ResponseDataAttributes
type PATCHGiftCardsGiftCardId200ResponseDataAttributes struct {
	// The international 3-letter currency code as defined by the ISO 4217 standard.
	CurrencyCode *string `json:"currency_code,omitempty"`
	// The gift card balance, in cents.
	BalanceCents *int32 `json:"balance_cents,omitempty"`
	// The gift card balance max, in cents.
	BalanceMaxCents *string `json:"balance_max_cents,omitempty"`
	// Indicates if the gift card can be used only one.
	SingleUse *bool `json:"single_use,omitempty"`
	// Indicates if the gift card can be recharged.
	Rechargeable *bool `json:"rechargeable,omitempty"`
	// The URL of an image that represents the gift card.
	ImageUrl *string `json:"image_url,omitempty"`
	// Time at which the gift card will expire.
	ExpiresAt *string `json:"expires_at,omitempty"`
	// The email address of the associated recipient. When creating or updating a gift card, this is a shortcut to find or create the associated recipient by email.
	RecipientEmail *string `json:"recipient_email,omitempty"`
	// Send this attribute if you want to confirm a draft gift card. The gift card becomes 'inactive', waiting to be activated.
	Purchase *bool `json:"_purchase,omitempty"`
	// Send this attribute if you want to activate a gift card.
	Activate *bool `json:"_activate,omitempty"`
	// Send this attribute if you want to deactivate a gift card.
	Deactivate *bool `json:"_deactivate,omitempty"`
	// The balance change, in cents. Send a negative value to reduces the card balance by the specified amount. Send a positive value to recharge the gift card (if rechargeable).
	BalanceChangeCents *int32 `json:"_balance_change_cents,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewPATCHGiftCardsGiftCardId200ResponseDataAttributes instantiates a new PATCHGiftCardsGiftCardId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHGiftCardsGiftCardId200ResponseDataAttributes() *PATCHGiftCardsGiftCardId200ResponseDataAttributes {
	this := PATCHGiftCardsGiftCardId200ResponseDataAttributes{}
	return &this
}

// NewPATCHGiftCardsGiftCardId200ResponseDataAttributesWithDefaults instantiates a new PATCHGiftCardsGiftCardId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHGiftCardsGiftCardId200ResponseDataAttributesWithDefaults() *PATCHGiftCardsGiftCardId200ResponseDataAttributes {
	this := PATCHGiftCardsGiftCardId200ResponseDataAttributes{}
	return &this
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetCurrencyCode() string {
	if o == nil || o.CurrencyCode == nil {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || o.CurrencyCode == nil {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) HasCurrencyCode() bool {
	if o != nil && o.CurrencyCode != nil {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetBalanceCents returns the BalanceCents field value if set, zero value otherwise.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetBalanceCents() int32 {
	if o == nil || o.BalanceCents == nil {
		var ret int32
		return ret
	}
	return *o.BalanceCents
}

// GetBalanceCentsOk returns a tuple with the BalanceCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetBalanceCentsOk() (*int32, bool) {
	if o == nil || o.BalanceCents == nil {
		return nil, false
	}
	return o.BalanceCents, true
}

// HasBalanceCents returns a boolean if a field has been set.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) HasBalanceCents() bool {
	if o != nil && o.BalanceCents != nil {
		return true
	}

	return false
}

// SetBalanceCents gets a reference to the given int32 and assigns it to the BalanceCents field.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetBalanceCents(v int32) {
	o.BalanceCents = &v
}

// GetBalanceMaxCents returns the BalanceMaxCents field value if set, zero value otherwise.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetBalanceMaxCents() string {
	if o == nil || o.BalanceMaxCents == nil {
		var ret string
		return ret
	}
	return *o.BalanceMaxCents
}

// GetBalanceMaxCentsOk returns a tuple with the BalanceMaxCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetBalanceMaxCentsOk() (*string, bool) {
	if o == nil || o.BalanceMaxCents == nil {
		return nil, false
	}
	return o.BalanceMaxCents, true
}

// HasBalanceMaxCents returns a boolean if a field has been set.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) HasBalanceMaxCents() bool {
	if o != nil && o.BalanceMaxCents != nil {
		return true
	}

	return false
}

// SetBalanceMaxCents gets a reference to the given string and assigns it to the BalanceMaxCents field.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetBalanceMaxCents(v string) {
	o.BalanceMaxCents = &v
}

// GetSingleUse returns the SingleUse field value if set, zero value otherwise.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetSingleUse() bool {
	if o == nil || o.SingleUse == nil {
		var ret bool
		return ret
	}
	return *o.SingleUse
}

// GetSingleUseOk returns a tuple with the SingleUse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetSingleUseOk() (*bool, bool) {
	if o == nil || o.SingleUse == nil {
		return nil, false
	}
	return o.SingleUse, true
}

// HasSingleUse returns a boolean if a field has been set.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) HasSingleUse() bool {
	if o != nil && o.SingleUse != nil {
		return true
	}

	return false
}

// SetSingleUse gets a reference to the given bool and assigns it to the SingleUse field.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetSingleUse(v bool) {
	o.SingleUse = &v
}

// GetRechargeable returns the Rechargeable field value if set, zero value otherwise.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetRechargeable() bool {
	if o == nil || o.Rechargeable == nil {
		var ret bool
		return ret
	}
	return *o.Rechargeable
}

// GetRechargeableOk returns a tuple with the Rechargeable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetRechargeableOk() (*bool, bool) {
	if o == nil || o.Rechargeable == nil {
		return nil, false
	}
	return o.Rechargeable, true
}

// HasRechargeable returns a boolean if a field has been set.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) HasRechargeable() bool {
	if o != nil && o.Rechargeable != nil {
		return true
	}

	return false
}

// SetRechargeable gets a reference to the given bool and assigns it to the Rechargeable field.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetRechargeable(v bool) {
	o.Rechargeable = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetImageUrl() string {
	if o == nil || o.ImageUrl == nil {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetImageUrlOk() (*string, bool) {
	if o == nil || o.ImageUrl == nil {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) HasImageUrl() bool {
	if o != nil && o.ImageUrl != nil {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetExpiresAt() string {
	if o == nil || o.ExpiresAt == nil {
		var ret string
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetExpiresAtOk() (*string, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt != nil {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given string and assigns it to the ExpiresAt field.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetExpiresAt(v string) {
	o.ExpiresAt = &v
}

// GetRecipientEmail returns the RecipientEmail field value if set, zero value otherwise.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetRecipientEmail() string {
	if o == nil || o.RecipientEmail == nil {
		var ret string
		return ret
	}
	return *o.RecipientEmail
}

// GetRecipientEmailOk returns a tuple with the RecipientEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetRecipientEmailOk() (*string, bool) {
	if o == nil || o.RecipientEmail == nil {
		return nil, false
	}
	return o.RecipientEmail, true
}

// HasRecipientEmail returns a boolean if a field has been set.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) HasRecipientEmail() bool {
	if o != nil && o.RecipientEmail != nil {
		return true
	}

	return false
}

// SetRecipientEmail gets a reference to the given string and assigns it to the RecipientEmail field.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetRecipientEmail(v string) {
	o.RecipientEmail = &v
}

// GetPurchase returns the Purchase field value if set, zero value otherwise.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetPurchase() bool {
	if o == nil || o.Purchase == nil {
		var ret bool
		return ret
	}
	return *o.Purchase
}

// GetPurchaseOk returns a tuple with the Purchase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetPurchaseOk() (*bool, bool) {
	if o == nil || o.Purchase == nil {
		return nil, false
	}
	return o.Purchase, true
}

// HasPurchase returns a boolean if a field has been set.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) HasPurchase() bool {
	if o != nil && o.Purchase != nil {
		return true
	}

	return false
}

// SetPurchase gets a reference to the given bool and assigns it to the Purchase field.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetPurchase(v bool) {
	o.Purchase = &v
}

// GetActivate returns the Activate field value if set, zero value otherwise.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetActivate() bool {
	if o == nil || o.Activate == nil {
		var ret bool
		return ret
	}
	return *o.Activate
}

// GetActivateOk returns a tuple with the Activate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetActivateOk() (*bool, bool) {
	if o == nil || o.Activate == nil {
		return nil, false
	}
	return o.Activate, true
}

// HasActivate returns a boolean if a field has been set.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) HasActivate() bool {
	if o != nil && o.Activate != nil {
		return true
	}

	return false
}

// SetActivate gets a reference to the given bool and assigns it to the Activate field.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetActivate(v bool) {
	o.Activate = &v
}

// GetDeactivate returns the Deactivate field value if set, zero value otherwise.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetDeactivate() bool {
	if o == nil || o.Deactivate == nil {
		var ret bool
		return ret
	}
	return *o.Deactivate
}

// GetDeactivateOk returns a tuple with the Deactivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetDeactivateOk() (*bool, bool) {
	if o == nil || o.Deactivate == nil {
		return nil, false
	}
	return o.Deactivate, true
}

// HasDeactivate returns a boolean if a field has been set.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) HasDeactivate() bool {
	if o != nil && o.Deactivate != nil {
		return true
	}

	return false
}

// SetDeactivate gets a reference to the given bool and assigns it to the Deactivate field.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetDeactivate(v bool) {
	o.Deactivate = &v
}

// GetBalanceChangeCents returns the BalanceChangeCents field value if set, zero value otherwise.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetBalanceChangeCents() int32 {
	if o == nil || o.BalanceChangeCents == nil {
		var ret int32
		return ret
	}
	return *o.BalanceChangeCents
}

// GetBalanceChangeCentsOk returns a tuple with the BalanceChangeCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetBalanceChangeCentsOk() (*int32, bool) {
	if o == nil || o.BalanceChangeCents == nil {
		return nil, false
	}
	return o.BalanceChangeCents, true
}

// HasBalanceChangeCents returns a boolean if a field has been set.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) HasBalanceChangeCents() bool {
	if o != nil && o.BalanceChangeCents != nil {
		return true
	}

	return false
}

// SetBalanceChangeCents gets a reference to the given int32 and assigns it to the BalanceChangeCents field.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetBalanceChangeCents(v int32) {
	o.BalanceChangeCents = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PATCHGiftCardsGiftCardId200ResponseDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o PATCHGiftCardsGiftCardId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CurrencyCode != nil {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	if o.BalanceCents != nil {
		toSerialize["balance_cents"] = o.BalanceCents
	}
	if o.BalanceMaxCents != nil {
		toSerialize["balance_max_cents"] = o.BalanceMaxCents
	}
	if o.SingleUse != nil {
		toSerialize["single_use"] = o.SingleUse
	}
	if o.Rechargeable != nil {
		toSerialize["rechargeable"] = o.Rechargeable
	}
	if o.ImageUrl != nil {
		toSerialize["image_url"] = o.ImageUrl
	}
	if o.ExpiresAt != nil {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if o.RecipientEmail != nil {
		toSerialize["recipient_email"] = o.RecipientEmail
	}
	if o.Purchase != nil {
		toSerialize["_purchase"] = o.Purchase
	}
	if o.Activate != nil {
		toSerialize["_activate"] = o.Activate
	}
	if o.Deactivate != nil {
		toSerialize["_deactivate"] = o.Deactivate
	}
	if o.BalanceChangeCents != nil {
		toSerialize["_balance_change_cents"] = o.BalanceChangeCents
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

type NullablePATCHGiftCardsGiftCardId200ResponseDataAttributes struct {
	value *PATCHGiftCardsGiftCardId200ResponseDataAttributes
	isSet bool
}

func (v NullablePATCHGiftCardsGiftCardId200ResponseDataAttributes) Get() *PATCHGiftCardsGiftCardId200ResponseDataAttributes {
	return v.value
}

func (v *NullablePATCHGiftCardsGiftCardId200ResponseDataAttributes) Set(val *PATCHGiftCardsGiftCardId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHGiftCardsGiftCardId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHGiftCardsGiftCardId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHGiftCardsGiftCardId200ResponseDataAttributes(val *PATCHGiftCardsGiftCardId200ResponseDataAttributes) *NullablePATCHGiftCardsGiftCardId200ResponseDataAttributes {
	return &NullablePATCHGiftCardsGiftCardId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHGiftCardsGiftCardId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHGiftCardsGiftCardId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
