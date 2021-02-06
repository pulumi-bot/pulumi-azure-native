// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * An Azure Monitor PrivateLinkScope definition.
 */
export class PrivateLinkScope extends pulumi.CustomResource {
    /**
     * Get an existing PrivateLinkScope resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PrivateLinkScope {
        return new PrivateLinkScope(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:insights/v20191017preview:PrivateLinkScope';

    /**
     * Returns true if the given object is an instance of PrivateLinkScope.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PrivateLinkScope {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PrivateLinkScope.__pulumiType;
    }

    /**
     * Resource location
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Azure resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * List of private endpoint connections.
     */
    public /*out*/ readonly privateEndpointConnections!: pulumi.Output<outputs.insights.v20191017preview.PrivateEndpointConnectionResponse[]>;
    /**
     * Current state of this PrivateLinkScope: whether or not is has been provisioned within the resource group it is defined. Users cannot change this value but are able to read from it. Values will include Provisioning ,Succeeded, Canceled and Failed.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Azure resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a PrivateLinkScope resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PrivateLinkScopeArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if ((!args || args.resourceGroupName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.scopeName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'scopeName'");
            }
            inputs["location"] = args ? args.location : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["scopeName"] = args ? args.scopeName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["privateEndpointConnections"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["privateEndpointConnections"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(PrivateLinkScope.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a PrivateLinkScope resource.
 */
export interface PrivateLinkScopeArgs {
    /**
     * Resource location
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the Azure Monitor PrivateLinkScope resource.
     */
    readonly scopeName: pulumi.Input<string>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
