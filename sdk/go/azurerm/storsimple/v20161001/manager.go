// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20161001

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The StorSimple Manager
type Manager struct {
	pulumi.CustomResourceState

	// ETag of the Manager
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The Geo location of the Manager
	Location pulumi.StringOutput `pulumi:"location"`
	// The Resource Name
	Name pulumi.StringOutput `pulumi:"name"`
	// List of properties of the Manager
	Properties ManagerPropertiesResponseOutput `pulumi:"properties"`
	// Tags attached to the Manager
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewManager registers a new resource with the given unique name, arguments, and options.
func NewManager(ctx *pulumi.Context,
	name string, args *ManagerArgs, opts ...pulumi.ResourceOption) (*Manager, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ManagerArgs{}
	}
	var resource Manager
	err := ctx.RegisterResource("azurerm:storsimple/v20161001:Manager", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManager gets an existing Manager resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManager(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagerState, opts ...pulumi.ResourceOption) (*Manager, error) {
	var resource Manager
	err := ctx.ReadResource("azurerm:storsimple/v20161001:Manager", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Manager resources.
type managerState struct {
	// ETag of the Manager
	Etag *string `pulumi:"etag"`
	// The Geo location of the Manager
	Location *string `pulumi:"location"`
	// The Resource Name
	Name *string `pulumi:"name"`
	// List of properties of the Manager
	Properties *ManagerPropertiesResponse `pulumi:"properties"`
	// Tags attached to the Manager
	Tags map[string]string `pulumi:"tags"`
	// The Resource type
	Type *string `pulumi:"type"`
}

type ManagerState struct {
	// ETag of the Manager
	Etag pulumi.StringPtrInput
	// The Geo location of the Manager
	Location pulumi.StringPtrInput
	// The Resource Name
	Name pulumi.StringPtrInput
	// List of properties of the Manager
	Properties ManagerPropertiesResponsePtrInput
	// Tags attached to the Manager
	Tags pulumi.StringMapInput
	// The Resource type
	Type pulumi.StringPtrInput
}

func (ManagerState) ElementType() reflect.Type {
	return reflect.TypeOf((*managerState)(nil)).Elem()
}

type managerArgs struct {
	// Specifies if the Manager is Garda or Helsinki
	CisIntrinsicSettings *ManagerIntrinsicSettings `pulumi:"cisIntrinsicSettings"`
	// ETag of the Manager
	Etag *string `pulumi:"etag"`
	// The Geo location of the Manager
	Location string `pulumi:"location"`
	// The manager name
	Name string `pulumi:"name"`
	// The resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies the Sku
	Sku *ManagerSku `pulumi:"sku"`
	// Tags attached to the Manager
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Manager resource.
type ManagerArgs struct {
	// Specifies if the Manager is Garda or Helsinki
	CisIntrinsicSettings ManagerIntrinsicSettingsPtrInput
	// ETag of the Manager
	Etag pulumi.StringPtrInput
	// The Geo location of the Manager
	Location pulumi.StringInput
	// The manager name
	Name pulumi.StringInput
	// The resource group name
	ResourceGroupName pulumi.StringInput
	// Specifies the Sku
	Sku ManagerSkuPtrInput
	// Tags attached to the Manager
	Tags pulumi.StringMapInput
}

func (ManagerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managerArgs)(nil)).Elem()
}
