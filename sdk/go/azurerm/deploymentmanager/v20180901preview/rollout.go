// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180901preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Defines the PUT rollout request body.
//
// ## Example Usage
// ### Create or update rollout
//
// ```go
// package main
//
// import (
// 	deploymentmanager "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/deploymentmanager/v20180901preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := deploymentmanager.NewRollout(ctx, "rollout", &deploymentmanager.RolloutArgs{
// 			ArtifactSourceId: pulumi.String("/subscriptions/caac1590-e859-444f-a9e0-62091c0f5929/resourceGroups/myResourceGroup/Microsoft.DeploymentManager/artifactSources/myArtifactSource"),
// 			BuildVersion:     pulumi.String("1.0.0.1"),
// 			Identity: &deploymentmanager.IdentityArgs{
// 				IdentityIds: pulumi.StringArray{
// 					pulumi.String("/subscriptions/caac1590-e859-444f-a9e0-62091c0f5929/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userassignedidentities/myuseridentity"),
// 				},
// 				Type: pulumi.String("userAssigned"),
// 			},
// 			Location:          pulumi.String("centralus"),
// 			ResourceGroupName: pulumi.String("myResourceGroup"),
// 			RolloutName:       pulumi.String("myRollout"),
// 			StepGroups: deploymentmanager.StepArray{
// 				&deploymentmanager.StepArgs{
// 					DeploymentTargetId: pulumi.String("Microsoft.DeploymentManager/serviceTopologies/myTopology/services/myService/serviceUnits/myServiceUnit1'"),
// 					Name:               pulumi.String("FirstRegion"),
// 					PostDeploymentSteps: deploymentmanager.PrePostStepArray{
// 						&deploymentmanager.PrePostStepArgs{
// 							StepId: pulumi.String("Microsoft.DeploymentManager/steps/postDeployStep1"),
// 						},
// 					},
// 					PreDeploymentSteps: deploymentmanager.PrePostStepArray{
// 						&deploymentmanager.PrePostStepArgs{
// 							StepId: pulumi.String("Microsoft.DeploymentManager/steps/preDeployStep1"),
// 						},
// 						&deploymentmanager.PrePostStepArgs{
// 							StepId: pulumi.String("Microsoft.DeploymentManager/steps/preDeployStep2"),
// 						},
// 					},
// 				},
// 				&deploymentmanager.StepArgs{
// 					DependsOnStepGroups: pulumi.StringArray{
// 						pulumi.String("FirstRegion"),
// 					},
// 					DeploymentTargetId: pulumi.String("Microsoft.DeploymentManager/serviceTopologies/myTopology/services/myService/serviceUnits/myServiceUnit2'"),
// 					Name:               pulumi.String("SecondRegion"),
// 					PostDeploymentSteps: deploymentmanager.PrePostStepArray{
// 						&deploymentmanager.PrePostStepArgs{
// 							StepId: pulumi.String("Microsoft.DeploymentManager/steps/postDeployStep5"),
// 						},
// 					},
// 					PreDeploymentSteps: deploymentmanager.PrePostStepArray{
// 						&deploymentmanager.PrePostStepArgs{
// 							StepId: pulumi.String("Microsoft.DeploymentManager/steps/preDeployStep3"),
// 						},
// 						&deploymentmanager.PrePostStepArgs{
// 							StepId: pulumi.String("Microsoft.DeploymentManager/steps/preDeployStep4"),
// 						},
// 					},
// 				},
// 			},
// 			Tags:                    nil,
// 			TargetServiceTopologyId: pulumi.String("/subscriptions/caac1590-e859-444f-a9e0-62091c0f5929/resourceGroups/myResourceGroup/Microsoft.DeploymentManager/serviceTopologies/myTopology"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type Rollout struct {
	pulumi.CustomResourceState

	// The reference to the artifact source resource Id where the payload is located.
	ArtifactSourceId pulumi.StringPtrOutput `pulumi:"artifactSourceId"`
	// The version of the build being deployed.
	BuildVersion pulumi.StringOutput `pulumi:"buildVersion"`
	// Identity for the resource.
	Identity IdentityResponseOutput `pulumi:"identity"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The list of step groups that define the orchestration.
	StepGroups StepResponseArrayOutput `pulumi:"stepGroups"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The resource Id of the service topology from which service units are being referenced in step groups to be deployed.
	TargetServiceTopologyId pulumi.StringOutput `pulumi:"targetServiceTopologyId"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRollout registers a new resource with the given unique name, arguments, and options.
func NewRollout(ctx *pulumi.Context,
	name string, args *RolloutArgs, opts ...pulumi.ResourceOption) (*Rollout, error) {
	if args == nil || args.BuildVersion == nil {
		return nil, errors.New("missing required argument 'BuildVersion'")
	}
	if args == nil || args.Identity == nil {
		return nil, errors.New("missing required argument 'Identity'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.RolloutName == nil {
		return nil, errors.New("missing required argument 'RolloutName'")
	}
	if args == nil || args.StepGroups == nil {
		return nil, errors.New("missing required argument 'StepGroups'")
	}
	if args == nil || args.TargetServiceTopologyId == nil {
		return nil, errors.New("missing required argument 'TargetServiceTopologyId'")
	}
	if args == nil {
		args = &RolloutArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:deploymentmanager/v20191101preview:Rollout"),
		},
	})
	opts = append(opts, aliases)
	var resource Rollout
	err := ctx.RegisterResource("azurerm:deploymentmanager/v20180901preview:Rollout", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRollout gets an existing Rollout resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRollout(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RolloutState, opts ...pulumi.ResourceOption) (*Rollout, error) {
	var resource Rollout
	err := ctx.ReadResource("azurerm:deploymentmanager/v20180901preview:Rollout", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Rollout resources.
type rolloutState struct {
	// The reference to the artifact source resource Id where the payload is located.
	ArtifactSourceId *string `pulumi:"artifactSourceId"`
	// The version of the build being deployed.
	BuildVersion *string `pulumi:"buildVersion"`
	// Identity for the resource.
	Identity *IdentityResponse `pulumi:"identity"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// The list of step groups that define the orchestration.
	StepGroups []StepResponse `pulumi:"stepGroups"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The resource Id of the service topology from which service units are being referenced in step groups to be deployed.
	TargetServiceTopologyId *string `pulumi:"targetServiceTopologyId"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `pulumi:"type"`
}

