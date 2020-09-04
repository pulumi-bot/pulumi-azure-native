// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.OperationsManagement.V20151101Preview.Inputs
{

    /// <summary>
    /// ManagementConfiguration properties supported by the OperationsManagement resource provider.
    /// </summary>
    public sealed class ManagementConfigurationPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The applicationId of the appliance for this Management.
        /// </summary>
        [Input("applicationId")]
        public Input<string>? ApplicationId { get; set; }

        [Input("parameters", required: true)]
        private InputList<Inputs.ArmTemplateParameterArgs>? _parameters;

        /// <summary>
        /// Parameters to run the ARM template
        /// </summary>
        public InputList<Inputs.ArmTemplateParameterArgs> Parameters
        {
            get => _parameters ?? (_parameters = new InputList<Inputs.ArmTemplateParameterArgs>());
            set => _parameters = value;
        }

        /// <summary>
        /// The type of the parent resource.
        /// </summary>
        [Input("parentResourceType", required: true)]
        public Input<string> ParentResourceType { get; set; } = null!;

        [Input("template", required: true)]
        private InputMap<object>? _template;

        /// <summary>
        /// The Json object containing the ARM template to deploy
        /// </summary>
        public InputMap<object> Template
        {
            get => _template ?? (_template = new InputMap<object>());
            set => _template = value;
        }

        public ManagementConfigurationPropertiesArgs()
        {
        }
    }
}
