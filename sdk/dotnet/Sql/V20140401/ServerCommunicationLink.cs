// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Sql.V20140401
{
    /// <summary>
    /// Server communication link.
    /// </summary>
    public partial class ServerCommunicationLink : Pulumi.CustomResource
    {
        /// <summary>
        /// Communication link kind.  This property is used for Azure Portal metadata.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Communication link location.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The properties of resource.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.ServerCommunicationLinkPropertiesResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ServerCommunicationLink resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServerCommunicationLink(string name, ServerCommunicationLinkArgs args, CustomResourceOptions? options = null)
            : base("azurerm:sql/v20140401:ServerCommunicationLink", name, args ?? new ServerCommunicationLinkArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServerCommunicationLink(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:sql/v20140401:ServerCommunicationLink", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ServerCommunicationLink resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServerCommunicationLink Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ServerCommunicationLink(name, id, options);
        }
    }

    public sealed class ServerCommunicationLinkArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the server communication link.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the partner server.
        /// </summary>
        [Input("partnerServer", required: true)]
        public Input<string> PartnerServer { get; set; } = null!;

        /// <summary>
        /// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the server.
        /// </summary>
        [Input("serverName", required: true)]
        public Input<string> ServerName { get; set; } = null!;

        public ServerCommunicationLinkArgs()
        {
        }
    }
}
