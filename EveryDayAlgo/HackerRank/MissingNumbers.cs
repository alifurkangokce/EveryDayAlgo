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
//      * Complete the 'missingNumbers' function below.
//      *
//      * The function is expected to return an INTEGER_ARRAY.
//      * The function accepts following parameters:
//      *  1. INTEGER_ARRAY arr
//      *  2. INTEGER_ARRAY brr
//      */

//     public static List<int> missingNumbers(List<int> arr, List<int> brr)
//     {
//         var dic = new Dictionary<int, int>();
//         var res=new List<int>();
//         foreach (var item in brr)
//         {
//             if (!dic.ContainsKey(item))
//             {
//                 dic.Add(item, 1);
//             }
//             else
//             {
//                 dic[item]++;
//             }
//         }
//         foreach (var item in arr)
//         {
//             if (dic.ContainsKey(item))
//             {
//                 dic[item]--;
//             }
//         }
//         foreach (var item in dic)
//         {
//             if (item.Value!=0)
//             {
//                 res.Add(item.Key);
//             }
//         }
//         return res.OrderBy(x=>x).ToList();
//     }

// }

// class Solution
// {
//     public static void Main(string[] args)
//     {
//         //TextWriter textWriter = new StreamWriter(@System.Environment.GetEnvironmentVariable("OUTPUT_PATH"), true);

//         int n = Convert.ToInt32(Console.ReadLine().Trim());

//         List<int> arr = Console.ReadLine().TrimEnd().Split(' ').ToList().Select(arrTemp => Convert.ToInt32(arrTemp)).ToList();

//         int m = Convert.ToInt32(Console.ReadLine().Trim());

//         List<int> brr = Console.ReadLine().TrimEnd().Split(' ').ToList().Select(brrTemp => Convert.ToInt32(brrTemp)).ToList();

//         List<int> result = Result.missingNumbers(arr, brr);
//         Console.WriteLine(result);

//         //textWriter.WriteLine(String.Join(" ", result));

//         //textWriter.Flush();
//         //textWriter.Close();
//     }
// }
