// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Protection container mapping object.
 *
 * ## Example Usage
 * ### Create protection container mapping.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const replicationProtectionContainerMapping = new azurerm.recoveryservices.v20180710.ReplicationProtectionContainerMapping("replicationProtectionContainerMapping", {
 *     fabricName: "cloud1",
 *     mappingName: "cloud1protectionprofile1",
 *     protectionContainerName: "cloud_6d224fc6-f326-5d35-96de-fbf51efb3179",
 *     resourceGroupName: "resourceGroupPS1",
 *     resourceName: "vault1",
 * });
 *
 * ```
 */
export class ReplicationProtectionContainerMapping extends pulumi.CustomResource {
    /**
     * Get an existing ReplicationProtectionContainerMapping resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ReplicationProtectionContainerMapping {
        return new ReplicationProtectionContainerMapping(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:recoveryservices/v20180710:ReplicationProtectionContainerMapping';

    /**
     * Returns true if the given object is an instance of ReplicationProtectionContainerMapping.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ReplicationProtectionContainerMapping {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ReplicationProtectionContainerMapping.__pulumiType;
    }

    /**
     * Resource Location
     */
    public /*out*/ readonly location!: pulumi.Output<string | undefined>;
    /**
     * Resource Name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The custom data.
     */
    public readonly properties!: pulumi.Output<outputs.recoveryservices.v20180710.ProtectionContainerMappingPropertiesResponse>;
    /**
     * Resource Type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ReplicationProtectionContainerMapping resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ReplicationProtectionContainerMappingArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.fabricName === undefined) {
                throw new Error("Missing required property 'fabricName'");
            }
            if (!args || args.mappingName === undefined) {
                throw new Error("Missing required property 'mappingName'");
            }
            if (!args || args.protectionContainerName === undefined) {
                throw new Error("Missing required property 'protectionContainerName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.resourceName === undefined) {
                throw new Error("Missing required property 'resourceName'");
            }
            inputs["fabricName"] = args ? args.fabricName : undefined;
            inputs["mappingName"] = args ? args.mappingName : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["protectionContainerName"] = args ? args.protectionContainerName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["resourceName"] = args ? args.resourceName : undefined;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["properties"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:recoveryservices/latest:ReplicationProtectionContainerMapping" }, { type: "azurerm:recoveryservices/v20160810:ReplicationProtectionContainerMapping" }, { type: "azurerm:recoveryservices/v20180110:ReplicationProtectionContainerMapping" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(ReplicationProtectionContainerMapping.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ReplicationProtectionContainerMapping resource.
 */
export interface ReplicationProtectionContainerMappingArgs {
    /**
     * Fabric name.
     */
    readonly fabricName: pulumi.Input<string>;
    /**
     * Protection container mapping name.
     */
    readonly mappingName: pulumi.Input<string>;
    /**
     * Configure protection input properties.
     */
    readonly properties?: pulumi.Input<inputs.recoveryservices.v20180710.CreateProtectionContainerMappingInputProperties>;
    /**
     * Protection container name.
     */
    readonly protectionContainerName: pulumi.Input<string>;
    /**
     * The name of the resource group where the recovery services vault is present.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the recovery services vault.
     */
    readonly resourceName: pulumi.Input<string>;
}
