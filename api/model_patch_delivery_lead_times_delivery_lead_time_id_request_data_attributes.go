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

// checks if the PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestDataAttributes{}

// PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestDataAttributes struct for PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestDataAttributes
type PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestDataAttributes struct {
	// The delivery lead minimum time (in hours) when shipping from the associated stock location with the associated shipping method.
	MinHours interface{} `json:"min_hours,omitempty"`
	// The delivery lead maximun time (in hours) when shipping from the associated stock location with the associated shipping method.
	MaxHours interface{} `json:"max_hours,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestDataAttributes instantiates a new PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestDataAttributes() *PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestDataAttributes {
	this := PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestDataAttributes{}
	return &this
}

// NewPATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestDataAttributesWithDefaults instantiates a new PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestDataAttributesWithDefaults() *PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestDataAttributes {
	this := PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestDataAttributes{}
	return &this
}

// GetMinHours returns the MinHours field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestDataAttributes) GetMinHours() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.MinHours
}

// GetMinHoursOk returns a tuple with the MinHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestDataAttributes) GetMinHoursOk() (*interface{}, bool) {
	if o == nil || IsNil(o.MinHours) {
		return nil, false
	}
	return &o.MinHours, true
}

// HasMinHours returns a boolean if a field has been set.
func (o *PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestDataAttributes) HasMinHours() bool {
	if o != nil && IsNil(o.MinHours) {
		return true
	}

	return false
}

// SetMinHours gets a reference to the given interface{} and assigns it to the MinHours field.
func (o *PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestDataAttributes) SetMinHours(v interface{}) {
	o.MinHours = v
}

// GetMaxHours returns the MaxHours field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestDataAttributes) GetMaxHours() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.MaxHours
}

// GetMaxHoursOk returns a tuple with the MaxHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestDataAttributes) GetMaxHoursOk() (*interface{}, bool) {
	if o == nil || IsNil(o.MaxHours) {
		return nil, false
	}
	return &o.MaxHours, true
}

// HasMaxHours returns a boolean if a field has been set.
func (o *PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestDataAttributes) HasMaxHours() bool {
	if o != nil && IsNil(o.MaxHours) {
		return true
	}

	return false
}

// SetMaxHours gets a reference to the given interface{} and assigns it to the MaxHours field.
func (o *PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestDataAttributes) SetMaxHours(v interface{}) {
	o.MaxHours = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.MinHours != nil {
		toSerialize["min_hours"] = o.MinHours
	}
	if o.MaxHours != nil {
		toSerialize["max_hours"] = o.MaxHours
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

type NullablePATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestDataAttributes struct {
	value *PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestDataAttributes
	isSet bool
}

func (v NullablePATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestDataAttributes) Get() *PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestDataAttributes {
	return v.value
}

func (v *NullablePATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestDataAttributes) Set(val *PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestDataAttributes(val *PATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestDataAttributes) *NullablePATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestDataAttributes {
	return &NullablePATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHDeliveryLeadTimesDeliveryLeadTimeIdRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
