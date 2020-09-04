// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * App certificate
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
        return new Certificate(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:web/v20150801:Certificate';

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
    public readonly cerBlob!: pulumi.Output<string | undefined>;
    /**
     * Certificate expiration date
     */
    public readonly expirationDate!: pulumi.Output<string | undefined>;
    /**
     * Friendly name of the certificate
     */
    public readonly friendlyName!: pulumi.Output<string | undefined>;
    /**
     * Host names the certificate applies to
     */
    public readonly hostNames!: pulumi.Output<string[] | undefined>;
    /**
     * Specification for the hosting environment (App Service Environment) to use for the certificate
     */
    public readonly hostingEnvironmentProfile!: pulumi.Output<outputs.web.v20150801.HostingEnvironmentProfileResponse | undefined>;
    /**
     * Certificate issue Date
     */
    public readonly issueDate!: pulumi.Output<string | undefined>;
    /**
     * Certificate issuer
     */
    public readonly issuer!: pulumi.Output<string | undefined>;
    /**
     * Kind of resource
     */
    public readonly kind!: pulumi.Output<string | undefined>;
    /**
     * Resource Location
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Resource Name
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * Certificate password
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * Pfx blob
     */
    public readonly pfxBlob!: pulumi.Output<string | undefined>;
    /**
     * Public key hash
     */
    public readonly publicKeyHash!: pulumi.Output<string | undefined>;
    /**
     * Self link
     */
    public readonly selfLink!: pulumi.Output<string | undefined>;
    /**
     * App name
     */
    public readonly siteName!: pulumi.Output<string | undefined>;
    /**
     * Subject name of the certificate
     */
    public readonly subjectName!: pulumi.Output<string | undefined>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Certificate thumbprint
     */
    public readonly thumbprint!: pulumi.Output<string | undefined>;
    /**
     * Resource type
     */
    public readonly type!: pulumi.Output<string | undefined>;
    /**
     * Is the certificate valid?
     */
    public readonly valid!: pulumi.Output<boolean | undefined>;

    /**
     * Create a Certificate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CertificateArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["cerBlob"] = args ? args.cerBlob : undefined;
            inputs["expirationDate"] = args ? args.expirationDate : undefined;
            inputs["friendlyName"] = args ? args.friendlyName : undefined;
            inputs["hostNames"] = args ? args.hostNames : undefined;
            inputs["hostingEnvironmentProfile"] = args ? args.hostingEnvironmentProfile : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["issueDate"] = args ? args.issueDate : undefined;
            inputs["issuer"] = args ? args.issuer : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["password"] = args ? args.password : undefined;
            inputs["pfxBlob"] = args ? args.pfxBlob : undefined;
            inputs["publicKeyHash"] = args ? args.publicKeyHash : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["selfLink"] = args ? args.selfLink : undefined;
            inputs["siteName"] = args ? args.siteName : undefined;
            inputs["subjectName"] = args ? args.subjectName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["thumbprint"] = args ? args.thumbprint : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["valid"] = args ? args.valid : undefined;
        } else {
            inputs["cerBlob"] = undefined /*out*/;
            inputs["expirationDate"] = undefined /*out*/;
            inputs["friendlyName"] = undefined /*out*/;
            inputs["hostNames"] = undefined /*out*/;
            inputs["hostingEnvironmentProfile"] = undefined /*out*/;
            inputs["issueDate"] = undefined /*out*/;
            inputs["issuer"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["password"] = undefined /*out*/;
            inputs["pfxBlob"] = undefined /*out*/;
            inputs["publicKeyHash"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
            inputs["siteName"] = undefined /*out*/;
            inputs["subjectName"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
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
        const aliasOpts = { aliases: [{ type: "azurerm:web/latest:Certificate" }, { type: "azurerm:web/v20160301:Certificate" }, { type: "azurerm:web/v20180201:Certificate" }, { type: "azurerm:web/v20181101:Certificate" }, { type: "azurerm:web/v20190801:Certificate" }, { type: "azurerm:web/v20200601:Certificate" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Certificate.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Certificate resource.
 */
export interface CertificateArgs {
    /**
     * Raw bytes of .cer file
     */
    readonly cerBlob?: pulumi.Input<string>;
    /**
     * Certificate expiration date
     */
    readonly expirationDate?: pulumi.Input<string>;
    /**
     * Friendly name of the certificate
     */
    readonly friendlyName?: pulumi.Input<string>;
    /**
     * Host names the certificate applies to
     */
    readonly hostNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specification for the hosting environment (App Service Environment) to use for the certificate
     */
    readonly hostingEnvironmentProfile?: pulumi.Input<inputs.web.v20150801.HostingEnvironmentProfile>;
    /**
     * Resource Id
     */
    readonly id?: pulumi.Input<string>;
    /**
     * Certificate issue Date
     */
    readonly issueDate?: pulumi.Input<string>;
    /**
     * Certificate issuer
     */
    readonly issuer?: pulumi.Input<string>;
    /**
     * Kind of resource
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * Resource Location
     */
    readonly location: pulumi.Input<string>;
    /**
     * Resource Name
     */
    readonly name: pulumi.Input<string>;
    /**
     * Certificate password
     */
    readonly password?: pulumi.Input<string>;
    /**
     * Pfx blob
     */
    readonly pfxBlob?: pulumi.Input<string>;
    /**
     * Public key hash
     */
    readonly publicKeyHash?: pulumi.Input<string>;
    /**
     * Name of the resource group
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Self link
     */
    readonly selfLink?: pulumi.Input<string>;
    /**
     * App name
     */
    readonly siteName?: pulumi.Input<string>;
    /**
     * Subject name of the certificate
     */
    readonly subjectName?: pulumi.Input<string>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Certificate thumbprint
     */
    readonly thumbprint?: pulumi.Input<string>;
    /**
     * Resource type
     */
    readonly type?: pulumi.Input<string>;
    /**
     * Is the certificate valid?
     */
    readonly valid?: pulumi.Input<boolean>;
}
