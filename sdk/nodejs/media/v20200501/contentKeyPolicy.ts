// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * A Content Key Policy resource.
 *
 * ## Example Usage
 * ### Creates a Content Key Policy with ClearKey option and Token Restriction
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const contentKeyPolicy = new azurerm.media.v20200501.ContentKeyPolicy("contentKeyPolicy", {
 *     accountName: "contosomedia",
 *     contentKeyPolicyName: "PolicyWithClearKeyOptionAndSwtTokenRestriction",
 *     description: "ArmPolicyDescription",
 *     options: [{
 *         configuration: {
 *             odataType: "#Microsoft.Media.ContentKeyPolicyClearKeyConfiguration",
 *         },
 *         name: "ClearKeyOption",
 *         restriction: {
 *             odataType: "#Microsoft.Media.ContentKeyPolicyTokenRestriction",
 *         },
 *     }],
 *     resourceGroupName: "contoso",
 * });
 *
 * ```
 * ### Creates a Content Key Policy with PlayReady option and Open Restriction
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const contentKeyPolicy = new azurerm.media.v20200501.ContentKeyPolicy("contentKeyPolicy", {
 *     accountName: "contosomedia",
 *     contentKeyPolicyName: "PolicyWithPlayReadyOptionAndOpenRestriction",
 *     description: "ArmPolicyDescription",
 *     options: [{
 *         configuration: {
 *             odataType: "#Microsoft.Media.ContentKeyPolicyPlayReadyConfiguration",
 *         },
 *         name: "ArmPolicyOptionName",
 *         restriction: {
 *             odataType: "#Microsoft.Media.ContentKeyPolicyOpenRestriction",
 *         },
 *     }],
 *     resourceGroupName: "contoso",
 * });
 *
 * ```
 * ### Creates a Content Key Policy with Widevine option and Token Restriction
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const contentKeyPolicy = new azurerm.media.v20200501.ContentKeyPolicy("contentKeyPolicy", {
 *     accountName: "contosomedia",
 *     contentKeyPolicyName: "PolicyWithWidevineOptionAndJwtTokenRestriction",
 *     description: "ArmPolicyDescription",
 *     options: [{
 *         configuration: {
 *             odataType: "#Microsoft.Media.ContentKeyPolicyWidevineConfiguration",
 *         },
 *         name: "widevineoption",
 *         restriction: {
 *             odataType: "#Microsoft.Media.ContentKeyPolicyTokenRestriction",
 *         },
 *     }],
 *     resourceGroupName: "contoso",
 * });
 *
 * ```
 * ### Creates a Content Key Policy with multiple options
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const contentKeyPolicy = new azurerm.media.v20200501.ContentKeyPolicy("contentKeyPolicy", {
 *     accountName: "contosomedia",
 *     contentKeyPolicyName: "PolicyCreatedWithMultipleOptions",
 *     description: "ArmPolicyDescription",
 *     options: [
 *         {
 *             configuration: {
 *                 odataType: "#Microsoft.Media.ContentKeyPolicyClearKeyConfiguration",
 *             },
 *             name: "ClearKeyOption",
 *             restriction: {
 *                 odataType: "#Microsoft.Media.ContentKeyPolicyTokenRestriction",
 *             },
 *         },
 *         {
 *             configuration: {
 *                 odataType: "#Microsoft.Media.ContentKeyPolicyWidevineConfiguration",
 *             },
 *             name: "widevineoption",
 *             restriction: {
 *                 odataType: "#Microsoft.Media.ContentKeyPolicyOpenRestriction",
 *             },
 *         },
 *     ],
 *     resourceGroupName: "contoso",
 * });
 *
 * ```
 */
export class ContentKeyPolicy extends pulumi.CustomResource {
    /**
     * Get an existing ContentKeyPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ContentKeyPolicy {
        return new ContentKeyPolicy(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:media/v20200501:ContentKeyPolicy';

    /**
     * Returns true if the given object is an instance of ContentKeyPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ContentKeyPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ContentKeyPolicy.__pulumiType;
    }

    /**
     * The creation date of the Policy
     */
    public /*out*/ readonly created!: pulumi.Output<string>;
    /**
     * A description for the Policy.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The last modified date of the Policy
     */
    public /*out*/ readonly lastModified!: pulumi.Output<string>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The Key Policy options.
     */
    public readonly options!: pulumi.Output<outputs.media.v20200501.ContentKeyPolicyOptionResponse[]>;
    /**
     * The legacy Policy ID.
     */
    public /*out*/ readonly policyId!: pulumi.Output<string>;
    /**
     * The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ContentKeyPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ContentKeyPolicyArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.accountName === undefined) {
                throw new Error("Missing required property 'accountName'");
            }
            if (!args || args.contentKeyPolicyName === undefined) {
                throw new Error("Missing required property 'contentKeyPolicyName'");
            }
            if (!args || args.options === undefined) {
                throw new Error("Missing required property 'options'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["accountName"] = args ? args.accountName : undefined;
            inputs["contentKeyPolicyName"] = args ? args.contentKeyPolicyName : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["options"] = args ? args.options : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["created"] = undefined /*out*/;
            inputs["lastModified"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["policyId"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["created"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["lastModified"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["options"] = undefined /*out*/;
            inputs["policyId"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:media/latest:ContentKeyPolicy" }, { type: "azurerm:media/v20180330preview:ContentKeyPolicy" }, { type: "azurerm:media/v20180601preview:ContentKeyPolicy" }, { type: "azurerm:media/v20180701:ContentKeyPolicy" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(ContentKeyPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ContentKeyPolicy resource.
 */
export interface ContentKeyPolicyArgs {
    /**
     * The Media Services account name.
     */
    readonly accountName: pulumi.Input<string>;
    /**
     * The Content Key Policy name.
     */
    readonly contentKeyPolicyName: pulumi.Input<string>;
    /**
     * A description for the Policy.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The Key Policy options.
     */
    readonly options: pulumi.Input<pulumi.Input<inputs.media.v20200501.ContentKeyPolicyOption>[]>;
    /**
     * The name of the resource group within the Azure subscription.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
