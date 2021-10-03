import { LoadPairService, GetPairService } from "../services/connections.mjs";
import { LoadVolume, LoadEmas } from "./series.mjs";
import { DrawSupport, DrawResistance, DrawLine, DrawWedge } from "./lines.mjs";
import { chart, candleSeries } from "../app.mjs";
import { RemoveLine } from "./lines.mjs";
import { capFirst } from "../utils/utils.mjs";
import { DrawTargets } from "../indicators/targets.mjs";

var ema56Line;
var ema200Line;
var ema500Line;
var Lines = [];
var clickData = [];
var sequenceData = [];
var tempLine;
var cont = 0;

function LoadChart(pair, chart, candleSeries, wedgeElement, bt, cb){
    var linesLength = Lines.length;
    for(var i = 0; i < linesLength; i++){
        try {
            chart.removeSeries(Lines.pop());
        }
        catch{
            
        }
    }


    var candlesData;
    if(bt == 0){
        candlesData = LoadPairService(pair);
    }
    else {
        candlesData = GetPairService();

        chart.removeSeries(candleSeries);
        chart.removeSeries(ema56Line);
        chart.removeSeries(ema200Line);
        chart.removeSeries(ema500Line);

        candleSeries = chart.addCandlestickSeries();
    }

    candlesData
    .then(data => {
        const cdata = data.map(d => {
            return {
                time: new Date(d.Period.Start).getTime()/1000,
                open: d.OpenPrice,
                high: d.MaxPrice,
                low: d.MinPrice,
                close: d.ClosePrice,
                volume: d.Volume
            };
        });

        candleSeries.setData(cdata);

        LoadEmas(chart, cdata, (ema56, ema200, ema500) => {
            ema56Line = ema56;
            ema200Line = ema200;
            ema500Line = ema500;
        });

        LoadVolume(chart, cdata);
        DrawSupport(candleSeries, 30);
        DrawSupport(candleSeries, 120);
        DrawResistance(candleSeries, 30);
        DrawResistance(candleSeries, 120);
        DrawWedge(chart, (topLine, bottomLine, type) => {
            Lines.push(topLine);
            Lines.push(bottomLine);

            wedgeElement.innerHTML += "<br/>" + capFirst(type);
            console.log(type);
        });

        let appendedData = cdata;

        var today = new Date();
        var time = new Date(Date.UTC(today.getFullYear(), today.getMonth(), today.getDay()+1, 0, 0, 0, 0));
        for (var i = 0; i < 500; ++i) {
            appendedData.push({
                time: time.getTime() / 1000,
            });
            time.setUTCDate(time.getUTCDate() + 1);
        }
        candleSeries.setData(appendedData);

        DrawTargets(chart, cdata, (targetTop, targetBottom) => {
            Lines.push(targetTop);
            Lines.push(targetBottom);
        });

        cb(chart, candleSeries);
    })
    .catch(err => console.log(err));
}


function handleClick(param) {
    if (!param.point) {
        return;
    }
    
    clickData.push({
        time: param.time,
        value: candleSeries.coordinateToPrice(param.point.y, param.point.x)
    });

    if(clickData.length == 2){
        Lines.push(DrawLine(chart, clickData[0], clickData[1]));
        clickData.pop();
        clickData.pop();
        chart.unsubscribeClick(handleClick);
    }

    /*console.log(`An user clicks at (${param.point.x}, ${param.point.y}) point, the time is ${param.time}`);
    console.log(candleSeries.coordinateToPrice(param.point.y, param.point.x))*/
}

function HandleSequenceClick(param){
    if (!param.point) {
        return;
    }
    
    sequenceData.push({
        time: param.time,
        value: candleSeries.coordinateToPrice(param.point.y, param.point.x)
    });

    if(sequenceData.length == 2){
        Lines.push(DrawLine(chart, sequenceData[0], sequenceData[1]));
        [sequenceData[0], sequenceData[1]] = [sequenceData[1], sequenceData[0]];
        sequenceData.pop();
    }
}

function DeleteLastLine(){
    RemoveLine(chart, Lines.pop());
}


export { LoadChart, handleClick, DeleteLastLine, HandleSequenceClick }