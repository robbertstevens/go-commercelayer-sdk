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

// StripeGatewayUpdateData struct for StripeGatewayUpdateData
type StripeGatewayUpdateData struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id            string                                                      `json:"id"`
	Attributes    PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes `json:"attributes"`
	Relationships map[string]interface{}                                      `json:"relationships,omitempty"`
}

// NewStripeGatewayUpdateData instantiates a new StripeGatewayUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStripeGatewayUpdateData(type_ string, id string, attributes PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) *StripeGatewayUpdateData {
	this := StripeGatewayUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewStripeGatewayUpdateDataWithDefaults instantiates a new StripeGatewayUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStripeGatewayUpdateDataWithDefaults() *StripeGatewayUpdateData {
	this := StripeGatewayUpdateData{}
	return &this
}

// GetType returns the Type field value
func (o *StripeGatewayUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StripeGatewayUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *StripeGatewayUpdateData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *StripeGatewayUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *StripeGatewayUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *StripeGatewayUpdateData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *StripeGatewayUpdateData) GetAttributes() PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes {
	if o == nil {
		var ret PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *StripeGatewayUpdateData) GetAttributesOk() (*PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *StripeGatewayUpdateData) SetAttributes(v PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *StripeGatewayUpdateData) GetRelationships() map[string]interface{} {
	if o == nil || o.Relationships == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripeGatewayUpdateData) GetRelationshipsOk() (map[string]interface{}, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *StripeGatewayUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]interface{} and assigns it to the Relationships field.
func (o *StripeGatewayUpdateData) SetRelationships(v map[string]interface{}) {
	o.Relationships = v
}

func (o StripeGatewayUpdateData) MarshalJSON() ([]byte, error) {
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

type NullableStripeGatewayUpdateData struct {
	value *StripeGatewayUpdateData
	isSet bool
}

func (v NullableStripeGatewayUpdateData) Get() *StripeGatewayUpdateData {
	return v.value
}

func (v *NullableStripeGatewayUpdateData) Set(val *StripeGatewayUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableStripeGatewayUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableStripeGatewayUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStripeGatewayUpdateData(val *StripeGatewayUpdateData) *NullableStripeGatewayUpdateData {
	return &NullableStripeGatewayUpdateData{value: val, isSet: true}
}

func (v NullableStripeGatewayUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStripeGatewayUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
