<?php

// 孤儿进程是指父进程在子进程结束之前提前退出，这些子进程将由init（进程id为1）收养并完成对其各种数据状态的收集
$pid = pcntl_fork();
if ($pid > 0) {
    // 显示父进程ID,可用getmypid()，或posix_getpid()
    echo "Father PID: ".getmypid().PHP_EOL;
    // 让父进程停止2秒，期间子进程的父进程id还是这个父进程
    sleep(2);
} elseif ($pid == 0) {
    // 让子进程循环10次，每次休眠1s，然后每秒钟获取一次父进程的ID
    for ($i = 1; $i <= 10; $i++) {
        sleep(1);
        echo posix_getppid().PHP_EOL;
    }
} else {
    echo "fork error".PHP_EOL;
}