// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20171103preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Specifies the disk information fo the HANA instance
type Disk struct {
	// Specifies the size of an empty data disk in gigabytes.
	DiskSizeGB *int `pulumi:"diskSizeGB"`
	// The disk name.
	Name *string `pulumi:"name"`
}

// DiskInput is an input type that accepts DiskArgs and DiskOutput values.
// You can construct a concrete instance of `DiskInput` via:
//
//          DiskArgs{...}
type DiskInput interface {
	pulumi.Input

	ToDiskOutput() DiskOutput
	ToDiskOutputWithContext(context.Context) DiskOutput
}

// Specifies the disk information fo the HANA instance
type DiskArgs struct {
	// Specifies the size of an empty data disk in gigabytes.
	DiskSizeGB pulumi.IntPtrInput `pulumi:"diskSizeGB"`
	// The disk name.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (DiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Disk)(nil)).Elem()
}

func (i DiskArgs) ToDiskOutput() DiskOutput {
	return i.ToDiskOutputWithContext(context.Background())
}

func (i DiskArgs) ToDiskOutputWithContext(ctx context.Context) DiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskOutput)
}

// DiskArrayInput is an input type that accepts DiskArray and DiskArrayOutput values.
// You can construct a concrete instance of `DiskArrayInput` via:
//
//          DiskArray{ DiskArgs{...} }
type DiskArrayInput interface {
	pulumi.Input

	ToDiskArrayOutput() DiskArrayOutput
	ToDiskArrayOutputWithContext(context.Context) DiskArrayOutput
}

type DiskArray []DiskInput

func (DiskArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Disk)(nil)).Elem()
}

func (i DiskArray) ToDiskArrayOutput() DiskArrayOutput {
	return i.ToDiskArrayOutputWithContext(context.Background())
}

func (i DiskArray) ToDiskArrayOutputWithContext(ctx context.Context) DiskArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskArrayOutput)
}

// Specifies the disk information fo the HANA instance
type DiskResponse struct {
	// Specifies the size of an empty data disk in gigabytes.
	DiskSizeGB *int `pulumi:"diskSizeGB"`
	// Specifies the logical unit number of the data disk. This value is used to identify data disks within the VM and therefore must be unique for each data disk attached to a VM.
	Lun int `pulumi:"lun"`
	// The disk name.
	Name *string `pulumi:"name"`
}

// Specifies the disk information fo the HANA instance
type DiskResponseOutput struct{ *pulumi.OutputState }

func (DiskResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskResponse)(nil)).Elem()
}

func (o DiskResponseOutput) ToDiskResponseOutput() DiskResponseOutput {
	return o
}

func (o DiskResponseOutput) ToDiskResponseOutputWithContext(ctx context.Context) DiskResponseOutput {
	return o
}

// Specifies the size of an empty data disk in gigabytes.
func (o DiskResponseOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DiskResponse) *int { return v.DiskSizeGB }).(pulumi.IntPtrOutput)
}

// Specifies the logical unit number of the data disk. This value is used to identify data disks within the VM and therefore must be unique for each data disk attached to a VM.
func (o DiskResponseOutput) Lun() pulumi.IntOutput {
	return o.ApplyT(func(v DiskResponse) int { return v.Lun }).(pulumi.IntOutput)
}

// The disk name.
func (o DiskResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DiskResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type DiskResponseArrayOutput struct{ *pulumi.OutputState }

func (DiskResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DiskResponse)(nil)).Elem()
}

func (o DiskResponseArrayOutput) ToDiskResponseArrayOutput() DiskResponseArrayOutput {
	return o
}

func (o DiskResponseArrayOutput) ToDiskResponseArrayOutputWithContext(ctx context.Context) DiskResponseArrayOutput {
	return o
}

func (o DiskResponseArrayOutput) Index(i pulumi.IntInput) DiskResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DiskResponse {
		return vs[0].([]DiskResponse)[vs[1].(int)]
	}).(DiskResponseOutput)
}

// Specifies the hardware settings for the HANA instance.
type HardwareProfileResponse struct {
	// Specifies the HANA instance SKU.
	HanaInstanceSize string `pulumi:"hanaInstanceSize"`
	// Name of the hardware type (vendor and/or their product name)
	HardwareType string `pulumi:"hardwareType"`
}

