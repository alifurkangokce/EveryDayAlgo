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
//      * Complete the 'jimOrders' function below.
//      *
//      * The function is expected to return an INTEGER_ARRAY.
//      * The function accepts 2D_INTEGER_ARRAY orders as parameter.
//      */

//     public static List<int> jimOrders(List<List<int>> orders)
//     {
//         var dic = new Dictionary<int, int>();
//         int i = 1;
//         foreach (var order in orders)
//         {
//             dic.Add(i++, order[0] + order[1]);
//         }
//         var sorted = dic.OrderBy(x => x.Value).ToDictionary(x=>x.Key,x=>x.Value);
        
//         var result=sorted.Keys.ToList();
//         return result;
//     }

// }

// class Solution
// {
//     public static void Main(string[] args)
//     {
//         //TextWriter textWriter = new StreamWriter(@System.Environment.GetEnvironmentVariable("OUTPUT_PATH"), true);

//         int n = Convert.ToInt32(Console.ReadLine().Trim());

//         List<List<int>> orders = new List<List<int>>();

//         for (int i = 0; i < n; i++)
//         {
//             orders.Add(Console.ReadLine().TrimEnd().Split(' ').ToList().Select(ordersTemp => Convert.ToInt32(ordersTemp)).ToList());
//         }

//         List<int> result = Result.jimOrders(orders);

//         Console.WriteLine(String.Join(" ", result));
//         //textWriter.WriteLine(String.Join(" ", result));

//         //textWriter.Flush();
//         //textWriter.Close();
//     }
// }