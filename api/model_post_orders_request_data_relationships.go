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

// checks if the POSTOrdersRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTOrdersRequestDataRelationships{}

// POSTOrdersRequestDataRelationships struct for POSTOrdersRequestDataRelationships
type POSTOrdersRequestDataRelationships struct {
	Market          *POSTBillingInfoValidationRulesRequestDataRelationshipsMarket    `json:"market,omitempty"`
	Customer        *POSTCouponRecipientsRequestDataRelationshipsCustomer            `json:"customer,omitempty"`
	ShippingAddress *POSTCustomerAddressesRequestDataRelationshipsAddress            `json:"shipping_address,omitempty"`
	BillingAddress  *POSTCustomerAddressesRequestDataRelationshipsAddress            `json:"billing_address,omitempty"`
	PaymentMethod   *POSTOrdersRequestDataRelationshipsPaymentMethod                 `json:"payment_method,omitempty"`
	PaymentSource   *POSTCustomerPaymentSourcesRequestDataRelationshipsPaymentSource `json:"payment_source,omitempty"`
}

// NewPOSTOrdersRequestDataRelationships instantiates a new POSTOrdersRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTOrdersRequestDataRelationships() *POSTOrdersRequestDataRelationships {
	this := POSTOrdersRequestDataRelationships{}
	return &this
}