// Specifies the hardware settings for the HANA instance.
type HardwareProfileResponseOutput struct{ *pulumi.OutputState }

func (HardwareProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HardwareProfileResponse)(nil)).Elem()
}

func (o HardwareProfileResponseOutput) ToHardwareProfileResponseOutput() HardwareProfileResponseOutput {
	return o
}

func (o HardwareProfileResponseOutput) ToHardwareProfileResponseOutputWithContext(ctx context.Context) HardwareProfileResponseOutput {
	return o
}

func (o HardwareProfileResponseOutput) ToHardwareProfileResponsePtrOutput() HardwareProfileResponsePtrOutput {
	return o.ToHardwareProfileResponsePtrOutputWithContext(context.Background())
}

func (o HardwareProfileResponseOutput) ToHardwareProfileResponsePtrOutputWithContext(ctx context.Context) HardwareProfileResponsePtrOutput {
	return o.ApplyT(func(v HardwareProfileResponse) *HardwareProfileResponse {
		return &v
	}).(HardwareProfileResponsePtrOutput)
}

// Specifies the HANA instance SKU.
func (o HardwareProfileResponseOutput) HanaInstanceSize() pulumi.StringOutput {
	return o.ApplyT(func(v HardwareProfileResponse) string { return v.HanaInstanceSize }).(pulumi.StringOutput)
}

// Name of the hardware type (vendor and/or their product name)
func (o HardwareProfileResponseOutput) HardwareType() pulumi.StringOutput {
	return o.ApplyT(func(v HardwareProfileResponse) string { return v.HardwareType }).(pulumi.StringOutput)
}

type HardwareProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (HardwareProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HardwareProfileResponse)(nil)).Elem()
}

func (o HardwareProfileResponsePtrOutput) ToHardwareProfileResponsePtrOutput() HardwareProfileResponsePtrOutput {
	return o
}

func (o HardwareProfileResponsePtrOutput) ToHardwareProfileResponsePtrOutputWithContext(ctx context.Context) HardwareProfileResponsePtrOutput {
	return o
}

func (o HardwareProfileResponsePtrOutput) Elem() HardwareProfileResponseOutput {
	return o.ApplyT(func(v *HardwareProfileResponse) HardwareProfileResponse { return *v }).(HardwareProfileResponseOutput)
}

// Specifies the HANA instance SKU.
func (o HardwareProfileResponsePtrOutput) HanaInstanceSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HardwareProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.HanaInstanceSize
	}).(pulumi.StringPtrOutput)
}

// Name of the hardware type (vendor and/or their product name)
func (o HardwareProfileResponsePtrOutput) HardwareType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HardwareProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.HardwareType
	}).(pulumi.StringPtrOutput)
}

// Specifies the IP address of the network interface.
type IpAddress struct {
	// Specifies the IP address of the network interface.
	IpAddress *string `pulumi:"ipAddress"`
}

// IpAddressInput is an input type that accepts IpAddressArgs and IpAddressOutput values.
// You can construct a concrete instance of `IpAddressInput` via:
//
//          IpAddressArgs{...}
type IpAddressInput interface {
	pulumi.Input

	ToIpAddressOutput() IpAddressOutput
	ToIpAddressOutputWithContext(context.Context) IpAddressOutput
}

// Specifies the IP address of the network interface.
type IpAddressArgs struct {
	// Specifies the IP address of the network interface.
	IpAddress pulumi.StringPtrInput `pulumi:"ipAddress"`
}

func (IpAddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IpAddress)(nil)).Elem()
}

func (i IpAddressArgs) ToIpAddressOutput() IpAddressOutput {
	return i.ToIpAddressOutputWithContext(context.Background())
}

func (i IpAddressArgs) ToIpAddressOutputWithContext(ctx context.Context) IpAddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpAddressOutput)
}

// IpAddressArrayInput is an input type that accepts IpAddressArray and IpAddressArrayOutput values.
// You can construct a concrete instance of `IpAddressArrayInput` via:
//
//          IpAddressArray{ IpAddressArgs{...} }
type IpAddressArrayInput interface {
	pulumi.Input

	ToIpAddressArrayOutput() IpAddressArrayOutput
	ToIpAddressArrayOutputWithContext(context.Context) IpAddressArrayOutput
}

