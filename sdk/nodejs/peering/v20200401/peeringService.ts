// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Peering Service
 */
export class PeeringService extends pulumi.CustomResource {
    /**
     * Get an existing PeeringService resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PeeringService {
        return new PeeringService(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:peering/v20200401:PeeringService';

    /**
     * Returns true if the given object is an instance of PeeringService.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PeeringService {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PeeringService.__pulumiType;
    }

    /**
     * The location of the resource.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The properties that define a peering service.
     */
    public /*out*/ readonly properties!: pulumi.Output<outputs.peering.v20200401.PeeringServicePropertiesResponse>;
    /**
     * The SKU that defines the type of the peering service.
     */
    public readonly sku!: pulumi.Output<outputs.peering.v20200401.PeeringServiceSkuResponse | undefined>;
    /**
     * The resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a PeeringService resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PeeringServiceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PeeringServiceArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as PeeringServiceArgs | undefined;
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
            inputs["peeringServiceLocation"] = args ? args.peeringServiceLocation : undefined;
            inputs["peeringServiceProvider"] = args ? args.peeringServiceProvider : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
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
        super(PeeringService.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a PeeringService resource.
 */
export interface PeeringServiceArgs {
    /**
     * The location of the resource.
     */
    readonly location: pulumi.Input<string>;
    /**
     * The name of the peering service.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The PeeringServiceLocation of the Customer.
     */
    readonly peeringServiceLocation?: pulumi.Input<string>;
    /**
     * The MAPS Provider Name.
     */
    readonly peeringServiceProvider?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The SKU that defines the type of the peering service.
     */
    readonly sku?: pulumi.Input<inputs.peering.v20200401.PeeringServiceSku>;
    /**
     * The resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
