<?php
// 程序从pcntl_fork()后父进程和子进程将各自继续往下执行
$pid = pcntl_fork();
if ($pid > 0) {
    echo "我是父亲".PHP_EOL;
} else if ($pid == 0) {
    echo "我是儿子".PHP_EOL;
} else {
    echo "fork失败".PHP_EOL;
}