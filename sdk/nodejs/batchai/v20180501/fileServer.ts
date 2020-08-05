// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * File Server information.
 */
export class FileServer extends pulumi.CustomResource {
    /**
     * Get an existing FileServer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): FileServer {
        return new FileServer(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:batchai/v20180501:FileServer';

    /**
     * Returns true if the given object is an instance of FileServer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FileServer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FileServer.__pulumiType;
    }

    /**
     * The name of the resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * File Server properties.
     */
    public /*out*/ readonly properties!: pulumi.Output<outputs.batchai.v20180501.FileServerPropertiesResponse>;
    /**
     * The type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a FileServer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FileServerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FileServerArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as FileServerArgs | undefined;
            if (!args || args.dataDisks === undefined) {
                throw new Error("Missing required property 'dataDisks'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.sshConfiguration === undefined) {
                throw new Error("Missing required property 'sshConfiguration'");
            }
            if (!args || args.vmSize === undefined) {
                throw new Error("Missing required property 'vmSize'");
            }
            if (!args || args.workspaceName === undefined) {
                throw new Error("Missing required property 'workspaceName'");
            }
            inputs["dataDisks"] = args ? args.dataDisks : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sshConfiguration"] = args ? args.sshConfiguration : undefined;
            inputs["subnet"] = args ? args.subnet : undefined;
            inputs["vmSize"] = args ? args.vmSize : undefined;
            inputs["workspaceName"] = args ? args.workspaceName : undefined;
            inputs["properties"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(FileServer.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a FileServer resource.
 */
export interface FileServerArgs {
    /**
     * Settings for the data disks which will be created for the File Server.
     */
    readonly dataDisks: pulumi.Input<inputs.batchai.v20180501.DataDisks>;
    /**
     * The name of the file server within the specified resource group. File server names can only contain a combination of alphanumeric characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
     */
    readonly name: pulumi.Input<string>;
    /**
     * Name of the resource group to which the resource belongs.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * SSH configuration for the File Server node.
     */
    readonly sshConfiguration: pulumi.Input<inputs.batchai.v20180501.SshConfiguration>;
    /**
     * Identifier of an existing virtual network subnet to put the File Server in. If not provided, a new virtual network and subnet will be created.
     */
    readonly subnet?: pulumi.Input<inputs.batchai.v20180501.ResourceId>;
    /**
     * The size of the virtual machine for the File Server. For information about available VM sizes from the Virtual Machines Marketplace, see Sizes for Virtual Machines (Linux).
     */
    readonly vmSize: pulumi.Input<string>;
    /**
     * The name of the workspace. Workspace names can only contain a combination of alphanumeric characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
     */
    readonly workspaceName: pulumi.Input<string>;
}
