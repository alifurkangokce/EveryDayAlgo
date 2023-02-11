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
     * Complete the 'alternate' function below.
     *
     * The function is expected to return an INTEGER.
     * The function accepts STRING s as parameter.
     */

    public static int alternate(string s)
    {
        var distinctL=s.GroupBy(x=>x).ToList();
        int max=0;
        for (int i = 0; i <= distinctL.Count-1; i++)
        {
            for (int j = 1; j <=distinctL.Count-1; j++)
            {
                int res=alternateHelper(distinctL[i].First(),distinctL[j].First(),s);
                max=Math.Max(max,res);
                
            }
        }
        return max;
    }
    public static int alternateHelper(char ch1,char ch2,string s){
        int count=0;
        char lastChar='0';
        for (int i = 0; i <= s.Length-1; i++)
        {
            char ch=s[i];
            if (ch==ch1 || ch==ch2)
            {
                if (lastChar==ch)
                {
                    return 0;
                }
                lastChar=ch;
                count++;
            }
        }
        return count;
    }

}

class Solution
{
    public static void Main(string[] args)
    {
        // TextWriter textWriter = new StreamWriter(@System.Environment.GetEnvironmentVariable("OUTPUT_PATH"), true);

        int l = Convert.ToInt32(Console.ReadLine().Trim());

        string s = Console.ReadLine();

        int result = Result.alternate(s);

        // textWriter.WriteLine(result);

        // textWriter.Flush();
        // textWriter.Close();
        Console.Write(result);
    }
}
