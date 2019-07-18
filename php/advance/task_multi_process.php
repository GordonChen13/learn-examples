<?php

// 设置umask为0，使当前进程创建的文件权限为777
umask(0);
$pid = pcntl_fork();
if ($pid < 0) {
    exit('fork error.');
} else if ($pid > 0) {
    // 主进程退出
    exit();
}
// 子进程继续执行
// 最关键的一步，执行setsid函数
/**
 * http://linux.die.net/man/2/setsid
[setsid详解][1] 主要目的脱离终端控制，自立门户。
创建一个新的会话，而且让这个pid统治这个会话，他既是会话组长，也是进程组长。
而且谁也没法控制这个会话，除了这个pid。当然关机除外。。
这时可以成做pid为这个无终端的会话组长。
注意这个函数需要当前进程不是父进程，或者说不是会话组长。
在这里当然不是，因为父进程已经被kill

换句话来说就是 : 调用进程必须是非当前进程组组长，调用后，产生一个新的会话期，且该会话期中只有一个进程组，且该进程组组长为调用进程，没有控制终端，新产生的group ID 和 session ID 被设置成调用进程的PID
 */
if (!posix_setsid()) {
    exit('setsid error.');
}
// 理论上fork一次就可以了
// 但fork两次，在基于system V的系统中，通过再次fork，父进程退出，子进程继续
// 保证形成的daemon进程绝对不会成为会话首进程，不会拥有控制终端。
$pid = pcntl_fork();
if ($pid < 0) {
    exit('fork error');
} else if ($pid > 0) {
    // 主进程退出
    exit();
}
// 子进程继续执行
// 给进程重新起个名字
cli_set_process_title('php master process');

// 生成守护进程后，fork 5个子进程执行任务
$child_pid = [];

// 父进程安装SIGCHLD信号处理分发器
pcntl_signal(SIGCHLD, function (){
    // 注意需使用global将child_pid全局化，不然读到的数组为空
    global $child_pid;
    // 如果子进程的数量大于0，也就是说如果还有子进程存活未退出，那么执行回收
    $child_pid_count = count($child_pid);
    if ($child_pid_count > 0) {
        foreach ($child_pid as $pid_key => $pid_item) {
            $wait_result = pcntl_waitpid($pid_item, $status, WNOHANG);
            // 子进程回收成功， 将其进程id从child_pid中移除
            if ($wait_result == $pid_item || -1 == $wait_result) {
                unset($child_pid[$pid_key]);
            }
        }
    }
});

for ($i = 1; $i <= 5; $i++) {
    $_pid = pcntl_fork();
    if ($_pid < 0) {
        exit();
    } elseif ($_pid == 0) {
        cli_set_process_title('php worker process');
        /**
         *  业务代码。。。。。。
         */
        exit();
    } elseif ($_pid > 0) {
        $child_pid[] = $_pid;
    }
}

// 主进程继续循环不断派遣信号
while (true) {
    pcntl_signal_dispatch();
    // 每派遣一次休眠一秒
    sleep(1);
}