<?php

/** @var \Laravel\Lumen\Routing\Router $router */

$router->group(['prefix' => 'api'], function () use($router) {
    $router->get('/', function () use ($router) {
        return $router->app->version();
    });

    $router->group(['prefix' => 'moedas'], function () use($router) {
        $router->get('', 'MoedaController@index');
        $router->get('{uid}', 'MoedaController@show');
        $router->post('', 'MoedaController@store');
        $router->put('{uid}', 'MoedaController@update');
        $router->delete('{uid}', 'MoedaController@destroy');
    });
});
