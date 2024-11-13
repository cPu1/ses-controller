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
	"github.com/aws/aws-sdk-go/aws"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = &aws.JSONValue{}
	_ = ackv1alpha1.AWSAccountID("")
)

// When included in a receipt rule, this action adds a header to the received
// email.
//
// For information about adding a header using a receipt rule, see the Amazon
// SES Developer Guide (https://docs.aws.amazon.com/ses/latest/dg/receiving-email-action-add-header.html).
type AddHeaderAction struct {
	HeaderName  *string `json:"headerName,omitempty"`
	HeaderValue *string `json:"headerValue,omitempty"`
}

// When included in a receipt rule, this action rejects the received email by
// returning a bounce response to the sender and, optionally, publishes a notification
// to Amazon Simple Notification Service (Amazon SNS).
//
// For information about sending a bounce message in response to a received
// email, see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/dg/receiving-email-action-bounce.html).
type BounceAction struct {
	Message       *string `json:"message,omitempty"`
	Sender        *string `json:"sender,omitempty"`
	SmtpReplyCode *string `json:"smtpReplyCode,omitempty"`
	StatusCode    *string `json:"statusCode,omitempty"`
	TopicARN      *string `json:"topicARN,omitempty"`
}

// Recipient-related information to include in the Delivery Status Notification
// (DSN) when an email that Amazon SES receives on your behalf bounces.
//
// For information about receiving email through Amazon SES, see the Amazon
// SES Developer Guide (https://docs.aws.amazon.com/ses/latest/dg/receiving-email.html).
type BouncedRecipientInfo struct {
	Recipient    *string `json:"recipient,omitempty"`
	RecipientARN *string `json:"recipientARN,omitempty"`
}

// Contains information associated with an Amazon CloudWatch event destination
// to which email sending events are published.
//
// Event destinations, such as Amazon CloudWatch, are associated with configuration
// sets, which enable you to publish email sending events. For information about
// using configuration sets, see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/dg/monitor-sending-activity.html).
type CloudWatchDestination struct {
	DimensionConfigurations []*CloudWatchDimensionConfiguration `json:"dimensionConfigurations,omitempty"`
}

// Contains the dimension configuration to use when you publish email sending
// events to Amazon CloudWatch.
//
// For information about publishing email sending events to Amazon CloudWatch,
// see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/dg/monitor-sending-activity.html).
type CloudWatchDimensionConfiguration struct {
	DefaultDimensionValue *string `json:"defaultDimensionValue,omitempty"`
	DimensionName         *string `json:"dimensionName,omitempty"`
	DimensionValueSource  *string `json:"dimensionValueSource,omitempty"`
}

// The name of the configuration set.
//
// Configuration sets let you create groups of rules that you can apply to the
// emails you send using Amazon SES. For more information about using configuration
// sets, see Using Amazon SES Configuration Sets (https://docs.aws.amazon.com/ses/latest/dg/using-configuration-sets.html)
// in the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/dg/).
type ConfigurationSet_SDK struct {
	Name *string `json:"name,omitempty"`
}

// Contains information about a custom verification email template.
type CustomVerificationEmailTemplate_SDK struct {
	FailureRedirectionURL *string `json:"failureRedirectionURL,omitempty"`
	FromEmailAddress      *string `json:"fromEmailAddress,omitempty"`
	SuccessRedirectionURL *string `json:"successRedirectionURL,omitempty"`
	TemplateName          *string `json:"templateName,omitempty"`
	TemplateSubject       *string `json:"templateSubject,omitempty"`
}

// Specifies whether messages that use the configuration set are required to
// use Transport Layer Security (TLS).
type DeliveryOptions struct {
	TLSPolicy *string `json:"tlsPolicy,omitempty"`
}

