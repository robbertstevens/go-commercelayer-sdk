/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.4.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETOrderCopies200ResponseDataInnerAttributes struct for GETOrderCopies200ResponseDataInnerAttributes
type GETOrderCopies200ResponseDataInnerAttributes struct {
	// The order copy status. One of 'pending' (default), 'in_progress', 'failed', or 'completed'.
	Status *string `json:"status,omitempty"`
	// Time at which the order copy was started.
	StartedAt *string `json:"started_at,omitempty"`
	// Time at which the order copy was completed.
	CompletedAt *string `json:"completed_at,omitempty"`
	// Time at which the order copy has failed.
	FailedAt *string `json:"failed_at,omitempty"`
	// Indicates if the target order must be placed upon copy.
	PlaceTargetOrder *bool `json:"place_target_order,omitempty"`
	// Indicates if the source order must be cancelled upon copy.
	CancelSourceOrder *bool `json:"cancel_source_order,omitempty"`
	// Indicates if the payment source within the source order customer's wallet must be copied.
	ReuseWallet *bool `json:"reuse_wallet,omitempty"`
	// Contains the order copy errors, if any.
	ErrorsLog map[string]interface{} `json:"errors_log,omitempty"`
	// Indicates the number of copy errors, if any.
	ErrorsCount *int32 `json:"errors_count,omitempty"`
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

// NewGETOrderCopies200ResponseDataInnerAttributes instantiates a new GETOrderCopies200ResponseDataInnerAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETOrderCopies200ResponseDataInnerAttributes() *GETOrderCopies200ResponseDataInnerAttributes {
	this := GETOrderCopies200ResponseDataInnerAttributes{}
	return &this
}

