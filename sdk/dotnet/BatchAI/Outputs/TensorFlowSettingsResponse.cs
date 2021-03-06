// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.BatchAI.Outputs
{

    [OutputType]
    public sealed class TensorFlowSettingsResponse
    {
        /// <summary>
        /// Command line arguments that need to be passed to the python script for the master task.
        /// </summary>
        public readonly string? MasterCommandLineArgs;
        /// <summary>
        /// Command line arguments that need to be passed to the python script for the parameter server. Optional for single process jobs.
        /// </summary>
        public readonly string? ParameterServerCommandLineArgs;
        /// <summary>
        /// The number of parameter server tasks. If specified, the value must be less than or equal to nodeCount. If not specified, the default value is equal to 1 for distributed TensorFlow training. This property can be specified only for distributed TensorFlow training.
        /// </summary>
        public readonly int? ParameterServerCount;
        /// <summary>
        /// The path to the Python interpreter.
        /// </summary>
        public readonly string? PythonInterpreterPath;
        /// <summary>
        /// The python script to execute.
        /// </summary>
        public readonly string PythonScriptFilePath;
        /// <summary>
        /// Command line arguments that need to be passed to the python script for the worker task. Optional for single process jobs.
        /// </summary>
        public readonly string? WorkerCommandLineArgs;
        /// <summary>
        /// The number of worker tasks. If specified, the value must be less than or equal to (nodeCount * numberOfGPUs per VM). If not specified, the default value is equal to nodeCount. This property can be specified only for distributed TensorFlow training.
        /// </summary>
        public readonly int? WorkerCount;

        [OutputConstructor]
        private TensorFlowSettingsResponse(
            string? masterCommandLineArgs,

            string? parameterServerCommandLineArgs,

            int? parameterServerCount,

            string? pythonInterpreterPath,

            string pythonScriptFilePath,

            string? workerCommandLineArgs,

            int? workerCount)
        {
            MasterCommandLineArgs = masterCommandLineArgs;
            ParameterServerCommandLineArgs = parameterServerCommandLineArgs;
            ParameterServerCount = parameterServerCount;
            PythonInterpreterPath = pythonInterpreterPath;
            PythonScriptFilePath = pythonScriptFilePath;
            WorkerCommandLineArgs = workerCommandLineArgs;
            WorkerCount = workerCount;
        }
    }
}
