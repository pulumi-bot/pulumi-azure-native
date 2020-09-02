// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * OpenShift Managed cluster.
 *
 * ## Create/Update OpenShift Managed Cluster
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const openShiftManagedCluster = new azurerm.containerservice.latest.OpenShiftManagedCluster("openShiftManagedCluster", {
 *     agentPoolProfiles: [
 *         {
 *             count: 2,
 *             name: "infra",
 *             osType: "Linux",
 *             role: "infra",
 *             subnetCidr: "10.0.0.0/24",
 *             vmSize: "Standard_D4s_v3",
 *         },
 *         {
 *             count: 4,
 *             name: "compute",
 *             osType: "Linux",
 *             role: "compute",
 *             subnetCidr: "10.0.0.0/24",
 *             vmSize: "Standard_D4s_v3",
 *         },
 *     ],
 *     authProfile: {
 *         identityProviders: [{
 *             name: "Azure AD",
 *             provider: {
 *                 kind: "AADIdentityProvider",
 *             },
 *         }],
 *     },
 *     location: "location1",
 *     masterPoolProfile: {
 *         count: 3,
 *         name: "master",
 *         osType: "Linux",
 *         subnetCidr: "10.0.0.0/24",
 *         vmSize: "Standard_D4s_v3",
 *     },
 *     networkProfile: {
 *         vnetCidr: "10.0.0.0/8",
 *     },
 *     openShiftVersion: "v3.11",
 *     resourceGroupName: "rg1",
 *     resourceName: "clustername1",
 *     routerProfiles: [{
 *         name: "default",
 *     }],
 *     tags: {
 *         archv2: "",
 *         tier: "production",
 *     },
 * });
 *
 * ```
 */
export class OpenShiftManagedCluster extends pulumi.CustomResource {
    /**
     * Get an existing OpenShiftManagedCluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): OpenShiftManagedCluster {
        return new OpenShiftManagedCluster(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:containerservice/latest:OpenShiftManagedCluster';

    /**
     * Returns true if the given object is an instance of OpenShiftManagedCluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OpenShiftManagedCluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OpenShiftManagedCluster.__pulumiType;
    }

    /**
     * Configuration of OpenShift cluster VMs.
     */
    public readonly agentPoolProfiles!: pulumi.Output<outputs.containerservice.latest.OpenShiftManagedClusterAgentPoolProfileResponse[] | undefined>;
    /**
     * Configures OpenShift authentication.
     */
    public readonly authProfile!: pulumi.Output<outputs.containerservice.latest.OpenShiftManagedClusterAuthProfileResponse | undefined>;
    /**
     * Version of OpenShift specified when creating the cluster.
     */
    public /*out*/ readonly clusterVersion!: pulumi.Output<string>;
    /**
     * Service generated FQDN for OpenShift API server loadbalancer internal hostname.
     */
    public /*out*/ readonly fqdn!: pulumi.Output<string>;
    /**
     * Resource location
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Configuration for OpenShift master VMs.
     */
    public readonly masterPoolProfile!: pulumi.Output<outputs.containerservice.latest.OpenShiftManagedClusterMasterPoolProfileResponse | undefined>;
    /**
     * Resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Configuration for OpenShift networking.
     */
    public readonly networkProfile!: pulumi.Output<outputs.containerservice.latest.NetworkProfileResponse | undefined>;
    /**
     * Version of OpenShift specified when creating the cluster.
     */
    public readonly openShiftVersion!: pulumi.Output<string>;
    /**
     * Define the resource plan as required by ARM for billing purposes
     */
    public readonly plan!: pulumi.Output<outputs.containerservice.latest.PurchasePlanResponse | undefined>;
    /**
     * The current deployment or provisioning state, which only appears in the response.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Service generated FQDN for OpenShift API server.
     */
    public /*out*/ readonly publicHostname!: pulumi.Output<string>;
    /**
     * Configuration for OpenShift router(s).
     */
    public readonly routerProfiles!: pulumi.Output<outputs.containerservice.latest.OpenShiftRouterProfileResponse[] | undefined>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a OpenShiftManagedCluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OpenShiftManagedClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OpenShiftManagedClusterArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as OpenShiftManagedClusterArgs | undefined;
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.openShiftVersion === undefined) {
                throw new Error("Missing required property 'openShiftVersion'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.resourceName === undefined) {
                throw new Error("Missing required property 'resourceName'");
            }
            inputs["agentPoolProfiles"] = args ? args.agentPoolProfiles : undefined;
            inputs["authProfile"] = args ? args.authProfile : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["masterPoolProfile"] = args ? args.masterPoolProfile : undefined;
            inputs["networkProfile"] = args ? args.networkProfile : undefined;
            inputs["openShiftVersion"] = args ? args.openShiftVersion : undefined;
            inputs["plan"] = args ? args.plan : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["resourceName"] = args ? args.resourceName : undefined;
            inputs["routerProfiles"] = args ? args.routerProfiles : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["clusterVersion"] = undefined /*out*/;
            inputs["fqdn"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["publicHostname"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:containerservice/v20190430:OpenShiftManagedCluster" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(OpenShiftManagedCluster.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a OpenShiftManagedCluster resource.
 */
export interface OpenShiftManagedClusterArgs {
    /**
     * Configuration of OpenShift cluster VMs.
     */
    readonly agentPoolProfiles?: pulumi.Input<pulumi.Input<inputs.containerservice.latest.OpenShiftManagedClusterAgentPoolProfile>[]>;
    /**
     * Configures OpenShift authentication.
     */
    readonly authProfile?: pulumi.Input<inputs.containerservice.latest.OpenShiftManagedClusterAuthProfile>;
    /**
     * Resource location
     */
    readonly location: pulumi.Input<string>;
    /**
     * Configuration for OpenShift master VMs.
     */
    readonly masterPoolProfile?: pulumi.Input<inputs.containerservice.latest.OpenShiftManagedClusterMasterPoolProfile>;
    /**
     * Configuration for OpenShift networking.
     */
    readonly networkProfile?: pulumi.Input<inputs.containerservice.latest.NetworkProfile>;
    /**
     * Version of OpenShift specified when creating the cluster.
     */
    readonly openShiftVersion: pulumi.Input<string>;
    /**
     * Define the resource plan as required by ARM for billing purposes
     */
    readonly plan?: pulumi.Input<inputs.containerservice.latest.PurchasePlan>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the OpenShift managed cluster resource.
     */
    readonly resourceName: pulumi.Input<string>;
    /**
     * Configuration for OpenShift router(s).
     */
    readonly routerProfiles?: pulumi.Input<pulumi.Input<inputs.containerservice.latest.OpenShiftRouterProfile>[]>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}