<?php

namespace App\Http\Controllers;

use App\Http\Requests\CurrencyRequest;
use App\Http\Requests\MoedaRequest;
use App\Repositories\MoedaRepository;
use App\Service\MoedaService;
use Illuminate\Database\Eloquent\ModelNotFoundException;
use Illuminate\Http\JsonResponse;
use Illuminate\Http\Request;
use Illuminate\Validation\ValidationException;
use Symfony\Component\HttpKernel\Exception\NotFoundHttpException;

class MoedaController extends BaseController
{
    public function __construct(MoedaRepository $moedaRepository, MoedaRequest $rules)
    {
        parent::__construct($moedaRepository, $rules);
    }

    public function converteMoedas(Request $request): JsonResponse
    {
        try {
            $this->validate($request, CurrencyRequest::create());
        } catch (ValidationException $e) {
            return $this->apiResponse(false, 'Parametros incorretos.', $e->errors(), 400);
        }
        
        try {
            $moedaService = new MoedaService($request->input('to'), $request->input('from'), $request->input('amount'));

            return $this->apiResponse(true, 'Dados retornados com sucesso', $moedaService->getConversion());
        } catch (ModelNotFoundException | NotFoundHttpException $e) {
            return $this->apiResponse(false, $e->getMessage(), [], 404);
        } catch (\Throwable $e) {
            return $this->apiResponse(false, $e->getMessage(), [], 500);
        }
    }
}
