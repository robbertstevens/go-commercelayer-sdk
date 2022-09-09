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

// CheckoutComGatewayData struct for CheckoutComGatewayData
type CheckoutComGatewayData struct {
	// The resource's type
	Type string `json:"type"`
	Attributes GETCheckoutComGateways200ResponseDataInnerAttributes `json:"attributes"`
	Relationships *CheckoutComGatewayDataRelationships `json:"relationships,omitempty"`
}

// NewCheckoutComGatewayData instantiates a new CheckoutComGatewayData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutComGatewayData(type_ string, attributes GETCheckoutComGateways200ResponseDataInnerAttributes) *CheckoutComGatewayData {
	this := CheckoutComGatewayData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewCheckoutComGatewayDataWithDefaults instantiates a new CheckoutComGatewayData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutComGatewayDataWithDefaults() *CheckoutComGatewayData {
	this := CheckoutComGatewayData{}
	var type_ string = "checkout_com_gateways"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *CheckoutComGatewayData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CheckoutComGatewayData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CheckoutComGatewayData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *CheckoutComGatewayData) GetAttributes() GETCheckoutComGateways200ResponseDataInnerAttributes {
	if o == nil {
		var ret GETCheckoutComGateways200ResponseDataInnerAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CheckoutComGatewayData) GetAttributesOk() (*GETCheckoutComGateways200ResponseDataInnerAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CheckoutComGatewayData) SetAttributes(v GETCheckoutComGateways200ResponseDataInnerAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *CheckoutComGatewayData) GetRelationships() CheckoutComGatewayDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret CheckoutComGatewayDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutComGatewayData) GetRelationshipsOk() (*CheckoutComGatewayDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CheckoutComGatewayData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given CheckoutComGatewayDataRelationships and assigns it to the Relationships field.
func (o *CheckoutComGatewayData) SetRelationships(v CheckoutComGatewayDataRelationships) {
	o.Relationships = &v
}

func (o CheckoutComGatewayData) MarshalJSON() ([]byte, error) {
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

type NullableCheckoutComGatewayData struct {
	value *CheckoutComGatewayData
	isSet bool
}

func (v NullableCheckoutComGatewayData) Get() *CheckoutComGatewayData {
	return v.value
}

func (v *NullableCheckoutComGatewayData) Set(val *CheckoutComGatewayData) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutComGatewayData) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutComGatewayData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutComGatewayData(val *CheckoutComGatewayData) *NullableCheckoutComGatewayData {
	return &NullableCheckoutComGatewayData{value: val, isSet: true}
}

func (v NullableCheckoutComGatewayData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutComGatewayData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


