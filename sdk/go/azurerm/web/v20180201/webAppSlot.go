// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180201

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A web app, a mobile app backend, or an API app.
type WebAppSlot struct {
	pulumi.CustomResourceState

	// Managed service identity.
	Identity ManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Site resource specific properties
	Properties SiteResponsePropertiesOutput `pulumi:"properties"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWebAppSlot registers a new resource with the given unique name, arguments, and options.
func NewWebAppSlot(ctx *pulumi.Context,
	name string, args *WebAppSlotArgs, opts ...pulumi.ResourceOption) (*WebAppSlot, error) {
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
		args = &WebAppSlotArgs{}
	}
	var resource WebAppSlot
	err := ctx.RegisterResource("azurerm:web/v20180201:WebAppSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebAppSlot gets an existing WebAppSlot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAppSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppSlotState, opts ...pulumi.ResourceOption) (*WebAppSlot, error) {
	var resource WebAppSlot
	err := ctx.ReadResource("azurerm:web/v20180201:WebAppSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAppSlot resources.
type webAppSlotState struct {
	// Managed service identity.
	Identity *ManagedServiceIdentityResponse `pulumi:"identity"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Location.
	Location *string `pulumi:"location"`
	// Resource Name.
	Name *string `pulumi:"name"`
	// Site resource specific properties
	Properties *SiteResponseProperties `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type WebAppSlotState struct {
	// Managed service identity.
	Identity ManagedServiceIdentityResponsePtrInput
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Resource Location.
	Location pulumi.StringPtrInput
	// Resource Name.
	Name pulumi.StringPtrInput
	// Site resource specific properties
	Properties SiteResponsePropertiesPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (WebAppSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppSlotState)(nil)).Elem()
}

type webAppSlotArgs struct {
	// <code>true</code> to enable client affinity; <code>false</code> to stop sending session affinity cookies, which route client requests in the same session to the same instance. Default is <code>true</code>.
	ClientAffinityEnabled *bool `pulumi:"clientAffinityEnabled"`
	// <code>true</code> to enable client certificate authentication (TLS mutual authentication); otherwise, <code>false</code>. Default is <code>false</code>.
	ClientCertEnabled *bool `pulumi:"clientCertEnabled"`
	// client certificate authentication comma-separated exclusion paths
	ClientCertExclusionPaths *string `pulumi:"clientCertExclusionPaths"`
	// If specified during app creation, the app is cloned from a source app.
	CloningInfo *CloningInfo `pulumi:"cloningInfo"`
	// Size of the function container.
	ContainerSize *int `pulumi:"containerSize"`
	// Maximum allowed daily memory-time quota (applicable on dynamic apps only).
	DailyMemoryTimeQuota *int `pulumi:"dailyMemoryTimeQuota"`
	// <code>true</code> if the app is enabled; otherwise, <code>false</code>. Setting this value to false disables the app (takes the app offline).
	Enabled *bool `pulumi:"enabled"`
	// GeoDistributions for this site
	GeoDistributions []GeoDistribution `pulumi:"geoDistributions"`
	// Hostname SSL states are used to manage the SSL bindings for app's hostnames.
	HostNameSslStates []HostNameSslState `pulumi:"hostNameSslStates"`
	// <code>true</code> to disable the public hostnames of the app; otherwise, <code>false</code>.
	//  If <code>true</code>, the app is only accessible via API management process.
	HostNamesDisabled *bool `pulumi:"hostNamesDisabled"`
	// App Service Environment to use for the app.
	HostingEnvironmentProfile *HostingEnvironmentProfile `pulumi:"hostingEnvironmentProfile"`
	// HttpsOnly: configures a web site to accept only https requests. Issues redirect for
	// http requests
	HttpsOnly *bool `pulumi:"httpsOnly"`
	// Hyper-V sandbox.
	HyperV *bool `pulumi:"hyperV"`
	// Managed service identity.
	Identity *ManagedServiceIdentity `pulumi:"identity"`
	// Obsolete: Hyper-V sandbox.
	IsXenon *bool `pulumi:"isXenon"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Location.
	Location string `pulumi:"location"`
	// Name of the deployment slot to create or update. By default, this API attempts to create or modify the production slot.
	Name string `pulumi:"name"`
	// Site redundancy mode
	RedundancyMode *string `pulumi:"redundancyMode"`
	// <code>true</code> if reserved; otherwise, <code>false</code>.
	Reserved *bool `pulumi:"reserved"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// <code>true</code> to stop SCM (KUDU) site when the app is stopped; otherwise, <code>false</code>. The default is <code>false</code>.
	ScmSiteAlsoStopped *bool `pulumi:"scmSiteAlsoStopped"`
	// Resource ID of the associated App Service plan, formatted as: "/subscriptions/{subscriptionID}/resourceGroups/{groupName}/providers/Microsoft.Web/serverfarms/{appServicePlanName}".
	ServerFarmId *string `pulumi:"serverFarmId"`
	// Configuration of the app.
	SiteConfig *SiteConfig `pulumi:"siteConfig"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a WebAppSlot resource.
type WebAppSlotArgs struct {
	// <code>true</code> to enable client affinity; <code>false</code> to stop sending session affinity cookies, which route client requests in the same session to the same instance. Default is <code>true</code>.
	ClientAffinityEnabled pulumi.BoolPtrInput
	// <code>true</code> to enable client certificate authentication (TLS mutual authentication); otherwise, <code>false</code>. Default is <code>false</code>.
	ClientCertEnabled pulumi.BoolPtrInput
	// client certificate authentication comma-separated exclusion paths
	ClientCertExclusionPaths pulumi.StringPtrInput
	// If specified during app creation, the app is cloned from a source app.
	CloningInfo CloningInfoPtrInput
	// Size of the function container.
	ContainerSize pulumi.IntPtrInput
	// Maximum allowed daily memory-time quota (applicable on dynamic apps only).
	DailyMemoryTimeQuota pulumi.IntPtrInput
	// <code>true</code> if the app is enabled; otherwise, <code>false</code>. Setting this value to false disables the app (takes the app offline).
	Enabled pulumi.BoolPtrInput
	// GeoDistributions for this site
	GeoDistributions GeoDistributionArrayInput
	// Hostname SSL states are used to manage the SSL bindings for app's hostnames.
	HostNameSslStates HostNameSslStateArrayInput
	// <code>true</code> to disable the public hostnames of the app; otherwise, <code>false</code>.
	//  If <code>true</code>, the app is only accessible via API management process.
	HostNamesDisabled pulumi.BoolPtrInput
	// App Service Environment to use for the app.
	HostingEnvironmentProfile HostingEnvironmentProfilePtrInput
	// HttpsOnly: configures a web site to accept only https requests. Issues redirect for
	// http requests
	HttpsOnly pulumi.BoolPtrInput
	// Hyper-V sandbox.
	HyperV pulumi.BoolPtrInput
	// Managed service identity.
	Identity ManagedServiceIdentityPtrInput
	// Obsolete: Hyper-V sandbox.
	IsXenon pulumi.BoolPtrInput
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Resource Location.
	Location pulumi.StringInput
	// Name of the deployment slot to create or update. By default, this API attempts to create or modify the production slot.
	Name pulumi.StringInput
	// Site redundancy mode
	RedundancyMode pulumi.StringPtrInput
	// <code>true</code> if reserved; otherwise, <code>false</code>.
	Reserved pulumi.BoolPtrInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
	// <code>true</code> to stop SCM (KUDU) site when the app is stopped; otherwise, <code>false</code>. The default is <code>false</code>.
	ScmSiteAlsoStopped pulumi.BoolPtrInput
	// Resource ID of the associated App Service plan, formatted as: "/subscriptions/{subscriptionID}/resourceGroups/{groupName}/providers/Microsoft.Web/serverfarms/{appServicePlanName}".
	ServerFarmId pulumi.StringPtrInput
	// Configuration of the app.
	SiteConfig SiteConfigPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (WebAppSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppSlotArgs)(nil)).Elem()
}
