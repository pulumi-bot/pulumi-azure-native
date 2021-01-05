// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Single item in List or Get Alias(Disaster Recovery configuration) operation
// Latest API Version: 2017-04-01.
type DisasterRecoveryConfig struct {
	pulumi.CustomResourceState

	// Alternate name specified when alias and namespace names are same.
	AlternateName pulumi.StringPtrOutput `pulumi:"alternateName"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// ARM Id of the Primary/Secondary eventhub namespace name, which is part of GEO DR pairing
	PartnerNamespace pulumi.StringPtrOutput `pulumi:"partnerNamespace"`
	// Number of entities pending to be replicated.
	PendingReplicationOperationsCount pulumi.Float64Output `pulumi:"pendingReplicationOperationsCount"`
	// Provisioning state of the Alias(Disaster Recovery configuration) - possible values 'Accepted' or 'Succeeded' or 'Failed'
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// role of namespace in GEO DR - possible values 'Primary' or 'PrimaryNotReplicating' or 'Secondary'
	Role pulumi.StringOutput `pulumi:"role"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDisasterRecoveryConfig registers a new resource with the given unique name, arguments, and options.
func NewDisasterRecoveryConfig(ctx *pulumi.Context,
	name string, args *DisasterRecoveryConfigArgs, opts ...pulumi.ResourceOption) (*DisasterRecoveryConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Alias == nil {
		return nil, errors.New("invalid value for required argument 'Alias'")
	}
	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:eventhub/v20170401:DisasterRecoveryConfig"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventhub/v20180101preview:DisasterRecoveryConfig"),
		},
	})
	opts = append(opts, aliases)
	var resource DisasterRecoveryConfig
	err := ctx.RegisterResource("azure-nextgen:eventhub/latest:DisasterRecoveryConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDisasterRecoveryConfig gets an existing DisasterRecoveryConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDisasterRecoveryConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DisasterRecoveryConfigState, opts ...pulumi.ResourceOption) (*DisasterRecoveryConfig, error) {
	var resource DisasterRecoveryConfig
	err := ctx.ReadResource("azure-nextgen:eventhub/latest:DisasterRecoveryConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DisasterRecoveryConfig resources.
type disasterRecoveryConfigState struct {
	// Alternate name specified when alias and namespace names are same.
	AlternateName *string `pulumi:"alternateName"`
	// Resource name.
	Name *string `pulumi:"name"`
	// ARM Id of the Primary/Secondary eventhub namespace name, which is part of GEO DR pairing
	PartnerNamespace *string `pulumi:"partnerNamespace"`
	// Number of entities pending to be replicated.
	PendingReplicationOperationsCount *float64 `pulumi:"pendingReplicationOperationsCount"`
	// Provisioning state of the Alias(Disaster Recovery configuration) - possible values 'Accepted' or 'Succeeded' or 'Failed'
	ProvisioningState *string `pulumi:"provisioningState"`
	// role of namespace in GEO DR - possible values 'Primary' or 'PrimaryNotReplicating' or 'Secondary'
	Role *string `pulumi:"role"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type DisasterRecoveryConfigState struct {
	// Alternate name specified when alias and namespace names are same.
	AlternateName pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// ARM Id of the Primary/Secondary eventhub namespace name, which is part of GEO DR pairing
	PartnerNamespace pulumi.StringPtrInput
	// Number of entities pending to be replicated.
	PendingReplicationOperationsCount pulumi.Float64PtrInput
	// Provisioning state of the Alias(Disaster Recovery configuration) - possible values 'Accepted' or 'Succeeded' or 'Failed'
	ProvisioningState pulumi.StringPtrInput
	// role of namespace in GEO DR - possible values 'Primary' or 'PrimaryNotReplicating' or 'Secondary'
	Role pulumi.StringPtrInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (DisasterRecoveryConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*disasterRecoveryConfigState)(nil)).Elem()
}

type disasterRecoveryConfigArgs struct {
	// The Disaster Recovery configuration name
	Alias string `pulumi:"alias"`
	// Alternate name specified when alias and namespace names are same.
	AlternateName *string `pulumi:"alternateName"`
	// The Namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// ARM Id of the Primary/Secondary eventhub namespace name, which is part of GEO DR pairing
	PartnerNamespace *string `pulumi:"partnerNamespace"`
	// Name of the resource group within the azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a DisasterRecoveryConfig resource.
type DisasterRecoveryConfigArgs struct {
	// The Disaster Recovery configuration name
	Alias pulumi.StringInput
	// Alternate name specified when alias and namespace names are same.
	AlternateName pulumi.StringPtrInput
	// The Namespace name
	NamespaceName pulumi.StringInput
	// ARM Id of the Primary/Secondary eventhub namespace name, which is part of GEO DR pairing
	PartnerNamespace pulumi.StringPtrInput
	// Name of the resource group within the azure subscription.
	ResourceGroupName pulumi.StringInput
}

func (DisasterRecoveryConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*disasterRecoveryConfigArgs)(nil)).Elem()
}

type DisasterRecoveryConfigInput interface {
	pulumi.Input

	ToDisasterRecoveryConfigOutput() DisasterRecoveryConfigOutput
	ToDisasterRecoveryConfigOutputWithContext(ctx context.Context) DisasterRecoveryConfigOutput
}

func (*DisasterRecoveryConfig) ElementType() reflect.Type {
	return reflect.TypeOf((*DisasterRecoveryConfig)(nil))
}

func (i *DisasterRecoveryConfig) ToDisasterRecoveryConfigOutput() DisasterRecoveryConfigOutput {
	return i.ToDisasterRecoveryConfigOutputWithContext(context.Background())
}

func (i *DisasterRecoveryConfig) ToDisasterRecoveryConfigOutputWithContext(ctx context.Context) DisasterRecoveryConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DisasterRecoveryConfigOutput)
}

type DisasterRecoveryConfigOutput struct {
	*pulumi.OutputState
}

func (DisasterRecoveryConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DisasterRecoveryConfig)(nil))
}

func (o DisasterRecoveryConfigOutput) ToDisasterRecoveryConfigOutput() DisasterRecoveryConfigOutput {
	return o
}

func (o DisasterRecoveryConfigOutput) ToDisasterRecoveryConfigOutputWithContext(ctx context.Context) DisasterRecoveryConfigOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DisasterRecoveryConfigOutput{})
}
