// cd ${TENSORFLOW_ROOT_DIR}/tensorflow/lite/tools/make
// ./download_dependencies.sh
// ./build_lib.sh
//
// g++
// -I${TENSORFLOW_ROOT_DIR}/tensorflow/lite/tools/make/downloads/flatbuffers/include
// -I${TENSORFLOW_ROOT_DIR} -pthread -Wall -Wextra -pedantic -o square square.cc
// -L${TENSORFLOW_ROOT_DIR}/tensorflow/lite/tools/make/gen/linux_x86_64/lib/
// -ltensorflow-lite -ldl

// ./square ./temp/tflite-q

#include <cstdio>
#include <unistd.h>

#include "tensorflow/lite/interpreter.h"
#include "tensorflow/lite/kernels/register.h"
#include "tensorflow/lite/model.h"
#include "tensorflow/lite/optional_debug_tools.h"

using namespace tflite;
extern float data[490];
char* words[] = {
    "silent", "unknown", "yes", "no",  "up",   "down",
    "left",   "right",   "on",  "off", "stop", "go",
};
int main(int argc, char* argv[]) {
    const char* filename = "./temp/output.tflite";
    std::unique_ptr<tflite::FlatBufferModel> model =
        tflite::FlatBufferModel::BuildFromFile(filename);

    tflite::ops::builtin::BuiltinOpResolver resolver;
    InterpreterBuilder builder(*model, resolver);
    std::unique_ptr<Interpreter> interpreter;
    builder(&interpreter);
    interpreter->AllocateTensors();
    printf("input: %s\n", interpreter->GetInputName(0));
    printf("output: %s\n", interpreter->GetOutputName(0));

    float* input = interpreter->typed_input_tensor<float>(0);
    memcpy(input, data, sizeof(data));

    interpreter->Invoke();
    float* output = interpreter->typed_output_tensor<float>(0);
    int argmax = 0;
    float max_val = 0.0;
    for (int i = 0; i < 12; i++) {
        printf("%f\n", output[i]);
        if (max_val < output[i]) {
            max_val = output[i];
            argmax = i;
        }
    }
    printf("%s\n", words[argmax]);
    return 0;
}