// NewGETOrderCopies200ResponseDataInnerAttributesWithDefaults instantiates a new GETOrderCopies200ResponseDataInnerAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETOrderCopies200ResponseDataInnerAttributesWithDefaults() *GETOrderCopies200ResponseDataInnerAttributes {
	this := GETOrderCopies200ResponseDataInnerAttributes{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GETOrderCopies200ResponseDataInnerAttributes) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrderCopies200ResponseDataInnerAttributes) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GETOrderCopies200ResponseDataInnerAttributes) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GETOrderCopies200ResponseDataInnerAttributes) SetStatus(v string) {
	o.Status = &v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *GETOrderCopies200ResponseDataInnerAttributes) GetStartedAt() string {
	if o == nil || o.StartedAt == nil {
		var ret string
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrderCopies200ResponseDataInnerAttributes) GetStartedAtOk() (*string, bool) {
	if o == nil || o.StartedAt == nil {
		return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *GETOrderCopies200ResponseDataInnerAttributes) HasStartedAt() bool {
	if o != nil && o.StartedAt != nil {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given string and assigns it to the StartedAt field.
func (o *GETOrderCopies200ResponseDataInnerAttributes) SetStartedAt(v string) {
	o.StartedAt = &v
}

// GetCompletedAt returns the CompletedAt field value if set, zero value otherwise.
func (o *GETOrderCopies200ResponseDataInnerAttributes) GetCompletedAt() string {
	if o == nil || o.CompletedAt == nil {
		var ret string
		return ret
	}
	return *o.CompletedAt
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrderCopies200ResponseDataInnerAttributes) GetCompletedAtOk() (*string, bool) {
	if o == nil || o.CompletedAt == nil {
		return nil, false
	}
	return o.CompletedAt, true
}

// HasCompletedAt returns a boolean if a field has been set.
func (o *GETOrderCopies200ResponseDataInnerAttributes) HasCompletedAt() bool {
	if o != nil && o.CompletedAt != nil {
		return true
	}

	return false
}

// SetCompletedAt gets a reference to the given string and assigns it to the CompletedAt field.
func (o *GETOrderCopies200ResponseDataInnerAttributes) SetCompletedAt(v string) {
	o.CompletedAt = &v
}

// GetFailedAt returns the FailedAt field value if set, zero value otherwise.
func (o *GETOrderCopies200ResponseDataInnerAttributes) GetFailedAt() string {
	if o == nil || o.FailedAt == nil {
		var ret string
		return ret
	}
	return *o.FailedAt
}

// GetFailedAtOk returns a tuple with the FailedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrderCopies200ResponseDataInnerAttributes) GetFailedAtOk() (*string, bool) {
	if o == nil || o.FailedAt == nil {
		return nil, false
	}
	return o.FailedAt, true
}

// HasFailedAt returns a boolean if a field has been set.
func (o *GETOrderCopies200ResponseDataInnerAttributes) HasFailedAt() bool {
	if o != nil && o.FailedAt != nil {
		return true
	}

	return false
}

// SetFailedAt gets a reference to the given string and assigns it to the FailedAt field.
func (o *GETOrderCopies200ResponseDataInnerAttributes) SetFailedAt(v string) {
	o.FailedAt = &v
}

// GetPlaceTargetOrder returns the PlaceTargetOrder field value if set, zero value otherwise.
func (o *GETOrderCopies200ResponseDataInnerAttributes) GetPlaceTargetOrder() bool {
	if o == nil || o.PlaceTargetOrder == nil {
		var ret bool
		return ret
	}
	return *o.PlaceTargetOrder
}

// GetPlaceTargetOrderOk returns a tuple with the PlaceTargetOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrderCopies200ResponseDataInnerAttributes) GetPlaceTargetOrderOk() (*bool, bool) {
	if o == nil || o.PlaceTargetOrder == nil {
		return nil, false
	}
	return o.PlaceTargetOrder, true
}

// HasPlaceTargetOrder returns a boolean if a field has been set.
func (o *GETOrderCopies200ResponseDataInnerAttributes) HasPlaceTargetOrder() bool {
	if o != nil && o.PlaceTargetOrder != nil {
		return true
	}

	return false
}

// SetPlaceTargetOrder gets a reference to the given bool and assigns it to the PlaceTargetOrder field.
func (o *GETOrderCopies200ResponseDataInnerAttributes) SetPlaceTargetOrder(v bool) {
	o.PlaceTargetOrder = &v
}

// GetCancelSourceOrder returns the CancelSourceOrder field value if set, zero value otherwise.
func (o *GETOrderCopies200ResponseDataInnerAttributes) GetCancelSourceOrder() bool {
	if o == nil || o.CancelSourceOrder == nil {
		var ret bool
		return ret
	}
	return *o.CancelSourceOrder
}

// GetCancelSourceOrderOk returns a tuple with the CancelSourceOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrderCopies200ResponseDataInnerAttributes) GetCancelSourceOrderOk() (*bool, bool) {
	if o == nil || o.CancelSourceOrder == nil {
		return nil, false
	}
	return o.CancelSourceOrder, true
}

// HasCancelSourceOrder returns a boolean if a field has been set.
func (o *GETOrderCopies200ResponseDataInnerAttributes) HasCancelSourceOrder() bool {
	if o != nil && o.CancelSourceOrder != nil {
		return true
	}

	return false
}

// SetCancelSourceOrder gets a reference to the given bool and assigns it to the CancelSourceOrder field.
func (o *GETOrderCopies200ResponseDataInnerAttributes) SetCancelSourceOrder(v bool) {
	o.CancelSourceOrder = &v
}

// GetReuseWallet returns the ReuseWallet field value if set, zero value otherwise.
func (o *GETOrderCopies200ResponseDataInnerAttributes) GetReuseWallet() bool {
	if o == nil || o.ReuseWallet == nil {
		var ret bool
		return ret
	}
	return *o.ReuseWallet
}

// GetReuseWalletOk returns a tuple with the ReuseWallet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrderCopies200ResponseDataInnerAttributes) GetReuseWalletOk() (*bool, bool) {
	if o == nil || o.ReuseWallet == nil {
		return nil, false
	}
	return o.ReuseWallet, true
}

// HasReuseWallet returns a boolean if a field has been set.
func (o *GETOrderCopies200ResponseDataInnerAttributes) HasReuseWallet() bool {
	if o != nil && o.ReuseWallet != nil {
		return true
	}

	return false
}

// SetReuseWallet gets a reference to the given bool and assigns it to the ReuseWallet field.
func (o *GETOrderCopies200ResponseDataInnerAttributes) SetReuseWallet(v bool) {
	o.ReuseWallet = &v
}

// GetErrorsLog returns the ErrorsLog field value if set, zero value otherwise.
func (o *GETOrderCopies200ResponseDataInnerAttributes) GetErrorsLog() map[string]interface{} {
	if o == nil || o.ErrorsLog == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ErrorsLog
}

// GetErrorsLogOk returns a tuple with the ErrorsLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrderCopies200ResponseDataInnerAttributes) GetErrorsLogOk() (map[string]interface{}, bool) {
	if o == nil || o.ErrorsLog == nil {
		return nil, false
	}
	return o.ErrorsLog, true
}

