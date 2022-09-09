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

// GETDeliveryLeadTimes200ResponseDataInnerAttributes struct for GETDeliveryLeadTimes200ResponseDataInnerAttributes
type GETDeliveryLeadTimes200ResponseDataInnerAttributes struct {
	// The delivery lead minimum time (in hours) when shipping from the associated stock location with the associated shipping method.
	MinHours *int32 `json:"min_hours,omitempty"`
	// The delivery lead maximun time (in hours) when shipping from the associated stock location with the associated shipping method.
	MaxHours *int32 `json:"max_hours,omitempty"`
	// The delivery lead minimum time, in days (rounded)
	MinDays *int32 `json:"min_days,omitempty"`
	// The delivery lead maximun time, in days (rounded)
	MaxDays *int32 `json:"max_days,omitempty"`
	// Unique identifier for the resource (hash).
	Id *string `json:"id,omitempty"`
	// Time at which the resource was created.
	CreatedAt *string `json:"created_at,omitempty"`
	// Time at which the resource was last updated.
	UpdatedAt *string `json:"updated_at,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewGETDeliveryLeadTimes200ResponseDataInnerAttributes instantiates a new GETDeliveryLeadTimes200ResponseDataInnerAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETDeliveryLeadTimes200ResponseDataInnerAttributes() *GETDeliveryLeadTimes200ResponseDataInnerAttributes {
	this := GETDeliveryLeadTimes200ResponseDataInnerAttributes{}
	return &this
}

// NewGETDeliveryLeadTimes200ResponseDataInnerAttributesWithDefaults instantiates a new GETDeliveryLeadTimes200ResponseDataInnerAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETDeliveryLeadTimes200ResponseDataInnerAttributesWithDefaults() *GETDeliveryLeadTimes200ResponseDataInnerAttributes {
	this := GETDeliveryLeadTimes200ResponseDataInnerAttributes{}
	return &this
}

// GetMinHours returns the MinHours field value if set, zero value otherwise.
func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) GetMinHours() int32 {
	if o == nil || o.MinHours == nil {
		var ret int32
		return ret
	}
	return *o.MinHours
}

// GetMinHoursOk returns a tuple with the MinHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) GetMinHoursOk() (*int32, bool) {
	if o == nil || o.MinHours == nil {
		return nil, false
	}
	return o.MinHours, true
}

// HasMinHours returns a boolean if a field has been set.
func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) HasMinHours() bool {
	if o != nil && o.MinHours != nil {
		return true
	}

	return false
}

// SetMinHours gets a reference to the given int32 and assigns it to the MinHours field.
func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) SetMinHours(v int32) {
	o.MinHours = &v
}

// GetMaxHours returns the MaxHours field value if set, zero value otherwise.
func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) GetMaxHours() int32 {
	if o == nil || o.MaxHours == nil {
		var ret int32
		return ret
	}
	return *o.MaxHours
}

// GetMaxHoursOk returns a tuple with the MaxHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) GetMaxHoursOk() (*int32, bool) {
	if o == nil || o.MaxHours == nil {
		return nil, false
	}
	return o.MaxHours, true
}

// HasMaxHours returns a boolean if a field has been set.
func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) HasMaxHours() bool {
	if o != nil && o.MaxHours != nil {
		return true
	}

	return false
}

// SetMaxHours gets a reference to the given int32 and assigns it to the MaxHours field.
func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) SetMaxHours(v int32) {
	o.MaxHours = &v
}

// GetMinDays returns the MinDays field value if set, zero value otherwise.
func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) GetMinDays() int32 {
	if o == nil || o.MinDays == nil {
		var ret int32
		return ret
	}
	return *o.MinDays
}

// GetMinDaysOk returns a tuple with the MinDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) GetMinDaysOk() (*int32, bool) {
	if o == nil || o.MinDays == nil {
		return nil, false
	}
	return o.MinDays, true
}

// HasMinDays returns a boolean if a field has been set.
func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) HasMinDays() bool {
	if o != nil && o.MinDays != nil {
		return true
	}

	return false
}

// SetMinDays gets a reference to the given int32 and assigns it to the MinDays field.
func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) SetMinDays(v int32) {
	o.MinDays = &v
}

// GetMaxDays returns the MaxDays field value if set, zero value otherwise.
func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) GetMaxDays() int32 {
	if o == nil || o.MaxDays == nil {
		var ret int32
		return ret
	}
	return *o.MaxDays
}

// GetMaxDaysOk returns a tuple with the MaxDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) GetMaxDaysOk() (*int32, bool) {
	if o == nil || o.MaxDays == nil {
		return nil, false
	}
	return o.MaxDays, true
}

// HasMaxDays returns a boolean if a field has been set.
func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) HasMaxDays() bool {
	if o != nil && o.MaxDays != nil {
		return true
	}

	return false
}

// SetMaxDays gets a reference to the given int32 and assigns it to the MaxDays field.
func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) SetMaxDays(v int32) {
	o.MaxDays = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o GETDeliveryLeadTimes200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MinHours != nil {
		toSerialize["min_hours"] = o.MinHours
	}
	if o.MaxHours != nil {
		toSerialize["max_hours"] = o.MaxHours
	}
	if o.MinDays != nil {
		toSerialize["min_days"] = o.MinDays
	}
	if o.MaxDays != nil {
		toSerialize["max_days"] = o.MaxDays
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
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
	return json.Marshal(toSerialize)
}

type NullableGETDeliveryLeadTimes200ResponseDataInnerAttributes struct {
	value *GETDeliveryLeadTimes200ResponseDataInnerAttributes
	isSet bool
}

func (v NullableGETDeliveryLeadTimes200ResponseDataInnerAttributes) Get() *GETDeliveryLeadTimes200ResponseDataInnerAttributes {
	return v.value
}

func (v *NullableGETDeliveryLeadTimes200ResponseDataInnerAttributes) Set(val *GETDeliveryLeadTimes200ResponseDataInnerAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETDeliveryLeadTimes200ResponseDataInnerAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETDeliveryLeadTimes200ResponseDataInnerAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETDeliveryLeadTimes200ResponseDataInnerAttributes(val *GETDeliveryLeadTimes200ResponseDataInnerAttributes) *NullableGETDeliveryLeadTimes200ResponseDataInnerAttributes {
	return &NullableGETDeliveryLeadTimes200ResponseDataInnerAttributes{value: val, isSet: true}
}

func (v NullableGETDeliveryLeadTimes200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETDeliveryLeadTimes200ResponseDataInnerAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