type IpAddressArray []IpAddressInput

func (IpAddressArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpAddress)(nil)).Elem()
}

func (i IpAddressArray) ToIpAddressArrayOutput() IpAddressArrayOutput {
	return i.ToIpAddressArrayOutputWithContext(context.Background())
}

func (i IpAddressArray) ToIpAddressArrayOutputWithContext(ctx context.Context) IpAddressArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpAddressArrayOutput)
}

// Specifies the IP address of the network interface.
type IpAddressResponse struct {
	// Specifies the IP address of the network interface.
	IpAddress *string `pulumi:"ipAddress"`
}

// Specifies the IP address of the network interface.
type IpAddressResponseOutput struct{ *pulumi.OutputState }

func (IpAddressResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpAddressResponse)(nil)).Elem()
}

func (o IpAddressResponseOutput) ToIpAddressResponseOutput() IpAddressResponseOutput {
	return o
}

func (o IpAddressResponseOutput) ToIpAddressResponseOutputWithContext(ctx context.Context) IpAddressResponseOutput {
	return o
}

// Specifies the IP address of the network interface.
func (o IpAddressResponseOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpAddressResponse) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

type IpAddressResponseArrayOutput struct{ *pulumi.OutputState }

func (IpAddressResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpAddressResponse)(nil)).Elem()
}

func (o IpAddressResponseArrayOutput) ToIpAddressResponseArrayOutput() IpAddressResponseArrayOutput {
	return o
}

func (o IpAddressResponseArrayOutput) ToIpAddressResponseArrayOutputWithContext(ctx context.Context) IpAddressResponseArrayOutput {
	return o
}

func (o IpAddressResponseArrayOutput) Index(i pulumi.IntInput) IpAddressResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IpAddressResponse {
		return vs[0].([]IpAddressResponse)[vs[1].(int)]
	}).(IpAddressResponseOutput)
}

// Specifies the network settings for the HANA instance disks.
type NetworkProfile struct {
	// Specifies the network interfaces for the HANA instance.
	NetworkInterfaces []IpAddress `pulumi:"networkInterfaces"`
}

// NetworkProfileInput is an input type that accepts NetworkProfileArgs and NetworkProfileOutput values.
// You can construct a concrete instance of `NetworkProfileInput` via:
//
//          NetworkProfileArgs{...}
type NetworkProfileInput interface {
	pulumi.Input

	ToNetworkProfileOutput() NetworkProfileOutput
	ToNetworkProfileOutputWithContext(context.Context) NetworkProfileOutput
}

// Specifies the network settings for the HANA instance disks.
type NetworkProfileArgs struct {
	// Specifies the network interfaces for the HANA instance.
	NetworkInterfaces IpAddressArrayInput `pulumi:"networkInterfaces"`
}

func (NetworkProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkProfile)(nil)).Elem()
}

func (i NetworkProfileArgs) ToNetworkProfileOutput() NetworkProfileOutput {
	return i.ToNetworkProfileOutputWithContext(context.Background())
}

func (i NetworkProfileArgs) ToNetworkProfileOutputWithContext(ctx context.Context) NetworkProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkProfileOutput)
}

func (i NetworkProfileArgs) ToNetworkProfilePtrOutput() NetworkProfilePtrOutput {
	return i.ToNetworkProfilePtrOutputWithContext(context.Background())
}

func (i NetworkProfileArgs) ToNetworkProfilePtrOutputWithContext(ctx context.Context) NetworkProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkProfileOutput).ToNetworkProfilePtrOutputWithContext(ctx)
}

// NetworkProfilePtrInput is an input type that accepts NetworkProfileArgs, NetworkProfilePtr and NetworkProfilePtrOutput values.
// You can construct a concrete instance of `NetworkProfilePtrInput` via:
//
//          NetworkProfileArgs{...}
//
//  or:
//
//          nil
type NetworkProfilePtrInput interface {
	pulumi.Input

	ToNetworkProfilePtrOutput() NetworkProfilePtrOutput
	ToNetworkProfilePtrOutputWithContext(context.Context) NetworkProfilePtrOutput
}

