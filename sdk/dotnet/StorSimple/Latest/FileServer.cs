// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.StorSimple.Latest
{
    /// <summary>
    /// The file server.
    /// 
    /// ## Example Usage
    /// ### FileServersCreateOrUpdate
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var fileServer = new AzureRM.StorSimple.Latest.FileServer("fileServer", new AzureRM.StorSimple.Latest.FileServerArgs
    ///         {
    ///             BackupScheduleGroupId = "/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/hsdk-4xy4fi2ivg/backupScheduleGroups/BackupSchGroupForSDKTest",
    ///             Description = "Demo FileServer for SDK Test",
    ///             DeviceName = "HSDK-4XY4FI2IVG",
    ///             DomainName = "fareast.corp.microsoft.com",
    ///             FileServerName = "HSDK-4XY4FI2IVG",
    ///             ManagerName = "hAzureSDKOperations",
    ///             ResourceGroupName = "ResourceGroupForSDKTest",
    ///             StorageDomainId = "/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/storageDomains/sd-fs-HSDK-4XY4FI2IVG",
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
        /// The backup policy id.
        /// </summary>
        [Output("backupScheduleGroupId")]
        public Output<string> BackupScheduleGroupId { get; private set; } = null!;

        /// <summary>
        /// The description of the file server
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Domain of the file server
        /// </summary>
        [Output("domainName")]
        public Output<string> DomainName { get; private set; } = null!;

        /// <summary>
        /// The name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The storage domain id.
        /// </summary>
        [Output("storageDomainId")]
        public Output<string> StorageDomainId { get; private set; } = null!;

        /// <summary>
        /// The type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a FileServer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FileServer(string name, FileServerArgs args, CustomResourceOptions? options = null)
            : base("azurerm:storsimple/latest:FileServer", name, args ?? new FileServerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FileServer(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:storsimple/latest:FileServer", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:storsimple/v20161001:FileServer"},
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
        /// The backup policy id.
        /// </summary>
        [Input("backupScheduleGroupId", required: true)]
        public Input<string> BackupScheduleGroupId { get; set; } = null!;

        /// <summary>
        /// The description of the file server
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The device name.
        /// </summary>
        [Input("deviceName", required: true)]
        public Input<string> DeviceName { get; set; } = null!;

        /// <summary>
        /// Domain of the file server
        /// </summary>
        [Input("domainName", required: true)]
        public Input<string> DomainName { get; set; } = null!;

        /// <summary>
        /// The file server name.
        /// </summary>
        [Input("fileServerName", required: true)]
        public Input<string> FileServerName { get; set; } = null!;

        /// <summary>
        /// The manager name
        /// </summary>
        [Input("managerName", required: true)]
        public Input<string> ManagerName { get; set; } = null!;

        /// <summary>
        /// The resource group name
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The storage domain id.
        /// </summary>
        [Input("storageDomainId", required: true)]
        public Input<string> StorageDomainId { get; set; } = null!;

        public FileServerArgs()
        {
        }
    }
}
