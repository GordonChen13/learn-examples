<?php
// pcntl_wait有阻塞问题，父进程只能挂起等待子进程结束或终止。
// pcntl_waitpid( $pid, &$status, $option = 0 )的第三个参数如果设置为WNOHANG，
//那么父进程不会阻塞一直等待到有子进程退出或终止，否则将会和pcntl_wait()的表现类似
$pid = pcntl_fork();
if ($pid > 0) {
    // 更改php进程名称
    cli_set_process_title('php father process');

    // 返回$wait_result,就是子进程的进程号，如果子进程已经是僵尸进程则为0
//    $wait_result = pcntl_wait($status);
    $wait_result = pcntl_waitpid($pid, $status, WNOHANG);
    print_r($wait_result);
    print_r($status);
    sleep(60);
    echo 'no blocking, running here'.PHP_EOL;
} elseif ($pid == 0) {
    cli_set_process_title('php child process');
    // 让子进程休息10s，但是进程结束后，父进程不对子进程做任何处理，这样子进程就会变成僵尸进程
    sleep(10);
} else {
    exit('fork error.' .PHP_EOL);
}
