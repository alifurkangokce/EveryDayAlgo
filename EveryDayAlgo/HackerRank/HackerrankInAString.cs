using System.CodeDom.Compiler;
using System.Collections.Generic;
using System.Collections;
using System.ComponentModel;
using System.Diagnostics.CodeAnalysis;
using System.Globalization;
using System.IO;
using System.Linq;
using System.Reflection;
using System.Runtime.Serialization;
using System.Text.RegularExpressions;
using System.Text;
using System;

class Result
{

    /*
     * Complete the 'hackerrankInString' function below.
     *
     * The function is expected to return a STRING.
     * The function accepts STRING s as parameter.
     */

    public static string hackerrankInString(string s)
    {
        List<char> hckrnk=new List<char>() {
            'h','a','c','k','e','r','r','a','n','k'
        };
        for (int i = 0; i <= s.Length-1; i++)
        {
            if (hckrnk[0]==s[i])
            {
                hckrnk.RemoveAt(0);
                if (!hckrnk.Any())
                {
                    return "YES";
                }
            }
        }

        return "NO";
    }

}

class Solution
{
    public static void Main(string[] args)
    {
        // TextWriter textWriter = new StreamWriter(@System.Environment.GetEnvironmentVariable("OUTPUT_PATH"), true);

        // int q = Convert.ToInt32(Console.ReadLine().Trim());

        // for (int qItr = 0; qItr < q; qItr++)
        // {
        //     string s = Console.ReadLine();

        //     string result = Result.hackerrankInString(s);

        //     textWriter.WriteLine(result);
        // }

        // textWriter.Flush();
        // textWriter.Close();
        Result.hackerrankInString("hereiamstackerrank");
    }
}
