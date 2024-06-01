#!/bin/bash  
#表示使用的bash 系统
# 直接./hello.sh   就会使用上面的解释器
# 没有则会报错  

echo "hello"

# 执行前要给脚本+ x 执行权限
chmod a+x hello.sh # a表示所有用户 可以看 tldr

date

whoami


#!/bin/bash

# 这是一个示例Shell脚本

# 输出欢迎消息
echo "欢迎使用示例Shell脚本！"

# 获取用户输入
echo "请输入您的名字："
read name

# 打印用户输入
echo "您输入的名字是：$name"

# 检查目录是否存在
echo "请输入一个目录路径："
read directory

if [ -d "$directory" ]; then
  echo "目录存在！"
else
  echo "目录不存在！"
fi

# 列出目录中的文件
echo "目录中的文件列表："
ls -l "$directory"

# 完成消息
echo "示例Shell脚本已完成。谢谢使用！"