type networkProfilePtrType NetworkProfileArgs

func NetworkProfilePtr(v *NetworkProfileArgs) NetworkProfilePtrInput {
	return (*networkProfilePtrType)(v)
}

func (*networkProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkProfile)(nil)).Elem()
}

func (i *networkProfilePtrType) ToNetworkProfilePtrOutput() NetworkProfilePtrOutput {
	return i.ToNetworkProfilePtrOutputWithContext(context.Background())
}

func (i *networkProfilePtrType) ToNetworkProfilePtrOutputWithContext(ctx context.Context) NetworkProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkProfilePtrOutput)
}

// Specifies the network settings for the HANA instance disks.
type NetworkProfileResponse struct {
	// Specifies the circuit id for connecting to express route.
	CircuitId string `pulumi:"circuitId"`
	// Specifies the network interfaces for the HANA instance.
	NetworkInterfaces []IpAddressResponse `pulumi:"networkInterfaces"`
}

// Specifies the network settings for the HANA instance disks.
type NetworkProfileResponseOutput struct{ *pulumi.OutputState }

func (NetworkProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkProfileResponse)(nil)).Elem()
}

func (o NetworkProfileResponseOutput) ToNetworkProfileResponseOutput() NetworkProfileResponseOutput {
	return o
}

func (o NetworkProfileResponseOutput) ToNetworkProfileResponseOutputWithContext(ctx context.Context) NetworkProfileResponseOutput {
	return o
}

func (o NetworkProfileResponseOutput) ToNetworkProfileResponsePtrOutput() NetworkProfileResponsePtrOutput {
	return o.ToNetworkProfileResponsePtrOutputWithContext(context.Background())
}

func (o NetworkProfileResponseOutput) ToNetworkProfileResponsePtrOutputWithContext(ctx context.Context) NetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v NetworkProfileResponse) *NetworkProfileResponse {
		return &v
	}).(NetworkProfileResponsePtrOutput)
}

// Specifies the circuit id for connecting to express route.
func (o NetworkProfileResponseOutput) CircuitId() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkProfileResponse) string { return v.CircuitId }).(pulumi.StringOutput)
}

// Specifies the network interfaces for the HANA instance.
func (o NetworkProfileResponseOutput) NetworkInterfaces() IpAddressResponseArrayOutput {
	return o.ApplyT(func(v NetworkProfileResponse) []IpAddressResponse { return v.NetworkInterfaces }).(IpAddressResponseArrayOutput)
}

type NetworkProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (NetworkProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkProfileResponse)(nil)).Elem()
}

func (o NetworkProfileResponsePtrOutput) ToNetworkProfileResponsePtrOutput() NetworkProfileResponsePtrOutput {
	return o
}

func (o NetworkProfileResponsePtrOutput) ToNetworkProfileResponsePtrOutputWithContext(ctx context.Context) NetworkProfileResponsePtrOutput {
	return o
}

func (o NetworkProfileResponsePtrOutput) Elem() NetworkProfileResponseOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) NetworkProfileResponse { return *v }).(NetworkProfileResponseOutput)
}

// Specifies the circuit id for connecting to express route.
func (o NetworkProfileResponsePtrOutput) CircuitId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CircuitId
	}).(pulumi.StringPtrOutput)
}

// Specifies the network interfaces for the HANA instance.
func (o NetworkProfileResponsePtrOutput) NetworkInterfaces() IpAddressResponseArrayOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) []IpAddressResponse {
		if v == nil {
			return nil
		}
		return v.NetworkInterfaces
	}).(IpAddressResponseArrayOutput)
}

// Specifies the operating system settings for the HANA instance.
type OSProfile struct {
	// Specifies the host OS name of the HANA instance.
	ComputerName *string `pulumi:"computerName"`
	// Specifies the SSH public key used to access the operating system.
	SshPublicKey *string `pulumi:"sshPublicKey"`
}

// OSProfileInput is an input type that accepts OSProfileArgs and OSProfileOutput values.
// You can construct a concrete instance of `OSProfileInput` via:
//
//          OSProfileArgs{...}
type OSProfileInput interface {
	pulumi.Input

	ToOSProfileOutput() OSProfileOutput
	ToOSProfileOutputWithContext(context.Context) OSProfileOutput
}

