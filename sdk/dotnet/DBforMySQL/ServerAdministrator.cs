// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DBforMySQL
{
    /// <summary>
    /// Represents a and external administrator to be created.
    /// </summary>
    public partial class ServerAdministrator : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Properties of the server AAD administrator.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.ServerAdministratorPropertiesResponse> Properties { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ServerAdministrator resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServerAdministrator(string name, ServerAdministratorArgs args, CustomResourceOptions? options = null)
            : base("azurerm:dbformysql:ServerAdministrator", name, args ?? new ServerAdministratorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServerAdministrator(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:dbformysql:ServerAdministrator", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ServerAdministrator resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServerAdministrator Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ServerAdministrator(name, id, options);
        }
    }

    public sealed class ServerAdministratorArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the server.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Properties of the server AAD administrator.
        /// </summary>
        [Input("properties")]
        public Input<Inputs.ServerAdministratorPropertiesArgs>? Properties { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public ServerAdministratorArgs()
        {
        }
    }
}
