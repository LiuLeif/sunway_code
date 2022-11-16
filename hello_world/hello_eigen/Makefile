CPPFLAGS := -Ieigen-3.4.0
CXXFLAGS := -g -Og

SRC := $(wildcard *.cc)
APP := $(patsubst %.cc,%,${SRC})

all:${APP}

clean:
	rm -rf ${APP}
