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
//     public static long strangeCounter(long t)
//     {
//         long startTime = 1;
//         long startValue = 3;
//         while(t>=(startTime*2)+2){
//             startTime=(startTime*2)+2;
//             startValue*=2;
//         }

//         long offset = (t - startTime);
        
//         return (startValue - offset);
//     }

// }

// class Solution
// {
//     public static void Main(string[] args)
//     {
//         TextWriter textWriter = new StreamWriter(@System.Environment.GetEnvironmentVariable("OUTPUT_PATH"), true);

//         long t = Convert.ToInt64(Console.ReadLine().Trim());

//         long result = Result.strangeCounter(t);

//         textWriter.WriteLine(result);

//         textWriter.Flush();
//         textWriter.Close();
//     }
// }
