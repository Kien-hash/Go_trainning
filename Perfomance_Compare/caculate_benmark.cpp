#include <stdio.h>
#include <time.h>
#include <string>
#include <iostream>
using namespace std;

#define N1 500000
#define N2 1000000
#define N3 1000000000
#define N4 10000000000
#define N5 100000000000

string charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789";
int charset_Size = charset.size();

double sum_benmark(double N);
double cal_benmark(double N);
double Fibonaci(int N);
string random_string(int N);
int strCompare(string str1, string str2);

int main()
{
   clock_t begin = clock();

   {
      // do smt in here to consume time
   }

   clock_t end = clock();
   double time_spent = (double)(end - begin) / CLOCKS_PER_SEC;
   printf("\nTime consumming: %f", time_spent);
   return 0;
}

/********************************************************************
   Ham tinh tong tu 1 den N 
*********************************************************************/
double sum_benmark(double N)
{
   double sum = 0;
   double count = 1;
   do
   {
      sum += count;
   } while (count++ < N);
   return sum;
}

/********************************************************************
   Ham tinh tong cua chuoi phan so  
*********************************************************************/
double cal_benmark(double N)
{
   double sum = 0;
   double count = 1;
   double num;
   int sign = -1;
   do
   {
      //printf("%d\n",sign);
      num = count / ((count + 1) * (count + 2));
      sum += sign * num;
      sign = -sign;
      //printf("\n%f",sum);
   } while (count++ < N);
   return sum;
}

/********************************************************************
   Ham tinh so thu N trong day Fibonaci 
*********************************************************************/
double Fibonaci(int N)
{
   if (N == 1 || N == 2)
   {
      return 1;
   }
   else
   {
      return Fibonaci(N - 1) + Fibonaci(N - 2);
   }
}

/********************************************************************
   Ham sinh ra chuoi co N ki tu ngau nhien
*********************************************************************/
string random_string(int N)
{
   string randString;
   for (int i = 0; i < N; i++)
   {
      randString += charset[rand() % charset_Size];
   }
   return randString;
}

/********************************************************************
   Ham so sanh 2 chuoi co cung do dai la N
*********************************************************************/
int strCompare(string str1, string str2)
{
   int count = 0;
   for (int i = 0; i < str1.size(); i++)
   {
      if (str1[i] == str2[i])
         count++;
   }
   return count;
}