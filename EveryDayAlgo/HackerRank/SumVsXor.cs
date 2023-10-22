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
// using System.Runtime.InteropServices;

// class Result
// {

//     /*
//      * Complete the 'sumXor' function below.
//      *
//      * The function is expected to return a LONG_INTEGER.
//      * The function accepts LONG_INTEGER n as parameter.
//      */

//     public static long sumXor(long n)
//     {
//         string bin = Convert.ToString(n, 2);
//         int zeroCount = 0;
//         if (bin.Length == 1)
//         {
//             return 1;
//         }
//         for (int i = 0; i < bin.Length; i++)
//         {

//             if (bin[i] == '0')
//             {
//                 zeroCount++;
//             }
//         }

//         if (zeroCount <= 1)
//         {
//             return 2;
//         }
//         long shifted = 2;
//         return shifted << zeroCount - 1;
//     }

// }

// class Solution
// {
//     public static void Main(string[] args)
//     {
//         //TextWriter textWriter = new StreamWriter(@System.Environment.GetEnvironmentVariable("OUTPUT_PATH"), true);

//         long n = Convert.ToInt64(Console.ReadLine().Trim());

//         long result = Result.sumXor(n);
//         Console.WriteLine(result);

//         //textWriter.WriteLine(result);

//         //textWriter.Flush();
//         //textWriter.Close();
//     }
// }