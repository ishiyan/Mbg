#include <stdio.h>
#include <windows.h>

// This is a console application ...

#define LEN 252
#define L1 6
#define L2 42

double input1C[LEN] = {
91.5000,
94.8150,
94.3750,
95.0950,
93.7800,
94.6250,
92.5300,
92.7500,
90.3150,
92.4700,
96.1250,
97.2500,
98.5000,
89.8750,
91.0000,
92.8150,
89.1550,
89.3450,
91.6250,
89.8750,
88.3750,
87.6250,
84.7800,
83.0000,
83.5000,
81.3750,
84.4400,
89.2500,
86.3750,
86.2500,
85.2500,
87.1250,
85.8150,
88.9700,
88.4700,
86.8750,
86.8150,
84.8750,
84.1900,
83.8750,
83.3750,
85.5000,
89.1900,
89.4400,
91.0950,
90.7500,
91.4400,
89.0000,
91.0000,
90.5000,
89.0300,
88.8150,
84.2800,
83.5000,
82.6900,
84.7500,
85.6550,
86.1900,
88.9400,
89.2800,
88.6250,
88.5000,
91.9700,
91.5000,
93.2500,
93.5000,
93.1550,
91.7200,
90.0000,
89.6900,
88.8750,
85.1900,
83.3750,
84.8750,
85.9400,
97.2500,
99.8750,
104.9400,
106.0000,
102.5000,
102.4050,
104.5950,
106.1250,
106.0000,
106.0650,
104.6250,
108.6250,
109.3150,
110.5000,
112.7500,
123.0000,
119.6250,
118.7500,
119.2500,
117.9400,
116.4400,
115.1900,
111.8750,
110.5950,
118.1250,
116.0000,
116.0000,
112.0000,
113.7500,
112.9400,
116.0000,
120.5000,
116.6200,
117.0000,
115.2500,
114.3100,
115.5000,
115.8700,
120.6900,
120.1900,
120.7500,
124.7500,
123.3700,
122.9400,
122.5600,
123.1200,
122.5600,
124.6200,
129.2500,
131.0000,
132.2500,
131.0000,
132.8100,
134.0000,
137.3800,
137.8100,
137.8800,
137.2500,
136.3100,
136.2500,
134.6300,
128.2500,
129.0000,
123.8700,
124.8100,
123.0000,
126.2500,
128.3800,
125.3700,
125.6900,
122.2500,
119.3700,
118.5000,
123.1900,
123.5000,
122.1900,
119.3100,
123.3100,
121.1200,
123.3700,
127.3700,
128.5000,
123.8700,
122.9400,
121.7500,
124.4400,
122.0000,
122.3700,
122.9400,
124.0000,
123.1900,
124.5600,
127.2500,
125.8700,
128.8600,
132.0000,
130.7500,
134.7500,
135.0000,
132.3800,
133.3100,
131.9400,
130.0000,
125.3700,
130.1300,
127.1200,
125.1900,
122.0000,
125.0000,
123.0000,
123.5000,
120.0600,
121.0000,
117.7500,
119.8700,
122.0000,
119.1900,
116.3700,
113.5000,
114.2500,
110.0000,
105.0600,
107.0000,
107.8700,
107.0000,
107.1200,
107.0000,
91.0000,
93.9400,
93.8700,
95.5000,
93.0000,
94.9400,
98.2500,
96.7500,
94.8100,
94.3700,
91.5600,
90.2500,
93.9400,
93.6200,
97.0000,
95.0000,
95.8700,
94.0600,
94.6200,
93.7500,
98.0000,
103.9400,
107.8700,
106.0600,
104.5000,
105.0000,
104.1900,
103.0600,
103.4200,
105.2700,
111.8700,
116.0000,
116.6200,
118.2800,
113.3700,
109.0000,
109.7000,
109.2500,
107.0000,
109.1900,
110.0000,
109.2000,
110.1200,
108.0000,
108.6200,
109.7500,
109.8100,
109.0000,
108.7500,
107.8700
};

double output1C[LEN];

// define function prototype
typedef int (__stdcall * LPFRSX)(unsigned int, double *, double *, double) ;
/*
EXPORT int WINAPI RSX(DWORD dwDataLength, double *pdSeries, double *pdRSX,
								double dLength) ;
*/
									
#define RSXSUCCESS 0

void print_output(void) {
	int i0 = 0;
	for  (int i=0; i<L2; i++)
	{
		for (int j=0; j <L1; j++) {
			printf(" %.15f,", output1C[i0+j]); 
		}
		i0 += L1;
		printf("\n");
	}
}

