// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * OpenId Connect Provider details.
 */
export class OpenIdConnectProvider extends pulumi.CustomResource {
    /**
     * Get an existing OpenIdConnectProvider resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): OpenIdConnectProvider {
        return new OpenIdConnectProvider(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:apimanagement/v20180601preview:OpenIdConnectProvider';

    /**
     * Returns true if the given object is an instance of OpenIdConnectProvider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OpenIdConnectProvider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OpenIdConnectProvider.__pulumiType;
    }

    /**
     * Client ID of developer console which is the client application.
     */
    public readonly clientId!: pulumi.Output<string>;
    /**
     * Client Secret of developer console which is the client application.
     */
    public readonly clientSecret!: pulumi.Output<string | undefined>;
    /**
     * User-friendly description of OpenID Connect Provider.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * User-friendly OpenID Connect Provider name.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Metadata endpoint URI.
     */
    public readonly metadataEndpoint!: pulumi.Output<string>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Resource type for API Management resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a OpenIdConnectProvider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OpenIdConnectProviderArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.clientId === undefined) {
                throw new Error("Missing required property 'clientId'");
            }
            if (!args || args.displayName === undefined) {
                throw new Error("Missing required property 'displayName'");
            }
            if (!args || args.metadataEndpoint === undefined) {
                throw new Error("Missing required property 'metadataEndpoint'");
            }
            if (!args || args.opid === undefined) {
                throw new Error("Missing required property 'opid'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.serviceName === undefined) {
                throw new Error("Missing required property 'serviceName'");
            }
            inputs["clientId"] = args ? args.clientId : undefined;
            inputs["clientSecret"] = args ? args.clientSecret : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["metadataEndpoint"] = args ? args.metadataEndpoint : undefined;
            inputs["opid"] = args ? args.opid : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["serviceName"] = args ? args.serviceName : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["clientId"] = undefined /*out*/;
            inputs["clientSecret"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["metadataEndpoint"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:apimanagement/latest:OpenIdConnectProvider" }, { type: "azure-nextgen:apimanagement/v20160707:OpenIdConnectProvider" }, { type: "azure-nextgen:apimanagement/v20161010:OpenIdConnectProvider" }, { type: "azure-nextgen:apimanagement/v20170301:OpenIdConnectProvider" }, { type: "azure-nextgen:apimanagement/v20180101:OpenIdConnectProvider" }, { type: "azure-nextgen:apimanagement/v20190101:OpenIdConnectProvider" }, { type: "azure-nextgen:apimanagement/v20191201:OpenIdConnectProvider" }, { type: "azure-nextgen:apimanagement/v20191201preview:OpenIdConnectProvider" }, { type: "azure-nextgen:apimanagement/v20200601preview:OpenIdConnectProvider" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(OpenIdConnectProvider.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a OpenIdConnectProvider resource.
 */
export interface OpenIdConnectProviderArgs {
    /**
     * Client ID of developer console which is the client application.
     */
    readonly clientId: pulumi.Input<string>;
    /**
     * Client Secret of developer console which is the client application.
     */
    readonly clientSecret?: pulumi.Input<string>;
    /**
     * User-friendly description of OpenID Connect Provider.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * User-friendly OpenID Connect Provider name.
     */
    readonly displayName: pulumi.Input<string>;
    /**
     * Metadata endpoint URI.
     */
    readonly metadataEndpoint: pulumi.Input<string>;
    /**
     * Identifier of the OpenID Connect Provider.
     */
    readonly opid: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the API Management service.
     */
    readonly serviceName: pulumi.Input<string>;
}
