// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191101preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupSourceControlConfiguration(ctx *pulumi.Context, args *LookupSourceControlConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupSourceControlConfigurationResult, error) {
	var rv LookupSourceControlConfigurationResult
	err := ctx.Invoke("azurerm:kubernetesconfiguration/v20191101preview:getSourceControlConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSourceControlConfigurationArgs struct {
	// The name of the kubernetes cluster.
	ClusterName string `pulumi:"clusterName"`
	// The Kubernetes cluster resource name - either managedClusters (for AKS clusters) or connectedClusters (for OnPrem K8S clusters).
	ClusterResourceName string `pulumi:"clusterResourceName"`
	// The Kubernetes cluster RP - either Microsoft.ContainerService (for AKS clusters) or Microsoft.Kubernetes (for OnPrem K8S clusters).
	ClusterRp string `pulumi:"clusterRp"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the Source Control Configuration.
	SourceControlConfigurationName string `pulumi:"sourceControlConfigurationName"`
}

// The SourceControl Configuration object.
type LookupSourceControlConfigurationResult struct {
	// Compliance Status of the Configuration
	ComplianceStatus ComplianceStatusResponse `pulumi:"complianceStatus"`
	// Option to enable Helm Operator for this git configuration.
	EnableHelmOperator *string `pulumi:"enableHelmOperator"`
	// Properties for Helm operator.
	HelmOperatorProperties *HelmOperatorPropertiesResponse `pulumi:"helmOperatorProperties"`
	// Resource name
	Name string `pulumi:"name"`
	// Instance name of the operator - identifying the specific configuration.
	OperatorInstanceName *string `pulumi:"operatorInstanceName"`
	// The namespace to which this operator is installed to. Maximum of 253 lower case alphanumeric characters, hyphen and period only.
	OperatorNamespace *string `pulumi:"operatorNamespace"`
	// Any Parameters for the Operator instance in string format.
	OperatorParams *string `pulumi:"operatorParams"`
	// Scope at which the operator will be installed.
	OperatorScope *string `pulumi:"operatorScope"`
	// Type of the operator
	OperatorType *string `pulumi:"operatorType"`
	// The provisioning state of the resource provider.
	ProvisioningState string `pulumi:"provisioningState"`
	// Public Key associated with this SourceControl configuration (either generated within the cluster or provided by the user).
	RepositoryPublicKey string `pulumi:"repositoryPublicKey"`
	// Url of the SourceControl Repository.
	RepositoryUrl *string `pulumi:"repositoryUrl"`
	// Resource type
	Type string `pulumi:"type"`
}
