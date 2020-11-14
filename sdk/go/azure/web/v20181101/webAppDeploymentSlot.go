// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20181101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// User credentials used for publishing activity.
type WebAppDeploymentSlot struct {
	pulumi.CustomResourceState

	// True if deployment is currently active, false if completed and null if not started.
	Active pulumi.BoolPtrOutput `pulumi:"active"`
	// Who authored the deployment.
	Author pulumi.StringPtrOutput `pulumi:"author"`
	// Author email.
	AuthorEmail pulumi.StringPtrOutput `pulumi:"authorEmail"`
	// Who performed the deployment.
	Deployer pulumi.StringPtrOutput `pulumi:"deployer"`
	// Details on deployment.
	Details pulumi.StringPtrOutput `pulumi:"details"`
	// End time.
	EndTime pulumi.StringPtrOutput `pulumi:"endTime"`
	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Details about deployment status.
	Message pulumi.StringPtrOutput `pulumi:"message"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Start time.
	StartTime pulumi.StringPtrOutput `pulumi:"startTime"`
	// Deployment status.
	Status pulumi.IntPtrOutput `pulumi:"status"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWebAppDeploymentSlot registers a new resource with the given unique name, arguments, and options.
func NewWebAppDeploymentSlot(ctx *pulumi.Context,
	name string, args *WebAppDeploymentSlotArgs, opts ...pulumi.ResourceOption) (*WebAppDeploymentSlot, error) {
	if args == nil || args.Id == nil {
		return nil, errors.New("missing required argument 'Id'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Slot == nil {
		return nil, errors.New("missing required argument 'Slot'")
	}
	if args == nil {
		args = &WebAppDeploymentSlotArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:web/latest:WebAppDeploymentSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20150801:WebAppDeploymentSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20160801:WebAppDeploymentSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20180201:WebAppDeploymentSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20190801:WebAppDeploymentSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20200601:WebAppDeploymentSlot"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppDeploymentSlot
	err := ctx.RegisterResource("azure-nextgen:web/v20181101:WebAppDeploymentSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebAppDeploymentSlot gets an existing WebAppDeploymentSlot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAppDeploymentSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppDeploymentSlotState, opts ...pulumi.ResourceOption) (*WebAppDeploymentSlot, error) {
	var resource WebAppDeploymentSlot
	err := ctx.ReadResource("azure-nextgen:web/v20181101:WebAppDeploymentSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAppDeploymentSlot resources.
type webAppDeploymentSlotState struct {
	// True if deployment is currently active, false if completed and null if not started.
	Active *bool `pulumi:"active"`
	// Who authored the deployment.
	Author *string `pulumi:"author"`
	// Author email.
	AuthorEmail *string `pulumi:"authorEmail"`
	// Who performed the deployment.
	Deployer *string `pulumi:"deployer"`
	// Details on deployment.
	Details *string `pulumi:"details"`
	// End time.
	EndTime *string `pulumi:"endTime"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Details about deployment status.
	Message *string `pulumi:"message"`
	// Resource Name.
	Name *string `pulumi:"name"`
	// Start time.
	StartTime *string `pulumi:"startTime"`
	// Deployment status.
	Status *int `pulumi:"status"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type WebAppDeploymentSlotState struct {
	// True if deployment is currently active, false if completed and null if not started.
	Active pulumi.BoolPtrInput
	// Who authored the deployment.
	Author pulumi.StringPtrInput
	// Author email.
	AuthorEmail pulumi.StringPtrInput
	// Who performed the deployment.
	Deployer pulumi.StringPtrInput
	// Details on deployment.
	Details pulumi.StringPtrInput
	// End time.
	EndTime pulumi.StringPtrInput
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Details about deployment status.
	Message pulumi.StringPtrInput
	// Resource Name.
	Name pulumi.StringPtrInput
	// Start time.
	StartTime pulumi.StringPtrInput
	// Deployment status.
	Status pulumi.IntPtrInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (WebAppDeploymentSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppDeploymentSlotState)(nil)).Elem()
}

type webAppDeploymentSlotArgs struct {
	// True if deployment is currently active, false if completed and null if not started.
	Active *bool `pulumi:"active"`
	// Who authored the deployment.
	Author *string `pulumi:"author"`
	// Author email.
	AuthorEmail *string `pulumi:"authorEmail"`
	// Who performed the deployment.
	Deployer *string `pulumi:"deployer"`
	// Details on deployment.
	Details *string `pulumi:"details"`
	// End time.
	EndTime *string `pulumi:"endTime"`
	// ID of an existing deployment.
	Id string `pulumi:"id"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Details about deployment status.
	Message *string `pulumi:"message"`
	// Name of the app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the deployment slot. If a slot is not specified, the API creates a deployment for the production slot.
	Slot string `pulumi:"slot"`
	// Start time.
	StartTime *string `pulumi:"startTime"`
	// Deployment status.
	Status *int `pulumi:"status"`
}

// The set of arguments for constructing a WebAppDeploymentSlot resource.
type WebAppDeploymentSlotArgs struct {
	// True if deployment is currently active, false if completed and null if not started.
	Active pulumi.BoolPtrInput
	// Who authored the deployment.
	Author pulumi.StringPtrInput
	// Author email.
	AuthorEmail pulumi.StringPtrInput
	// Who performed the deployment.
	Deployer pulumi.StringPtrInput
	// Details on deployment.
	Details pulumi.StringPtrInput
	// End time.
	EndTime pulumi.StringPtrInput
	// ID of an existing deployment.
	Id pulumi.StringInput
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Details about deployment status.
	Message pulumi.StringPtrInput
	// Name of the app.
	Name pulumi.StringInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
	// Name of the deployment slot. If a slot is not specified, the API creates a deployment for the production slot.
	Slot pulumi.StringInput
	// Start time.
	StartTime pulumi.StringPtrInput
	// Deployment status.
	Status pulumi.IntPtrInput
}

func (WebAppDeploymentSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppDeploymentSlotArgs)(nil)).Elem()
}

type WebAppDeploymentSlotInput interface {
	pulumi.Input

	ToWebAppDeploymentSlotOutput() WebAppDeploymentSlotOutput
	ToWebAppDeploymentSlotOutputWithContext(ctx context.Context) WebAppDeploymentSlotOutput
}

func (WebAppDeploymentSlot) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAppDeploymentSlot)(nil)).Elem()
}

func (i WebAppDeploymentSlot) ToWebAppDeploymentSlotOutput() WebAppDeploymentSlotOutput {
	return i.ToWebAppDeploymentSlotOutputWithContext(context.Background())
}

func (i WebAppDeploymentSlot) ToWebAppDeploymentSlotOutputWithContext(ctx context.Context) WebAppDeploymentSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppDeploymentSlotOutput)
}

type WebAppDeploymentSlotOutput struct {
	*pulumi.OutputState
}

func (WebAppDeploymentSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAppDeploymentSlotOutput)(nil)).Elem()
}

func (o WebAppDeploymentSlotOutput) ToWebAppDeploymentSlotOutput() WebAppDeploymentSlotOutput {
	return o
}

func (o WebAppDeploymentSlotOutput) ToWebAppDeploymentSlotOutputWithContext(ctx context.Context) WebAppDeploymentSlotOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WebAppDeploymentSlotOutput{})
}
