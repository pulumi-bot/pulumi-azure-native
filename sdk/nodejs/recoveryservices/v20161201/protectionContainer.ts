// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Base class for container with backup items. Containers with specific workloads are derived from this class.
 */
export class ProtectionContainer extends pulumi.CustomResource {
    /**
     * Get an existing ProtectionContainer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ProtectionContainer {
        return new ProtectionContainer(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:recoveryservices/v20161201:ProtectionContainer';

    /**
     * Returns true if the given object is an instance of ProtectionContainer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProtectionContainer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProtectionContainer.__pulumiType;
    }

    /**
     * Optional ETag.
     */
    public readonly eTag!: pulumi.Output<string | undefined>;
    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Resource name associated with the resource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * ProtectionContainerResource properties
     */
    public readonly properties!: pulumi.Output<outputs.recoveryservices.v20161201.AzureSqlContainerResponse | outputs.recoveryservices.v20161201.AzureStorageContainerResponse | outputs.recoveryservices.v20161201.AzureWorkloadContainerResponse | outputs.recoveryservices.v20161201.DpmContainerResponse | outputs.recoveryservices.v20161201.GenericContainerResponse | outputs.recoveryservices.v20161201.IaaSVMContainerResponse | outputs.recoveryservices.v20161201.MabContainerResponse>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ProtectionContainer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProtectionContainerArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if ((!args || args.containerName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'containerName'");
            }
            if ((!args || args.fabricName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'fabricName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.vaultName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'vaultName'");
            }
            inputs["containerName"] = args ? args.containerName : undefined;
            inputs["eTag"] = args ? args.eTag : undefined;
            inputs["fabricName"] = args ? args.fabricName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["vaultName"] = args ? args.vaultName : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["eTag"] = undefined /*out*/;
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
        const aliasOpts = { aliases: [{ type: "azure-nextgen:recoveryservices/latest:ProtectionContainer" }, { type: "azure-nextgen:recoveryservices/v20201001:ProtectionContainer" }, { type: "azure-nextgen:recoveryservices/v20201201:ProtectionContainer" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(ProtectionContainer.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ProtectionContainer resource.
 */
export interface ProtectionContainerArgs {
    /**
     * Name of the container to be registered.
     */
    readonly containerName: pulumi.Input<string>;
    /**
     * Optional ETag.
     */
    readonly eTag?: pulumi.Input<string>;
    /**
     * Fabric name associated with the container.
     */
    readonly fabricName: pulumi.Input<string>;
    /**
     * Resource location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * ProtectionContainerResource properties
     */
    readonly properties?: pulumi.Input<inputs.recoveryservices.v20161201.AzureSqlContainer | inputs.recoveryservices.v20161201.AzureStorageContainer | inputs.recoveryservices.v20161201.AzureWorkloadContainer | inputs.recoveryservices.v20161201.DpmContainer | inputs.recoveryservices.v20161201.GenericContainer | inputs.recoveryservices.v20161201.IaaSVMContainer | inputs.recoveryservices.v20161201.MabContainer>;
    /**
     * The name of the resource group where the recovery services vault is present.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the recovery services vault.
     */
    readonly vaultName: pulumi.Input<string>;
}