// HasErrorsLog returns a boolean if a field has been set.
func (o *GETOrderCopies200ResponseDataInnerAttributes) HasErrorsLog() bool {
	if o != nil && o.ErrorsLog != nil {
		return true
	}

	return false
}

// SetErrorsLog gets a reference to the given map[string]interface{} and assigns it to the ErrorsLog field.
func (o *GETOrderCopies200ResponseDataInnerAttributes) SetErrorsLog(v map[string]interface{}) {
	o.ErrorsLog = v
}

// GetErrorsCount returns the ErrorsCount field value if set, zero value otherwise.
func (o *GETOrderCopies200ResponseDataInnerAttributes) GetErrorsCount() int32 {
	if o == nil || o.ErrorsCount == nil {
		var ret int32
		return ret
	}
	return *o.ErrorsCount
}

// GetErrorsCountOk returns a tuple with the ErrorsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrderCopies200ResponseDataInnerAttributes) GetErrorsCountOk() (*int32, bool) {
	if o == nil || o.ErrorsCount == nil {
		return nil, false
	}
	return o.ErrorsCount, true
}

// HasErrorsCount returns a boolean if a field has been set.
func (o *GETOrderCopies200ResponseDataInnerAttributes) HasErrorsCount() bool {
	if o != nil && o.ErrorsCount != nil {
		return true
	}

	return false
}

// SetErrorsCount gets a reference to the given int32 and assigns it to the ErrorsCount field.
func (o *GETOrderCopies200ResponseDataInnerAttributes) SetErrorsCount(v int32) {
	o.ErrorsCount = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GETOrderCopies200ResponseDataInnerAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrderCopies200ResponseDataInnerAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETOrderCopies200ResponseDataInnerAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *GETOrderCopies200ResponseDataInnerAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GETOrderCopies200ResponseDataInnerAttributes) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrderCopies200ResponseDataInnerAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETOrderCopies200ResponseDataInnerAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *GETOrderCopies200ResponseDataInnerAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *GETOrderCopies200ResponseDataInnerAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrderCopies200ResponseDataInnerAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETOrderCopies200ResponseDataInnerAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *GETOrderCopies200ResponseDataInnerAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *GETOrderCopies200ResponseDataInnerAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrderCopies200ResponseDataInnerAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETOrderCopies200ResponseDataInnerAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *GETOrderCopies200ResponseDataInnerAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GETOrderCopies200ResponseDataInnerAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrderCopies200ResponseDataInnerAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETOrderCopies200ResponseDataInnerAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *GETOrderCopies200ResponseDataInnerAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o GETOrderCopies200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.StartedAt != nil {
		toSerialize["started_at"] = o.StartedAt
	}
	if o.CompletedAt != nil {
		toSerialize["completed_at"] = o.CompletedAt
	}
	if o.FailedAt != nil {
		toSerialize["failed_at"] = o.FailedAt
	}
	if o.PlaceTargetOrder != nil {
		toSerialize["place_target_order"] = o.PlaceTargetOrder
	}
	if o.CancelSourceOrder != nil {
		toSerialize["cancel_source_order"] = o.CancelSourceOrder
	}
	if o.ReuseWallet != nil {
		toSerialize["reuse_wallet"] = o.ReuseWallet
	}
	if o.ErrorsLog != nil {
		toSerialize["errors_log"] = o.ErrorsLog
	}
	if o.ErrorsCount != nil {
		toSerialize["errors_count"] = o.ErrorsCount
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

type NullableGETOrderCopies200ResponseDataInnerAttributes struct {
	value *GETOrderCopies200ResponseDataInnerAttributes
	isSet bool
}

func (v NullableGETOrderCopies200ResponseDataInnerAttributes) Get() *GETOrderCopies200ResponseDataInnerAttributes {
	return v.value
}

func (v *NullableGETOrderCopies200ResponseDataInnerAttributes) Set(val *GETOrderCopies200ResponseDataInnerAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETOrderCopies200ResponseDataInnerAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETOrderCopies200ResponseDataInnerAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETOrderCopies200ResponseDataInnerAttributes(val *GETOrderCopies200ResponseDataInnerAttributes) *NullableGETOrderCopies200ResponseDataInnerAttributes {
	return &NullableGETOrderCopies200ResponseDataInnerAttributes{value: val, isSet: true}
}

func (v NullableGETOrderCopies200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETOrderCopies200ResponseDataInnerAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
