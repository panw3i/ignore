# 定义变量
SOURCES := $(wildcard *.go)
BINARY_NAME := ignore
PLATFORMS := linux darwin windows
ARCHS := amd64 arm64

# 编译所有平台和架构的二进制文件
all: clean
	$(foreach platform,$(PLATFORMS),$(foreach arch,$(ARCHS), \
		GOOS=$(platform) GOARCH=$(arch) go build -o out/$(BINARY_NAME)-$(platform)-$(arch)$(if $(filter $(platform),windows),.exe) $(SOURCES);))

# 清理所有二进制文件
clean:
	$(foreach platform,$(PLATFORMS),$(foreach arch,$(ARCHS), \
		rm -f out/$(BINARY_NAME)-$(platform)-$(arch)$(if $(filter $(platform),windows),.exe);))

# 编译指定平台和架构的二进制文件
build:
	GOOS=$(os) GOARCH=$(arch) go build -o out/$(BINARY_NAME)-$(os)-$(arch)$(if $(filter $(os),windows),.exe) $(SOURCES)

.PHONY: all clean build
