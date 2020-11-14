// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Represents a Watchlist in Azure Security Insights.
type Watchlist struct {
	pulumi.CustomResourceState

	// The content type of the raw content. Example : text/csv or text/tsv
	ContentType pulumi.StringPtrOutput `pulumi:"contentType"`
	// The time the watchlist was created
	Created pulumi.StringPtrOutput `pulumi:"created"`
	// Describes a user that created the watchlist
	CreatedBy UserInfoResponsePtrOutput `pulumi:"createdBy"`
	// The default duration of a watchlist (in ISO 8601 duration format)
	DefaultDuration pulumi.StringPtrOutput `pulumi:"defaultDuration"`
	// A description of the watchlist
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The display name of the watchlist
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// A flag that indicates if the watchlist is deleted or not
	IsDeleted pulumi.BoolPtrOutput `pulumi:"isDeleted"`
	// List of labels relevant to this watchlist
	Labels pulumi.StringArrayOutput `pulumi:"labels"`
	// Azure resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// The number of lines in a csv/tsv content to skip before the header
	NumberOfLinesToSkip pulumi.IntPtrOutput `pulumi:"numberOfLinesToSkip"`
	// The provider of the watchlist
	Provider pulumi.StringOutput `pulumi:"provider"`
	// The raw content that represents to watchlist items to create. In case of csv/tsv content type, it's the content of the file that will parsed by the endpoint
	RawContent pulumi.StringPtrOutput `pulumi:"rawContent"`
	// The source of the watchlist
	Source pulumi.StringOutput `pulumi:"source"`
	// The tenantId where the watchlist belongs to
	TenantId pulumi.StringPtrOutput `pulumi:"tenantId"`
	// Azure resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// The last time the watchlist was updated
	Updated pulumi.StringPtrOutput `pulumi:"updated"`
	// Describes a user that updated the watchlist
	UpdatedBy UserInfoResponsePtrOutput `pulumi:"updatedBy"`
	// The alias of the watchlist
	WatchlistAlias pulumi.StringPtrOutput `pulumi:"watchlistAlias"`
	// The id (a Guid) of the watchlist
	WatchlistId pulumi.StringPtrOutput `pulumi:"watchlistId"`
	// The type of the watchlist
	WatchlistType pulumi.StringPtrOutput `pulumi:"watchlistType"`
}

