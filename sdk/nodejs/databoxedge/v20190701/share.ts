// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Represents a share on the  Data Box Edge/Gateway device.
 *
 * ## Example Usage
 * ### SharePut
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const share = new azurerm.databoxedge.v20190701.Share("share", {
 *     accessProtocol: "SMB",
 *     azureContainerInfo: {
 *         containerName: "testContainerSMB",
 *         dataFormat: "BlockBlob",
 *         storageAccountCredentialId: "/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/GroupForEdgeAutomation/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/testedgedevice/storageAccountCredentials/sac1",
 *     },
 *     dataPolicy: "Cloud",
 *     description: "",
 *     deviceName: "testedgedevice",
 *     monitoringStatus: "Enabled",
 *     name: "smbshare",
 *     resourceGroupName: "GroupForEdgeAutomation",
 *     shareStatus: "Online",
 *     userAccessRights: [{
 *         accessType: "Change",
 *         userId: "/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/GroupForEdgeAutomation/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/testedgedevice/users/user2",
 *     }],
 * });
 *
 * ```
 */
export class Share extends pulumi.CustomResource {
    /**
     * Get an existing Share resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Share {
        return new Share(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:databoxedge/v20190701:Share';

    /**
     * Returns true if the given object is an instance of Share.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Share {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Share.__pulumiType;
    }

    /**
     * Access protocol to be used by the share.
     */
    public readonly accessProtocol!: pulumi.Output<string>;
    /**
     * Azure container mapping for the share.
     */
    public readonly azureContainerInfo!: pulumi.Output<outputs.databoxedge.v20190701.AzureContainerInfoResponse | undefined>;
    /**
     * List of IP addresses and corresponding access rights on the share(required for NFS protocol).
     */
    public readonly clientAccessRights!: pulumi.Output<outputs.databoxedge.v20190701.ClientAccessRightResponse[] | undefined>;
    /**
     * Data policy of the share.
     */
    public readonly dataPolicy!: pulumi.Output<string | undefined>;
    /**
     * Description for the share.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Current monitoring status of the share.
     */
    public readonly monitoringStatus!: pulumi.Output<string>;
    /**
     * The object name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Details of the refresh job on this share.
     */
    public readonly refreshDetails!: pulumi.Output<outputs.databoxedge.v20190701.RefreshDetailsResponse | undefined>;
    /**
     * Share mount point to the role.
     */
    public /*out*/ readonly shareMappings!: pulumi.Output<outputs.databoxedge.v20190701.MountPointMapResponse[]>;
    /**
     * Current status of the share.
     */
    public readonly shareStatus!: pulumi.Output<string>;
    /**
     * The hierarchical type of the object.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Mapping of users and corresponding access rights on the share (required for SMB protocol).
     */
    public readonly userAccessRights!: pulumi.Output<outputs.databoxedge.v20190701.UserAccessRightResponse[] | undefined>;

    /**
     * Create a Share resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ShareArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ShareArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as ShareArgs | undefined;
            if (!args || args.accessProtocol === undefined) {
                throw new Error("Missing required property 'accessProtocol'");
            }
            if (!args || args.deviceName === undefined) {
                throw new Error("Missing required property 'deviceName'");
            }
            if (!args || args.monitoringStatus === undefined) {
                throw new Error("Missing required property 'monitoringStatus'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.shareStatus === undefined) {
                throw new Error("Missing required property 'shareStatus'");
            }
            inputs["accessProtocol"] = args ? args.accessProtocol : undefined;
            inputs["azureContainerInfo"] = args ? args.azureContainerInfo : undefined;
            inputs["clientAccessRights"] = args ? args.clientAccessRights : undefined;
            inputs["dataPolicy"] = args ? args.dataPolicy : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["deviceName"] = args ? args.deviceName : undefined;
            inputs["monitoringStatus"] = args ? args.monitoringStatus : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["refreshDetails"] = args ? args.refreshDetails : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["shareStatus"] = args ? args.shareStatus : undefined;
            inputs["userAccessRights"] = args ? args.userAccessRights : undefined;
            inputs["shareMappings"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:databoxedge/latest:Share" }, { type: "azurerm:databoxedge/v20190301:Share" }, { type: "azurerm:databoxedge/v20190801:Share" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Share.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Share resource.
 */
export interface ShareArgs {
    /**
     * Access protocol to be used by the share.
     */
    readonly accessProtocol: pulumi.Input<string>;
    /**
     * Azure container mapping for the share.
     */
    readonly azureContainerInfo?: pulumi.Input<inputs.databoxedge.v20190701.AzureContainerInfo>;
    /**
     * List of IP addresses and corresponding access rights on the share(required for NFS protocol).
     */
    readonly clientAccessRights?: pulumi.Input<pulumi.Input<inputs.databoxedge.v20190701.ClientAccessRight>[]>;
    /**
     * Data policy of the share.
     */
    readonly dataPolicy?: pulumi.Input<string>;
    /**
     * Description for the share.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The device name.
     */
    readonly deviceName: pulumi.Input<string>;
    /**
     * Current monitoring status of the share.
     */
    readonly monitoringStatus: pulumi.Input<string>;
    /**
     * The share name.
     */
    readonly name: pulumi.Input<string>;
    /**
     * Details of the refresh job on this share.
     */
    readonly refreshDetails?: pulumi.Input<inputs.databoxedge.v20190701.RefreshDetails>;
    /**
     * The resource group name.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Current status of the share.
     */
    readonly shareStatus: pulumi.Input<string>;
    /**
     * Mapping of users and corresponding access rights on the share (required for SMB protocol).
     */
    readonly userAccessRights?: pulumi.Input<pulumi.Input<inputs.databoxedge.v20190701.UserAccessRight>[]>;
}
