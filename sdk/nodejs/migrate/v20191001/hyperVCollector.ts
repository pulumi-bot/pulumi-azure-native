// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * ## Example Usage
 * ### HyperVCollectors_Create
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const hyperVCollector = new azurerm.migrate.v20191001.HyperVCollector("hyperVCollector", {
 *     eTag: "\"00000981-0000-0300-0000-5d74cd5f0000\"",
 *     hyperVCollectorName: "migrateprojectce73collector",
 *     projectName: "migrateprojectce73project",
 *     resourceGroupName: "contosoithyperv",
 * });
 *
 * ```
 */
export class HyperVCollector extends pulumi.CustomResource {
    /**
     * Get an existing HyperVCollector resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): HyperVCollector {
        return new HyperVCollector(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:migrate/v20191001:HyperVCollector';

    /**
     * Returns true if the given object is an instance of HyperVCollector.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is HyperVCollector {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === HyperVCollector.__pulumiType;
    }

    public readonly eTag!: pulumi.Output<string | undefined>;
    public /*out*/ readonly name!: pulumi.Output<string>;
    public readonly properties!: pulumi.Output<outputs.migrate.v20191001.CollectorPropertiesResponse>;
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a HyperVCollector resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HyperVCollectorArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.hyperVCollectorName === undefined) {
                throw new Error("Missing required property 'hyperVCollectorName'");
            }
            if (!args || args.projectName === undefined) {
                throw new Error("Missing required property 'projectName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["eTag"] = args ? args.eTag : undefined;
            inputs["hyperVCollectorName"] = args ? args.hyperVCollectorName : undefined;
            inputs["projectName"] = args ? args.projectName : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["eTag"] = undefined /*out*/;
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
        const aliasOpts = { aliases: [{ type: "azurerm:migrate/latest:HyperVCollector" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(HyperVCollector.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a HyperVCollector resource.
 */
export interface HyperVCollectorArgs {
    readonly eTag?: pulumi.Input<string>;
    /**
     * Unique name of a Hyper-V collector within a project.
     */
    readonly hyperVCollectorName: pulumi.Input<string>;
    /**
     * Name of the Azure Migrate project.
     */
    readonly projectName: pulumi.Input<string>;
    readonly properties?: pulumi.Input<inputs.migrate.v20191001.CollectorProperties>;
    /**
     * Name of the Azure Resource Group that project is part of.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
