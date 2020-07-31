// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190701

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The remediation definition.
type RemediationAtResource struct {
	pulumi.CustomResourceState

	// The name of the remediation.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties for the remediation.
	Properties RemediationPropertiesResponseOutput `pulumi:"properties"`
	// The type of the remediation.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRemediationAtResource registers a new resource with the given unique name, arguments, and options.
func NewRemediationAtResource(ctx *pulumi.Context,
	name string, args *RemediationAtResourceArgs, opts ...pulumi.ResourceOption) (*RemediationAtResource, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceId == nil {
		return nil, errors.New("missing required argument 'ResourceId'")
	}
	if args == nil {
		args = &RemediationAtResourceArgs{}
	}
	var resource RemediationAtResource
	err := ctx.RegisterResource("azurerm:policyinsights/v20190701:RemediationAtResource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRemediationAtResource gets an existing RemediationAtResource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRemediationAtResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RemediationAtResourceState, opts ...pulumi.ResourceOption) (*RemediationAtResource, error) {
	var resource RemediationAtResource
	err := ctx.ReadResource("azurerm:policyinsights/v20190701:RemediationAtResource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RemediationAtResource resources.
type remediationAtResourceState struct {
	// The name of the remediation.
	Name *string `pulumi:"name"`
	// Properties for the remediation.
	Properties *RemediationPropertiesResponse `pulumi:"properties"`
	// The type of the remediation.
	Type *string `pulumi:"type"`
}

type RemediationAtResourceState struct {
	// The name of the remediation.
	Name pulumi.StringPtrInput
	// Properties for the remediation.
	Properties RemediationPropertiesResponsePtrInput
	// The type of the remediation.
	Type pulumi.StringPtrInput
}

func (RemediationAtResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*remediationAtResourceState)(nil)).Elem()
}

type remediationAtResourceArgs struct {
	// The name of the remediation.
	Name string `pulumi:"name"`
	// Properties for the remediation.
	Properties *RemediationProperties `pulumi:"properties"`
	// Resource ID.
	ResourceId string `pulumi:"resourceId"`
}

// The set of arguments for constructing a RemediationAtResource resource.
type RemediationAtResourceArgs struct {
	// The name of the remediation.
	Name pulumi.StringInput
	// Properties for the remediation.
	Properties RemediationPropertiesPtrInput
	// Resource ID.
	ResourceId pulumi.StringInput
}

func (RemediationAtResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*remediationAtResourceArgs)(nil)).Elem()
}