<?php

namespace App\Http\Controllers;

use App\Http\Requests\GenerateMissionRequest;
use Exception;
use Pb\GenerateMissionRequest as PbGenerateMissionRequest;
use Pb\MissionGeneratorClient;

class MissionController extends Controller
{
    public function random(GenerateMissionRequest $request, MissionGeneratorClient $generator)
    {
        //@todo rewrite, looks like a GO-code
        list($mission, $error) = $generator->Generate(
            new PbGenerateMissionRequest(
                $request->validated()
            )
        )->wait();

        if (! $mission) {
            throw new Exception("Couldn't generate mission: {$error}");
        }

        return  $mission->serializeToJsonString();
    }
}
