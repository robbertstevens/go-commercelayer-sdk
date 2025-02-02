/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.3.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes{}

// GETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes struct for GETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes
type GETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes struct {
	// The promotion rule's type.
	Type interface{} `json:"type,omitempty"`
	// Time at which the resource was created.
	CreatedAt interface{} `json:"created_at,omitempty"`
	// Time at which the resource was last updated.
	UpdatedAt interface{} `json:"updated_at,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
	// The filters used to trigger promotion on the matching order and its relationships attributes.
	Filters interface{} `json:"filters,omitempty"`
}

// NewGETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes instantiates a new GETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes() *GETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes {
	this := GETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes{}
	return &this
}

// NewGETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributesWithDefaults instantiates a new GETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributesWithDefaults() *GETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes {
	this := GETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *GETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes) SetType(v interface{}) {
	o.Type = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes) GetCreatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return &o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes) HasCreatedAt() bool {
	if o != nil && IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given interface{} and assigns it to the CreatedAt field.
func (o *GETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes) SetCreatedAt(v interface{}) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes) GetUpdatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes) HasUpdatedAt() bool {
	if o != nil && IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given interface{} and assigns it to the UpdatedAt field.
func (o *GETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes) SetUpdatedAt(v interface{}) {
	o.UpdatedAt = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *GETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *GETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *GETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

// GetFilters returns the Filters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes) GetFilters() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes) GetFiltersOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return &o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *GETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes) HasFilters() bool {
	if o != nil && IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given interface{} and assigns it to the Filters field.
func (o *GETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes) SetFilters(v interface{}) {
	o.Filters = v
}

func (o GETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
	}
	if o.ReferenceOrigin != nil {
		toSerialize["reference_origin"] = o.ReferenceOrigin
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Filters != nil {
		toSerialize["filters"] = o.Filters
	}
	return toSerialize, nil
}

type NullableGETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes struct {
	value *GETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes
	isSet bool
}

func (v NullableGETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes) Get() *GETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes {
	return v.value
}

func (v *NullableGETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes) Set(val *GETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes(val *GETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes) *NullableGETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes {
	return &NullableGETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullableGETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
