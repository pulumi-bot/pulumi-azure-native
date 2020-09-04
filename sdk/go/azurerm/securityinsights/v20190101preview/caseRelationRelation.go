// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190101preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Represents a case relation
type CaseRelationRelation struct {
	pulumi.CustomResourceState

	// The case related bookmark id
	BookmarkId pulumi.StringOutput `pulumi:"bookmarkId"`
	// The case related bookmark name
	BookmarkName pulumi.StringPtrOutput `pulumi:"bookmarkName"`
	// The case identifier
	CaseIdentifier pulumi.StringOutput `pulumi:"caseIdentifier"`
	// ETag for relation
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The type of relation node
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Azure resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Name of relation
	RelationName pulumi.StringOutput `pulumi:"relationName"`
	// Azure resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCaseRelationRelation registers a new resource with the given unique name, arguments, and options.
func NewCaseRelationRelation(ctx *pulumi.Context,
	name string, args *CaseRelationRelationArgs, opts ...pulumi.ResourceOption) (*CaseRelationRelation, error) {
	if args == nil || args.CaseId == nil {
		return nil, errors.New("missing required argument 'CaseId'")
	}
	if args == nil || args.OperationalInsightsResourceProvider == nil {
		return nil, errors.New("missing required argument 'OperationalInsightsResourceProvider'")
	}
	if args == nil || args.RelationName == nil {
		return nil, errors.New("missing required argument 'RelationName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.WorkspaceName == nil {
		return nil, errors.New("missing required argument 'WorkspaceName'")
	}
	if args == nil {
		args = &CaseRelationRelationArgs{}
	}
	var resource CaseRelationRelation
	err := ctx.RegisterResource("azurerm:securityinsights/v20190101preview:CaseRelationRelation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCaseRelationRelation gets an existing CaseRelationRelation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCaseRelationRelation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CaseRelationRelationState, opts ...pulumi.ResourceOption) (*CaseRelationRelation, error) {
	var resource CaseRelationRelation
	err := ctx.ReadResource("azurerm:securityinsights/v20190101preview:CaseRelationRelation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CaseRelationRelation resources.
type caseRelationRelationState struct {
	// The case related bookmark id
	BookmarkId *string `pulumi:"bookmarkId"`
	// The case related bookmark name
	BookmarkName *string `pulumi:"bookmarkName"`
	// The case identifier
	CaseIdentifier *string `pulumi:"caseIdentifier"`
	// ETag for relation
	Etag *string `pulumi:"etag"`
	// The type of relation node
	Kind *string `pulumi:"kind"`
	// Azure resource name
	Name *string `pulumi:"name"`
	// Name of relation
	RelationName *string `pulumi:"relationName"`
	// Azure resource type
	Type *string `pulumi:"type"`
}

type CaseRelationRelationState struct {
	// The case related bookmark id
	BookmarkId pulumi.StringPtrInput
	// The case related bookmark name
	BookmarkName pulumi.StringPtrInput
	// The case identifier
	CaseIdentifier pulumi.StringPtrInput
	// ETag for relation
	Etag pulumi.StringPtrInput
	// The type of relation node
	Kind pulumi.StringPtrInput
	// Azure resource name
	Name pulumi.StringPtrInput
	// Name of relation
	RelationName pulumi.StringPtrInput
	// Azure resource type
	Type pulumi.StringPtrInput
}

func (CaseRelationRelationState) ElementType() reflect.Type {
	return reflect.TypeOf((*caseRelationRelationState)(nil)).Elem()
}

type caseRelationRelationArgs struct {
	// Case ID
	CaseId string `pulumi:"caseId"`
	// ETag for relation
	Etag *string `pulumi:"etag"`
	// The namespace of workspaces resource provider- Microsoft.OperationalInsights.
	OperationalInsightsResourceProvider string `pulumi:"operationalInsightsResourceProvider"`
	// Name of relation
	RelationName string `pulumi:"relationName"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Relation source node
	SourceRelationNode *RelationNode `pulumi:"sourceRelationNode"`
	// Relation target node
	TargetRelationNode *RelationNode `pulumi:"targetRelationNode"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a CaseRelationRelation resource.
type CaseRelationRelationArgs struct {
	// Case ID
	CaseId pulumi.StringInput
	// ETag for relation
	Etag pulumi.StringPtrInput
	// The namespace of workspaces resource provider- Microsoft.OperationalInsights.
	OperationalInsightsResourceProvider pulumi.StringInput
	// Name of relation
	RelationName pulumi.StringInput
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Relation source node
	SourceRelationNode RelationNodePtrInput
	// Relation target node
	TargetRelationNode RelationNodePtrInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (CaseRelationRelationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*caseRelationRelationArgs)(nil)).Elem()
}
