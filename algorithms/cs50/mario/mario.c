#include <cs50.h>
#include <stdio.h>


int main(void)
{
    //start a forever loop to read the entry of the user
for (int g = 0; g >= 0; g++)
{
    int input = get_int("Enter a number between 1 and 8:");

    if (input >= 1 && input <= 8)
    {
     int line = 1;
     while (line <= input)
     {
     int space = input - line;
        for (int sp = space; sp >0; sp--)
{
        printf(" ");
};
        for (int i = input; i > space; i--)
{
        printf("#");
};
     printf("\n");
     line++;
    }
    }
}
}
