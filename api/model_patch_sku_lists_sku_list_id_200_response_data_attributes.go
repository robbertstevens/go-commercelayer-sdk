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

// PATCHSkuListsSkuListId200ResponseDataAttributes struct for PATCHSkuListsSkuListId200ResponseDataAttributes
type PATCHSkuListsSkuListId200ResponseDataAttributes struct {
	// The SKU list's internal name.
	Name *string `json:"name,omitempty"`
	// An internal description of the SKU list.
	Description *string `json:"description,omitempty"`
	// The URL of an image that represents the SKU list.
	ImageUrl *string `json:"image_url,omitempty"`
	// Indicates if the SKU list is populated manually.
	Manual *bool `json:"manual,omitempty"`
	// The regex that will be evaluated to match the SKU codes.
	SkuCodeRegex *string `json:"sku_code_regex,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewPATCHSkuListsSkuListId200ResponseDataAttributes instantiates a new PATCHSkuListsSkuListId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHSkuListsSkuListId200ResponseDataAttributes() *PATCHSkuListsSkuListId200ResponseDataAttributes {
	this := PATCHSkuListsSkuListId200ResponseDataAttributes{}
	return &this
}

// NewPATCHSkuListsSkuListId200ResponseDataAttributesWithDefaults instantiates a new PATCHSkuListsSkuListId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHSkuListsSkuListId200ResponseDataAttributesWithDefaults() *PATCHSkuListsSkuListId200ResponseDataAttributes {
	this := PATCHSkuListsSkuListId200ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PATCHSkuListsSkuListId200ResponseDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHSkuListsSkuListId200ResponseDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PATCHSkuListsSkuListId200ResponseDataAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PATCHSkuListsSkuListId200ResponseDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PATCHSkuListsSkuListId200ResponseDataAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHSkuListsSkuListId200ResponseDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PATCHSkuListsSkuListId200ResponseDataAttributes) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PATCHSkuListsSkuListId200ResponseDataAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *PATCHSkuListsSkuListId200ResponseDataAttributes) GetImageUrl() string {
	if o == nil || o.ImageUrl == nil {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHSkuListsSkuListId200ResponseDataAttributes) GetImageUrlOk() (*string, bool) {
	if o == nil || o.ImageUrl == nil {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *PATCHSkuListsSkuListId200ResponseDataAttributes) HasImageUrl() bool {
	if o != nil && o.ImageUrl != nil {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *PATCHSkuListsSkuListId200ResponseDataAttributes) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetManual returns the Manual field value if set, zero value otherwise.
func (o *PATCHSkuListsSkuListId200ResponseDataAttributes) GetManual() bool {
	if o == nil || o.Manual == nil {
		var ret bool
		return ret
	}
	return *o.Manual
}

// GetManualOk returns a tuple with the Manual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHSkuListsSkuListId200ResponseDataAttributes) GetManualOk() (*bool, bool) {
	if o == nil || o.Manual == nil {
		return nil, false
	}
	return o.Manual, true
}

// HasManual returns a boolean if a field has been set.
func (o *PATCHSkuListsSkuListId200ResponseDataAttributes) HasManual() bool {
	if o != nil && o.Manual != nil {
		return true
	}

	return false
}

// SetManual gets a reference to the given bool and assigns it to the Manual field.
func (o *PATCHSkuListsSkuListId200ResponseDataAttributes) SetManual(v bool) {
	o.Manual = &v
}

// GetSkuCodeRegex returns the SkuCodeRegex field value if set, zero value otherwise.
func (o *PATCHSkuListsSkuListId200ResponseDataAttributes) GetSkuCodeRegex() string {
	if o == nil || o.SkuCodeRegex == nil {
		var ret string
		return ret
	}
	return *o.SkuCodeRegex
}

// GetSkuCodeRegexOk returns a tuple with the SkuCodeRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHSkuListsSkuListId200ResponseDataAttributes) GetSkuCodeRegexOk() (*string, bool) {
	if o == nil || o.SkuCodeRegex == nil {
		return nil, false
	}
	return o.SkuCodeRegex, true
}

// HasSkuCodeRegex returns a boolean if a field has been set.
func (o *PATCHSkuListsSkuListId200ResponseDataAttributes) HasSkuCodeRegex() bool {
	if o != nil && o.SkuCodeRegex != nil {
		return true
	}

	return false
}

// SetSkuCodeRegex gets a reference to the given string and assigns it to the SkuCodeRegex field.
func (o *PATCHSkuListsSkuListId200ResponseDataAttributes) SetSkuCodeRegex(v string) {
	o.SkuCodeRegex = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *PATCHSkuListsSkuListId200ResponseDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHSkuListsSkuListId200ResponseDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHSkuListsSkuListId200ResponseDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *PATCHSkuListsSkuListId200ResponseDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *PATCHSkuListsSkuListId200ResponseDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHSkuListsSkuListId200ResponseDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHSkuListsSkuListId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *PATCHSkuListsSkuListId200ResponseDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PATCHSkuListsSkuListId200ResponseDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHSkuListsSkuListId200ResponseDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHSkuListsSkuListId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PATCHSkuListsSkuListId200ResponseDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o PATCHSkuListsSkuListId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullablePATCHSkuListsSkuListId200ResponseDataAttributes struct {
	value *PATCHSkuListsSkuListId200ResponseDataAttributes
	isSet bool
}

func (v NullablePATCHSkuListsSkuListId200ResponseDataAttributes) Get() *PATCHSkuListsSkuListId200ResponseDataAttributes {
	return v.value
}

func (v *NullablePATCHSkuListsSkuListId200ResponseDataAttributes) Set(val *PATCHSkuListsSkuListId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHSkuListsSkuListId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHSkuListsSkuListId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHSkuListsSkuListId200ResponseDataAttributes(val *PATCHSkuListsSkuListId200ResponseDataAttributes) *NullablePATCHSkuListsSkuListId200ResponseDataAttributes {
	return &NullablePATCHSkuListsSkuListId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHSkuListsSkuListId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHSkuListsSkuListId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
