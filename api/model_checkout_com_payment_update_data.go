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

// CheckoutComPaymentUpdateData struct for CheckoutComPaymentUpdateData
type CheckoutComPaymentUpdateData struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id string `json:"id"`
	Attributes PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes `json:"attributes"`
	Relationships *AdyenPaymentUpdateDataRelationships `json:"relationships,omitempty"`
}

// NewCheckoutComPaymentUpdateData instantiates a new CheckoutComPaymentUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutComPaymentUpdateData(type_ string, id string, attributes PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) *CheckoutComPaymentUpdateData {
	this := CheckoutComPaymentUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewCheckoutComPaymentUpdateDataWithDefaults instantiates a new CheckoutComPaymentUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutComPaymentUpdateDataWithDefaults() *CheckoutComPaymentUpdateData {
	this := CheckoutComPaymentUpdateData{}
	var type_ string = "checkout_com_payments"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *CheckoutComPaymentUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CheckoutComPaymentUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CheckoutComPaymentUpdateData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *CheckoutComPaymentUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CheckoutComPaymentUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CheckoutComPaymentUpdateData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *CheckoutComPaymentUpdateData) GetAttributes() PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes {
	if o == nil {
		var ret PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CheckoutComPaymentUpdateData) GetAttributesOk() (*PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CheckoutComPaymentUpdateData) SetAttributes(v PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *CheckoutComPaymentUpdateData) GetRelationships() AdyenPaymentUpdateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret AdyenPaymentUpdateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutComPaymentUpdateData) GetRelationshipsOk() (*AdyenPaymentUpdateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CheckoutComPaymentUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AdyenPaymentUpdateDataRelationships and assigns it to the Relationships field.
func (o *CheckoutComPaymentUpdateData) SetRelationships(v AdyenPaymentUpdateDataRelationships) {
	o.Relationships = &v
}

func (o CheckoutComPaymentUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableCheckoutComPaymentUpdateData struct {
	value *CheckoutComPaymentUpdateData
	isSet bool
}

func (v NullableCheckoutComPaymentUpdateData) Get() *CheckoutComPaymentUpdateData {
	return v.value
}

func (v *NullableCheckoutComPaymentUpdateData) Set(val *CheckoutComPaymentUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutComPaymentUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutComPaymentUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutComPaymentUpdateData(val *CheckoutComPaymentUpdateData) *NullableCheckoutComPaymentUpdateData {
	return &NullableCheckoutComPaymentUpdateData{value: val, isSet: true}
}

func (v NullableCheckoutComPaymentUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutComPaymentUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


