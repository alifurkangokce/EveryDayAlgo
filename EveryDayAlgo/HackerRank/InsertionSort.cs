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
//      * Complete the 'insertionSort1' function below.
//      *
//      * The function accepts following parameters:
//      *  1. INTEGER n
//      *  2. INTEGER_ARRAY arr
//      */

//     public static void insertionSort1(int n, List<int> arr)
//     {
//             int j = 0, t;
            
//             for(int i = 0; i < n; i++){
//                 t = arr[i];
//                 for(j = i - 1; (j >= 0 && arr[j] > t); j--){
//                     arr[j + 1] = arr[j];
//                     for (int x = 0; x < n; x++)
//                     {
//                         Console.Write(arr[x]+" ");
//                     }
//                     Console.WriteLine("");
//                 }
//                 arr[j + 1] = t;
//             }
//             for (int y = 0; y < n; y++)
//             {
//                 Console.Write(arr[y]+" ");
//             }
//     }

// }

// class Solution
// {
//     public static void Main(string[] args)
//     {
//         int n = Convert.ToInt32(Console.ReadLine().Trim());

//         List<int> arr = Console.ReadLine().TrimEnd().Split(' ').ToList().Select(arrTemp => Convert.ToInt32(arrTemp)).ToList();

//         Result.insertionSort1(n, arr);
//     }
// }
