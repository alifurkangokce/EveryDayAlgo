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
//      * Complete the 'gridChallenge' function below.
//      *
//      * The function is expected to return a STRING.
//      * The function accepts STRING_ARRAY grid as parameter.
//      */

//     public static string gridChallenge(List<string> grid)
//     {
//         for (int i = 0; i <grid.Count-1; i++)
//         {
//             char[] charX = grid[i].ToCharArray();

//             Array.Sort(charX);
//             grid[i] = new string(charX);

//             char[] charY = grid[i+1].ToCharArray();

//             Array.Sort(charY);
//             grid[i+1] = new string(charY);
//             for (int j = 0; j < grid[0].Length; j++)
//             {

//                 if (grid[i][j] > grid[i + 1][j])
//                 {
//                     return "NO";
//                 }
//             }
//         }

//         return "YES";
//     }

// }

// class Solution
// {
//     public static void Main(string[] args)
//     {
//         TextWriter textWriter = new StreamWriter(@System.Environment.GetEnvironmentVariable("OUTPUT_PATH"), true);

//         int t = Convert.ToInt32(Console.ReadLine().Trim());

//         for (int tItr = 0; tItr < t; tItr++)
//         {
//             int n = Convert.ToInt32(Console.ReadLine().Trim());

//             List<string> grid = new List<string>();

//             for (int i = 0; i < n; i++)
//             {
//                 string gridItem = Console.ReadLine();
//                 grid.Add(gridItem);
//             }

//             string result = Result.gridChallenge(grid);

//             //textWriter.WriteLine(result);
//             Console.WriteLine(result);
//         }

//         //textWriter.Flush();
//         //textWriter.Close();
//     }
// }
