#
# This is free and unencumbered software released into the public domain.
#
# Anyone is free to copy, modify, publish, use, compile, sell, or
# distribute this software, either in source code form or as a compiled
# binary, for any purpose, commercial or non-commercial, and by any
# means.
#
# In jurisdictions that recognize copyright laws, the author or authors
# of this software dedicate any and all copyright interest in the
# software to the public domain. We make this dedication for the benefit
# of the public at large and to the detriment of our heirs and
# successors. We intend this dedication to be an overt act of
# relinquishment in perpetuity of all present and future rights to this
# software under copyright law.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
# EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
# MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
# IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR
# OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
# ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
# OTHER DEALINGS IN THE SOFTWARE.
#
# For more information, please refer to <http://unlicense.org/>
#

HANDLER ?= handler
PACKAGE ?= $(HANDLER)
GOPATH  ?= $(HOME)/go
RELEASE_DIR ?= release

docker:
	@docker run --rm                                                             \
	  -e HANDLER=$(HANDLER)                                                      \
	  -e PACKAGE=$(PACKAGE)                                                      \
	  -e GOPATH=$(GOPATH)                                                        \
	  -v $(CURDIR):$(CURDIR)                                                     \
	  $(foreach GP,$(subst :, ,$(GOPATH)),-v $(GP):$(GP))                        \
	  -w $(CURDIR)                                                               \
	  eawsy/aws-lambda-go-shim:latest make all

.PHONY: docker

all: build pack copy perm

.PHONY: all

build:
	@go build -buildmode=plugin -ldflags='-w -s' -o $(HANDLER).so

.PHONY: build

pack:
	@pack $(HANDLER) $(HANDLER).so $(PACKAGE).zip

.PHONY: pack

copy:
	  @mkdir -p $(RELEASE_DIR)
		@cp $(PACKAGE).zip $(RELEASE_DIR)
		
.PHONY: copy

perm:
	@chown $(shell stat -c '%u:%g' .) $(HANDLER).so $(PACKAGE).zip
	@chown -R $(shell stat -c '%u:%g' .) $(RELEASE_DIR)

.PHONY: perm

clean:
	@rm -rf $(HANDLER).so $(PACKAGE).zip
	@rm -rf $(RELEASE)

.PHONY: clean
