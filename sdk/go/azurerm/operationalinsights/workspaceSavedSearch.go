// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package operationalinsights

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Value object for saved search results.
type WorkspaceSavedSearch struct {
	pulumi.CustomResourceState

	// The ETag of the saved search.
	ETag pulumi.StringPtrOutput `pulumi:"eTag"`
	// The name of the saved search.
	Name pulumi.StringOutput `pulumi:"name"`
	// The properties of the saved search.
	Properties SavedSearchPropertiesResponseOutput `pulumi:"properties"`
	// The type of the saved search.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWorkspaceSavedSearch registers a new resource with the given unique name, arguments, and options.
func NewWorkspaceSavedSearch(ctx *pulumi.Context,
	name string, args *WorkspaceSavedSearchArgs, opts ...pulumi.ResourceOption) (*WorkspaceSavedSearch, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.Properties == nil {
		return nil, errors.New("missing required argument 'Properties'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.WorkspaceName == nil {
		return nil, errors.New("missing required argument 'WorkspaceName'")
	}
	if args == nil {
		args = &WorkspaceSavedSearchArgs{}
	}
	var resource WorkspaceSavedSearch
	err := ctx.RegisterResource("azurerm:operationalinsights:WorkspaceSavedSearch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkspaceSavedSearch gets an existing WorkspaceSavedSearch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkspaceSavedSearch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceSavedSearchState, opts ...pulumi.ResourceOption) (*WorkspaceSavedSearch, error) {
	var resource WorkspaceSavedSearch
	err := ctx.ReadResource("azurerm:operationalinsights:WorkspaceSavedSearch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkspaceSavedSearch resources.
type workspaceSavedSearchState struct {
	// The ETag of the saved search.
	ETag *string `pulumi:"eTag"`
	// The name of the saved search.
	Name *string `pulumi:"name"`
	// The properties of the saved search.
	Properties *SavedSearchPropertiesResponse `pulumi:"properties"`
	// The type of the saved search.
	Type *string `pulumi:"type"`
}

type WorkspaceSavedSearchState struct {
	// The ETag of the saved search.
	ETag pulumi.StringPtrInput
	// The name of the saved search.
	Name pulumi.StringPtrInput
	// The properties of the saved search.
	Properties SavedSearchPropertiesResponsePtrInput
	// The type of the saved search.
	Type pulumi.StringPtrInput
}

func (WorkspaceSavedSearchState) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceSavedSearchState)(nil)).Elem()
}

type workspaceSavedSearchArgs struct {
	// The ETag of the saved search.
	ETag *string `pulumi:"eTag"`
	// The id of the saved search.
	Name string `pulumi:"name"`
	// The properties of the saved search.
	Properties SavedSearchProperties `pulumi:"properties"`
	// The Resource Group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Log Analytics Workspace name.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a WorkspaceSavedSearch resource.
type WorkspaceSavedSearchArgs struct {
	// The ETag of the saved search.
	ETag pulumi.StringPtrInput
	// The id of the saved search.
	Name pulumi.StringInput
	// The properties of the saved search.
	Properties SavedSearchPropertiesInput
	// The Resource Group name.
	ResourceGroupName pulumi.StringInput
	// The Log Analytics Workspace name.
	WorkspaceName pulumi.StringInput
}

func (WorkspaceSavedSearchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceSavedSearchArgs)(nil)).Elem()
}