// Specifies the operating system settings for the HANA instance.
type OSProfileArgs struct {
	// Specifies the host OS name of the HANA instance.
	ComputerName pulumi.StringPtrInput `pulumi:"computerName"`
	// Specifies the SSH public key used to access the operating system.
	SshPublicKey pulumi.StringPtrInput `pulumi:"sshPublicKey"`
}

func (OSProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OSProfile)(nil)).Elem()
}

func (i OSProfileArgs) ToOSProfileOutput() OSProfileOutput {
	return i.ToOSProfileOutputWithContext(context.Background())
}

func (i OSProfileArgs) ToOSProfileOutputWithContext(ctx context.Context) OSProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OSProfileOutput)
}

func (i OSProfileArgs) ToOSProfilePtrOutput() OSProfilePtrOutput {
	return i.ToOSProfilePtrOutputWithContext(context.Background())
}

func (i OSProfileArgs) ToOSProfilePtrOutputWithContext(ctx context.Context) OSProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OSProfileOutput).ToOSProfilePtrOutputWithContext(ctx)
}

// OSProfilePtrInput is an input type that accepts OSProfileArgs, OSProfilePtr and OSProfilePtrOutput values.
// You can construct a concrete instance of `OSProfilePtrInput` via:
//
//          OSProfileArgs{...}
//
//  or:
//
//          nil
type OSProfilePtrInput interface {
	pulumi.Input

	ToOSProfilePtrOutput() OSProfilePtrOutput
	ToOSProfilePtrOutputWithContext(context.Context) OSProfilePtrOutput
}

type osprofilePtrType OSProfileArgs

func OSProfilePtr(v *OSProfileArgs) OSProfilePtrInput {
	return (*osprofilePtrType)(v)
}

func (*osprofilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OSProfile)(nil)).Elem()
}

func (i *osprofilePtrType) ToOSProfilePtrOutput() OSProfilePtrOutput {
	return i.ToOSProfilePtrOutputWithContext(context.Background())
}

func (i *osprofilePtrType) ToOSProfilePtrOutputWithContext(ctx context.Context) OSProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OSProfilePtrOutput)
}

// Specifies the operating system settings for the HANA instance.
type OSProfileResponse struct {
	// Specifies the host OS name of the HANA instance.
	ComputerName *string `pulumi:"computerName"`
	// This property allows you to specify the type of the OS.
	OsType string `pulumi:"osType"`
	// Specifies the SSH public key used to access the operating system.
	SshPublicKey *string `pulumi:"sshPublicKey"`
	// Specifies version of operating system.
	Version string `pulumi:"version"`
}

// Specifies the operating system settings for the HANA instance.
type OSProfileResponseOutput struct{ *pulumi.OutputState }

func (OSProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OSProfileResponse)(nil)).Elem()
}

func (o OSProfileResponseOutput) ToOSProfileResponseOutput() OSProfileResponseOutput {
	return o
}

func (o OSProfileResponseOutput) ToOSProfileResponseOutputWithContext(ctx context.Context) OSProfileResponseOutput {
	return o
}

func (o OSProfileResponseOutput) ToOSProfileResponsePtrOutput() OSProfileResponsePtrOutput {
	return o.ToOSProfileResponsePtrOutputWithContext(context.Background())
}

func (o OSProfileResponseOutput) ToOSProfileResponsePtrOutputWithContext(ctx context.Context) OSProfileResponsePtrOutput {
	return o.ApplyT(func(v OSProfileResponse) *OSProfileResponse {
		return &v
	}).(OSProfileResponsePtrOutput)
}

// Specifies the host OS name of the HANA instance.
func (o OSProfileResponseOutput) ComputerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSProfileResponse) *string { return v.ComputerName }).(pulumi.StringPtrOutput)
}

// This property allows you to specify the type of the OS.
func (o OSProfileResponseOutput) OsType() pulumi.StringOutput {
	return o.ApplyT(func(v OSProfileResponse) string { return v.OsType }).(pulumi.StringOutput)
}

// Specifies the SSH public key used to access the operating system.
func (o OSProfileResponseOutput) SshPublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSProfileResponse) *string { return v.SshPublicKey }).(pulumi.StringPtrOutput)
}

