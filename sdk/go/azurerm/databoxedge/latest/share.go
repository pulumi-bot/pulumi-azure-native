// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Represents a share on the  Data Box Edge/Gateway device.
type Share struct {
	pulumi.CustomResourceState

	// Access protocol to be used by the share.
	AccessProtocol pulumi.StringOutput `pulumi:"accessProtocol"`
	// Azure container mapping for the share.
	AzureContainerInfo AzureContainerInfoResponsePtrOutput `pulumi:"azureContainerInfo"`
	// List of IP addresses and corresponding access rights on the share(required for NFS protocol).
	ClientAccessRights ClientAccessRightResponseArrayOutput `pulumi:"clientAccessRights"`
	// Data policy of the share.
	DataPolicy pulumi.StringPtrOutput `pulumi:"dataPolicy"`
	// Description for the share.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Current monitoring status of the share.
	MonitoringStatus pulumi.StringOutput `pulumi:"monitoringStatus"`
	// The object name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Details of the refresh job on this share.
	RefreshDetails RefreshDetailsResponsePtrOutput `pulumi:"refreshDetails"`
	// Share mount point to the role.
	ShareMappings MountPointMapResponseArrayOutput `pulumi:"shareMappings"`
	// Current status of the share.
	ShareStatus pulumi.StringOutput `pulumi:"shareStatus"`
	// The hierarchical type of the object.
	Type pulumi.StringOutput `pulumi:"type"`
	// Mapping of users and corresponding access rights on the share (required for SMB protocol).
	UserAccessRights UserAccessRightResponseArrayOutput `pulumi:"userAccessRights"`
}

