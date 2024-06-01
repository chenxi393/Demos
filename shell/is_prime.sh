#!/bin/bash

# 判断是否为素数
is_prime() {
  local num=$1
  is_prime=true

  if [ $num -lt 2 ]; then
    is_prime=false
  fi

  # ((...)) 和 [[...]] 是 Bash 构造，而不是所有 sh 环境
  for ((i=2; i*i<=num; i++)); do
    if ((num % i == 0)); then  # 纠正此处 
      is_prime=false
      break
    fi
  done
}

# 获取用户输入
echo "请输入一个正整数："
read number

# 调用函数判断是否为素数
is_prime $number

# 输出结果
if [ "$is_prime" = true ]; then
  echo "$number 是素数。"
else
  echo "$number 不是素数。"
fi

