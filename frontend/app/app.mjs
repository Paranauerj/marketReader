import { chartProperties, rsiProperties } from "./config.mjs";
import { LoadChart, handleClick, DeleteLastLine } from "./chart/load.mjs";
import { BacktrackService } from "./services/connections.mjs";
import { LoadRSI } from "./chart/rsi.mjs";
import { GetTrends } from "./indicators/trends.mjs";
import { ClearDoms } from "./utils/clearDom.mjs";
import { HandleSequenceClick } from "./chart/load.mjs";

var chart;
var candleSeries;
const domElement = document.getElementById("tvchart");
var LightweightCharts;

const domRSIElement = document.getElementById("rsichart");
var rsiChart;
var rsiSeries;
var pair = "btcusd";

const domShortElement = document.getElementById("short_term");
const domMidElement = document.getElementById("mid_term");
const domLongElement = document.getElementById("long_term");
const domWedgeElement = document.getElementById("wedge");

function Start(paire, light){

    ClearDoms([domElement, domRSIElement, domShortElement, domMidElement, domLongElement, domWedgeElement]);

    LightweightCharts = light;
    pair = paire;
    chart = LightweightCharts.createChart(domElement, chartProperties);
    candleSeries = chart.addCandlestickSeries();

    LoadChart(pair, chart, candleSeries, domWedgeElement, 0, (charte, series) => {
        chart = charte;
        candleSeries = series;

        rsiChart = LightweightCharts.createChart(domRSIElement, rsiProperties);
        rsiSeries = rsiChart.addLineSeries();
        
        LoadRSI(rsiChart, rsiSeries, (charte, series) => {
            rsiChart = charte;
            rsiSeries = series;
        });

        GetTrends(domShortElement, domMidElement, domLongElement);

    });
}

function Backtrack(days){
    if(days != ""){
        var days_back = parseInt(days) + 1;
        BacktrackService(days_back);
        ClearDoms([domShortElement, domMidElement, domLongElement, domWedgeElement]);

        LoadChart(pair, chart, candleSeries, domWedgeElement, days_back, (charte, series) => {
            chart = charte;
            candleSeries = series;

            LoadRSI(rsiChart, rsiSeries, (charte, series) => {
                rsiChart = charte;
                rsiSeries = series;
            });

            GetTrends(domShortElement, domMidElement, domLongElement);
        });
    }   
}

function NewLine(){
    chart.subscribeClick(handleClick);
}

function NewSequence(){
    chart.subscribeClick(HandleSequenceClick);
}

function StopSequence(){
    chart.unsubscribeClick(HandleSequenceClick);
}

export { Start, Backtrack, NewLine, DeleteLastLine, candleSeries, chart, NewSequence, StopSequence }
