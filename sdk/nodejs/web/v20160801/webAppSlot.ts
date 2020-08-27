// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * A web app, a mobile app backend, or an API app.
 */
export class WebAppSlot extends pulumi.CustomResource {
    /**
     * Get an existing WebAppSlot resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): WebAppSlot {
        return new WebAppSlot(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:web/v20160801:WebAppSlot';

    /**
     * Returns true if the given object is an instance of WebAppSlot.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WebAppSlot {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WebAppSlot.__pulumiType;
    }

    /**
     * Management information availability state for the app.
     */
    public /*out*/ readonly availabilityState!: pulumi.Output<SiteAvailabilityState>;
    /**
     * <code>true</code> to enable client affinity; <code>false</code> to stop sending session affinity cookies, which route client requests in the same session to the same instance. Default is <code>true</code>.
     */
    public readonly clientAffinityEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * <code>true</code> to enable client certificate authentication (TLS mutual authentication); otherwise, <code>false</code>. Default is <code>false</code>.
     */
    public readonly clientCertEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * If specified during app creation, the app is cloned from a source app.
     */
    public readonly cloningInfo!: pulumi.Output<outputs.web.v20160801.CloningInfoResponse | undefined>;
    /**
     * Size of the function container.
     */
    public readonly containerSize!: pulumi.Output<number | undefined>;
    /**
     * Maximum allowed daily memory-time quota (applicable on dynamic apps only).
     */
    public readonly dailyMemoryTimeQuota!: pulumi.Output<number | undefined>;
    /**
     * Default hostname of the app. Read-only.
     */
    public /*out*/ readonly defaultHostName!: pulumi.Output<string>;
    /**
     * <code>true</code> if the app is enabled; otherwise, <code>false</code>. Setting this value to false disables the app (takes the app offline).
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * Enabled hostnames for the app.Hostnames need to be assigned (see HostNames) AND enabled. Otherwise,
     * the app is not served on those hostnames.
     */
    public /*out*/ readonly enabledHostNames!: pulumi.Output<string[]>;
    /**
     * Hostname SSL states are used to manage the SSL bindings for app's hostnames.
     */
    public readonly hostNameSslStates!: pulumi.Output<outputs.web.v20160801.HostNameSslStateResponse[] | undefined>;
    /**
     * Hostnames associated with the app.
     */
    public /*out*/ readonly hostNames!: pulumi.Output<string[]>;
    /**
     * <code>true</code> to disable the public hostnames of the app; otherwise, <code>false</code>.
     *  If <code>true</code>, the app is only accessible via API management process.
     */
    public readonly hostNamesDisabled!: pulumi.Output<boolean | undefined>;
    /**
     * App Service Environment to use for the app.
     */
    public readonly hostingEnvironmentProfile!: pulumi.Output<outputs.web.v20160801.HostingEnvironmentProfileResponse | undefined>;
    /**
     * HttpsOnly: configures a web site to accept only https requests. Issues redirect for
     * http requests
     */
    public readonly httpsOnly!: pulumi.Output<boolean | undefined>;
    /**
     * Managed service identity.
     */
    public readonly identity!: pulumi.Output<outputs.web.v20160801.ManagedServiceIdentityResponse | undefined>;
    /**
     * <code>true</code> if the app is a default container; otherwise, <code>false</code>.
     */
    public /*out*/ readonly isDefaultContainer!: pulumi.Output<boolean>;
    /**
     * Kind of resource.
     */
    public readonly kind!: pulumi.Output<string | undefined>;
    /**
     * Last time the app was modified, in UTC. Read-only.
     */
    public /*out*/ readonly lastModifiedTimeUtc!: pulumi.Output<string>;
    /**
     * Resource Location.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Maximum number of workers.
     * This only applies to Functions container.
     */
    public /*out*/ readonly maxNumberOfWorkers!: pulumi.Output<number>;
    /**
     * Resource Name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * List of IP addresses that the app uses for outbound connections (e.g. database access). Includes VIPs from tenants that site can be hosted with current settings. Read-only.
     */
    public /*out*/ readonly outboundIpAddresses!: pulumi.Output<string>;
    /**
     * List of IP addresses that the app uses for outbound connections (e.g. database access). Includes VIPs from all tenants. Read-only.
     */
    public /*out*/ readonly possibleOutboundIpAddresses!: pulumi.Output<string>;
    /**
     * Name of the repository site.
     */
    public /*out*/ readonly repositorySiteName!: pulumi.Output<string>;
    /**
     * <code>true</code> if reserved; otherwise, <code>false</code>.
     */
    public readonly reserved!: pulumi.Output<boolean | undefined>;
    /**
     * Name of the resource group the app belongs to. Read-only.
     */
    public /*out*/ readonly resourceGroup!: pulumi.Output<string>;
    /**
     * <code>true</code> to stop SCM (KUDU) site when the app is stopped; otherwise, <code>false</code>. The default is <code>false</code>.
     */
    public readonly scmSiteAlsoStopped!: pulumi.Output<boolean | undefined>;
    /**
     * Resource ID of the associated App Service plan, formatted as: "/subscriptions/{subscriptionID}/resourceGroups/{groupName}/providers/Microsoft.Web/serverfarms/{appServicePlanName}".
     */
    public readonly serverFarmId!: pulumi.Output<string | undefined>;
    /**
     * Configuration of the app.
     */
    public readonly siteConfig!: pulumi.Output<outputs.web.v20160801.SiteConfigResponse | undefined>;
    /**
     * Status of the last deployment slot swap operation.
     */
    public /*out*/ readonly slotSwapStatus!: pulumi.Output<outputs.web.v20160801.SlotSwapStatusResponse>;
    /**
     * If specified during app creation, the app is created from a previous snapshot.
     */
    public readonly snapshotInfo!: pulumi.Output<outputs.web.v20160801.SnapshotRecoveryRequestResponse | undefined>;
    /**
     * Current state of the app.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * App suspended till in case memory-time quota is exceeded.
     */
    public /*out*/ readonly suspendedTill!: pulumi.Output<string>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Specifies which deployment slot this app will swap into. Read-only.
     */
    public /*out*/ readonly targetSwapSlot!: pulumi.Output<string>;
    /**
     * Azure Traffic Manager hostnames associated with the app. Read-only.
     */
    public /*out*/ readonly trafficManagerHostNames!: pulumi.Output<string[]>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * State indicating whether the app has exceeded its quota usage. Read-only.
     */
    public /*out*/ readonly usageState!: pulumi.Output<UsageState>;