// Contains information about an event destination.
//
// When you create or update an event destination, you must provide one, and
// only one, destination. The destination can be Amazon CloudWatch, Amazon Kinesis
// Firehose or Amazon Simple Notification Service (Amazon SNS).
//
// Event destinations are associated with configuration sets, which enable you
// to publish email sending events to Amazon CloudWatch, Amazon Kinesis Firehose,
// or Amazon Simple Notification Service (Amazon SNS). For information about
// using configuration sets, see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/dg/monitor-sending-activity.html).
type EventDestination struct {
	// Contains information associated with an Amazon CloudWatch event destination
	// to which email sending events are published.
	//
	// Event destinations, such as Amazon CloudWatch, are associated with configuration
	// sets, which enable you to publish email sending events. For information about
	// using configuration sets, see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/dg/monitor-sending-activity.html).
	CloudWatchDestination *CloudWatchDestination `json:"cloudWatchDestination,omitempty"`
	Enabled               *bool                  `json:"enabled,omitempty"`
	// Contains the delivery stream ARN and the IAM role ARN associated with an
	// Amazon Kinesis Firehose event destination.
	//
	// Event destinations, such as Amazon Kinesis Firehose, are associated with
	// configuration sets, which enable you to publish email sending events. For
	// information about using configuration sets, see the Amazon SES Developer
	// Guide (https://docs.aws.amazon.com/ses/latest/dg/monitor-sending-activity.html).
	KinesisFirehoseDestination *KinesisFirehoseDestination `json:"kinesisFirehoseDestination,omitempty"`
	MatchingEventTypes         []*string                   `json:"matchingEventTypes,omitempty"`
	Name                       *string                     `json:"name,omitempty"`
	// Contains the topic ARN associated with an Amazon Simple Notification Service
	// (Amazon SNS) event destination.
	//
	// Event destinations, such as Amazon SNS, are associated with configuration
	// sets, which enable you to publish email sending events. For information about
	// using configuration sets, see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/dg/monitor-sending-activity.html).
	SNSDestination *SNSDestination `json:"snsDestination,omitempty"`
}

// Represents the DKIM attributes of a verified email address or a domain.
type IdentityDkimAttributes struct {
	DkimEnabled *bool `json:"dkimEnabled,omitempty"`
}

// Represents the notification attributes of an identity, including whether
// an identity has Amazon Simple Notification Service (Amazon SNS) topics set
// for bounce, complaint, and/or delivery notifications, and whether feedback
// forwarding is enabled for bounce and complaint notifications.
type IdentityNotificationAttributes struct {
	ForwardingEnabled                      *bool `json:"forwardingEnabled,omitempty"`
	HeadersInBounceNotificationsEnabled    *bool `json:"headersInBounceNotificationsEnabled,omitempty"`
	HeadersInComplaintNotificationsEnabled *bool `json:"headersInComplaintNotificationsEnabled,omitempty"`
	HeadersInDeliveryNotificationsEnabled  *bool `json:"headersInDeliveryNotificationsEnabled,omitempty"`
}

// Contains the delivery stream ARN and the IAM role ARN associated with an
// Amazon Kinesis Firehose event destination.
//
// Event destinations, such as Amazon Kinesis Firehose, are associated with
// configuration sets, which enable you to publish email sending events. For
// information about using configuration sets, see the Amazon SES Developer
// Guide (https://docs.aws.amazon.com/ses/latest/dg/monitor-sending-activity.html).
type KinesisFirehoseDestination struct {
	DeliveryStreamARN *string `json:"deliveryStreamARN,omitempty"`
	IAMRoleARN        *string `json:"iamRoleARN,omitempty"`
}

// When included in a receipt rule, this action calls an Amazon Web Services
// Lambda function and, optionally, publishes a notification to Amazon Simple
// Notification Service (Amazon SNS).
//
// To enable Amazon SES to call your Amazon Web Services Lambda function or
// to publish to an Amazon SNS topic of another account, Amazon SES must have
// permission to access those resources. For information about giving permissions,
// see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/dg/receiving-email-permissions.html).
//
// For information about using Amazon Web Services Lambda actions in receipt
// rules, see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/dg/receiving-email-action-lambda.html).
type LambdaAction struct {
	FunctionARN    *string `json:"functionARN,omitempty"`
	InvocationType *string `json:"invocationType,omitempty"`
	TopicARN       *string `json:"topicARN,omitempty"`
}

