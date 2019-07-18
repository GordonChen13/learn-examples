<?php

// 子进程在退出的时候，会向父进程发送一个信号，叫做SIGCHLD，
//那么父进程一旦收到了这个信号，就可以作出相应的回收动作，
//也就是执行pcntl_waitpid()，从而解决掉僵尸进程
$pid = pcntl_fork();
if ($pid > 0) {
    pcntl_signal(SIGCHLD, function () use($pid) {
        echo "receive SIGCHLD signal".PHP_EOL;
        pcntl_waitpid($pid, $status, WNOHANG);
    });
    // 更改php进程名称
    cli_set_process_title('php father process');

//    while (true) {
//        sleep(1);
//        pcntl_signal_dispatch();
//    }
} elseif ($pid == 0) {
    cli_set_process_title('php child process');
    // 让子进程休息20s后退出
    sleep(10);
    exit();
} else {
    exit('fork error.' .PHP_EOL);
}
