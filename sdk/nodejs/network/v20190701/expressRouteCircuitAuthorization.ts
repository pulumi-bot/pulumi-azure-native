// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Authorization in an ExpressRouteCircuit resource.
 *
 * ## Example Usage
 * ### Create ExpressRouteCircuit Authorization
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const expressRouteCircuitAuthorization = new azurerm.network.v20190701.ExpressRouteCircuitAuthorization("expressRouteCircuitAuthorization", {
 *     authorizationKey: "authKey",
 *     authorizationName: "authorizatinName",
 *     authorizationUseStatus: "Available",
 *     circuitName: "circuitName",
 *     resourceGroupName: "rg1",
 * });
 *
 * ```
 */
export class ExpressRouteCircuitAuthorization extends pulumi.CustomResource {
    /**
     * Get an existing ExpressRouteCircuitAuthorization resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ExpressRouteCircuitAuthorization {
        return new ExpressRouteCircuitAuthorization(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network/v20190701:ExpressRouteCircuitAuthorization';

    /**
     * Returns true if the given object is an instance of ExpressRouteCircuitAuthorization.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ExpressRouteCircuitAuthorization {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ExpressRouteCircuitAuthorization.__pulumiType;
    }

    /**
     * The authorization key.
     */
    public readonly authorizationKey!: pulumi.Output<string | undefined>;
    /**
     * The authorization use status.
     */
    public readonly authorizationUseStatus!: pulumi.Output<string | undefined>;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The name of the resource that is unique within a resource group. This name can be used to access the resource.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * The provisioning state of the authorization resource.
     */
    public readonly provisioningState!: pulumi.Output<string | undefined>;
    /**
     * Type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ExpressRouteCircuitAuthorization resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ExpressRouteCircuitAuthorizationArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.authorizationName === undefined) {
                throw new Error("Missing required property 'authorizationName'");
            }
            if (!args || args.circuitName === undefined) {
                throw new Error("Missing required property 'circuitName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["authorizationKey"] = args ? args.authorizationKey : undefined;
            inputs["authorizationName"] = args ? args.authorizationName : undefined;
            inputs["authorizationUseStatus"] = args ? args.authorizationUseStatus : undefined;
            inputs["circuitName"] = args ? args.circuitName : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["provisioningState"] = args ? args.provisioningState : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["etag"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["authorizationKey"] = undefined /*out*/;
            inputs["authorizationUseStatus"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:network/latest:ExpressRouteCircuitAuthorization" }, { type: "azurerm:network/v20150501preview:ExpressRouteCircuitAuthorization" }, { type: "azurerm:network/v20150615:ExpressRouteCircuitAuthorization" }, { type: "azurerm:network/v20160330:ExpressRouteCircuitAuthorization" }, { type: "azurerm:network/v20160601:ExpressRouteCircuitAuthorization" }, { type: "azurerm:network/v20160901:ExpressRouteCircuitAuthorization" }, { type: "azurerm:network/v20161201:ExpressRouteCircuitAuthorization" }, { type: "azurerm:network/v20170301:ExpressRouteCircuitAuthorization" }, { type: "azurerm:network/v20170601:ExpressRouteCircuitAuthorization" }, { type: "azurerm:network/v20170801:ExpressRouteCircuitAuthorization" }, { type: "azurerm:network/v20170901:ExpressRouteCircuitAuthorization" }, { type: "azurerm:network/v20171001:ExpressRouteCircuitAuthorization" }, { type: "azurerm:network/v20171101:ExpressRouteCircuitAuthorization" }, { type: "azurerm:network/v20180101:ExpressRouteCircuitAuthorization" }, { type: "azurerm:network/v20180201:ExpressRouteCircuitAuthorization" }, { type: "azurerm:network/v20180401:ExpressRouteCircuitAuthorization" }, { type: "azurerm:network/v20180601:ExpressRouteCircuitAuthorization" }, { type: "azurerm:network/v20180701:ExpressRouteCircuitAuthorization" }, { type: "azurerm:network/v20180801:ExpressRouteCircuitAuthorization" }, { type: "azurerm:network/v20181001:ExpressRouteCircuitAuthorization" }, { type: "azurerm:network/v20181101:ExpressRouteCircuitAuthorization" }, { type: "azurerm:network/v20181201:ExpressRouteCircuitAuthorization" }, { type: "azurerm:network/v20190201:ExpressRouteCircuitAuthorization" }, { type: "azurerm:network/v20190401:ExpressRouteCircuitAuthorization" }, { type: "azurerm:network/v20190601:ExpressRouteCircuitAuthorization" }, { type: "azurerm:network/v20190801:ExpressRouteCircuitAuthorization" }, { type: "azurerm:network/v20190901:ExpressRouteCircuitAuthorization" }, { type: "azurerm:network/v20191101:ExpressRouteCircuitAuthorization" }, { type: "azurerm:network/v20191201:ExpressRouteCircuitAuthorization" }, { type: "azurerm:network/v20200301:ExpressRouteCircuitAuthorization" }, { type: "azurerm:network/v20200401:ExpressRouteCircuitAuthorization" }, { type: "azurerm:network/v20200501:ExpressRouteCircuitAuthorization" }, { type: "azurerm:network/v20200601:ExpressRouteCircuitAuthorization" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(ExpressRouteCircuitAuthorization.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ExpressRouteCircuitAuthorization resource.
 */
export interface ExpressRouteCircuitAuthorizationArgs {
    /**
     * The authorization key.
     */
    readonly authorizationKey?: pulumi.Input<string>;
    /**
     * The name of the authorization.
     */
    readonly authorizationName: pulumi.Input<string>;
    /**
     * The authorization use status.
     */
    readonly authorizationUseStatus?: pulumi.Input<string>;
    /**
     * The name of the express route circuit.
     */
    readonly circuitName: pulumi.Input<string>;
    /**
     * Resource ID.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * The name of the resource that is unique within a resource group. This name can be used to access the resource.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The provisioning state of the authorization resource.
     */
    readonly provisioningState?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
