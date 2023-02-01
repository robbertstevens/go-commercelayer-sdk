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

// GETAvalaraAccounts200ResponseDataInnerAttributes struct for GETAvalaraAccounts200ResponseDataInnerAttributes
type GETAvalaraAccounts200ResponseDataInnerAttributes struct {
	// The tax calculator's internal name.
	Name *string `json:"name,omitempty"`
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
	// The Avalara account username.
	Username *string `json:"username,omitempty"`
	// The Avalara company code.
	CompanyCode *string `json:"company_code,omitempty"`
	// Indicates if the transaction will be recorded and visible on the Avalara website.
	CommitInvoice *string `json:"commit_invoice,omitempty"`
	// Indicates if the seller is responsible for paying/remitting the customs duty & import tax to the customs authorities.
	Ddp *string `json:"ddp,omitempty"`
}

// NewGETAvalaraAccounts200ResponseDataInnerAttributes instantiates a new GETAvalaraAccounts200ResponseDataInnerAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETAvalaraAccounts200ResponseDataInnerAttributes() *GETAvalaraAccounts200ResponseDataInnerAttributes {
	this := GETAvalaraAccounts200ResponseDataInnerAttributes{}
	return &this
}

// NewGETAvalaraAccounts200ResponseDataInnerAttributesWithDefaults instantiates a new GETAvalaraAccounts200ResponseDataInnerAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETAvalaraAccounts200ResponseDataInnerAttributesWithDefaults() *GETAvalaraAccounts200ResponseDataInnerAttributes {
	this := GETAvalaraAccounts200ResponseDataInnerAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) SetName(v string) {
	o.Name = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) SetUsername(v string) {
	o.Username = &v
}

// GetCompanyCode returns the CompanyCode field value if set, zero value otherwise.
func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) GetCompanyCode() string {
	if o == nil || o.CompanyCode == nil {
		var ret string
		return ret
	}
	return *o.CompanyCode
}

// GetCompanyCodeOk returns a tuple with the CompanyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) GetCompanyCodeOk() (*string, bool) {
	if o == nil || o.CompanyCode == nil {
		return nil, false
	}
	return o.CompanyCode, true
}

// HasCompanyCode returns a boolean if a field has been set.
func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) HasCompanyCode() bool {
	if o != nil && o.CompanyCode != nil {
		return true
	}

	return false
}

// SetCompanyCode gets a reference to the given string and assigns it to the CompanyCode field.
func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) SetCompanyCode(v string) {
	o.CompanyCode = &v
}

// GetCommitInvoice returns the CommitInvoice field value if set, zero value otherwise.
func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) GetCommitInvoice() string {
	if o == nil || o.CommitInvoice == nil {
		var ret string
		return ret
	}
	return *o.CommitInvoice
}

// GetCommitInvoiceOk returns a tuple with the CommitInvoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) GetCommitInvoiceOk() (*string, bool) {
	if o == nil || o.CommitInvoice == nil {
		return nil, false
	}
	return o.CommitInvoice, true
}

// HasCommitInvoice returns a boolean if a field has been set.
func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) HasCommitInvoice() bool {
	if o != nil && o.CommitInvoice != nil {
		return true
	}

	return false
}

// SetCommitInvoice gets a reference to the given string and assigns it to the CommitInvoice field.
func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) SetCommitInvoice(v string) {
	o.CommitInvoice = &v
}

// GetDdp returns the Ddp field value if set, zero value otherwise.
func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) GetDdp() string {
	if o == nil || o.Ddp == nil {
		var ret string
		return ret
	}
	return *o.Ddp
}

// GetDdpOk returns a tuple with the Ddp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) GetDdpOk() (*string, bool) {
	if o == nil || o.Ddp == nil {
		return nil, false
	}
	return o.Ddp, true
}

// HasDdp returns a boolean if a field has been set.
func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) HasDdp() bool {
	if o != nil && o.Ddp != nil {
		return true
	}

	return false
}

// SetDdp gets a reference to the given string and assigns it to the Ddp field.
func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) SetDdp(v string) {
	o.Ddp = &v
}

func (o GETAvalaraAccounts200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
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
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.CompanyCode != nil {
		toSerialize["company_code"] = o.CompanyCode
	}
	if o.CommitInvoice != nil {
		toSerialize["commit_invoice"] = o.CommitInvoice
	}
	if o.Ddp != nil {
		toSerialize["ddp"] = o.Ddp
	}
	return json.Marshal(toSerialize)
}

type NullableGETAvalaraAccounts200ResponseDataInnerAttributes struct {
	value *GETAvalaraAccounts200ResponseDataInnerAttributes
	isSet bool
}

func (v NullableGETAvalaraAccounts200ResponseDataInnerAttributes) Get() *GETAvalaraAccounts200ResponseDataInnerAttributes {
	return v.value
}

func (v *NullableGETAvalaraAccounts200ResponseDataInnerAttributes) Set(val *GETAvalaraAccounts200ResponseDataInnerAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETAvalaraAccounts200ResponseDataInnerAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETAvalaraAccounts200ResponseDataInnerAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETAvalaraAccounts200ResponseDataInnerAttributes(val *GETAvalaraAccounts200ResponseDataInnerAttributes) *NullableGETAvalaraAccounts200ResponseDataInnerAttributes {
	return &NullableGETAvalaraAccounts200ResponseDataInnerAttributes{value: val, isSet: true}
}

func (v NullableGETAvalaraAccounts200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETAvalaraAccounts200ResponseDataInnerAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