// NewShare registers a new resource with the given unique name, arguments, and options.
func NewShare(ctx *pulumi.Context,
	name string, args *ShareArgs, opts ...pulumi.ResourceOption) (*Share, error) {
	if args == nil || args.AccessProtocol == nil {
		return nil, errors.New("missing required argument 'AccessProtocol'")
	}
	if args == nil || args.DeviceName == nil {
		return nil, errors.New("missing required argument 'DeviceName'")
	}
	if args == nil || args.MonitoringStatus == nil {
		return nil, errors.New("missing required argument 'MonitoringStatus'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ShareStatus == nil {
		return nil, errors.New("missing required argument 'ShareStatus'")
	}
	if args == nil {
		args = &ShareArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:databoxedge/v20190301:Share"),
		},
		{
			Type: pulumi.String("azurerm:databoxedge/v20190701:Share"),
		},
		{
			Type: pulumi.String("azurerm:databoxedge/v20190801:Share"),
		},
		{
			Type: pulumi.String("azurerm:databoxedge/v20200501preview:Share"),
		},
	})
	opts = append(opts, aliases)
	var resource Share
	err := ctx.RegisterResource("azurerm:databoxedge/latest:Share", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetShare gets an existing Share resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetShare(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ShareState, opts ...pulumi.ResourceOption) (*Share, error) {
	var resource Share
	err := ctx.ReadResource("azurerm:databoxedge/latest:Share", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Share resources.
type shareState struct {
	// Access protocol to be used by the share.
	AccessProtocol *string `pulumi:"accessProtocol"`
	// Azure container mapping for the share.
	AzureContainerInfo *AzureContainerInfoResponse `pulumi:"azureContainerInfo"`
	// List of IP addresses and corresponding access rights on the share(required for NFS protocol).
	ClientAccessRights []ClientAccessRightResponse `pulumi:"clientAccessRights"`
	// Data policy of the share.
	DataPolicy *string `pulumi:"dataPolicy"`
	// Description for the share.
	Description *string `pulumi:"description"`
	// Current monitoring status of the share.
	MonitoringStatus *string `pulumi:"monitoringStatus"`
	// The object name.
	Name *string `pulumi:"name"`
	// Details of the refresh job on this share.
	RefreshDetails *RefreshDetailsResponse `pulumi:"refreshDetails"`
	// Share mount point to the role.
	ShareMappings []MountPointMapResponse `pulumi:"shareMappings"`
	// Current status of the share.
	ShareStatus *string `pulumi:"shareStatus"`
	// The hierarchical type of the object.
	Type *string `pulumi:"type"`
	// Mapping of users and corresponding access rights on the share (required for SMB protocol).
	UserAccessRights []UserAccessRightResponse `pulumi:"userAccessRights"`
}

type ShareState struct {
	// Access protocol to be used by the share.
	AccessProtocol pulumi.StringPtrInput
	// Azure container mapping for the share.
	AzureContainerInfo AzureContainerInfoResponsePtrInput
	// List of IP addresses and corresponding access rights on the share(required for NFS protocol).
	ClientAccessRights ClientAccessRightResponseArrayInput
	// Data policy of the share.
	DataPolicy pulumi.StringPtrInput
	// Description for the share.
	Description pulumi.StringPtrInput
	// Current monitoring status of the share.
	MonitoringStatus pulumi.StringPtrInput
	// The object name.
	Name pulumi.StringPtrInput
	// Details of the refresh job on this share.
	RefreshDetails RefreshDetailsResponsePtrInput
	// Share mount point to the role.
	ShareMappings MountPointMapResponseArrayInput
	// Current status of the share.
	ShareStatus pulumi.StringPtrInput
	// The hierarchical type of the object.
	Type pulumi.StringPtrInput
	// Mapping of users and corresponding access rights on the share (required for SMB protocol).
	UserAccessRights UserAccessRightResponseArrayInput
}

func (ShareState) ElementType() reflect.Type {
	return reflect.TypeOf((*shareState)(nil)).Elem()
}

type shareArgs struct {
	// Access protocol to be used by the share.
	AccessProtocol string `pulumi:"accessProtocol"`
	// Azure container mapping for the share.
	AzureContainerInfo *AzureContainerInfo `pulumi:"azureContainerInfo"`
	// List of IP addresses and corresponding access rights on the share(required for NFS protocol).
	ClientAccessRights []ClientAccessRight `pulumi:"clientAccessRights"`
	// Data policy of the share.
	DataPolicy *string `pulumi:"dataPolicy"`
	// Description for the share.
	Description *string `pulumi:"description"`
	// The device name.
	DeviceName string `pulumi:"deviceName"`
	// Current monitoring status of the share.
	MonitoringStatus string `pulumi:"monitoringStatus"`
	// The share name.
	Name string `pulumi:"name"`
	// Details of the refresh job on this share.
	RefreshDetails *RefreshDetails `pulumi:"refreshDetails"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Current status of the share.
	ShareStatus string `pulumi:"shareStatus"`
	// Mapping of users and corresponding access rights on the share (required for SMB protocol).
	UserAccessRights []UserAccessRight `pulumi:"userAccessRights"`
}

// The set of arguments for constructing a Share resource.
type ShareArgs struct {
	// Access protocol to be used by the share.
	AccessProtocol pulumi.StringInput
	// Azure container mapping for the share.
	AzureContainerInfo AzureContainerInfoPtrInput
	// List of IP addresses and corresponding access rights on the share(required for NFS protocol).
	ClientAccessRights ClientAccessRightArrayInput
	// Data policy of the share.
	DataPolicy pulumi.StringPtrInput
	// Description for the share.
	Description pulumi.StringPtrInput
	// The device name.
	DeviceName pulumi.StringInput
	// Current monitoring status of the share.
	MonitoringStatus pulumi.StringInput
	// The share name.
	Name pulumi.StringInput
	// Details of the refresh job on this share.
	RefreshDetails RefreshDetailsPtrInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// Current status of the share.
	ShareStatus pulumi.StringInput
	// Mapping of users and corresponding access rights on the share (required for SMB protocol).
	UserAccessRights UserAccessRightArrayInput
}

func (ShareArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*shareArgs)(nil)).Elem()
}
