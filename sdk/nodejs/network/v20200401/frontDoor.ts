// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Front Door represents a collection of backend endpoints to route traffic to along with rules that specify how traffic is sent there.
 *
 * ## Example Usage
 * ### Create or update specific Front Door
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const frontDoor = new azurerm.network.v20200401.FrontDoor("frontDoor", {
 *     backendPools: [{
 *         name: "backendPool1",
 *     }],
 *     backendPoolsSettings: {
 *         enforceCertificateNameCheck: "Enabled",
 *         sendRecvTimeoutSeconds: 60,
 *     },
 *     enabledState: "Enabled",
 *     frontDoorName: "frontDoor1",
 *     frontendEndpoints: [
 *         {
 *             name: "frontendEndpoint1",
 *         },
 *         {
 *             name: "default",
 *         },
 *     ],
 *     healthProbeSettings: [{
 *         name: "healthProbeSettings1",
 *     }],
 *     loadBalancingSettings: [{
 *         name: "loadBalancingSettings1",
 *     }],
 *     location: "westus",
 *     resourceGroupName: "rg1",
 *     routingRules: [{
 *         name: "routingRule1",
 *     }],
 *     tags: {
 *         tag1: "value1",
 *         tag2: "value2",
 *     },
 * });
 *
 * ```
 */
export class FrontDoor extends pulumi.CustomResource {
    /**
     * Get an existing FrontDoor resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): FrontDoor {
        return new FrontDoor(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network/v20200401:FrontDoor';

    /**
     * Returns true if the given object is an instance of FrontDoor.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FrontDoor {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FrontDoor.__pulumiType;
    }

    /**
     * Backend pools available to routing rules.
     */
    public readonly backendPools!: pulumi.Output<outputs.network.v20200401.BackendPoolResponse[] | undefined>;
    /**
     * Settings for all backendPools
     */
    public readonly backendPoolsSettings!: pulumi.Output<outputs.network.v20200401.BackendPoolsSettingsResponse | undefined>;
    /**
     * The host that each frontendEndpoint must CNAME to.
     */
    public /*out*/ readonly cname!: pulumi.Output<string>;
    /**
     * Operational status of the Front Door load balancer. Permitted values are 'Enabled' or 'Disabled'
     */
    public readonly enabledState!: pulumi.Output<string | undefined>;
    /**
     * A friendly name for the frontDoor
     */
    public readonly friendlyName!: pulumi.Output<string | undefined>;
    /**
     * The Id of the frontdoor.
     */
    public /*out*/ readonly frontdoorId!: pulumi.Output<string>;
    /**
     * Frontend endpoints available to routing rules.
     */
    public readonly frontendEndpoints!: pulumi.Output<outputs.network.v20200401.FrontendEndpointResponse[] | undefined>;
    /**
     * Health probe settings associated with this Front Door instance.
     */
    public readonly healthProbeSettings!: pulumi.Output<outputs.network.v20200401.HealthProbeSettingsModelResponse[] | undefined>;
    /**
     * Load balancing settings associated with this Front Door instance.
     */
    public readonly loadBalancingSettings!: pulumi.Output<outputs.network.v20200401.LoadBalancingSettingsModelResponse[] | undefined>;
    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Provisioning state of the Front Door.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Resource status of the Front Door.
     */
    public readonly resourceState!: pulumi.Output<string | undefined>;
    /**
     * Routing rules associated with this Front Door.
     */
    public readonly routingRules!: pulumi.Output<outputs.network.v20200401.RoutingRuleResponse[] | undefined>;
    /**
     * Rules Engine Configurations available to routing rules.
     */
    public /*out*/ readonly rulesEngines!: pulumi.Output<outputs.network.v20200401.RulesEngineResponse[]>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a FrontDoor resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FrontDoorArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.frontDoorName === undefined) {
                throw new Error("Missing required property 'frontDoorName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["backendPools"] = args ? args.backendPools : undefined;
            inputs["backendPoolsSettings"] = args ? args.backendPoolsSettings : undefined;
            inputs["enabledState"] = args ? args.enabledState : undefined;
            inputs["friendlyName"] = args ? args.friendlyName : undefined;
            inputs["frontDoorName"] = args ? args.frontDoorName : undefined;
            inputs["frontendEndpoints"] = args ? args.frontendEndpoints : undefined;
            inputs["healthProbeSettings"] = args ? args.healthProbeSettings : undefined;
            inputs["loadBalancingSettings"] = args ? args.loadBalancingSettings : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["resourceState"] = args ? args.resourceState : undefined;
            inputs["routingRules"] = args ? args.routingRules : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["cname"] = undefined /*out*/;
            inputs["frontdoorId"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["rulesEngines"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["backendPools"] = undefined /*out*/;
            inputs["backendPoolsSettings"] = undefined /*out*/;
            inputs["cname"] = undefined /*out*/;
            inputs["enabledState"] = undefined /*out*/;
            inputs["friendlyName"] = undefined /*out*/;
            inputs["frontdoorId"] = undefined /*out*/;
            inputs["frontendEndpoints"] = undefined /*out*/;
            inputs["healthProbeSettings"] = undefined /*out*/;
            inputs["loadBalancingSettings"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["resourceState"] = undefined /*out*/;
            inputs["routingRules"] = undefined /*out*/;
            inputs["rulesEngines"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:network/latest:FrontDoor" }, { type: "azurerm:network/v20180801:FrontDoor" }, { type: "azurerm:network/v20190401:FrontDoor" }, { type: "azurerm:network/v20190501:FrontDoor" }, { type: "azurerm:network/v20200101:FrontDoor" }, { type: "azurerm:network/v20200501:FrontDoor" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(FrontDoor.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a FrontDoor resource.
 */
export interface FrontDoorArgs {
    /**
     * Backend pools available to routing rules.
     */
    readonly backendPools?: pulumi.Input<pulumi.Input<inputs.network.v20200401.BackendPool>[]>;
    /**
     * Settings for all backendPools
     */
    readonly backendPoolsSettings?: pulumi.Input<inputs.network.v20200401.BackendPoolsSettings>;
    /**
     * Operational status of the Front Door load balancer. Permitted values are 'Enabled' or 'Disabled'
     */
    readonly enabledState?: pulumi.Input<string>;
    /**
     * A friendly name for the frontDoor
     */
    readonly friendlyName?: pulumi.Input<string>;
    /**
     * Name of the Front Door which is globally unique.
     */
    readonly frontDoorName: pulumi.Input<string>;
    /**
     * Frontend endpoints available to routing rules.
     */
    readonly frontendEndpoints?: pulumi.Input<pulumi.Input<inputs.network.v20200401.FrontendEndpoint>[]>;
    /**
     * Health probe settings associated with this Front Door instance.
     */
    readonly healthProbeSettings?: pulumi.Input<pulumi.Input<inputs.network.v20200401.HealthProbeSettingsModel>[]>;
    /**
     * Load balancing settings associated with this Front Door instance.
     */
    readonly loadBalancingSettings?: pulumi.Input<pulumi.Input<inputs.network.v20200401.LoadBalancingSettingsModel>[]>;
    /**
     * Resource location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Name of the Resource group within the Azure subscription.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Resource status of the Front Door.
     */
    readonly resourceState?: pulumi.Input<string>;
    /**
     * Routing rules associated with this Front Door.
     */
    readonly routingRules?: pulumi.Input<pulumi.Input<inputs.network.v20200401.RoutingRule>[]>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
