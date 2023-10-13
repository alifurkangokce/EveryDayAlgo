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
//      * Complete the 'luckBalance' function below.
//      *
//      * The function is expected to return an INTEGER.
//      * The function accepts following parameters:
//      *  1. INTEGER k
//      *  2. 2D_INTEGER_ARRAY contests
//      */

//     public static int luckBalance(int k, List<List<int>> contests)
//     {
//         int luck_balance = 0;

//         var siralamaCiftleri = contests
//             .Select((liste, indeks) => new { Liste = liste, Indeks = indeks })
//             .ToList();

   
//         siralamaCiftleri.Sort((x, y) => y.Liste[0].CompareTo(x.Liste[0]));


//         for (int i = 0; i < siralamaCiftleri.Count; i++)
//         {
//             contests[i] = siralamaCiftleri[i].Liste;
//         }
//         for (int i = 0; i < contests.Count; i++)
//         {
//             int luck = contests[i][0];
//             int importance = contests[i][1];
//             if (importance==1 && k>0)
//             {
//                 k--;
//                 luck_balance += luck;
//             }else if (importance == 1 && k == 0)
//             {
//                 luck_balance-=luck;
//             }

//             if (importance==0)
//             {
//                 luck_balance += luck;
//             }
//         }
//         return luck_balance;
//     }

// }

// class Solution
// {
//     public static void Main(string[] args)
//     {
//         TextWriter textWriter = new StreamWriter(@System.Environment.GetEnvironmentVariable("OUTPUT_PATH"), true);

//         string[] firstMultipleInput = Console.ReadLine().TrimEnd().Split(' ');

//         int n = Convert.ToInt32(firstMultipleInput[0]);

//         int k = Convert.ToInt32(firstMultipleInput[1]);

//         List<List<int>> contests = new List<List<int>>();

//         for (int i = 0; i < n; i++)
//         {
//             contests.Add(Console.ReadLine().TrimEnd().Split(' ').ToList().Select(contestsTemp => Convert.ToInt32(contestsTemp)).ToList());
//         }

//         int result = Result.luckBalance(k, contests);

//         textWriter.WriteLine(result);

//         textWriter.Flush();
//         textWriter.Close();
//     }
// }
