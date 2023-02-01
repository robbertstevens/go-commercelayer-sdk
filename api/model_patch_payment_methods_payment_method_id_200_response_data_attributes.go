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

// PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes struct for PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes
type PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes struct {
	// The payment source type, can be one of: 'AdyenPayment', 'BraintreePayment', 'CheckoutComPayment', 'CreditCard', 'ExternalPayment', 'KlarnaPayment', 'PaypalPayment', 'StripePayment', or 'WireTransfer'.
	PaymentSourceType *string `json:"payment_source_type,omitempty"`
	// The international 3-letter currency code as defined by the ISO 4217 standard.
	CurrencyCode *string `json:"currency_code,omitempty"`
	// Send this attribute if you want to mark the payment as MOTO, must be supported by payment gateway.
	Moto *bool `json:"moto,omitempty"`
	// Send this attribute if you want to mark the payment method as disabled.
	Disable *bool `json:"_disable,omitempty"`
	// Send this attribute if you want to mark the payment method as enabled.
	Enable *bool `json:"_enable,omitempty"`
	// The payment method's price, in cents
	PriceAmountCents *int32 `json:"price_amount_cents,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewPATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes instantiates a new PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes() *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes {
	this := PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes{}
	return &this
}

// NewPATCHPaymentMethodsPaymentMethodId200ResponseDataAttributesWithDefaults instantiates a new PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHPaymentMethodsPaymentMethodId200ResponseDataAttributesWithDefaults() *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes {
	this := PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes{}
	return &this
}

// GetPaymentSourceType returns the PaymentSourceType field value if set, zero value otherwise.
func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetPaymentSourceType() string {
	if o == nil || o.PaymentSourceType == nil {
		var ret string
		return ret
	}
	return *o.PaymentSourceType
}

// GetPaymentSourceTypeOk returns a tuple with the PaymentSourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetPaymentSourceTypeOk() (*string, bool) {
	if o == nil || o.PaymentSourceType == nil {
		return nil, false
	}
	return o.PaymentSourceType, true
}

// HasPaymentSourceType returns a boolean if a field has been set.
func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) HasPaymentSourceType() bool {
	if o != nil && o.PaymentSourceType != nil {
		return true
	}

	return false
}

// SetPaymentSourceType gets a reference to the given string and assigns it to the PaymentSourceType field.
func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) SetPaymentSourceType(v string) {
	o.PaymentSourceType = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetCurrencyCode() string {
	if o == nil || o.CurrencyCode == nil {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || o.CurrencyCode == nil {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) HasCurrencyCode() bool {
	if o != nil && o.CurrencyCode != nil {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetMoto returns the Moto field value if set, zero value otherwise.
func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetMoto() bool {
	if o == nil || o.Moto == nil {
		var ret bool
		return ret
	}
	return *o.Moto
}

// GetMotoOk returns a tuple with the Moto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetMotoOk() (*bool, bool) {
	if o == nil || o.Moto == nil {
		return nil, false
	}
	return o.Moto, true
}

// HasMoto returns a boolean if a field has been set.
func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) HasMoto() bool {
	if o != nil && o.Moto != nil {
		return true
	}

	return false
}

// SetMoto gets a reference to the given bool and assigns it to the Moto field.
func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) SetMoto(v bool) {
	o.Moto = &v
}

// GetDisable returns the Disable field value if set, zero value otherwise.
func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetDisable() bool {
	if o == nil || o.Disable == nil {
		var ret bool
		return ret
	}
	return *o.Disable
}

// GetDisableOk returns a tuple with the Disable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetDisableOk() (*bool, bool) {
	if o == nil || o.Disable == nil {
		return nil, false
	}
	return o.Disable, true
}

// HasDisable returns a boolean if a field has been set.
func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) HasDisable() bool {
	if o != nil && o.Disable != nil {
		return true
	}

	return false
}

// SetDisable gets a reference to the given bool and assigns it to the Disable field.
func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) SetDisable(v bool) {
	o.Disable = &v
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetEnable() bool {
	if o == nil || o.Enable == nil {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetEnableOk() (*bool, bool) {
	if o == nil || o.Enable == nil {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) HasEnable() bool {
	if o != nil && o.Enable != nil {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) SetEnable(v bool) {
	o.Enable = &v
}

// GetPriceAmountCents returns the PriceAmountCents field value if set, zero value otherwise.
func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetPriceAmountCents() int32 {
	if o == nil || o.PriceAmountCents == nil {
		var ret int32
		return ret
	}
	return *o.PriceAmountCents
}

// GetPriceAmountCentsOk returns a tuple with the PriceAmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetPriceAmountCentsOk() (*int32, bool) {
	if o == nil || o.PriceAmountCents == nil {
		return nil, false
	}
	return o.PriceAmountCents, true
}

// HasPriceAmountCents returns a boolean if a field has been set.
func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) HasPriceAmountCents() bool {
	if o != nil && o.PriceAmountCents != nil {
		return true
	}

	return false
}

// SetPriceAmountCents gets a reference to the given int32 and assigns it to the PriceAmountCents field.
func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) SetPriceAmountCents(v int32) {
	o.PriceAmountCents = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PaymentSourceType != nil {
		toSerialize["payment_source_type"] = o.PaymentSourceType
	}
	if o.CurrencyCode != nil {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	if o.Moto != nil {
		toSerialize["moto"] = o.Moto
	}
	if o.Disable != nil {
		toSerialize["_disable"] = o.Disable
	}
	if o.Enable != nil {
		toSerialize["_enable"] = o.Enable
	}
	if o.PriceAmountCents != nil {
		toSerialize["price_amount_cents"] = o.PriceAmountCents
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

type NullablePATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes struct {
	value *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes
	isSet bool
}

func (v NullablePATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) Get() *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes {
	return v.value
}

func (v *NullablePATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) Set(val *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes(val *PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) *NullablePATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes {
	return &NullablePATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
