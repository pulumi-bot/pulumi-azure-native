// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20160401

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Describes a DNS zone.
type Zone struct {
	pulumi.CustomResourceState

	// The etag of the zone.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The maximum number of record sets that can be created in this DNS zone.  This is a read-only property and any attempt to set this value will be ignored.
	MaxNumberOfRecordSets pulumi.IntPtrOutput `pulumi:"maxNumberOfRecordSets"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The name servers for this DNS zone. This is a read-only property and any attempt to set this value will be ignored.
	NameServers pulumi.StringArrayOutput `pulumi:"nameServers"`
	// The current number of record sets in this DNS zone.  This is a read-only property and any attempt to set this value will be ignored.
	NumberOfRecordSets pulumi.IntPtrOutput `pulumi:"numberOfRecordSets"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewZone registers a new resource with the given unique name, arguments, and options.
func NewZone(ctx *pulumi.Context,
	name string, args *ZoneArgs, opts ...pulumi.ResourceOption) (*Zone, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ZoneName == nil {
		return nil, errors.New("missing required argument 'ZoneName'")
	}
	if args == nil {
		args = &ZoneArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:network/latest:Zone"),
		},
		{
			Type: pulumi.String("azurerm:network/v20150504preview:Zone"),
		},
		{
			Type: pulumi.String("azurerm:network/v20170901:Zone"),
		},
		{
			Type: pulumi.String("azurerm:network/v20171001:Zone"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180301preview:Zone"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180501:Zone"),
		},
	})
	opts = append(opts, aliases)
	var resource Zone
	err := ctx.RegisterResource("azurerm:network/v20160401:Zone", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetZone gets an existing Zone resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetZone(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ZoneState, opts ...pulumi.ResourceOption) (*Zone, error) {
	var resource Zone
	err := ctx.ReadResource("azurerm:network/v20160401:Zone", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Zone resources.
type zoneState struct {
	// The etag of the zone.
	Etag *string `pulumi:"etag"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The maximum number of record sets that can be created in this DNS zone.  This is a read-only property and any attempt to set this value will be ignored.
	MaxNumberOfRecordSets *int `pulumi:"maxNumberOfRecordSets"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// The name servers for this DNS zone. This is a read-only property and any attempt to set this value will be ignored.
	NameServers []string `pulumi:"nameServers"`
	// The current number of record sets in this DNS zone.  This is a read-only property and any attempt to set this value will be ignored.
	NumberOfRecordSets *int `pulumi:"numberOfRecordSets"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `pulumi:"type"`
}

type ZoneState struct {
	// The etag of the zone.
	Etag pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The maximum number of record sets that can be created in this DNS zone.  This is a read-only property and any attempt to set this value will be ignored.
	MaxNumberOfRecordSets pulumi.IntPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// The name servers for this DNS zone. This is a read-only property and any attempt to set this value will be ignored.
	NameServers pulumi.StringArrayInput
	// The current number of record sets in this DNS zone.  This is a read-only property and any attempt to set this value will be ignored.
	NumberOfRecordSets pulumi.IntPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringPtrInput
}

func (ZoneState) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneState)(nil)).Elem()
}

type zoneArgs struct {
	// The etag of the zone.
	Etag *string `pulumi:"etag"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The maximum number of record sets that can be created in this DNS zone.  This is a read-only property and any attempt to set this value will be ignored.
	MaxNumberOfRecordSets *int `pulumi:"maxNumberOfRecordSets"`
	// The current number of record sets in this DNS zone.  This is a read-only property and any attempt to set this value will be ignored.
	NumberOfRecordSets *int `pulumi:"numberOfRecordSets"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The name of the DNS zone (without a terminating dot).
	ZoneName string `pulumi:"zoneName"`
}

// The set of arguments for constructing a Zone resource.
type ZoneArgs struct {
	// The etag of the zone.
	Etag pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringInput
	// The maximum number of record sets that can be created in this DNS zone.  This is a read-only property and any attempt to set this value will be ignored.
	MaxNumberOfRecordSets pulumi.IntPtrInput
	// The current number of record sets in this DNS zone.  This is a read-only property and any attempt to set this value will be ignored.
	NumberOfRecordSets pulumi.IntPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The name of the DNS zone (without a terminating dot).
	ZoneName pulumi.StringInput
}

func (ZoneArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneArgs)(nil)).Elem()
}
