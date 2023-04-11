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

// checks if the POSTInventoryReturnLocationsRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTInventoryReturnLocationsRequestDataAttributes{}

// POSTInventoryReturnLocationsRequestDataAttributes struct for POSTInventoryReturnLocationsRequestDataAttributes
type POSTInventoryReturnLocationsRequestDataAttributes struct {
	// The inventory return location priority within the associated invetory model.
	Priority interface{} `json:"priority"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPOSTInventoryReturnLocationsRequestDataAttributes instantiates a new POSTInventoryReturnLocationsRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTInventoryReturnLocationsRequestDataAttributes(priority interface{}) *POSTInventoryReturnLocationsRequestDataAttributes {
	this := POSTInventoryReturnLocationsRequestDataAttributes{}
	this.Priority = priority
	return &this
}

// NewPOSTInventoryReturnLocationsRequestDataAttributesWithDefaults instantiates a new POSTInventoryReturnLocationsRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTInventoryReturnLocationsRequestDataAttributesWithDefaults() *POSTInventoryReturnLocationsRequestDataAttributes {
	this := POSTInventoryReturnLocationsRequestDataAttributes{}
	return &this
}

// GetPriority returns the Priority field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTInventoryReturnLocationsRequestDataAttributes) GetPriority() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTInventoryReturnLocationsRequestDataAttributes) GetPriorityOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *POSTInventoryReturnLocationsRequestDataAttributes) SetPriority(v interface{}) {
	o.Priority = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTInventoryReturnLocationsRequestDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTInventoryReturnLocationsRequestDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTInventoryReturnLocationsRequestDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *POSTInventoryReturnLocationsRequestDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTInventoryReturnLocationsRequestDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTInventoryReturnLocationsRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTInventoryReturnLocationsRequestDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *POSTInventoryReturnLocationsRequestDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTInventoryReturnLocationsRequestDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTInventoryReturnLocationsRequestDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTInventoryReturnLocationsRequestDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *POSTInventoryReturnLocationsRequestDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o POSTInventoryReturnLocationsRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTInventoryReturnLocationsRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
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

type NullablePOSTInventoryReturnLocationsRequestDataAttributes struct {
	value *POSTInventoryReturnLocationsRequestDataAttributes
	isSet bool
}

func (v NullablePOSTInventoryReturnLocationsRequestDataAttributes) Get() *POSTInventoryReturnLocationsRequestDataAttributes {
	return v.value
}

func (v *NullablePOSTInventoryReturnLocationsRequestDataAttributes) Set(val *POSTInventoryReturnLocationsRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTInventoryReturnLocationsRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTInventoryReturnLocationsRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTInventoryReturnLocationsRequestDataAttributes(val *POSTInventoryReturnLocationsRequestDataAttributes) *NullablePOSTInventoryReturnLocationsRequestDataAttributes {
	return &NullablePOSTInventoryReturnLocationsRequestDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTInventoryReturnLocationsRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTInventoryReturnLocationsRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
