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
    public sealed class CNTKsettingsResponse
    {
        /// <summary>
        /// Command line arguments that need to be passed to the python script or cntk executable.
        /// </summary>
        public readonly string? CommandLineArgs;
        /// <summary>
        /// Specifies the path of the BrainScript config file. This property can be specified only if the languageType is 'BrainScript'.
        /// </summary>
        public readonly string? ConfigFilePath;
        /// <summary>
        /// The language to use for launching CNTK (aka Microsoft Cognitive Toolkit) job. Valid values are 'BrainScript' or 'Python'.
        /// </summary>
        public readonly string? LanguageType;
        /// <summary>
        /// Number of processes to launch for the job execution. The default value for this property is equal to nodeCount property
        /// </summary>
        public readonly int? ProcessCount;
        /// <summary>
        /// The path to the Python interpreter. This property can be specified only if the languageType is 'Python'.
        /// </summary>
        public readonly string? PythonInterpreterPath;
        /// <summary>
        /// Python script to execute. This property can be specified only if the languageType is 'Python'.
        /// </summary>
        public readonly string? PythonScriptFilePath;

        [OutputConstructor]
        private CNTKsettingsResponse(
            string? commandLineArgs,

            string? configFilePath,

            string? languageType,

            int? processCount,

            string? pythonInterpreterPath,

            string? pythonScriptFilePath)
        {
            CommandLineArgs = commandLineArgs;
            ConfigFilePath = configFilePath;
            LanguageType = languageType;
            ProcessCount = processCount;
            PythonInterpreterPath = pythonInterpreterPath;
            PythonScriptFilePath = pythonScriptFilePath;
        }
    }
}