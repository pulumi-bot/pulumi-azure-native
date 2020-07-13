// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.BatchAI.Outputs
{

    [OutputType]
    public sealed class CaffeSettingsResponseResult
    {
        /// <summary>
        /// Command line arguments that need to be passed to the Caffe job.
        /// </summary>
        public readonly string? CommandLineArgs;
        /// <summary>
        /// Path of the config file for the job. This property cannot be specified if pythonScriptFilePath is specified.
        /// </summary>
        public readonly string? ConfigFilePath;
        /// <summary>
        /// Number of processes to launch for the job execution. The default value for this property is equal to nodeCount property
        /// </summary>
        public readonly int? ProcessCount;
        /// <summary>
        /// The path to the Python interpreter. The property can be specified only if the pythonScriptFilePath is specified.
        /// </summary>
        public readonly string? PythonInterpreterPath;
        /// <summary>
        /// Python script to execute. This property cannot be specified if configFilePath is specified.
        /// </summary>
        public readonly string? PythonScriptFilePath;

        [OutputConstructor]
        private CaffeSettingsResponseResult(
            string? commandLineArgs,

            string? configFilePath,

            int? processCount,

            string? pythonInterpreterPath,

            string? pythonScriptFilePath)
        {
            CommandLineArgs = commandLineArgs;
            ConfigFilePath = configFilePath;
            ProcessCount = processCount;
            PythonInterpreterPath = pythonInterpreterPath;
            PythonScriptFilePath = pythonScriptFilePath;
        }
    }
}