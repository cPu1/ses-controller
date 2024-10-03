// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CustomVerificationEmailTemplateSpec defines the desired state of CustomVerificationEmailTemplate.
//
// Contains information about a custom verification email template.
type CustomVerificationEmailTemplateSpec struct {

	// The URL that the recipient of the verification email is sent to if his or
	// her address is not successfully verified.
	// +kubebuilder:validation:Required
	FailureRedirectionURL *string `json:"failureRedirectionURL"`
	// The email address that the custom verification email is sent from.
	// +kubebuilder:validation:Required
	FromEmailAddress *string `json:"fromEmailAddress"`
	// The URL that the recipient of the verification email is sent to if his or
	// her address is successfully verified.
	// +kubebuilder:validation:Required
	SuccessRedirectionURL *string `json:"successRedirectionURL"`
	// The content of the custom verification email. The total size of the email
	// must be less than 10 MB. The message body may contain HTML, with some limitations.
	// For more information, see Custom Verification Email Frequently Asked Questions
	// (https://docs.aws.amazon.com/ses/latest/dg/creating-identities.html#send-email-verify-address-custom)
	// in the Amazon SES Developer Guide.
	// +kubebuilder:validation:Required
	TemplateContent *string `json:"templateContent"`
	// The name of the custom verification email template.
	// +kubebuilder:validation:Required
	TemplateName *string `json:"templateName"`
	// The subject line of the custom verification email.
	// +kubebuilder:validation:Required
	TemplateSubject *string `json:"templateSubject"`
}

// CustomVerificationEmailTemplateStatus defines the observed state of CustomVerificationEmailTemplate
type CustomVerificationEmailTemplateStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
}

// CustomVerificationEmailTemplate is the Schema for the CustomVerificationEmailTemplates API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type CustomVerificationEmailTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CustomVerificationEmailTemplateSpec   `json:"spec,omitempty"`
	Status            CustomVerificationEmailTemplateStatus `json:"status,omitempty"`
}

// CustomVerificationEmailTemplateList contains a list of CustomVerificationEmailTemplate
// +kubebuilder:object:root=true
type CustomVerificationEmailTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CustomVerificationEmailTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CustomVerificationEmailTemplate{}, &CustomVerificationEmailTemplateList{})
}
