// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * The profile resource format.
 *
 * ## Example Usage
 * ### Profiles_CreateOrUpdate
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const profile = new azurerm.customerinsights.latest.Profile("profile", {
 *     apiEntitySetName: "TestProfileType396",
 *     fields: [
 *         {
 *             fieldName: "Id",
 *             fieldType: "Edm.String",
 *             isArray: false,
 *             isRequired: true,
 *         },
 *         {
 *             fieldName: "ProfileId",
 *             fieldType: "Edm.String",
 *             isArray: false,
 *             isRequired: true,
 *         },
 *         {
 *             fieldName: "LastName",
 *             fieldType: "Edm.String",
 *             isArray: false,
 *             isRequired: true,
 *         },
 *         {
 *             fieldName: "TestProfileType396",
 *             fieldType: "Edm.String",
 *             isArray: false,
 *             isRequired: true,
 *         },
 *         {
 *             fieldName: "SavingAccountBalance",
 *             fieldType: "Edm.Int32",
 *             isArray: false,
 *             isRequired: true,
 *         },
 *     ],
 *     hubName: "sdkTestHub",
 *     largeImage: "\\\\Images\\\\LargeImage",
 *     mediumImage: "\\\\Images\\\\MediumImage",
 *     profileName: "TestProfileType396",
 *     resourceGroupName: "TestHubRG",
 *     schemaItemTypeLink: "SchemaItemTypeLink",
 *     smallImage: "\\\\Images\\\\smallImage",
 *     strongIds: [
 *         {
 *             keyPropertyNames: [
 *                 "Id",
 *                 "SavingAccountBalance",
 *             ],
 *             strongIdName: "Id",
 *         },
 *         {
 *             keyPropertyNames: [
 *                 "ProfileId",
 *                 "LastName",
 *             ],
 *             strongIdName: "ProfileId",
 *         },
 *     ],
 * });
 *
 * ```
 */
export class Profile extends pulumi.CustomResource {
    /**
     * Get an existing Profile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Profile {
        return new Profile(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:customerinsights/latest:Profile';

    /**
     * Returns true if the given object is an instance of Profile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Profile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Profile.__pulumiType;
    }

    /**
     * The api entity set name. This becomes the odata entity set name for the entity Type being referred in this object.
     */
    public readonly apiEntitySetName!: pulumi.Output<string | undefined>;
    /**
     * The attributes for the Type.
     */
    public readonly attributes!: pulumi.Output<{[key: string]: string[]} | undefined>;
    /**
     * Localized descriptions for the property.
     */
    public readonly description!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Localized display names for the property.
     */
    public readonly displayName!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Type of entity.
     */
    public readonly entityType!: pulumi.Output<string | undefined>;
    /**
     * The properties of the Profile.
     */
    public readonly fields!: pulumi.Output<outputs.customerinsights.latest.PropertyDefinitionResponse[] | undefined>;
    /**
     * The instance count.
     */
    public readonly instancesCount!: pulumi.Output<number | undefined>;
    /**
     * Large Image associated with the Property or EntityType.
     */
    public readonly largeImage!: pulumi.Output<string | undefined>;
    /**
     * The last changed time for the type definition.
     */
    public /*out*/ readonly lastChangedUtc!: pulumi.Output<string>;
    /**
     * Any custom localized attributes for the Type.
     */
    public readonly localizedAttributes!: pulumi.Output<{[key: string]: {[key: string]: string}} | undefined>;
    /**
     * Medium Image associated with the Property or EntityType.
     */
    public readonly mediumImage!: pulumi.Output<string | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Provisioning state.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The schema org link. This helps ACI identify and suggest semantic models.
     */
    public readonly schemaItemTypeLink!: pulumi.Output<string | undefined>;
    /**
     * Small Image associated with the Property or EntityType.
     */
    public readonly smallImage!: pulumi.Output<string | undefined>;
    /**
     * The strong IDs.
     */
    public readonly strongIds!: pulumi.Output<outputs.customerinsights.latest.StrongIdResponse[] | undefined>;
    /**
     * The hub name.
     */
    public /*out*/ readonly tenantId!: pulumi.Output<string>;
    /**
     * The timestamp property name. Represents the time when the interaction or profile update happened.
     */
    public readonly timestampFieldName!: pulumi.Output<string | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The name of the entity.
     */
    public readonly typeName!: pulumi.Output<string | undefined>;

    /**
     * Create a Profile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProfileArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.hubName === undefined) {
                throw new Error("Missing required property 'hubName'");
            }
            if (!args || args.profileName === undefined) {
                throw new Error("Missing required property 'profileName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["apiEntitySetName"] = args ? args.apiEntitySetName : undefined;
            inputs["attributes"] = args ? args.attributes : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["entityType"] = args ? args.entityType : undefined;
            inputs["fields"] = args ? args.fields : undefined;
            inputs["hubName"] = args ? args.hubName : undefined;
            inputs["instancesCount"] = args ? args.instancesCount : undefined;
            inputs["largeImage"] = args ? args.largeImage : undefined;
            inputs["localizedAttributes"] = args ? args.localizedAttributes : undefined;
            inputs["mediumImage"] = args ? args.mediumImage : undefined;
            inputs["profileName"] = args ? args.profileName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["schemaItemTypeLink"] = args ? args.schemaItemTypeLink : undefined;
            inputs["smallImage"] = args ? args.smallImage : undefined;
            inputs["strongIds"] = args ? args.strongIds : undefined;
            inputs["timestampFieldName"] = args ? args.timestampFieldName : undefined;
            inputs["typeName"] = args ? args.typeName : undefined;
            inputs["lastChangedUtc"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["tenantId"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["apiEntitySetName"] = undefined /*out*/;
            inputs["attributes"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["entityType"] = undefined /*out*/;
            inputs["fields"] = undefined /*out*/;
            inputs["instancesCount"] = undefined /*out*/;
            inputs["largeImage"] = undefined /*out*/;
            inputs["lastChangedUtc"] = undefined /*out*/;
            inputs["localizedAttributes"] = undefined /*out*/;
            inputs["mediumImage"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["schemaItemTypeLink"] = undefined /*out*/;
            inputs["smallImage"] = undefined /*out*/;
            inputs["strongIds"] = undefined /*out*/;
            inputs["tenantId"] = undefined /*out*/;
            inputs["timestampFieldName"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["typeName"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:customerinsights/v20170101:Profile" }, { type: "azurerm:customerinsights/v20170426:Profile" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Profile.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Profile resource.
 */
export interface ProfileArgs {
    /**
     * The api entity set name. This becomes the odata entity set name for the entity Type being referred in this object.
     */
    readonly apiEntitySetName?: pulumi.Input<string>;
    /**
     * The attributes for the Type.
     */
    readonly attributes?: pulumi.Input<{[key: string]: pulumi.Input<pulumi.Input<string>[]>}>;
    /**
     * Localized descriptions for the property.
     */
    readonly description?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Localized display names for the property.
     */
    readonly displayName?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Type of entity.
     */
    readonly entityType?: pulumi.Input<string>;
    /**
     * The properties of the Profile.
     */
    readonly fields?: pulumi.Input<pulumi.Input<inputs.customerinsights.latest.PropertyDefinition>[]>;
    /**
     * The name of the hub.
     */
    readonly hubName: pulumi.Input<string>;
    /**
     * The instance count.
     */
    readonly instancesCount?: pulumi.Input<number>;
    /**
     * Large Image associated with the Property or EntityType.
     */
    readonly largeImage?: pulumi.Input<string>;
    /**
     * Any custom localized attributes for the Type.
     */
    readonly localizedAttributes?: pulumi.Input<{[key: string]: pulumi.Input<{[key: string]: pulumi.Input<string>}>}>;
    /**
     * Medium Image associated with the Property or EntityType.
     */
    readonly mediumImage?: pulumi.Input<string>;
    /**
     * The name of the profile.
     */
    readonly profileName: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The schema org link. This helps ACI identify and suggest semantic models.
     */
    readonly schemaItemTypeLink?: pulumi.Input<string>;
    /**
     * Small Image associated with the Property or EntityType.
     */
    readonly smallImage?: pulumi.Input<string>;
    /**
     * The strong IDs.
     */
    readonly strongIds?: pulumi.Input<pulumi.Input<inputs.customerinsights.latest.StrongId>[]>;
    /**
     * The timestamp property name. Represents the time when the interaction or profile update happened.
     */
    readonly timestampFieldName?: pulumi.Input<string>;
    /**
     * The name of the entity.
     */
    readonly typeName?: pulumi.Input<string>;
}
