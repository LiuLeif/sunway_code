#!/bin/bash
export TENSORFLOW_ROOT_DIR=~/source/tensorflow
pushd $TENSORFLOW_ROOT_DIR/tensorflow/lite/tools/make
./build_lib.sh && popd && g++  -O2 -I${TENSORFLOW_ROOT_DIR}/tensorflow/lite/tools/make/downloads/flatbuffers/include -I${TENSORFLOW_ROOT_DIR} -pthread -Wall -Wextra -pedantic -o inference inference.cc ./temp/output.cc  -L${TENSORFLOW_ROOT_DIR}/tensorflow/lite/tools/make/gen/linux_x86_64/lib/ -ltensorflow-lite -ldl
