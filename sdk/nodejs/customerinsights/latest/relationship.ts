// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * The relationship resource format.
 *
 * ## Example Usage
 * ### Relationships_CreateOrUpdate
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const relationship = new azurerm.customerinsights.latest.Relationship("relationship", {
 *     cardinality: "OneToOne",
 *     description: {
 *         "en-us": "Relationship Description",
 *     },
 *     displayName: {
 *         "en-us": "Relationship DisplayName",
 *     },
 *     fields: [],
 *     hubName: "sdkTestHub",
 *     profileType: "testProfile2326994",
 *     relatedProfileType: "testProfile2326994",
 *     relationshipName: "SomeRelationship",
 *     resourceGroupName: "TestHubRG",
 * });
 *
 * ```
 */
export class Relationship extends pulumi.CustomResource {
    /**
     * Get an existing Relationship resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Relationship {
        return new Relationship(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:customerinsights/latest:Relationship';

    /**
     * Returns true if the given object is an instance of Relationship.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Relationship {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Relationship.__pulumiType;
    }

    /**
     * The Relationship Cardinality.
     */
    public readonly cardinality!: pulumi.Output<string | undefined>;
    /**
     * Localized descriptions for the Relationship.
     */
    public readonly description!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Localized display name for the Relationship.
     */
    public readonly displayName!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The expiry date time in UTC.
     */
    public readonly expiryDateTimeUtc!: pulumi.Output<string | undefined>;
    /**
     * The properties of the Relationship.
     */
    public readonly fields!: pulumi.Output<outputs.customerinsights.latest.PropertyDefinitionResponse[] | undefined>;
    /**
     * Optional property to be used to map fields in profile to their strong ids in related profile.
     */
    public readonly lookupMappings!: pulumi.Output<outputs.customerinsights.latest.RelationshipTypeMappingResponse[] | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Profile type.
     */
    public readonly profileType!: pulumi.Output<string>;
    /**
     * Provisioning state.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Related profile being referenced.
     */
    public readonly relatedProfileType!: pulumi.Output<string>;
    /**
     * The relationship guid id.
     */
    public /*out*/ readonly relationshipGuidId!: pulumi.Output<string>;
    /**
     * The Relationship name.
     */
    public readonly relationshipName!: pulumi.Output<string>;
    /**
     * The hub name.
     */
    public /*out*/ readonly tenantId!: pulumi.Output<string>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Relationship resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RelationshipArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.hubName === undefined) {
                throw new Error("Missing required property 'hubName'");
            }
            if (!args || args.profileType === undefined) {
                throw new Error("Missing required property 'profileType'");
            }
            if (!args || args.relatedProfileType === undefined) {
                throw new Error("Missing required property 'relatedProfileType'");
            }
            if (!args || args.relationshipName === undefined) {
                throw new Error("Missing required property 'relationshipName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["cardinality"] = args ? args.cardinality : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["expiryDateTimeUtc"] = args ? args.expiryDateTimeUtc : undefined;
            inputs["fields"] = args ? args.fields : undefined;
            inputs["hubName"] = args ? args.hubName : undefined;
            inputs["lookupMappings"] = args ? args.lookupMappings : undefined;
            inputs["profileType"] = args ? args.profileType : undefined;
            inputs["relatedProfileType"] = args ? args.relatedProfileType : undefined;
            inputs["relationshipName"] = args ? args.relationshipName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["relationshipGuidId"] = undefined /*out*/;
            inputs["tenantId"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["cardinality"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["expiryDateTimeUtc"] = undefined /*out*/;
            inputs["fields"] = undefined /*out*/;
            inputs["lookupMappings"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["profileType"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["relatedProfileType"] = undefined /*out*/;
            inputs["relationshipGuidId"] = undefined /*out*/;
            inputs["relationshipName"] = undefined /*out*/;
            inputs["tenantId"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:customerinsights/v20170101:Relationship" }, { type: "azurerm:customerinsights/v20170426:Relationship" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Relationship.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Relationship resource.
 */
export interface RelationshipArgs {
    /**
     * The Relationship Cardinality.
     */
    readonly cardinality?: pulumi.Input<string>;
    /**
     * Localized descriptions for the Relationship.
     */
    readonly description?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Localized display name for the Relationship.
     */
    readonly displayName?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The expiry date time in UTC.
     */
    readonly expiryDateTimeUtc?: pulumi.Input<string>;
    /**
     * The properties of the Relationship.
     */
    readonly fields?: pulumi.Input<pulumi.Input<inputs.customerinsights.latest.PropertyDefinition>[]>;
    /**
     * The name of the hub.
     */
    readonly hubName: pulumi.Input<string>;
    /**
     * Optional property to be used to map fields in profile to their strong ids in related profile.
     */
    readonly lookupMappings?: pulumi.Input<pulumi.Input<inputs.customerinsights.latest.RelationshipTypeMapping>[]>;
    /**
     * Profile type.
     */
    readonly profileType: pulumi.Input<string>;
    /**
     * Related profile being referenced.
     */
    readonly relatedProfileType: pulumi.Input<string>;
    /**
     * The name of the Relationship.
     */
    readonly relationshipName: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
