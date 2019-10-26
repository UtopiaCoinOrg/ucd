
//#include <windows.h>
#include "x19r.h"
#include <stdint.h>
#include <stdio.h>
#include <ctype.h>
void HexStrToByte(const char* source, unsigned char* dest, int sourceLen);

int main()
{
	for (int i = 0; i < 10; i++)
	{
		unsigned char buf[33] = { 0 };
		unsigned char src[80] = {0};// { 7, 0, 0, 0, 50, 102, 115, 9, 210, 174, 170, 73, 93, 4, 169, 77, 53, 117, 121, 177, 119, 25, 244, 87, 241, 232, 140, 99, 232, 76, 109, 61, 189, 194, 158, 77, 26, 223, 229, 151, 52, 127, 173, 86, 11, 85, 118, 245, 175, 78, 105, 192, 119, 180, 230, 165, 97, 68, 239, 166, 17, 245, 119, 218, 51, 28, 126, 218, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 101, 102, 102, 32, 32, 78, 0, 0, 0, 0, 0, 0, 17, 0, 0, 0, 125, 1, 0, 0, 48, 189, 119, 91, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 };
		
		/*
		00000000ff45a3435c103bc673d4967f518e6568c1b785390c97e8d0125c9ba624cae17a00000000000000000000000000000000ffff0f1e00743ba40b00000000000000808f675b0000000000000000
4d73c94bac2f3b6a817e54676a64f5cc67c3c7c3fb42f66fa0c88eeeac7b3409
		*/
		unsigned char temp[80] = {0};// = "00000000ff45a3435c103bc673d4967f518e6568c1b785390c97e8d0125c9ba624cae17a00000000000000000000000000000000ffff0f1e00743ba40b00000000000000808f675b0000000000000000";
		unsigned char result = "e24b6c013f7511398ec5644e1a86bea745efdb392a2a061ac36c508b57d5b151";

		HexStrToByte("00000000ff45a3435c103bc673d4967f518e6568c1b785390c97e8d0125c9ba624cae17a00000000000000000000000000000000ffff0f1e00743ba40b00000000000000808f675b0000000000000000",
			temp, 160);
		//4d73c94bac2f3b6a817e54676a64f5cc67c3c7c3fb42f66fa0c88eeeac7b3409
		

		
		uint32_t input[20] = {65537, 0,0,0,0,504365055, 658067456, 1164, 730,4133696795,
			2753075618, 2994799119, 1224175243, 2551598650, 3925753936, 3045750475, 2410842596, 1570783376, 0, 163491};
	
/*		uint32_t input[20] = { 6553, 0,0,0,0,504365055, 658067456, 1164, 730,4133696795,
			2753075618, 2994799119, 1224175243, 2551598650, 3925753936, 3045750475, 2410842596, 1570783376, 0, 163491 }*/;

		x19r_hash(buf, &temp, 80);

		for (int j = 0; j < 32; j++)
		{
			printf("%02x", buf[31 - j]);
		}
		printf("\n");

		int j = 0;
	}
	getchar();
	return 0;
}


void HexStrToByte(const char* source, unsigned char* dest, int sourceLen)
{
	short i;
	unsigned char highByte, lowByte;

	for (i = 0; i < sourceLen; i += 2)
	{
		highByte = toupper(source[i]);
		lowByte = toupper(source[i + 1]);

		if (highByte > 0x39)
			highByte -= 0x37;
		else
			highByte -= 0x30;

		if (lowByte > 0x39)
			lowByte -= 0x37;
		else
			lowByte -= 0x30;

		dest[i / 2] = (highByte << 4) | lowByte;
	}
	return;
}