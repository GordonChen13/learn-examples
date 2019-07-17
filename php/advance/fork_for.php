<?php
// for循环
for ($i = 1; $i <= 3; $i++) {
    $pid = pcntl_fork();
    if ($pid > 0) {
        echo "father $i".PHP_EOL;
    } elseif ($pid == 0) {
        echo "son $i".PHP_EOL;
//        exit();
    } else {
        echo "fork fail $i".PHP_EOL;
    }
}

// 打印7father 7son
//
// i=1的时候父进程的pid不为0 这时候fork了一个pid=0的子进程a, 子进程数量1
//
//i=2的时候父进程fork了一个子进程b, 子进程a又fork了一个子进程c, 子进程数量1+2
//
//i=3的时候父进程fork了一个子进程d, a子进程fork了e, b子进程fork了f, c子进程fork了g子进程数量1+2+4=7
//
//至于在fork子进程退出的时候 i=1 =2 =3的时候都只有一个父进程fork一个子进程 所以只有三个儿子