// An action that Amazon SES can take when it receives an email on behalf of
// one or more email addresses or domains that you own. An instance of this
// data type can represent only one action.
//
// For information about setting up receipt rules, see the Amazon SES Developer
// Guide (https://docs.aws.amazon.com/ses/latest/dg/receiving-email-receipt-rules-console-walkthrough.html).
type ReceiptAction struct {
	// When included in a receipt rule, this action adds a header to the received
	// email.
	//
	// For information about adding a header using a receipt rule, see the Amazon
	// SES Developer Guide (https://docs.aws.amazon.com/ses/latest/dg/receiving-email-action-add-header.html).
	AddHeaderAction *AddHeaderAction `json:"addHeaderAction,omitempty"`
	// When included in a receipt rule, this action rejects the received email by
	// returning a bounce response to the sender and, optionally, publishes a notification
	// to Amazon Simple Notification Service (Amazon SNS).
	//
	// For information about sending a bounce message in response to a received
	// email, see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/dg/receiving-email-action-bounce.html).
	BounceAction *BounceAction `json:"bounceAction,omitempty"`
	// When included in a receipt rule, this action calls an Amazon Web Services
	// Lambda function and, optionally, publishes a notification to Amazon Simple
	// Notification Service (Amazon SNS).
	//
	// To enable Amazon SES to call your Amazon Web Services Lambda function or
	// to publish to an Amazon SNS topic of another account, Amazon SES must have
	// permission to access those resources. For information about giving permissions,
	// see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/dg/receiving-email-permissions.html).
	//
	// For information about using Amazon Web Services Lambda actions in receipt
	// rules, see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/dg/receiving-email-action-lambda.html).
	LambdaAction *LambdaAction `json:"lambdaAction,omitempty"`
	// When included in a receipt rule, this action saves the received message to
	// an Amazon Simple Storage Service (Amazon S3) bucket and, optionally, publishes
	// a notification to Amazon Simple Notification Service (Amazon SNS).
	//
	// To enable Amazon SES to write emails to your Amazon S3 bucket, use an Amazon
	// Web Services KMS key to encrypt your emails, or publish to an Amazon SNS
	// topic of another account, Amazon SES must have permission to access those
	// resources. For information about granting permissions, see the Amazon SES
	// Developer Guide (https://docs.aws.amazon.com/ses/latest/dg/receiving-email-permissions.html).
	//
	// When you save your emails to an Amazon S3 bucket, the maximum email size
	// (including headers) is 40 MB. Emails larger than that bounces.
	//
	// For information about specifying Amazon S3 actions in receipt rules, see
	// the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/dg/receiving-email-action-s3.html).
	S3Action *S3Action `json:"s3Action,omitempty"`
	// When included in a receipt rule, this action publishes a notification to
	// Amazon Simple Notification Service (Amazon SNS). This action includes a complete
	// copy of the email content in the Amazon SNS notifications. Amazon SNS notifications
	// for all other actions simply provide information about the email. They do
	// not include the email content itself.
	//
	// If you own the Amazon SNS topic, you don't need to do anything to give Amazon
	// SES permission to publish emails to it. However, if you don't own the Amazon
	// SNS topic, you need to attach a policy to the topic to give Amazon SES permissions
	// to access it. For information about giving permissions, see the Amazon SES
	// Developer Guide (https://docs.aws.amazon.com/ses/latest/dg/receiving-email-permissions.html).
	//
	// You can only publish emails that are 150 KB or less (including the header)
	// to Amazon SNS. Larger emails bounce. If you anticipate emails larger than
	// 150 KB, use the S3 action instead.
	//
	// For information about using a receipt rule to publish an Amazon SNS notification,
	// see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/dg/receiving-email-action-sns.html).
	SNSAction *SNSAction `json:"snsAction,omitempty"`
	// When included in a receipt rule, this action terminates the evaluation of
	// the receipt rule set and, optionally, publishes a notification to Amazon
	// Simple Notification Service (Amazon SNS).
	//
	// For information about setting a stop action in a receipt rule, see the Amazon
	// SES Developer Guide (https://docs.aws.amazon.com/ses/latest/dg/receiving-email-action-stop.html).
	StopAction *StopAction `json:"stopAction,omitempty"`
	// When included in a receipt rule, this action calls Amazon WorkMail and, optionally,
	// publishes a notification to Amazon Simple Notification Service (Amazon SNS).
	// It usually isn't necessary to set this up manually, because Amazon WorkMail
	// adds the rule automatically during its setup procedure.
	//
	// For information using a receipt rule to call Amazon WorkMail, see the Amazon
	// SES Developer Guide (https://docs.aws.amazon.com/ses/latest/dg/receiving-email-action-workmail.html).
	WorkmailAction *WorkmailAction `json:"workmailAction,omitempty"`
}

