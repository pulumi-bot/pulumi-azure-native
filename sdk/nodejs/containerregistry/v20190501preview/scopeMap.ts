// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * An object that represents a scope map for a container registry.
 */
export class ScopeMap extends pulumi.CustomResource {
    /**
     * Get an existing ScopeMap resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ScopeMap {
        return new ScopeMap(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:containerregistry/v20190501preview:ScopeMap';

    /**
     * Returns true if the given object is an instance of ScopeMap.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ScopeMap {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ScopeMap.__pulumiType;
    }

    /**
     * The list of scoped permissions for registry artifacts.
     * E.g. repositories/repository-name/content/read,
     * repositories/repository-name/metadata/write
     */
    public readonly actions!: pulumi.Output<string[]>;
    /**
     * The creation date of scope map.
     */
    public /*out*/ readonly creationDate!: pulumi.Output<string>;
    /**
     * The user friendly description of the scope map.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Provisioning state of the resource.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ScopeMap resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ScopeMapArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.actions === undefined) {
                throw new Error("Missing required property 'actions'");
            }
            if (!args || args.registryName === undefined) {
                throw new Error("Missing required property 'registryName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.scopeMapName === undefined) {
                throw new Error("Missing required property 'scopeMapName'");
            }
            inputs["actions"] = args ? args.actions : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["registryName"] = args ? args.registryName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["scopeMapName"] = args ? args.scopeMapName : undefined;
            inputs["creationDate"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["actions"] = undefined /*out*/;
            inputs["creationDate"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
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
        const aliasOpts = { aliases: [{ type: "azurerm:containerregistry/latest:ScopeMap" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(ScopeMap.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ScopeMap resource.
 */
export interface ScopeMapArgs {
    /**
     * The list of scoped permissions for registry artifacts.
     * E.g. repositories/repository-name/content/read,
     * repositories/repository-name/metadata/write
     */
    readonly actions: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The user friendly description of the scope map.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The name of the container registry.
     */
    readonly registryName: pulumi.Input<string>;
    /**
     * The name of the resource group to which the container registry belongs.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the scope map.
     */
    readonly scopeMapName: pulumi.Input<string>;
}
