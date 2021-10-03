import { SupportService, ResistanceService, GetWedgeService } from "../services/connections.mjs";

function DrawSupport(candleSeries, value){
    SupportService(value)
    .then(data => {
        var minPriceLine = {
            price: data,
            color: 'red',
            lineWidth: 2,
            lineStyle: LightweightCharts.LineStyle.Solid,
            axisLabelVisible: true,
            title: value + 'days support',
        };
        candleSeries.createPriceLine(minPriceLine);
    });
}

function DrawResistance(candleSeries, value){
    ResistanceService(value)
    .then(data => {
        var maxPriceLine = {
            price: data,
            color: 'green',
            lineWidth: 2,
            lineStyle: LightweightCharts.LineStyle.Solid,
            axisLabelVisible: true,
            title: value + 'days resistance',
        };
        candleSeries.createPriceLine(maxPriceLine);
    });
}

function DrawWedge(chart, cb){
    GetWedgeService()
    .then(data => {
        if(data.exists == true){
            var topLine = DrawLine(
                chart, 
                {time: data.bottom[0].time, value: data.bottom[0].value}, 
                {time: data.bottom[data.bottom.length-1].time, value: data.bottom[data.bottom.length-1].value},
                {color: "purple", lineWidth: 1}
            );

            var bottomLine = DrawLine(
                chart, 
                {time: data.top[0].time, value: data.top[0].value}, 
                {time: data.top[data.top.length-1].time, value: data.top[data.top.length-1].value},
                {color: "purple", lineWidth: 1}
            );
            
            cb(topLine, bottomLine, data.type);
        }
    });
}

function DrawInfiniteLine(series, value, opt){
    var lineType = LightweightCharts.LineStyle.Solid;
    if(opt){
        switch(opt){
            case "solid":
                lineType = LightweightCharts.LineStyle.Solid;
                break;
            case "dotted":
                lineType = LightweightCharts.LineStyle.Dotted;
                break;
            case "dashed":
                lineType = LightweightCharts.LineStyle.Dashed;
                break;
            case "largeDashed":
                lineType = LightweightCharts.LineStyle.LargeDashed;
                break;
            case "sparseDotted":
                lineType = LightweightCharts.LineStyle.SparseDotted;
                break;
            default:
                lineType = LightweightCharts.LineStyle.Solid;
        }
    }

    var maxPriceLine = {
        price: value,
        color: 'blue',
        lineWidth: 1,
        lineStyle: lineType,
        axisLabelVisible: true,
    };
    series.createPriceLine(maxPriceLine);
}

function DrawLine(chart, from, to, opts){
    var mainSeries = chart.addLineSeries(opts);

    try{
        mainSeries.setData(generateLine(from, to));
    }
    catch(e){
        console.error(e)
    }

    return mainSeries;
}

function DrawLineSeries(chart, data, opts){
    var mainSeries = chart.addLineSeries(opts);

    try{
        mainSeries.setData(data);
    }
    catch(e){
        console.error(e);
    }

    return mainSeries;
}

function RemoveLine(chart, line){
    chart.removeSeries(line);
}

function generateLine(from, to) {
    var res = [];
    
    res.push({
        time: from.time,
        value: from.value,
    });

    res.push({
        time: to.time,
        value: to.value,
    });

    return res;
  }


export { DrawSupport, DrawResistance, DrawLine, RemoveLine, DrawInfiniteLine, DrawWedge, DrawLineSeries }