// NewPOSTOrdersRequestDataRelationshipsWithDefaults instantiates a new POSTOrdersRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTOrdersRequestDataRelationshipsWithDefaults() *POSTOrdersRequestDataRelationships {
	this := POSTOrdersRequestDataRelationships{}
	return &this
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *POSTOrdersRequestDataRelationships) GetMarket() POSTBillingInfoValidationRulesRequestDataRelationshipsMarket {
	if o == nil || IsNil(o.Market) {
		var ret POSTBillingInfoValidationRulesRequestDataRelationshipsMarket
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrdersRequestDataRelationships) GetMarketOk() (*POSTBillingInfoValidationRulesRequestDataRelationshipsMarket, bool) {
	if o == nil || IsNil(o.Market) {
		return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *POSTOrdersRequestDataRelationships) HasMarket() bool {
	if o != nil && !IsNil(o.Market) {
		return true
	}

	return false
}

// SetMarket gets a reference to the given POSTBillingInfoValidationRulesRequestDataRelationshipsMarket and assigns it to the Market field.
func (o *POSTOrdersRequestDataRelationships) SetMarket(v POSTBillingInfoValidationRulesRequestDataRelationshipsMarket) {
	o.Market = &v
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *POSTOrdersRequestDataRelationships) GetCustomer() POSTCouponRecipientsRequestDataRelationshipsCustomer {
	if o == nil || IsNil(o.Customer) {
		var ret POSTCouponRecipientsRequestDataRelationshipsCustomer
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrdersRequestDataRelationships) GetCustomerOk() (*POSTCouponRecipientsRequestDataRelationshipsCustomer, bool) {
	if o == nil || IsNil(o.Customer) {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *POSTOrdersRequestDataRelationships) HasCustomer() bool {
	if o != nil && !IsNil(o.Customer) {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given POSTCouponRecipientsRequestDataRelationshipsCustomer and assigns it to the Customer field.
func (o *POSTOrdersRequestDataRelationships) SetCustomer(v POSTCouponRecipientsRequestDataRelationshipsCustomer) {
	o.Customer = &v
}

// GetShippingAddress returns the ShippingAddress field value if set, zero value otherwise.
func (o *POSTOrdersRequestDataRelationships) GetShippingAddress() POSTCustomerAddressesRequestDataRelationshipsAddress {
	if o == nil || IsNil(o.ShippingAddress) {
		var ret POSTCustomerAddressesRequestDataRelationshipsAddress
		return ret
	}
	return *o.ShippingAddress
}

// GetShippingAddressOk returns a tuple with the ShippingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrdersRequestDataRelationships) GetShippingAddressOk() (*POSTCustomerAddressesRequestDataRelationshipsAddress, bool) {
	if o == nil || IsNil(o.ShippingAddress) {
		return nil, false
	}
	return o.ShippingAddress, true
}

// HasShippingAddress returns a boolean if a field has been set.
func (o *POSTOrdersRequestDataRelationships) HasShippingAddress() bool {
	if o != nil && !IsNil(o.ShippingAddress) {
		return true
	}

	return false
}

// SetShippingAddress gets a reference to the given POSTCustomerAddressesRequestDataRelationshipsAddress and assigns it to the ShippingAddress field.
func (o *POSTOrdersRequestDataRelationships) SetShippingAddress(v POSTCustomerAddressesRequestDataRelationshipsAddress) {
	o.ShippingAddress = &v
}

// GetBillingAddress returns the BillingAddress field value if set, zero value otherwise.
func (o *POSTOrdersRequestDataRelationships) GetBillingAddress() POSTCustomerAddressesRequestDataRelationshipsAddress {
	if o == nil || IsNil(o.BillingAddress) {
		var ret POSTCustomerAddressesRequestDataRelationshipsAddress
		return ret
	}
	return *o.BillingAddress
}

// GetBillingAddressOk returns a tuple with the BillingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrdersRequestDataRelationships) GetBillingAddressOk() (*POSTCustomerAddressesRequestDataRelationshipsAddress, bool) {
	if o == nil || IsNil(o.BillingAddress) {
		return nil, false
	}
	return o.BillingAddress, true
}

// HasBillingAddress returns a boolean if a field has been set.
func (o *POSTOrdersRequestDataRelationships) HasBillingAddress() bool {
	if o != nil && !IsNil(o.BillingAddress) {
		return true
	}

	return false
}

// SetBillingAddress gets a reference to the given POSTCustomerAddressesRequestDataRelationshipsAddress and assigns it to the BillingAddress field.
func (o *POSTOrdersRequestDataRelationships) SetBillingAddress(v POSTCustomerAddressesRequestDataRelationshipsAddress) {
	o.BillingAddress = &v
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise.
func (o *POSTOrdersRequestDataRelationships) GetPaymentMethod() POSTOrdersRequestDataRelationshipsPaymentMethod {
	if o == nil || IsNil(o.PaymentMethod) {
		var ret POSTOrdersRequestDataRelationshipsPaymentMethod
		return ret
	}
	return *o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrdersRequestDataRelationships) GetPaymentMethodOk() (*POSTOrdersRequestDataRelationshipsPaymentMethod, bool) {
	if o == nil || IsNil(o.PaymentMethod) {
		return nil, false
	}
	return o.PaymentMethod, true
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *POSTOrdersRequestDataRelationships) HasPaymentMethod() bool {
	if o != nil && !IsNil(o.PaymentMethod) {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given POSTOrdersRequestDataRelationshipsPaymentMethod and assigns it to the PaymentMethod field.
func (o *POSTOrdersRequestDataRelationships) SetPaymentMethod(v POSTOrdersRequestDataRelationshipsPaymentMethod) {
	o.PaymentMethod = &v
}

// GetPaymentSource returns the PaymentSource field value if set, zero value otherwise.
func (o *POSTOrdersRequestDataRelationships) GetPaymentSource() POSTCustomerPaymentSourcesRequestDataRelationshipsPaymentSource {
	if o == nil || IsNil(o.PaymentSource) {
		var ret POSTCustomerPaymentSourcesRequestDataRelationshipsPaymentSource
		return ret
	}
	return *o.PaymentSource
}

// GetPaymentSourceOk returns a tuple with the PaymentSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrdersRequestDataRelationships) GetPaymentSourceOk() (*POSTCustomerPaymentSourcesRequestDataRelationshipsPaymentSource, bool) {
	if o == nil || IsNil(o.PaymentSource) {
		return nil, false
	}
	return o.PaymentSource, true
}

// HasPaymentSource returns a boolean if a field has been set.
func (o *POSTOrdersRequestDataRelationships) HasPaymentSource() bool {
	if o != nil && !IsNil(o.PaymentSource) {
		return true
	}

	return false
}

// SetPaymentSource gets a reference to the given POSTCustomerPaymentSourcesRequestDataRelationshipsPaymentSource and assigns it to the PaymentSource field.
func (o *POSTOrdersRequestDataRelationships) SetPaymentSource(v POSTCustomerPaymentSourcesRequestDataRelationshipsPaymentSource) {
	o.PaymentSource = &v
}

func (o POSTOrdersRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTOrdersRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Market) {
		toSerialize["market"] = o.Market
	}
	if !IsNil(o.Customer) {
		toSerialize["customer"] = o.Customer
	}
	if !IsNil(o.ShippingAddress) {
		toSerialize["shipping_address"] = o.ShippingAddress
	}
	if !IsNil(o.BillingAddress) {
		toSerialize["billing_address"] = o.BillingAddress
	}
	if !IsNil(o.PaymentMethod) {
		toSerialize["payment_method"] = o.PaymentMethod
	}
	if !IsNil(o.PaymentSource) {
		toSerialize["payment_source"] = o.PaymentSource
	}
	return toSerialize, nil
}

type NullablePOSTOrdersRequestDataRelationships struct {
	value *POSTOrdersRequestDataRelationships
	isSet bool
}

func (v NullablePOSTOrdersRequestDataRelationships) Get() *POSTOrdersRequestDataRelationships {
	return v.value
}

func (v *NullablePOSTOrdersRequestDataRelationships) Set(val *POSTOrdersRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTOrdersRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTOrdersRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTOrdersRequestDataRelationships(val *POSTOrdersRequestDataRelationships) *NullablePOSTOrdersRequestDataRelationships {
	return &NullablePOSTOrdersRequestDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTOrdersRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTOrdersRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
