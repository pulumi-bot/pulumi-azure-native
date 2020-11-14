// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Threat intelligence resource.
type ThreatIntelligenceIndicator struct {
	pulumi.CustomResourceState

	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The kind of the entity.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Azure resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewThreatIntelligenceIndicator registers a new resource with the given unique name, arguments, and options.
func NewThreatIntelligenceIndicator(ctx *pulumi.Context,
	name string, args *ThreatIntelligenceIndicatorArgs, opts ...pulumi.ResourceOption) (*ThreatIntelligenceIndicator, error) {
	if args == nil || args.Kind == nil {
		return nil, errors.New("missing required argument 'Kind'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.OperationalInsightsResourceProvider == nil {
		return nil, errors.New("missing required argument 'OperationalInsightsResourceProvider'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.WorkspaceName == nil {
		return nil, errors.New("missing required argument 'WorkspaceName'")
	}
	if args == nil {
		args = &ThreatIntelligenceIndicatorArgs{}
	}
	var resource ThreatIntelligenceIndicator
	err := ctx.RegisterResource("azure-nextgen:securityinsights/v20190101preview:ThreatIntelligenceIndicator", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetThreatIntelligenceIndicator gets an existing ThreatIntelligenceIndicator resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetThreatIntelligenceIndicator(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ThreatIntelligenceIndicatorState, opts ...pulumi.ResourceOption) (*ThreatIntelligenceIndicator, error) {
	var resource ThreatIntelligenceIndicator
	err := ctx.ReadResource("azure-nextgen:securityinsights/v20190101preview:ThreatIntelligenceIndicator", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ThreatIntelligenceIndicator resources.
type threatIntelligenceIndicatorState struct {
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// The kind of the entity.
	Kind *string `pulumi:"kind"`
	// Azure resource name
	Name *string `pulumi:"name"`
	// Azure resource type
	Type *string `pulumi:"type"`
}

type ThreatIntelligenceIndicatorState struct {
	// Etag of the azure resource
	Etag pulumi.StringPtrInput
	// The kind of the entity.
	Kind pulumi.StringPtrInput
	// Azure resource name
	Name pulumi.StringPtrInput
	// Azure resource type
	Type pulumi.StringPtrInput
}

func (ThreatIntelligenceIndicatorState) ElementType() reflect.Type {
	return reflect.TypeOf((*threatIntelligenceIndicatorState)(nil)).Elem()
}

type threatIntelligenceIndicatorArgs struct {
	// Confidence of threat intelligence entity
	Confidence *int `pulumi:"confidence"`
	// Created by
	Created *string `pulumi:"created"`
	// Created by reference of threat intelligence entity
	CreatedByRef *string `pulumi:"createdByRef"`
	// Description of a threat intelligence entity
	Description *string `pulumi:"description"`
	// Display name of a threat intelligence entity
	DisplayName *string `pulumi:"displayName"`
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// External ID of threat intelligence entity
	ExternalId *string `pulumi:"externalId"`
	// External References
	ExternalReferences []string `pulumi:"externalReferences"`
	// Granular Markings
	GranularMarkings []ThreatIntelligenceGranularMarkingModel `pulumi:"granularMarkings"`
	// Indicator types of threat intelligence entities
	IndicatorTypes []string `pulumi:"indicatorTypes"`
	// Kill chain phases
	KillChainPhases []ThreatIntelligenceKillChainPhase `pulumi:"killChainPhases"`
	// The kind of the entity.
	Kind string `pulumi:"kind"`
	// Labels  of threat intelligence entity
	Labels []string `pulumi:"labels"`
	// Last updated time in UTC
	LastUpdatedTimeUtc *string `pulumi:"lastUpdatedTimeUtc"`
	// Modified by
	Modified *string `pulumi:"modified"`
	// Threat Intelligence Identifier
	Name string `pulumi:"name"`
	// The namespace of workspaces resource provider- Microsoft.OperationalInsights.
	OperationalInsightsResourceProvider string `pulumi:"operationalInsightsResourceProvider"`
	// Pattern of a threat intelligence entity
	Pattern *string `pulumi:"pattern"`
	// Pattern type of a threat intelligence entity
	PatternType *string `pulumi:"patternType"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Is threat intelligence entity revoked
	Revoked *bool `pulumi:"revoked"`
	// Source of a threat intelligence entity
	Source *string `pulumi:"source"`
	// List of tags
	ThreatIntelligenceTags []string `pulumi:"threatIntelligenceTags"`
	// Threat types
	ThreatTypes []string `pulumi:"threatTypes"`
	// Valid from
	ValidFrom *string `pulumi:"validFrom"`
	// Valid until
	ValidUntil *string `pulumi:"validUntil"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a ThreatIntelligenceIndicator resource.
type ThreatIntelligenceIndicatorArgs struct {
	// Confidence of threat intelligence entity
	Confidence pulumi.IntPtrInput
	// Created by
	Created pulumi.StringPtrInput
	// Created by reference of threat intelligence entity
	CreatedByRef pulumi.StringPtrInput
	// Description of a threat intelligence entity
	Description pulumi.StringPtrInput
	// Display name of a threat intelligence entity
	DisplayName pulumi.StringPtrInput
	// Etag of the azure resource
	Etag pulumi.StringPtrInput
	// External ID of threat intelligence entity
	ExternalId pulumi.StringPtrInput
	// External References
	ExternalReferences pulumi.StringArrayInput
	// Granular Markings
	GranularMarkings ThreatIntelligenceGranularMarkingModelArrayInput
	// Indicator types of threat intelligence entities
	IndicatorTypes pulumi.StringArrayInput
	// Kill chain phases
	KillChainPhases ThreatIntelligenceKillChainPhaseArrayInput
	// The kind of the entity.
	Kind pulumi.StringInput
	// Labels  of threat intelligence entity
	Labels pulumi.StringArrayInput
	// Last updated time in UTC
	LastUpdatedTimeUtc pulumi.StringPtrInput
	// Modified by
	Modified pulumi.StringPtrInput
	// Threat Intelligence Identifier
	Name pulumi.StringInput
	// The namespace of workspaces resource provider- Microsoft.OperationalInsights.
	OperationalInsightsResourceProvider pulumi.StringInput
	// Pattern of a threat intelligence entity
	Pattern pulumi.StringPtrInput
	// Pattern type of a threat intelligence entity
	PatternType pulumi.StringPtrInput
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Is threat intelligence entity revoked
	Revoked pulumi.BoolPtrInput
	// Source of a threat intelligence entity
	Source pulumi.StringPtrInput
	// List of tags
	ThreatIntelligenceTags pulumi.StringArrayInput
	// Threat types
	ThreatTypes pulumi.StringArrayInput
	// Valid from
	ValidFrom pulumi.StringPtrInput
	// Valid until
	ValidUntil pulumi.StringPtrInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (ThreatIntelligenceIndicatorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*threatIntelligenceIndicatorArgs)(nil)).Elem()
}

type ThreatIntelligenceIndicatorInput interface {
	pulumi.Input

	ToThreatIntelligenceIndicatorOutput() ThreatIntelligenceIndicatorOutput
	ToThreatIntelligenceIndicatorOutputWithContext(ctx context.Context) ThreatIntelligenceIndicatorOutput
}

func (ThreatIntelligenceIndicator) ElementType() reflect.Type {
	return reflect.TypeOf((*ThreatIntelligenceIndicator)(nil)).Elem()
}

func (i ThreatIntelligenceIndicator) ToThreatIntelligenceIndicatorOutput() ThreatIntelligenceIndicatorOutput {
	return i.ToThreatIntelligenceIndicatorOutputWithContext(context.Background())
}

func (i ThreatIntelligenceIndicator) ToThreatIntelligenceIndicatorOutputWithContext(ctx context.Context) ThreatIntelligenceIndicatorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThreatIntelligenceIndicatorOutput)
}

type ThreatIntelligenceIndicatorOutput struct {
	*pulumi.OutputState
}

func (ThreatIntelligenceIndicatorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ThreatIntelligenceIndicatorOutput)(nil)).Elem()
}

func (o ThreatIntelligenceIndicatorOutput) ToThreatIntelligenceIndicatorOutput() ThreatIntelligenceIndicatorOutput {
	return o
}

func (o ThreatIntelligenceIndicatorOutput) ToThreatIntelligenceIndicatorOutputWithContext(ctx context.Context) ThreatIntelligenceIndicatorOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ThreatIntelligenceIndicatorOutput{})
}
