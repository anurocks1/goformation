package networkfirewall

import (
	"github.com/dexty007/goformation/v4/cloudformation/policies"
)

// RuleGroup_Addresses AWS CloudFormation Resource (AWS::NetworkFirewall::RuleGroup.Addresses)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-addresses.html
type RuleGroup_Addresses struct {

	// Addresses AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-addresses.html#cfn-networkfirewall-rulegroup-addresses-addresses
	Addresses []RuleGroup_Address `json:"Addresses,omitempty"`

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
func (r *RuleGroup_Addresses) AWSCloudFormationType() string {
	return "AWS::NetworkFirewall::RuleGroup.Addresses"
}
