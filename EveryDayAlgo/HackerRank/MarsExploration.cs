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
     * Complete the 'marsExploration' function below.
     *
     * The function is expected to return an INTEGER.
     * The function accepts STRING s as parameter.
     */

    public static int marsExploration(string s)
    {
        int counter=0;
        for (int i = 0; i < s.Length/3; i++)
        {
            if (s[3*i]!='S')  
            {
                counter++;
            }
            if ( s[(3*i)+1]!='O')
            {
                counter++;
            }
            if ( s[(3*i)+2]!='S')
            {
                counter++;
            }
        }
        return counter;
    }

}

class Solution
{
    public static void Main(string[] args)
    {
        // TextWriter textWriter = new StreamWriter(@System.Environment.GetEnvironmentVariable("OUTPUT_PATH"), true);

        //string s = Console.ReadLine();
        string s="SOSOOSOSOSOSOSSOSOSOSOSOSOS";
        int result = Result.marsExploration(s);

        // textWriter.WriteLine(result);

        // textWriter.Flush();
        // textWriter.Close();
        Console.Write(result);
    }
}
