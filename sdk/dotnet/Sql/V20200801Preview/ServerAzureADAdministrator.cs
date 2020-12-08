// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Sql.V20200801Preview
{
    /// <summary>
    /// Azure Active Directory administrator.
    /// </summary>
    public partial class ServerAzureADAdministrator : Pulumi.CustomResource
    {
        /// <summary>
        /// Type of the sever administrator.
        /// </summary>
        [Output("administratorType")]
        public Output<string> AdministratorType { get; private set; } = null!;

        /// <summary>
        /// Azure Active Directory only Authentication enabled.
        /// </summary>
        [Output("azureADOnlyAuthentication")]
        public Output<bool> AzureADOnlyAuthentication { get; private set; } = null!;

        /// <summary>
        /// Login name of the server administrator.
        /// </summary>
        [Output("login")]
        public Output<string> Login { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// SID (object ID) of the server administrator.
        /// </summary>
        [Output("sid")]
        public Output<string> Sid { get; private set; } = null!;

        /// <summary>
        /// Tenant ID of the administrator.
        /// </summary>
        [Output("tenantId")]
        public Output<string?> TenantId { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ServerAzureADAdministrator resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServerAzureADAdministrator(string name, ServerAzureADAdministratorArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:sql/v20200801preview:ServerAzureADAdministrator", name, args ?? new ServerAzureADAdministratorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServerAzureADAdministrator(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:sql/v20200801preview:ServerAzureADAdministrator", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:sql/latest:ServerAzureADAdministrator"},
                    new Pulumi.Alias { Type = "azure-nextgen:sql/v20140401:ServerAzureADAdministrator"},
                    new Pulumi.Alias { Type = "azure-nextgen:sql/v20180601preview:ServerAzureADAdministrator"},
                    new Pulumi.Alias { Type = "azure-nextgen:sql/v20190601preview:ServerAzureADAdministrator"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ServerAzureADAdministrator resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServerAzureADAdministrator Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ServerAzureADAdministrator(name, id, options);
        }
    }

    public sealed class ServerAzureADAdministratorArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of server active directory administrator.
        /// </summary>
        [Input("administratorName", required: true)]
        public Input<string> AdministratorName { get; set; } = null!;

        /// <summary>
        /// Type of the sever administrator.
        /// </summary>
        [Input("administratorType", required: true)]
        public InputUnion<string, Pulumi.AzureNextGen.Sql.V20200801Preview.AdministratorType> AdministratorType { get; set; } = null!;

        /// <summary>
        /// Login name of the server administrator.
        /// </summary>
        [Input("login", required: true)]
        public Input<string> Login { get; set; } = null!;

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

        /// <summary>
        /// SID (object ID) of the server administrator.
        /// </summary>
        [Input("sid", required: true)]
        public Input<string> Sid { get; set; } = null!;

        /// <summary>
        /// Tenant ID of the administrator.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        public ServerAzureADAdministratorArgs()
        {
        }
    }
}
