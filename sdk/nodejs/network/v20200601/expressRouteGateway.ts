// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * ExpressRoute gateway resource.
 */
export class ExpressRouteGateway extends pulumi.CustomResource {
    /**
     * Get an existing ExpressRouteGateway resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ExpressRouteGateway {
        return new ExpressRouteGateway(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network/v20200601:ExpressRouteGateway';

    /**
     * Returns true if the given object is an instance of ExpressRouteGateway.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ExpressRouteGateway {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ExpressRouteGateway.__pulumiType;
    }

    /**
     * Configuration for auto scaling.
     */
    public readonly autoScaleConfiguration!: pulumi.Output<outputs.network.v20200601.ExpressRouteGatewayPropertiesResponseAutoScaleConfiguration | undefined>;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * List of ExpressRoute connections to the ExpressRoute gateway.
     */
    public /*out*/ readonly expressRouteConnections!: pulumi.Output<outputs.network.v20200601.ExpressRouteConnectionResponse[]>;
    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The provisioning state of the express route gateway resource.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The Virtual Hub where the ExpressRoute gateway is or will be deployed.
     */
    public readonly virtualHub!: pulumi.Output<outputs.network.v20200601.VirtualHubIdResponse>;

    /**
     * Create a ExpressRouteGateway resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ExpressRouteGatewayArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.expressRouteGatewayName === undefined) {
                throw new Error("Missing required property 'expressRouteGatewayName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.virtualHub === undefined) {
                throw new Error("Missing required property 'virtualHub'");
            }
            inputs["autoScaleConfiguration"] = args ? args.autoScaleConfiguration : undefined;
            inputs["expressRouteGatewayName"] = args ? args.expressRouteGatewayName : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["virtualHub"] = args ? args.virtualHub : undefined;
            inputs["etag"] = undefined /*out*/;
            inputs["expressRouteConnections"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["autoScaleConfiguration"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["expressRouteConnections"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["virtualHub"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:network/latest:ExpressRouteGateway" }, { type: "azurerm:network/v20180801:ExpressRouteGateway" }, { type: "azurerm:network/v20181001:ExpressRouteGateway" }, { type: "azurerm:network/v20181101:ExpressRouteGateway" }, { type: "azurerm:network/v20181201:ExpressRouteGateway" }, { type: "azurerm:network/v20190201:ExpressRouteGateway" }, { type: "azurerm:network/v20190401:ExpressRouteGateway" }, { type: "azurerm:network/v20190601:ExpressRouteGateway" }, { type: "azurerm:network/v20190701:ExpressRouteGateway" }, { type: "azurerm:network/v20190801:ExpressRouteGateway" }, { type: "azurerm:network/v20190901:ExpressRouteGateway" }, { type: "azurerm:network/v20191101:ExpressRouteGateway" }, { type: "azurerm:network/v20191201:ExpressRouteGateway" }, { type: "azurerm:network/v20200301:ExpressRouteGateway" }, { type: "azurerm:network/v20200401:ExpressRouteGateway" }, { type: "azurerm:network/v20200501:ExpressRouteGateway" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(ExpressRouteGateway.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ExpressRouteGateway resource.
 */
export interface ExpressRouteGatewayArgs {
    /**
     * Configuration for auto scaling.
     */
    readonly autoScaleConfiguration?: pulumi.Input<inputs.network.v20200601.ExpressRouteGatewayPropertiesAutoScaleConfiguration>;
    /**
     * The name of the ExpressRoute gateway.
     */
    readonly expressRouteGatewayName: pulumi.Input<string>;
    /**
     * Resource ID.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * Resource location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The Virtual Hub where the ExpressRoute gateway is or will be deployed.
     */
    readonly virtualHub: pulumi.Input<inputs.network.v20200601.VirtualHubId>;
}
