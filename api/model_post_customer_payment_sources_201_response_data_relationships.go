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

// checks if the POSTCustomerPaymentSources201ResponseDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTCustomerPaymentSources201ResponseDataRelationships{}

// POSTCustomerPaymentSources201ResponseDataRelationships struct for POSTCustomerPaymentSources201ResponseDataRelationships
type POSTCustomerPaymentSources201ResponseDataRelationships struct {
	Customer      *POSTCouponRecipients201ResponseDataRelationshipsCustomer            `json:"customer,omitempty"`
	PaymentMethod *POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentMethod `json:"payment_method,omitempty"`
	PaymentSource *POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSource `json:"payment_source,omitempty"`
	Versions      *POSTAddresses201ResponseDataRelationshipsVersions                   `json:"versions,omitempty"`
}

// NewPOSTCustomerPaymentSources201ResponseDataRelationships instantiates a new POSTCustomerPaymentSources201ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTCustomerPaymentSources201ResponseDataRelationships() *POSTCustomerPaymentSources201ResponseDataRelationships {
	this := POSTCustomerPaymentSources201ResponseDataRelationships{}
	return &this
}

// NewPOSTCustomerPaymentSources201ResponseDataRelationshipsWithDefaults instantiates a new POSTCustomerPaymentSources201ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTCustomerPaymentSources201ResponseDataRelationshipsWithDefaults() *POSTCustomerPaymentSources201ResponseDataRelationships {
	this := POSTCustomerPaymentSources201ResponseDataRelationships{}
	return &this
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *POSTCustomerPaymentSources201ResponseDataRelationships) GetCustomer() POSTCouponRecipients201ResponseDataRelationshipsCustomer {
	if o == nil || IsNil(o.Customer) {
		var ret POSTCouponRecipients201ResponseDataRelationshipsCustomer
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCustomerPaymentSources201ResponseDataRelationships) GetCustomerOk() (*POSTCouponRecipients201ResponseDataRelationshipsCustomer, bool) {
	if o == nil || IsNil(o.Customer) {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *POSTCustomerPaymentSources201ResponseDataRelationships) HasCustomer() bool {
	if o != nil && !IsNil(o.Customer) {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given POSTCouponRecipients201ResponseDataRelationshipsCustomer and assigns it to the Customer field.
func (o *POSTCustomerPaymentSources201ResponseDataRelationships) SetCustomer(v POSTCouponRecipients201ResponseDataRelationshipsCustomer) {
	o.Customer = &v
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise.
func (o *POSTCustomerPaymentSources201ResponseDataRelationships) GetPaymentMethod() POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentMethod {
	if o == nil || IsNil(o.PaymentMethod) {
		var ret POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentMethod
		return ret
	}
	return *o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCustomerPaymentSources201ResponseDataRelationships) GetPaymentMethodOk() (*POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentMethod, bool) {
	if o == nil || IsNil(o.PaymentMethod) {
		return nil, false
	}
	return o.PaymentMethod, true
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *POSTCustomerPaymentSources201ResponseDataRelationships) HasPaymentMethod() bool {
	if o != nil && !IsNil(o.PaymentMethod) {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentMethod and assigns it to the PaymentMethod field.
func (o *POSTCustomerPaymentSources201ResponseDataRelationships) SetPaymentMethod(v POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentMethod) {
	o.PaymentMethod = &v
}

// GetPaymentSource returns the PaymentSource field value if set, zero value otherwise.
func (o *POSTCustomerPaymentSources201ResponseDataRelationships) GetPaymentSource() POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSource {
	if o == nil || IsNil(o.PaymentSource) {
		var ret POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSource
		return ret
	}
	return *o.PaymentSource
}

// GetPaymentSourceOk returns a tuple with the PaymentSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCustomerPaymentSources201ResponseDataRelationships) GetPaymentSourceOk() (*POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSource, bool) {
	if o == nil || IsNil(o.PaymentSource) {
		return nil, false
	}
	return o.PaymentSource, true
}

// HasPaymentSource returns a boolean if a field has been set.
func (o *POSTCustomerPaymentSources201ResponseDataRelationships) HasPaymentSource() bool {
	if o != nil && !IsNil(o.PaymentSource) {
		return true
	}

	return false
}

// SetPaymentSource gets a reference to the given POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSource and assigns it to the PaymentSource field.
func (o *POSTCustomerPaymentSources201ResponseDataRelationships) SetPaymentSource(v POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSource) {
	o.PaymentSource = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *POSTCustomerPaymentSources201ResponseDataRelationships) GetVersions() POSTAddresses201ResponseDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret POSTAddresses201ResponseDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCustomerPaymentSources201ResponseDataRelationships) GetVersionsOk() (*POSTAddresses201ResponseDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *POSTCustomerPaymentSources201ResponseDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given POSTAddresses201ResponseDataRelationshipsVersions and assigns it to the Versions field.
func (o *POSTCustomerPaymentSources201ResponseDataRelationships) SetVersions(v POSTAddresses201ResponseDataRelationshipsVersions) {
	o.Versions = &v
}

func (o POSTCustomerPaymentSources201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTCustomerPaymentSources201ResponseDataRelationships) ToMap() (map[string]interface{}, error) {
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

type NullablePOSTCustomerPaymentSources201ResponseDataRelationships struct {
	value *POSTCustomerPaymentSources201ResponseDataRelationships
	isSet bool
}

func (v NullablePOSTCustomerPaymentSources201ResponseDataRelationships) Get() *POSTCustomerPaymentSources201ResponseDataRelationships {
	return v.value
}

func (v *NullablePOSTCustomerPaymentSources201ResponseDataRelationships) Set(val *POSTCustomerPaymentSources201ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTCustomerPaymentSources201ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTCustomerPaymentSources201ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTCustomerPaymentSources201ResponseDataRelationships(val *POSTCustomerPaymentSources201ResponseDataRelationships) *NullablePOSTCustomerPaymentSources201ResponseDataRelationships {
	return &NullablePOSTCustomerPaymentSources201ResponseDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTCustomerPaymentSources201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTCustomerPaymentSources201ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
