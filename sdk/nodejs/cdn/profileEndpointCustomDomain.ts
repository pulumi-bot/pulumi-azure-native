// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Friendly domain name mapping to the endpoint hostname that the customer provides for branding purposes, e.g. www.contoso.com.
 */
export class ProfileEndpointCustomDomain extends pulumi.CustomResource {
    /**
     * Get an existing ProfileEndpointCustomDomain resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProfileEndpointCustomDomainState, opts?: pulumi.CustomResourceOptions): ProfileEndpointCustomDomain {
        return new ProfileEndpointCustomDomain(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:cdn:ProfileEndpointCustomDomain';

    /**
     * Returns true if the given object is an instance of ProfileEndpointCustomDomain.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProfileEndpointCustomDomain {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProfileEndpointCustomDomain.__pulumiType;
    }

    /**
     * Resource name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The JSON object that contains the properties of the custom domain to create.
     */
    public readonly properties!: pulumi.Output<outputs.cdn.CustomDomainPropertiesResponse>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ProfileEndpointCustomDomain resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProfileEndpointCustomDomainArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProfileEndpointCustomDomainArgs | ProfileEndpointCustomDomainState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ProfileEndpointCustomDomainState | undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["properties"] = state ? state.properties : undefined;
            inputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as ProfileEndpointCustomDomainArgs | undefined;
            if (!args || args.endpointName === undefined) {
                throw new Error("Missing required property 'endpointName'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.profileName === undefined) {
                throw new Error("Missing required property 'profileName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["endpointName"] = args ? args.endpointName : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["profileName"] = args ? args.profileName : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ProfileEndpointCustomDomain.__pulumiType, name, inputs, opts);
    }
}

/**
 * Friendly domain name mapping to the endpoint hostname that the customer provides for branding purposes, e.g. www.contoso.com.
 */
export interface ProfileEndpointCustomDomainState {
    /**
     * Resource name.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The JSON object that contains the properties of the custom domain to create.
     */
    readonly properties: pulumi.Input<inputs.cdn.CustomDomainPropertiesResponse>;
    /**
     * Resource type.
     */
    readonly type: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProfileEndpointCustomDomain resource.
 */
export interface ProfileEndpointCustomDomainArgs {
    /**
     * Name of the endpoint under the profile which is unique globally.
     */
    readonly endpointName: pulumi.Input<string>;
    /**
     * Name of the custom domain within an endpoint.
     */
    readonly name: pulumi.Input<string>;
    /**
     * Name of the CDN profile which is unique within the resource group.
     */
    readonly profileName: pulumi.Input<string>;
    /**
     * The JSON object that contains the properties of the custom domain to create.
     */
    readonly properties?: pulumi.Input<inputs.cdn.CustomDomainPropertiesParameters>;
    /**
     * Name of the Resource group within the Azure subscription.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
