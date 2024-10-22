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

// checks if the POSTPaymentMethods201ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTPaymentMethods201ResponseDataAttributes{}

// POSTPaymentMethods201ResponseDataAttributes struct for POSTPaymentMethods201ResponseDataAttributes
type POSTPaymentMethods201ResponseDataAttributes struct {
	// The payment method's internal name.
	Name interface{} `json:"name,omitempty"`
	// The payment source type. One of 'adyen_payments', 'axerve_payments', 'braintree_payments', 'checkout_com_payments', 'credit_cards', 'external_payments', 'klarna_payments', 'paypal_payments', 'satispay_payments', 'stripe_payments', or 'wire_transfers'.
	PaymentSourceType interface{} `json:"payment_source_type"`
	// The international 3-letter currency code as defined by the ISO 4217 standard.
	CurrencyCode interface{} `json:"currency_code,omitempty"`
	// Send this attribute if you want to mark the payment as MOTO, must be supported by payment gateway.
	Moto interface{} `json:"moto,omitempty"`
	// Send this attribute if you want to require the payment capture before fulfillment.
	RequireCapture interface{} `json:"require_capture,omitempty"`
	// Send this attribute if you want to automatically place the order upon authorization performed asynchronously.
	AutoPlace interface{} `json:"auto_place,omitempty"`
	// Send this attribute if you want to automatically capture the payment upon authorization.
	AutoCapture interface{} `json:"auto_capture,omitempty"`
	// The payment method's price, in cents.
	PriceAmountCents interface{} `json:"price_amount_cents"`
	// Send this attribute if you want to limit automatic capture to orders for which the total amount is equal or less than the specified value, in cents.
	AutoCaptureMaxAmountCents interface{} `json:"auto_capture_max_amount_cents,omitempty"`
	// Send this attribute if you want to mark this resource as disabled.
	Disable interface{} `json:"_disable,omitempty"`
	// Send this attribute if you want to mark this resource as enabled.
	Enable interface{} `json:"_enable,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPOSTPaymentMethods201ResponseDataAttributes instantiates a new POSTPaymentMethods201ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTPaymentMethods201ResponseDataAttributes(paymentSourceType interface{}, priceAmountCents interface{}) *POSTPaymentMethods201ResponseDataAttributes {
	this := POSTPaymentMethods201ResponseDataAttributes{}
	this.PaymentSourceType = paymentSourceType
	this.PriceAmountCents = priceAmountCents
	return &this
}

// NewPOSTPaymentMethods201ResponseDataAttributesWithDefaults instantiates a new POSTPaymentMethods201ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTPaymentMethods201ResponseDataAttributesWithDefaults() *POSTPaymentMethods201ResponseDataAttributes {
	this := POSTPaymentMethods201ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTPaymentMethods201ResponseDataAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPaymentMethods201ResponseDataAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *POSTPaymentMethods201ResponseDataAttributes) HasName() bool {
	if o != nil && IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given interface{} and assigns it to the Name field.
func (o *POSTPaymentMethods201ResponseDataAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetPaymentSourceType returns the PaymentSourceType field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTPaymentMethods201ResponseDataAttributes) GetPaymentSourceType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.PaymentSourceType
}

// GetPaymentSourceTypeOk returns a tuple with the PaymentSourceType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPaymentMethods201ResponseDataAttributes) GetPaymentSourceTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.PaymentSourceType) {
		return nil, false
	}
	return &o.PaymentSourceType, true
}

// SetPaymentSourceType sets field value
func (o *POSTPaymentMethods201ResponseDataAttributes) SetPaymentSourceType(v interface{}) {
	o.PaymentSourceType = v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTPaymentMethods201ResponseDataAttributes) GetCurrencyCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPaymentMethods201ResponseDataAttributes) GetCurrencyCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *POSTPaymentMethods201ResponseDataAttributes) HasCurrencyCode() bool {
	if o != nil && IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given interface{} and assigns it to the CurrencyCode field.
func (o *POSTPaymentMethods201ResponseDataAttributes) SetCurrencyCode(v interface{}) {
	o.CurrencyCode = v
}

// GetMoto returns the Moto field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTPaymentMethods201ResponseDataAttributes) GetMoto() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Moto
}

// GetMotoOk returns a tuple with the Moto field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPaymentMethods201ResponseDataAttributes) GetMotoOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Moto) {
		return nil, false
	}
	return &o.Moto, true
}

// HasMoto returns a boolean if a field has been set.
func (o *POSTPaymentMethods201ResponseDataAttributes) HasMoto() bool {
	if o != nil && IsNil(o.Moto) {
		return true
	}

	return false
}

// SetMoto gets a reference to the given interface{} and assigns it to the Moto field.
func (o *POSTPaymentMethods201ResponseDataAttributes) SetMoto(v interface{}) {
	o.Moto = v
}

// GetRequireCapture returns the RequireCapture field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTPaymentMethods201ResponseDataAttributes) GetRequireCapture() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.RequireCapture
}

// GetRequireCaptureOk returns a tuple with the RequireCapture field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPaymentMethods201ResponseDataAttributes) GetRequireCaptureOk() (*interface{}, bool) {
	if o == nil || IsNil(o.RequireCapture) {
		return nil, false
	}
	return &o.RequireCapture, true
}

// HasRequireCapture returns a boolean if a field has been set.
func (o *POSTPaymentMethods201ResponseDataAttributes) HasRequireCapture() bool {
	if o != nil && IsNil(o.RequireCapture) {
		return true
	}

	return false
}

// SetRequireCapture gets a reference to the given interface{} and assigns it to the RequireCapture field.
func (o *POSTPaymentMethods201ResponseDataAttributes) SetRequireCapture(v interface{}) {
	o.RequireCapture = v
}

// GetAutoPlace returns the AutoPlace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTPaymentMethods201ResponseDataAttributes) GetAutoPlace() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.AutoPlace
}

// GetAutoPlaceOk returns a tuple with the AutoPlace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPaymentMethods201ResponseDataAttributes) GetAutoPlaceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.AutoPlace) {
		return nil, false
	}
	return &o.AutoPlace, true
}

// HasAutoPlace returns a boolean if a field has been set.
func (o *POSTPaymentMethods201ResponseDataAttributes) HasAutoPlace() bool {
	if o != nil && IsNil(o.AutoPlace) {
		return true
	}

	return false
}

// SetAutoPlace gets a reference to the given interface{} and assigns it to the AutoPlace field.
func (o *POSTPaymentMethods201ResponseDataAttributes) SetAutoPlace(v interface{}) {
	o.AutoPlace = v
}

// GetAutoCapture returns the AutoCapture field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTPaymentMethods201ResponseDataAttributes) GetAutoCapture() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.AutoCapture
}

// GetAutoCaptureOk returns a tuple with the AutoCapture field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPaymentMethods201ResponseDataAttributes) GetAutoCaptureOk() (*interface{}, bool) {
	if o == nil || IsNil(o.AutoCapture) {
		return nil, false
	}
	return &o.AutoCapture, true
}

// HasAutoCapture returns a boolean if a field has been set.
func (o *POSTPaymentMethods201ResponseDataAttributes) HasAutoCapture() bool {
	if o != nil && IsNil(o.AutoCapture) {
		return true
	}

	return false
}

// SetAutoCapture gets a reference to the given interface{} and assigns it to the AutoCapture field.
func (o *POSTPaymentMethods201ResponseDataAttributes) SetAutoCapture(v interface{}) {
	o.AutoCapture = v
}

// GetPriceAmountCents returns the PriceAmountCents field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTPaymentMethods201ResponseDataAttributes) GetPriceAmountCents() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.PriceAmountCents
}

// GetPriceAmountCentsOk returns a tuple with the PriceAmountCents field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPaymentMethods201ResponseDataAttributes) GetPriceAmountCentsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.PriceAmountCents) {
		return nil, false
	}
	return &o.PriceAmountCents, true
}

// SetPriceAmountCents sets field value
func (o *POSTPaymentMethods201ResponseDataAttributes) SetPriceAmountCents(v interface{}) {
	o.PriceAmountCents = v
}

// GetAutoCaptureMaxAmountCents returns the AutoCaptureMaxAmountCents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTPaymentMethods201ResponseDataAttributes) GetAutoCaptureMaxAmountCents() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.AutoCaptureMaxAmountCents
}

// GetAutoCaptureMaxAmountCentsOk returns a tuple with the AutoCaptureMaxAmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPaymentMethods201ResponseDataAttributes) GetAutoCaptureMaxAmountCentsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.AutoCaptureMaxAmountCents) {
		return nil, false
	}
	return &o.AutoCaptureMaxAmountCents, true
}

// HasAutoCaptureMaxAmountCents returns a boolean if a field has been set.
func (o *POSTPaymentMethods201ResponseDataAttributes) HasAutoCaptureMaxAmountCents() bool {
	if o != nil && IsNil(o.AutoCaptureMaxAmountCents) {
		return true
	}

	return false
}

// SetAutoCaptureMaxAmountCents gets a reference to the given interface{} and assigns it to the AutoCaptureMaxAmountCents field.
func (o *POSTPaymentMethods201ResponseDataAttributes) SetAutoCaptureMaxAmountCents(v interface{}) {
	o.AutoCaptureMaxAmountCents = v
}

// GetDisable returns the Disable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTPaymentMethods201ResponseDataAttributes) GetDisable() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Disable
}

// GetDisableOk returns a tuple with the Disable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPaymentMethods201ResponseDataAttributes) GetDisableOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Disable) {
		return nil, false
	}
	return &o.Disable, true
}

// HasDisable returns a boolean if a field has been set.
func (o *POSTPaymentMethods201ResponseDataAttributes) HasDisable() bool {
	if o != nil && IsNil(o.Disable) {
		return true
	}

	return false
}

// SetDisable gets a reference to the given interface{} and assigns it to the Disable field.
func (o *POSTPaymentMethods201ResponseDataAttributes) SetDisable(v interface{}) {
	o.Disable = v
}

// GetEnable returns the Enable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTPaymentMethods201ResponseDataAttributes) GetEnable() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPaymentMethods201ResponseDataAttributes) GetEnableOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Enable) {
		return nil, false
	}
	return &o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *POSTPaymentMethods201ResponseDataAttributes) HasEnable() bool {
	if o != nil && IsNil(o.Enable) {
		return true
	}

	return false
}

// SetEnable gets a reference to the given interface{} and assigns it to the Enable field.
func (o *POSTPaymentMethods201ResponseDataAttributes) SetEnable(v interface{}) {
	o.Enable = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTPaymentMethods201ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPaymentMethods201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTPaymentMethods201ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *POSTPaymentMethods201ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTPaymentMethods201ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPaymentMethods201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTPaymentMethods201ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *POSTPaymentMethods201ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTPaymentMethods201ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPaymentMethods201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTPaymentMethods201ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *POSTPaymentMethods201ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o POSTPaymentMethods201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTPaymentMethods201ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
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
	if o.AutoPlace != nil {
		toSerialize["auto_place"] = o.AutoPlace
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
	if o.Disable != nil {
		toSerialize["_disable"] = o.Disable
	}
	if o.Enable != nil {
		toSerialize["_enable"] = o.Enable
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

type NullablePOSTPaymentMethods201ResponseDataAttributes struct {
	value *POSTPaymentMethods201ResponseDataAttributes
	isSet bool
}

func (v NullablePOSTPaymentMethods201ResponseDataAttributes) Get() *POSTPaymentMethods201ResponseDataAttributes {
	return v.value
}

func (v *NullablePOSTPaymentMethods201ResponseDataAttributes) Set(val *POSTPaymentMethods201ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTPaymentMethods201ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTPaymentMethods201ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTPaymentMethods201ResponseDataAttributes(val *POSTPaymentMethods201ResponseDataAttributes) *NullablePOSTPaymentMethods201ResponseDataAttributes {
	return &NullablePOSTPaymentMethods201ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTPaymentMethods201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTPaymentMethods201ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
