// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Resource for OuContainer.
 */
export class OuContainer extends pulumi.CustomResource {
    /**
     * Get an existing OuContainer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): OuContainer {
        return new OuContainer(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:aad/v20170601:OuContainer';

    /**
     * Returns true if the given object is an instance of OuContainer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OuContainer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OuContainer.__pulumiType;
    }

    /**
     * The list of container accounts
     */
    public /*out*/ readonly accounts!: pulumi.Output<outputs.aad.v20170601.ContainerAccountResponse[] | undefined>;
    /**
     * The OuContainer name
     */
    public /*out*/ readonly containerId!: pulumi.Output<string>;
    /**
     * The Deployment id
     */
    public /*out*/ readonly deploymentId!: pulumi.Output<string>;
    /**
     * The domain name of Domain Services.
     */
    public /*out*/ readonly domainName!: pulumi.Output<string>;
    /**
     * Resource etag
     */
    public /*out*/ readonly etag!: pulumi.Output<string | undefined>;
    /**
     * Resource location
     */
    public /*out*/ readonly location!: pulumi.Output<string | undefined>;
    /**
     * Resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The current deployment or provisioning state, which only appears in the response.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Status of OuContainer instance
     */
    public /*out*/ readonly serviceStatus!: pulumi.Output<string>;
    /**
     * Resource tags
     */
    public /*out*/ readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Azure Active Directory tenant id
     */
    public /*out*/ readonly tenantId!: pulumi.Output<string>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a OuContainer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OuContainerArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.domainServiceName === undefined) {
                throw new Error("Missing required property 'domainServiceName'");
            }
            if (!args || args.ouContainerName === undefined) {
                throw new Error("Missing required property 'ouContainerName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["accountName"] = args ? args.accountName : undefined;
            inputs["domainServiceName"] = args ? args.domainServiceName : undefined;
            inputs["ouContainerName"] = args ? args.ouContainerName : undefined;
            inputs["password"] = args ? args.password : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["spn"] = args ? args.spn : undefined;
            inputs["accounts"] = undefined /*out*/;
            inputs["containerId"] = undefined /*out*/;
            inputs["deploymentId"] = undefined /*out*/;
            inputs["domainName"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["serviceStatus"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["tenantId"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["accounts"] = undefined /*out*/;
            inputs["containerId"] = undefined /*out*/;
            inputs["deploymentId"] = undefined /*out*/;
            inputs["domainName"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["serviceStatus"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["tenantId"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:aad/latest:OuContainer" }, { type: "azurerm:aad/v20200101:OuContainer" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(OuContainer.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a OuContainer resource.
 */
export interface OuContainerArgs {
    /**
     * The account name
     */
    readonly accountName?: pulumi.Input<string>;
    /**
     * The name of the domain service.
     */
    readonly domainServiceName: pulumi.Input<string>;
    /**
     * The name of the OuContainer.
     */
    readonly ouContainerName: pulumi.Input<string>;
    /**
     * The account password
     */
    readonly password?: pulumi.Input<string>;
    /**
     * The name of the resource group within the user's subscription. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The account spn
     */
    readonly spn?: pulumi.Input<string>;
}
