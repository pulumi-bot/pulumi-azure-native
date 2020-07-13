// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Authorization.Inputs
{

    /// <summary>
    /// The policy definition reference.
    /// </summary>
    public sealed class PolicyDefinitionReferenceArgs : Pulumi.ResourceArgs
    {
        [Input("groupNames")]
        private InputList<string>? _groupNames;

        /// <summary>
        /// The name of the groups that this policy definition reference belongs to.
        /// </summary>
        public InputList<string> GroupNames
        {
            get => _groupNames ?? (_groupNames = new InputList<string>());
            set => _groupNames = value;
        }

        /// <summary>
        /// The parameter values for the referenced policy rule. The keys are the parameter names.
        /// </summary>
        [Input("parameters")]
        public Input<Inputs.ParameterValuesArgs>? Parameters { get; set; }

        /// <summary>
        /// The ID of the policy definition or policy set definition.
        /// </summary>
        [Input("policyDefinitionId", required: true)]
        public Input<string> PolicyDefinitionId { get; set; } = null!;

        /// <summary>
        /// A unique id (within the policy set definition) for this policy definition reference.
        /// </summary>
        [Input("policyDefinitionReferenceId")]
        public Input<string>? PolicyDefinitionReferenceId { get; set; }

        public PolicyDefinitionReferenceArgs()
        {
        }
    }
}