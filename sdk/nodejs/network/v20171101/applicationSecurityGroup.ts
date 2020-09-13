// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * An application security group in a resource group.
 */
export class ApplicationSecurityGroup extends pulumi.CustomResource {
    /**
     * Get an existing ApplicationSecurityGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ApplicationSecurityGroup {
        return new ApplicationSecurityGroup(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network/v20171101:ApplicationSecurityGroup';

    /**
     * Returns true if the given object is an instance of ApplicationSecurityGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApplicationSecurityGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApplicationSecurityGroup.__pulumiType;
    }

    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The provisioning state of the application security group resource. Possible values are: 'Succeeded', 'Updating', 'Deleting', and 'Failed'.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The resource GUID property of the application security group resource. It uniquely identifies a resource, even if the user changes its name or migrate the resource across subscriptions or resource groups.
     */
    public /*out*/ readonly resourceGuid!: pulumi.Output<string>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ApplicationSecurityGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApplicationSecurityGroupArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.applicationSecurityGroupName === undefined) {
                throw new Error("Missing required property 'applicationSecurityGroupName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["applicationSecurityGroupName"] = args ? args.applicationSecurityGroupName : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["etag"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["resourceGuid"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["etag"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["resourceGuid"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:network/latest:ApplicationSecurityGroup" }, { type: "azurerm:network/v20170901:ApplicationSecurityGroup" }, { type: "azurerm:network/v20171001:ApplicationSecurityGroup" }, { type: "azurerm:network/v20180101:ApplicationSecurityGroup" }, { type: "azurerm:network/v20180201:ApplicationSecurityGroup" }, { type: "azurerm:network/v20180401:ApplicationSecurityGroup" }, { type: "azurerm:network/v20180601:ApplicationSecurityGroup" }, { type: "azurerm:network/v20180701:ApplicationSecurityGroup" }, { type: "azurerm:network/v20180801:ApplicationSecurityGroup" }, { type: "azurerm:network/v20181001:ApplicationSecurityGroup" }, { type: "azurerm:network/v20181101:ApplicationSecurityGroup" }, { type: "azurerm:network/v20181201:ApplicationSecurityGroup" }, { type: "azurerm:network/v20190201:ApplicationSecurityGroup" }, { type: "azurerm:network/v20190401:ApplicationSecurityGroup" }, { type: "azurerm:network/v20190601:ApplicationSecurityGroup" }, { type: "azurerm:network/v20190701:ApplicationSecurityGroup" }, { type: "azurerm:network/v20190801:ApplicationSecurityGroup" }, { type: "azurerm:network/v20190901:ApplicationSecurityGroup" }, { type: "azurerm:network/v20191101:ApplicationSecurityGroup" }, { type: "azurerm:network/v20191201:ApplicationSecurityGroup" }, { type: "azurerm:network/v20200301:ApplicationSecurityGroup" }, { type: "azurerm:network/v20200401:ApplicationSecurityGroup" }, { type: "azurerm:network/v20200501:ApplicationSecurityGroup" }, { type: "azurerm:network/v20200601:ApplicationSecurityGroup" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(ApplicationSecurityGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ApplicationSecurityGroup resource.
 */
export interface ApplicationSecurityGroupArgs {
    /**
     * The name of the application security group.
     */
    readonly applicationSecurityGroupName: pulumi.Input<string>;
    /**
     * Resource ID.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * Resource location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
