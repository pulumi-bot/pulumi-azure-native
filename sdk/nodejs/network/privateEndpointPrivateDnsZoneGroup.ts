// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Private dns zone group resource.
 */
export class PrivateEndpointPrivateDnsZoneGroup extends pulumi.CustomResource {
    /**
     * Get an existing PrivateEndpointPrivateDnsZoneGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PrivateEndpointPrivateDnsZoneGroup {
        return new PrivateEndpointPrivateDnsZoneGroup(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network:PrivateEndpointPrivateDnsZoneGroup';

    /**
     * Returns true if the given object is an instance of PrivateEndpointPrivateDnsZoneGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PrivateEndpointPrivateDnsZoneGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PrivateEndpointPrivateDnsZoneGroup.__pulumiType;
    }

    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * Name of the resource that is unique within a resource group. This name can be used to access the resource.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * Properties of the private dns zone group.
     */
    public readonly properties!: pulumi.Output<outputs.network.PrivateDnsZoneGroupPropertiesFormatResponse>;

    /**
     * Create a PrivateEndpointPrivateDnsZoneGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: PrivateEndpointPrivateDnsZoneGroupArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
            if (!args || args.privateDnsZoneGroupName === undefined) {
                throw new Error("Missing required property 'privateDnsZoneGroupName'");
            }
            if (!args || args.privateEndpointName === undefined) {
                throw new Error("Missing required property 'privateEndpointName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
        inputs["id"] = args ? args.id : undefined;
        inputs["name"] = args ? args.name : undefined;
        inputs["privateDnsZoneGroupName"] = args ? args.privateDnsZoneGroupName : undefined;
        inputs["privateEndpointName"] = args ? args.privateEndpointName : undefined;
        inputs["properties"] = args ? args.properties : undefined;
        inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
        inputs["etag"] = undefined /*out*/;
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(PrivateEndpointPrivateDnsZoneGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a PrivateEndpointPrivateDnsZoneGroup resource.
 */
export interface PrivateEndpointPrivateDnsZoneGroupArgs {
    /**
     * Resource ID.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * Name of the resource that is unique within a resource group. This name can be used to access the resource.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the private dns zone group.
     */
    readonly privateDnsZoneGroupName: pulumi.Input<string>;
    /**
     * The name of the private endpoint.
     */
    readonly privateEndpointName: pulumi.Input<string>;
    /**
     * Properties of the private dns zone group.
     */
    readonly properties?: pulumi.Input<inputs.network.PrivateDnsZoneGroupPropertiesFormat>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}