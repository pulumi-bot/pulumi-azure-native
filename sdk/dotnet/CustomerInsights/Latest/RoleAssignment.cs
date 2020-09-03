// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.CustomerInsights.Latest
{
    /// <summary>
    /// The Role Assignment resource format.
    /// 
    /// ## Example Usage
    /// ### RoleAssignments_CreateOrUpdate
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var roleAssignment = new AzureRM.CustomerInsights.Latest.RoleAssignment("roleAssignment", new AzureRM.CustomerInsights.Latest.RoleAssignmentArgs
    ///         {
    ///             AssignmentName = "assignmentName8976",
    ///             HubName = "sdkTestHub",
    ///             Principals = 
    ///             {
    ///                 new AzureRM.CustomerInsights.Latest.Inputs.AssignmentPrincipalArgs
    ///                 {
    ///                     PrincipalId = "4c54c38ffa9b416ba5a6d6c8a20cbe7e",
    ///                     PrincipalType = "User",
    ///                 },
    ///                 new AzureRM.CustomerInsights.Latest.Inputs.AssignmentPrincipalArgs
    ///                 {
    ///                     PrincipalId = "93061d15a5054f2b9948ae25724cf9d5",
    ///                     PrincipalType = "User",
    ///                 },
    ///             },
    ///             ResourceGroupName = "TestHubRG",
    ///             Role = "Admin",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class RoleAssignment : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the metadata object.
        /// </summary>
        [Output("assignmentName")]
        public Output<string> AssignmentName { get; private set; } = null!;

        /// <summary>
        /// Widget types set for the assignment.
        /// </summary>
        [Output("conflationPolicies")]
        public Output<Outputs.ResourceSetDescriptionResponseResult?> ConflationPolicies { get; private set; } = null!;

        /// <summary>
        /// Connectors set for the assignment.
        /// </summary>
        [Output("connectors")]
        public Output<Outputs.ResourceSetDescriptionResponseResult?> Connectors { get; private set; } = null!;

        /// <summary>
        /// Localized description for the metadata.
        /// </summary>
        [Output("description")]
        public Output<ImmutableDictionary<string, string>?> Description { get; private set; } = null!;

        /// <summary>
        /// Localized display names for the metadata.
        /// </summary>
        [Output("displayName")]
        public Output<ImmutableDictionary<string, string>?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Interactions set for the assignment.
        /// </summary>
        [Output("interactions")]
        public Output<Outputs.ResourceSetDescriptionResponseResult?> Interactions { get; private set; } = null!;

        /// <summary>
        /// Kpis set for the assignment.
        /// </summary>
        [Output("kpis")]
        public Output<Outputs.ResourceSetDescriptionResponseResult?> Kpis { get; private set; } = null!;

        /// <summary>
        /// Links set for the assignment.
        /// </summary>
        [Output("links")]
        public Output<Outputs.ResourceSetDescriptionResponseResult?> Links { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The principals being assigned to.
        /// </summary>
        [Output("principals")]
        public Output<ImmutableArray<Outputs.AssignmentPrincipalResponseResult>> Principals { get; private set; } = null!;

        /// <summary>
        /// Profiles set for the assignment.
        /// </summary>
        [Output("profiles")]
        public Output<Outputs.ResourceSetDescriptionResponseResult?> Profiles { get; private set; } = null!;

        /// <summary>
        /// Provisioning state.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The Role assignments set for the relationship links.
        /// </summary>
        [Output("relationshipLinks")]
        public Output<Outputs.ResourceSetDescriptionResponseResult?> RelationshipLinks { get; private set; } = null!;

        /// <summary>
        /// The Role assignments set for the relationships.
        /// </summary>
        [Output("relationships")]
        public Output<Outputs.ResourceSetDescriptionResponseResult?> Relationships { get; private set; } = null!;

        /// <summary>
        /// Type of roles.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;

        /// <summary>
        /// The Role assignments set for the assignment.
        /// </summary>
        [Output("roleAssignments")]
        public Output<Outputs.ResourceSetDescriptionResponseResult?> RoleAssignments { get; private set; } = null!;

        /// <summary>
        /// Sas Policies set for the assignment.
        /// </summary>
        [Output("sasPolicies")]
        public Output<Outputs.ResourceSetDescriptionResponseResult?> SasPolicies { get; private set; } = null!;

        /// <summary>
        /// The Role assignments set for the assignment.
        /// </summary>
        [Output("segments")]
        public Output<Outputs.ResourceSetDescriptionResponseResult?> Segments { get; private set; } = null!;

        /// <summary>
        /// The hub name.
        /// </summary>
        [Output("tenantId")]
        public Output<string> TenantId { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Views set for the assignment.
        /// </summary>
        [Output("views")]
        public Output<Outputs.ResourceSetDescriptionResponseResult?> Views { get; private set; } = null!;

        /// <summary>
        /// Widget types set for the assignment.
        /// </summary>
        [Output("widgetTypes")]
        public Output<Outputs.ResourceSetDescriptionResponseResult?> WidgetTypes { get; private set; } = null!;


        /// <summary>
        /// Create a RoleAssignment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RoleAssignment(string name, RoleAssignmentArgs args, CustomResourceOptions? options = null)
            : base("azurerm:customerinsights/latest:RoleAssignment", name, args ?? new RoleAssignmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RoleAssignment(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:customerinsights/latest:RoleAssignment", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:customerinsights/v20170101:RoleAssignment"},
                    new Pulumi.Alias { Type = "azurerm:customerinsights/v20170426:RoleAssignment"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing RoleAssignment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RoleAssignment Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new RoleAssignment(name, id, options);
        }
    }

    public sealed class RoleAssignmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The assignment name
        /// </summary>
        [Input("assignmentName", required: true)]
        public Input<string> AssignmentName { get; set; } = null!;

        /// <summary>
        /// Widget types set for the assignment.
        /// </summary>
        [Input("conflationPolicies")]
        public Input<Inputs.ResourceSetDescriptionArgs>? ConflationPolicies { get; set; }

        /// <summary>
        /// Connectors set for the assignment.
        /// </summary>
        [Input("connectors")]
        public Input<Inputs.ResourceSetDescriptionArgs>? Connectors { get; set; }

        [Input("description")]
        private InputMap<string>? _description;

        /// <summary>
        /// Localized description for the metadata.
        /// </summary>
        public InputMap<string> Description
        {
            get => _description ?? (_description = new InputMap<string>());
            set => _description = value;
        }

        [Input("displayName")]
        private InputMap<string>? _displayName;

        /// <summary>
        /// Localized display names for the metadata.
        /// </summary>
        public InputMap<string> DisplayName
        {
            get => _displayName ?? (_displayName = new InputMap<string>());
            set => _displayName = value;
        }

        /// <summary>
        /// The name of the hub.
        /// </summary>
        [Input("hubName", required: true)]
        public Input<string> HubName { get; set; } = null!;

        /// <summary>
        /// Interactions set for the assignment.
        /// </summary>
        [Input("interactions")]
        public Input<Inputs.ResourceSetDescriptionArgs>? Interactions { get; set; }

        /// <summary>
        /// Kpis set for the assignment.
        /// </summary>
        [Input("kpis")]
        public Input<Inputs.ResourceSetDescriptionArgs>? Kpis { get; set; }

        /// <summary>
        /// Links set for the assignment.
        /// </summary>
        [Input("links")]
        public Input<Inputs.ResourceSetDescriptionArgs>? Links { get; set; }

        [Input("principals", required: true)]
        private InputList<Inputs.AssignmentPrincipalArgs>? _principals;

        /// <summary>
        /// The principals being assigned to.
        /// </summary>
        public InputList<Inputs.AssignmentPrincipalArgs> Principals
        {
            get => _principals ?? (_principals = new InputList<Inputs.AssignmentPrincipalArgs>());
            set => _principals = value;
        }

        /// <summary>
        /// Profiles set for the assignment.
        /// </summary>
        [Input("profiles")]
        public Input<Inputs.ResourceSetDescriptionArgs>? Profiles { get; set; }

        /// <summary>
        /// The Role assignments set for the relationship links.
        /// </summary>
        [Input("relationshipLinks")]
        public Input<Inputs.ResourceSetDescriptionArgs>? RelationshipLinks { get; set; }

        /// <summary>
        /// The Role assignments set for the relationships.
        /// </summary>
        [Input("relationships")]
        public Input<Inputs.ResourceSetDescriptionArgs>? Relationships { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Type of roles.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        /// <summary>
        /// The Role assignments set for the assignment.
        /// </summary>
        [Input("roleAssignments")]
        public Input<Inputs.ResourceSetDescriptionArgs>? RoleAssignments { get; set; }

        /// <summary>
        /// Sas Policies set for the assignment.
        /// </summary>
        [Input("sasPolicies")]
        public Input<Inputs.ResourceSetDescriptionArgs>? SasPolicies { get; set; }

        /// <summary>
        /// The Role assignments set for the assignment.
        /// </summary>
        [Input("segments")]
        public Input<Inputs.ResourceSetDescriptionArgs>? Segments { get; set; }

        /// <summary>
        /// Views set for the assignment.
        /// </summary>
        [Input("views")]
        public Input<Inputs.ResourceSetDescriptionArgs>? Views { get; set; }

        /// <summary>
        /// Widget types set for the assignment.
        /// </summary>
        [Input("widgetTypes")]
        public Input<Inputs.ResourceSetDescriptionArgs>? WidgetTypes { get; set; }

        public RoleAssignmentArgs()
        {
        }
    }
}
