// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200701

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// VpnServerConfiguration Resource.
type VpnServerConfiguration struct {
	pulumi.CustomResourceState

	// The set of aad vpn authentication parameters.
	AadAuthenticationParameters AadAuthenticationParametersResponsePtrOutput `pulumi:"aadAuthenticationParameters"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// List of references to P2SVpnGateways.
	P2SVpnGateways P2SVpnGatewayResponseArrayOutput `pulumi:"p2SVpnGateways"`
	// The provisioning state of the VpnServerConfiguration resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Radius client root certificate of VpnServerConfiguration.
	RadiusClientRootCertificates VpnServerConfigRadiusClientRootCertificateResponseArrayOutput `pulumi:"radiusClientRootCertificates"`
	// The radius server address property of the VpnServerConfiguration resource for point to site client connection.
	RadiusServerAddress pulumi.StringPtrOutput `pulumi:"radiusServerAddress"`
	// Radius Server root certificate of VpnServerConfiguration.
	RadiusServerRootCertificates VpnServerConfigRadiusServerRootCertificateResponseArrayOutput `pulumi:"radiusServerRootCertificates"`
	// The radius secret property of the VpnServerConfiguration resource for point to site client connection.
	RadiusServerSecret pulumi.StringPtrOutput `pulumi:"radiusServerSecret"`
	// Multiple Radius Server configuration for VpnServerConfiguration.
	RadiusServers RadiusServerResponseArrayOutput `pulumi:"radiusServers"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// VPN authentication types for the VpnServerConfiguration.
	VpnAuthenticationTypes pulumi.StringArrayOutput `pulumi:"vpnAuthenticationTypes"`
	// VpnClientIpsecPolicies for VpnServerConfiguration.
	VpnClientIpsecPolicies IpsecPolicyResponseArrayOutput `pulumi:"vpnClientIpsecPolicies"`
	// VPN client revoked certificate of VpnServerConfiguration.
	VpnClientRevokedCertificates VpnServerConfigVpnClientRevokedCertificateResponseArrayOutput `pulumi:"vpnClientRevokedCertificates"`
	// VPN client root certificate of VpnServerConfiguration.
	VpnClientRootCertificates VpnServerConfigVpnClientRootCertificateResponseArrayOutput `pulumi:"vpnClientRootCertificates"`
	// VPN protocols for the VpnServerConfiguration.
	VpnProtocols pulumi.StringArrayOutput `pulumi:"vpnProtocols"`
}

