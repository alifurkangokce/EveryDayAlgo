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
//      * Complete the 'anagram' function below.
//      *
//      * The function is expected to return an INTEGER.
//      * The function accepts STRING s as parameter.
//      */

//     public static int anagram(string s)
//     {
//         int n=s.Length;
//         if (n%2!=0)
//         {
//             return -1;
//         }
//         StringBuilder sb=new StringBuilder(s,n/2,n/2,n/2);
//         for (int i = 0; i < n/2; i++)
//         {
//             for (int j = 0; j < sb.Length; j++)
//             {
//                 if (s[i]==s[j])
//                 {
//                     sb.Remove(j,1);
//                     break;
//                 }
//             }
//         }
//         return sb.Length;
//     }

// }

// class Solution
// {
//     public static void Main(string[] args)
//     {
//         // TextWriter textWriter = new StreamWriter(@System.Environment.GetEnvironmentVariable("OUTPUT_PATH"), true);

//         int q = Convert.ToInt32(Console.ReadLine().Trim());

//         for (int qItr = 0; qItr < q; qItr++)
//         {
//             string s = Console.ReadLine();

//             int result = Result.anagram(s);

//             // textWriter.WriteLine(result);
//             Console.WriteLine(result);
//         }

//         // textWriter.Flush();
//         // textWriter.Close();
//     }
// }
 