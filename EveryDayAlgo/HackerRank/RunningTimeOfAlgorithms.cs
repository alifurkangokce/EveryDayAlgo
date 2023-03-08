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
//      * Complete the 'runningTime' function below.
//      *
//      * The function is expected to return an INTEGER.
//      * The function accepts INTEGER_ARRAY arr as parameter.
//      */

//     public static int runningTime(List<int> arr)
//     {
//         int ret = 0;

//         while (arr.Count != 0)
//         {
//             ret += arr.IndexOf(arr.Min());
//             arr.Remove(arr.Min());
//         }
//         return ret;
//     }


// }

// class Solution
// {
//     public static void Main(string[] args)
//     {
//         // TextWriter textWriter = new StreamWriter(@System.Environment.GetEnvironmentVariable("OUTPUT_PATH"), true);

//         //int n = 5;

//         List<int> arr = new List<int>(){
//             2,1,3,1,2
//         };

//         int result = Result.runningTime(arr);

//         // textWriter.WriteLine(result);

//         // textWriter.Flush();
//         // textWriter.Close();
//     }
// }
