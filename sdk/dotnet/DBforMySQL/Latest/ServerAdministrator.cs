// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DBforMySQL.Latest
{
    /// <summary>
    /// Represents a and external administrator to be created.
    /// 
    /// ## Example Usage
    /// ### ServerAdministratorCreate
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var serverAdministrator = new AzureRM.DBforMySQL.Latest.ServerAdministrator("serverAdministrator", new AzureRM.DBforMySQL.Latest.ServerAdministratorArgs
    ///         {
    ///             AdministratorType = "ActiveDirectory",
    ///             Login = "bob@contoso.com",
    ///             ResourceGroupName = "testrg",
    ///             ServerName = "mysqltestsvc4",
    ///             Sid = "c6b82b90-a647-49cb-8a62-0d2d3cb7ac7c",
    ///             TenantId = "c6b82b90-a647-49cb-8a62-0d2d3cb7ac7c",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class ServerAdministrator : Pulumi.CustomResource
    {
        /// <summary>
        /// The type of administrator.
        /// </summary>
        [Output("administratorType")]
        public Output<string> AdministratorType { get; private set; } = null!;

        /// <summary>
        /// The server administrator login account name.
        /// </summary>
        [Output("login")]
        public Output<string> Login { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The server administrator Sid (Secure ID).
        /// </summary>
        [Output("sid")]
        public Output<string> Sid { get; private set; } = null!;

        /// <summary>
        /// The server Active Directory Administrator tenant id.
        /// </summary>
        [Output("tenantId")]
        public Output<string> TenantId { get; private set; } = null!;

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
            : base("azurerm:dbformysql/latest:ServerAdministrator", name, args ?? new ServerAdministratorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServerAdministrator(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:dbformysql/latest:ServerAdministrator", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:dbformysql/v20171201:ServerAdministrator"},
                    new Pulumi.Alias { Type = "azurerm:dbformysql/v20171201preview:ServerAdministrator"},
                },
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
        /// The type of administrator.
        /// </summary>
        [Input("administratorType", required: true)]
        public Input<string> AdministratorType { get; set; } = null!;

        /// <summary>
        /// The server administrator login account name.
        /// </summary>
        [Input("login", required: true)]
        public Input<string> Login { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the server.
        /// </summary>
        [Input("serverName", required: true)]
        public Input<string> ServerName { get; set; } = null!;

        /// <summary>
        /// The server administrator Sid (Secure ID).
        /// </summary>
        [Input("sid", required: true)]
        public Input<string> Sid { get; set; } = null!;

        /// <summary>
        /// The server Active Directory Administrator tenant id.
        /// </summary>
        [Input("tenantId", required: true)]
        public Input<string> TenantId { get; set; } = null!;

        public ServerAdministratorArgs()
        {
        }
    }
}
