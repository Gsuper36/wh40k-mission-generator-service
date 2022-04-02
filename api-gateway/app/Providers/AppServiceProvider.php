<?php

namespace App\Providers;

use Illuminate\Support\ServiceProvider;
use Pb\MissionGeneratorClient;
use Grpc\ChannelCredentials;

class AppServiceProvider extends ServiceProvider
{
    /**
     * Register any application services.
     *
     * @return void
     */
    public function register()
    {
        $this->app->bind(MissionGeneratorClient::class, function ($app) {
            return new MissionGeneratorClient(config('grpc.mission_generator.address'), [
                "credentials" => ChannelCredentials::createInsecure()
            ]);
        });
    }

    /**
     * Bootstrap any application services.
     *
     * @return void
     */
    public function boot()
    {
        //
    }
}
