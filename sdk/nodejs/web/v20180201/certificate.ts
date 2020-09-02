// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * SSL certificate for an app.
 *
 * ## Example Usage
 * ### Create Or Update Certificate
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const certificate = new azurerm.web.v20180201.Certificate("certificate", {
 *     hostNames: ["ServerCert"],
 *     location: "East US",
 *     name: "testc6282",
 *     password: `SWsSsd__233$Sdsds#%Sd!`,
 *     resourceGroupName: "testrg123",
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
    public static readonly __pulumiType = 'azurerm:web/v20180201:Certificate';

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
     * Raw bytes of .cer file
     */
    public /*out*/ readonly cerBlob!: pulumi.Output<string>;
    /**
     * Certificate expiration date.
     */
    public /*out*/ readonly expirationDate!: pulumi.Output<string>;
    /**
     * Friendly name of the certificate.
     */
    public /*out*/ readonly friendlyName!: pulumi.Output<string>;
    /**
     * Host names the certificate applies to.
     */
    public readonly hostNames!: pulumi.Output<string[] | undefined>;
    /**
     * Specification for the App Service Environment to use for the certificate.
     */
    public /*out*/ readonly hostingEnvironmentProfile!: pulumi.Output<outputs.web.v20180201.HostingEnvironmentProfileResponse>;
    /**
     * Certificate issue Date.
     */
    public /*out*/ readonly issueDate!: pulumi.Output<string>;
    /**
     * Certificate issuer.
     */
    public /*out*/ readonly issuer!: pulumi.Output<string>;
    /**
     * Key Vault Csm resource Id.
     */
    public readonly keyVaultId!: pulumi.Output<string | undefined>;
    /**
     * Key Vault secret name.
     */
    public readonly keyVaultSecretName!: pulumi.Output<string | undefined>;
    /**
     * Status of the Key Vault secret.
     */
    public /*out*/ readonly keyVaultSecretStatus!: pulumi.Output<string>;
    /**
     * Kind of resource.
     */
    public readonly kind!: pulumi.Output<string | undefined>;
    /**
     * Resource Location.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Resource Name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Certificate password.
     */
    public readonly password!: pulumi.Output<string>;
    /**
     * Pfx blob.
     */
    public readonly pfxBlob!: pulumi.Output<string | undefined>;
    /**
     * Public key hash.
     */
    public /*out*/ readonly publicKeyHash!: pulumi.Output<string>;
    /**
     * Self link.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * Resource ID of the associated App Service plan, formatted as: "/subscriptions/{subscriptionID}/resourceGroups/{groupName}/providers/Microsoft.Web/serverfarms/{appServicePlanName}".
     */
    public readonly serverFarmId!: pulumi.Output<string | undefined>;
    /**
     * App name.
     */
    public /*out*/ readonly siteName!: pulumi.Output<string>;
    /**
     * Subject name of the certificate.
     */
    public /*out*/ readonly subjectName!: pulumi.Output<string>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Certificate thumbprint.
     */
    public /*out*/ readonly thumbprint!: pulumi.Output<string>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Is the certificate valid?.
     */
    public /*out*/ readonly valid!: pulumi.Output<boolean>;

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
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.password === undefined) {
                throw new Error("Missing required property 'password'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["hostNames"] = args ? args.hostNames : undefined;
            inputs["keyVaultId"] = args ? args.keyVaultId : undefined;
            inputs["keyVaultSecretName"] = args ? args.keyVaultSecretName : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["password"] = args ? args.password : undefined;
            inputs["pfxBlob"] = args ? args.pfxBlob : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["serverFarmId"] = args ? args.serverFarmId : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["cerBlob"] = undefined /*out*/;
            inputs["expirationDate"] = undefined /*out*/;
            inputs["friendlyName"] = undefined /*out*/;
            inputs["hostingEnvironmentProfile"] = undefined /*out*/;
            inputs["issueDate"] = undefined /*out*/;
            inputs["issuer"] = undefined /*out*/;
            inputs["keyVaultSecretStatus"] = undefined /*out*/;
            inputs["publicKeyHash"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
            inputs["siteName"] = undefined /*out*/;
            inputs["subjectName"] = undefined /*out*/;
            inputs["thumbprint"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["valid"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:web/latest:Certificate" }, { type: "azurerm:web/v20150801:Certificate" }, { type: "azurerm:web/v20160301:Certificate" }, { type: "azurerm:web/v20181101:Certificate" }, { type: "azurerm:web/v20190801:Certificate" }, { type: "azurerm:web/v20200601:Certificate" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Certificate.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Certificate resource.
 */
export interface CertificateArgs {
    /**
     * Host names the certificate applies to.
     */
    readonly hostNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Key Vault Csm resource Id.
     */
    readonly keyVaultId?: pulumi.Input<string>;
    /**
     * Key Vault secret name.
     */
    readonly keyVaultSecretName?: pulumi.Input<string>;
    /**
     * Kind of resource.
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * Resource Location.
     */
    readonly location: pulumi.Input<string>;
    /**
     * Name of the certificate.
     */
    readonly name: pulumi.Input<string>;
    /**
     * Certificate password.
     */
    readonly password: pulumi.Input<string>;
    /**
     * Pfx blob.
     */
    readonly pfxBlob?: pulumi.Input<string>;
    /**
     * Name of the resource group to which the resource belongs.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Resource ID of the associated App Service plan, formatted as: "/subscriptions/{subscriptionID}/resourceGroups/{groupName}/providers/Microsoft.Web/serverfarms/{appServicePlanName}".
     */
    readonly serverFarmId?: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
