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
//      * Complete the 'beautifulBinaryString' function below.
//      *
//      * The function is expected to return an INTEGER.
//      * The function accepts STRING b as parameter.
//      */

//     public static int beautifulBinaryString(string b)
//     {
//         var counter = 0;
//         for (int i = 0; i <= b.Length-3; i++)
//         {
//             if (b[i]=='0' && b[i+1]=='1' && b[i+2]=='0')
//             {
//                 counter++;
//                 i += 2;
//             }
//         }
//         return counter;
//     }

// }

// class Solution
// {
//     public static void Main(string[] args)
//     {
//         TextWriter textWriter = new StreamWriter(@System.Environment.GetEnvironmentVariable("OUTPUT_PATH"), true);

//         int n = Convert.ToInt32(Console.ReadLine().Trim());

//         string b = Console.ReadLine();

//         int result = Result.beautifulBinaryString(b);

//         textWriter.WriteLine(result);

//         textWriter.Flush();
//         textWriter.Close();
//     }
// }
