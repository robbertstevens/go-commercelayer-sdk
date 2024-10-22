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

// checks if the CustomerPaymentSourceDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerPaymentSourceDataRelationships{}

// CustomerPaymentSourceDataRelationships struct for CustomerPaymentSourceDataRelationships
type CustomerPaymentSourceDataRelationships struct {
	Customer      *CouponRecipientDataRelationshipsCustomer            `json:"customer,omitempty"`
	PaymentMethod *AdyenGatewayDataRelationshipsPaymentMethods         `json:"payment_method,omitempty"`
	PaymentSource *CustomerPaymentSourceDataRelationshipsPaymentSource `json:"payment_source,omitempty"`
	Versions      *AddressDataRelationshipsVersions                    `json:"versions,omitempty"`
}

// NewCustomerPaymentSourceDataRelationships instantiates a new CustomerPaymentSourceDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerPaymentSourceDataRelationships() *CustomerPaymentSourceDataRelationships {
	this := CustomerPaymentSourceDataRelationships{}
	return &this
}

// NewCustomerPaymentSourceDataRelationshipsWithDefaults instantiates a new CustomerPaymentSourceDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerPaymentSourceDataRelationshipsWithDefaults() *CustomerPaymentSourceDataRelationships {
	this := CustomerPaymentSourceDataRelationships{}
	return &this
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *CustomerPaymentSourceDataRelationships) GetCustomer() CouponRecipientDataRelationshipsCustomer {
	if o == nil || IsNil(o.Customer) {
		var ret CouponRecipientDataRelationshipsCustomer
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerPaymentSourceDataRelationships) GetCustomerOk() (*CouponRecipientDataRelationshipsCustomer, bool) {
	if o == nil || IsNil(o.Customer) {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *CustomerPaymentSourceDataRelationships) HasCustomer() bool {
	if o != nil && !IsNil(o.Customer) {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given CouponRecipientDataRelationshipsCustomer and assigns it to the Customer field.
func (o *CustomerPaymentSourceDataRelationships) SetCustomer(v CouponRecipientDataRelationshipsCustomer) {
	o.Customer = &v
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise.
func (o *CustomerPaymentSourceDataRelationships) GetPaymentMethod() AdyenGatewayDataRelationshipsPaymentMethods {
	if o == nil || IsNil(o.PaymentMethod) {
		var ret AdyenGatewayDataRelationshipsPaymentMethods
		return ret
	}
	return *o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerPaymentSourceDataRelationships) GetPaymentMethodOk() (*AdyenGatewayDataRelationshipsPaymentMethods, bool) {
	if o == nil || IsNil(o.PaymentMethod) {
		return nil, false
	}
	return o.PaymentMethod, true
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *CustomerPaymentSourceDataRelationships) HasPaymentMethod() bool {
	if o != nil && !IsNil(o.PaymentMethod) {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given AdyenGatewayDataRelationshipsPaymentMethods and assigns it to the PaymentMethod field.
func (o *CustomerPaymentSourceDataRelationships) SetPaymentMethod(v AdyenGatewayDataRelationshipsPaymentMethods) {
	o.PaymentMethod = &v
}

// GetPaymentSource returns the PaymentSource field value if set, zero value otherwise.
func (o *CustomerPaymentSourceDataRelationships) GetPaymentSource() CustomerPaymentSourceDataRelationshipsPaymentSource {
	if o == nil || IsNil(o.PaymentSource) {
		var ret CustomerPaymentSourceDataRelationshipsPaymentSource
		return ret
	}
	return *o.PaymentSource
}

// GetPaymentSourceOk returns a tuple with the PaymentSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerPaymentSourceDataRelationships) GetPaymentSourceOk() (*CustomerPaymentSourceDataRelationshipsPaymentSource, bool) {
	if o == nil || IsNil(o.PaymentSource) {
		return nil, false
	}
	return o.PaymentSource, true
}

// HasPaymentSource returns a boolean if a field has been set.
func (o *CustomerPaymentSourceDataRelationships) HasPaymentSource() bool {
	if o != nil && !IsNil(o.PaymentSource) {
		return true
	}

	return false
}

// SetPaymentSource gets a reference to the given CustomerPaymentSourceDataRelationshipsPaymentSource and assigns it to the PaymentSource field.
func (o *CustomerPaymentSourceDataRelationships) SetPaymentSource(v CustomerPaymentSourceDataRelationshipsPaymentSource) {
	o.PaymentSource = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *CustomerPaymentSourceDataRelationships) GetVersions() AddressDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret AddressDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerPaymentSourceDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *CustomerPaymentSourceDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given AddressDataRelationshipsVersions and assigns it to the Versions field.
func (o *CustomerPaymentSourceDataRelationships) SetVersions(v AddressDataRelationshipsVersions) {
	o.Versions = &v
}

func (o CustomerPaymentSourceDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerPaymentSourceDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Customer) {
		toSerialize["customer"] = o.Customer
	}
	if !IsNil(o.PaymentMethod) {
		toSerialize["payment_method"] = o.PaymentMethod
	}
	if !IsNil(o.PaymentSource) {
		toSerialize["payment_source"] = o.PaymentSource
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	return toSerialize, nil
}

type NullableCustomerPaymentSourceDataRelationships struct {
	value *CustomerPaymentSourceDataRelationships
	isSet bool
}

func (v NullableCustomerPaymentSourceDataRelationships) Get() *CustomerPaymentSourceDataRelationships {
	return v.value
}

func (v *NullableCustomerPaymentSourceDataRelationships) Set(val *CustomerPaymentSourceDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerPaymentSourceDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerPaymentSourceDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerPaymentSourceDataRelationships(val *CustomerPaymentSourceDataRelationships) *NullableCustomerPaymentSourceDataRelationships {
	return &NullableCustomerPaymentSourceDataRelationships{value: val, isSet: true}
}

func (v NullableCustomerPaymentSourceDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerPaymentSourceDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
