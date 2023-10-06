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
//      * Complete the 'decentNumber' function below.
//      *
//      * The function accepts INTEGER n as parameter.
//      */

//     public static void decentNumber(int n)
//     {
//         StringBuilder sb = new StringBuilder();

//         int m = n;
//         while (m % 3 != 0)
//         {
//             m = m - 5;
//         }

//         if (m < 0) sb = sb.Append(-1);

//         else
//         {
//             for (int i = 0; i < m; i++)
//             {
//                 sb.Append(5);
//             }

//             for (int i = 0; i < n - m; i++)
//             {
//                 sb.Append(3);
//             }
//         }
//         Console.WriteLine(sb.ToString());
       
//     }

// }

// class Solution
// {
//     public static void Main(string[] args)
//     {
//         int t = Convert.ToInt32(Console.ReadLine().Trim());

//         for (int tItr = 0; tItr < t; tItr++)
//         {
//             int n = Convert.ToInt32(Console.ReadLine().Trim());

//             Result.decentNumber(n);
//         }
//     }
// }