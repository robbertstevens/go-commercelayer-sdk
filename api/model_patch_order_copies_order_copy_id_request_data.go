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

// checks if the PATCHOrderCopiesOrderCopyIdRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHOrderCopiesOrderCopyIdRequestData{}

// PATCHOrderCopiesOrderCopyIdRequestData struct for PATCHOrderCopiesOrderCopyIdRequestData
type PATCHOrderCopiesOrderCopyIdRequestData struct {
	// The resource's type
	Type interface{} `json:"type"`
	// The resource's id
	Id            interface{}                                                                     `json:"id"`
	Attributes    PATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequestDataAttributes `json:"attributes"`
	Relationships interface{}                                                                     `json:"relationships,omitempty"`
}

// NewPATCHOrderCopiesOrderCopyIdRequestData instantiates a new PATCHOrderCopiesOrderCopyIdRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHOrderCopiesOrderCopyIdRequestData(type_ interface{}, id interface{}, attributes PATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequestDataAttributes) *PATCHOrderCopiesOrderCopyIdRequestData {
	this := PATCHOrderCopiesOrderCopyIdRequestData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewPATCHOrderCopiesOrderCopyIdRequestDataWithDefaults instantiates a new PATCHOrderCopiesOrderCopyIdRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHOrderCopiesOrderCopyIdRequestDataWithDefaults() *PATCHOrderCopiesOrderCopyIdRequestData {
	this := PATCHOrderCopiesOrderCopyIdRequestData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *PATCHOrderCopiesOrderCopyIdRequestData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHOrderCopiesOrderCopyIdRequestData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PATCHOrderCopiesOrderCopyIdRequestData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *PATCHOrderCopiesOrderCopyIdRequestData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHOrderCopiesOrderCopyIdRequestData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PATCHOrderCopiesOrderCopyIdRequestData) SetId(v interface{}) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *PATCHOrderCopiesOrderCopyIdRequestData) GetAttributes() PATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequestDataAttributes {
	if o == nil {
		var ret PATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *PATCHOrderCopiesOrderCopyIdRequestData) GetAttributesOk() (*PATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *PATCHOrderCopiesOrderCopyIdRequestData) SetAttributes(v PATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHOrderCopiesOrderCopyIdRequestData) GetRelationships() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHOrderCopiesOrderCopyIdRequestData) GetRelationshipsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return &o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PATCHOrderCopiesOrderCopyIdRequestData) HasRelationships() bool {
	if o != nil && IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given interface{} and assigns it to the Relationships field.
func (o *PATCHOrderCopiesOrderCopyIdRequestData) SetRelationships(v interface{}) {
	o.Relationships = v
}

func (o PATCHOrderCopiesOrderCopyIdRequestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHOrderCopiesOrderCopyIdRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["attributes"] = o.Attributes
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullablePATCHOrderCopiesOrderCopyIdRequestData struct {
	value *PATCHOrderCopiesOrderCopyIdRequestData
	isSet bool
}

func (v NullablePATCHOrderCopiesOrderCopyIdRequestData) Get() *PATCHOrderCopiesOrderCopyIdRequestData {
	return v.value
}

func (v *NullablePATCHOrderCopiesOrderCopyIdRequestData) Set(val *PATCHOrderCopiesOrderCopyIdRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHOrderCopiesOrderCopyIdRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHOrderCopiesOrderCopyIdRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHOrderCopiesOrderCopyIdRequestData(val *PATCHOrderCopiesOrderCopyIdRequestData) *NullablePATCHOrderCopiesOrderCopyIdRequestData {
	return &NullablePATCHOrderCopiesOrderCopyIdRequestData{value: val, isSet: true}
}

func (v NullablePATCHOrderCopiesOrderCopyIdRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHOrderCopiesOrderCopyIdRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
