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
//      * Complete the 'caesarCipher' function below.
//      *
//      * The function is expected to return a STRING.
//      * The function accepts following parameters:
//      *  1. STRING s
//      *  2. INTEGER k
//      */

// public static string caesarCipher(string s, int k)
//     {
//         string alphabet = "abcdefghijklmnopqrstuvwxyz";
//         k = k % 26;
//         Regex re = new Regex("[a-zA-Z]");
//         string result = string.Empty;
        
//         for(int i = 0; i < s.Length; i++){
//             if(!re.IsMatch($"{s[i]}")) result += s[i];
//             else if(char.IsUpper(s[i])){
//                 result += char.ToUpper(alphabet[(alphabet.IndexOf(char.ToLower(s[i])) + k) % 26]);
//             } else {
//                 result += alphabet[(alphabet.IndexOf(s[i]) + k) % 26];
//             }
//         }
//         return result;
//     }

// }

// class Solution
// {
//     public static void Main(string[] args)
//     {
//         TextWriter textWriter = new StreamWriter(@System.Environment.GetEnvironmentVariable("OUTPUT_PATH"), true);

//         int n = Convert.ToInt32(Console.ReadLine().Trim());

//         string s = Console.ReadLine();

//         int k = Convert.ToInt32(Console.ReadLine().Trim());

//         string result = Result.caesarCipher(s, k);

//         textWriter.WriteLine(result);

//         textWriter.Flush();
//         textWriter.Close();
//     }
// }
