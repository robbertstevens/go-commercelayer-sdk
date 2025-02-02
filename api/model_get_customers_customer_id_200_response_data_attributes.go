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

// checks if the GETCustomersCustomerId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETCustomersCustomerId200ResponseDataAttributes{}

// GETCustomersCustomerId200ResponseDataAttributes struct for GETCustomersCustomerId200ResponseDataAttributes
type GETCustomersCustomerId200ResponseDataAttributes struct {
	// The customer's email address.
	Email interface{} `json:"email,omitempty"`
	// The customer's status. One of 'prospect' (default), 'acquired', or 'repeat'.
	Status interface{} `json:"status,omitempty"`
	// Indicates if the customer has a password.
	HasPassword interface{} `json:"has_password,omitempty"`
	// The total number of orders for the customer.
	TotalOrdersCount interface{} `json:"total_orders_count,omitempty"`
	// A reference to uniquely identify the shopper during payment sessions.
	ShopperReference interface{} `json:"shopper_reference,omitempty"`
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

// NewGETCustomersCustomerId200ResponseDataAttributes instantiates a new GETCustomersCustomerId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCustomersCustomerId200ResponseDataAttributes() *GETCustomersCustomerId200ResponseDataAttributes {
	this := GETCustomersCustomerId200ResponseDataAttributes{}
	return &this
}

// NewGETCustomersCustomerId200ResponseDataAttributesWithDefaults instantiates a new GETCustomersCustomerId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCustomersCustomerId200ResponseDataAttributesWithDefaults() *GETCustomersCustomerId200ResponseDataAttributes {
	this := GETCustomersCustomerId200ResponseDataAttributes{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCustomersCustomerId200ResponseDataAttributes) GetEmail() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCustomersCustomerId200ResponseDataAttributes) GetEmailOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return &o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *GETCustomersCustomerId200ResponseDataAttributes) HasEmail() bool {
	if o != nil && IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given interface{} and assigns it to the Email field.
func (o *GETCustomersCustomerId200ResponseDataAttributes) SetEmail(v interface{}) {
	o.Email = v
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCustomersCustomerId200ResponseDataAttributes) GetStatus() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCustomersCustomerId200ResponseDataAttributes) GetStatusOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return &o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GETCustomersCustomerId200ResponseDataAttributes) HasStatus() bool {
	if o != nil && IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given interface{} and assigns it to the Status field.
func (o *GETCustomersCustomerId200ResponseDataAttributes) SetStatus(v interface{}) {
	o.Status = v
}

// GetHasPassword returns the HasPassword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCustomersCustomerId200ResponseDataAttributes) GetHasPassword() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.HasPassword
}

// GetHasPasswordOk returns a tuple with the HasPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCustomersCustomerId200ResponseDataAttributes) GetHasPasswordOk() (*interface{}, bool) {
	if o == nil || IsNil(o.HasPassword) {
		return nil, false
	}
	return &o.HasPassword, true
}

// HasHasPassword returns a boolean if a field has been set.
func (o *GETCustomersCustomerId200ResponseDataAttributes) HasHasPassword() bool {
	if o != nil && IsNil(o.HasPassword) {
		return true
	}

	return false
}

// SetHasPassword gets a reference to the given interface{} and assigns it to the HasPassword field.
func (o *GETCustomersCustomerId200ResponseDataAttributes) SetHasPassword(v interface{}) {
	o.HasPassword = v
}

// GetTotalOrdersCount returns the TotalOrdersCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCustomersCustomerId200ResponseDataAttributes) GetTotalOrdersCount() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.TotalOrdersCount
}

// GetTotalOrdersCountOk returns a tuple with the TotalOrdersCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCustomersCustomerId200ResponseDataAttributes) GetTotalOrdersCountOk() (*interface{}, bool) {
	if o == nil || IsNil(o.TotalOrdersCount) {
		return nil, false
	}
	return &o.TotalOrdersCount, true
}

// HasTotalOrdersCount returns a boolean if a field has been set.
func (o *GETCustomersCustomerId200ResponseDataAttributes) HasTotalOrdersCount() bool {
	if o != nil && IsNil(o.TotalOrdersCount) {
		return true
	}

	return false
}

// SetTotalOrdersCount gets a reference to the given interface{} and assigns it to the TotalOrdersCount field.
func (o *GETCustomersCustomerId200ResponseDataAttributes) SetTotalOrdersCount(v interface{}) {
	o.TotalOrdersCount = v
}

// GetShopperReference returns the ShopperReference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCustomersCustomerId200ResponseDataAttributes) GetShopperReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ShopperReference
}

// GetShopperReferenceOk returns a tuple with the ShopperReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCustomersCustomerId200ResponseDataAttributes) GetShopperReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ShopperReference) {
		return nil, false
	}
	return &o.ShopperReference, true
}

// HasShopperReference returns a boolean if a field has been set.
func (o *GETCustomersCustomerId200ResponseDataAttributes) HasShopperReference() bool {
	if o != nil && IsNil(o.ShopperReference) {
		return true
	}

	return false
}

// SetShopperReference gets a reference to the given interface{} and assigns it to the ShopperReference field.
func (o *GETCustomersCustomerId200ResponseDataAttributes) SetShopperReference(v interface{}) {
	o.ShopperReference = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCustomersCustomerId200ResponseDataAttributes) GetCreatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCustomersCustomerId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return &o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETCustomersCustomerId200ResponseDataAttributes) HasCreatedAt() bool {
	if o != nil && IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given interface{} and assigns it to the CreatedAt field.
func (o *GETCustomersCustomerId200ResponseDataAttributes) SetCreatedAt(v interface{}) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCustomersCustomerId200ResponseDataAttributes) GetUpdatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCustomersCustomerId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETCustomersCustomerId200ResponseDataAttributes) HasUpdatedAt() bool {
	if o != nil && IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given interface{} and assigns it to the UpdatedAt field.
func (o *GETCustomersCustomerId200ResponseDataAttributes) SetUpdatedAt(v interface{}) {
	o.UpdatedAt = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCustomersCustomerId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCustomersCustomerId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETCustomersCustomerId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *GETCustomersCustomerId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCustomersCustomerId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCustomersCustomerId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETCustomersCustomerId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *GETCustomersCustomerId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCustomersCustomerId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCustomersCustomerId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETCustomersCustomerId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *GETCustomersCustomerId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o GETCustomersCustomerId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETCustomersCustomerId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.HasPassword != nil {
		toSerialize["has_password"] = o.HasPassword
	}
	if o.TotalOrdersCount != nil {
		toSerialize["total_orders_count"] = o.TotalOrdersCount
	}
	if o.ShopperReference != nil {
		toSerialize["shopper_reference"] = o.ShopperReference
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

type NullableGETCustomersCustomerId200ResponseDataAttributes struct {
	value *GETCustomersCustomerId200ResponseDataAttributes
	isSet bool
}

func (v NullableGETCustomersCustomerId200ResponseDataAttributes) Get() *GETCustomersCustomerId200ResponseDataAttributes {
	return v.value
}

func (v *NullableGETCustomersCustomerId200ResponseDataAttributes) Set(val *GETCustomersCustomerId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCustomersCustomerId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCustomersCustomerId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCustomersCustomerId200ResponseDataAttributes(val *GETCustomersCustomerId200ResponseDataAttributes) *NullableGETCustomersCustomerId200ResponseDataAttributes {
	return &NullableGETCustomersCustomerId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullableGETCustomersCustomerId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCustomersCustomerId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
