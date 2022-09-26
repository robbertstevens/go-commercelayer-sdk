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

// RefundDataRelationships struct for RefundDataRelationships
type RefundDataRelationships struct {
	Order            *AdyenPaymentDataRelationshipsOrder     `json:"order,omitempty"`
	ReferenceCapture *AuthorizationDataRelationshipsCaptures `json:"reference_capture,omitempty"`
	Events           *CustomerAddressDataRelationshipsEvents `json:"events,omitempty"`
}

// NewRefundDataRelationships instantiates a new RefundDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefundDataRelationships() *RefundDataRelationships {
	this := RefundDataRelationships{}
	return &this
}

// NewRefundDataRelationshipsWithDefaults instantiates a new RefundDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefundDataRelationshipsWithDefaults() *RefundDataRelationships {
	this := RefundDataRelationships{}
	return &this
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *RefundDataRelationships) GetOrder() AdyenPaymentDataRelationshipsOrder {
	if o == nil || o.Order == nil {
		var ret AdyenPaymentDataRelationshipsOrder
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefundDataRelationships) GetOrderOk() (*AdyenPaymentDataRelationshipsOrder, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *RefundDataRelationships) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given AdyenPaymentDataRelationshipsOrder and assigns it to the Order field.
func (o *RefundDataRelationships) SetOrder(v AdyenPaymentDataRelationshipsOrder) {
	o.Order = &v
}

// GetReferenceCapture returns the ReferenceCapture field value if set, zero value otherwise.
func (o *RefundDataRelationships) GetReferenceCapture() AuthorizationDataRelationshipsCaptures {
	if o == nil || o.ReferenceCapture == nil {
		var ret AuthorizationDataRelationshipsCaptures
		return ret
	}
	return *o.ReferenceCapture
}

// GetReferenceCaptureOk returns a tuple with the ReferenceCapture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefundDataRelationships) GetReferenceCaptureOk() (*AuthorizationDataRelationshipsCaptures, bool) {
	if o == nil || o.ReferenceCapture == nil {
		return nil, false
	}
	return o.ReferenceCapture, true
}

// HasReferenceCapture returns a boolean if a field has been set.
func (o *RefundDataRelationships) HasReferenceCapture() bool {
	if o != nil && o.ReferenceCapture != nil {
		return true
	}

	return false
}

// SetReferenceCapture gets a reference to the given AuthorizationDataRelationshipsCaptures and assigns it to the ReferenceCapture field.
func (o *RefundDataRelationships) SetReferenceCapture(v AuthorizationDataRelationshipsCaptures) {
	o.ReferenceCapture = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *RefundDataRelationships) GetEvents() CustomerAddressDataRelationshipsEvents {
	if o == nil || o.Events == nil {
		var ret CustomerAddressDataRelationshipsEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefundDataRelationships) GetEventsOk() (*CustomerAddressDataRelationshipsEvents, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *RefundDataRelationships) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given CustomerAddressDataRelationshipsEvents and assigns it to the Events field.
func (o *RefundDataRelationships) SetEvents(v CustomerAddressDataRelationshipsEvents) {
	o.Events = &v
}

func (o RefundDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Order != nil {
		toSerialize["order"] = o.Order
	}
	if o.ReferenceCapture != nil {
		toSerialize["reference_capture"] = o.ReferenceCapture
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableRefundDataRelationships struct {
	value *RefundDataRelationships
	isSet bool
}

func (v NullableRefundDataRelationships) Get() *RefundDataRelationships {
	return v.value
}

func (v *NullableRefundDataRelationships) Set(val *RefundDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableRefundDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableRefundDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefundDataRelationships(val *RefundDataRelationships) *NullableRefundDataRelationships {
	return &NullableRefundDataRelationships{value: val, isSet: true}
}

func (v NullableRefundDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefundDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