// NewWatchlist registers a new resource with the given unique name, arguments, and options.
func NewWatchlist(ctx *pulumi.Context,
	name string, args *WatchlistArgs, opts ...pulumi.ResourceOption) (*Watchlist, error) {
	if args == nil || args.DisplayName == nil {
		return nil, errors.New("missing required argument 'DisplayName'")
	}
	if args == nil || args.OperationalInsightsResourceProvider == nil {
		return nil, errors.New("missing required argument 'OperationalInsightsResourceProvider'")
	}
	if args == nil || args.Provider == nil {
		return nil, errors.New("missing required argument 'Provider'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Source == nil {
		return nil, errors.New("missing required argument 'Source'")
	}
	if args == nil || args.WatchlistAlias == nil {
		return nil, errors.New("missing required argument 'WatchlistAlias'")
	}
	if args == nil || args.WorkspaceName == nil {
		return nil, errors.New("missing required argument 'WorkspaceName'")
	}
	if args == nil {
		args = &WatchlistArgs{}
	}
	var resource Watchlist
	err := ctx.RegisterResource("azure-nextgen:securityinsights/v20190101preview:Watchlist", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWatchlist gets an existing Watchlist resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWatchlist(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WatchlistState, opts ...pulumi.ResourceOption) (*Watchlist, error) {
	var resource Watchlist
	err := ctx.ReadResource("azure-nextgen:securityinsights/v20190101preview:Watchlist", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Watchlist resources.
type watchlistState struct {
	// The content type of the raw content. Example : text/csv or text/tsv
	ContentType *string `pulumi:"contentType"`
	// The time the watchlist was created
	Created *string `pulumi:"created"`
	// Describes a user that created the watchlist
	CreatedBy *UserInfoResponse `pulumi:"createdBy"`
	// The default duration of a watchlist (in ISO 8601 duration format)
	DefaultDuration *string `pulumi:"defaultDuration"`
	// A description of the watchlist
	Description *string `pulumi:"description"`
	// The display name of the watchlist
	DisplayName *string `pulumi:"displayName"`
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// A flag that indicates if the watchlist is deleted or not
	IsDeleted *bool `pulumi:"isDeleted"`
	// List of labels relevant to this watchlist
	Labels []string `pulumi:"labels"`
	// Azure resource name
	Name *string `pulumi:"name"`
	// The number of lines in a csv/tsv content to skip before the header
	NumberOfLinesToSkip *int `pulumi:"numberOfLinesToSkip"`
	// The provider of the watchlist
	Provider *string `pulumi:"provider"`
	// The raw content that represents to watchlist items to create. In case of csv/tsv content type, it's the content of the file that will parsed by the endpoint
	RawContent *string `pulumi:"rawContent"`
	// The source of the watchlist
	Source *string `pulumi:"source"`
	// The tenantId where the watchlist belongs to
	TenantId *string `pulumi:"tenantId"`
	// Azure resource type
	Type *string `pulumi:"type"`
	// The last time the watchlist was updated
	Updated *string `pulumi:"updated"`
	// Describes a user that updated the watchlist
	UpdatedBy *UserInfoResponse `pulumi:"updatedBy"`
	// The alias of the watchlist
	WatchlistAlias *string `pulumi:"watchlistAlias"`
	// The id (a Guid) of the watchlist
	WatchlistId *string `pulumi:"watchlistId"`
	// The type of the watchlist
	WatchlistType *string `pulumi:"watchlistType"`
}

type WatchlistState struct {
	// The content type of the raw content. Example : text/csv or text/tsv
	ContentType pulumi.StringPtrInput
	// The time the watchlist was created
	Created pulumi.StringPtrInput
	// Describes a user that created the watchlist
	CreatedBy UserInfoResponsePtrInput
	// The default duration of a watchlist (in ISO 8601 duration format)
	DefaultDuration pulumi.StringPtrInput
	// A description of the watchlist
	Description pulumi.StringPtrInput
	// The display name of the watchlist
	DisplayName pulumi.StringPtrInput
	// Etag of the azure resource
	Etag pulumi.StringPtrInput
	// A flag that indicates if the watchlist is deleted or not
	IsDeleted pulumi.BoolPtrInput
	// List of labels relevant to this watchlist
	Labels pulumi.StringArrayInput
	// Azure resource name
	Name pulumi.StringPtrInput
	// The number of lines in a csv/tsv content to skip before the header
	NumberOfLinesToSkip pulumi.IntPtrInput
	// The provider of the watchlist
	Provider pulumi.StringPtrInput
	// The raw content that represents to watchlist items to create. In case of csv/tsv content type, it's the content of the file that will parsed by the endpoint
	RawContent pulumi.StringPtrInput
	// The source of the watchlist
	Source pulumi.StringPtrInput
	// The tenantId where the watchlist belongs to
	TenantId pulumi.StringPtrInput
	// Azure resource type
	Type pulumi.StringPtrInput
	// The last time the watchlist was updated
	Updated pulumi.StringPtrInput
	// Describes a user that updated the watchlist
	UpdatedBy UserInfoResponsePtrInput
	// The alias of the watchlist
	WatchlistAlias pulumi.StringPtrInput
	// The id (a Guid) of the watchlist
	WatchlistId pulumi.StringPtrInput
	// The type of the watchlist
	WatchlistType pulumi.StringPtrInput
}

func (WatchlistState) ElementType() reflect.Type {
	return reflect.TypeOf((*watchlistState)(nil)).Elem()
}

type watchlistArgs struct {
	// The content type of the raw content. Example : text/csv or text/tsv
	ContentType *string `pulumi:"contentType"`
	// The time the watchlist was created
	Created *string `pulumi:"created"`
	// Describes a user that created the watchlist
	CreatedBy *UserInfo `pulumi:"createdBy"`
	// The default duration of a watchlist (in ISO 8601 duration format)
	DefaultDuration *string `pulumi:"defaultDuration"`
	// A description of the watchlist
	Description *string `pulumi:"description"`
	// The display name of the watchlist
	DisplayName string `pulumi:"displayName"`
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// A flag that indicates if the watchlist is deleted or not
	IsDeleted *bool `pulumi:"isDeleted"`
	// List of labels relevant to this watchlist
	Labels []string `pulumi:"labels"`
	// The number of lines in a csv/tsv content to skip before the header
	NumberOfLinesToSkip *int `pulumi:"numberOfLinesToSkip"`
	// The namespace of workspaces resource provider- Microsoft.OperationalInsights.
	OperationalInsightsResourceProvider string `pulumi:"operationalInsightsResourceProvider"`
	// The provider of the watchlist
	Provider string `pulumi:"provider"`
	// The raw content that represents to watchlist items to create. In case of csv/tsv content type, it's the content of the file that will parsed by the endpoint
	RawContent *string `pulumi:"rawContent"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The source of the watchlist
	Source string `pulumi:"source"`
	// The tenantId where the watchlist belongs to
	TenantId *string `pulumi:"tenantId"`
	// The last time the watchlist was updated
	Updated *string `pulumi:"updated"`
	// Describes a user that updated the watchlist
	UpdatedBy *UserInfo `pulumi:"updatedBy"`
	// The alias of the watchlist
	WatchlistAlias string `pulumi:"watchlistAlias"`
	// The id (a Guid) of the watchlist
	WatchlistId *string `pulumi:"watchlistId"`
	// The type of the watchlist
	WatchlistType *string `pulumi:"watchlistType"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a Watchlist resource.
type WatchlistArgs struct {
	// The content type of the raw content. Example : text/csv or text/tsv
	ContentType pulumi.StringPtrInput
	// The time the watchlist was created
	Created pulumi.StringPtrInput
	// Describes a user that created the watchlist
	CreatedBy UserInfoPtrInput
	// The default duration of a watchlist (in ISO 8601 duration format)
	DefaultDuration pulumi.StringPtrInput
	// A description of the watchlist
	Description pulumi.StringPtrInput
	// The display name of the watchlist
	DisplayName pulumi.StringInput
	// Etag of the azure resource
	Etag pulumi.StringPtrInput
	// A flag that indicates if the watchlist is deleted or not
	IsDeleted pulumi.BoolPtrInput
	// List of labels relevant to this watchlist
	Labels pulumi.StringArrayInput
	// The number of lines in a csv/tsv content to skip before the header
	NumberOfLinesToSkip pulumi.IntPtrInput
	// The namespace of workspaces resource provider- Microsoft.OperationalInsights.
	OperationalInsightsResourceProvider pulumi.StringInput
	// The provider of the watchlist
	Provider pulumi.StringInput
	// The raw content that represents to watchlist items to create. In case of csv/tsv content type, it's the content of the file that will parsed by the endpoint
	RawContent pulumi.StringPtrInput
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The source of the watchlist
	Source pulumi.StringInput
	// The tenantId where the watchlist belongs to
	TenantId pulumi.StringPtrInput
	// The last time the watchlist was updated
	Updated pulumi.StringPtrInput
	// Describes a user that updated the watchlist
	UpdatedBy UserInfoPtrInput
	// The alias of the watchlist
	WatchlistAlias pulumi.StringInput
	// The id (a Guid) of the watchlist
	WatchlistId pulumi.StringPtrInput
	// The type of the watchlist
	WatchlistType pulumi.StringPtrInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (WatchlistArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*watchlistArgs)(nil)).Elem()
}

type WatchlistInput interface {
	pulumi.Input

	ToWatchlistOutput() WatchlistOutput
	ToWatchlistOutputWithContext(ctx context.Context) WatchlistOutput
}

func (Watchlist) ElementType() reflect.Type {
	return reflect.TypeOf((*Watchlist)(nil)).Elem()
}

func (i Watchlist) ToWatchlistOutput() WatchlistOutput {
	return i.ToWatchlistOutputWithContext(context.Background())
}

func (i Watchlist) ToWatchlistOutputWithContext(ctx context.Context) WatchlistOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WatchlistOutput)
}

type WatchlistOutput struct {
	*pulumi.OutputState
}

func (WatchlistOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WatchlistOutput)(nil)).Elem()
}

func (o WatchlistOutput) ToWatchlistOutput() WatchlistOutput {
	return o
}

func (o WatchlistOutput) ToWatchlistOutputWithContext(ctx context.Context) WatchlistOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WatchlistOutput{})
}
