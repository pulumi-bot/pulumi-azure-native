// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * The management group details.
 */
export class ManagementGroup extends pulumi.CustomResource {
    /**
     * Get an existing ManagementGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ManagementGroup {
        return new ManagementGroup(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:management/v20171101preview:ManagementGroup';

    /**
     * Returns true if the given object is an instance of ManagementGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ManagementGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ManagementGroup.__pulumiType;
    }

    /**
     * The list of children.
     */
    public /*out*/ readonly children!: pulumi.Output<outputs.management.v20171101preview.ManagementGroupChildInfoResponse[] | undefined>;
    /**
     * The details of a management group.
     */
    public /*out*/ readonly details!: pulumi.Output<outputs.management.v20171101preview.ManagementGroupDetailsResponse | undefined>;
    /**
     * The friendly name of the management group.
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * The name of the management group. For example, 00000000-0000-0000-0000-000000000000
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The AAD Tenant ID associated with the management group. For example, 00000000-0000-0000-0000-000000000000
     */
    public /*out*/ readonly tenantId!: pulumi.Output<string | undefined>;
    /**
     * The type of the resource.  For example, /providers/Microsoft.Management/managementGroups
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ManagementGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ManagementGroupArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.groupId === undefined) {
                throw new Error("Missing required property 'groupId'");
            }
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["groupId"] = args ? args.groupId : undefined;
            inputs["parentId"] = args ? args.parentId : undefined;
            inputs["children"] = undefined /*out*/;
            inputs["details"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["tenantId"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["children"] = undefined /*out*/;
            inputs["details"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["tenantId"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:management/latest:ManagementGroup" }, { type: "azurerm:management/v20180101preview:ManagementGroup" }, { type: "azurerm:management/v20180301preview:ManagementGroup" }, { type: "azurerm:management/v20191101:ManagementGroup" }, { type: "azurerm:management/v20200201:ManagementGroup" }, { type: "azurerm:management/v20200501:ManagementGroup" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(ManagementGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ManagementGroup resource.
 */
export interface ManagementGroupArgs {
    /**
     * The friendly name of the management group.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * Management Group ID.
     */
    readonly groupId: pulumi.Input<string>;
    /**
     * (Optional) The fully qualified ID for the parent management group.  For example, /providers/Microsoft.Management/managementGroups/0000000-0000-0000-0000-000000000000
     */
    readonly parentId?: pulumi.Input<string>;
}
