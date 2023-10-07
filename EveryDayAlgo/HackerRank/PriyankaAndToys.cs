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
//      * Complete the 'toys' function below.
//      *
//      * The function is expected to return an INTEGER.
//      * The function accepts INTEGER_ARRAY w as parameter.
//      */

//     public static int toys(List<int> w)
//     {
//         w = w.OrderBy(x=>x).ToList();
//         var ff = w[0] + 4;
//         var cnt = 1;
//         for (int i = 1; i <=w.Count-1; i++)
//         {
//             if (w[i]>ff)
//             {
//                 ff = w[i] + 4;
//                 cnt++;
//             }
//         }

//         return cnt;

//     }

// }

// class Solution
// {
//     public static void Main(string[] args)
//     {
//         //TextWriter textWriter = new StreamWriter(@System.Environment.GetEnvironmentVariable("OUTPUT_PATH"), true);

//         int n = Convert.ToInt32(Console.ReadLine().Trim());

//         List<int> w = Console.ReadLine().TrimEnd().Split(' ').ToList().Select(wTemp => Convert.ToInt32(wTemp)).ToList();

//         int result = Result.toys(w);
//         Console.WriteLine(result);

//         //textWriter.WriteLine(result);

//         //textWriter.Flush();
//         //textWriter.Close();
//     }
// }