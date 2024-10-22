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

// checks if the OrderCopyDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderCopyDataRelationships{}

// OrderCopyDataRelationships struct for OrderCopyDataRelationships
type OrderCopyDataRelationships struct {
	SourceOrder       *AdyenPaymentDataRelationshipsOrder          `json:"source_order,omitempty"`
	TargetOrder       *AdyenPaymentDataRelationshipsOrder          `json:"target_order,omitempty"`
	Events            *AddressDataRelationshipsEvents              `json:"events,omitempty"`
	OrderSubscription *CustomerDataRelationshipsOrderSubscriptions `json:"order_subscription,omitempty"`
}

// NewOrderCopyDataRelationships instantiates a new OrderCopyDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderCopyDataRelationships() *OrderCopyDataRelationships {
	this := OrderCopyDataRelationships{}
	return &this
}

// NewOrderCopyDataRelationshipsWithDefaults instantiates a new OrderCopyDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderCopyDataRelationshipsWithDefaults() *OrderCopyDataRelationships {
	this := OrderCopyDataRelationships{}
	return &this
}

// GetSourceOrder returns the SourceOrder field value if set, zero value otherwise.
func (o *OrderCopyDataRelationships) GetSourceOrder() AdyenPaymentDataRelationshipsOrder {
	if o == nil || IsNil(o.SourceOrder) {
		var ret AdyenPaymentDataRelationshipsOrder
		return ret
	}
	return *o.SourceOrder
}

// GetSourceOrderOk returns a tuple with the SourceOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCopyDataRelationships) GetSourceOrderOk() (*AdyenPaymentDataRelationshipsOrder, bool) {
	if o == nil || IsNil(o.SourceOrder) {
		return nil, false
	}
	return o.SourceOrder, true
}

// HasSourceOrder returns a boolean if a field has been set.
func (o *OrderCopyDataRelationships) HasSourceOrder() bool {
	if o != nil && !IsNil(o.SourceOrder) {
		return true
	}

	return false
}

// SetSourceOrder gets a reference to the given AdyenPaymentDataRelationshipsOrder and assigns it to the SourceOrder field.
func (o *OrderCopyDataRelationships) SetSourceOrder(v AdyenPaymentDataRelationshipsOrder) {
	o.SourceOrder = &v
}

// GetTargetOrder returns the TargetOrder field value if set, zero value otherwise.
func (o *OrderCopyDataRelationships) GetTargetOrder() AdyenPaymentDataRelationshipsOrder {
	if o == nil || IsNil(o.TargetOrder) {
		var ret AdyenPaymentDataRelationshipsOrder
		return ret
	}
	return *o.TargetOrder
}

// GetTargetOrderOk returns a tuple with the TargetOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCopyDataRelationships) GetTargetOrderOk() (*AdyenPaymentDataRelationshipsOrder, bool) {
	if o == nil || IsNil(o.TargetOrder) {
		return nil, false
	}
	return o.TargetOrder, true
}

// HasTargetOrder returns a boolean if a field has been set.
func (o *OrderCopyDataRelationships) HasTargetOrder() bool {
	if o != nil && !IsNil(o.TargetOrder) {
		return true
	}

	return false
}

// SetTargetOrder gets a reference to the given AdyenPaymentDataRelationshipsOrder and assigns it to the TargetOrder field.
func (o *OrderCopyDataRelationships) SetTargetOrder(v AdyenPaymentDataRelationshipsOrder) {
	o.TargetOrder = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *OrderCopyDataRelationships) GetEvents() AddressDataRelationshipsEvents {
	if o == nil || IsNil(o.Events) {
		var ret AddressDataRelationshipsEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCopyDataRelationships) GetEventsOk() (*AddressDataRelationshipsEvents, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *OrderCopyDataRelationships) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given AddressDataRelationshipsEvents and assigns it to the Events field.
func (o *OrderCopyDataRelationships) SetEvents(v AddressDataRelationshipsEvents) {
	o.Events = &v
}

// GetOrderSubscription returns the OrderSubscription field value if set, zero value otherwise.
func (o *OrderCopyDataRelationships) GetOrderSubscription() CustomerDataRelationshipsOrderSubscriptions {
	if o == nil || IsNil(o.OrderSubscription) {
		var ret CustomerDataRelationshipsOrderSubscriptions
		return ret
	}
	return *o.OrderSubscription
}

// GetOrderSubscriptionOk returns a tuple with the OrderSubscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCopyDataRelationships) GetOrderSubscriptionOk() (*CustomerDataRelationshipsOrderSubscriptions, bool) {
	if o == nil || IsNil(o.OrderSubscription) {
		return nil, false
	}
	return o.OrderSubscription, true
}

// HasOrderSubscription returns a boolean if a field has been set.
func (o *OrderCopyDataRelationships) HasOrderSubscription() bool {
	if o != nil && !IsNil(o.OrderSubscription) {
		return true
	}

	return false
}

// SetOrderSubscription gets a reference to the given CustomerDataRelationshipsOrderSubscriptions and assigns it to the OrderSubscription field.
func (o *OrderCopyDataRelationships) SetOrderSubscription(v CustomerDataRelationshipsOrderSubscriptions) {
	o.OrderSubscription = &v
}

func (o OrderCopyDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderCopyDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SourceOrder) {
		toSerialize["source_order"] = o.SourceOrder
	}
	if !IsNil(o.TargetOrder) {
		toSerialize["target_order"] = o.TargetOrder
	}
	if !IsNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	if !IsNil(o.OrderSubscription) {
		toSerialize["order_subscription"] = o.OrderSubscription
	}
	return toSerialize, nil
}

type NullableOrderCopyDataRelationships struct {
	value *OrderCopyDataRelationships
	isSet bool
}

func (v NullableOrderCopyDataRelationships) Get() *OrderCopyDataRelationships {
	return v.value
}

func (v *NullableOrderCopyDataRelationships) Set(val *OrderCopyDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCopyDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCopyDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCopyDataRelationships(val *OrderCopyDataRelationships) *NullableOrderCopyDataRelationships {
	return &NullableOrderCopyDataRelationships{value: val, isSet: true}
}

func (v NullableOrderCopyDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCopyDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
