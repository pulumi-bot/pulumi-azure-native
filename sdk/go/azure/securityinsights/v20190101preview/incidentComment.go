// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Represents an incident comment
type IncidentComment struct {
	pulumi.CustomResourceState

	// Describes the client that created the comment
	Author ClientInfoResponseOutput `pulumi:"author"`
	// The time the comment was created
	CreatedTimeUtc pulumi.StringOutput `pulumi:"createdTimeUtc"`
	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The time the comment was updated
	LastModifiedTimeUtc pulumi.StringOutput `pulumi:"lastModifiedTimeUtc"`
	// The comment message
	Message pulumi.StringOutput `pulumi:"message"`
	// Azure resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewIncidentComment registers a new resource with the given unique name, arguments, and options.
func NewIncidentComment(ctx *pulumi.Context,
	name string, args *IncidentCommentArgs, opts ...pulumi.ResourceOption) (*IncidentComment, error) {
	if args == nil || args.IncidentCommentId == nil {
		return nil, errors.New("missing required argument 'IncidentCommentId'")
	}
	if args == nil || args.IncidentId == nil {
		return nil, errors.New("missing required argument 'IncidentId'")
	}
	if args == nil || args.Message == nil {
		return nil, errors.New("missing required argument 'Message'")
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
		args = &IncidentCommentArgs{}
	}
	var resource IncidentComment
	err := ctx.RegisterResource("azure-nextgen:securityinsights/v20190101preview:IncidentComment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIncidentComment gets an existing IncidentComment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIncidentComment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IncidentCommentState, opts ...pulumi.ResourceOption) (*IncidentComment, error) {
	var resource IncidentComment
	err := ctx.ReadResource("azure-nextgen:securityinsights/v20190101preview:IncidentComment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IncidentComment resources.
type incidentCommentState struct {
	// Describes the client that created the comment
	Author *ClientInfoResponse `pulumi:"author"`
	// The time the comment was created
	CreatedTimeUtc *string `pulumi:"createdTimeUtc"`
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// The time the comment was updated
	LastModifiedTimeUtc *string `pulumi:"lastModifiedTimeUtc"`
	// The comment message
	Message *string `pulumi:"message"`
	// Azure resource name
	Name *string `pulumi:"name"`
	// Azure resource type
	Type *string `pulumi:"type"`
}

type IncidentCommentState struct {
	// Describes the client that created the comment
	Author ClientInfoResponsePtrInput
	// The time the comment was created
	CreatedTimeUtc pulumi.StringPtrInput
	// Etag of the azure resource
	Etag pulumi.StringPtrInput
	// The time the comment was updated
	LastModifiedTimeUtc pulumi.StringPtrInput
	// The comment message
	Message pulumi.StringPtrInput
	// Azure resource name
	Name pulumi.StringPtrInput
	// Azure resource type
	Type pulumi.StringPtrInput
}

func (IncidentCommentState) ElementType() reflect.Type {
	return reflect.TypeOf((*incidentCommentState)(nil)).Elem()
}

type incidentCommentArgs struct {
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// Incident comment ID
	IncidentCommentId string `pulumi:"incidentCommentId"`
	// Incident ID
	IncidentId string `pulumi:"incidentId"`
	// The comment message
	Message string `pulumi:"message"`
	// The namespace of workspaces resource provider- Microsoft.OperationalInsights.
	OperationalInsightsResourceProvider string `pulumi:"operationalInsightsResourceProvider"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a IncidentComment resource.
type IncidentCommentArgs struct {
	// Etag of the azure resource
	Etag pulumi.StringPtrInput
	// Incident comment ID
	IncidentCommentId pulumi.StringInput
	// Incident ID
	IncidentId pulumi.StringInput
	// The comment message
	Message pulumi.StringInput
	// The namespace of workspaces resource provider- Microsoft.OperationalInsights.
	OperationalInsightsResourceProvider pulumi.StringInput
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (IncidentCommentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*incidentCommentArgs)(nil)).Elem()
}

type IncidentCommentInput interface {
	pulumi.Input

	ToIncidentCommentOutput() IncidentCommentOutput
	ToIncidentCommentOutputWithContext(ctx context.Context) IncidentCommentOutput
}

func (IncidentComment) ElementType() reflect.Type {
	return reflect.TypeOf((*IncidentComment)(nil)).Elem()
}

func (i IncidentComment) ToIncidentCommentOutput() IncidentCommentOutput {
	return i.ToIncidentCommentOutputWithContext(context.Background())
}

func (i IncidentComment) ToIncidentCommentOutputWithContext(ctx context.Context) IncidentCommentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IncidentCommentOutput)
}

type IncidentCommentOutput struct {
	*pulumi.OutputState
}

func (IncidentCommentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IncidentCommentOutput)(nil)).Elem()
}

func (o IncidentCommentOutput) ToIncidentCommentOutput() IncidentCommentOutput {
	return o
}

func (o IncidentCommentOutput) ToIncidentCommentOutputWithContext(ctx context.Context) IncidentCommentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(IncidentCommentOutput{})
}
