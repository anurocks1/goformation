package autoscaling

import (
	"github.com/dexty007/goformation/v4/cloudformation/policies"
)

// LaunchConfiguration_MetadataOption AWS CloudFormation Resource (AWS::AutoScaling::LaunchConfiguration.MetadataOption)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-launchconfig-metadataoption.html
type LaunchConfiguration_MetadataOption struct {

	// HttpEndpoint AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-launchconfig-metadataoption.html#cfn-autoscaling-launchconfig-metadataoption-httpendpoint
	HttpEndpoint string `json:"HttpEndpoint,omitempty"`

	// HttpPutResponseHopLimit AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-launchconfig-metadataoption.html#cfn-autoscaling-launchconfig-metadataoption-httpputresponsehoplimit
	HttpPutResponseHopLimit int `json:"HttpPutResponseHopLimit,omitempty"`

	// HttpTokens AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-launchconfig-metadataoption.html#cfn-autoscaling-launchconfig-metadataoption-httptokens
	HttpTokens string `json:"HttpTokens,omitempty"`

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
func (r *LaunchConfiguration_MetadataOption) AWSCloudFormationType() string {
	return "AWS::AutoScaling::LaunchConfiguration.MetadataOption"
}
