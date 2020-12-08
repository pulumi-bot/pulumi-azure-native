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
    /// An Azure SQL managed instance administrator.
    /// </summary>
    public partial class ManagedInstanceAdministrator : Pulumi.CustomResource
    {
        /// <summary>
        /// Type of the managed instance administrator.
        /// </summary>
        [Output("administratorType")]
        public Output<string> AdministratorType { get; private set; } = null!;

        /// <summary>
        /// Login name of the managed instance administrator.
        /// </summary>
        [Output("login")]
        public Output<string> Login { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// SID (object ID) of the managed instance administrator.
        /// </summary>
        [Output("sid")]
        public Output<string> Sid { get; private set; } = null!;

        /// <summary>
        /// Tenant ID of the managed instance administrator.
        /// </summary>
        [Output("tenantId")]
        public Output<string?> TenantId { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ManagedInstanceAdministrator resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ManagedInstanceAdministrator(string name, ManagedInstanceAdministratorArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:sql/v20200801preview:ManagedInstanceAdministrator", name, args ?? new ManagedInstanceAdministratorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ManagedInstanceAdministrator(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:sql/v20200801preview:ManagedInstanceAdministrator", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:sql/v20170301preview:ManagedInstanceAdministrator"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ManagedInstanceAdministrator resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ManagedInstanceAdministrator Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ManagedInstanceAdministrator(name, id, options);
        }
    }

    public sealed class ManagedInstanceAdministratorArgs : Pulumi.ResourceArgs
    {
        [Input("administratorName", required: true)]
        public Input<string> AdministratorName { get; set; } = null!;

        /// <summary>
        /// Type of the managed instance administrator.
        /// </summary>
        [Input("administratorType", required: true)]
        public InputUnion<string, Pulumi.AzureNextGen.Sql.V20200801Preview.ManagedInstanceAdministratorType> AdministratorType { get; set; } = null!;

        /// <summary>
        /// Login name of the managed instance administrator.
        /// </summary>
        [Input("login", required: true)]
        public Input<string> Login { get; set; } = null!;

        /// <summary>
        /// The name of the managed instance.
        /// </summary>
        [Input("managedInstanceName", required: true)]
        public Input<string> ManagedInstanceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// SID (object ID) of the managed instance administrator.
        /// </summary>
        [Input("sid", required: true)]
        public Input<string> Sid { get; set; } = null!;

        /// <summary>
        /// Tenant ID of the managed instance administrator.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        public ManagedInstanceAdministratorArgs()
        {
        }
    }
}
