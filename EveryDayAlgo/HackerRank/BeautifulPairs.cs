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
//      * Complete the 'beautifulPairs' function below.
//      *
//      * The function is expected to return an INTEGER.
//      * The function accepts following parameters:
//      *  1. INTEGER_ARRAY A
//      *  2. INTEGER_ARRAY B
//      */

//     public static int beautifulPairs(List<int> A, List<int> B)
//     {
//         Dictionary<int,int> dc=new Dictionary<int,int>();
//         for (int i = 0; i <A.Count; i++)
//         {
//             if (!dc.ContainsKey(A[i]))
//             {
//                 dc.Add(A[i],1);
//             }
//             else
//             {
//                 dc[A[i]]++;
//             }
//         }
//         var res = 0;

//         for (int i = 0; i < B.Count; i++)
//         {
//             if (dc.ContainsKey(B[i]))
//             {
//                 if (dc[B[i]]>0)
//                 {
//                     dc[B[i]]--;
//                     res++;
//                 }
//             }

//         }


//         return res == A.Count ? res - 1 : res + 1;
//            }

// }

// class Solution
// {
//     public static void Main(string[] args)
//     {
//         TextWriter textWriter = new StreamWriter(@System.Environment.GetEnvironmentVariable("OUTPUT_PATH"), true);

//         int n = Convert.ToInt32(Console.ReadLine().Trim());

//         List<int> A = Console.ReadLine().TrimEnd().Split(' ').ToList().Select(ATemp => Convert.ToInt32(ATemp)).ToList();

//         List<int> B = Console.ReadLine().TrimEnd().Split(' ').ToList().Select(BTemp => Convert.ToInt32(BTemp)).ToList();

//         int result = Result.beautifulPairs(A, B);

//         textWriter.WriteLine(result);

//         textWriter.Flush();
//         textWriter.Close();
//     }
// }
