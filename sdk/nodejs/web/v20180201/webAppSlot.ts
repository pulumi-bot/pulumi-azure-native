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
    public static readonly __pulumiType = 'azurerm:web/v20180201:WebAppSlot';

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
     * Managed service identity.
     */
    public readonly identity!: pulumi.Output<outputs.web.v20180201.ManagedServiceIdentityResponse | undefined>;
    /**
     * Kind of resource.
     */
    public readonly kind!: pulumi.Output<string | undefined>;
    /**
     * Resource Location.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Resource Name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Site resource specific properties
     */
    public /*out*/ readonly properties!: pulumi.Output<outputs.web.v20180201.SiteResponseProperties>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

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
            inputs["clientAffinityEnabled"] = args ? args.clientAffinityEnabled : undefined;
            inputs["clientCertEnabled"] = args ? args.clientCertEnabled : undefined;
            inputs["clientCertExclusionPaths"] = args ? args.clientCertExclusionPaths : undefined;
            inputs["cloningInfo"] = args ? args.cloningInfo : undefined;
            inputs["containerSize"] = args ? args.containerSize : undefined;
            inputs["dailyMemoryTimeQuota"] = args ? args.dailyMemoryTimeQuota : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["geoDistributions"] = args ? args.geoDistributions : undefined;
            inputs["hostNameSslStates"] = args ? args.hostNameSslStates : undefined;
            inputs["hostNamesDisabled"] = args ? args.hostNamesDisabled : undefined;
            inputs["hostingEnvironmentProfile"] = args ? args.hostingEnvironmentProfile : undefined;
            inputs["httpsOnly"] = args ? args.httpsOnly : undefined;
            inputs["hyperV"] = args ? args.hyperV : undefined;
            inputs["identity"] = args ? args.identity : undefined;
            inputs["isXenon"] = args ? args.isXenon : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["redundancyMode"] = args ? args.redundancyMode : undefined;
            inputs["reserved"] = args ? args.reserved : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["scmSiteAlsoStopped"] = args ? args.scmSiteAlsoStopped : undefined;
            inputs["serverFarmId"] = args ? args.serverFarmId : undefined;
            inputs["siteConfig"] = args ? args.siteConfig : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["properties"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
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
     * client certificate authentication comma-separated exclusion paths
     */
    readonly clientCertExclusionPaths?: pulumi.Input<string>;
    /**
     * If specified during app creation, the app is cloned from a source app.
     */
    readonly cloningInfo?: pulumi.Input<inputs.web.v20180201.CloningInfo>;
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
     * GeoDistributions for this site
     */
    readonly geoDistributions?: pulumi.Input<pulumi.Input<inputs.web.v20180201.GeoDistribution>[]>;
    /**
     * Hostname SSL states are used to manage the SSL bindings for app's hostnames.
     */
    readonly hostNameSslStates?: pulumi.Input<pulumi.Input<inputs.web.v20180201.HostNameSslState>[]>;
    /**
     * <code>true</code> to disable the public hostnames of the app; otherwise, <code>false</code>.
     *  If <code>true</code>, the app is only accessible via API management process.
     */
    readonly hostNamesDisabled?: pulumi.Input<boolean>;
    /**
     * App Service Environment to use for the app.
     */
    readonly hostingEnvironmentProfile?: pulumi.Input<inputs.web.v20180201.HostingEnvironmentProfile>;
    /**
     * HttpsOnly: configures a web site to accept only https requests. Issues redirect for
     * http requests
     */
    readonly httpsOnly?: pulumi.Input<boolean>;
    /**
     * Hyper-V sandbox.
     */
    readonly hyperV?: pulumi.Input<boolean>;
    /**
     * Managed service identity.
     */
    readonly identity?: pulumi.Input<inputs.web.v20180201.ManagedServiceIdentity>;
    /**
     * Obsolete: Hyper-V sandbox.
     */
    readonly isXenon?: pulumi.Input<boolean>;
    /**
     * Kind of resource.
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * Resource Location.
     */
    readonly location: pulumi.Input<string>;
    /**
     * Name of the deployment slot to create or update. By default, this API attempts to create or modify the production slot.
     */
    readonly name: pulumi.Input<string>;
    /**
     * Site redundancy mode
     */
    readonly redundancyMode?: pulumi.Input<string>;
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
    readonly siteConfig?: pulumi.Input<inputs.web.v20180201.SiteConfig>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
