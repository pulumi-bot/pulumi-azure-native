// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * A class represent a SignalR service resource.
 */
export class SignalR extends pulumi.CustomResource {
    /**
     * Get an existing SignalR resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): SignalR {
        return new SignalR(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:signalrservice/v20181001:SignalR';

    /**
     * Returns true if the given object is an instance of SignalR.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SignalR {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SignalR.__pulumiType;
    }

    /**
     * The GEO location of the SignalR service. e.g. West US | East US | North Central US | South Central US.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The properties of the service.
     */
    public readonly properties!: pulumi.Output<outputs.signalrservice.v20181001.SignalRPropertiesResponse>;
    /**
     * SKU of the service.
     */
    public readonly sku!: pulumi.Output<outputs.signalrservice.v20181001.ResourceSkuResponse>;
    /**
     * Tags of the service which is a list of key value pairs that describe the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string}>;
    /**
     * The type of the service - e.g. "Microsoft.SignalRService/SignalR"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a SignalR resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SignalRArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SignalRArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as SignalRArgs | undefined;
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(SignalR.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a SignalR resource.
 */
export interface SignalRArgs {
    /**
     * Azure GEO region: e.g. West US | East US | North Central US | South Central US | West Europe | North Europe | East Asia | Southeast Asia | etc. 
     * The geo region of a resource never changes after it is created.
     */
    readonly location: pulumi.Input<string>;
    /**
     * The name of the SignalR resource.
     */
    readonly name: pulumi.Input<string>;
    /**
     * Settings used to provision or configure the resource
     */
    readonly properties?: pulumi.Input<inputs.signalrservice.v20181001.SignalRCreateOrUpdateProperties>;
    /**
     * The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The billing information of the resource.(e.g. basic vs. standard)
     */
    readonly sku?: pulumi.Input<inputs.signalrservice.v20181001.ResourceSku>;
    /**
     * A list of key value pairs that describe the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}