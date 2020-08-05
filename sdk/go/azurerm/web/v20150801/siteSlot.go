// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150801

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Represents a web app
type SiteSlot struct {
	pulumi.CustomResourceState

	// Kind of resource
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource Name
	Name       pulumi.StringPtrOutput       `pulumi:"name"`
	Properties SiteResponsePropertiesOutput `pulumi:"properties"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewSiteSlot registers a new resource with the given unique name, arguments, and options.
func NewSiteSlot(ctx *pulumi.Context,
	name string, args *SiteSlotArgs, opts ...pulumi.ResourceOption) (*SiteSlot, error) {
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
		args = &SiteSlotArgs{}
	}
	var resource SiteSlot
	err := ctx.RegisterResource("azurerm:web/v20150801:SiteSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSiteSlot gets an existing SiteSlot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSiteSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SiteSlotState, opts ...pulumi.ResourceOption) (*SiteSlot, error) {
	var resource SiteSlot
	err := ctx.ReadResource("azurerm:web/v20150801:SiteSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SiteSlot resources.
type siteSlotState struct {
	// Kind of resource
	Kind *string `pulumi:"kind"`
	// Resource Location
	Location *string `pulumi:"location"`
	// Resource Name
	Name       *string                 `pulumi:"name"`
	Properties *SiteResponseProperties `pulumi:"properties"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}

type SiteSlotState struct {
	// Kind of resource
	Kind pulumi.StringPtrInput
	// Resource Location
	Location pulumi.StringPtrInput
	// Resource Name
	Name       pulumi.StringPtrInput
	Properties SiteResponsePropertiesPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (SiteSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*siteSlotState)(nil)).Elem()
}

type siteSlotArgs struct {
	// Specifies if the client affinity is enabled when load balancing http request for multiple instances of the web app
	ClientAffinityEnabled *bool `pulumi:"clientAffinityEnabled"`
	// Specifies if the client certificate is enabled for the web app
	ClientCertEnabled *bool `pulumi:"clientCertEnabled"`
	// This is only valid for web app creation. If specified, web app is cloned from
	//             a source web app
	CloningInfo *CloningInfo `pulumi:"cloningInfo"`
	// Size of a function container
	ContainerSize *int `pulumi:"containerSize"`
	// True if the site is enabled; otherwise, false. Setting this  value to false disables the site (takes the site off line).
	Enabled *bool `pulumi:"enabled"`
	// If true, web app hostname is force registered with DNS
	ForceDnsRegistration *string `pulumi:"forceDnsRegistration"`
	// Name of gateway app associated with web app
	GatewaySiteName *string `pulumi:"gatewaySiteName"`
	// Hostname SSL states are  used to manage the SSL bindings for site's hostnames.
	HostNameSslStates []HostNameSslState `pulumi:"hostNameSslStates"`
	// Specifies if the public hostnames are disabled the web app.
	//             If set to true the app is only accessible via API Management process
	HostNamesDisabled *bool `pulumi:"hostNamesDisabled"`
	// Specification for the hosting environment (App Service Environment) to use for the web app
	HostingEnvironmentProfile *HostingEnvironmentProfile `pulumi:"hostingEnvironmentProfile"`
	// Resource Id
	Id *string `pulumi:"id"`
	// Kind of resource
	Kind *string `pulumi:"kind"`
	// Resource Location
	Location string `pulumi:"location"`
	// Maximum number of workers
	//             This only applies to function container
	MaxNumberOfWorkers *int    `pulumi:"maxNumberOfWorkers"`
	MicroService       *string `pulumi:"microService"`
	// Name of web app slot. If not specified then will default to production slot.
	Name string `pulumi:"name"`
	// Name of the resource group
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// If set indicates whether to stop SCM (KUDU) site when the web app is stopped. Default is false.
	ScmSiteAlsoStopped *bool   `pulumi:"scmSiteAlsoStopped"`
	ServerFarmId       *string `pulumi:"serverFarmId"`
	// Configuration of web app
	SiteConfig *SiteConfig `pulumi:"siteConfig"`
	// If true, custom (non *.azurewebsites.net) domains associated with web app are not verified.
	SkipCustomDomainVerification *string `pulumi:"skipCustomDomainVerification"`
	// If true web app hostname is not registered with DNS on creation. This parameter is
	//             only used for app creation
	SkipDnsRegistration *string `pulumi:"skipDnsRegistration"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Time to live in seconds for web app's default domain name
	TtlInSeconds *string `pulumi:"ttlInSeconds"`
	// Resource type
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a SiteSlot resource.
type SiteSlotArgs struct {
	// Specifies if the client affinity is enabled when load balancing http request for multiple instances of the web app
	ClientAffinityEnabled pulumi.BoolPtrInput
	// Specifies if the client certificate is enabled for the web app
	ClientCertEnabled pulumi.BoolPtrInput
	// This is only valid for web app creation. If specified, web app is cloned from
	//             a source web app
	CloningInfo CloningInfoPtrInput
	// Size of a function container
	ContainerSize pulumi.IntPtrInput
	// True if the site is enabled; otherwise, false. Setting this  value to false disables the site (takes the site off line).
	Enabled pulumi.BoolPtrInput
	// If true, web app hostname is force registered with DNS
	ForceDnsRegistration pulumi.StringPtrInput
	// Name of gateway app associated with web app
	GatewaySiteName pulumi.StringPtrInput
	// Hostname SSL states are  used to manage the SSL bindings for site's hostnames.
	HostNameSslStates HostNameSslStateArrayInput
	// Specifies if the public hostnames are disabled the web app.
	//             If set to true the app is only accessible via API Management process
	HostNamesDisabled pulumi.BoolPtrInput
	// Specification for the hosting environment (App Service Environment) to use for the web app
	HostingEnvironmentProfile HostingEnvironmentProfilePtrInput
	// Resource Id
	Id pulumi.StringPtrInput
	// Kind of resource
	Kind pulumi.StringPtrInput
	// Resource Location
	Location pulumi.StringInput
	// Maximum number of workers
	//             This only applies to function container
	MaxNumberOfWorkers pulumi.IntPtrInput
	MicroService       pulumi.StringPtrInput
	// Name of web app slot. If not specified then will default to production slot.
	Name pulumi.StringInput
	// Name of the resource group
	ResourceGroupName pulumi.StringInput
	// If set indicates whether to stop SCM (KUDU) site when the web app is stopped. Default is false.
	ScmSiteAlsoStopped pulumi.BoolPtrInput
	ServerFarmId       pulumi.StringPtrInput
	// Configuration of web app
	SiteConfig SiteConfigPtrInput
	// If true, custom (non *.azurewebsites.net) domains associated with web app are not verified.
	SkipCustomDomainVerification pulumi.StringPtrInput
	// If true web app hostname is not registered with DNS on creation. This parameter is
	//             only used for app creation
	SkipDnsRegistration pulumi.StringPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Time to live in seconds for web app's default domain name
	TtlInSeconds pulumi.StringPtrInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (SiteSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*siteSlotArgs)(nil)).Elem()
}
