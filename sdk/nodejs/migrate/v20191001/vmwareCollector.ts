// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * ## Example Usage
 * ### VMwareCollectors_Create
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const vMwareCollector = new azurerm.migrate.v20191001.VMwareCollector("vMwareCollector", {
 *     eTag: "\"01003d32-0000-0d00-0000-5d74d2e50000\"",
 *     projectName: "abgoyalWEselfhostb72bproject",
 *     resourceGroupName: "abgoyal-westEurope",
 *     vmWareCollectorName: "PortalvCenterbc2fcollector",
 * });
 *
 * ```
 */
export class VMwareCollector extends pulumi.CustomResource {
    /**
     * Get an existing VMwareCollector resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): VMwareCollector {
        return new VMwareCollector(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:migrate/v20191001:VMwareCollector';

    /**
     * Returns true if the given object is an instance of VMwareCollector.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VMwareCollector {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VMwareCollector.__pulumiType;
    }

    public readonly eTag!: pulumi.Output<string | undefined>;
    public /*out*/ readonly name!: pulumi.Output<string>;
    public readonly properties!: pulumi.Output<outputs.migrate.v20191001.CollectorPropertiesResponse>;
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a VMwareCollector resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VMwareCollectorArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VMwareCollectorArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as VMwareCollectorArgs | undefined;
            if (!args || args.projectName === undefined) {
                throw new Error("Missing required property 'projectName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.vmWareCollectorName === undefined) {
                throw new Error("Missing required property 'vmWareCollectorName'");
            }
            inputs["eTag"] = args ? args.eTag : undefined;
            inputs["projectName"] = args ? args.projectName : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["vmWareCollectorName"] = args ? args.vmWareCollectorName : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:migrate/latest:VMwareCollector" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(VMwareCollector.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a VMwareCollector resource.
 */
export interface VMwareCollectorArgs {
    readonly eTag?: pulumi.Input<string>;
    /**
     * Name of the Azure Migrate project.
     */
    readonly projectName: pulumi.Input<string>;
    readonly properties?: pulumi.Input<inputs.migrate.v20191001.CollectorProperties>;
    /**
     * Name of the Azure Resource Group that project is part of.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Unique name of a VMware collector within a project.
     */
    readonly vmWareCollectorName: pulumi.Input<string>;
}
