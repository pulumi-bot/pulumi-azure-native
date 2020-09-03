// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Compute.V20190301
{
    /// <summary>
    /// The source user image virtual hard disk. The virtual hard disk will be copied before being attached to the virtual machine. If SourceImage is provided, the destination virtual hard drive must not exist.
    /// 
    /// ## Example Usage
    /// ### Create a virtual machine image from a blob.
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var image = new AzureRM.Compute.V20190301.Image("image", new AzureRM.Compute.V20190301.ImageArgs
    ///         {
    ///             ImageName = "myImage",
    ///             Location = "West US",
    ///             ResourceGroupName = "myResourceGroup",
    ///             StorageProfile = new AzureRM.Compute.V20190301.Inputs.ImageStorageProfileArgs
    ///             {
    ///                 OsDisk = new AzureRM.Compute.V20190301.Inputs.ImageOSDiskArgs
    ///                 {
    ///                     BlobUri = "https://mystorageaccount.blob.core.windows.net/osimages/osimage.vhd",
    ///                     OsState = "Generalized",
    ///                     OsType = "Linux",
    ///                 },
    ///                 ZoneResilient = true,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// ### Create a virtual machine image from a managed disk.
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var image = new AzureRM.Compute.V20190301.Image("image", new AzureRM.Compute.V20190301.ImageArgs
    ///         {
    ///             ImageName = "myImage",
    ///             Location = "West US",
    ///             ResourceGroupName = "myResourceGroup",
    ///             StorageProfile = new AzureRM.Compute.V20190301.Inputs.ImageStorageProfileArgs
    ///             {
    ///                 OsDisk = new AzureRM.Compute.V20190301.Inputs.ImageOSDiskArgs
    ///                 {
    ///                     ManagedDisk = new AzureRM.Compute.V20190301.Inputs.SubResourceArgs
    ///                     {
    ///                         Id = "subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myManagedDisk",
    ///                     },
    ///                     OsState = "Generalized",
    ///                     OsType = "Linux",
    ///                 },
    ///                 ZoneResilient = true,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// ### Create a virtual machine image from a snapshot.
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var image = new AzureRM.Compute.V20190301.Image("image", new AzureRM.Compute.V20190301.ImageArgs
    ///         {
    ///             ImageName = "myImage",
    ///             Location = "West US",
    ///             ResourceGroupName = "myResourceGroup",
    ///             StorageProfile = new AzureRM.Compute.V20190301.Inputs.ImageStorageProfileArgs
    ///             {
    ///                 OsDisk = new AzureRM.Compute.V20190301.Inputs.ImageOSDiskArgs
    ///                 {
    ///                     OsState = "Generalized",
    ///                     OsType = "Linux",
    ///                     Snapshot = new AzureRM.Compute.V20190301.Inputs.SubResourceArgs
    ///                     {
    ///                         Id = "subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot",
    ///                     },
    ///                 },
    ///                 ZoneResilient = false,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// ### Create a virtual machine image from an existing virtual machine.
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var image = new AzureRM.Compute.V20190301.Image("image", new AzureRM.Compute.V20190301.ImageArgs
    ///         {
    ///             ImageName = "myImage",
    ///             Location = "West US",
    ///             ResourceGroupName = "myResourceGroup",
    ///             SourceVirtualMachine = new AzureRM.Compute.V20190301.Inputs.SubResourceArgs
    ///             {
    ///                 Id = "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// ### Create a virtual machine image that includes a data disk from a blob.
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var image = new AzureRM.Compute.V20190301.Image("image", new AzureRM.Compute.V20190301.ImageArgs
    ///         {
    ///             ImageName = "myImage",
    ///             Location = "West US",
    ///             ResourceGroupName = "myResourceGroup",
    ///             StorageProfile = new AzureRM.Compute.V20190301.Inputs.ImageStorageProfileArgs
    ///             {
    ///                 DataDisks = 
    ///                 {
    ///                     new AzureRM.Compute.V20190301.Inputs.ImageDataDiskArgs
    ///                     {
    ///                         BlobUri = "https://mystorageaccount.blob.core.windows.net/dataimages/dataimage.vhd",
    ///                         Lun = 1,
    ///                     },
    ///                 },
    ///                 OsDisk = new AzureRM.Compute.V20190301.Inputs.ImageOSDiskArgs
    ///                 {
    ///                     BlobUri = "https://mystorageaccount.blob.core.windows.net/osimages/osimage.vhd",
    ///                     OsState = "Generalized",
    ///                     OsType = "Linux",
    ///                 },
    ///                 ZoneResilient = false,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// ### Create a virtual machine image that includes a data disk from a managed disk.
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var image = new AzureRM.Compute.V20190301.Image("image", new AzureRM.Compute.V20190301.ImageArgs
    ///         {
    ///             ImageName = "myImage",
    ///             Location = "West US",
    ///             ResourceGroupName = "myResourceGroup",
    ///             StorageProfile = new AzureRM.Compute.V20190301.Inputs.ImageStorageProfileArgs
    ///             {
    ///                 DataDisks = 
    ///                 {
    ///                     new AzureRM.Compute.V20190301.Inputs.ImageDataDiskArgs
    ///                     {
    ///                         Lun = 1,
    ///                         ManagedDisk = new AzureRM.Compute.V20190301.Inputs.SubResourceArgs
    ///                         {
    ///                             Id = "subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myManagedDisk2",
    ///                         },
    ///                     },
    ///                 },
    ///                 OsDisk = new AzureRM.Compute.V20190301.Inputs.ImageOSDiskArgs
    ///                 {
    ///                     ManagedDisk = new AzureRM.Compute.V20190301.Inputs.SubResourceArgs
    ///                     {
    ///                         Id = "subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myManagedDisk",
    ///                     },
    ///                     OsState = "Generalized",
    ///                     OsType = "Linux",
    ///                 },
    ///                 ZoneResilient = false,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// ### Create a virtual machine image that includes a data disk from a snapshot.
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var image = new AzureRM.Compute.V20190301.Image("image", new AzureRM.Compute.V20190301.ImageArgs
    ///         {
    ///             ImageName = "myImage",
    ///             Location = "West US",
    ///             ResourceGroupName = "myResourceGroup",
    ///             StorageProfile = new AzureRM.Compute.V20190301.Inputs.ImageStorageProfileArgs
    ///             {
    ///                 DataDisks = 
    ///                 {
    ///                     new AzureRM.Compute.V20190301.Inputs.ImageDataDiskArgs
    ///                     {
    ///                         Lun = 1,
    ///                         Snapshot = new AzureRM.Compute.V20190301.Inputs.SubResourceArgs
    ///                         {
    ///                             Id = "subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot2",
    ///                         },
    ///                     },
    ///                 },
    ///                 OsDisk = new AzureRM.Compute.V20190301.Inputs.ImageOSDiskArgs
    ///                 {
    ///                     OsState = "Generalized",
    ///                     OsType = "Linux",
    ///                     Snapshot = new AzureRM.Compute.V20190301.Inputs.SubResourceArgs
    ///                     {
    ///                         Id = "subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot",
    ///                     },
    ///                 },
    ///                 ZoneResilient = true,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class Image : Pulumi.CustomResource
    {
        /// <summary>
        /// Gets the HyperVGenerationType of the VirtualMachine created from the image
        /// </summary>
        [Output("hyperVGeneration")]
        public Output<string?> HyperVGeneration { get; private set; } = null!;

        /// <summary>
        /// Resource location
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Resource name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The provisioning state.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The source virtual machine from which Image is created.
        /// </summary>
        [Output("sourceVirtualMachine")]
        public Output<Outputs.SubResourceResponseResult?> SourceVirtualMachine { get; private set; } = null!;

        /// <summary>
        /// Specifies the storage settings for the virtual machine disks.
        /// </summary>
        [Output("storageProfile")]
        public Output<Outputs.ImageStorageProfileResponseResult?> StorageProfile { get; private set; } = null!;

        /// <summary>
        /// Resource tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Resource type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Image resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Image(string name, ImageArgs args, CustomResourceOptions? options = null)
            : base("azurerm:compute/v20190301:Image", name, args ?? new ImageArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Image(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:compute/v20190301:Image", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:compute/latest:Image"},
                    new Pulumi.Alias { Type = "azurerm:compute/v20160430preview:Image"},
                    new Pulumi.Alias { Type = "azurerm:compute/v20170330:Image"},
                    new Pulumi.Alias { Type = "azurerm:compute/v20171201:Image"},
                    new Pulumi.Alias { Type = "azurerm:compute/v20180401:Image"},
                    new Pulumi.Alias { Type = "azurerm:compute/v20180601:Image"},
                    new Pulumi.Alias { Type = "azurerm:compute/v20181001:Image"},
                    new Pulumi.Alias { Type = "azurerm:compute/v20190701:Image"},
                    new Pulumi.Alias { Type = "azurerm:compute/v20191201:Image"},
                    new Pulumi.Alias { Type = "azurerm:compute/v20200601:Image"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Image resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Image Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Image(name, id, options);
        }
    }

    public sealed class ImageArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Gets the HyperVGenerationType of the VirtualMachine created from the image
        /// </summary>
        [Input("hyperVGeneration")]
        public Input<string>? HyperVGeneration { get; set; }

        /// <summary>
        /// The name of the image.
        /// </summary>
        [Input("imageName", required: true)]
        public Input<string> ImageName { get; set; } = null!;

        /// <summary>
        /// Resource location
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The source virtual machine from which Image is created.
        /// </summary>
        [Input("sourceVirtualMachine")]
        public Input<Inputs.SubResourceArgs>? SourceVirtualMachine { get; set; }

        /// <summary>
        /// Specifies the storage settings for the virtual machine disks.
        /// </summary>
        [Input("storageProfile")]
        public Input<Inputs.ImageStorageProfileArgs>? StorageProfile { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ImageArgs()
        {
        }
    }
}