// Specifies version of operating system.
func (o OSProfileResponseOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v OSProfileResponse) string { return v.Version }).(pulumi.StringOutput)
}

type OSProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (OSProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OSProfileResponse)(nil)).Elem()
}

func (o OSProfileResponsePtrOutput) ToOSProfileResponsePtrOutput() OSProfileResponsePtrOutput {
	return o
}

func (o OSProfileResponsePtrOutput) ToOSProfileResponsePtrOutputWithContext(ctx context.Context) OSProfileResponsePtrOutput {
	return o
}

func (o OSProfileResponsePtrOutput) Elem() OSProfileResponseOutput {
	return o.ApplyT(func(v *OSProfileResponse) OSProfileResponse { return *v }).(OSProfileResponseOutput)
}

// Specifies the host OS name of the HANA instance.
func (o OSProfileResponsePtrOutput) ComputerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.ComputerName
	}).(pulumi.StringPtrOutput)
}

// This property allows you to specify the type of the OS.
func (o OSProfileResponsePtrOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.OsType
	}).(pulumi.StringPtrOutput)
}

// Specifies the SSH public key used to access the operating system.
func (o OSProfileResponsePtrOutput) SshPublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.SshPublicKey
	}).(pulumi.StringPtrOutput)
}

// Specifies version of operating system.
func (o OSProfileResponsePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Version
	}).(pulumi.StringPtrOutput)
}

// Specifies the storage settings for the HANA instance disks.
type StorageProfile struct {
	// Specifies information about the operating system disk used by the hana instance.
	OsDisks []Disk `pulumi:"osDisks"`
}

// StorageProfileInput is an input type that accepts StorageProfileArgs and StorageProfileOutput values.
// You can construct a concrete instance of `StorageProfileInput` via:
//
//          StorageProfileArgs{...}
type StorageProfileInput interface {
	pulumi.Input

	ToStorageProfileOutput() StorageProfileOutput
	ToStorageProfileOutputWithContext(context.Context) StorageProfileOutput
}

// Specifies the storage settings for the HANA instance disks.
type StorageProfileArgs struct {
	// Specifies information about the operating system disk used by the hana instance.
	OsDisks DiskArrayInput `pulumi:"osDisks"`
}

func (StorageProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageProfile)(nil)).Elem()
}

func (i StorageProfileArgs) ToStorageProfileOutput() StorageProfileOutput {
	return i.ToStorageProfileOutputWithContext(context.Background())
}

func (i StorageProfileArgs) ToStorageProfileOutputWithContext(ctx context.Context) StorageProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageProfileOutput)
}

func (i StorageProfileArgs) ToStorageProfilePtrOutput() StorageProfilePtrOutput {
	return i.ToStorageProfilePtrOutputWithContext(context.Background())
}

func (i StorageProfileArgs) ToStorageProfilePtrOutputWithContext(ctx context.Context) StorageProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageProfileOutput).ToStorageProfilePtrOutputWithContext(ctx)
}

// StorageProfilePtrInput is an input type that accepts StorageProfileArgs, StorageProfilePtr and StorageProfilePtrOutput values.
// You can construct a concrete instance of `StorageProfilePtrInput` via:
//
//          StorageProfileArgs{...}
//
//  or:
//
//          nil
type StorageProfilePtrInput interface {
	pulumi.Input

	ToStorageProfilePtrOutput() StorageProfilePtrOutput
	ToStorageProfilePtrOutputWithContext(context.Context) StorageProfilePtrOutput
}

type storageProfilePtrType StorageProfileArgs

func StorageProfilePtr(v *StorageProfileArgs) StorageProfilePtrInput {
	return (*storageProfilePtrType)(v)
}

func (*storageProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageProfile)(nil)).Elem()
}

func (i *storageProfilePtrType) ToStorageProfilePtrOutput() StorageProfilePtrOutput {
	return i.ToStorageProfilePtrOutputWithContext(context.Background())
}

func (i *storageProfilePtrType) ToStorageProfilePtrOutputWithContext(ctx context.Context) StorageProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageProfilePtrOutput)
}

