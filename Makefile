# !!!MAKE SURE YOUR GOPATH ENVIRONMENT VARIABLE IS SET FIRST!!!

# Merlin Server & Agent version number
VERSION=1.5.0

BUILD=$(shell git rev-parse HEAD)
DIR=bin/v${VERSION}/${BUILD}
XBUILD=-X main.build=${BUILD} -X github.com/Ne0nd0g/merlin/pkg/agent.build=${BUILD}

# Merlin Agent Variables
URL ?= https://127.0.0.1:443
XURL=-X main.url=${URL}
PSK ?= merlin
XPSK=-X main.psk=${PSK}
PROXY ?=
XPROXY =-X main.proxy=$(PROXY)
HOST ?=
XHOST =-X main.host=$(HOST)
PROTO ?= h2
XPROTO =-X main.protocol=$(PROTO)
JA3 ?=
XJA3 =-X main.ja3=$(JA3)
KILLDATE ?= 0
XKILLDATE =-X main.killdate=$(KILLDATE)
MAXRETRY ?= 7
XMAXRETRY =-X main.maxretry=$(MAXRETRY)
PADDING ?= 4096
XPADDING =-X main.padding=$(PADDING)
SKEW ?= 0
XSKEW =-X main.skew=$(SKEW)
USERAGENT ?= Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.85 Safari/537.36
XUSERAGENT =-X "main.useragent=$(USERAGENT)"
SLEEP ?= 30s
XSLEEP = -X main.sleep=$(SLEEP)

# Compile Flags
LDFLAGS=-ldflags '-s -w ${XBUILD} ${XPROTO} ${XURL} ${XHOST} ${XPSK} ${XPROXY} ${XKILLDATE} ${XMAXRETRY} ${XPADDING} ${XSKEW} ${XUSERAGENT} ${XSLEEP} -buildid='
GCFLAGS=-gcflags=all=-trimpath=$(GOPATH)
ASMFLAGS=-asmflags=all=-trimpath=$(GOPATH)# -asmflags=-trimpath=$(GOPATH)
PACKAGE=7za a -p${PASSWORD} -mhe -mx=9
F=README.MD LICENSE data/modules docs data/README.MD data/agents/README.MD data/db/ data/log/README.MD data/x509 data/src data/bin data/html
F2=LICENSE

# Make Directory to store executables
$(shell mkdir -p ${DIR})

# Misc
# GOGARBLE contains a list of all the packages to obfuscate
GOGARBLE=golang.org,gopkg.in,github.com
# The Merlin server and agent MUST be built with the same seed value
# Set during build with "make linux-garble SEED=<insert seed>
SEED=d0d03a0ae4722535a0e1d5d0c8385ce42015511e68d960fadef4b4eaf5942feb

# Compile Agent - Windows x64 DLL - main() - Console
default:
	export GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ CGO_ENABLED=1; \
	go build ${LDFLAGS} ${GCFLAGS} ${ASMFLAGS} -buildmode=c-archive -o ${DIR}/main.a main.go; \
	cp merlin.c ${DIR}; \
	x86_64-w64-mingw32-gcc -shared -pthread -o ${DIR}/merlin.dll ${DIR}/merlin.c ${DIR}/main.a -lwinmm -lntdll -lws2_32

garble:
	export GOGARBLE=${GOGARBLE}; export GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ CGO_ENABLED=1; \
	garble -tiny -literals -seed ${SEED} build ${LDFLAGS} ${GCFLAGS} ${ASMFLAGS} -buildmode=c-archive -o ${DIR}/main.a main.go; \
	cp merlin.c ${DIR}; \
	x86_64-w64-mingw32-gcc -shared -pthread -o ${DIR}/merlin.dll ${DIR}/merlin.c ${DIR}/main.a -lwinmm -lntdll -lws2_32

clean:
	rm -rf ${DIR}*
