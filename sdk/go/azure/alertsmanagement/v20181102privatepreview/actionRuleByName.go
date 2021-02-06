// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20181102privatepreview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Action rule object containing target scope, conditions and suppression logic
type ActionRuleByName struct {
	pulumi.CustomResourceState

	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Azure resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Action rule properties defining scope, conditions, suppression logic for action rule
	Properties ActionRulePropertiesResponseOutput `pulumi:"properties"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Azure resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewActionRuleByName registers a new resource with the given unique name, arguments, and options.
func NewActionRuleByName(ctx *pulumi.Context,
	name string, args *ActionRuleByNameArgs, opts ...pulumi.ResourceOption) (*ActionRuleByName, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ActionRuleName == nil {
		return nil, errors.New("invalid value for required argument 'ActionRuleName'")
	}
	if args.ResourceGroup == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroup'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:alertsmanagement/v20190505preview:ActionRuleByName"),
		},
	})
	opts = append(opts, aliases)
	var resource ActionRuleByName
	err := ctx.RegisterResource("azure-nextgen:alertsmanagement/v20181102privatepreview:ActionRuleByName", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetActionRuleByName gets an existing ActionRuleByName resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetActionRuleByName(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ActionRuleByNameState, opts ...pulumi.ResourceOption) (*ActionRuleByName, error) {
	var resource ActionRuleByName
	err := ctx.ReadResource("azure-nextgen:alertsmanagement/v20181102privatepreview:ActionRuleByName", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ActionRuleByName resources.
type actionRuleByNameState struct {
	// Resource location
	Location *string `pulumi:"location"`
	// Azure resource name
	Name *string `pulumi:"name"`
	// Action rule properties defining scope, conditions, suppression logic for action rule
	Properties *ActionRulePropertiesResponse `pulumi:"properties"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Azure resource type
	Type *string `pulumi:"type"`
}

type ActionRuleByNameState struct {
	// Resource location
	Location pulumi.StringPtrInput
	// Azure resource name
	Name pulumi.StringPtrInput
	// Action rule properties defining scope, conditions, suppression logic for action rule
	Properties ActionRulePropertiesResponsePtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Azure resource type
	Type pulumi.StringPtrInput
}

func (ActionRuleByNameState) ElementType() reflect.Type {
	return reflect.TypeOf((*actionRuleByNameState)(nil)).Elem()
}

type actionRuleByNameArgs struct {
	// The name of action rule that needs to be created/updated
	ActionRuleName string `pulumi:"actionRuleName"`
	// Resource location
	Location *string `pulumi:"location"`
	// Action rule properties defining scope, conditions, suppression logic for action rule
	Properties *ActionRuleProperties `pulumi:"properties"`
	// Resource group name where the resource is created.
	ResourceGroup string `pulumi:"resourceGroup"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ActionRuleByName resource.
type ActionRuleByNameArgs struct {
	// The name of action rule that needs to be created/updated
	ActionRuleName pulumi.StringInput
	// Resource location
	Location pulumi.StringPtrInput
	// Action rule properties defining scope, conditions, suppression logic for action rule
	Properties ActionRulePropertiesPtrInput
	// Resource group name where the resource is created.
	ResourceGroup pulumi.StringInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (ActionRuleByNameArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*actionRuleByNameArgs)(nil)).Elem()
}

type ActionRuleByNameInput interface {
	pulumi.Input

	ToActionRuleByNameOutput() ActionRuleByNameOutput
	ToActionRuleByNameOutputWithContext(ctx context.Context) ActionRuleByNameOutput
}

func (*ActionRuleByName) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionRuleByName)(nil))
}

func (i *ActionRuleByName) ToActionRuleByNameOutput() ActionRuleByNameOutput {
	return i.ToActionRuleByNameOutputWithContext(context.Background())
}

func (i *ActionRuleByName) ToActionRuleByNameOutputWithContext(ctx context.Context) ActionRuleByNameOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionRuleByNameOutput)
}

type ActionRuleByNameOutput struct {
	*pulumi.OutputState
}

func (ActionRuleByNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionRuleByName)(nil))
}

func (o ActionRuleByNameOutput) ToActionRuleByNameOutput() ActionRuleByNameOutput {
	return o
}

func (o ActionRuleByNameOutput) ToActionRuleByNameOutputWithContext(ctx context.Context) ActionRuleByNameOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ActionRuleByNameOutput{})
}
