CPPFLAGS := -Ieigen-3.4.0 -I.
CXXFLAGS := -g -O0

SRC := $(wildcard *.cc)
APP := $(patsubst %.cc,%,${SRC})

all:${APP}

clean:
	rm -rf ${APP}