    /**
     * Create a WebAppSlot resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WebAppSlotArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WebAppSlotArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as WebAppSlotArgs | undefined;
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.slot === undefined) {
                throw new Error("Missing required property 'slot'");
            }
            inputs["clientAffinityEnabled"] = args ? args.clientAffinityEnabled : undefined;
            inputs["clientCertEnabled"] = args ? args.clientCertEnabled : undefined;
            inputs["cloningInfo"] = args ? args.cloningInfo : undefined;
            inputs["containerSize"] = args ? args.containerSize : undefined;
            inputs["dailyMemoryTimeQuota"] = args ? args.dailyMemoryTimeQuota : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["forceDnsRegistration"] = args ? args.forceDnsRegistration : undefined;
            inputs["hostNameSslStates"] = args ? args.hostNameSslStates : undefined;
            inputs["hostNamesDisabled"] = args ? args.hostNamesDisabled : undefined;
            inputs["hostingEnvironmentProfile"] = args ? args.hostingEnvironmentProfile : undefined;
            inputs["httpsOnly"] = args ? args.httpsOnly : undefined;
            inputs["identity"] = args ? args.identity : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["reserved"] = args ? args.reserved : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["scmSiteAlsoStopped"] = args ? args.scmSiteAlsoStopped : undefined;
            inputs["serverFarmId"] = args ? args.serverFarmId : undefined;
            inputs["siteConfig"] = args ? args.siteConfig : undefined;
            inputs["skipCustomDomainVerification"] = args ? args.skipCustomDomainVerification : undefined;
            inputs["skipDnsRegistration"] = args ? args.skipDnsRegistration : undefined;
            inputs["slot"] = args ? args.slot : undefined;
            inputs["snapshotInfo"] = args ? args.snapshotInfo : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["ttlInSeconds"] = args ? args.ttlInSeconds : undefined;
            inputs["availabilityState"] = undefined /*out*/;
            inputs["defaultHostName"] = undefined /*out*/;
            inputs["enabledHostNames"] = undefined /*out*/;
            inputs["hostNames"] = undefined /*out*/;
            inputs["isDefaultContainer"] = undefined /*out*/;
            inputs["lastModifiedTimeUtc"] = undefined /*out*/;
            inputs["maxNumberOfWorkers"] = undefined /*out*/;
            inputs["outboundIpAddresses"] = undefined /*out*/;
            inputs["possibleOutboundIpAddresses"] = undefined /*out*/;
            inputs["repositorySiteName"] = undefined /*out*/;
            inputs["resourceGroup"] = undefined /*out*/;
            inputs["slotSwapStatus"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["suspendedTill"] = undefined /*out*/;
            inputs["targetSwapSlot"] = undefined /*out*/;
            inputs["trafficManagerHostNames"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["usageState"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:web/v20150801:WebAppSlot" }, { type: "azurerm:web/v20180201:WebAppSlot" }, { type: "azurerm:web/v20181101:WebAppSlot" }, { type: "azurerm:web/v20190801:WebAppSlot" }, { type: "azurerm:web/v20200601:WebAppSlot" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(WebAppSlot.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a WebAppSlot resource.
 */
export interface WebAppSlotArgs {
    /**
     * <code>true</code> to enable client affinity; <code>false</code> to stop sending session affinity cookies, which route client requests in the same session to the same instance. Default is <code>true</code>.
     */
    readonly clientAffinityEnabled?: pulumi.Input<boolean>;
    /**
     * <code>true</code> to enable client certificate authentication (TLS mutual authentication); otherwise, <code>false</code>. Default is <code>false</code>.
     */
    readonly clientCertEnabled?: pulumi.Input<boolean>;
    /**
     * If specified during app creation, the app is cloned from a source app.
     */
    readonly cloningInfo?: pulumi.Input<inputs.web.v20160801.CloningInfo>;
    /**
     * Size of the function container.
     */
    readonly containerSize?: pulumi.Input<number>;
    /**
     * Maximum allowed daily memory-time quota (applicable on dynamic apps only).
     */
    readonly dailyMemoryTimeQuota?: pulumi.Input<number>;
    /**
     * <code>true</code> if the app is enabled; otherwise, <code>false</code>. Setting this value to false disables the app (takes the app offline).
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * If true, web app hostname is force registered with DNS.
     */
    readonly forceDnsRegistration?: pulumi.Input<boolean>;
    /**
     * Hostname SSL states are used to manage the SSL bindings for app's hostnames.
     */
    readonly hostNameSslStates?: pulumi.Input<pulumi.Input<inputs.web.v20160801.HostNameSslState>[]>;
    /**
     * <code>true</code> to disable the public hostnames of the app; otherwise, <code>false</code>.
     *  If <code>true</code>, the app is only accessible via API management process.
     */
    readonly hostNamesDisabled?: pulumi.Input<boolean>;
    /**
     * App Service Environment to use for the app.
     */
    readonly hostingEnvironmentProfile?: pulumi.Input<inputs.web.v20160801.HostingEnvironmentProfile>;
    /**
     * HttpsOnly: configures a web site to accept only https requests. Issues redirect for
     * http requests
     */
    readonly httpsOnly?: pulumi.Input<boolean>;
    /**
     * Managed service identity.
     */
    readonly identity?: pulumi.Input<inputs.web.v20160801.ManagedServiceIdentity>;
    /**
     * Kind of resource.
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * Resource Location.
     */
    readonly location: pulumi.Input<string>;
    /**
     * Unique name of the app to create or update. To create or update a deployment slot, use the {slot} parameter.
     */
    readonly name: pulumi.Input<string>;
    /**
     * <code>true</code> if reserved; otherwise, <code>false</code>.
     */
    readonly reserved?: pulumi.Input<boolean>;
    /**
     * Name of the resource group to which the resource belongs.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * <code>true</code> to stop SCM (KUDU) site when the app is stopped; otherwise, <code>false</code>. The default is <code>false</code>.
     */
    readonly scmSiteAlsoStopped?: pulumi.Input<boolean>;
    /**
     * Resource ID of the associated App Service plan, formatted as: "/subscriptions/{subscriptionID}/resourceGroups/{groupName}/providers/Microsoft.Web/serverfarms/{appServicePlanName}".
     */
    readonly serverFarmId?: pulumi.Input<string>;
    /**
     * Configuration of the app.
     */
    readonly siteConfig?: pulumi.Input<inputs.web.v20160801.SiteConfig>;
    /**
     * If true, custom (non *.azurewebsites.net) domains associated with web app are not verified.
     */
    readonly skipCustomDomainVerification?: pulumi.Input<boolean>;
    /**
     * If true web app hostname is not registered with DNS on creation. This parameter is
     *  only used for app creation.
     */
    readonly skipDnsRegistration?: pulumi.Input<boolean>;
    /**
     * Name of the deployment slot to create or update. By default, this API attempts to create or modify the production slot.
     */
    readonly slot: pulumi.Input<string>;
    /**
     * If specified during app creation, the app is created from a previous snapshot.
     */
    readonly snapshotInfo?: pulumi.Input<inputs.web.v20160801.SnapshotRecoveryRequest>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Time to live in seconds for web app's default domain name.
     */
    readonly ttlInSeconds?: pulumi.Input<string>;
}
