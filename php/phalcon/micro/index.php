<?php

use Phalcon\Mvc\Micro;
use Phalcon\Loader;
use Phalcon\Di\FactoryDefault;
use Phalcon\Db\Adapter\Pdo\Mysql as PdoMysql;
use Phalcon\Http\Response;

/**
 * 自动加载model
 **/
$loader = new Loader();
$loader->registerNamespaces(
    [
        'Store\Toys' => __DIR__.'/models',
    ]
);
$loader->register();

/**
 * 数据库
 */
$di = new FactoryDefault();
$di->set(
    'db',
    function ()  {
        return new PdoMysql(
            [
                'host' => 'localhost',
                'username' => 'root',
                'password' => '123456',
                'dbname' => 'robotics',
            ]
        );
    }
);

$app = new Micro();

/**
 * 定义路由
 */
// 获取所有robots
$app->get(
    '/api/robots',
    function() use($app) {
        $phql = 'SELECT *  FROM Store\Toys\Robots ORDER BY name';
        $robots = $app->modelsManager->executeQuery($phql);
        $data = [];
        foreach ($robots as $robot) {
            $data[] = [
                'id' => $robot->id,
                'name' => $robot->name,
            ];
        }
        echo json_encode($data);
    }
);
// 按名称搜索robot
$app->get(
    '/api/robots/search/{name}',
    function($name) use($app) {
        $phql = 'SELECT * FROM Store\Toys\Robots WHERE name LIKE :name ORDER BY name';
        $robots = $app->modelsManager->executeQuery(
            $phql,
            [
                'name' => '%'.$name.'%',
            ]
        );
        $data = [];
        foreach ($robots as $robot) {
            $data[] = [
                'id' => $robot->id,
                'name' => $robot->name,
            ];
        }
        echo json_encode($data);
    }
);
// 获取指定Id的robot
$app->get(
    '/api/robots/{id:[0-9]+}',
    function($id) use ($app) {
        $phql = 'SELECT * FROM Store\Toys\Robots WHERE id = :id';
        $robot = $app->modelsManager->executeQuery(
            $phql,
            [
                'id' => $id,
            ]
        )->getFirst();
        $response = new Response();
        if ($robot === false) {
            $response->setJsonContent(
                [
                    'status' => 'NOT-FOUND',
                ]
            );
        } else {
            $response->setJsonContent(
                [
                    'status' => 'FOUND',
                    'data' => [
                        'id' => $robot->id,
                        'name' => $robot->name,
                    ]
                ]
            );
        }
        return $response;

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