// Specifies the storage settings for the HANA instance disks.
type StorageProfileResponse struct {
	// IP Address to connect to storage.
	NfsIpAddress string `pulumi:"nfsIpAddress"`
	// Specifies information about the operating system disk used by the hana instance.
	OsDisks []DiskResponse `pulumi:"osDisks"`
}

// Specifies the storage settings for the HANA instance disks.
type StorageProfileResponseOutput struct{ *pulumi.OutputState }

func (StorageProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageProfileResponse)(nil)).Elem()
}

func (o StorageProfileResponseOutput) ToStorageProfileResponseOutput() StorageProfileResponseOutput {
	return o
}

func (o StorageProfileResponseOutput) ToStorageProfileResponseOutputWithContext(ctx context.Context) StorageProfileResponseOutput {
	return o
}

func (o StorageProfileResponseOutput) ToStorageProfileResponsePtrOutput() StorageProfileResponsePtrOutput {
	return o.ToStorageProfileResponsePtrOutputWithContext(context.Background())
}

func (o StorageProfileResponseOutput) ToStorageProfileResponsePtrOutputWithContext(ctx context.Context) StorageProfileResponsePtrOutput {
	return o.ApplyT(func(v StorageProfileResponse) *StorageProfileResponse {
		return &v
	}).(StorageProfileResponsePtrOutput)
}

// IP Address to connect to storage.
func (o StorageProfileResponseOutput) NfsIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v StorageProfileResponse) string { return v.NfsIpAddress }).(pulumi.StringOutput)
}

// Specifies information about the operating system disk used by the hana instance.
func (o StorageProfileResponseOutput) OsDisks() DiskResponseArrayOutput {
	return o.ApplyT(func(v StorageProfileResponse) []DiskResponse { return v.OsDisks }).(DiskResponseArrayOutput)
}

type StorageProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (StorageProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageProfileResponse)(nil)).Elem()
}

func (o StorageProfileResponsePtrOutput) ToStorageProfileResponsePtrOutput() StorageProfileResponsePtrOutput {
	return o
}

func (o StorageProfileResponsePtrOutput) ToStorageProfileResponsePtrOutputWithContext(ctx context.Context) StorageProfileResponsePtrOutput {
	return o
}

func (o StorageProfileResponsePtrOutput) Elem() StorageProfileResponseOutput {
	return o.ApplyT(func(v *StorageProfileResponse) StorageProfileResponse { return *v }).(StorageProfileResponseOutput)
}

// IP Address to connect to storage.
func (o StorageProfileResponsePtrOutput) NfsIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.NfsIpAddress
	}).(pulumi.StringPtrOutput)
}

// Specifies information about the operating system disk used by the hana instance.
func (o StorageProfileResponsePtrOutput) OsDisks() DiskResponseArrayOutput {
	return o.ApplyT(func(v *StorageProfileResponse) []DiskResponse {
		if v == nil {
			return nil
		}
		return v.OsDisks
	}).(DiskResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(DiskOutput{})
	pulumi.RegisterOutputType(DiskArrayOutput{})
	pulumi.RegisterOutputType(DiskResponseOutput{})
	pulumi.RegisterOutputType(DiskResponseArrayOutput{})
	pulumi.RegisterOutputType(HardwareProfileResponseOutput{})
	pulumi.RegisterOutputType(HardwareProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(IpAddressOutput{})
	pulumi.RegisterOutputType(IpAddressArrayOutput{})
	pulumi.RegisterOutputType(IpAddressResponseOutput{})
	pulumi.RegisterOutputType(IpAddressResponseArrayOutput{})
	pulumi.RegisterOutputType(NetworkProfileOutput{})
	pulumi.RegisterOutputType(NetworkProfilePtrOutput{})
	pulumi.RegisterOutputType(NetworkProfileResponseOutput{})
	pulumi.RegisterOutputType(NetworkProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(OSProfileOutput{})
	pulumi.RegisterOutputType(OSProfilePtrOutput{})
	pulumi.RegisterOutputType(OSProfileResponseOutput{})
	pulumi.RegisterOutputType(OSProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(StorageProfileOutput{})
	pulumi.RegisterOutputType(StorageProfilePtrOutput{})
	pulumi.RegisterOutputType(StorageProfileResponseOutput{})
	pulumi.RegisterOutputType(StorageProfileResponsePtrOutput{})
}
