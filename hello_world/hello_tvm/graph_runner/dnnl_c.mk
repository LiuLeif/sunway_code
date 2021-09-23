USE_DNNL=1
include c.mk
LDLIBS += -lstdc++ -ldnnl
c_runtime/kws: ${TVM_ROOT}/src/runtime/contrib/dnnl/dnnl.o
