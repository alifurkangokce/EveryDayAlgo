// using System.CodeDom.Compiler;
// using System.Collections.Generic;
// using System.Collections;
// using System.ComponentModel;
// using System.Diagnostics.CodeAnalysis;
// using System.Globalization;
// using System.IO;
// using System.Linq;
// using System.Reflection;
// using System.Runtime.Serialization;
// using System.Text.RegularExpressions;
// using System.Text;
// using System;

// class Result
// {

//     /*
//      * Complete the 'organizingContainers' function below.
//      *
//      * The function is expected to return a STRING.
//      * The function accepts 2D_INTEGER_ARRAY container as parameter.
//      */

//     public static string organizingContainers(List<List<int>> container)
//     {
//         if (container.Any())
//         {
//             var containerCount = container.Count();
//             var ballTypeCount = new int[containerCount];
//             var containerCapacityCount = new int[containerCount];

//             for (int i = 0; i < containerCount; ++i)
//             {
//                 var row = container.ElementAt(i);
//                 for (int j = 0; j < row.Count(); ++j)
//                 {
//                     var ele = row.ElementAt(j);
//                     containerCapacityCount[i] += ele;
//                     ballTypeCount[j] += ele;
//                 }
//             }

//             Array.Sort(ballTypeCount);
//             Array.Sort(containerCapacityCount);

//             for (int i = 0; i < containerCount; ++i)
//             {
//                 if (ballTypeCount[i] > containerCapacityCount[i])
//                     return "Impossible";
//             }

//             return "Possible";
//         }

//         return "Impossible";
//     }

// }

// class Solution
// {
//     public static void Main(string[] args)
//     {
//         TextWriter textWriter = new StreamWriter(@System.Environment.GetEnvironmentVariable("OUTPUT_PATH"), true);

//         int q = Convert.ToInt32(Console.ReadLine().Trim());

//         for (int qItr = 0; qItr < q; qItr++)
//         {
//             int n = Convert.ToInt32(Console.ReadLine().Trim());

//             List<List<int>> container = new List<List<int>>();

//             for (int i = 0; i < n; i++)
//             {
//                 container.Add(Console.ReadLine().TrimEnd().Split(' ').ToList().Select(containerTemp => Convert.ToInt32(containerTemp)).ToList());
//             }

//             string result = Result.organizingContainers(container);

//             textWriter.WriteLine(result);
//         }

//         textWriter.Flush();
//         textWriter.Close();
//     }
// }
