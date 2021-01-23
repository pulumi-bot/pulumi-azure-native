// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Synapse.V20201201.Outputs
{

    [OutputType]
    public sealed class EnvironmentVariableSetupResponse
    {
        /// <summary>
        /// The type of custom setup.
        /// Expected value is 'EnvironmentVariableSetup'.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The name of the environment variable.
        /// </summary>
        public readonly string VariableName;
        /// <summary>
        /// The value of the environment variable.
        /// </summary>
        public readonly string VariableValue;

        [OutputConstructor]
        private EnvironmentVariableSetupResponse(
            string type,

            string variableName,

            string variableValue)
        {
            Type = type;
            VariableName = variableName;
            VariableValue = variableValue;
        }
    }
}
