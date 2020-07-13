// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.BatchAI.Inputs
{

    /// <summary>
    /// Specifies a setup task which can be used to customize the compute nodes of the cluster.
    /// </summary>
    public sealed class SetupTaskArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The command line to be executed on each cluster's node after it being allocated or rebooted. The command is executed in a bash subshell as a root.
        /// </summary>
        [Input("commandLine", required: true)]
        public Input<string> CommandLine { get; set; } = null!;

        [Input("environmentVariables")]
        private InputList<Inputs.EnvironmentVariableArgs>? _environmentVariables;

        /// <summary>
        /// A collection of user defined environment variables to be set for setup task.
        /// </summary>
        public InputList<Inputs.EnvironmentVariableArgs> EnvironmentVariables
        {
            get => _environmentVariables ?? (_environmentVariables = new InputList<Inputs.EnvironmentVariableArgs>());
            set => _environmentVariables = value;
        }

        [Input("secrets")]
        private InputList<Inputs.EnvironmentVariableWithSecretValueArgs>? _secrets;

        /// <summary>
        /// A collection of user defined environment variables with secret values to be set for the setup task. Server will never report values of these variables back.
        /// </summary>
        public InputList<Inputs.EnvironmentVariableWithSecretValueArgs> Secrets
        {
            get => _secrets ?? (_secrets = new InputList<Inputs.EnvironmentVariableWithSecretValueArgs>());
            set => _secrets = value;
        }

        /// <summary>
        /// The prefix of a path where the Batch AI service will upload the stdout, stderr and execution log of the setup task.
        /// </summary>
        [Input("stdOutErrPathPrefix", required: true)]
        public Input<string> StdOutErrPathPrefix { get; set; } = null!;

        public SetupTaskArgs()
        {
        }
    }
}