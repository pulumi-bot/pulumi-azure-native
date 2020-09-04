// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * The link resource format.
 *
 * ## Example Usage
 * ### Links_CreateOrUpdate
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const link = new azurerm.customerinsights.latest.Link("link", {
 *     description: {
 *         "en-us": "Link Description",
 *     },
 *     displayName: {
 *         "en-us": "Link DisplayName",
 *     },
 *     hubName: "sdkTestHub",
 *     linkName: "linkTest4806",
 *     mappings: [{
 *         linkType: "UpdateAlways",
 *         sourcePropertyName: "testInteraction1949",
 *         targetPropertyName: "testProfile1446",
 *     }],
 *     participantPropertyReferences: [{
 *         sourcePropertyName: "testInteraction1949",
 *         targetPropertyName: "ProfileId",
 *     }],
 *     resourceGroupName: "TestHubRG",
 *     sourceEntityType: "Interaction",
 *     sourceEntityTypeName: "testInteraction1949",
 *     targetEntityType: "Profile",
 *     targetEntityTypeName: "testProfile1446",
 * });
 *
 * ```
 */
export class Link extends pulumi.CustomResource {
    /**
     * Get an existing Link resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Link {
        return new Link(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:customerinsights/latest:Link';

    /**
     * Returns true if the given object is an instance of Link.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Link {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Link.__pulumiType;
    }

    /**
     * Localized descriptions for the Link.
     */
    public readonly description!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Localized display name for the Link.
     */
    public readonly displayName!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The link name.
     */
    public readonly linkName!: pulumi.Output<string>;
    /**
     * The set of properties mappings between the source and target Types.
     */
    public readonly mappings!: pulumi.Output<outputs.customerinsights.latest.TypePropertiesMappingResponse[] | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Determines whether this link is supposed to create or delete instances if Link is NOT Reference Only.
     */
    public readonly operationType!: pulumi.Output<string | undefined>;
    /**
     * The properties that represent the participating profile.
     */
    public readonly participantPropertyReferences!: pulumi.Output<outputs.customerinsights.latest.ParticipantPropertyReferenceResponse[]>;
    /**
     * Provisioning state.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Indicating whether the link is reference only link. This flag is ignored if the Mappings are defined. If the mappings are not defined and it is set to true, links processing will not create or update profiles.
     */
    public readonly referenceOnly!: pulumi.Output<boolean | undefined>;
    /**
     * Type of source entity.
     */
    public readonly sourceEntityType!: pulumi.Output<string>;
    /**
     * Name of the source Entity Type.
     */
    public readonly sourceEntityTypeName!: pulumi.Output<string>;
    /**
     * Type of target entity.
     */
    public readonly targetEntityType!: pulumi.Output<string>;
    /**
     * Name of the target Entity Type.
     */
    public readonly targetEntityTypeName!: pulumi.Output<string>;
    /**
     * The hub name.
     */
    public /*out*/ readonly tenantId!: pulumi.Output<string>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Link resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LinkArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.hubName === undefined) {
                throw new Error("Missing required property 'hubName'");
            }
            if (!args || args.linkName === undefined) {
                throw new Error("Missing required property 'linkName'");
            }
            if (!args || args.participantPropertyReferences === undefined) {
                throw new Error("Missing required property 'participantPropertyReferences'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.sourceEntityType === undefined) {
                throw new Error("Missing required property 'sourceEntityType'");
            }
            if (!args || args.sourceEntityTypeName === undefined) {
                throw new Error("Missing required property 'sourceEntityTypeName'");
            }
            if (!args || args.targetEntityType === undefined) {
                throw new Error("Missing required property 'targetEntityType'");
            }
            if (!args || args.targetEntityTypeName === undefined) {
                throw new Error("Missing required property 'targetEntityTypeName'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["hubName"] = args ? args.hubName : undefined;
            inputs["linkName"] = args ? args.linkName : undefined;
            inputs["mappings"] = args ? args.mappings : undefined;
            inputs["operationType"] = args ? args.operationType : undefined;
            inputs["participantPropertyReferences"] = args ? args.participantPropertyReferences : undefined;
            inputs["referenceOnly"] = args ? args.referenceOnly : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sourceEntityType"] = args ? args.sourceEntityType : undefined;
            inputs["sourceEntityTypeName"] = args ? args.sourceEntityTypeName : undefined;
            inputs["targetEntityType"] = args ? args.targetEntityType : undefined;
            inputs["targetEntityTypeName"] = args ? args.targetEntityTypeName : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["tenantId"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["description"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["linkName"] = undefined /*out*/;
            inputs["mappings"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["operationType"] = undefined /*out*/;
            inputs["participantPropertyReferences"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["referenceOnly"] = undefined /*out*/;
            inputs["sourceEntityType"] = undefined /*out*/;
            inputs["sourceEntityTypeName"] = undefined /*out*/;
            inputs["targetEntityType"] = undefined /*out*/;
            inputs["targetEntityTypeName"] = undefined /*out*/;
            inputs["tenantId"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:customerinsights/v20170101:Link" }, { type: "azurerm:customerinsights/v20170426:Link" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Link.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Link resource.
 */
export interface LinkArgs {
    /**
     * Localized descriptions for the Link.
     */
    readonly description?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Localized display name for the Link.
     */
    readonly displayName?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the hub.
     */
    readonly hubName: pulumi.Input<string>;
    /**
     * The name of the link.
     */
    readonly linkName: pulumi.Input<string>;
    /**
     * The set of properties mappings between the source and target Types.
     */
    readonly mappings?: pulumi.Input<pulumi.Input<inputs.customerinsights.latest.TypePropertiesMapping>[]>;
    /**
     * Determines whether this link is supposed to create or delete instances if Link is NOT Reference Only.
     */
    readonly operationType?: pulumi.Input<string>;
    /**
     * The properties that represent the participating profile.
     */
    readonly participantPropertyReferences: pulumi.Input<pulumi.Input<inputs.customerinsights.latest.ParticipantPropertyReference>[]>;
    /**
     * Indicating whether the link is reference only link. This flag is ignored if the Mappings are defined. If the mappings are not defined and it is set to true, links processing will not create or update profiles.
     */
    readonly referenceOnly?: pulumi.Input<boolean>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Type of source entity.
     */
    readonly sourceEntityType: pulumi.Input<string>;
    /**
     * Name of the source Entity Type.
     */
    readonly sourceEntityTypeName: pulumi.Input<string>;
    /**
     * Type of target entity.
     */
    readonly targetEntityType: pulumi.Input<string>;
    /**
     * Name of the target Entity Type.
     */
    readonly targetEntityTypeName: pulumi.Input<string>;
}
