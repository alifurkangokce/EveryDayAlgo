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
     * Complete the 'minimumNumber' function below.
     *
     * The function is expected to return an INTEGER.
     * The function accepts following parameters:
     *  1. INTEGER n
     *  2. STRING password
     */

    public static int minimumNumber(int n, string password)
    {
        int result=0;
        if (!password.Any(x=>x>=48&&x<=57))
        {
            result++;
        }
        if(!password.Any(x=>x>=97&&x<=122)){
            result++;
        }
        if(!password.Any(x=>x>=65&&x<=90)){
            result++;
        }
        if(!password.Any(x=>x>=33&&x<=46)){
            result++;
        }
        if (password.Length+result<6)
        {
            return 6-(password.Length);
        }else{
            return result;
        }
    }

}

class Solution
{
    public static void Main(string[] args)
    {
        // TextWriter textWriter = new StreamWriter(@System.Environment.GetEnvironmentVariable("OUTPUT_PATH"), true);

        // int n = Convert.ToInt32(Console.ReadLine().Trim());

        // string password = Console.ReadLine();

        // int answer = Result.minimumNumber(n, password);
        Result.minimumNumber(3,"Ab1");

        // textWriter.WriteLine(answer);

        // textWriter.Flush();
        // textWriter.Close();
    }
}
