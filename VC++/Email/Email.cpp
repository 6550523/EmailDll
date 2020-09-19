#include "stdafx.h"
#include <stdio.h>
#include <stdlib.h>
#include <Windows.h>

#define ERR_LEN 1024

#ifdef _WIN64
#define DLL_PATH _T("..\\..\\Release\\email-win-64.dll")
#else
#define DLL_PATH _T("..\\..\\Release\\email-win-32.dll")
#endif

int main()
{
	HMODULE module = LoadLibrary(DLL_PATH);
	if (module == NULL)
	{
		DWORD err = GetLastError();
		printf("Load email.dll failed %d\n", err);
		return -1;
	}

	//SendEmail
	typedef void(*SendEmailFunc)(char *, char*, char*, char*, char*, char*, char*, char*, char*, char*, char*, char*, char*, char*);
	SendEmailFunc SendEmail;
	SendEmail = (SendEmailFunc)GetProcAddress(module, "SendEmail");
	char output[ERR_LEN];
	memset(output, 0, ERR_LEN);
	SendEmail("foo <a@qq.com>", "", "b@qq.com|c@qq.com", "d@qq.com|e@qq.com", "Subject", "Text Body", "<h1>HTML is supported</h1>", "C:\\file1.txt||C:\\file2.txt", "smtp.exmail.qq.com:587", "", "a@qq.com", "password", "smtp.exmail.qq.com", output);
	printf(output);

	FreeLibrary(module);
	return 0;
}

