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
//      * Complete the 'stones' function below.
//      *
//      * The function is expected to return an INTEGER_ARRAY.
//      * The function accepts following parameters:
//      *  1. INTEGER n
//      *  2. INTEGER a
//      *  3. INTEGER b
//      */

//     public static List<int> stones(int n, int a, int b)
//     {
//          n -= 1;
//         List<int> result = new List<int>();
//         result.Add(a * n);
//         result.Add(b * n);
//         for(int i = 1; i <= (n/2); i++){
//             result.Add((a * i) + (b * (n - i)));
//             result.Add((b * i) + (a * (n - i)));
//         }
//         result.Sort();
//         return result.Distinct().ToList();
//     }

// }

// class Solution
// {
//     public static void Main(string[] args)
//     {
//         // TextWriter textWriter = new StreamWriter(@System.Environment.GetEnvironmentVariable("OUTPUT_PATH"), true);

//         int T = Convert.ToInt32(Console.ReadLine().Trim());

//         for (int TItr = 0; TItr < T; TItr++)
//         {
//             int n = Convert.ToInt32(Console.ReadLine().Trim());

//             int a = Convert.ToInt32(Console.ReadLine().Trim());

//             int b = Convert.ToInt32(Console.ReadLine().Trim());

//             List<int> result = Result.stones(n, a, b);

//             // textWriter.WriteLine(String.Join(" ", result));
//             Console.WriteLine(result);
//         }

//         // textWriter.Flush();
//         // textWriter.Close();
//     }
// }
