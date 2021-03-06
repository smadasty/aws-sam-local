package cloudformation

// AWSEMRCluster_InstanceFleetProvisioningSpecifications AWS CloudFormation Resource (AWS::EMR::Cluster.InstanceFleetProvisioningSpecifications)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-instancefleetprovisioningspecifications.html
type AWSEMRCluster_InstanceFleetProvisioningSpecifications struct {

	// SpotSpecification AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-instancefleetprovisioningspecifications.html#cfn-elasticmapreduce-cluster-instancefleetprovisioningspecifications-spotspecification
	SpotSpecification *AWSEMRCluster_SpotProvisioningSpecification `json:"SpotSpecification,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRCluster_InstanceFleetProvisioningSpecifications) AWSCloudFormationType() string {
	return "AWS::EMR::Cluster.InstanceFleetProvisioningSpecifications"
}
