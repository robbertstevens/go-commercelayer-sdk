/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 4.1.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the POSTPaymentMethodsRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTPaymentMethodsRequestDataAttributes{}

// POSTPaymentMethodsRequestDataAttributes struct for POSTPaymentMethodsRequestDataAttributes
type POSTPaymentMethodsRequestDataAttributes struct {
	// The payment source type, can be one of: 'AdyenPayment', 'AxervePayment', 'BraintreePayment', 'CheckoutComPayment', 'CreditCard', 'ExternalPayment', 'KlarnaPayment', 'PaypalPayment', 'SatispayPayment', 'StripePayment', or 'WireTransfer'.
	PaymentSourceType interface{} `json:"payment_source_type"`
	// The international 3-letter currency code as defined by the ISO 4217 standard.
	CurrencyCode interface{} `json:"currency_code,omitempty"`
	// Send this attribute if you want to mark the payment as MOTO, must be supported by payment gateway.
	Moto interface{} `json:"moto,omitempty"`
	// Send this attribute if you want to require the payment capture before fulfillment.
	RequireCapture interface{} `json:"require_capture,omitempty"`
	// Send this attribute if you want to automatically capture the payment upon authorization.
	AutoCapture interface{} `json:"auto_capture,omitempty"`
	// The payment method's price, in cents
	PriceAmountCents interface{} `json:"price_amount_cents"`
	// Send this attribute if you want to limit automatic capture to orders for which the total amount is equal or less than the specified value, in cents.
	AutoCaptureMaxAmountCents interface{} `json:"auto_capture_max_amount_cents,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPOSTPaymentMethodsRequestDataAttributes instantiates a new POSTPaymentMethodsRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTPaymentMethodsRequestDataAttributes(paymentSourceType interface{}, priceAmountCents interface{}) *POSTPaymentMethodsRequestDataAttributes {
	this := POSTPaymentMethodsRequestDataAttributes{}
	this.PaymentSourceType = paymentSourceType
	this.PriceAmountCents = priceAmountCents
	return &this
}

// NewPOSTPaymentMethodsRequestDataAttributesWithDefaults instantiates a new POSTPaymentMethodsRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTPaymentMethodsRequestDataAttributesWithDefaults() *POSTPaymentMethodsRequestDataAttributes {
	this := POSTPaymentMethodsRequestDataAttributes{}
	return &this
}

// GetPaymentSourceType returns the PaymentSourceType field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTPaymentMethodsRequestDataAttributes) GetPaymentSourceType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.PaymentSourceType
}

// GetPaymentSourceTypeOk returns a tuple with the PaymentSourceType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPaymentMethodsRequestDataAttributes) GetPaymentSourceTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.PaymentSourceType) {
		return nil, false
	}
	return &o.PaymentSourceType, true
}

// SetPaymentSourceType sets field value
func (o *POSTPaymentMethodsRequestDataAttributes) SetPaymentSourceType(v interface{}) {
	o.PaymentSourceType = v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTPaymentMethodsRequestDataAttributes) GetCurrencyCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPaymentMethodsRequestDataAttributes) GetCurrencyCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *POSTPaymentMethodsRequestDataAttributes) HasCurrencyCode() bool {
	if o != nil && IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given interface{} and assigns it to the CurrencyCode field.
func (o *POSTPaymentMethodsRequestDataAttributes) SetCurrencyCode(v interface{}) {
	o.CurrencyCode = v
}

// GetMoto returns the Moto field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTPaymentMethodsRequestDataAttributes) GetMoto() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Moto
}

// GetMotoOk returns a tuple with the Moto field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPaymentMethodsRequestDataAttributes) GetMotoOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Moto) {
		return nil, false
	}
	return &o.Moto, true
}

// HasMoto returns a boolean if a field has been set.
func (o *POSTPaymentMethodsRequestDataAttributes) HasMoto() bool {
	if o != nil && IsNil(o.Moto) {
		return true
	}

	return false
}

// SetMoto gets a reference to the given interface{} and assigns it to the Moto field.
func (o *POSTPaymentMethodsRequestDataAttributes) SetMoto(v interface{}) {
	o.Moto = v
}

// GetRequireCapture returns the RequireCapture field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTPaymentMethodsRequestDataAttributes) GetRequireCapture() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.RequireCapture
}

// GetRequireCaptureOk returns a tuple with the RequireCapture field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPaymentMethodsRequestDataAttributes) GetRequireCaptureOk() (*interface{}, bool) {
	if o == nil || IsNil(o.RequireCapture) {
		return nil, false
	}
	return &o.RequireCapture, true
}

// HasRequireCapture returns a boolean if a field has been set.
func (o *POSTPaymentMethodsRequestDataAttributes) HasRequireCapture() bool {
	if o != nil && IsNil(o.RequireCapture) {
		return true
	}

	return false
}

// SetRequireCapture gets a reference to the given interface{} and assigns it to the RequireCapture field.
func (o *POSTPaymentMethodsRequestDataAttributes) SetRequireCapture(v interface{}) {
	o.RequireCapture = v
}

// GetAutoCapture returns the AutoCapture field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTPaymentMethodsRequestDataAttributes) GetAutoCapture() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.AutoCapture
}

// GetAutoCaptureOk returns a tuple with the AutoCapture field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPaymentMethodsRequestDataAttributes) GetAutoCaptureOk() (*interface{}, bool) {
	if o == nil || IsNil(o.AutoCapture) {
		return nil, false
	}
	return &o.AutoCapture, true
}

// HasAutoCapture returns a boolean if a field has been set.
func (o *POSTPaymentMethodsRequestDataAttributes) HasAutoCapture() bool {
	if o != nil && IsNil(o.AutoCapture) {
		return true
	}

	return false
}

// SetAutoCapture gets a reference to the given interface{} and assigns it to the AutoCapture field.
func (o *POSTPaymentMethodsRequestDataAttributes) SetAutoCapture(v interface{}) {
	o.AutoCapture = v
}

// GetPriceAmountCents returns the PriceAmountCents field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTPaymentMethodsRequestDataAttributes) GetPriceAmountCents() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.PriceAmountCents
}

// GetPriceAmountCentsOk returns a tuple with the PriceAmountCents field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPaymentMethodsRequestDataAttributes) GetPriceAmountCentsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.PriceAmountCents) {
		return nil, false
	}
	return &o.PriceAmountCents, true
}

// SetPriceAmountCents sets field value
func (o *POSTPaymentMethodsRequestDataAttributes) SetPriceAmountCents(v interface{}) {
	o.PriceAmountCents = v
}

// GetAutoCaptureMaxAmountCents returns the AutoCaptureMaxAmountCents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTPaymentMethodsRequestDataAttributes) GetAutoCaptureMaxAmountCents() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.AutoCaptureMaxAmountCents
}

// GetAutoCaptureMaxAmountCentsOk returns a tuple with the AutoCaptureMaxAmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPaymentMethodsRequestDataAttributes) GetAutoCaptureMaxAmountCentsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.AutoCaptureMaxAmountCents) {
		return nil, false
	}
	return &o.AutoCaptureMaxAmountCents, true
}

// HasAutoCaptureMaxAmountCents returns a boolean if a field has been set.
func (o *POSTPaymentMethodsRequestDataAttributes) HasAutoCaptureMaxAmountCents() bool {
	if o != nil && IsNil(o.AutoCaptureMaxAmountCents) {
		return true
	}

	return false
}

// SetAutoCaptureMaxAmountCents gets a reference to the given interface{} and assigns it to the AutoCaptureMaxAmountCents field.
func (o *POSTPaymentMethodsRequestDataAttributes) SetAutoCaptureMaxAmountCents(v interface{}) {
	o.AutoCaptureMaxAmountCents = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTPaymentMethodsRequestDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPaymentMethodsRequestDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTPaymentMethodsRequestDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *POSTPaymentMethodsRequestDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTPaymentMethodsRequestDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPaymentMethodsRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTPaymentMethodsRequestDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *POSTPaymentMethodsRequestDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTPaymentMethodsRequestDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPaymentMethodsRequestDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTPaymentMethodsRequestDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *POSTPaymentMethodsRequestDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o POSTPaymentMethodsRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTPaymentMethodsRequestDataAttributes) ToMap() (map[string]interface{}, error) {
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
	if o.RequireCapture != nil {
		toSerialize["require_capture"] = o.RequireCapture
	}
	if o.AutoCapture != nil {
		toSerialize["auto_capture"] = o.AutoCapture
	}
	if o.PriceAmountCents != nil {
		toSerialize["price_amount_cents"] = o.PriceAmountCents
	}
	if o.AutoCaptureMaxAmountCents != nil {
		toSerialize["auto_capture_max_amount_cents"] = o.AutoCaptureMaxAmountCents
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

type NullablePOSTPaymentMethodsRequestDataAttributes struct {
	value *POSTPaymentMethodsRequestDataAttributes
	isSet bool
}

func (v NullablePOSTPaymentMethodsRequestDataAttributes) Get() *POSTPaymentMethodsRequestDataAttributes {
	return v.value
}

func (v *NullablePOSTPaymentMethodsRequestDataAttributes) Set(val *POSTPaymentMethodsRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTPaymentMethodsRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTPaymentMethodsRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTPaymentMethodsRequestDataAttributes(val *POSTPaymentMethodsRequestDataAttributes) *NullablePOSTPaymentMethodsRequestDataAttributes {
	return &NullablePOSTPaymentMethodsRequestDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTPaymentMethodsRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTPaymentMethodsRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
