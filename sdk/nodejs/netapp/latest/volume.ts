// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Volume resource
 *
 * ## Example Usage
 * ### Volumes_CreateOrUpdate
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const volume = new azurerm.netapp.latest.Volume("volume", {
 *     accountName: "account1",
 *     creationToken: "my-unique-file-path",
 *     location: "eastus",
 *     poolName: "pool1",
 *     resourceGroupName: "myRG",
 *     serviceLevel: "Premium",
 *     subnetId: "/subscriptions/9760acf5-4638-11e7-9bdb-020073ca7778/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3",
 *     throughputMibps: 128,
 *     usageThreshold: 107374182400,
 *     volumeName: "volume1",
 * });
 *
 * ```
 */
export class Volume extends pulumi.CustomResource {
    /**
     * Get an existing Volume resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Volume {
        return new Volume(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:netapp/latest:Volume';

    /**
     * Returns true if the given object is an instance of Volume.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Volume {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Volume.__pulumiType;
    }

    /**
     * UUID v4 or resource identifier used to identify the Backup.
     */
    public readonly backupId!: pulumi.Output<string | undefined>;
    /**
     * Unique Baremetal Tenant Identifier.
     */
    public /*out*/ readonly baremetalTenantId!: pulumi.Output<string>;
    /**
     * A unique file path for the volume. Used when creating mount targets
     */
    public readonly creationToken!: pulumi.Output<string>;
    /**
     * DataProtection type volumes include an object containing details of the replication
     */
    public readonly dataProtection!: pulumi.Output<outputs.netapp.latest.VolumePropertiesResponseDataProtection | undefined>;
    /**
     * Set of export policy rules
     */
    public readonly exportPolicy!: pulumi.Output<outputs.netapp.latest.VolumePropertiesResponseExportPolicy | undefined>;
    /**
     * Unique FileSystem Identifier.
     */
    public /*out*/ readonly fileSystemId!: pulumi.Output<string>;
    /**
     * Restoring
     */
    public readonly isRestoring!: pulumi.Output<boolean | undefined>;
    /**
     * Describe if a volume is KerberosEnabled. To be use with swagger version 2020-05-01 or later
     */
    public readonly kerberosEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Resource location
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * List of mount targets
     */
    public readonly mountTargets!: pulumi.Output<outputs.netapp.latest.MountTargetPropertiesResponse[] | undefined>;
    /**
     * Resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Set of protocol types
     */
    public readonly protocolTypes!: pulumi.Output<string[] | undefined>;
    /**
     * Azure lifecycle management
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The security style of volume
     */
    public readonly securityStyle!: pulumi.Output<string | undefined>;
    /**
     * The service level of the file system
     */
    public readonly serviceLevel!: pulumi.Output<string | undefined>;
    /**
     * If enabled (true) the volume will contain a read-only .snapshot directory which provides access to each of the volume's snapshots (default to true).
     */
    public readonly snapshotDirectoryVisible!: pulumi.Output<boolean | undefined>;
    /**
     * UUID v4 or resource identifier used to identify the Snapshot.
     */
    public readonly snapshotId!: pulumi.Output<string | undefined>;
    /**
     * The Azure Resource URI for a delegated subnet. Must have the delegation Microsoft.NetApp/volumes
     */
    public readonly subnetId!: pulumi.Output<string>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly throughputMibps!: pulumi.Output<number | undefined>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Maximum storage quota allowed for a file system in bytes. This is a soft quota used for alerting only. Minimum size is 100 GiB. Upper limit is 100TiB. Specified in bytes.
     */
    public readonly usageThreshold!: pulumi.Output<number>;
    /**
     * What type of volume is this
     */
    public readonly volumeType!: pulumi.Output<string | undefined>;

    /**
     * Create a Volume resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VolumeArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.accountName === undefined) {
                throw new Error("Missing required property 'accountName'");
            }
            if (!args || args.creationToken === undefined) {
                throw new Error("Missing required property 'creationToken'");
            }
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.poolName === undefined) {
                throw new Error("Missing required property 'poolName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.subnetId === undefined) {
                throw new Error("Missing required property 'subnetId'");
            }
            if (!args || args.usageThreshold === undefined) {
                throw new Error("Missing required property 'usageThreshold'");
            }
            if (!args || args.volumeName === undefined) {
                throw new Error("Missing required property 'volumeName'");
            }
            inputs["accountName"] = args ? args.accountName : undefined;
            inputs["backupId"] = args ? args.backupId : undefined;
            inputs["creationToken"] = args ? args.creationToken : undefined;
            inputs["dataProtection"] = args ? args.dataProtection : undefined;
            inputs["exportPolicy"] = args ? args.exportPolicy : undefined;
            inputs["isRestoring"] = args ? args.isRestoring : undefined;
            inputs["kerberosEnabled"] = args ? args.kerberosEnabled : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["mountTargets"] = args ? args.mountTargets : undefined;
            inputs["poolName"] = args ? args.poolName : undefined;
            inputs["protocolTypes"] = args ? args.protocolTypes : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["securityStyle"] = args ? args.securityStyle : undefined;
            inputs["serviceLevel"] = args ? args.serviceLevel : undefined;
            inputs["snapshotDirectoryVisible"] = args ? args.snapshotDirectoryVisible : undefined;
            inputs["snapshotId"] = args ? args.snapshotId : undefined;
            inputs["subnetId"] = args ? args.subnetId : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["throughputMibps"] = args ? args.throughputMibps : undefined;
            inputs["usageThreshold"] = args ? args.usageThreshold : undefined;
            inputs["volumeName"] = args ? args.volumeName : undefined;
            inputs["volumeType"] = args ? args.volumeType : undefined;
            inputs["baremetalTenantId"] = undefined /*out*/;
            inputs["fileSystemId"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["backupId"] = undefined /*out*/;
            inputs["baremetalTenantId"] = undefined /*out*/;
            inputs["creationToken"] = undefined /*out*/;
            inputs["dataProtection"] = undefined /*out*/;
            inputs["exportPolicy"] = undefined /*out*/;
            inputs["fileSystemId"] = undefined /*out*/;
            inputs["isRestoring"] = undefined /*out*/;
            inputs["kerberosEnabled"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["mountTargets"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["protocolTypes"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["securityStyle"] = undefined /*out*/;
            inputs["serviceLevel"] = undefined /*out*/;
            inputs["snapshotDirectoryVisible"] = undefined /*out*/;
            inputs["snapshotId"] = undefined /*out*/;
            inputs["subnetId"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["throughputMibps"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["usageThreshold"] = undefined /*out*/;
            inputs["volumeType"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:netapp/v20170815:Volume" }, { type: "azurerm:netapp/v20190501:Volume" }, { type: "azurerm:netapp/v20190601:Volume" }, { type: "azurerm:netapp/v20190701:Volume" }, { type: "azurerm:netapp/v20190801:Volume" }, { type: "azurerm:netapp/v20191001:Volume" }, { type: "azurerm:netapp/v20191101:Volume" }, { type: "azurerm:netapp/v20200201:Volume" }, { type: "azurerm:netapp/v20200601:Volume" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Volume.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Volume resource.
 */
export interface VolumeArgs {
    /**
     * The name of the NetApp account
     */
    readonly accountName: pulumi.Input<string>;
    /**
     * UUID v4 or resource identifier used to identify the Backup.
     */
    readonly backupId?: pulumi.Input<string>;
    /**
     * A unique file path for the volume. Used when creating mount targets
     */
    readonly creationToken: pulumi.Input<string>;
    /**
     * DataProtection type volumes include an object containing details of the replication
     */
    readonly dataProtection?: pulumi.Input<inputs.netapp.latest.VolumePropertiesDataProtection>;
    /**
     * Set of export policy rules
     */
    readonly exportPolicy?: pulumi.Input<inputs.netapp.latest.VolumePropertiesExportPolicy>;
    /**
     * Restoring
     */
    readonly isRestoring?: pulumi.Input<boolean>;
    /**
     * Describe if a volume is KerberosEnabled. To be use with swagger version 2020-05-01 or later
     */
    readonly kerberosEnabled?: pulumi.Input<boolean>;
    /**
     * Resource location
     */
    readonly location: pulumi.Input<string>;
    /**
     * List of mount targets
     */
    readonly mountTargets?: pulumi.Input<pulumi.Input<inputs.netapp.latest.MountTargetProperties>[]>;
    /**
     * The name of the capacity pool
     */
    readonly poolName: pulumi.Input<string>;
    /**
     * Set of protocol types
     */
    readonly protocolTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The security style of volume
     */
    readonly securityStyle?: pulumi.Input<string>;
    /**
     * The service level of the file system
     */
    readonly serviceLevel?: pulumi.Input<string>;
    /**
     * If enabled (true) the volume will contain a read-only .snapshot directory which provides access to each of the volume's snapshots (default to true).
     */
    readonly snapshotDirectoryVisible?: pulumi.Input<boolean>;
    /**
     * UUID v4 or resource identifier used to identify the Snapshot.
     */
    readonly snapshotId?: pulumi.Input<string>;
    /**
     * The Azure Resource URI for a delegated subnet. Must have the delegation Microsoft.NetApp/volumes
     */
    readonly subnetId: pulumi.Input<string>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly throughputMibps?: pulumi.Input<number>;
    /**
     * Maximum storage quota allowed for a file system in bytes. This is a soft quota used for alerting only. Minimum size is 100 GiB. Upper limit is 100TiB. Specified in bytes.
     */
    readonly usageThreshold: pulumi.Input<number>;
    /**
     * The name of the volume
     */
    readonly volumeName: pulumi.Input<string>;
    /**
     * What type of volume is this
     */
    readonly volumeType?: pulumi.Input<string>;
}
