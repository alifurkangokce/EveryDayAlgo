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
     * Complete the 'superReducedString' function below.
     *
     * The function is expected to return a STRING.
     * The function accepts STRING s as parameter.
     */

    public static string superReducedString(string s)
    {
        var sb = new StringBuilder(s);
        var i = 0;

        while (i < sb.Length - 1)
        {
            if (sb[i] == sb[i + 1])
            {
                sb = sb.Remove(i, 2);
                i = 0;
                continue;
            }
            i++;
        }

        return sb.Length == 0 ? "Empty String" : sb.ToString();

    }

}

class Solution
{
    public static void Main(string[] args)
    {
        // TextWriter textWriter = new StreamWriter(@System.Environment.GetEnvironmentVariable("OUTPUT_PATH"), true);

        string s = Console.ReadLine();

        string result = Result.superReducedString(s);

        // textWriter.WriteLine(result);

        // textWriter.Flush();
        // textWriter.Close();
        Console.WriteLine(result);
    }
}
