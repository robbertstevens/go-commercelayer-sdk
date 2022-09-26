/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// POSTAvalaraAccounts201ResponseDataAttributes struct for POSTAvalaraAccounts201ResponseDataAttributes
type POSTAvalaraAccounts201ResponseDataAttributes struct {
	// The tax calculator's internal name.
	Name string `json:"name"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// The Avalara account username.
	Username string `json:"username"`
	// The Avalara account password.
	Password string `json:"password"`
	// The Avalara company code.
	CompanyCode string `json:"company_code"`
	// Indicates if the transaction will be recorded and visible on the Avalara website.
	CommitInvoice *string `json:"commit_invoice,omitempty"`
	// Indicates if the seller is responsible for paying/remitting the customs duty & import tax to the customs authorities.
	Ddp *string `json:"ddp,omitempty"`
}

// NewPOSTAvalaraAccounts201ResponseDataAttributes instantiates a new POSTAvalaraAccounts201ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTAvalaraAccounts201ResponseDataAttributes(name string, username string, password string, companyCode string) *POSTAvalaraAccounts201ResponseDataAttributes {
	this := POSTAvalaraAccounts201ResponseDataAttributes{}
	this.Name = name
	this.Username = username
	this.Password = password
	this.CompanyCode = companyCode
	return &this
}

// NewPOSTAvalaraAccounts201ResponseDataAttributesWithDefaults instantiates a new POSTAvalaraAccounts201ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTAvalaraAccounts201ResponseDataAttributesWithDefaults() *POSTAvalaraAccounts201ResponseDataAttributes {
	this := POSTAvalaraAccounts201ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value
func (o *POSTAvalaraAccounts201ResponseDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *POSTAvalaraAccounts201ResponseDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *POSTAvalaraAccounts201ResponseDataAttributes) SetName(v string) {
	o.Name = v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *POSTAvalaraAccounts201ResponseDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTAvalaraAccounts201ResponseDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTAvalaraAccounts201ResponseDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *POSTAvalaraAccounts201ResponseDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *POSTAvalaraAccounts201ResponseDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTAvalaraAccounts201ResponseDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTAvalaraAccounts201ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *POSTAvalaraAccounts201ResponseDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *POSTAvalaraAccounts201ResponseDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTAvalaraAccounts201ResponseDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTAvalaraAccounts201ResponseDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *POSTAvalaraAccounts201ResponseDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetUsername returns the Username field value
func (o *POSTAvalaraAccounts201ResponseDataAttributes) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *POSTAvalaraAccounts201ResponseDataAttributes) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *POSTAvalaraAccounts201ResponseDataAttributes) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value
func (o *POSTAvalaraAccounts201ResponseDataAttributes) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *POSTAvalaraAccounts201ResponseDataAttributes) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *POSTAvalaraAccounts201ResponseDataAttributes) SetPassword(v string) {
	o.Password = v
}

// GetCompanyCode returns the CompanyCode field value
func (o *POSTAvalaraAccounts201ResponseDataAttributes) GetCompanyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CompanyCode
}

// GetCompanyCodeOk returns a tuple with the CompanyCode field value
// and a boolean to check if the value has been set.
func (o *POSTAvalaraAccounts201ResponseDataAttributes) GetCompanyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompanyCode, true
}

// SetCompanyCode sets field value
func (o *POSTAvalaraAccounts201ResponseDataAttributes) SetCompanyCode(v string) {
	o.CompanyCode = v
}

// GetCommitInvoice returns the CommitInvoice field value if set, zero value otherwise.
func (o *POSTAvalaraAccounts201ResponseDataAttributes) GetCommitInvoice() string {
	if o == nil || o.CommitInvoice == nil {
		var ret string
		return ret
	}
	return *o.CommitInvoice
}

// GetCommitInvoiceOk returns a tuple with the CommitInvoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTAvalaraAccounts201ResponseDataAttributes) GetCommitInvoiceOk() (*string, bool) {
	if o == nil || o.CommitInvoice == nil {
		return nil, false
	}
	return o.CommitInvoice, true
}

// HasCommitInvoice returns a boolean if a field has been set.
func (o *POSTAvalaraAccounts201ResponseDataAttributes) HasCommitInvoice() bool {
	if o != nil && o.CommitInvoice != nil {
		return true
	}

	return false
}

// SetCommitInvoice gets a reference to the given string and assigns it to the CommitInvoice field.
func (o *POSTAvalaraAccounts201ResponseDataAttributes) SetCommitInvoice(v string) {
	o.CommitInvoice = &v
}

// GetDdp returns the Ddp field value if set, zero value otherwise.
func (o *POSTAvalaraAccounts201ResponseDataAttributes) GetDdp() string {
	if o == nil || o.Ddp == nil {
		var ret string
		return ret
	}
	return *o.Ddp
}

// GetDdpOk returns a tuple with the Ddp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTAvalaraAccounts201ResponseDataAttributes) GetDdpOk() (*string, bool) {
	if o == nil || o.Ddp == nil {
		return nil, false
	}
	return o.Ddp, true
}

// HasDdp returns a boolean if a field has been set.
func (o *POSTAvalaraAccounts201ResponseDataAttributes) HasDdp() bool {
	if o != nil && o.Ddp != nil {
		return true
	}

	return false
}

// SetDdp gets a reference to the given string and assigns it to the Ddp field.
func (o *POSTAvalaraAccounts201ResponseDataAttributes) SetDdp(v string) {
	o.Ddp = &v
}

func (o POSTAvalaraAccounts201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
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
	if true {
		toSerialize["username"] = o.Username
	}
	if true {
		toSerialize["password"] = o.Password
	}
	if true {
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

type NullablePOSTAvalaraAccounts201ResponseDataAttributes struct {
	value *POSTAvalaraAccounts201ResponseDataAttributes
	isSet bool
}

func (v NullablePOSTAvalaraAccounts201ResponseDataAttributes) Get() *POSTAvalaraAccounts201ResponseDataAttributes {
	return v.value
}

func (v *NullablePOSTAvalaraAccounts201ResponseDataAttributes) Set(val *POSTAvalaraAccounts201ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTAvalaraAccounts201ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTAvalaraAccounts201ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTAvalaraAccounts201ResponseDataAttributes(val *POSTAvalaraAccounts201ResponseDataAttributes) *NullablePOSTAvalaraAccounts201ResponseDataAttributes {
	return &NullablePOSTAvalaraAccounts201ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTAvalaraAccounts201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTAvalaraAccounts201ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
