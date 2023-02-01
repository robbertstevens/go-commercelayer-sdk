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

// PaymentMethodCreateDataRelationships struct for PaymentMethodCreateDataRelationships
type PaymentMethodCreateDataRelationships struct {
	Market         *BillingInfoValidationRuleCreateDataRelationshipsMarket `json:"market,omitempty"`
	PaymentGateway PaymentMethodCreateDataRelationshipsPaymentGateway      `json:"payment_gateway"`
}

// NewPaymentMethodCreateDataRelationships instantiates a new PaymentMethodCreateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodCreateDataRelationships(paymentGateway PaymentMethodCreateDataRelationshipsPaymentGateway) *PaymentMethodCreateDataRelationships {
	this := PaymentMethodCreateDataRelationships{}
	this.PaymentGateway = paymentGateway
	return &this
}

// NewPaymentMethodCreateDataRelationshipsWithDefaults instantiates a new PaymentMethodCreateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodCreateDataRelationshipsWithDefaults() *PaymentMethodCreateDataRelationships {
	this := PaymentMethodCreateDataRelationships{}
	return &this
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *PaymentMethodCreateDataRelationships) GetMarket() BillingInfoValidationRuleCreateDataRelationshipsMarket {
	if o == nil || o.Market == nil {
		var ret BillingInfoValidationRuleCreateDataRelationshipsMarket
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCreateDataRelationships) GetMarketOk() (*BillingInfoValidationRuleCreateDataRelationshipsMarket, bool) {
	if o == nil || o.Market == nil {
		return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *PaymentMethodCreateDataRelationships) HasMarket() bool {
	if o != nil && o.Market != nil {
		return true
	}

	return false
}

// SetMarket gets a reference to the given BillingInfoValidationRuleCreateDataRelationshipsMarket and assigns it to the Market field.
func (o *PaymentMethodCreateDataRelationships) SetMarket(v BillingInfoValidationRuleCreateDataRelationshipsMarket) {
	o.Market = &v
}

// GetPaymentGateway returns the PaymentGateway field value
func (o *PaymentMethodCreateDataRelationships) GetPaymentGateway() PaymentMethodCreateDataRelationshipsPaymentGateway {
	if o == nil {
		var ret PaymentMethodCreateDataRelationshipsPaymentGateway
		return ret
	}

	return o.PaymentGateway
}

// GetPaymentGatewayOk returns a tuple with the PaymentGateway field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodCreateDataRelationships) GetPaymentGatewayOk() (*PaymentMethodCreateDataRelationshipsPaymentGateway, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentGateway, true
}

// SetPaymentGateway sets field value
func (o *PaymentMethodCreateDataRelationships) SetPaymentGateway(v PaymentMethodCreateDataRelationshipsPaymentGateway) {
	o.PaymentGateway = v
}

func (o PaymentMethodCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Market != nil {
		toSerialize["market"] = o.Market
	}
	if true {
		toSerialize["payment_gateway"] = o.PaymentGateway
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentMethodCreateDataRelationships struct {
	value *PaymentMethodCreateDataRelationships
	isSet bool
}

func (v NullablePaymentMethodCreateDataRelationships) Get() *PaymentMethodCreateDataRelationships {
	return v.value
}

func (v *NullablePaymentMethodCreateDataRelationships) Set(val *PaymentMethodCreateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodCreateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodCreateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodCreateDataRelationships(val *PaymentMethodCreateDataRelationships) *NullablePaymentMethodCreateDataRelationships {
	return &NullablePaymentMethodCreateDataRelationships{value: val, isSet: true}
}

func (v NullablePaymentMethodCreateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodCreateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
