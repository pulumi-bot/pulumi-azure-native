// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DBforMySQL.V20200101privatePreview
{
    /// <summary>
    /// A MySQL Server key.
    /// </summary>
    public partial class ServerKey : Pulumi.CustomResource
    {
        /// <summary>
        /// The key creation date.
        /// </summary>
        [Output("creationDate")]
        public Output<string> CreationDate { get; private set; } = null!;

        /// <summary>
        /// Kind of encryption protector. This is metadata used for the Azure portal experience.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The key type like 'AzureKeyVault'.
        /// </summary>
        [Output("serverKeyType")]
        public Output<string> ServerKeyType { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The URI of the key.
        /// </summary>
        [Output("uri")]
        public Output<string?> Uri { get; private set; } = null!;


        /// <summary>
        /// Create a ServerKey resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServerKey(string name, ServerKeyArgs args, CustomResourceOptions? options = null)
            : base("azurerm:dbformysql/v20200101privatepreview:ServerKey", name, args ?? new ServerKeyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServerKey(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:dbformysql/v20200101privatepreview:ServerKey", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:dbformysql/latest:ServerKey"},
                    new Pulumi.Alias { Type = "azurerm:dbformysql/v20200101:ServerKey"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ServerKey resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServerKey Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ServerKey(name, id, options);
        }
    }

    public sealed class ServerKeyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the MySQL Server key to be operated on (updated or created).
        /// </summary>
        [Input("keyName", required: true)]
        public Input<string> KeyName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The key type like 'AzureKeyVault'.
        /// </summary>
        [Input("serverKeyType", required: true)]
        public Input<string> ServerKeyType { get; set; } = null!;

        /// <summary>
        /// The name of the server.
        /// </summary>
        [Input("serverName", required: true)]
        public Input<string> ServerName { get; set; } = null!;

        /// <summary>
        /// The URI of the key.
        /// </summary>
        [Input("uri")]
        public Input<string>? Uri { get; set; }

        public ServerKeyArgs()
        {
        }
    }
}
