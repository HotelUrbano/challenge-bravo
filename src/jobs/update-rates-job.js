var schedule = require('node-schedule');
const ExchangeRatesService = require( "../services/exchange-rates-service");
const Configuration = require('../config/config');
const { Container } = require("typedi")
class UpdateRatesJob
{
	initJob(){
		this.job = schedule.scheduleJob( Configuration.CRON_JOB_STRING , function(){
			Container.get(ExchangeRatesService).getExchangeRates();
			console.log("Successfully obtained updated exchange rates.");
		});
		console.log("UpdateRatesJob initialized.");
	}
	cancelJob()
	{
		if(this.job)
		{
			this.job.cancel();
			return;
		}
		console.warn("UpdateRatesJob cancel failed! Job is not running!")
	}
}
let instance = new UpdateRatesJob();

module.exports = instance;