/*
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
using System.Numerics;

class Result
{

    /*
     * Complete the 'separateNumbers' function below.
     *
     * The function accepts STRING s as parameter.
     #1#

        public static void separateNumbers(string s)
        {
            string newS = "";
            for (int i = 0; i < s.Length - 1; i++)
            {
                newS = "";
                BigInteger returnInt = BigInteger.Parse(new String(s.Skip(0).Take(i + 1).ToArray()));
                BigInteger firstInt = BigInteger.Parse(new String(s.Skip(0).Take(i + 1).ToArray()));
                while (newS.Length < s.Length)
                {
                    BigInteger nextInt = firstInt++;
                    newS += nextInt.ToString();
                }
                if (s == newS)
                {
                    Console.WriteLine("YES " + returnInt.ToString());
                    break;
                }
            }
            if (s != newS)
            {
                Console.WriteLine("NO");
            }
        }

}

class Solution
{
    public static void Main(string[] args)
    {
        // int q = Convert.ToInt32(Console.ReadLine().Trim());

        // for (int qItr = 0; qItr < q; qItr++)
        // {
        //     string s = Console.ReadLine();

        //     Result.separateNumbers(s);
        // }
        
        Result.separateNumbers("99910001001");
        
    }
}
*/
