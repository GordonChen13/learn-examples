<?php

// 子进程拥有父进程的数据副本，而不是共享
$number = 1;
$pid = pcntl_fork();
if ($pid > 0) {
    $number += 1;
    echo "father's number+1: $number".PHP_EOL;
} else if ($pid == 0) {
    $number += 2; //写时复制
    echo "son's number+2: $number".PHP_EOL;
} else {
    echo "fork fail".PHP_EOL;
}