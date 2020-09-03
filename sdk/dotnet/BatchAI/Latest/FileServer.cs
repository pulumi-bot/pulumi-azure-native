// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.BatchAI.Latest
{
    /// <summary>
    /// File Server information.
    /// 
    /// ## Example Usage
    /// ### Create a file server
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var fileServer = new AzureRM.BatchAI.Latest.FileServer("fileServer", new AzureRM.BatchAI.Latest.FileServerArgs
    ///         {
    ///             DataDisks = new AzureRM.BatchAI.Latest.Inputs.DataDisksArgs
    ///             {
    ///                 DiskCount = 2,
    ///                 DiskSizeInGB = 10,
    ///                 StorageAccountType = "Standard_LRS",
    ///             },
    ///             FileServerName = "demo_nfs",
    ///             ResourceGroupName = "demo_resource_group",
    ///             SshConfiguration = new AzureRM.BatchAI.Latest.Inputs.SshConfigurationArgs
    ///             {
    ///                 UserAccountSettings = new AzureRM.BatchAI.Latest.Inputs.UserAccountSettingsArgs
    ///                 {
    ///                     AdminUserName = "admin_user_name",
    ///                     AdminUserPassword = "admin_user_password",
    ///                 },
    ///             },
    ///             VmSize = "STANDARD_NC6",
    ///             WorkspaceName = "demo_workspace",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class FileServer : Pulumi.CustomResource
    {
        /// <summary>
        /// Time when the FileServer was created.
        /// </summary>
        [Output("creationTime")]
        public Output<string> CreationTime { get; private set; } = null!;

        /// <summary>
        /// Information about disks attached to File Server VM.
        /// </summary>
        [Output("dataDisks")]
        public Output<Outputs.DataDisksResponseResult?> DataDisks { get; private set; } = null!;

        /// <summary>
        /// File Server mount settings.
        /// </summary>
        [Output("mountSettings")]
        public Output<Outputs.MountSettingsResponseResult> MountSettings { get; private set; } = null!;

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Provisioning state of the File Server. Possible values: creating - The File Server is getting created; updating - The File Server creation has been accepted and it is getting updated; deleting - The user has requested that the File Server be deleted, and it is in the process of being deleted; failed - The File Server creation has failed with the specified error code. Details about the error code are specified in the message field; succeeded - The File Server creation has succeeded.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Time when the provisioning state was changed.
        /// </summary>
        [Output("provisioningStateTransitionTime")]
        public Output<string> ProvisioningStateTransitionTime { get; private set; } = null!;

        /// <summary>
        /// SSH configuration for accessing the File Server node.
        /// </summary>
        [Output("sshConfiguration")]
        public Output<Outputs.SshConfigurationResponseResult?> SshConfiguration { get; private set; } = null!;

        /// <summary>
        /// File Server virtual network subnet resource ID.
        /// </summary>
        [Output("subnet")]
        public Output<Outputs.ResourceIdResponseResult?> Subnet { get; private set; } = null!;

        /// <summary>
        /// The type of the resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// VM size of the File Server.
        /// </summary>
        [Output("vmSize")]
        public Output<string?> VmSize { get; private set; } = null!;


        /// <summary>
        /// Create a FileServer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FileServer(string name, FileServerArgs args, CustomResourceOptions? options = null)
            : base("azurerm:batchai/latest:FileServer", name, args ?? new FileServerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FileServer(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:batchai/latest:FileServer", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:batchai/v20180501:FileServer"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing FileServer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FileServer Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new FileServer(name, id, options);
        }
    }

    public sealed class FileServerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Settings for the data disks which will be created for the File Server.
        /// </summary>
        [Input("dataDisks", required: true)]
        public Input<Inputs.DataDisksArgs> DataDisks { get; set; } = null!;

        /// <summary>
        /// The name of the file server within the specified resource group. File server names can only contain a combination of alphanumeric characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
        /// </summary>
        [Input("fileServerName", required: true)]
        public Input<string> FileServerName { get; set; } = null!;

        /// <summary>
        /// Name of the resource group to which the resource belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// SSH configuration for the File Server node.
        /// </summary>
        [Input("sshConfiguration", required: true)]
        public Input<Inputs.SshConfigurationArgs> SshConfiguration { get; set; } = null!;

        /// <summary>
        /// Identifier of an existing virtual network subnet to put the File Server in. If not provided, a new virtual network and subnet will be created.
        /// </summary>
        [Input("subnet")]
        public Input<Inputs.ResourceIdArgs>? Subnet { get; set; }

        /// <summary>
        /// The size of the virtual machine for the File Server. For information about available VM sizes from the Virtual Machines Marketplace, see Sizes for Virtual Machines (Linux).
        /// </summary>
        [Input("vmSize", required: true)]
        public Input<string> VmSize { get; set; } = null!;

        /// <summary>
        /// The name of the workspace. Workspace names can only contain a combination of alphanumeric characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
        /// </summary>
        [Input("workspaceName", required: true)]
        public Input<string> WorkspaceName { get; set; } = null!;

        public FileServerArgs()
        {
        }
    }
}
