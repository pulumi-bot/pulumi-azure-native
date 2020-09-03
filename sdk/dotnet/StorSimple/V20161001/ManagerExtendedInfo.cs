// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.StorSimple.V20161001
{
    /// <summary>
    /// The extended info of the manager.
    /// 
    /// ## Example Usage
    /// ### ManagersCreateExtendedInfo
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var managerExtendedInfo = new AzureRM.StorSimple.V20161001.ManagerExtendedInfo("managerExtendedInfo", new AzureRM.StorSimple.V20161001.ManagerExtendedInfoArgs
    ///         {
    ///             Algorithm = "SHA256",
    ///             Etag = "6531d5d7-3ced-4f78-83b6-804368f2ca0c",
    ///             IntegrityKey = "e6501980-7efe-4602-bb0e-3cb9a08a6003",
    ///             ManagerName = "ManagerForSDKTest2",
    ///             ResourceGroupName = "ResourceGroupForSDKTest",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class ManagerExtendedInfo : Pulumi.CustomResource
    {
        /// <summary>
        /// Represents the encryption algorithm used to encrypt the other keys. None - if EncryptionKey is saved in plain text format. AlgorithmName - if encryption is used
        /// </summary>
        [Output("algorithm")]
        public Output<string> Algorithm { get; private set; } = null!;

        /// <summary>
        /// Represents the CEK of the resource
        /// </summary>
        [Output("encryptionKey")]
        public Output<string?> EncryptionKey { get; private set; } = null!;

        /// <summary>
        /// Represents the Cert thumbprint that was used to encrypt the CEK
        /// </summary>
        [Output("encryptionKeyThumbprint")]
        public Output<string?> EncryptionKeyThumbprint { get; private set; } = null!;

        /// <summary>
        /// ETag of the Resource
        /// </summary>
        [Output("etag")]
        public Output<string?> Etag { get; private set; } = null!;

        /// <summary>
        /// Represents the CIK of the resource
        /// </summary>
        [Output("integrityKey")]
        public Output<string> IntegrityKey { get; private set; } = null!;

        /// <summary>
        /// The name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Represents the portal thumbprint which can be used optionally to encrypt the entire data before storing it.
        /// </summary>
        [Output("portalCertificateThumbprint")]
        public Output<string?> PortalCertificateThumbprint { get; private set; } = null!;

        /// <summary>
        /// The type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Represents the version of the ExtendedInfo object being persisted
        /// </summary>
        [Output("version")]
        public Output<string?> Version { get; private set; } = null!;


        /// <summary>
        /// Create a ManagerExtendedInfo resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ManagerExtendedInfo(string name, ManagerExtendedInfoArgs args, CustomResourceOptions? options = null)
            : base("azurerm:storsimple/v20161001:ManagerExtendedInfo", name, args ?? new ManagerExtendedInfoArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ManagerExtendedInfo(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:storsimple/v20161001:ManagerExtendedInfo", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:storsimple/latest:ManagerExtendedInfo"},
                    new Pulumi.Alias { Type = "azurerm:storsimple/v20170601:ManagerExtendedInfo"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ManagerExtendedInfo resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ManagerExtendedInfo Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ManagerExtendedInfo(name, id, options);
        }
    }

    public sealed class ManagerExtendedInfoArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Represents the encryption algorithm used to encrypt the other keys. None - if EncryptionKey is saved in plain text format. AlgorithmName - if encryption is used
        /// </summary>
        [Input("algorithm", required: true)]
        public Input<string> Algorithm { get; set; } = null!;

        /// <summary>
        /// Represents the CEK of the resource
        /// </summary>
        [Input("encryptionKey")]
        public Input<string>? EncryptionKey { get; set; }

        /// <summary>
        /// Represents the Cert thumbprint that was used to encrypt the CEK
        /// </summary>
        [Input("encryptionKeyThumbprint")]
        public Input<string>? EncryptionKeyThumbprint { get; set; }

        /// <summary>
        /// ETag of the Resource
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// Represents the CIK of the resource
        /// </summary>
        [Input("integrityKey", required: true)]
        public Input<string> IntegrityKey { get; set; } = null!;

        /// <summary>
        /// The manager name
        /// </summary>
        [Input("managerName", required: true)]
        public Input<string> ManagerName { get; set; } = null!;

        /// <summary>
        /// Represents the portal thumbprint which can be used optionally to encrypt the entire data before storing it.
        /// </summary>
        [Input("portalCertificateThumbprint")]
        public Input<string>? PortalCertificateThumbprint { get; set; }

        /// <summary>
        /// The resource group name
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Represents the version of the ExtendedInfo object being persisted
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public ManagerExtendedInfoArgs()
        {
        }
    }
}
