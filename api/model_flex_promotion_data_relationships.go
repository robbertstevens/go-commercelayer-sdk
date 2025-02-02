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

// checks if the FlexPromotionDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FlexPromotionDataRelationships{}

// FlexPromotionDataRelationships struct for FlexPromotionDataRelationships
type FlexPromotionDataRelationships struct {
	CouponCodesPromotionRule *BuyXPayYPromotionDataRelationshipsCouponCodesPromotionRule `json:"coupon_codes_promotion_rule,omitempty"`
	Coupons                  *BuyXPayYPromotionDataRelationshipsCoupons                  `json:"coupons,omitempty"`
	Attachments              *AuthorizationDataRelationshipsAttachments                  `json:"attachments,omitempty"`
	Events                   *AddressDataRelationshipsEvents                             `json:"events,omitempty"`
	Tags                     *AddressDataRelationshipsTags                               `json:"tags,omitempty"`
	Versions                 *AddressDataRelationshipsVersions                           `json:"versions,omitempty"`
}

// NewFlexPromotionDataRelationships instantiates a new FlexPromotionDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlexPromotionDataRelationships() *FlexPromotionDataRelationships {
	this := FlexPromotionDataRelationships{}
	return &this
}

// NewFlexPromotionDataRelationshipsWithDefaults instantiates a new FlexPromotionDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlexPromotionDataRelationshipsWithDefaults() *FlexPromotionDataRelationships {
	this := FlexPromotionDataRelationships{}
	return &this
}

// GetCouponCodesPromotionRule returns the CouponCodesPromotionRule field value if set, zero value otherwise.
func (o *FlexPromotionDataRelationships) GetCouponCodesPromotionRule() BuyXPayYPromotionDataRelationshipsCouponCodesPromotionRule {
	if o == nil || IsNil(o.CouponCodesPromotionRule) {
		var ret BuyXPayYPromotionDataRelationshipsCouponCodesPromotionRule
		return ret
	}
	return *o.CouponCodesPromotionRule
}

// GetCouponCodesPromotionRuleOk returns a tuple with the CouponCodesPromotionRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexPromotionDataRelationships) GetCouponCodesPromotionRuleOk() (*BuyXPayYPromotionDataRelationshipsCouponCodesPromotionRule, bool) {
	if o == nil || IsNil(o.CouponCodesPromotionRule) {
		return nil, false
	}
	return o.CouponCodesPromotionRule, true
}

// HasCouponCodesPromotionRule returns a boolean if a field has been set.
func (o *FlexPromotionDataRelationships) HasCouponCodesPromotionRule() bool {
	if o != nil && !IsNil(o.CouponCodesPromotionRule) {
		return true
	}

	return false
}

// SetCouponCodesPromotionRule gets a reference to the given BuyXPayYPromotionDataRelationshipsCouponCodesPromotionRule and assigns it to the CouponCodesPromotionRule field.
func (o *FlexPromotionDataRelationships) SetCouponCodesPromotionRule(v BuyXPayYPromotionDataRelationshipsCouponCodesPromotionRule) {
	o.CouponCodesPromotionRule = &v
}

// GetCoupons returns the Coupons field value if set, zero value otherwise.
func (o *FlexPromotionDataRelationships) GetCoupons() BuyXPayYPromotionDataRelationshipsCoupons {
	if o == nil || IsNil(o.Coupons) {
		var ret BuyXPayYPromotionDataRelationshipsCoupons
		return ret
	}
	return *o.Coupons
}

// GetCouponsOk returns a tuple with the Coupons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexPromotionDataRelationships) GetCouponsOk() (*BuyXPayYPromotionDataRelationshipsCoupons, bool) {
	if o == nil || IsNil(o.Coupons) {
		return nil, false
	}
	return o.Coupons, true
}

// HasCoupons returns a boolean if a field has been set.
func (o *FlexPromotionDataRelationships) HasCoupons() bool {
	if o != nil && !IsNil(o.Coupons) {
		return true
	}

	return false
}

// SetCoupons gets a reference to the given BuyXPayYPromotionDataRelationshipsCoupons and assigns it to the Coupons field.
func (o *FlexPromotionDataRelationships) SetCoupons(v BuyXPayYPromotionDataRelationshipsCoupons) {
	o.Coupons = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *FlexPromotionDataRelationships) GetAttachments() AuthorizationDataRelationshipsAttachments {
	if o == nil || IsNil(o.Attachments) {
		var ret AuthorizationDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexPromotionDataRelationships) GetAttachmentsOk() (*AuthorizationDataRelationshipsAttachments, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *FlexPromotionDataRelationships) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given AuthorizationDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *FlexPromotionDataRelationships) SetAttachments(v AuthorizationDataRelationshipsAttachments) {
	o.Attachments = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *FlexPromotionDataRelationships) GetEvents() AddressDataRelationshipsEvents {
	if o == nil || IsNil(o.Events) {
		var ret AddressDataRelationshipsEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexPromotionDataRelationships) GetEventsOk() (*AddressDataRelationshipsEvents, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *FlexPromotionDataRelationships) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given AddressDataRelationshipsEvents and assigns it to the Events field.
func (o *FlexPromotionDataRelationships) SetEvents(v AddressDataRelationshipsEvents) {
	o.Events = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *FlexPromotionDataRelationships) GetTags() AddressDataRelationshipsTags {
	if o == nil || IsNil(o.Tags) {
		var ret AddressDataRelationshipsTags
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexPromotionDataRelationships) GetTagsOk() (*AddressDataRelationshipsTags, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *FlexPromotionDataRelationships) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given AddressDataRelationshipsTags and assigns it to the Tags field.
func (o *FlexPromotionDataRelationships) SetTags(v AddressDataRelationshipsTags) {
	o.Tags = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *FlexPromotionDataRelationships) GetVersions() AddressDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret AddressDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexPromotionDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *FlexPromotionDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given AddressDataRelationshipsVersions and assigns it to the Versions field.
func (o *FlexPromotionDataRelationships) SetVersions(v AddressDataRelationshipsVersions) {
	o.Versions = &v
}

func (o FlexPromotionDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlexPromotionDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CouponCodesPromotionRule) {
		toSerialize["coupon_codes_promotion_rule"] = o.CouponCodesPromotionRule
	}
	if !IsNil(o.Coupons) {
		toSerialize["coupons"] = o.Coupons
	}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	if !IsNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	return toSerialize, nil
}

type NullableFlexPromotionDataRelationships struct {
	value *FlexPromotionDataRelationships
	isSet bool
}

func (v NullableFlexPromotionDataRelationships) Get() *FlexPromotionDataRelationships {
	return v.value
}

func (v *NullableFlexPromotionDataRelationships) Set(val *FlexPromotionDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableFlexPromotionDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableFlexPromotionDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlexPromotionDataRelationships(val *FlexPromotionDataRelationships) *NullableFlexPromotionDataRelationships {
	return &NullableFlexPromotionDataRelationships{value: val, isSet: true}
}

func (v NullableFlexPromotionDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlexPromotionDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
