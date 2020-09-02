// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * A hostname binding object.
 */
export class WebAppHostNameBinding extends pulumi.CustomResource {
    /**
     * Get an existing WebAppHostNameBinding resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): WebAppHostNameBinding {
        return new WebAppHostNameBinding(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:web/v20200601:WebAppHostNameBinding';

    /**
     * Returns true if the given object is an instance of WebAppHostNameBinding.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WebAppHostNameBinding {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WebAppHostNameBinding.__pulumiType;
    }

    /**
     * Azure resource name.
     */
    public readonly azureResourceName!: pulumi.Output<string | undefined>;
    /**
     * Azure resource type.
     */
    public readonly azureResourceType!: pulumi.Output<string | undefined>;
    /**
     * Custom DNS record type.
     */
    public readonly customHostNameDnsRecordType!: pulumi.Output<string | undefined>;
    /**
     * Fully qualified ARM domain resource URI.
     */
    public readonly domainId!: pulumi.Output<string | undefined>;
    /**
     * Hostname type.
     */
    public readonly hostNameType!: pulumi.Output<string | undefined>;
    /**
     * Kind of resource.
     */
    public readonly kind!: pulumi.Output<string | undefined>;
    /**
     * Resource Name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * App Service app name.
     */
    public readonly siteName!: pulumi.Output<string | undefined>;
    /**
     * SSL type
     */
    public readonly sslState!: pulumi.Output<string | undefined>;
    /**
     * SSL certificate thumbprint
     */
    public readonly thumbprint!: pulumi.Output<string | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Virtual IP address assigned to the hostname if IP based SSL is enabled.
     */
    public /*out*/ readonly virtualIP!: pulumi.Output<string>;

    /**
     * Create a WebAppHostNameBinding resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WebAppHostNameBindingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WebAppHostNameBindingArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as WebAppHostNameBindingArgs | undefined;
            if (!args || args.hostName === undefined) {
                throw new Error("Missing required property 'hostName'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["azureResourceName"] = args ? args.azureResourceName : undefined;
            inputs["azureResourceType"] = args ? args.azureResourceType : undefined;
            inputs["customHostNameDnsRecordType"] = args ? args.customHostNameDnsRecordType : undefined;
            inputs["domainId"] = args ? args.domainId : undefined;
            inputs["hostName"] = args ? args.hostName : undefined;
            inputs["hostNameType"] = args ? args.hostNameType : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["siteName"] = args ? args.siteName : undefined;
            inputs["sslState"] = args ? args.sslState : undefined;
            inputs["thumbprint"] = args ? args.thumbprint : undefined;
            inputs["type"] = undefined /*out*/;
            inputs["virtualIP"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:web/latest:WebAppHostNameBinding" }, { type: "azurerm:web/v20150801:WebAppHostNameBinding" }, { type: "azurerm:web/v20160801:WebAppHostNameBinding" }, { type: "azurerm:web/v20180201:WebAppHostNameBinding" }, { type: "azurerm:web/v20181101:WebAppHostNameBinding" }, { type: "azurerm:web/v20190801:WebAppHostNameBinding" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(WebAppHostNameBinding.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a WebAppHostNameBinding resource.
 */
export interface WebAppHostNameBindingArgs {
    /**
     * Azure resource name.
     */
    readonly azureResourceName?: pulumi.Input<string>;
    /**
     * Azure resource type.
     */
    readonly azureResourceType?: pulumi.Input<string>;
    /**
     * Custom DNS record type.
     */
    readonly customHostNameDnsRecordType?: pulumi.Input<string>;
    /**
     * Fully qualified ARM domain resource URI.
     */
    readonly domainId?: pulumi.Input<string>;
    /**
     * Hostname in the hostname binding.
     */
    readonly hostName: pulumi.Input<string>;
    /**
     * Hostname type.
     */
    readonly hostNameType?: pulumi.Input<string>;
    /**
     * Kind of resource.
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * Name of the app.
     */
    readonly name: pulumi.Input<string>;
    /**
     * Name of the resource group to which the resource belongs.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * App Service app name.
     */
    readonly siteName?: pulumi.Input<string>;
    /**
     * SSL type
     */
    readonly sslState?: pulumi.Input<string>;
    /**
     * SSL certificate thumbprint
     */
    readonly thumbprint?: pulumi.Input<string>;
}