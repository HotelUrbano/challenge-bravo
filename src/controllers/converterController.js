const Currency = require('../models/currencyModel');
const converterService = require('../services/converterService');
const { validationResult } = require('express-validator');
require('dotenv').config();

exports.convert = async function (req, res, next) {

    const { errors } = validationResult(req);

    if (errors.length > 0) {
        return res.status(400).send({ 
            error: errors 
        })
    }

    // Salva moeda de referência numa variável
    const ref = process.env.CURRENCY_REF

    let { from, to, amount } = req.query;

    // Checa se moeda de origem existe no banco de dados
    const confirmFrom = await Currency
        .findOne({ 
            sigla: from 
        })
        .then(dbCurrencyFrom => {
            if (!dbCurrencyFrom) {
                res.status(404).send({
                    message: `A moeda de origem ${from} não consta no banco de dados.`
                });
            } return dbCurrencyFrom
        })
        .catch(err => {
            res.status(500).send({
                message: 
                    err.message || `Erro ao procurar moeda de origem com codigo: ${from}`
            });
        });

    // Checa se moeda de destino existe no banco de dados
    const confirmTo = await Currency
        .findOne({ 
            sigla: to 
        })
        .then(dbCurrencyTo => {
            if (!dbCurrencyTo) {
                res.status(404).send({
                    message: `A moeda de destino ${to} não consta no banco de dados.` 
                });
            } return dbCurrencyTo
        })
        .catch(err => {
            res.status(500).send({
                message: 
                    err.message || `Erro ao procurar moeda de destino com codigo: ${to}`
            });
        });
        
    if ((confirmFrom != null) && (confirmTo != null)) { 
        await converterService.convert(from, to, amount, ref)
            .then(converted => {
                res.status(200).send({
                    message: `Conversão de ${amount} em ${from} para ${to}`,
                    [from]: parseFloat(amount), // Moeda de origem
                    [to]: converted // Moeda de destino
                })
            })
            .catch(err => {
                res.status(500).send({ 
                    message: 
                        err.message || 'Erro ao realizar a conversão' 
                })
            });
    }
}
