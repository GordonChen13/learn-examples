<?php

// 子进程比父进程先结束，但父进程没有回收子进程，释放子进程占用的资源，子进程就会成为僵尸进程
$pid = pcntl_fork();
if ($pid > 0) {
    // 更改php进程名称
    cli_set_process_title('php father process');
    // 主进程休息60s
    sleep(60);
} elseif ($pid == 0) {
    cli_set_process_title('php child process');
    // 让子进程休息10s，但是进程结束后，父进程不对子进程做任何处理，这样子进程就会变成僵尸进程
    sleep(10);
} else {
    exit('fork error.' .PHP_EOL);
}