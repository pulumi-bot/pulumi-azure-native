// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Devices.Latest
{
    /// <summary>
    /// The X509 Certificate.
    /// 
    /// ## Example Usage
    /// ### Certificates_CreateOrUpdate
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var certificate = new AzureRM.Devices.Latest.Certificate("certificate", new AzureRM.Devices.Latest.CertificateArgs
    ///         {
    ///             CertificateName = "cert",
    ///             ResourceGroupName = "myResourceGroup",
    ///             ResourceName = "iothub",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class Certificate : Pulumi.CustomResource
    {
        /// <summary>
        /// The entity tag.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The name of the certificate.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The description of an X509 CA Certificate.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.CertificatePropertiesResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// The resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Certificate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Certificate(string name, CertificateArgs args, CustomResourceOptions? options = null)
            : base("azurerm:devices/latest:Certificate", name, args ?? new CertificateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Certificate(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:devices/latest:Certificate", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:devices/v20170701:Certificate"},
                    new Pulumi.Alias { Type = "azurerm:devices/v20180122:Certificate"},
                    new Pulumi.Alias { Type = "azurerm:devices/v20180401:Certificate"},
                    new Pulumi.Alias { Type = "azurerm:devices/v20181201preview:Certificate"},
                    new Pulumi.Alias { Type = "azurerm:devices/v20190322:Certificate"},
                    new Pulumi.Alias { Type = "azurerm:devices/v20190322preview:Certificate"},
                    new Pulumi.Alias { Type = "azurerm:devices/v20190701preview:Certificate"},
                    new Pulumi.Alias { Type = "azurerm:devices/v20191104:Certificate"},
                    new Pulumi.Alias { Type = "azurerm:devices/v20200301:Certificate"},
                    new Pulumi.Alias { Type = "azurerm:devices/v20200401:Certificate"},
                    new Pulumi.Alias { Type = "azurerm:devices/v20200615:Certificate"},
                    new Pulumi.Alias { Type = "azurerm:devices/v20200710preview:Certificate"},
                    new Pulumi.Alias { Type = "azurerm:devices/v20200801:Certificate"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Certificate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Certificate Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Certificate(name, id, options);
        }
    }

    public sealed class CertificateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the certificate
        /// </summary>
        [Input("certificateName", required: true)]
        public Input<string> CertificateName { get; set; } = null!;

        /// <summary>
        /// The description of an X509 CA Certificate.
        /// </summary>
        [Input("properties")]
        public Input<Inputs.CertificatePropertiesArgs>? Properties { get; set; }

        /// <summary>
        /// The name of the resource group that contains the IoT hub.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the IoT hub.
        /// </summary>
        [Input("resourceName", required: true)]
        public Input<string> ResourceName { get; set; } = null!;

        public CertificateArgs()
        {
        }
    }
}
