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
     * Complete the 'makingAnagrams' function below.
     *
     * The function is expected to return an INTEGER.
     * The function accepts following parameters:
     *  1. STRING s1
     *  2. STRING s2
     */

    public static int makingAnagrams(string s1, string s2)
    {
        var mm= new Dictionary<char, int>();
        var mm2=new Dictionary<char, int>();
        var res = 0;
        foreach (var item in s1)
        {
            if (mm.ContainsKey(item))
            {
                mm[item]++;
            }
            else
            {
                mm.Add(item, 1);
            }
            
        }
        foreach (var item in s2)
        {
            if (mm2.ContainsKey(item))
            {
                mm2[item]++;
            }
            else
            {
                mm2.Add(item, 1);
            }
        }
        foreach (var item in mm)
        {
            if (mm2.ContainsKey(item.Key))
            {
                res += Math.Abs(item.Value - mm2[item.Key]);
                mm2.Remove(item.Key);
            }
            else
            {
                res += item.Value;
            }
            mm.Remove(item.Key); 
        }
        foreach (var item in mm2)
        {
            res += item.Value;
        }
        return res;
    }

}

class Solution
{
    public static void Main(string[] args)
    {
        // TextWriter textWriter = new StreamWriter(@System.Environment.GetEnvironmentVariable("OUTPUT_PATH"), true);

        string s1 = Console.ReadLine();

        string s2 = Console.ReadLine();

        int result = Result.makingAnagrams(s1, s2);

        // textWriter.WriteLine(result);

        // textWriter.Flush();
        // textWriter.Close();
        Console.WriteLine(result);
    }
}
