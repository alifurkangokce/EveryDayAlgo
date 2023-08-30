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
//      * Complete the 'palindromeIndex' function below.
//      *
//      * The function is expected to return an INTEGER.
//      * The function accepts STRING s as parameter.
//      */

//     public static int palindromeIndex(string s)
//     {
//         string r = new string(s.Reverse().ToArray());
//         if (s == r) return -1;
//         int n = s.Length;
//         for (int i = 0; i < n / 2; i++)
//             if (s[i] != s[n - i - 1])
//             return (s[i] == s[n - i - 2] && s[i + 1] == s[n - i - 3]) ?
//                 n - i - 1 : i;
//         return -2;
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
//             string s = Console.ReadLine();

//             int result = Result.palindromeIndex(s);

//             textWriter.WriteLine(result);
//         }

//         textWriter.Flush();
//         textWriter.Close();
//     }
// }
