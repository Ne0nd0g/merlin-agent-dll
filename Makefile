# !!!MAKE SURE YOUR GOPATH ENVIRONMENT VARIABLE IS SET FIRST!!!

# Merlin Agent DLL
VERSION=2.1.0-dll

BUILD=$(shell git rev-parse HEAD)
DIR=bin/v${VERSION}/${BUILD}
XBUILD=-X "github.com/Ne0nd0g/merlin-agent/v2/core.Build=${BUILD}"

# Merlin Agent Variables
URL ?= https://127.0.0.1:443
XURL=-X "main.url=${URL}"
PSK ?= merlin
XPSK=-X "main.psk=${PSK}"
PROXY ?=
XPROXY =-X "main.proxy=$(PROXY)"
SLEEP ?= 30s
XSLEEP =-X "main.sleep=$(SLEEP)"
HOST ?=
XHOST =-X "main.host=$(HOST)"
PROTO ?= h2
XPROTO =-X "main.protocol=$(PROTO)"
JA3 ?=
XJA3 =-X "main.ja3=$(JA3)"
USERAGENT = Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.85 Safari/537.36
XUSERAGENT =-X "main.useragent=$(USERAGENT)"
HEADERS =
XHEADERS =-X "main.headers=$(HEADERS)"
SKEW ?= 3000
XSKEW=-X "main.skew=${SKEW}"
PAD ?= 4096
XPAD=-X "main.padding=${PAD}"
KILLDATE ?= 0
XKILLDATE=-X "main.killdate=${KILLDATE}"
RETRY ?= 7
XRETRY=-X "main.maxretry=${RETRY}"
PARROT ?=
XPARROT=-X "main.parrot=${PARROT}"
AUTH ?= opaque
XAUTH=-X "main.auth=${AUTH}"
ADDR ?= 127.0.0.1:4444
XADDR=-X "main.addr=${ADDR}"
TRANSFORMS ?= jwe,gob-base
XTRANSFORMS=-X "main.transforms=${TRANSFORMS}"
LISTENER ?=
XLISTENER=-X "main.listener=${LISTENER}"
SECURE ?= false
XSECURE =-X "main.secure=${SECURE}"

# Compile Flags
LDFLAGS=-ldflags '-s -w ${XSECURE} ${XPARROT} ${XADDR} ${XAUTH} ${XTRANSFORMS} ${XLISTENER} ${XBUILD} ${XPROTO} ${XURL} ${XHOST} ${XPSK} ${XSLEEP} ${XPROXY} $(XUSERAGENT) $(XHEADERS) ${XSKEW} ${XPAD} ${XKILLDATE} ${XRETRY} -buildid='
GCFLAGS=-gcflags=all=-trimpath=$(GOPATH)
ASMFLAGS=-asmflags=all=-trimpath=$(GOPATH)# -asmflags=-trimpath=$(GOPATH)
PASSWORD=merlin
PACKAGE=7za a -p${PASSWORD} -mhe -mx=9
F=LICENSE

# Make Directory to store executables
$(shell mkdir -p ${DIR})

# Misc
# GOGARBLE contains a list of all the packages to obfuscate
GOGARBLE=golang.org,gopkg.in,github.com,go.dedis.ch
# The Merlin server and agent MUST be built with the same seed value
# Set during build with "make linux-garble SEED=<insert seed>
SEED=d0d03a0ae4722535a0e1d5d0c8385ce42015511e68d960fadef4b4eaf5942feb

# Compile Agent - Windows x64 DLL - main() - Console
default:
	export GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ CGO_ENABLED=1; \
	go build ${LDFLAGS} ${GCFLAGS} ${ASMFLAGS} -buildmode=c-archive -o ${DIR}/main.a main.go && \
	cp merlin.c ${DIR} && \
	x86_64-w64-mingw32-gcc -shared -pthread -o ${DIR}/merlin.dll ${DIR}/merlin.c ${DIR}/main.a -lwinmm -lntdll -lws2_32 && \
	cp ${DIR}/merlin.dll .

distro: clean default package

garble:
	export GOGARBLE=${GOGARBLE}; export GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ CGO_ENABLED=1; \
	garble -tiny -literals -seed ${SEED} build ${LDFLAGS} ${GCFLAGS} ${ASMFLAGS} -buildmode=c-archive -o ${DIR}/main.a main.go; \
	cp merlin.c ${DIR}; \
	x86_64-w64-mingw32-gcc -shared -pthread -o ${DIR}/merlin.dll ${DIR}/merlin.c ${DIR}/main.a -lwinmm -lntdll -lws2_32

package:
	${PACKAGE} ${DIR}/merlin-agent-dll.7z ${DIR}/merlin.dll ${F}
	cp ${DIR}/merlin-agent-dll.7z .

clean:
	rm -rf ${DIR}*
