<?php

use Phalcon\Mvc\Micro;

$app = new Micro();

/**
 * 定义路由
 */
// 获取所有robots
$app->get(
    '/api/robots',
    function() {

    }
);
// 按名称搜索robot
$app->get(
    '/api/robots/search/{name}',
    function($name) {

    }
);
// 获取指定Id的robot
$app->get(
    '/api/robots/{id:[0-9]+}',
    function($id) {

    }
);
// 添加新的robot
$app->post(
    '/api/robots',
    function() {

    }
);
// 更新指定id的robot
$app->put(
    '/api/robots/{id:[0-9]+}',
    function($id) {

    }
);
// 删除指定id的robot
$app->delete(
    '/api/robots/{id:[0-9]+}',
    function($id) {

    }
);
$app->handle();