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

// GETRefunds200ResponseDataInnerAttributes struct for GETRefunds200ResponseDataInnerAttributes
type GETRefunds200ResponseDataInnerAttributes struct {
	// The transaction number, auto generated
	Number *string `json:"number,omitempty"`
	// The international 3-letter currency code as defined by the ISO 4217 standard, inherited from the associated order.
	CurrencyCode *string `json:"currency_code,omitempty"`
	// The transaction amount, in cents.
	AmountCents *int32 `json:"amount_cents,omitempty"`
	// The transaction amount, float.
	AmountFloat *float32 `json:"amount_float,omitempty"`
	// The transaction amount, formatted.
	FormattedAmount *string `json:"formatted_amount,omitempty"`
	// Indicates if the transaction is successful
	Succeeded *bool `json:"succeeded,omitempty"`
	// The message returned by the payment gateway
	Message *string `json:"message,omitempty"`
	// The error code, if any, returned by the payment gateway
	ErrorCode *string `json:"error_code,omitempty"`
	// The error detail, if any, returned by the payment gateway
	ErrorDetail *string `json:"error_detail,omitempty"`
	// The token identifying the transaction, returned by the payment gateway
	Token *string `json:"token,omitempty"`
	// The ID identifying the transaction, returned by the payment gateway
	GatewayTransactionId *string `json:"gateway_transaction_id,omitempty"`
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

// NewGETRefunds200ResponseDataInnerAttributes instantiates a new GETRefunds200ResponseDataInnerAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETRefunds200ResponseDataInnerAttributes() *GETRefunds200ResponseDataInnerAttributes {
	this := GETRefunds200ResponseDataInnerAttributes{}
	return &this
}

// NewGETRefunds200ResponseDataInnerAttributesWithDefaults instantiates a new GETRefunds200ResponseDataInnerAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETRefunds200ResponseDataInnerAttributesWithDefaults() *GETRefunds200ResponseDataInnerAttributes {
	this := GETRefunds200ResponseDataInnerAttributes{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *GETRefunds200ResponseDataInnerAttributes) GetNumber() string {
	if o == nil || o.Number == nil {
		var ret string
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETRefunds200ResponseDataInnerAttributes) GetNumberOk() (*string, bool) {
	if o == nil || o.Number == nil {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *GETRefunds200ResponseDataInnerAttributes) HasNumber() bool {
	if o != nil && o.Number != nil {
		return true
	}

	return false
}

// SetNumber gets a reference to the given string and assigns it to the Number field.
func (o *GETRefunds200ResponseDataInnerAttributes) SetNumber(v string) {
	o.Number = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *GETRefunds200ResponseDataInnerAttributes) GetCurrencyCode() string {
	if o == nil || o.CurrencyCode == nil {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETRefunds200ResponseDataInnerAttributes) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || o.CurrencyCode == nil {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *GETRefunds200ResponseDataInnerAttributes) HasCurrencyCode() bool {
	if o != nil && o.CurrencyCode != nil {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *GETRefunds200ResponseDataInnerAttributes) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetAmountCents returns the AmountCents field value if set, zero value otherwise.
func (o *GETRefunds200ResponseDataInnerAttributes) GetAmountCents() int32 {
	if o == nil || o.AmountCents == nil {
		var ret int32
		return ret
	}
	return *o.AmountCents
}

// GetAmountCentsOk returns a tuple with the AmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETRefunds200ResponseDataInnerAttributes) GetAmountCentsOk() (*int32, bool) {
	if o == nil || o.AmountCents == nil {
		return nil, false
	}
	return o.AmountCents, true
}

// HasAmountCents returns a boolean if a field has been set.
func (o *GETRefunds200ResponseDataInnerAttributes) HasAmountCents() bool {
	if o != nil && o.AmountCents != nil {
		return true
	}

	return false
}

// SetAmountCents gets a reference to the given int32 and assigns it to the AmountCents field.
func (o *GETRefunds200ResponseDataInnerAttributes) SetAmountCents(v int32) {
	o.AmountCents = &v
}

// GetAmountFloat returns the AmountFloat field value if set, zero value otherwise.
func (o *GETRefunds200ResponseDataInnerAttributes) GetAmountFloat() float32 {
	if o == nil || o.AmountFloat == nil {
		var ret float32
		return ret
	}
	return *o.AmountFloat
}

// GetAmountFloatOk returns a tuple with the AmountFloat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETRefunds200ResponseDataInnerAttributes) GetAmountFloatOk() (*float32, bool) {
	if o == nil || o.AmountFloat == nil {
		return nil, false
	}
	return o.AmountFloat, true
}

// HasAmountFloat returns a boolean if a field has been set.
func (o *GETRefunds200ResponseDataInnerAttributes) HasAmountFloat() bool {
	if o != nil && o.AmountFloat != nil {
		return true
	}

	return false
}

// SetAmountFloat gets a reference to the given float32 and assigns it to the AmountFloat field.
func (o *GETRefunds200ResponseDataInnerAttributes) SetAmountFloat(v float32) {
	o.AmountFloat = &v
}

// GetFormattedAmount returns the FormattedAmount field value if set, zero value otherwise.
func (o *GETRefunds200ResponseDataInnerAttributes) GetFormattedAmount() string {
	if o == nil || o.FormattedAmount == nil {
		var ret string
		return ret
	}
	return *o.FormattedAmount
}

// GetFormattedAmountOk returns a tuple with the FormattedAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETRefunds200ResponseDataInnerAttributes) GetFormattedAmountOk() (*string, bool) {
	if o == nil || o.FormattedAmount == nil {
		return nil, false
	}
	return o.FormattedAmount, true
}

// HasFormattedAmount returns a boolean if a field has been set.
func (o *GETRefunds200ResponseDataInnerAttributes) HasFormattedAmount() bool {
	if o != nil && o.FormattedAmount != nil {
		return true
	}

	return false
}

// SetFormattedAmount gets a reference to the given string and assigns it to the FormattedAmount field.
func (o *GETRefunds200ResponseDataInnerAttributes) SetFormattedAmount(v string) {
	o.FormattedAmount = &v
}

// GetSucceeded returns the Succeeded field value if set, zero value otherwise.
func (o *GETRefunds200ResponseDataInnerAttributes) GetSucceeded() bool {
	if o == nil || o.Succeeded == nil {
		var ret bool
		return ret
	}
	return *o.Succeeded
}

// GetSucceededOk returns a tuple with the Succeeded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETRefunds200ResponseDataInnerAttributes) GetSucceededOk() (*bool, bool) {
	if o == nil || o.Succeeded == nil {
		return nil, false
	}
	return o.Succeeded, true
}

