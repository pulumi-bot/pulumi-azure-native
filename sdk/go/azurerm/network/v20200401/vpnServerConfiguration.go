// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200401

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// VpnServerConfiguration Resource.
//
// ## Example Usage
// ### VpnServerConfigurationCreate
//
// ```go
// package main
//
// import (
// 	network "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/network/v20200401"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := network.NewVpnServerConfiguration(ctx, "vpnServerConfiguration", &network.VpnServerConfigurationArgs{
// 			Location: pulumi.String("West US"),
// 			RadiusClientRootCertificates: network.VpnServerConfigRadiusClientRootCertificateArray{
// 				&network.VpnServerConfigRadiusClientRootCertificateArgs{
// 					Name:       pulumi.String("vpnServerConfigRadiusClientRootCert1"),
// 					Thumbprint: pulumi.String("83FFBFC8848B5A5836C94D0112367E16148A286F"),
// 				},
// 			},
// 			RadiusServerRootCertificates: network.VpnServerConfigRadiusServerRootCertificateArray{
// 				&network.VpnServerConfigRadiusServerRootCertificateArgs{
// 					Name:           pulumi.String("vpnServerConfigRadiusServerRootCer1"),
// 					PublicCertData: pulumi.String("MIIC5zCCAc+gAwIBAgIQErQ0Hk4aDJxIA+Q5RagB+jANBgkqhkiG9w0BAQsFADAWMRQwEgYDVQQDDAtQMlNSb290Q2VydDAeFw0xNzEyMTQyMTA3MzhaFw0xODEyMTQyMTI3MzhaMBYxFDASBgNVBAMMC1AyU1Jvb3RDZXJ0MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArP7/NQXmW7cQ/ZR1mv3Y3I29Lt7HTOqzo/1KUOoVH3NItbQIRAQbwKy3UWrOFz4eGNX2GWtNRMdCyWsKeqy9Ltsdfcm1IbKXkl84DFeU/ZacXu4Dl3xX3gV5du4TLZjEowJELyur11Ea2YcjPRQ/FzAF9/hGuboS1HZQEPLx4FdUs9OxCYOtc0MxBCwLfVTTRqarb0Ne+arNYd4kCzIhAke1nOyKAJBda5ZL+VHy3S5S8qGlD46jm8HXugmAkUygS4oIIXOmj/1O9sNAi3LN60zufSzCmP8Rm/iUGX+DHAGGiXxwZOKQLEDaZXKqoHjMPP0XudmSWwOIbyeQVrLhkwIDAQABozEwLzAOBgNVHQ8BAf8EBAMCAgQwHQYDVR0OBBYEFEfeNU2trYxNLF9ONmuJUsT13pKDMA0GCSqGSIb3DQEBCwUAA4IBAQBmM6RJzsGGipxyMhimHKN2xlkejhVsgBoTAhOU0llW9aUSwINJ9zFUGgI8IzUFy1VG776fchHp0LMRmPSIUYk5btEPxbsrPtumPuMH8EQGrS+Rt4pD+78c8H1fEPkq5CmDl/PKu4JoFGv+aFcE+Od0hlILstIF10Qysf++QXDolKfzJa/56bgMeYKFiju73loiRM57ns8ddXpfLl792UVpRkFU62LNns6Y1LKTwapmUF4IvIuAIzd6LZNOQng64LAKXtKnViJ1JQiXwf4CEzhgvAti3/ejpb3U90hsrUcyZi6wBv9bZLcAJRWpz61JNYliM1d1grSwQDKGXNQE4xuM"),
// 				},
// 			},
// 			RadiusServers: network.RadiusServerArray{
// 				&network.RadiusServerArgs{
// 					RadiusServerAddress: pulumi.String("10.0.0.0"),
// 					RadiusServerScore:   pulumi.Int(25),
// 					RadiusServerSecret:  pulumi.String("radiusServerSecret"),
// 				},
// 			},
// 			ResourceGroupName: pulumi.String("rg1"),
// 			Tags: pulumi.StringMap{
// 				"key1": pulumi.String("value1"),
// 			},
// 			VpnClientIpsecPolicies: network.IpsecPolicyArray{
// 				&network.IpsecPolicyArgs{
// 					DhGroup:             pulumi.String("DHGroup14"),
// 					IkeEncryption:       pulumi.String("AES256"),
// 					IkeIntegrity:        pulumi.String("SHA384"),
// 					IpsecEncryption:     pulumi.String("AES256"),
// 					IpsecIntegrity:      pulumi.String("SHA256"),
// 					PfsGroup:            pulumi.String("PFS14"),
// 					SaDataSizeKilobytes: pulumi.Int(429497),
// 					SaLifeTimeSeconds:   pulumi.Int(86472),
// 				},
// 			},
// 			VpnClientRevokedCertificates: network.VpnServerConfigVpnClientRevokedCertificateArray{
// 				&network.VpnServerConfigVpnClientRevokedCertificateArgs{
// 					Name:       pulumi.String("vpnServerConfigVpnClientRevokedCert1"),
// 					Thumbprint: pulumi.String("83FFBFC8848B5A5836C94D0112367E16148A286F"),
// 				},
// 			},
// 			VpnClientRootCertificates: network.VpnServerConfigVpnClientRootCertificateArray{
// 				&network.VpnServerConfigVpnClientRootCertificateArgs{
// 					Name:           pulumi.String("vpnServerConfigVpnClientRootCert1"),
// 					PublicCertData: pulumi.String("MIIC5zCCAc+gAwIBAgIQErQ0Hk4aDJxIA+Q5RagB+jANBgkqhkiG9w0BAQsFADAWMRQwEgYDVQQDDAtQMlNSb290Q2VydDAeFw0xNzEyMTQyMTA3MzhaFw0xODEyMTQyMTI3MzhaMBYxFDASBgNVBAMMC1AyU1Jvb3RDZXJ0MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArP7/NQXmW7cQ/ZR1mv3Y3I29Lt7HTOqzo/1KUOoVH3NItbQIRAQbwKy3UWrOFz4eGNX2GWtNRMdCyWsKeqy9Ltsdfcm1IbKXkl84DFeU/ZacXu4Dl3xX3gV5du4TLZjEowJELyur11Ea2YcjPRQ/FzAF9/hGuboS1HZQEPLx4FdUs9OxCYOtc0MxBCwLfVTTRqarb0Ne+arNYd4kCzIhAke1nOyKAJBda5ZL+VHy3S5S8qGlD46jm8HXugmAkUygS4oIIXOmj/1O9sNAi3LN60zufSzCmP8Rm/iUGX+DHAGGiXxwZOKQLEDaZXKqoHjMPP0XudmSWwOIbyeQVrLhkwIDAQABozEwLzAOBgNVHQ8BAf8EBAMCAgQwHQYDVR0OBBYEFEfeNU2trYxNLF9ONmuJUsT13pKDMA0GCSqGSIb3DQEBCwUAA4IBAQBmM6RJzsGGipxyMhimHKN2xlkejhVsgBoTAhOU0llW9aUSwINJ9zFUGgI8IzUFy1VG776fchHp0LMRmPSIUYk5btEPxbsrPtumPuMH8EQGrS+Rt4pD+78c8H1fEPkq5CmDl/PKu4JoFGv+aFcE+Od0hlILstIF10Qysf++QXDolKfzJa/56bgMeYKFiju73loiRM57ns8ddXpfLl792UVpRkFU62LNns6Y1LKTwapmUF4IvIuAIzd6LZNOQng64LAKXtKnViJ1JQiXwf4CEzhgvAti3/ejpb3U90hsrUcyZi6wBv9bZLcAJRWpz61JNYliM1d1grSwQDKGXNQE4xuN"),
// 				},
// 			},
// 			VpnProtocols: pulumi.StringArray{
// 				pulumi.String("IkeV2"),
// 			},
// 			VpnServerConfigurationName: pulumi.String("vpnServerConfiguration1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
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
			Type: pulumi.String("azurerm:network/latest:VpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190801:VpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190901:VpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azurerm:network/v20191101:VpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azurerm:network/v20191201:VpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200301:VpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200501:VpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200601:VpnServerConfiguration"),
		},
	})
	opts = append(opts, aliases)
	var resource VpnServerConfiguration
	err := ctx.RegisterResource("azurerm:network/v20200401:VpnServerConfiguration", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azurerm:network/v20200401:VpnServerConfiguration", name, id, state, &resource, opts...)
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
