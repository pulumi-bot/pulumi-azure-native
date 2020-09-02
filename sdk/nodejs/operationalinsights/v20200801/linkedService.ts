// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * The top level Linked service resource container.
 *
 * ## Example Usage
 * ### LinkedServicesCreate
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const linkedService = new azurerm.operationalinsights.v20200801.LinkedService("linkedService", {
 *     linkedServiceName: "Cluster",
 *     resourceGroupName: "mms-eus",
 *     workspaceName: "TestLinkWS",
 *     writeAccessResourceId: "/subscriptions/00000000-0000-0000-0000-00000000000/resourceGroups/mms-eus/providers/Microsoft.OperationalInsights/clusters/testcluster",
 * });
 *
 * ```
 */
export class LinkedService extends pulumi.CustomResource {
    /**
     * Get an existing LinkedService resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): LinkedService {
        return new LinkedService(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:operationalinsights/v20200801:LinkedService';

    /**
     * Returns true if the given object is an instance of LinkedService.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LinkedService {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LinkedService.__pulumiType;
    }

    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The provisioning state of the linked service.
     */
    public readonly provisioningState!: pulumi.Output<string | undefined>;
    /**
     * The resource id of the resource that will be linked to the workspace. This should be used for linking resources which require read access
     */
    public readonly resourceId!: pulumi.Output<string | undefined>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The resource id of the resource that will be linked to the workspace. This should be used for linking resources which require write access
     */
    public readonly writeAccessResourceId!: pulumi.Output<string | undefined>;

    /**
     * Create a LinkedService resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LinkedServiceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LinkedServiceArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as LinkedServiceArgs | undefined;
            if (!args || args.linkedServiceName === undefined) {
                throw new Error("Missing required property 'linkedServiceName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.workspaceName === undefined) {
                throw new Error("Missing required property 'workspaceName'");
            }
            inputs["linkedServiceName"] = args ? args.linkedServiceName : undefined;
            inputs["provisioningState"] = args ? args.provisioningState : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["resourceId"] = args ? args.resourceId : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["workspaceName"] = args ? args.workspaceName : undefined;
            inputs["writeAccessResourceId"] = args ? args.writeAccessResourceId : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:operationalinsights/latest:LinkedService" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(LinkedService.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a LinkedService resource.
 */
export interface LinkedServiceArgs {
    /**
     * Name of the linkedServices resource
     */
    readonly linkedServiceName: pulumi.Input<string>;
    /**
     * The provisioning state of the linked service.
     */
    readonly provisioningState?: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The resource id of the resource that will be linked to the workspace. This should be used for linking resources which require read access
     */
    readonly resourceId?: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the workspace.
     */
    readonly workspaceName: pulumi.Input<string>;
    /**
     * The resource id of the resource that will be linked to the workspace. This should be used for linking resources which require write access
     */
    readonly writeAccessResourceId?: pulumi.Input<string>;
}