type RolloutState struct {
	// The reference to the artifact source resource Id where the payload is located.
	ArtifactSourceId pulumi.StringPtrInput
	// The version of the build being deployed.
	BuildVersion pulumi.StringPtrInput
	// Identity for the resource.
	Identity IdentityResponsePtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// The list of step groups that define the orchestration.
	StepGroups StepResponseArrayInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The resource Id of the service topology from which service units are being referenced in step groups to be deployed.
	TargetServiceTopologyId pulumi.StringPtrInput
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringPtrInput
}

func (RolloutState) ElementType() reflect.Type {
	return reflect.TypeOf((*rolloutState)(nil)).Elem()
}

type rolloutArgs struct {
	// The reference to the artifact source resource Id where the payload is located.
	ArtifactSourceId *string `pulumi:"artifactSourceId"`
	// The version of the build being deployed.
	BuildVersion string `pulumi:"buildVersion"`
	// Identity for the resource.
	Identity Identity `pulumi:"identity"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The rollout name.
	RolloutName string `pulumi:"rolloutName"`
	// The list of step groups that define the orchestration.
	StepGroups []StepType `pulumi:"stepGroups"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The resource Id of the service topology from which service units are being referenced in step groups to be deployed.
	TargetServiceTopologyId string `pulumi:"targetServiceTopologyId"`
}

// The set of arguments for constructing a Rollout resource.
type RolloutArgs struct {
	// The reference to the artifact source resource Id where the payload is located.
	ArtifactSourceId pulumi.StringPtrInput
	// The version of the build being deployed.
	BuildVersion pulumi.StringInput
	// Identity for the resource.
	Identity IdentityInput
	// The geo-location where the resource lives
	Location pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The rollout name.
	RolloutName pulumi.StringInput
	// The list of step groups that define the orchestration.
	StepGroups StepTypeArrayInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The resource Id of the service topology from which service units are being referenced in step groups to be deployed.
	TargetServiceTopologyId pulumi.StringInput
}

func (RolloutArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rolloutArgs)(nil)).Elem()
}