// NewVpnServerConfiguration registers a new resource with the given unique name, arguments, and options.
func NewVpnServerConfiguration(ctx *pulumi.Context,
	name string, args *VpnServerConfigurationArgs, opts ...pulumi.ResourceOption) (*VpnServerConfiguration, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.VpnServerConfigurationName == nil {
		return nil, errors.New("missing required argument 'VpnServerConfigurationName'")
	}
	if args == nil {
		args = &VpnServerConfigurationArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:network/latest:VpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190801:VpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190901:VpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20191101:VpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20191201:VpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200301:VpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200401:VpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200501:VpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200601:VpnServerConfiguration"),
		},
	})
	opts = append(opts, aliases)
	var resource VpnServerConfiguration
	err := ctx.RegisterResource("azure-nextgen:network/v20200701:VpnServerConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpnServerConfiguration gets an existing VpnServerConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpnServerConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpnServerConfigurationState, opts ...pulumi.ResourceOption) (*VpnServerConfiguration, error) {
	var resource VpnServerConfiguration
	err := ctx.ReadResource("azure-nextgen:network/v20200701:VpnServerConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpnServerConfiguration resources.
type vpnServerConfigurationState struct {
	// The set of aad vpn authentication parameters.
	AadAuthenticationParameters *AadAuthenticationParametersResponse `pulumi:"aadAuthenticationParameters"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name *string `pulumi:"name"`
	// List of references to P2SVpnGateways.
	P2SVpnGateways []P2SVpnGatewayResponse `pulumi:"p2SVpnGateways"`
	// The provisioning state of the VpnServerConfiguration resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Radius client root certificate of VpnServerConfiguration.
	RadiusClientRootCertificates []VpnServerConfigRadiusClientRootCertificateResponse `pulumi:"radiusClientRootCertificates"`
	// The radius server address property of the VpnServerConfiguration resource for point to site client connection.
	RadiusServerAddress *string `pulumi:"radiusServerAddress"`
	// Radius Server root certificate of VpnServerConfiguration.
	RadiusServerRootCertificates []VpnServerConfigRadiusServerRootCertificateResponse `pulumi:"radiusServerRootCertificates"`
	// The radius secret property of the VpnServerConfiguration resource for point to site client connection.
	RadiusServerSecret *string `pulumi:"radiusServerSecret"`
	// Multiple Radius Server configuration for VpnServerConfiguration.
	RadiusServers []RadiusServerResponse `pulumi:"radiusServers"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
	// VPN authentication types for the VpnServerConfiguration.
	VpnAuthenticationTypes []string `pulumi:"vpnAuthenticationTypes"`
	// VpnClientIpsecPolicies for VpnServerConfiguration.
	VpnClientIpsecPolicies []IpsecPolicyResponse `pulumi:"vpnClientIpsecPolicies"`
	// VPN client revoked certificate of VpnServerConfiguration.
	VpnClientRevokedCertificates []VpnServerConfigVpnClientRevokedCertificateResponse `pulumi:"vpnClientRevokedCertificates"`
	// VPN client root certificate of VpnServerConfiguration.
	VpnClientRootCertificates []VpnServerConfigVpnClientRootCertificateResponse `pulumi:"vpnClientRootCertificates"`
	// VPN protocols for the VpnServerConfiguration.
	VpnProtocols []string `pulumi:"vpnProtocols"`
}

type VpnServerConfigurationState struct {
	// The set of aad vpn authentication parameters.
	AadAuthenticationParameters AadAuthenticationParametersResponsePtrInput
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// List of references to P2SVpnGateways.
	P2SVpnGateways P2SVpnGatewayResponseArrayInput
	// The provisioning state of the VpnServerConfiguration resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState pulumi.StringPtrInput
	// Radius client root certificate of VpnServerConfiguration.
	RadiusClientRootCertificates VpnServerConfigRadiusClientRootCertificateResponseArrayInput
	// The radius server address property of the VpnServerConfiguration resource for point to site client connection.
	RadiusServerAddress pulumi.StringPtrInput
	// Radius Server root certificate of VpnServerConfiguration.
	RadiusServerRootCertificates VpnServerConfigRadiusServerRootCertificateResponseArrayInput
	// The radius secret property of the VpnServerConfiguration resource for point to site client connection.
	RadiusServerSecret pulumi.StringPtrInput
	// Multiple Radius Server configuration for VpnServerConfiguration.
	RadiusServers RadiusServerResponseArrayInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
	// VPN authentication types for the VpnServerConfiguration.
	VpnAuthenticationTypes pulumi.StringArrayInput
	// VpnClientIpsecPolicies for VpnServerConfiguration.
	VpnClientIpsecPolicies IpsecPolicyResponseArrayInput
	// VPN client revoked certificate of VpnServerConfiguration.
	VpnClientRevokedCertificates VpnServerConfigVpnClientRevokedCertificateResponseArrayInput
	// VPN client root certificate of VpnServerConfiguration.
	VpnClientRootCertificates VpnServerConfigVpnClientRootCertificateResponseArrayInput
	// VPN protocols for the VpnServerConfiguration.
	VpnProtocols pulumi.StringArrayInput
}

func (VpnServerConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnServerConfigurationState)(nil)).Elem()
}

