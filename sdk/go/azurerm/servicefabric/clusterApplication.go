// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicefabric

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The application resource.
type ClusterApplication struct {
	pulumi.CustomResourceState

	// Azure resource etag.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Describes the managed identities for an Azure resource.
	Identity ManagedIdentityResponsePtrOutput `pulumi:"identity"`
	// It will be deprecated in New API, resource location depends on the parent resource.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Azure resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The application resource properties.
	Properties ApplicationResourcePropertiesResponseOutput `pulumi:"properties"`
	// Azure resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Azure resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewClusterApplication registers a new resource with the given unique name, arguments, and options.
func NewClusterApplication(ctx *pulumi.Context,
	name string, args *ClusterApplicationArgs, opts ...pulumi.ResourceOption) (*ClusterApplication, error) {
	if args == nil || args.ClusterName == nil {
		return nil, errors.New("missing required argument 'ClusterName'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ClusterApplicationArgs{}
	}
	var resource ClusterApplication
	err := ctx.RegisterResource("azurerm:servicefabric:ClusterApplication", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterApplication gets an existing ClusterApplication resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterApplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterApplicationState, opts ...pulumi.ResourceOption) (*ClusterApplication, error) {
	var resource ClusterApplication
	err := ctx.ReadResource("azurerm:servicefabric:ClusterApplication", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterApplication resources.
type clusterApplicationState struct {
	// Azure resource etag.
	Etag *string `pulumi:"etag"`
	// Describes the managed identities for an Azure resource.
	Identity *ManagedIdentityResponse `pulumi:"identity"`
	// It will be deprecated in New API, resource location depends on the parent resource.
	Location *string `pulumi:"location"`
	// Azure resource name.
	Name *string `pulumi:"name"`
	// The application resource properties.
	Properties *ApplicationResourcePropertiesResponse `pulumi:"properties"`
	// Azure resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Azure resource type.
	Type *string `pulumi:"type"`
}

type ClusterApplicationState struct {
	// Azure resource etag.
	Etag pulumi.StringPtrInput
	// Describes the managed identities for an Azure resource.
	Identity ManagedIdentityResponsePtrInput
	// It will be deprecated in New API, resource location depends on the parent resource.
	Location pulumi.StringPtrInput
	// Azure resource name.
	Name pulumi.StringPtrInput
	// The application resource properties.
	Properties ApplicationResourcePropertiesResponsePtrInput
	// Azure resource tags.
	Tags pulumi.StringMapInput
	// Azure resource type.
	Type pulumi.StringPtrInput
}

func (ClusterApplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterApplicationState)(nil)).Elem()
}

type clusterApplicationArgs struct {
	// The name of the cluster resource.
	ClusterName string `pulumi:"clusterName"`
	// Describes the managed identities for an Azure resource.
	Identity *ManagedIdentity `pulumi:"identity"`
	// It will be deprecated in New API, resource location depends on the parent resource.
	Location *string `pulumi:"location"`
	// The name of the application resource.
	Name string `pulumi:"name"`
	// The application resource properties.
	Properties *ApplicationResourceProperties `pulumi:"properties"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Azure resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ClusterApplication resource.
type ClusterApplicationArgs struct {
	// The name of the cluster resource.
	ClusterName pulumi.StringInput
	// Describes the managed identities for an Azure resource.
	Identity ManagedIdentityPtrInput
	// It will be deprecated in New API, resource location depends on the parent resource.
	Location pulumi.StringPtrInput
	// The name of the application resource.
	Name pulumi.StringInput
	// The application resource properties.
	Properties ApplicationResourcePropertiesPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Azure resource tags.
	Tags pulumi.StringMapInput
}

func (ClusterApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterApplicationArgs)(nil)).Elem()
}