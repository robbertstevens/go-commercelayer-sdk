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

// checks if the PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestData{}

// PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestData struct for PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestData
type PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestData struct {
	// The resource's type
	Type interface{} `json:"type"`
	// The resource's id
	Id            interface{}                                                               `json:"id"`
	Attributes    PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestDataAttributes     `json:"attributes"`
	Relationships *PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestDataRelationships `json:"relationships,omitempty"`
}

// NewPATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestData instantiates a new PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestData(type_ interface{}, id interface{}, attributes PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestDataAttributes) *PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestData {
	this := PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewPATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestDataWithDefaults instantiates a new PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestDataWithDefaults() *PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestData {
	this := PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestData) SetId(v interface{}) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestData) GetAttributes() PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestDataAttributes {
	if o == nil {
		var ret PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestData) GetAttributesOk() (*PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestData) SetAttributes(v PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestData) GetRelationships() PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestData) GetRelationshipsOk() (*PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestDataRelationships and assigns it to the Relationships field.
func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestData) SetRelationships(v PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestDataRelationships) {
	o.Relationships = &v
}

func (o PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["attributes"] = o.Attributes
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullablePATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestData struct {
	value *PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestData
	isSet bool
}

func (v NullablePATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestData) Get() *PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestData {
	return v.value
}

func (v *NullablePATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestData) Set(val *PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestData(val *PATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestData) *NullablePATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestData {
	return &NullablePATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestData{value: val, isSet: true}
}

func (v NullablePATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHSkuListPromotionRulesSkuListPromotionRuleIdRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
