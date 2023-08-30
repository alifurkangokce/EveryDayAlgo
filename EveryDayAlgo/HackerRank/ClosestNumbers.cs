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
//      * Complete the 'closestNumbers' function below.
//      *
//      * The function is expected to return an INTEGER_ARRAY.
//      * The function accepts INTEGER_ARRAY arr as parameter.
//      */

//     public static List<int> closestNumbers(List<int> arr)
//     {
//         var maxInt = int.MaxValue;
//         var resArr=new List<int>();

//         var sort=arr.OrderByDescending(x => x).ToList();
//         for (int i = 0; i < sort.Count-1; i++)
//         {
//             if (sort[i] - sort[i+1]<maxInt)
//             {
//                 resArr=new List<int>();
//                 maxInt = sort[i] - sort[i + 1];
//                 resArr.Add(sort[i]);
//                 resArr.Add(sort[i+1]);
//             }else if (sort[i] - sort[i + 1]==maxInt)
//             {
//                 resArr.Add(sort[i]);
//                 resArr.Add(sort[i + 1]);
//             }
//         }

//         return resArr.OrderBy(x=>x).ToList();   
//     }

// }

// class Solution
// {
//     public static void Main(string[] args)
//     {
//         TextWriter textWriter = new StreamWriter(@System.Environment.GetEnvironmentVariable("OUTPUT_PATH"), true);

//         int n = Convert.ToInt32(Console.ReadLine().Trim());

//         List<int> arr = Console.ReadLine().TrimEnd().Split(' ').ToList().Select(arrTemp => Convert.ToInt32(arrTemp)).ToList();

//         List<int> result = Result.closestNumbers(arr);

//         textWriter.WriteLine(String.Join(" ", result));

//         textWriter.Flush();
//         textWriter.Close();
//     }
// }