// Receipt rules enable you to specify which actions Amazon SES should take
// when it receives mail on behalf of one or more email addresses or domains
// that you own.
//
// Each receipt rule defines a set of email addresses or domains that it applies
// to. If the email addresses or domains match at least one recipient address
// of the message, Amazon SES executes all of the receipt rule's actions on
// the message.
//
// For information about setting up receipt rules, see the Amazon SES Developer
// Guide (https://docs.aws.amazon.com/ses/latest/dg/receiving-email-receipt-rules-console-walkthrough.html).
type ReceiptRule struct {
	Actions     []*ReceiptAction `json:"actions,omitempty"`
	Enabled     *bool            `json:"enabled,omitempty"`
	Name        *string          `json:"name,omitempty"`
	Recipients  []*string        `json:"recipients,omitempty"`
	ScanEnabled *bool            `json:"scanEnabled,omitempty"`
	TLSPolicy   *string          `json:"tlsPolicy,omitempty"`
}

// Information about a receipt rule set.
//
// A receipt rule set is a collection of rules that specify what Amazon SES
// should do with mail it receives on behalf of your account's verified domains.
//
// For information about setting up receipt rule sets, see the Amazon SES Developer
// Guide (https://docs.aws.amazon.com/ses/latest/dg/receiving-email-concepts.html#receiving-email-concepts-rules).
type ReceiptRuleSetMetadata struct {
	CreatedTimestamp *metav1.Time `json:"createdTimestamp,omitempty"`
	Name             *string      `json:"name,omitempty"`
}

// Recipient-related information to include in the Delivery Status Notification
// (DSN) when an email that Amazon SES receives on your behalf bounces.
//
// For information about receiving email through Amazon SES, see the Amazon
// SES Developer Guide (https://docs.aws.amazon.com/ses/latest/dg/receiving-email.html).
type RecipientDsnFields struct {
	FinalRecipient *string `json:"finalRecipient,omitempty"`
}

// Contains information about the reputation settings for a configuration set.
type ReputationOptions struct {
	LastFreshStart           *metav1.Time `json:"lastFreshStart,omitempty"`
	ReputationMetricsEnabled *bool        `json:"reputationMetricsEnabled,omitempty"`
	SendingEnabled           *bool        `json:"sendingEnabled,omitempty"`
}

// When included in a receipt rule, this action saves the received message to
// an Amazon Simple Storage Service (Amazon S3) bucket and, optionally, publishes
// a notification to Amazon Simple Notification Service (Amazon SNS).
//
// To enable Amazon SES to write emails to your Amazon S3 bucket, use an Amazon
// Web Services KMS key to encrypt your emails, or publish to an Amazon SNS
// topic of another account, Amazon SES must have permission to access those
// resources. For information about granting permissions, see the Amazon SES
// Developer Guide (https://docs.aws.amazon.com/ses/latest/dg/receiving-email-permissions.html).
//
// When you save your emails to an Amazon S3 bucket, the maximum email size
// (including headers) is 40 MB. Emails larger than that bounces.
//
// For information about specifying Amazon S3 actions in receipt rules, see
// the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/dg/receiving-email-action-s3.html).
type S3Action struct {
	BucketName      *string `json:"bucketName,omitempty"`
	KMSKeyARN       *string `json:"kmsKeyARN,omitempty"`
	ObjectKeyPrefix *string `json:"objectKeyPrefix,omitempty"`
	TopicARN        *string `json:"topicARN,omitempty"`
}