type vpnServerConfigurationArgs struct {
	// The set of aad vpn authentication parameters.
	AadAuthenticationParameters *AadAuthenticationParameters `pulumi:"aadAuthenticationParameters"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the VpnServerConfiguration that is unique within a resource group.
	Name *string `pulumi:"name"`
	// Radius client root certificate of VpnServerConfiguration.
	RadiusClientRootCertificates []VpnServerConfigRadiusClientRootCertificate `pulumi:"radiusClientRootCertificates"`
	// The radius server address property of the VpnServerConfiguration resource for point to site client connection.
	RadiusServerAddress *string `pulumi:"radiusServerAddress"`
	// Radius Server root certificate of VpnServerConfiguration.
	RadiusServerRootCertificates []VpnServerConfigRadiusServerRootCertificate `pulumi:"radiusServerRootCertificates"`
	// The radius secret property of the VpnServerConfiguration resource for point to site client connection.
	RadiusServerSecret *string `pulumi:"radiusServerSecret"`
	// Multiple Radius Server configuration for VpnServerConfiguration.
	RadiusServers []RadiusServer `pulumi:"radiusServers"`
	// The resource group name of the VpnServerConfiguration.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// VPN authentication types for the VpnServerConfiguration.
	VpnAuthenticationTypes []string `pulumi:"vpnAuthenticationTypes"`
	// VpnClientIpsecPolicies for VpnServerConfiguration.
	VpnClientIpsecPolicies []IpsecPolicy `pulumi:"vpnClientIpsecPolicies"`
	// VPN client revoked certificate of VpnServerConfiguration.
	VpnClientRevokedCertificates []VpnServerConfigVpnClientRevokedCertificate `pulumi:"vpnClientRevokedCertificates"`
	// VPN client root certificate of VpnServerConfiguration.
	VpnClientRootCertificates []VpnServerConfigVpnClientRootCertificate `pulumi:"vpnClientRootCertificates"`
	// VPN protocols for the VpnServerConfiguration.
	VpnProtocols []string `pulumi:"vpnProtocols"`
	// The name of the VpnServerConfiguration being created or updated.
	VpnServerConfigurationName string `pulumi:"vpnServerConfigurationName"`
}

// The set of arguments for constructing a VpnServerConfiguration resource.
type VpnServerConfigurationArgs struct {
	// The set of aad vpn authentication parameters.
	AadAuthenticationParameters AadAuthenticationParametersPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the VpnServerConfiguration that is unique within a resource group.
	Name pulumi.StringPtrInput
	// Radius client root certificate of VpnServerConfiguration.
	RadiusClientRootCertificates VpnServerConfigRadiusClientRootCertificateArrayInput
	// The radius server address property of the VpnServerConfiguration resource for point to site client connection.
	RadiusServerAddress pulumi.StringPtrInput
	// Radius Server root certificate of VpnServerConfiguration.
	RadiusServerRootCertificates VpnServerConfigRadiusServerRootCertificateArrayInput
	// The radius secret property of the VpnServerConfiguration resource for point to site client connection.
	RadiusServerSecret pulumi.StringPtrInput
	// Multiple Radius Server configuration for VpnServerConfiguration.
	RadiusServers RadiusServerArrayInput
	// The resource group name of the VpnServerConfiguration.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// VPN authentication types for the VpnServerConfiguration.
	VpnAuthenticationTypes pulumi.StringArrayInput
	// VpnClientIpsecPolicies for VpnServerConfiguration.
	VpnClientIpsecPolicies IpsecPolicyArrayInput
	// VPN client revoked certificate of VpnServerConfiguration.
	VpnClientRevokedCertificates VpnServerConfigVpnClientRevokedCertificateArrayInput
	// VPN client root certificate of VpnServerConfiguration.
	VpnClientRootCertificates VpnServerConfigVpnClientRootCertificateArrayInput
	// VPN protocols for the VpnServerConfiguration.
	VpnProtocols pulumi.StringArrayInput
	// The name of the VpnServerConfiguration being created or updated.
	VpnServerConfigurationName pulumi.StringInput
}

func (VpnServerConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnServerConfigurationArgs)(nil)).Elem()
}

type VpnServerConfigurationInput interface {
	pulumi.Input

	ToVpnServerConfigurationOutput() VpnServerConfigurationOutput
	ToVpnServerConfigurationOutputWithContext(ctx context.Context) VpnServerConfigurationOutput
}

func (VpnServerConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnServerConfiguration)(nil)).Elem()
}

func (i VpnServerConfiguration) ToVpnServerConfigurationOutput() VpnServerConfigurationOutput {
	return i.ToVpnServerConfigurationOutputWithContext(context.Background())
}

func (i VpnServerConfiguration) ToVpnServerConfigurationOutputWithContext(ctx context.Context) VpnServerConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnServerConfigurationOutput)
}

type VpnServerConfigurationOutput struct {
	*pulumi.OutputState
}

func (VpnServerConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnServerConfigurationOutput)(nil)).Elem()
}

func (o VpnServerConfigurationOutput) ToVpnServerConfigurationOutput() VpnServerConfigurationOutput {
	return o
}

func (o VpnServerConfigurationOutput) ToVpnServerConfigurationOutputWithContext(ctx context.Context) VpnServerConfigurationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(VpnServerConfigurationOutput{})
}
