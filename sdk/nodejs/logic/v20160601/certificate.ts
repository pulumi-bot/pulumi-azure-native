// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * The integration account certificate.
 *
 * ## Example Usage
 * ### Create or update a certificate
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const certificate = new azurerm.logic.v20160601.Certificate("certificate", {
 *     certificateName: "testCertificate",
 *     integrationAccountName: "testIntegrationAccount",
 *     key: {
 *         keyName: "<keyName>",
 *         keyVault: {
 *             id: "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourcegroups/testResourceGroup/providers/microsoft.keyvault/vaults/<keyVaultName>",
 *         },
 *         keyVersion: "87d9764197604449b9b8eb7bd8710868",
 *     },
 *     location: "brazilsouth",
 *     publicCertificate: "<publicCertificateValue>",
 *     resourceGroupName: "testResourceGroup",
 * });
 *
 * ```
 */
export class Certificate extends pulumi.CustomResource {
    /**
     * Get an existing Certificate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Certificate {
        return new Certificate(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:logic/v20160601:Certificate';

    /**
     * Returns true if the given object is an instance of Certificate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Certificate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Certificate.__pulumiType;
    }

    /**
     * The changed time.
     */
    public /*out*/ readonly changedTime!: pulumi.Output<string>;
    /**
     * The created time.
     */
    public /*out*/ readonly createdTime!: pulumi.Output<string>;
    /**
     * The key details in the key vault.
     */
    public readonly key!: pulumi.Output<outputs.logic.v20160601.KeyVaultKeyReferenceResponse | undefined>;
    /**
     * The resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * The metadata.
     */
    public readonly metadata!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Gets the resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The public certificate.
     */
    public readonly publicCertificate!: pulumi.Output<string | undefined>;
    /**
     * The resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Gets the resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Certificate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CertificateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CertificateArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as CertificateArgs | undefined;
            if (!args || args.certificateName === undefined) {
                throw new Error("Missing required property 'certificateName'");
            }
            if (!args || args.integrationAccountName === undefined) {
                throw new Error("Missing required property 'integrationAccountName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["certificateName"] = args ? args.certificateName : undefined;
            inputs["integrationAccountName"] = args ? args.integrationAccountName : undefined;
            inputs["key"] = args ? args.key : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["metadata"] = args ? args.metadata : undefined;
            inputs["publicCertificate"] = args ? args.publicCertificate : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["changedTime"] = undefined /*out*/;
            inputs["createdTime"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:logic/latest:Certificate" }, { type: "azurerm:logic/v20190501:Certificate" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Certificate.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Certificate resource.
 */
export interface CertificateArgs {
    /**
     * The integration account certificate name.
     */
    readonly certificateName: pulumi.Input<string>;
    /**
     * The integration account name.
     */
    readonly integrationAccountName: pulumi.Input<string>;
    /**
     * The key details in the key vault.
     */
    readonly key?: pulumi.Input<inputs.logic.v20160601.KeyVaultKeyReference>;
    /**
     * The resource location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The metadata.
     */
    readonly metadata?: pulumi.Input<{[key: string]: any}>;
    /**
     * The public certificate.
     */
    readonly publicCertificate?: pulumi.Input<string>;
    /**
     * The resource group name.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