// When included in a receipt rule, this action publishes a notification to
// Amazon Simple Notification Service (Amazon SNS). This action includes a complete
// copy of the email content in the Amazon SNS notifications. Amazon SNS notifications
// for all other actions simply provide information about the email. They do
// not include the email content itself.
//
// If you own the Amazon SNS topic, you don't need to do anything to give Amazon
// SES permission to publish emails to it. However, if you don't own the Amazon
// SNS topic, you need to attach a policy to the topic to give Amazon SES permissions
// to access it. For information about giving permissions, see the Amazon SES
// Developer Guide (https://docs.aws.amazon.com/ses/latest/dg/receiving-email-permissions.html).
//
// You can only publish emails that are 150 KB or less (including the header)
// to Amazon SNS. Larger emails bounce. If you anticipate emails larger than
// 150 KB, use the S3 action instead.
//
// For information about using a receipt rule to publish an Amazon SNS notification,
// see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/dg/receiving-email-action-sns.html).
type SNSAction struct {
	Encoding *string `json:"encoding,omitempty"`
	TopicARN *string `json:"topicARN,omitempty"`
}

// Contains the topic ARN associated with an Amazon Simple Notification Service
// (Amazon SNS) event destination.
//
// Event destinations, such as Amazon SNS, are associated with configuration
// sets, which enable you to publish email sending events. For information about
// using configuration sets, see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/dg/monitor-sending-activity.html).
type SNSDestination struct {
	TopicARN *string `json:"topicARN,omitempty"`
	// Reference field for TopicARN
	TopicRef *ackv1alpha1.AWSResourceReferenceWrapper `json:"topicRef,omitempty"`
}

// Represents sending statistics data. Each SendDataPoint contains statistics
// for a 15-minute period of sending activity.
type SendDataPoint struct {
	Timestamp *metav1.Time `json:"timestamp,omitempty"`
}

// When included in a receipt rule, this action terminates the evaluation of
// the receipt rule set and, optionally, publishes a notification to Amazon
// Simple Notification Service (Amazon SNS).
//
// For information about setting a stop action in a receipt rule, see the Amazon
// SES Developer Guide (https://docs.aws.amazon.com/ses/latest/dg/receiving-email-action-stop.html).
type StopAction struct {
	Scope    *string `json:"scope,omitempty"`
	TopicARN *string `json:"topicARN,omitempty"`
}

// Contains information about an email template.
type TemplateMetadata struct {
	CreatedTimestamp *metav1.Time `json:"createdTimestamp,omitempty"`
	Name             *string      `json:"name,omitempty"`
}

// The content of the email, composed of a subject line and either an HTML part
// or a text-only part.
type Template_SDK struct {
	HTMLPart     *string `json:"htmlPart,omitempty"`
	SubjectPart  *string `json:"subjectPart,omitempty"`
	TemplateName *string `json:"templateName,omitempty"`
	TextPart     *string `json:"textPart,omitempty"`
}

// A domain that is used to redirect email recipients to an Amazon SES-operated
// domain. This domain captures open and click events generated by Amazon SES
// emails.
//
// For more information, see Configuring Custom Domains to Handle Open and Click
// Tracking (https://docs.aws.amazon.com/ses/latest/dg/configure-custom-open-click-domains.html)
// in the Amazon SES Developer Guide.
type TrackingOptions struct {
	CustomRedirectDomain *string `json:"customRedirectDomain,omitempty"`
}

// When included in a receipt rule, this action calls Amazon WorkMail and, optionally,
// publishes a notification to Amazon Simple Notification Service (Amazon SNS).
// It usually isn't necessary to set this up manually, because Amazon WorkMail
// adds the rule automatically during its setup procedure.
//
// For information using a receipt rule to call Amazon WorkMail, see the Amazon
// SES Developer Guide (https://docs.aws.amazon.com/ses/latest/dg/receiving-email-action-workmail.html).
type WorkmailAction struct {
	OrganizationARN *string `json:"organizationARN,omitempty"`
	TopicARN        *string `json:"topicARN,omitempty"`
}
