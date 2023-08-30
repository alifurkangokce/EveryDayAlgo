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
//      * Complete the 'funnyString' function below.
//      *
//      * The function is expected to return a STRING.
//      * The function accepts STRING s as parameter.
//      */

//     public static string funnyString(string s)
//     {
//         var ss = s.Reverse().ToList();
//         for (int i = 0; i <s.Length-1; i++)
//         {
//             if (Math.Abs(s[i] - s[i + 1]) != Math.Abs(ss[i] - ss[i+1]))
//             {
//                 return "Not Funny";
//             }
//         }

//         return "Funny";
//     }

// }

// class Solution
// {
//     public static void Main(string[] args)
//     {
//         // TextWriter textWriter = new StreamWriter(@System.Environment.GetEnvironmentVariable("OUTPUT_PATH"), true);

//         // int q = Convert.ToInt32(Console.ReadLine().Trim());

//         // for (int qItr = 0; qItr < q; qItr++)
//         // {
//         //     string s = Console.ReadLine();

//         //     string result = Result.funnyString(s);

//         //     textWriter.WriteLine(result);
//         // }

//         // textWriter.Flush();
//         // textWriter.Close();
//         var res=Result.funnyString("bcxz");
//         Console.Write(res);
//     }
// }
