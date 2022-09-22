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

// POSTImports201ResponseDataAttributes struct for POSTImports201ResponseDataAttributes
type POSTImports201ResponseDataAttributes struct {
	// The type of resource being imported.
	ResourceType string `json:"resource_type"`
	// The ID of the parent resource to be associated with imported data.
	ParentResourceId *string `json:"parent_resource_id,omitempty"`
	// Array of objects representing the resources that are being imported.
	Inputs []map[string]interface{} `json:"inputs"`
	// Indicates if the import should cleanup records that are not included in the inputs array.
	CleanupRecords *bool `json:"cleanup_records,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewPOSTImports201ResponseDataAttributes instantiates a new POSTImports201ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTImports201ResponseDataAttributes(resourceType string, inputs []map[string]interface{}) *POSTImports201ResponseDataAttributes {
	this := POSTImports201ResponseDataAttributes{}
	this.ResourceType = resourceType
	this.Inputs = inputs
	return &this
}

// NewPOSTImports201ResponseDataAttributesWithDefaults instantiates a new POSTImports201ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTImports201ResponseDataAttributesWithDefaults() *POSTImports201ResponseDataAttributes {
	this := POSTImports201ResponseDataAttributes{}
	return &this
}

// GetResourceType returns the ResourceType field value
func (o *POSTImports201ResponseDataAttributes) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *POSTImports201ResponseDataAttributes) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *POSTImports201ResponseDataAttributes) SetResourceType(v string) {
	o.ResourceType = v
}

// GetParentResourceId returns the ParentResourceId field value if set, zero value otherwise.
func (o *POSTImports201ResponseDataAttributes) GetParentResourceId() string {
	if o == nil || o.ParentResourceId == nil {
		var ret string
		return ret
	}
	return *o.ParentResourceId
}

// GetParentResourceIdOk returns a tuple with the ParentResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTImports201ResponseDataAttributes) GetParentResourceIdOk() (*string, bool) {
	if o == nil || o.ParentResourceId == nil {
		return nil, false
	}
	return o.ParentResourceId, true
}

// HasParentResourceId returns a boolean if a field has been set.
func (o *POSTImports201ResponseDataAttributes) HasParentResourceId() bool {
	if o != nil && o.ParentResourceId != nil {
		return true
	}

	return false
}

// SetParentResourceId gets a reference to the given string and assigns it to the ParentResourceId field.
func (o *POSTImports201ResponseDataAttributes) SetParentResourceId(v string) {
	o.ParentResourceId = &v
}

// GetInputs returns the Inputs field value
func (o *POSTImports201ResponseDataAttributes) GetInputs() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *POSTImports201ResponseDataAttributes) GetInputsOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Inputs, true
}

// SetInputs sets field value
func (o *POSTImports201ResponseDataAttributes) SetInputs(v []map[string]interface{}) {
	o.Inputs = v
}

// GetCleanupRecords returns the CleanupRecords field value if set, zero value otherwise.
func (o *POSTImports201ResponseDataAttributes) GetCleanupRecords() bool {
	if o == nil || o.CleanupRecords == nil {
		var ret bool
		return ret
	}
	return *o.CleanupRecords
}

// GetCleanupRecordsOk returns a tuple with the CleanupRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTImports201ResponseDataAttributes) GetCleanupRecordsOk() (*bool, bool) {
	if o == nil || o.CleanupRecords == nil {
		return nil, false
	}
	return o.CleanupRecords, true
}

// HasCleanupRecords returns a boolean if a field has been set.
func (o *POSTImports201ResponseDataAttributes) HasCleanupRecords() bool {
	if o != nil && o.CleanupRecords != nil {
		return true
	}

	return false
}

// SetCleanupRecords gets a reference to the given bool and assigns it to the CleanupRecords field.
func (o *POSTImports201ResponseDataAttributes) SetCleanupRecords(v bool) {
	o.CleanupRecords = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *POSTImports201ResponseDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTImports201ResponseDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTImports201ResponseDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *POSTImports201ResponseDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *POSTImports201ResponseDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTImports201ResponseDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTImports201ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *POSTImports201ResponseDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *POSTImports201ResponseDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTImports201ResponseDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTImports201ResponseDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *POSTImports201ResponseDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o POSTImports201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["resource_type"] = o.ResourceType
	}
	if o.ParentResourceId != nil {
		toSerialize["parent_resource_id"] = o.ParentResourceId
	}
	if true {
		toSerialize["inputs"] = o.Inputs
	}
	if o.CleanupRecords != nil {
		toSerialize["cleanup_records"] = o.CleanupRecords
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

type NullablePOSTImports201ResponseDataAttributes struct {
	value *POSTImports201ResponseDataAttributes
	isSet bool
}

func (v NullablePOSTImports201ResponseDataAttributes) Get() *POSTImports201ResponseDataAttributes {
	return v.value
}

func (v *NullablePOSTImports201ResponseDataAttributes) Set(val *POSTImports201ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTImports201ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTImports201ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTImports201ResponseDataAttributes(val *POSTImports201ResponseDataAttributes) *NullablePOSTImports201ResponseDataAttributes {
	return &NullablePOSTImports201ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTImports201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTImports201ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
