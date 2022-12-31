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
     * Complete the 'happyLadybugs' function below.
     *
     * The function is expected to return a STRING.
     * The function accepts STRING b as parameter.
     */

    public static string happyLadybugs(string b)
    {
        var dic=new Dictionary<char,int>();
        if (b.Contains('_'))
        {
            for (int i = 0; i <= b.Length-1; i++)
            {
                if (b[i]!='_')
                {
                    if (!dic.ContainsKey(b[i]))
                    {
                         dic.Add(b[i],1);
                    }else{
                        dic[b[i]]++;
                    }
                   
                }
            }
            foreach (var item in dic)
            {
                if (item.Value<=1)
                {
                    return "NO";
                }
            }
        }else{
            if (b.Length==1)
            {
                return "NO";
            }
            for (int i = 1; i <= b.Length-1; i++)
            {
                if (b[i]!=b[i-1] && i!=b.Length-1)
                {
                    if (b[i]!=b[i+1])
                    {
                        return "NO";
                    }
                }else if(b[i]!=b[i-1] && i==b.Length-1){
                    return "NO";
                }
            }
        }
        return "YES";

    }

}

class Solution
{
    public static void Main(string[] args)
    {
        // TextWriter textWriter = new StreamWriter(@System.Environment.GetEnvironmentVariable("OUTPUT_PATH"), true);

        int g = Convert.ToInt32(Console.ReadLine().Trim());

        for (int gItr = 0; gItr < g; gItr++)
        {
            int n = Convert.ToInt32(Console.ReadLine().Trim());

            string b = Console.ReadLine();

            string result = Result.happyLadybugs(b);

            // textWriter.WriteLine(result);
            Console.WriteLine(result);
        }

        // textWriter.Flush();
        // textWriter.Close();
    }
}
