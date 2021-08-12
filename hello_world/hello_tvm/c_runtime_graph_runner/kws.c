#include <float.h>
#include <stdio.h>
#include <stdlib.h>

#include "tvm_runtime.h"

int main(int argc, char** argv) {
    TVMGraphExecutor* executor = tvm_runtime_create();
    float input_data[1 * 99 * 12];

    DLTensor input;
    input.data = input_data;
    DLDevice dev = {kDLCPU, 0};
    input.device = dev;
    input.ndim = 3;
    DLDataType dtype = {kDLFloat, 32, 1};
    input.dtype = dtype;
    input.shape = (int64_t[]){1, 99, 12};
    input.strides = NULL;
    input.byte_offset = 0;

    TVMGraphExecutor_SetInput(executor, "input_1", &input);
    TVMGraphExecutor_Run(executor);

    float output_data[1];
    DLTensor output;
    output.data = output_data;
    DLDevice out_dev = {kDLCPU, 0};
    output.device = out_dev;
    output.ndim = 2;
    DLDataType out_dtype = {kDLFloat, 32, 1};
    output.dtype = out_dtype;
    output.shape = (int64_t[]){1, 1};
    output.strides = NULL;
    output.byte_offset = 0;

    TVMGraphExecutor_GetOutput(executor, 0, &output);

    printf("output: %f\n", output_data[0]);
}
