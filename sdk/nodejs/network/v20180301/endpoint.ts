// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Class representing a Traffic Manager endpoint.
 *
 * ## Example Usage
 * ### Endpoint-PUT-External-WithCustomHeaders
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const endpoint = new azurerm.network.v20180301.Endpoint("endpoint", {
 *     customHeaders: [
 *         {
 *             name: "header-1",
 *             value: "value-1",
 *         },
 *         {
 *             name: "header-2",
 *             value: "value-2",
 *         },
 *     ],
 *     endpointLocation: "North Europe",
 *     endpointName: "azsmnet7187",
 *     endpointStatus: "Enabled",
 *     endpointType: "ExternalEndpoints",
 *     name: "azsmnet7187",
 *     profileName: "azsmnet6386",
 *     resourceGroupName: "azuresdkfornetautoresttrafficmanager1421",
 *     target: "foobar.contoso.com",
 *     type: "Microsoft.network/TrafficManagerProfiles/ExternalEndpoints",
 * });
 *
 * ```
 * ### Endpoint-PUT-External-WithGeoMapping
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const endpoint = new azurerm.network.v20180301.Endpoint("endpoint", {
 *     endpointName: `My%20external%20endpoint`,
 *     endpointStatus: "Enabled",
 *     endpointType: "ExternalEndpoints",
 *     geoMapping: [
 *         "GEO-AS",
 *         "GEO-AF",
 *     ],
 *     name: "My external endpoint",
 *     profileName: "azuresdkfornetautoresttrafficmanager8224",
 *     resourceGroupName: "azuresdkfornetautoresttrafficmanager2191",
 *     target: "foobar.contoso.com",
 *     type: "Microsoft.network/TrafficManagerProfiles/ExternalEndpoints",
 * });
 *
 * ```
 * ### Endpoint-PUT-External-WithLocation
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const endpoint = new azurerm.network.v20180301.Endpoint("endpoint", {
 *     endpointLocation: "North Europe",
 *     endpointName: "azsmnet7187",
 *     endpointStatus: "Enabled",
 *     endpointType: "ExternalEndpoints",
 *     name: "azsmnet7187",
 *     profileName: "azsmnet6386",
 *     resourceGroupName: "azuresdkfornetautoresttrafficmanager1421",
 *     target: "foobar.contoso.com",
 *     type: "Microsoft.network/TrafficManagerProfiles/ExternalEndpoints",
 * });
 *
 * ```
 */
export class Endpoint extends pulumi.CustomResource {
    /**
     * Get an existing Endpoint resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Endpoint {
        return new Endpoint(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network/v20180301:Endpoint';

    /**
     * Returns true if the given object is an instance of Endpoint.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Endpoint {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Endpoint.__pulumiType;
    }

    /**
     * List of custom headers.
     */
    public readonly customHeaders!: pulumi.Output<outputs.network.v20180301.EndpointPropertiesResponseCustomHeaders[] | undefined>;
    /**
     * Specifies the location of the external or nested endpoints when using the ‘Performance’ traffic routing method.
     */
    public readonly endpointLocation!: pulumi.Output<string | undefined>;
    /**
     * The monitoring status of the endpoint.
     */
    public readonly endpointMonitorStatus!: pulumi.Output<string | undefined>;
    /**
     * The status of the endpoint. If the endpoint is Enabled, it is probed for endpoint health and is included in the traffic routing method.
     */
    public readonly endpointStatus!: pulumi.Output<string | undefined>;
    /**
     * The list of countries/regions mapped to this endpoint when using the ‘Geographic’ traffic routing method. Please consult Traffic Manager Geographic documentation for a full list of accepted values.
     */
    public readonly geoMapping!: pulumi.Output<string[] | undefined>;
    /**
     * The minimum number of endpoints that must be available in the child profile in order for the parent profile to be considered available. Only applicable to endpoint of type 'NestedEndpoints'.
     */
    public readonly minChildEndpoints!: pulumi.Output<number | undefined>;
    /**
     * The name of the resource
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * The priority of this endpoint when using the ‘Priority’ traffic routing method. Possible values are from 1 to 1000, lower values represent higher priority. This is an optional parameter.  If specified, it must be specified on all endpoints, and no two endpoints can share the same priority value.
     */
    public readonly priority!: pulumi.Output<number | undefined>;
    /**
     * The fully-qualified DNS name or IP address of the endpoint. Traffic Manager returns this value in DNS responses to direct traffic to this endpoint.
     */
    public readonly target!: pulumi.Output<string | undefined>;
    /**
     * The Azure Resource URI of the of the endpoint. Not applicable to endpoints of type 'ExternalEndpoints'.
     */
    public readonly targetResourceId!: pulumi.Output<string | undefined>;
    /**
     * The type of the resource. Ex- Microsoft.Network/trafficManagerProfiles.
     */
    public readonly type!: pulumi.Output<string | undefined>;
    /**
     * The weight of this endpoint when using the 'Weighted' traffic routing method. Possible values are from 1 to 1000.
     */
    public readonly weight!: pulumi.Output<number | undefined>;

    /**
     * Create a Endpoint resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EndpointArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.endpointName === undefined) {
                throw new Error("Missing required property 'endpointName'");
            }
            if (!args || args.endpointType === undefined) {
                throw new Error("Missing required property 'endpointType'");
            }
            if (!args || args.profileName === undefined) {
                throw new Error("Missing required property 'profileName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["customHeaders"] = args ? args.customHeaders : undefined;
            inputs["endpointLocation"] = args ? args.endpointLocation : undefined;
            inputs["endpointMonitorStatus"] = args ? args.endpointMonitorStatus : undefined;
            inputs["endpointName"] = args ? args.endpointName : undefined;
            inputs["endpointStatus"] = args ? args.endpointStatus : undefined;
            inputs["endpointType"] = args ? args.endpointType : undefined;
            inputs["geoMapping"] = args ? args.geoMapping : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["minChildEndpoints"] = args ? args.minChildEndpoints : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["priority"] = args ? args.priority : undefined;
            inputs["profileName"] = args ? args.profileName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["target"] = args ? args.target : undefined;
            inputs["targetResourceId"] = args ? args.targetResourceId : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["weight"] = args ? args.weight : undefined;
        } else {
            inputs["customHeaders"] = undefined /*out*/;
            inputs["endpointLocation"] = undefined /*out*/;
            inputs["endpointMonitorStatus"] = undefined /*out*/;
            inputs["endpointStatus"] = undefined /*out*/;
            inputs["geoMapping"] = undefined /*out*/;
            inputs["minChildEndpoints"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["priority"] = undefined /*out*/;
            inputs["target"] = undefined /*out*/;
            inputs["targetResourceId"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["weight"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:network/latest:Endpoint" }, { type: "azurerm:network/v20151101:Endpoint" }, { type: "azurerm:network/v20170301:Endpoint" }, { type: "azurerm:network/v20170501:Endpoint" }, { type: "azurerm:network/v20180201:Endpoint" }, { type: "azurerm:network/v20180401:Endpoint" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Endpoint.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Endpoint resource.
 */
export interface EndpointArgs {
    /**
     * List of custom headers.
     */
    readonly customHeaders?: pulumi.Input<pulumi.Input<inputs.network.v20180301.EndpointPropertiesCustomHeaders>[]>;
    /**
     * Specifies the location of the external or nested endpoints when using the ‘Performance’ traffic routing method.
     */
    readonly endpointLocation?: pulumi.Input<string>;
    /**
     * The monitoring status of the endpoint.
     */
    readonly endpointMonitorStatus?: pulumi.Input<string>;
    /**
     * The name of the Traffic Manager endpoint to be created or updated.
     */
    readonly endpointName: pulumi.Input<string>;
    /**
     * The status of the endpoint. If the endpoint is Enabled, it is probed for endpoint health and is included in the traffic routing method.
     */
    readonly endpointStatus?: pulumi.Input<string>;
    /**
     * The type of the Traffic Manager endpoint to be created or updated.
     */
    readonly endpointType: pulumi.Input<string>;
    /**
     * The list of countries/regions mapped to this endpoint when using the ‘Geographic’ traffic routing method. Please consult Traffic Manager Geographic documentation for a full list of accepted values.
     */
    readonly geoMapping?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficManagerProfiles/{resourceName}
     */
    readonly id?: pulumi.Input<string>;
    /**
     * The minimum number of endpoints that must be available in the child profile in order for the parent profile to be considered available. Only applicable to endpoint of type 'NestedEndpoints'.
     */
    readonly minChildEndpoints?: pulumi.Input<number>;
    /**
     * The name of the resource
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The priority of this endpoint when using the ‘Priority’ traffic routing method. Possible values are from 1 to 1000, lower values represent higher priority. This is an optional parameter.  If specified, it must be specified on all endpoints, and no two endpoints can share the same priority value.
     */
    readonly priority?: pulumi.Input<number>;
    /**
     * The name of the Traffic Manager profile.
     */
    readonly profileName: pulumi.Input<string>;
    /**
     * The name of the resource group containing the Traffic Manager endpoint to be created or updated.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The fully-qualified DNS name or IP address of the endpoint. Traffic Manager returns this value in DNS responses to direct traffic to this endpoint.
     */
    readonly target?: pulumi.Input<string>;
    /**
     * The Azure Resource URI of the of the endpoint. Not applicable to endpoints of type 'ExternalEndpoints'.
     */
    readonly targetResourceId?: pulumi.Input<string>;
    /**
     * The type of the resource. Ex- Microsoft.Network/trafficManagerProfiles.
     */
    readonly type?: pulumi.Input<string>;
    /**
     * The weight of this endpoint when using the 'Weighted' traffic routing method. Possible values are from 1 to 1000.
     */
    readonly weight?: pulumi.Input<number>;
}
