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
     * Complete the 'cavityMap' function below.
     *
     * The function is expected to return a STRING_ARRAY.
     * The function accepts STRING_ARRAY grid as parameter.
     */

    public static List<string> cavityMap(List<string> grid)
    {
        List<string> result = new List<string>();
        for (int i = 0; i < grid.Count; i++)
        {
            string row = string.Empty;
            for (int j = 0; j < grid.Count; j++)
            {
                if (i.Equals(0) || i.Equals(grid.Count - 1) ||
                    j.Equals(0) || j.Equals(grid.Count - 1))
                {
                    row += grid[i][j];
                } else if (grid[i][j] > grid[i - 1][j] &&
                 grid[i][j] > grid[i + 1][j] &&
                 grid[i][j] > grid[i][j - 1] &&
                 grid[i][j] > grid[i][j + 1])
                {
                    row += 'X';
                }
                else row += grid[i][j];
            }
            result.Add(row);
        }
        return result;
    }

}

class Solution
{
    public static void Main(string[] args)
    {
        // TextWriter textWriter = new StreamWriter(@System.Environment.GetEnvironmentVariable("OUTPUT_PATH"), true);

        //int n = Convert.ToInt32(Console.ReadLine().Trim());

        List<string> grid = new List<string>(){
            "989",
            "191",
            "111"
        };

        // for (int i = 0; i < n; i++)
        // {
        //     string gridItem = Console.ReadLine();
        //     grid.Add(gridItem);
        // }

        List<string> result = Result.cavityMap(grid);

        // textWriter.WriteLine(String.Join("\n", result));

        // textWriter.Flush();
        // textWriter.Close();
        Console.Write(result);
    }
}
