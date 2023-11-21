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
//      * Complete the 'timeInWords' function below.
//      *
//      * The function is expected to return a STRING.
//      * The function accepts following parameters:
//      *  1. INTEGER h
//      *  2. INTEGER m
//      */

//     public static string timeInWords(int h, int m)
//     {

//         string? time;
//         var hours = new Dictionary<int, string>() { { 1, "one" }, { 2, "two" }, { 3, "three" }, { 4, "four" }, { 5, "five" }, { 6, "six" }, { 7, "seven" }, { 8, "eight" }, { 9, "nine" }, { 10, "ten" }, { 11, "eleven" }, { 12, "twelve" } };
//         var minutes = new Dictionary<int, string>() { { 1, "one" }, { 2, "two" }, { 3, "three" }, { 4, "four" }, { 5, "five" }, { 6, "six" }, { 7, "seven" }, { 8, "eight" }, { 9, "nine" }, { 10, "ten" }, { 11, "eleven" }, { 12, "twelve" }, { 13, "thirteen" }, { 14, "fourteen" }, { 15, "quarter" }, { 16, "sixteen" }, { 17, "seventeen" }, { 18, "eighteen" }, { 19, "nineteen" }, { 20, "twenty" }, { 30, "half" } };

//         var realM = m;

//         if (m >= 31)
//         {
//             m = 60 - m;

//             if (realM >= 30)
//                 h = (h == 12) ? 1 : h + 1;
//         }

//         var hrs = hours.GetValueOrDefault(h);
//         var min = string.Empty;

//         if ((realM >= 1 && realM <= 20) || (m >= 1 && m <= 20))
//             min = minutes.GetValueOrDefault(m);
//         else if ((21 <= realM && realM <= 29) || (21 <= m && m <= 29))
//         {

//             var first = minutes.GetValueOrDefault(20);
//             var next = Convert.ToInt32(m.ToString().Substring(1, 1));
//             min = $"{first} {minutes.GetValueOrDefault(next)}";
//         }
//         else if (realM == 30)
//             min = minutes.GetValueOrDefault(30);

//         var minutesWord = (realM == 1 || realM == 59) ? "minute" : "minutes";

//         if (realM == 0)
//             time = $"{hrs} o' clock";
//         else if (realM == 15 || realM == 30)
//             time = $"{min} past {hrs}";
//         else if (realM >= 1 && realM <= 29)
//             time = $"{min} {minutesWord} past {hrs}";
//         else if (m == 15)
//             time = $"{min} to {hrs}";
//         else
//             time = $"{min} {minutesWord} to {hrs}";

//         return time;
//     }

// }

// class Solution
// {
//     public static void Main(string[] args)
//     {
//         TextWriter textWriter = new StreamWriter(@System.Environment.GetEnvironmentVariable("OUTPUT_PATH"), true);

//         int h = Convert.ToInt32(Console.ReadLine().Trim());

//         int m = Convert.ToInt32(Console.ReadLine().Trim());

//         string result = Result.timeInWords(h, m);

//         textWriter.WriteLine(result);

//         textWriter.Flush();
//         textWriter.Close();
//     }
// }