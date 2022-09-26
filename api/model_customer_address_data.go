/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// CustomerAddressData struct for CustomerAddressData
type CustomerAddressData struct {
	// The resource's type
	Type          string                                             `json:"type"`
	Attributes    GETCustomerAddresses200ResponseDataInnerAttributes `json:"attributes"`
	Relationships *CustomerAddressDataRelationships                  `json:"relationships,omitempty"`
}

// NewCustomerAddressData instantiates a new CustomerAddressData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerAddressData(type_ string, attributes GETCustomerAddresses200ResponseDataInnerAttributes) *CustomerAddressData {
	this := CustomerAddressData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewCustomerAddressDataWithDefaults instantiates a new CustomerAddressData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerAddressDataWithDefaults() *CustomerAddressData {
	this := CustomerAddressData{}
	return &this
}

// GetType returns the Type field value
func (o *CustomerAddressData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CustomerAddressData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CustomerAddressData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *CustomerAddressData) GetAttributes() GETCustomerAddresses200ResponseDataInnerAttributes {
	if o == nil {
		var ret GETCustomerAddresses200ResponseDataInnerAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CustomerAddressData) GetAttributesOk() (*GETCustomerAddresses200ResponseDataInnerAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CustomerAddressData) SetAttributes(v GETCustomerAddresses200ResponseDataInnerAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *CustomerAddressData) GetRelationships() CustomerAddressDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret CustomerAddressDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressData) GetRelationshipsOk() (*CustomerAddressDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CustomerAddressData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given CustomerAddressDataRelationships and assigns it to the Relationships field.
func (o *CustomerAddressData) SetRelationships(v CustomerAddressDataRelationships) {
	o.Relationships = &v
}

func (o CustomerAddressData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerAddressData struct {
	value *CustomerAddressData
	isSet bool
}

func (v NullableCustomerAddressData) Get() *CustomerAddressData {
	return v.value
}

func (v *NullableCustomerAddressData) Set(val *CustomerAddressData) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerAddressData) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerAddressData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerAddressData(val *CustomerAddressData) *NullableCustomerAddressData {
	return &NullableCustomerAddressData{value: val, isSet: true}
}

func (v NullableCustomerAddressData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerAddressData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