// HasSucceeded returns a boolean if a field has been set.
func (o *GETRefunds200ResponseDataInnerAttributes) HasSucceeded() bool {
	if o != nil && o.Succeeded != nil {
		return true
	}

	return false
}

// SetSucceeded gets a reference to the given bool and assigns it to the Succeeded field.
func (o *GETRefunds200ResponseDataInnerAttributes) SetSucceeded(v bool) {
	o.Succeeded = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GETRefunds200ResponseDataInnerAttributes) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETRefunds200ResponseDataInnerAttributes) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GETRefunds200ResponseDataInnerAttributes) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *GETRefunds200ResponseDataInnerAttributes) SetMessage(v string) {
	o.Message = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *GETRefunds200ResponseDataInnerAttributes) GetErrorCode() string {
	if o == nil || o.ErrorCode == nil {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETRefunds200ResponseDataInnerAttributes) GetErrorCodeOk() (*string, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *GETRefunds200ResponseDataInnerAttributes) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *GETRefunds200ResponseDataInnerAttributes) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetErrorDetail returns the ErrorDetail field value if set, zero value otherwise.
func (o *GETRefunds200ResponseDataInnerAttributes) GetErrorDetail() string {
	if o == nil || o.ErrorDetail == nil {
		var ret string
		return ret
	}
	return *o.ErrorDetail
}

// GetErrorDetailOk returns a tuple with the ErrorDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETRefunds200ResponseDataInnerAttributes) GetErrorDetailOk() (*string, bool) {
	if o == nil || o.ErrorDetail == nil {
		return nil, false
	}
	return o.ErrorDetail, true
}

// HasErrorDetail returns a boolean if a field has been set.
func (o *GETRefunds200ResponseDataInnerAttributes) HasErrorDetail() bool {
	if o != nil && o.ErrorDetail != nil {
		return true
	}

	return false
}

