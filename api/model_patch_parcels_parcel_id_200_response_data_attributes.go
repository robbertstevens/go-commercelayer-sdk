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

// PATCHParcelsParcelId200ResponseDataAttributes struct for PATCHParcelsParcelId200ResponseDataAttributes
type PATCHParcelsParcelId200ResponseDataAttributes struct {
	// The parcel weight, used to automatically calculate the tax rates from the available carrier accounts.
	Weight *float32 `json:"weight,omitempty"`
	// Can be one of 'gr', 'lb', or 'oz'
	UnitOfWeight *string `json:"unit_of_weight,omitempty"`
	// When shipping outside the US, you need to provide either an Exemption and Exclusion Legend (EEL) code or a Proof of Filing Citation (PFC). Which you need is based on the value of the goods being shipped. Value can be one of \"EEL\" o \"PFC\".
	EelPfc *string `json:"eel_pfc,omitempty"`
	// The type of item you are sending. Can be one of 'merchandise', 'gift', 'documents', 'returned_goods', 'sample', or 'other'.
	ContentsType *string `json:"contents_type,omitempty"`
	// If you specify 'other' in the 'contents_type' attribute, you must supply a brief description in this attribute.
	ContentsExplanation *string `json:"contents_explanation,omitempty"`
	// Indicates if the provided information is accurate
	CustomsCertify *bool `json:"customs_certify,omitempty"`
	// This is the name of the person who is certifying that the information provided on the customs form is accurate. Use a name of the person in your organization who is responsible for this.
	CustomsSigner *string `json:"customs_signer,omitempty"`
	// In case the shipment cannot be delivered, this option tells the carrier what you want to happen to the parcel. You can pass either 'return', or 'abandon'. The value defaults to 'return'. If you pass 'abandon', you will not receive the parcel back if it cannot be delivered.
	NonDeliveryOption *string `json:"non_delivery_option,omitempty"`
	// Describes if your parcel requires any special treatment or quarantine when entering the country. Can be one of 'none', 'other', 'quarantine', or 'sanitary_phytosanitary_inspection'.
	RestrictionType *string `json:"restriction_type,omitempty"`
	// If you specify 'other' in the restriction type, you must supply a brief description of what is required.
	RestrictionComments *string `json:"restriction_comments,omitempty"`
	// Indicates if the parcel requires customs info to get the shipping rates.
	CustomsInfoRequired *bool `json:"customs_info_required,omitempty"`
	// The shipping label url, ready to be downloaded and printed.
	ShippingLabelUrl *string `json:"shipping_label_url,omitempty"`
	// The shipping label file type. One of 'application/pdf', 'application/zpl', 'application/epl2', or 'image/png'.
	ShippingLabelFileType *string `json:"shipping_label_file_type,omitempty"`
	// The shipping label size.
	ShippingLabelSize *string `json:"shipping_label_size,omitempty"`
	// The shipping label resolution.
	ShippingLabelResolution *string `json:"shipping_label_resolution,omitempty"`
	// The tracking number associated to this parcel.
	TrackingNumber *string `json:"tracking_number,omitempty"`
	// The tracking status for this parcel, automatically updated in real time by the shipping carrier.
	TrackingStatus *string `json:"tracking_status,omitempty"`
	// Additional information about the tracking status, automatically updated in real time by the shipping carrier.
	TrackingStatusDetail *string `json:"tracking_status_detail,omitempty"`
	// Time at which the parcel's tracking status was last updated.
	TrackingStatusUpdatedAt *string `json:"tracking_status_updated_at,omitempty"`
	// The parcel's full tracking history, automatically updated in real time by the shipping carrier.
	TrackingDetails *string `json:"tracking_details,omitempty"`
	// The weight of the parcel as measured by the carrier in ounces (if available)
	CarrierWeightOz *string `json:"carrier_weight_oz,omitempty"`
	// The name of the person who signed for the parcel (if available)
	SignedBy *string `json:"signed_by,omitempty"`
	// The type of Incoterm (if available).
	Incoterm *string `json:"incoterm,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewPATCHParcelsParcelId200ResponseDataAttributes instantiates a new PATCHParcelsParcelId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHParcelsParcelId200ResponseDataAttributes() *PATCHParcelsParcelId200ResponseDataAttributes {
	this := PATCHParcelsParcelId200ResponseDataAttributes{}
	return &this
}

// NewPATCHParcelsParcelId200ResponseDataAttributesWithDefaults instantiates a new PATCHParcelsParcelId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHParcelsParcelId200ResponseDataAttributesWithDefaults() *PATCHParcelsParcelId200ResponseDataAttributes {
	this := PATCHParcelsParcelId200ResponseDataAttributes{}
	return &this
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) GetWeight() float32 {
	if o == nil || o.Weight == nil {
		var ret float32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) GetWeightOk() (*float32, bool) {
	if o == nil || o.Weight == nil {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) HasWeight() bool {
	if o != nil && o.Weight != nil {
		return true
	}

	return false
}

// SetWeight gets a reference to the given float32 and assigns it to the Weight field.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) SetWeight(v float32) {
	o.Weight = &v
}

// GetUnitOfWeight returns the UnitOfWeight field value if set, zero value otherwise.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) GetUnitOfWeight() string {
	if o == nil || o.UnitOfWeight == nil {
		var ret string
		return ret
	}
	return *o.UnitOfWeight
}

// GetUnitOfWeightOk returns a tuple with the UnitOfWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) GetUnitOfWeightOk() (*string, bool) {
	if o == nil || o.UnitOfWeight == nil {
		return nil, false
	}
	return o.UnitOfWeight, true
}

// HasUnitOfWeight returns a boolean if a field has been set.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) HasUnitOfWeight() bool {
	if o != nil && o.UnitOfWeight != nil {
		return true
	}

	return false
}

// SetUnitOfWeight gets a reference to the given string and assigns it to the UnitOfWeight field.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) SetUnitOfWeight(v string) {
	o.UnitOfWeight = &v
}

// GetEelPfc returns the EelPfc field value if set, zero value otherwise.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) GetEelPfc() string {
	if o == nil || o.EelPfc == nil {
		var ret string
		return ret
	}
	return *o.EelPfc
}

// GetEelPfcOk returns a tuple with the EelPfc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) GetEelPfcOk() (*string, bool) {
	if o == nil || o.EelPfc == nil {
		return nil, false
	}
	return o.EelPfc, true
}

// HasEelPfc returns a boolean if a field has been set.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) HasEelPfc() bool {
	if o != nil && o.EelPfc != nil {
		return true
	}

	return false
}

// SetEelPfc gets a reference to the given string and assigns it to the EelPfc field.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) SetEelPfc(v string) {
	o.EelPfc = &v
}

// GetContentsType returns the ContentsType field value if set, zero value otherwise.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) GetContentsType() string {
	if o == nil || o.ContentsType == nil {
		var ret string
		return ret
	}
	return *o.ContentsType
}

// GetContentsTypeOk returns a tuple with the ContentsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) GetContentsTypeOk() (*string, bool) {
	if o == nil || o.ContentsType == nil {
		return nil, false
	}
	return o.ContentsType, true
}

// HasContentsType returns a boolean if a field has been set.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) HasContentsType() bool {
	if o != nil && o.ContentsType != nil {
		return true
	}

	return false
}

// SetContentsType gets a reference to the given string and assigns it to the ContentsType field.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) SetContentsType(v string) {
	o.ContentsType = &v
}

// GetContentsExplanation returns the ContentsExplanation field value if set, zero value otherwise.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) GetContentsExplanation() string {
	if o == nil || o.ContentsExplanation == nil {
		var ret string
		return ret
	}
	return *o.ContentsExplanation
}

// GetContentsExplanationOk returns a tuple with the ContentsExplanation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) GetContentsExplanationOk() (*string, bool) {
	if o == nil || o.ContentsExplanation == nil {
		return nil, false
	}
	return o.ContentsExplanation, true
}

// HasContentsExplanation returns a boolean if a field has been set.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) HasContentsExplanation() bool {
	if o != nil && o.ContentsExplanation != nil {
		return true
	}

	return false
}

// SetContentsExplanation gets a reference to the given string and assigns it to the ContentsExplanation field.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) SetContentsExplanation(v string) {
	o.ContentsExplanation = &v
}

// GetCustomsCertify returns the CustomsCertify field value if set, zero value otherwise.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) GetCustomsCertify() bool {
	if o == nil || o.CustomsCertify == nil {
		var ret bool
		return ret
	}
	return *o.CustomsCertify
}

// GetCustomsCertifyOk returns a tuple with the CustomsCertify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) GetCustomsCertifyOk() (*bool, bool) {
	if o == nil || o.CustomsCertify == nil {
		return nil, false
	}
	return o.CustomsCertify, true
}

// HasCustomsCertify returns a boolean if a field has been set.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) HasCustomsCertify() bool {
	if o != nil && o.CustomsCertify != nil {
		return true
	}

	return false
}

// SetCustomsCertify gets a reference to the given bool and assigns it to the CustomsCertify field.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) SetCustomsCertify(v bool) {
	o.CustomsCertify = &v
}

// GetCustomsSigner returns the CustomsSigner field value if set, zero value otherwise.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) GetCustomsSigner() string {
	if o == nil || o.CustomsSigner == nil {
		var ret string
		return ret
	}
	return *o.CustomsSigner
}

// GetCustomsSignerOk returns a tuple with the CustomsSigner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) GetCustomsSignerOk() (*string, bool) {
	if o == nil || o.CustomsSigner == nil {
		return nil, false
	}
	return o.CustomsSigner, true
}

// HasCustomsSigner returns a boolean if a field has been set.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) HasCustomsSigner() bool {
	if o != nil && o.CustomsSigner != nil {
		return true
	}

	return false
}

// SetCustomsSigner gets a reference to the given string and assigns it to the CustomsSigner field.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) SetCustomsSigner(v string) {
	o.CustomsSigner = &v
}

// GetNonDeliveryOption returns the NonDeliveryOption field value if set, zero value otherwise.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) GetNonDeliveryOption() string {
	if o == nil || o.NonDeliveryOption == nil {
		var ret string
		return ret
	}
	return *o.NonDeliveryOption
}

// GetNonDeliveryOptionOk returns a tuple with the NonDeliveryOption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) GetNonDeliveryOptionOk() (*string, bool) {
	if o == nil || o.NonDeliveryOption == nil {
		return nil, false
	}
	return o.NonDeliveryOption, true
}

// HasNonDeliveryOption returns a boolean if a field has been set.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) HasNonDeliveryOption() bool {
	if o != nil && o.NonDeliveryOption != nil {
		return true
	}

	return false
}

// SetNonDeliveryOption gets a reference to the given string and assigns it to the NonDeliveryOption field.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) SetNonDeliveryOption(v string) {
	o.NonDeliveryOption = &v
}

// GetRestrictionType returns the RestrictionType field value if set, zero value otherwise.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) GetRestrictionType() string {
	if o == nil || o.RestrictionType == nil {
		var ret string
		return ret
	}
	return *o.RestrictionType
}

// GetRestrictionTypeOk returns a tuple with the RestrictionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) GetRestrictionTypeOk() (*string, bool) {
	if o == nil || o.RestrictionType == nil {
		return nil, false
	}
	return o.RestrictionType, true
}

// HasRestrictionType returns a boolean if a field has been set.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) HasRestrictionType() bool {
	if o != nil && o.RestrictionType != nil {
		return true
	}

	return false
}

// SetRestrictionType gets a reference to the given string and assigns it to the RestrictionType field.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) SetRestrictionType(v string) {
	o.RestrictionType = &v
}

// GetRestrictionComments returns the RestrictionComments field value if set, zero value otherwise.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) GetRestrictionComments() string {
	if o == nil || o.RestrictionComments == nil {
		var ret string
		return ret
	}
	return *o.RestrictionComments
}

// GetRestrictionCommentsOk returns a tuple with the RestrictionComments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) GetRestrictionCommentsOk() (*string, bool) {
	if o == nil || o.RestrictionComments == nil {
		return nil, false
	}
	return o.RestrictionComments, true
}

// HasRestrictionComments returns a boolean if a field has been set.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) HasRestrictionComments() bool {
	if o != nil && o.RestrictionComments != nil {
		return true
	}

	return false
}

// SetRestrictionComments gets a reference to the given string and assigns it to the RestrictionComments field.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) SetRestrictionComments(v string) {
	o.RestrictionComments = &v
}

// GetCustomsInfoRequired returns the CustomsInfoRequired field value if set, zero value otherwise.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) GetCustomsInfoRequired() bool {
	if o == nil || o.CustomsInfoRequired == nil {
		var ret bool
		return ret
	}
	return *o.CustomsInfoRequired
}

// GetCustomsInfoRequiredOk returns a tuple with the CustomsInfoRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) GetCustomsInfoRequiredOk() (*bool, bool) {
	if o == nil || o.CustomsInfoRequired == nil {
		return nil, false
	}
	return o.CustomsInfoRequired, true
}

// HasCustomsInfoRequired returns a boolean if a field has been set.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) HasCustomsInfoRequired() bool {
	if o != nil && o.CustomsInfoRequired != nil {
		return true
	}

	return false
}

// SetCustomsInfoRequired gets a reference to the given bool and assigns it to the CustomsInfoRequired field.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) SetCustomsInfoRequired(v bool) {
	o.CustomsInfoRequired = &v
}

// GetShippingLabelUrl returns the ShippingLabelUrl field value if set, zero value otherwise.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) GetShippingLabelUrl() string {
	if o == nil || o.ShippingLabelUrl == nil {
		var ret string
		return ret
	}
	return *o.ShippingLabelUrl
}

// GetShippingLabelUrlOk returns a tuple with the ShippingLabelUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) GetShippingLabelUrlOk() (*string, bool) {
	if o == nil || o.ShippingLabelUrl == nil {
		return nil, false
	}
	return o.ShippingLabelUrl, true
}

// HasShippingLabelUrl returns a boolean if a field has been set.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) HasShippingLabelUrl() bool {
	if o != nil && o.ShippingLabelUrl != nil {
		return true
	}

	return false
}

// SetShippingLabelUrl gets a reference to the given string and assigns it to the ShippingLabelUrl field.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) SetShippingLabelUrl(v string) {
	o.ShippingLabelUrl = &v
}

// GetShippingLabelFileType returns the ShippingLabelFileType field value if set, zero value otherwise.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) GetShippingLabelFileType() string {
	if o == nil || o.ShippingLabelFileType == nil {
		var ret string
		return ret
	}
	return *o.ShippingLabelFileType
}

// GetShippingLabelFileTypeOk returns a tuple with the ShippingLabelFileType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) GetShippingLabelFileTypeOk() (*string, bool) {
	if o == nil || o.ShippingLabelFileType == nil {
		return nil, false
	}
	return o.ShippingLabelFileType, true
}

// HasShippingLabelFileType returns a boolean if a field has been set.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) HasShippingLabelFileType() bool {
	if o != nil && o.ShippingLabelFileType != nil {
		return true
	}

	return false
}

// SetShippingLabelFileType gets a reference to the given string and assigns it to the ShippingLabelFileType field.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) SetShippingLabelFileType(v string) {
	o.ShippingLabelFileType = &v
}

// GetShippingLabelSize returns the ShippingLabelSize field value if set, zero value otherwise.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) GetShippingLabelSize() string {
	if o == nil || o.ShippingLabelSize == nil {
		var ret string
		return ret
	}
	return *o.ShippingLabelSize
}

// GetShippingLabelSizeOk returns a tuple with the ShippingLabelSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) GetShippingLabelSizeOk() (*string, bool) {
	if o == nil || o.ShippingLabelSize == nil {
		return nil, false
	}
	return o.ShippingLabelSize, true
}

// HasShippingLabelSize returns a boolean if a field has been set.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) HasShippingLabelSize() bool {
	if o != nil && o.ShippingLabelSize != nil {
		return true
	}

	return false
}

// SetShippingLabelSize gets a reference to the given string and assigns it to the ShippingLabelSize field.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) SetShippingLabelSize(v string) {
	o.ShippingLabelSize = &v
}

// GetShippingLabelResolution returns the ShippingLabelResolution field value if set, zero value otherwise.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) GetShippingLabelResolution() string {
	if o == nil || o.ShippingLabelResolution == nil {
		var ret string
		return ret
	}
	return *o.ShippingLabelResolution
}

// GetShippingLabelResolutionOk returns a tuple with the ShippingLabelResolution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) GetShippingLabelResolutionOk() (*string, bool) {
	if o == nil || o.ShippingLabelResolution == nil {
		return nil, false
	}
	return o.ShippingLabelResolution, true
}

// HasShippingLabelResolution returns a boolean if a field has been set.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) HasShippingLabelResolution() bool {
	if o != nil && o.ShippingLabelResolution != nil {
		return true
	}

	return false
}

// SetShippingLabelResolution gets a reference to the given string and assigns it to the ShippingLabelResolution field.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) SetShippingLabelResolution(v string) {
	o.ShippingLabelResolution = &v
}

// GetTrackingNumber returns the TrackingNumber field value if set, zero value otherwise.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) GetTrackingNumber() string {
	if o == nil || o.TrackingNumber == nil {
		var ret string
		return ret
	}
	return *o.TrackingNumber
}

// GetTrackingNumberOk returns a tuple with the TrackingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) GetTrackingNumberOk() (*string, bool) {
	if o == nil || o.TrackingNumber == nil {
		return nil, false
	}
	return o.TrackingNumber, true
}

// HasTrackingNumber returns a boolean if a field has been set.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) HasTrackingNumber() bool {
	if o != nil && o.TrackingNumber != nil {
		return true
	}

	return false
}

// SetTrackingNumber gets a reference to the given string and assigns it to the TrackingNumber field.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) SetTrackingNumber(v string) {
	o.TrackingNumber = &v
}

// GetTrackingStatus returns the TrackingStatus field value if set, zero value otherwise.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) GetTrackingStatus() string {
	if o == nil || o.TrackingStatus == nil {
		var ret string
		return ret
	}
	return *o.TrackingStatus
}

// GetTrackingStatusOk returns a tuple with the TrackingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) GetTrackingStatusOk() (*string, bool) {
	if o == nil || o.TrackingStatus == nil {
		return nil, false
	}
	return o.TrackingStatus, true
}

// HasTrackingStatus returns a boolean if a field has been set.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) HasTrackingStatus() bool {
	if o != nil && o.TrackingStatus != nil {
		return true
	}

	return false
}

// SetTrackingStatus gets a reference to the given string and assigns it to the TrackingStatus field.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) SetTrackingStatus(v string) {
	o.TrackingStatus = &v
}

// GetTrackingStatusDetail returns the TrackingStatusDetail field value if set, zero value otherwise.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) GetTrackingStatusDetail() string {
	if o == nil || o.TrackingStatusDetail == nil {
		var ret string
		return ret
	}
	return *o.TrackingStatusDetail
}

// GetTrackingStatusDetailOk returns a tuple with the TrackingStatusDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) GetTrackingStatusDetailOk() (*string, bool) {
	if o == nil || o.TrackingStatusDetail == nil {
		return nil, false
	}
	return o.TrackingStatusDetail, true
}

// HasTrackingStatusDetail returns a boolean if a field has been set.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) HasTrackingStatusDetail() bool {
	if o != nil && o.TrackingStatusDetail != nil {
		return true
	}

	return false
}

// SetTrackingStatusDetail gets a reference to the given string and assigns it to the TrackingStatusDetail field.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) SetTrackingStatusDetail(v string) {
	o.TrackingStatusDetail = &v
}

// GetTrackingStatusUpdatedAt returns the TrackingStatusUpdatedAt field value if set, zero value otherwise.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) GetTrackingStatusUpdatedAt() string {
	if o == nil || o.TrackingStatusUpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.TrackingStatusUpdatedAt
}

// GetTrackingStatusUpdatedAtOk returns a tuple with the TrackingStatusUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) GetTrackingStatusUpdatedAtOk() (*string, bool) {
	if o == nil || o.TrackingStatusUpdatedAt == nil {
		return nil, false
	}
	return o.TrackingStatusUpdatedAt, true
}

// HasTrackingStatusUpdatedAt returns a boolean if a field has been set.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) HasTrackingStatusUpdatedAt() bool {
	if o != nil && o.TrackingStatusUpdatedAt != nil {
		return true
	}

	return false
}

// SetTrackingStatusUpdatedAt gets a reference to the given string and assigns it to the TrackingStatusUpdatedAt field.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) SetTrackingStatusUpdatedAt(v string) {
	o.TrackingStatusUpdatedAt = &v
}

// GetTrackingDetails returns the TrackingDetails field value if set, zero value otherwise.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) GetTrackingDetails() string {
	if o == nil || o.TrackingDetails == nil {
		var ret string
		return ret
	}
	return *o.TrackingDetails
}

// GetTrackingDetailsOk returns a tuple with the TrackingDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) GetTrackingDetailsOk() (*string, bool) {
	if o == nil || o.TrackingDetails == nil {
		return nil, false
	}
	return o.TrackingDetails, true
}

// HasTrackingDetails returns a boolean if a field has been set.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) HasTrackingDetails() bool {
	if o != nil && o.TrackingDetails != nil {
		return true
	}

	return false
}

// SetTrackingDetails gets a reference to the given string and assigns it to the TrackingDetails field.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) SetTrackingDetails(v string) {
	o.TrackingDetails = &v
}

// GetCarrierWeightOz returns the CarrierWeightOz field value if set, zero value otherwise.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) GetCarrierWeightOz() string {
	if o == nil || o.CarrierWeightOz == nil {
		var ret string
		return ret
	}
	return *o.CarrierWeightOz
}

// GetCarrierWeightOzOk returns a tuple with the CarrierWeightOz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) GetCarrierWeightOzOk() (*string, bool) {
	if o == nil || o.CarrierWeightOz == nil {
		return nil, false
	}
	return o.CarrierWeightOz, true
}

// HasCarrierWeightOz returns a boolean if a field has been set.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) HasCarrierWeightOz() bool {
	if o != nil && o.CarrierWeightOz != nil {
		return true
	}

	return false
}

// SetCarrierWeightOz gets a reference to the given string and assigns it to the CarrierWeightOz field.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) SetCarrierWeightOz(v string) {
	o.CarrierWeightOz = &v
}

// GetSignedBy returns the SignedBy field value if set, zero value otherwise.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) GetSignedBy() string {
	if o == nil || o.SignedBy == nil {
		var ret string
		return ret
	}
	return *o.SignedBy
}

// GetSignedByOk returns a tuple with the SignedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) GetSignedByOk() (*string, bool) {
	if o == nil || o.SignedBy == nil {
		return nil, false
	}
	return o.SignedBy, true
}

// HasSignedBy returns a boolean if a field has been set.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) HasSignedBy() bool {
	if o != nil && o.SignedBy != nil {
		return true
	}

	return false
}

// SetSignedBy gets a reference to the given string and assigns it to the SignedBy field.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) SetSignedBy(v string) {
	o.SignedBy = &v
}

// GetIncoterm returns the Incoterm field value if set, zero value otherwise.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) GetIncoterm() string {
	if o == nil || o.Incoterm == nil {
		var ret string
		return ret
	}
	return *o.Incoterm
}

// GetIncotermOk returns a tuple with the Incoterm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) GetIncotermOk() (*string, bool) {
	if o == nil || o.Incoterm == nil {
		return nil, false
	}
	return o.Incoterm, true
}

// HasIncoterm returns a boolean if a field has been set.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) HasIncoterm() bool {
	if o != nil && o.Incoterm != nil {
		return true
	}

	return false
}

// SetIncoterm gets a reference to the given string and assigns it to the Incoterm field.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) SetIncoterm(v string) {
	o.Incoterm = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PATCHParcelsParcelId200ResponseDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o PATCHParcelsParcelId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Weight != nil {
		toSerialize["weight"] = o.Weight
	}
	if o.UnitOfWeight != nil {
		toSerialize["unit_of_weight"] = o.UnitOfWeight
	}
	if o.EelPfc != nil {
		toSerialize["eel_pfc"] = o.EelPfc
	}
	if o.ContentsType != nil {
		toSerialize["contents_type"] = o.ContentsType
	}
	if o.ContentsExplanation != nil {
		toSerialize["contents_explanation"] = o.ContentsExplanation
	}
	if o.CustomsCertify != nil {
		toSerialize["customs_certify"] = o.CustomsCertify
	}
	if o.CustomsSigner != nil {
		toSerialize["customs_signer"] = o.CustomsSigner
	}
	if o.NonDeliveryOption != nil {
		toSerialize["non_delivery_option"] = o.NonDeliveryOption
	}
	if o.RestrictionType != nil {
		toSerialize["restriction_type"] = o.RestrictionType
	}
	if o.RestrictionComments != nil {
		toSerialize["restriction_comments"] = o.RestrictionComments
	}
	if o.CustomsInfoRequired != nil {
		toSerialize["customs_info_required"] = o.CustomsInfoRequired
	}
	if o.ShippingLabelUrl != nil {
		toSerialize["shipping_label_url"] = o.ShippingLabelUrl
	}
	if o.ShippingLabelFileType != nil {
		toSerialize["shipping_label_file_type"] = o.ShippingLabelFileType
	}
	if o.ShippingLabelSize != nil {
		toSerialize["shipping_label_size"] = o.ShippingLabelSize
	}
	if o.ShippingLabelResolution != nil {
		toSerialize["shipping_label_resolution"] = o.ShippingLabelResolution
	}
	if o.TrackingNumber != nil {
		toSerialize["tracking_number"] = o.TrackingNumber
	}
	if o.TrackingStatus != nil {
		toSerialize["tracking_status"] = o.TrackingStatus
	}
	if o.TrackingStatusDetail != nil {
		toSerialize["tracking_status_detail"] = o.TrackingStatusDetail
	}
	if o.TrackingStatusUpdatedAt != nil {
		toSerialize["tracking_status_updated_at"] = o.TrackingStatusUpdatedAt
	}
	if o.TrackingDetails != nil {
		toSerialize["tracking_details"] = o.TrackingDetails
	}
	if o.CarrierWeightOz != nil {
		toSerialize["carrier_weight_oz"] = o.CarrierWeightOz
	}
	if o.SignedBy != nil {
		toSerialize["signed_by"] = o.SignedBy
	}
	if o.Incoterm != nil {
		toSerialize["incoterm"] = o.Incoterm
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

type NullablePATCHParcelsParcelId200ResponseDataAttributes struct {
	value *PATCHParcelsParcelId200ResponseDataAttributes
	isSet bool
}

func (v NullablePATCHParcelsParcelId200ResponseDataAttributes) Get() *PATCHParcelsParcelId200ResponseDataAttributes {
	return v.value
}

func (v *NullablePATCHParcelsParcelId200ResponseDataAttributes) Set(val *PATCHParcelsParcelId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHParcelsParcelId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHParcelsParcelId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHParcelsParcelId200ResponseDataAttributes(val *PATCHParcelsParcelId200ResponseDataAttributes) *NullablePATCHParcelsParcelId200ResponseDataAttributes {
	return &NullablePATCHParcelsParcelId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHParcelsParcelId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHParcelsParcelId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
