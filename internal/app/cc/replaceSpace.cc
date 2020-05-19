#include <string>
#include <iostream>

using namespace std;

void replaceSpace(char *str,int length) {
	int blankNumber = 0;
	int oldstringLen = 0;

	for (oldstringLen = 0; str[oldstringLen] != '\0'; oldstringLen++)
	{
		if(str[oldstringLen] == ' ')
		{
			blankNumber++;
		}
	}

	int newstringLen = oldstringLen + blankNumber*2;
	printf("%d,%d \n",newstringLen, oldstringLen);
	
	str[newstringLen] = '\0';

	int point1 = oldstringLen -1;
	int point2 = newstringLen -1;
	printf("--%s--,%d\n","", point1);
	while(point1 >= 0 && point2 > point1){
		if (str[point1] == ' ') {
			str[point2--] = '0';
			str[point2--] = '2';
			str[point2--] = '%';
		}
		else {
			str[point2--] = str[point1];
		}
		point1--;
	}
}

int main(int argc, char** argv)
{
	//char str[] = "hello world!";
	char str[] = "   ";
	int len = strlen(str);
	printf("%s\n",str);

	replaceSpace(str,len);
	
	printf("%s\n",str);
	//printf("1111%s", str);
	return 0;
}