// SetErrorDetail gets a reference to the given string and assigns it to the ErrorDetail field.
func (o *GETRefunds200ResponseDataInnerAttributes) SetErrorDetail(v string) {
	o.ErrorDetail = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GETRefunds200ResponseDataInnerAttributes) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETRefunds200ResponseDataInnerAttributes) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GETRefunds200ResponseDataInnerAttributes) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GETRefunds200ResponseDataInnerAttributes) SetToken(v string) {
	o.Token = &v
}

// GetGatewayTransactionId returns the GatewayTransactionId field value if set, zero value otherwise.
func (o *GETRefunds200ResponseDataInnerAttributes) GetGatewayTransactionId() string {
	if o == nil || o.GatewayTransactionId == nil {
		var ret string
		return ret
	}
	return *o.GatewayTransactionId
}

// GetGatewayTransactionIdOk returns a tuple with the GatewayTransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETRefunds200ResponseDataInnerAttributes) GetGatewayTransactionIdOk() (*string, bool) {
	if o == nil || o.GatewayTransactionId == nil {
		return nil, false
	}
	return o.GatewayTransactionId, true
}

// HasGatewayTransactionId returns a boolean if a field has been set.
func (o *GETRefunds200ResponseDataInnerAttributes) HasGatewayTransactionId() bool {
	if o != nil && o.GatewayTransactionId != nil {
		return true
	}

	return false
}

// SetGatewayTransactionId gets a reference to the given string and assigns it to the GatewayTransactionId field.
func (o *GETRefunds200ResponseDataInnerAttributes) SetGatewayTransactionId(v string) {
	o.GatewayTransactionId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETRefunds200ResponseDataInnerAttributes) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETRefunds200ResponseDataInnerAttributes) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETRefunds200ResponseDataInnerAttributes) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETRefunds200ResponseDataInnerAttributes) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GETRefunds200ResponseDataInnerAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETRefunds200ResponseDataInnerAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETRefunds200ResponseDataInnerAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *GETRefunds200ResponseDataInnerAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GETRefunds200ResponseDataInnerAttributes) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETRefunds200ResponseDataInnerAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETRefunds200ResponseDataInnerAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *GETRefunds200ResponseDataInnerAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *GETRefunds200ResponseDataInnerAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETRefunds200ResponseDataInnerAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETRefunds200ResponseDataInnerAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *GETRefunds200ResponseDataInnerAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *GETRefunds200ResponseDataInnerAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETRefunds200ResponseDataInnerAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETRefunds200ResponseDataInnerAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *GETRefunds200ResponseDataInnerAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GETRefunds200ResponseDataInnerAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETRefunds200ResponseDataInnerAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETRefunds200ResponseDataInnerAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *GETRefunds200ResponseDataInnerAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o GETRefunds200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Number != nil {
		toSerialize["number"] = o.Number
	}
	if o.CurrencyCode != nil {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	if o.AmountCents != nil {
		toSerialize["amount_cents"] = o.AmountCents
	}
	if o.AmountFloat != nil {
		toSerialize["amount_float"] = o.AmountFloat
	}
	if o.FormattedAmount != nil {
		toSerialize["formatted_amount"] = o.FormattedAmount
	}
	if o.Succeeded != nil {
		toSerialize["succeeded"] = o.Succeeded
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.ErrorCode != nil {
		toSerialize["error_code"] = o.ErrorCode
	}
	if o.ErrorDetail != nil {
		toSerialize["error_detail"] = o.ErrorDetail
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.GatewayTransactionId != nil {
		toSerialize["gateway_transaction_id"] = o.GatewayTransactionId
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

type NullableGETRefunds200ResponseDataInnerAttributes struct {
	value *GETRefunds200ResponseDataInnerAttributes
	isSet bool
}

func (v NullableGETRefunds200ResponseDataInnerAttributes) Get() *GETRefunds200ResponseDataInnerAttributes {
	return v.value
}

func (v *NullableGETRefunds200ResponseDataInnerAttributes) Set(val *GETRefunds200ResponseDataInnerAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETRefunds200ResponseDataInnerAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETRefunds200ResponseDataInnerAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETRefunds200ResponseDataInnerAttributes(val *GETRefunds200ResponseDataInnerAttributes) *NullableGETRefunds200ResponseDataInnerAttributes {
	return &NullableGETRefunds200ResponseDataInnerAttributes{value: val, isSet: true}
}

func (v NullableGETRefunds200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETRefunds200ResponseDataInnerAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
