// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * The X509 Certificate.
 */
export class DpsCertificate extends pulumi.CustomResource {
    /**
     * Get an existing DpsCertificate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DpsCertificate {
        return new DpsCertificate(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:devices/v20200901preview:DpsCertificate';

    /**
     * Returns true if the given object is an instance of DpsCertificate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DpsCertificate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DpsCertificate.__pulumiType;
    }

    /**
     * The entity tag.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The name of the certificate.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * properties of a certificate
     */
    public /*out*/ readonly properties!: pulumi.Output<outputs.devices.v20200901preview.CertificatePropertiesResponse>;
    /**
     * The resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a DpsCertificate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DpsCertificateArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.certificateName === undefined) {
                throw new Error("Missing required property 'certificateName'");
            }
            if (!args || args.provisioningServiceName === undefined) {
                throw new Error("Missing required property 'provisioningServiceName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["certificate"] = args ? args.certificate : undefined;
            inputs["certificateName"] = args ? args.certificateName : undefined;
            inputs["provisioningServiceName"] = args ? args.provisioningServiceName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["etag"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["properties"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["etag"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["properties"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:devices/latest:DpsCertificate" }, { type: "azurerm:devices/v20170821preview:DpsCertificate" }, { type: "azurerm:devices/v20171115:DpsCertificate" }, { type: "azurerm:devices/v20180122:DpsCertificate" }, { type: "azurerm:devices/v20200101:DpsCertificate" }, { type: "azurerm:devices/v20200301:DpsCertificate" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(DpsCertificate.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a DpsCertificate resource.
 */
export interface DpsCertificateArgs {
    /**
     * Base-64 representation of the X509 leaf certificate .cer file or just .pem file content.
     */
    readonly certificate?: pulumi.Input<string>;
    /**
     * The name of the certificate create or update.
     */
    readonly certificateName: pulumi.Input<string>;
    /**
     * The name of the provisioning service.
     */
    readonly provisioningServiceName: pulumi.Input<string>;
    /**
     * Resource group identifier.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
