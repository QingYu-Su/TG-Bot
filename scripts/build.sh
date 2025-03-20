#!/bin/sh
# 声明这是一个bash脚本，使用/bin/sh作为解释器

# Exit when any command fails
set -e
# 设置脚本在执行过程中，如果任何一个命令返回非零状态（即失败），脚本会立即退出

# Get the script directory and change to the project root
cd "$(dirname "$0")/../"
# 获取当前脚本所在的目录路径，并切换到项目根目录
# $0 是脚本本身的路径，dirname 获取其目录部分，再通过 ../ 跳转到项目根目录

# Detect the operating system
OS=$(uname -s)
# 使用 uname -s 命令获取当前操作系统名称，并将其赋值给变量 OS
# uname -s 会返回类似 "Linux"、"Darwin"（macOS）或 "CYGWIN_NT-10.0"（Windows）等值

# Set output file name based on the OS
if [[ "$OS" == *"CYGWIN"* || "$OS" == *"MINGW"* || "$OS" == *"MSYS"* ]]; then
  OUTPUT="./bin/bot.exe"
# 如果操作系统是 CYGWIN、MINGW 或 MSYS（均为 Windows 环境下的兼容层），则将输出文件名设置为 memogram.exe
else
  OUTPUT="./bin/bot"
# 对于其他操作系统（如 Linux 或 macOS），输出文件名设置为 memogram
fi

echo "Building for $OS..."
# 输出当前正在构建的目标操作系统信息

# Build the executable
go build -o "$OUTPUT" ./main.go
# 使用 Go 语言的构建工具 go build 来编译项目
# -o 参数指定输出的可执行文件路径为 $OUTPUT
# ./bin/memogram/main.go 是需要编译的主程序文件路径

# Output the success message
echo "Build successful!"
# 构建成功后，输出提示信息

# Output the command to run
echo "To run the application, execute the following command:"
echo "$OUTPUT"
# 提示用户运行生成的可执行文件的命令
# 输出 $OUTPUT 变量的值，即生成的可执行文件路径