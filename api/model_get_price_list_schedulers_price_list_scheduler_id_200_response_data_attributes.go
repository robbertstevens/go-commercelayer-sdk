/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.3.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes{}

// GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes struct for GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes
type GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes struct {
	// The price list scheduler's internal name.
	Name interface{} `json:"name,omitempty"`
	// The activation date/time of this price list scheduler.
	StartsAt interface{} `json:"starts_at,omitempty"`
	// The expiration date/time of this price list scheduler (must be after starts_at).
	ExpiresAt interface{} `json:"expires_at,omitempty"`
	// Indicates if the price list scheduler is active (enabled and not expired).
	Active interface{} `json:"active,omitempty"`
	// The price list scheduler status. One of 'disabled', 'expired', 'pending', or 'active'.
	Status interface{} `json:"status,omitempty"`
	// Time at which this resource was disabled.
	DisabledAt interface{} `json:"disabled_at,omitempty"`
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
}

// NewGETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes instantiates a new GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes() *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes {
	this := GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes{}
	return &this
}

// NewGETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributesWithDefaults instantiates a new GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributesWithDefaults() *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes {
	this := GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) HasName() bool {
	if o != nil && IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given interface{} and assigns it to the Name field.
func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetStartsAt returns the StartsAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) GetStartsAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.StartsAt
}

// GetStartsAtOk returns a tuple with the StartsAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) GetStartsAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.StartsAt) {
		return nil, false
	}
	return &o.StartsAt, true
}

// HasStartsAt returns a boolean if a field has been set.
func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) HasStartsAt() bool {
	if o != nil && IsNil(o.StartsAt) {
		return true
	}

	return false
}

// SetStartsAt gets a reference to the given interface{} and assigns it to the StartsAt field.
func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) SetStartsAt(v interface{}) {
	o.StartsAt = v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) GetExpiresAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) GetExpiresAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) HasExpiresAt() bool {
	if o != nil && IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given interface{} and assigns it to the ExpiresAt field.
func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) SetExpiresAt(v interface{}) {
	o.ExpiresAt = v
}

// GetActive returns the Active field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) GetActive() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) GetActiveOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return &o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) HasActive() bool {
	if o != nil && IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given interface{} and assigns it to the Active field.
func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) SetActive(v interface{}) {
	o.Active = v
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) GetStatus() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) GetStatusOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return &o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) HasStatus() bool {
	if o != nil && IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given interface{} and assigns it to the Status field.
func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) SetStatus(v interface{}) {
	o.Status = v
}

// GetDisabledAt returns the DisabledAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) GetDisabledAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.DisabledAt
}

// GetDisabledAtOk returns a tuple with the DisabledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) GetDisabledAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.DisabledAt) {
		return nil, false
	}
	return &o.DisabledAt, true
}

// HasDisabledAt returns a boolean if a field has been set.
func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) HasDisabledAt() bool {
	if o != nil && IsNil(o.DisabledAt) {
		return true
	}

	return false
}

// SetDisabledAt gets a reference to the given interface{} and assigns it to the DisabledAt field.
func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) SetDisabledAt(v interface{}) {
	o.DisabledAt = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) GetCreatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return &o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) HasCreatedAt() bool {
	if o != nil && IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given interface{} and assigns it to the CreatedAt field.
func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) SetCreatedAt(v interface{}) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) GetUpdatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) HasUpdatedAt() bool {
	if o != nil && IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given interface{} and assigns it to the UpdatedAt field.
func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) SetUpdatedAt(v interface{}) {
	o.UpdatedAt = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.StartsAt != nil {
		toSerialize["starts_at"] = o.StartsAt
	}
	if o.ExpiresAt != nil {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.DisabledAt != nil {
		toSerialize["disabled_at"] = o.DisabledAt
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
	return toSerialize, nil
}

type NullableGETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes struct {
	value *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes
	isSet bool
}

func (v NullableGETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) Get() *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes {
	return v.value
}

func (v *NullableGETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) Set(val *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes(val *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) *NullableGETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes {
	return &NullableGETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullableGETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
