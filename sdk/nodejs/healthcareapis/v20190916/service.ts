// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * The description of the service.
 *
 * ## Example Usage
 * ### Create or Update a service with all parameters
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const service = new azurerm.healthcareapis.v20190916.Service("service", {
 *     identity: {
 *         type: "SystemAssigned",
 *     },
 *     kind: "fhir-R4",
 *     location: "westus2",
 *     resourceGroupName: "rg1",
 *     resourceName: "service1",
 *     tags: {},
 * });
 *
 * ```
 * ### Create or Update a service with minimum parameters
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const service = new azurerm.healthcareapis.v20190916.Service("service", {
 *     kind: "fhir-R4",
 *     location: "westus2",
 *     resourceGroupName: "rg1",
 *     resourceName: "service2",
 *     tags: {},
 * });
 *
 * ```
 */
export class Service extends pulumi.CustomResource {
    /**
     * Get an existing Service resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Service {
        return new Service(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:healthcareapis/v20190916:Service';

    /**
     * Returns true if the given object is an instance of Service.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Service {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Service.__pulumiType;
    }

    /**
     * An etag associated with the resource, used for optimistic concurrency when editing it.
     */
    public readonly etag!: pulumi.Output<string | undefined>;
    /**
     * Setting indicating whether the service has a managed identity associated with it.
     */
    public readonly identity!: pulumi.Output<outputs.healthcareapis.v20190916.ResourceResponseIdentity | undefined>;
    /**
     * The kind of the service.
     */
    public readonly kind!: pulumi.Output<string>;
    /**
     * The resource location.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The common properties of a service.
     */
    public readonly properties!: pulumi.Output<outputs.healthcareapis.v20190916.ServicesPropertiesResponse>;
    /**
     * The resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Service resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.kind === undefined) {
                throw new Error("Missing required property 'kind'");
            }
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.resourceName === undefined) {
                throw new Error("Missing required property 'resourceName'");
            }
            inputs["etag"] = args ? args.etag : undefined;
            inputs["identity"] = args ? args.identity : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["resourceName"] = args ? args.resourceName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["etag"] = undefined /*out*/;
            inputs["identity"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["properties"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:healthcareapis/latest:Service" }, { type: "azurerm:healthcareapis/v20180820preview:Service" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Service.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Service resource.
 */
export interface ServiceArgs {
    /**
     * An etag associated with the resource, used for optimistic concurrency when editing it.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * Setting indicating whether the service has a managed identity associated with it.
     */
    readonly identity?: pulumi.Input<inputs.healthcareapis.v20190916.ResourceIdentity>;
    /**
     * The kind of the service.
     */
    readonly kind: pulumi.Input<string>;
    /**
     * The resource location.
     */
    readonly location: pulumi.Input<string>;
    /**
     * The common properties of a service.
     */
    readonly properties?: pulumi.Input<inputs.healthcareapis.v20190916.ServicesProperties>;
    /**
     * The name of the resource group that contains the service instance.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the service instance.
     */
    readonly resourceName: pulumi.Input<string>;
    /**
     * The resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
