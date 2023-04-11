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

// checks if the PATCHSkuListsSkuListIdRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHSkuListsSkuListIdRequestDataAttributes{}

// PATCHSkuListsSkuListIdRequestDataAttributes struct for PATCHSkuListsSkuListIdRequestDataAttributes
type PATCHSkuListsSkuListIdRequestDataAttributes struct {
	// The SKU list's internal name.
	Name interface{} `json:"name,omitempty"`
	// An internal description of the SKU list.
	Description interface{} `json:"description,omitempty"`
	// The URL of an image that represents the SKU list.
	ImageUrl interface{} `json:"image_url,omitempty"`
	// Indicates if the SKU list is populated manually.
	Manual interface{} `json:"manual,omitempty"`
	// The regex that will be evaluated to match the SKU codes.
	SkuCodeRegex interface{} `json:"sku_code_regex,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPATCHSkuListsSkuListIdRequestDataAttributes instantiates a new PATCHSkuListsSkuListIdRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHSkuListsSkuListIdRequestDataAttributes() *PATCHSkuListsSkuListIdRequestDataAttributes {
	this := PATCHSkuListsSkuListIdRequestDataAttributes{}
	return &this
}

// NewPATCHSkuListsSkuListIdRequestDataAttributesWithDefaults instantiates a new PATCHSkuListsSkuListIdRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHSkuListsSkuListIdRequestDataAttributesWithDefaults() *PATCHSkuListsSkuListIdRequestDataAttributes {
	this := PATCHSkuListsSkuListIdRequestDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHSkuListsSkuListIdRequestDataAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHSkuListsSkuListIdRequestDataAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PATCHSkuListsSkuListIdRequestDataAttributes) HasName() bool {
	if o != nil && IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given interface{} and assigns it to the Name field.
func (o *PATCHSkuListsSkuListIdRequestDataAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHSkuListsSkuListIdRequestDataAttributes) GetDescription() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHSkuListsSkuListIdRequestDataAttributes) GetDescriptionOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return &o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PATCHSkuListsSkuListIdRequestDataAttributes) HasDescription() bool {
	if o != nil && IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given interface{} and assigns it to the Description field.
func (o *PATCHSkuListsSkuListIdRequestDataAttributes) SetDescription(v interface{}) {
	o.Description = v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHSkuListsSkuListIdRequestDataAttributes) GetImageUrl() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHSkuListsSkuListIdRequestDataAttributes) GetImageUrlOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ImageUrl) {
		return nil, false
	}
	return &o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *PATCHSkuListsSkuListIdRequestDataAttributes) HasImageUrl() bool {
	if o != nil && IsNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given interface{} and assigns it to the ImageUrl field.
func (o *PATCHSkuListsSkuListIdRequestDataAttributes) SetImageUrl(v interface{}) {
	o.ImageUrl = v
}

// GetManual returns the Manual field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHSkuListsSkuListIdRequestDataAttributes) GetManual() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Manual
}

// GetManualOk returns a tuple with the Manual field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHSkuListsSkuListIdRequestDataAttributes) GetManualOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Manual) {
		return nil, false
	}
	return &o.Manual, true
}

// HasManual returns a boolean if a field has been set.
func (o *PATCHSkuListsSkuListIdRequestDataAttributes) HasManual() bool {
	if o != nil && IsNil(o.Manual) {
		return true
	}

	return false
}

// SetManual gets a reference to the given interface{} and assigns it to the Manual field.
func (o *PATCHSkuListsSkuListIdRequestDataAttributes) SetManual(v interface{}) {
	o.Manual = v
}

// GetSkuCodeRegex returns the SkuCodeRegex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHSkuListsSkuListIdRequestDataAttributes) GetSkuCodeRegex() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.SkuCodeRegex
}

// GetSkuCodeRegexOk returns a tuple with the SkuCodeRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHSkuListsSkuListIdRequestDataAttributes) GetSkuCodeRegexOk() (*interface{}, bool) {
	if o == nil || IsNil(o.SkuCodeRegex) {
		return nil, false
	}
	return &o.SkuCodeRegex, true
}

// HasSkuCodeRegex returns a boolean if a field has been set.
func (o *PATCHSkuListsSkuListIdRequestDataAttributes) HasSkuCodeRegex() bool {
	if o != nil && IsNil(o.SkuCodeRegex) {
		return true
	}

	return false
}

// SetSkuCodeRegex gets a reference to the given interface{} and assigns it to the SkuCodeRegex field.
func (o *PATCHSkuListsSkuListIdRequestDataAttributes) SetSkuCodeRegex(v interface{}) {
	o.SkuCodeRegex = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHSkuListsSkuListIdRequestDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHSkuListsSkuListIdRequestDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHSkuListsSkuListIdRequestDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *PATCHSkuListsSkuListIdRequestDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHSkuListsSkuListIdRequestDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHSkuListsSkuListIdRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHSkuListsSkuListIdRequestDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *PATCHSkuListsSkuListIdRequestDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHSkuListsSkuListIdRequestDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHSkuListsSkuListIdRequestDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHSkuListsSkuListIdRequestDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *PATCHSkuListsSkuListIdRequestDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o PATCHSkuListsSkuListIdRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHSkuListsSkuListIdRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ImageUrl != nil {
		toSerialize["image_url"] = o.ImageUrl
	}
	if o.Manual != nil {
		toSerialize["manual"] = o.Manual
	}
	if o.SkuCodeRegex != nil {
		toSerialize["sku_code_regex"] = o.SkuCodeRegex
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
	return toSerialize, nil
}

type NullablePATCHSkuListsSkuListIdRequestDataAttributes struct {
	value *PATCHSkuListsSkuListIdRequestDataAttributes
	isSet bool
}

func (v NullablePATCHSkuListsSkuListIdRequestDataAttributes) Get() *PATCHSkuListsSkuListIdRequestDataAttributes {
	return v.value
}

func (v *NullablePATCHSkuListsSkuListIdRequestDataAttributes) Set(val *PATCHSkuListsSkuListIdRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHSkuListsSkuListIdRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHSkuListsSkuListIdRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHSkuListsSkuListIdRequestDataAttributes(val *PATCHSkuListsSkuListIdRequestDataAttributes) *NullablePATCHSkuListsSkuListIdRequestDataAttributes {
	return &NullablePATCHSkuListsSkuListIdRequestDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHSkuListsSkuListIdRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHSkuListsSkuListIdRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