void main(void)
{
	HINSTANCE hDLL;   // Handle to DLL
	LPFRSX    RSX;    // Function pointer
	int iRes;

	// get handle to DLL module
	hDLL = LoadLibraryA("JRS_32.DLL");
	if (hDLL != NULL) {
		printf("// JRS_32.DLL loaded\n");

		// get pointer to JMA function
		RSX = (LPFRSX)GetProcAddress(hDLL, "RSX");
		if (!RSX) {
			FreeLibrary(hDLL);
			printf("// RSX function not found\n");
		}
		else {
			printf("// RSX function found\n");

			// call RSX function

			iRes = (*RSX)(LEN, input1C, output1C, 2);
			if (iRes == RSXSUCCESS) {
				printf("// RSX Successful.\n");
				printf("/////////////////////////////\n");
				printf("// length=2\n");
				printf("/////////////////////////////\n");
				print_output();
			}
			else
				printf("// RSX error %d\n", iRes);

			iRes = (*RSX)(LEN, input1C, output1C, 3);
			if (iRes == RSXSUCCESS) {
				printf("// RSX Successful.\n");
				printf("/////////////////////////////\n");
				printf("// length=3\n");
				printf("/////////////////////////////\n");
				print_output();
			}
			else
				printf("// RSX error %d\n", iRes);

			iRes = (*RSX)(LEN, input1C, output1C, 4);
			if (iRes == RSXSUCCESS) {
				printf("// RSX Successful.\n");
				printf("/////////////////////////////\n");
				printf("// length=4\n");
				printf("/////////////////////////////\n");
				print_output();
			}
			else
				printf("// RSX error %d\n", iRes);

			iRes = (*RSX)(LEN, input1C, output1C, 5);
			if (iRes == RSXSUCCESS) {
				printf("// RSX Successful.\n");
				printf("/////////////////////////////\n");
				printf("// length=5\n");
				printf("/////////////////////////////\n");
				print_output();
			}
			else
				printf("// RSX error %d\n", iRes);

			iRes = (*RSX)(LEN, input1C, output1C, 6);
			if (iRes == RSXSUCCESS) {
				printf("// RSX Successful.\n");
				printf("/////////////////////////////\n");
				printf("// length=6\n");
				printf("/////////////////////////////\n");
				print_output();
			}
			else
				printf("// RSX error %d\n", iRes);

			iRes = (*RSX)(LEN, input1C, output1C, 7);
			if (iRes == RSXSUCCESS) {
				printf("// RSX Successful.\n");
				printf("/////////////////////////////\n");
				printf("// length=7\n");
				printf("/////////////////////////////\n");
				print_output();
			}
			else
				printf("// RSX error %d\n", iRes);

			iRes = (*RSX)(LEN, input1C, output1C, 8);
			if (iRes == RSXSUCCESS) {
				printf("// RSX Successful.\n");
				printf("/////////////////////////////\n");
				printf("// length=8\n");
				printf("/////////////////////////////\n");
				print_output();
			}
			else
				printf("// RSX error %d\n", iRes);

			iRes = (*RSX)(LEN, input1C, output1C, 9);
			if (iRes == RSXSUCCESS) {
				printf("// RSX Successful.\n");
				printf("/////////////////////////////\n");
				printf("// length=9\n");
				printf("/////////////////////////////\n");
				print_output();
			}
			else
				printf("// RSX error %d\n", iRes);

			iRes = (*RSX)(LEN, input1C, output1C, 10);
			if (iRes == RSXSUCCESS) {
				printf("// RSX Successful.\n");
				printf("/////////////////////////////\n");
				printf("// length=10\n");
				printf("/////////////////////////////\n");
				print_output();
			}
			else
				printf("// RSX error %d\n", iRes);

			iRes = (*RSX)(LEN, input1C, output1C, 11);
			if (iRes == RSXSUCCESS) {
				printf("// RSX Successful.\n");
				printf("/////////////////////////////\n");
				printf("// length=11\n");
				printf("/////////////////////////////\n");
				print_output();
			}
			else
				printf("// RSX error %d\n", iRes);

			iRes = (*RSX)(LEN, input1C, output1C, 12);
			if (iRes == RSXSUCCESS) {
				printf("// RSX Successful.\n");
				printf("/////////////////////////////\n");
				printf("// length=12\n");
				printf("/////////////////////////////\n");
				print_output();
			}
			else
				printf("// RSX error %d\n", iRes);

			iRes = (*RSX)(LEN, input1C, output1C, 13);
			if (iRes == RSXSUCCESS) {
				printf("// RSX Successful.\n");
				printf("/////////////////////////////\n");
				printf("// length=13\n");
				printf("/////////////////////////////\n");
				print_output();
			}
			else
				printf("// RSX error %d\n", iRes);

			iRes = (*RSX)(LEN, input1C, output1C, 14);
			if (iRes == RSXSUCCESS) {
				printf("// RSX Successful.\n");
				printf("/////////////////////////////\n");
				printf("// length=14\n");
				printf("/////////////////////////////\n");
				print_output();
			}
			else
				printf("// RSX error %d\n", iRes);

			iRes = (*RSX)(LEN, input1C, output1C, 15);
			if (iRes == RSXSUCCESS) {
				printf("// RSX Successful.\n");
				printf("/////////////////////////////\n");
				printf("// length=15\n");
				printf("/////////////////////////////\n");
				print_output();
			}
			else
				printf("// RSX error %d\n", iRes);

			FreeLibrary(hDLL);
			printf("// Completed \n");
		}
	}
	else
		printf("// JRS_32.DLL not loaded \n");

	return;
}
