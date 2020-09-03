// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * The X509 Certificate.
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
    public static readonly __pulumiType = 'azurerm:devices/v20170701:Certificate';

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
     * The entity tag.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The name of the certificate.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The description of an X509 CA Certificate.
     */
    public /*out*/ readonly properties!: pulumi.Output<outputs.devices.v20170701.CertificatePropertiesResponse>;
    /**
     * The resource type.
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
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.resourceName === undefined) {
                throw new Error("Missing required property 'resourceName'");
            }
            inputs["certificate"] = args ? args.certificate : undefined;
            inputs["certificateName"] = args ? args.certificateName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["resourceName"] = args ? args.resourceName : undefined;
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
        const aliasOpts = { aliases: [{ type: "azurerm:devices/latest:Certificate" }, { type: "azurerm:devices/v20180122:Certificate" }, { type: "azurerm:devices/v20180401:Certificate" }, { type: "azurerm:devices/v20181201preview:Certificate" }, { type: "azurerm:devices/v20190322:Certificate" }, { type: "azurerm:devices/v20190322preview:Certificate" }, { type: "azurerm:devices/v20190701preview:Certificate" }, { type: "azurerm:devices/v20191104:Certificate" }, { type: "azurerm:devices/v20200301:Certificate" }, { type: "azurerm:devices/v20200401:Certificate" }, { type: "azurerm:devices/v20200615:Certificate" }, { type: "azurerm:devices/v20200710preview:Certificate" }, { type: "azurerm:devices/v20200801:Certificate" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Certificate.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Certificate resource.
 */
export interface CertificateArgs {
    /**
     * base-64 representation of the X509 leaf certificate .cer file or just .pem file content.
     */
    readonly certificate?: pulumi.Input<string>;
    /**
     * The name of the certificate
     */
    readonly certificateName: pulumi.Input<string>;
    /**
     * The name of the resource group that contains the IoT hub.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the IoT hub.
     */
    readonly resourceName: pulumi.Input<string>;
}
