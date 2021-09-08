package kendra

import (
	"github.com/wizrocket/goformation/v4/cloudformation/policies"
)

// Index_JwtTokenTypeConfiguration AWS CloudFormation Resource (AWS::Kendra::Index.JwtTokenTypeConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-jwttokentypeconfiguration.html
type Index_JwtTokenTypeConfiguration struct {

	// ClaimRegex AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-jwttokentypeconfiguration.html#cfn-kendra-index-jwttokentypeconfiguration-claimregex
	ClaimRegex string `json:"ClaimRegex,omitempty"`

	// GroupAttributeField AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-jwttokentypeconfiguration.html#cfn-kendra-index-jwttokentypeconfiguration-groupattributefield
	GroupAttributeField string `json:"GroupAttributeField,omitempty"`

	// Issuer AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-jwttokentypeconfiguration.html#cfn-kendra-index-jwttokentypeconfiguration-issuer
	Issuer string `json:"Issuer,omitempty"`

	// KeyLocation AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-jwttokentypeconfiguration.html#cfn-kendra-index-jwttokentypeconfiguration-keylocation
	KeyLocation string `json:"KeyLocation,omitempty"`

	// SecretManagerArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-jwttokentypeconfiguration.html#cfn-kendra-index-jwttokentypeconfiguration-secretmanagerarn
	SecretManagerArn string `json:"SecretManagerArn,omitempty"`

	// URL AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-jwttokentypeconfiguration.html#cfn-kendra-index-jwttokentypeconfiguration-url
	URL string `json:"URL,omitempty"`

	// UserNameAttributeField AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-jwttokentypeconfiguration.html#cfn-kendra-index-jwttokentypeconfiguration-usernameattributefield
	UserNameAttributeField string `json:"UserNameAttributeField,omitempty"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationUpdateReplacePolicy represents a CloudFormation UpdateReplacePolicy
	AWSCloudFormationUpdateReplacePolicy policies.UpdateReplacePolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`

	// AWSCloudFormationCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	AWSCloudFormationCondition string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *Index_JwtTokenTypeConfiguration) AWSCloudFormationType() string {
	return "AWS::Kendra::Index.JwtTokenTypeConfiguration"
}
