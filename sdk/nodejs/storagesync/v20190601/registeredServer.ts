// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Registered Server resource.
 *
 * ## Example Usage
 * ### RegisteredServers_Create
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const registeredServer = new azurerm.storagesync.v20190601.RegisteredServer("registeredServer", {
 *     agentVersion: "1.0.277.0",
 *     lastHeartBeat: "\"2017-08-08T18:29:06.470652Z\"",
 *     resourceGroupName: "SampleResourceGroup_1",
 *     serverCertificate: "\"MIIDFjCCAf6gAwIBAgIQQS+DS8uhc4VNzUkTw7wbRjANBgkqhkiG9w0BAQ0FADAzMTEwLwYDVQQDEyhhbmt1c2hiLXByb2QzLnJlZG1vbmQuY29ycC5taWNyb3NvZnQuY29tMB4XDTE3MDgwMzE3MDQyNFoXDTE4MDgwNDE3MDQyNFowMzExMC8GA1UEAxMoYW5rdXNoYi1wcm9kMy5yZWRtb25kLmNvcnAubWljcm9zb2Z0LmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBALDRvV4gmsIy6jGDPiHsXmvgVP749NNP7DopdlbHaNhjFmYINHl0uWylyaZmgJrROt2mnxN/zEyJtGnqYHlzUr4xvGq/qV5pqgdB9tag/sw9i22gfe9PRZ0FmSOZnXMbLYgLiDFqLtut5gHcOuWMj03YnkfoBEKlFBxWbagvW2yxz/Sxi9OVSJOKCaXra0RpcIHrO/KFl6ho2eE1/7Ykmfa8hZvSdoPd5gHdLiQcMB/pxq+mWp1fI6c8vFZoDu7Atn+NXTzYPKUxKzaisF12TsaKpohUsJpbB3Wocb0F5frn614D2pg14ERB5otjAMWw1m65csQWPI6dP8KIYe0+QPkCAwEAAaMmMCQwIgYDVR0lAQH/BBgwFgYIKwYBBQUHAwIGCisGAQQBgjcKAwwwDQYJKoZIhvcNAQENBQADggEBAA4RhVIBkw34M1RwakJgHvtjsOFxF1tVQA941NtLokx1l2Z8+GFQkcG4xpZSt+UN6wLerdCbnNhtkCErWUDeaT0jxk4g71Ofex7iM04crT4iHJr8mi96/XnhnkTUs+GDk12VgdeeNEczMZz+8Mxw9dJ5NCnYgTwO0SzGlclRsDvjzkLo8rh2ZG6n/jKrEyNXXo+hOqhupij0QbRP2Tvexdfw201kgN1jdZify8XzJ8Oi0bTS0KpJf2pNPOlooK2bjMUei9ANtEdXwwfVZGWvVh6tJjdv6k14wWWJ1L7zhA1IIVb1J+sQUzJji5iX0DrezjTz1Fg+gAzITaA/WsuujlM=\"",
 *     serverId: "\"080d4133-bdb5-40a0-96a0-71a6057bfe9a\"",
 *     serverOSVersion: "10.0.14393.0",
 *     serverRole: "Standalone",
 *     storageSyncServiceName: "SampleStorageSyncService_1",
 * });
 *
 * ```
 */
export class RegisteredServer extends pulumi.CustomResource {
    /**
     * Get an existing RegisteredServer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): RegisteredServer {
        return new RegisteredServer(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:storagesync/v20190601:RegisteredServer';

    /**
     * Returns true if the given object is an instance of RegisteredServer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RegisteredServer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RegisteredServer.__pulumiType;
    }

    /**
     * Registered Server Agent Version
     */
    public readonly agentVersion!: pulumi.Output<string | undefined>;
    /**
     * Registered Server clusterId
     */
    public readonly clusterId!: pulumi.Output<string | undefined>;
    /**
     * Registered Server clusterName
     */
    public readonly clusterName!: pulumi.Output<string | undefined>;
    /**
     * Resource discoveryEndpointUri
     */
    public /*out*/ readonly discoveryEndpointUri!: pulumi.Output<string | undefined>;
    /**
     * Friendly Name
     */
    public readonly friendlyName!: pulumi.Output<string | undefined>;
    /**
     * Registered Server last heart beat
     */
    public readonly lastHeartBeat!: pulumi.Output<string | undefined>;
    /**
     * Resource Last Operation Name
     */
    public /*out*/ readonly lastOperationName!: pulumi.Output<string | undefined>;
    /**
     * Registered Server lastWorkflowId
     */
    public /*out*/ readonly lastWorkflowId!: pulumi.Output<string | undefined>;
    /**
     * Management Endpoint Uri
     */
    public /*out*/ readonly managementEndpointUri!: pulumi.Output<string | undefined>;
    /**
     * Monitoring Configuration
     */
    public /*out*/ readonly monitoringConfiguration!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Registered Server Provisioning State
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string | undefined>;
    /**
     * Resource Location
     */
    public /*out*/ readonly resourceLocation!: pulumi.Output<string | undefined>;
    /**
     * Registered Server Certificate
     */
    public readonly serverCertificate!: pulumi.Output<string | undefined>;
    /**
     * Registered Server serverId
     */
    public readonly serverId!: pulumi.Output<string | undefined>;
    /**
     * Registered Server Management Error Code
     */
    public /*out*/ readonly serverManagementErrorCode!: pulumi.Output<number | undefined>;
    /**
     * Registered Server OS Version
     */
    public readonly serverOSVersion!: pulumi.Output<string | undefined>;
    /**
     * Registered Server serverRole
     */
    public readonly serverRole!: pulumi.Output<string | undefined>;
    /**
     * Service Location
     */
    public /*out*/ readonly serviceLocation!: pulumi.Output<string | undefined>;
    /**
     * Registered Server storageSyncServiceUid
     */
    public /*out*/ readonly storageSyncServiceUid!: pulumi.Output<string | undefined>;
    /**
     * The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a RegisteredServer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RegisteredServerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RegisteredServerArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as RegisteredServerArgs | undefined;
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.serverId === undefined) {
                throw new Error("Missing required property 'serverId'");
            }
            if (!args || args.storageSyncServiceName === undefined) {
                throw new Error("Missing required property 'storageSyncServiceName'");
            }
            inputs["agentVersion"] = args ? args.agentVersion : undefined;
            inputs["clusterId"] = args ? args.clusterId : undefined;
            inputs["clusterName"] = args ? args.clusterName : undefined;
            inputs["friendlyName"] = args ? args.friendlyName : undefined;
            inputs["lastHeartBeat"] = args ? args.lastHeartBeat : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["serverCertificate"] = args ? args.serverCertificate : undefined;
            inputs["serverId"] = args ? args.serverId : undefined;
            inputs["serverOSVersion"] = args ? args.serverOSVersion : undefined;
            inputs["serverRole"] = args ? args.serverRole : undefined;
            inputs["storageSyncServiceName"] = args ? args.storageSyncServiceName : undefined;
            inputs["discoveryEndpointUri"] = undefined /*out*/;
            inputs["lastOperationName"] = undefined /*out*/;
            inputs["lastWorkflowId"] = undefined /*out*/;
            inputs["managementEndpointUri"] = undefined /*out*/;
            inputs["monitoringConfiguration"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["resourceLocation"] = undefined /*out*/;
            inputs["serverManagementErrorCode"] = undefined /*out*/;
            inputs["serviceLocation"] = undefined /*out*/;
            inputs["storageSyncServiceUid"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:storagesync/latest:RegisteredServer" }, { type: "azurerm:storagesync/v20170605preview:RegisteredServer" }, { type: "azurerm:storagesync/v20180402:RegisteredServer" }, { type: "azurerm:storagesync/v20180701:RegisteredServer" }, { type: "azurerm:storagesync/v20181001:RegisteredServer" }, { type: "azurerm:storagesync/v20190201:RegisteredServer" }, { type: "azurerm:storagesync/v20190301:RegisteredServer" }, { type: "azurerm:storagesync/v20191001:RegisteredServer" }, { type: "azurerm:storagesync/v20200301:RegisteredServer" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(RegisteredServer.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a RegisteredServer resource.
 */
export interface RegisteredServerArgs {
    /**
     * Registered Server Agent Version
     */
    readonly agentVersion?: pulumi.Input<string>;
    /**
     * Registered Server clusterId
     */
    readonly clusterId?: pulumi.Input<string>;
    /**
     * Registered Server clusterName
     */
    readonly clusterName?: pulumi.Input<string>;
    /**
     * Friendly Name
     */
    readonly friendlyName?: pulumi.Input<string>;
    /**
     * Registered Server last heart beat
     */
    readonly lastHeartBeat?: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Registered Server Certificate
     */
    readonly serverCertificate?: pulumi.Input<string>;
    /**
     * Registered Server serverId
     */
    readonly serverId: pulumi.Input<string>;
    /**
     * Registered Server OS Version
     */
    readonly serverOSVersion?: pulumi.Input<string>;
    /**
     * Registered Server serverRole
     */
    readonly serverRole?: pulumi.Input<string>;
    /**
     * Name of Storage Sync Service resource.
     */
    readonly storageSyncServiceName: pulumi.Input<string>;
}
