// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190101preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Represents a relation between two resources
type BookmarkRelationRelation struct {
	pulumi.CustomResourceState

	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Azure resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource ID of the related resource
	RelatedResourceId pulumi.StringOutput `pulumi:"relatedResourceId"`
	// The resource kind of the related resource
	RelatedResourceKind pulumi.StringOutput `pulumi:"relatedResourceKind"`
	// The name of the related resource
	RelatedResourceName pulumi.StringOutput `pulumi:"relatedResourceName"`
	// The resource type of the related resource
	RelatedResourceType pulumi.StringOutput `pulumi:"relatedResourceType"`
	// Azure resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewBookmarkRelationRelation registers a new resource with the given unique name, arguments, and options.
func NewBookmarkRelationRelation(ctx *pulumi.Context,
	name string, args *BookmarkRelationRelationArgs, opts ...pulumi.ResourceOption) (*BookmarkRelationRelation, error) {
	if args == nil || args.BookmarkId == nil {
		return nil, errors.New("missing required argument 'BookmarkId'")
	}
	if args == nil || args.OperationalInsightsResourceProvider == nil {
		return nil, errors.New("missing required argument 'OperationalInsightsResourceProvider'")
	}
	if args == nil || args.RelatedResourceId == nil {
		return nil, errors.New("missing required argument 'RelatedResourceId'")
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
		args = &BookmarkRelationRelationArgs{}
	}
	var resource BookmarkRelationRelation
	err := ctx.RegisterResource("azurerm:securityinsights/v20190101preview:BookmarkRelationRelation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBookmarkRelationRelation gets an existing BookmarkRelationRelation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBookmarkRelationRelation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BookmarkRelationRelationState, opts ...pulumi.ResourceOption) (*BookmarkRelationRelation, error) {
	var resource BookmarkRelationRelation
	err := ctx.ReadResource("azurerm:securityinsights/v20190101preview:BookmarkRelationRelation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BookmarkRelationRelation resources.
type bookmarkRelationRelationState struct {
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// Azure resource name
	Name *string `pulumi:"name"`
	// The resource ID of the related resource
	RelatedResourceId *string `pulumi:"relatedResourceId"`
	// The resource kind of the related resource
	RelatedResourceKind *string `pulumi:"relatedResourceKind"`
	// The name of the related resource
	RelatedResourceName *string `pulumi:"relatedResourceName"`
	// The resource type of the related resource
	RelatedResourceType *string `pulumi:"relatedResourceType"`
	// Azure resource type
	Type *string `pulumi:"type"`
}

type BookmarkRelationRelationState struct {
	// Etag of the azure resource
	Etag pulumi.StringPtrInput
	// Azure resource name
	Name pulumi.StringPtrInput
	// The resource ID of the related resource
	RelatedResourceId pulumi.StringPtrInput
	// The resource kind of the related resource
	RelatedResourceKind pulumi.StringPtrInput
	// The name of the related resource
	RelatedResourceName pulumi.StringPtrInput
	// The resource type of the related resource
	RelatedResourceType pulumi.StringPtrInput
	// Azure resource type
	Type pulumi.StringPtrInput
}

func (BookmarkRelationRelationState) ElementType() reflect.Type {
	return reflect.TypeOf((*bookmarkRelationRelationState)(nil)).Elem()
}

type bookmarkRelationRelationArgs struct {
	// Bookmark ID
	BookmarkId string `pulumi:"bookmarkId"`
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// The namespace of workspaces resource provider- Microsoft.OperationalInsights.
	OperationalInsightsResourceProvider string `pulumi:"operationalInsightsResourceProvider"`
	// The resource ID of the related resource
	RelatedResourceId string `pulumi:"relatedResourceId"`
	// Relation Name
	RelationName string `pulumi:"relationName"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a BookmarkRelationRelation resource.
type BookmarkRelationRelationArgs struct {
	// Bookmark ID
	BookmarkId pulumi.StringInput
	// Etag of the azure resource
	Etag pulumi.StringPtrInput
	// The namespace of workspaces resource provider- Microsoft.OperationalInsights.
	OperationalInsightsResourceProvider pulumi.StringInput
	// The resource ID of the related resource
	RelatedResourceId pulumi.StringInput
	// Relation Name
	RelationName pulumi.StringInput
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (BookmarkRelationRelationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bookmarkRelationRelationArgs)(nil)).Elem()
}
