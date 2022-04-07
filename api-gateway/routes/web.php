<?php

use Illuminate\Support\Facades\Route;
use Pb\GenerateMissionRequest;
use Pb\MissionGeneratorClient;

/*
|--------------------------------------------------------------------------
| Web Routes
|--------------------------------------------------------------------------
|
| Here is where you can register web routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| contains the "web" middleware group. Now create something great!
|
*/

Route::get('/', function (MissionGeneratorClient $client) {
    return "DO NOT JUDGE MY CODE, I'M HAVING FUN!!!";
});
