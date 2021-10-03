const chartProperties = {
    width: 1500, 
    height: 600,
    timeScale: {
        timeVisible: true,
        secondsVisible: false,
    },
    crosshair: {
		mode: LightweightCharts.CrosshairMode.Normal,
	},
	timeScale: {
		borderColor: 'rgba(197, 203, 206, 1)',
	},
	handleScroll: {
		vertTouchDrag: false,
	}
}

const domElement = document.getElementById("tvchart");
const chart = LightweightCharts.createChart(domElement, chartProperties);
candleSeries = chart.addCandlestickSeries();

var ema56Line;
var ema200Line;

LoadChart(0);

function LoadChart(bt){
    var resp;
    if(bt == 0){
        resp = fetch("http://localhost:8090/load/btcusd");
    }
    else {
        resp = fetch("http://localhost:8090/get");

        /*for(var i = 0; i < resistanceLines.length; i++){
            candleSeries.removePriceLine(resistanceLines[i]);
            candleSeries.removePriceLine(supportLines[i]);
        }

        for(var i = 0; i < averages.length; i++){
            candleSeries.removeLineSeries(averages[i]);
        }*/

        chart.removeSeries(candleSeries);
        chart.removeSeries(ema56Line);
        chart.removeSeries(ema200Line);

        candleSeries = chart.addCandlestickSeries();
    }

    resp
    .then(res => res.json())
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
        LoadEmas(cdata);
        LoadVolume(cdata);
        DrawSupport(30);
        DrawSupport(120);
        DrawResistance(30);
        DrawResistance(120);
        /*var minPriceLine = {
            price: 7000,
            color: 'red',
            lineWidth: 1,
            lineStyle: LightweightCharts.LineStyle.Solid,
            axisLabelVisible: true,
            // title: 'minimum price',
        };
        

        candleSeries.createPriceLine(minPriceLine);

        var mainSeries = chart.addLineSeries();
        candleSeries.setData(generateData());
        */

    })
    .catch(err => console.log(err));
}



function LoadVolume(chData){
    var volumeSeries = chart.addHistogramSeries({
        color: '#26a69a',
        priceFormat: {
            type: 'volume',
        },
        priceScaleId: '',
        scaleMargins: {
            top: 0.8,
            bottom: 0,
        },
    });

    var result = [];
    for(var i = 0; i < chData.length; i++)
    {
        result.push(
            {
                time: chData[i].time, 
                value: chData[i].volume,
                color: (chData[i].close > chData[i].open ? "#00cc66" : "#ff5050" )
        });
    }

    volumeSeries.setData(result);

}

function LoadEmas(chData){
    fetch("http://localhost:8090/emas")
    .then(res => res.json())
    .then(data => {

        var result56 = [];
        var result200 = [];
        for(var i = 0; i < data.ema56.length; i++)
        {
            result56.push({
                time: chData[i].time,
                value: data.ema56[i]
            });

            result200.push({
                time: chData[i].time,
                value: data.ema200[i]
            });
        }

        ema56Line = chart.addLineSeries({
            color: 'blue',
            lineWidth: 2
        });

        ema200Line = chart.addLineSeries({
            color: 'orange',
            lineWidth: 2
        });

        ema56Line.setData(result56);
        ema200Line.setData(result200);

})
.catch(err => console.log(err));
}

function DrawSupport(value){
    fetch("http://localhost:8090/support/" + value)
    .then(res => res.json())
    .then(data => {
        // console.log(data);
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

function DrawResistance(value){
    fetch("http://localhost:8090/resistance/" + value)
    .then(res => res.json())
    .then(data => {
        // console.log(data);
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

function backtrack(){
    backtrack_days = document.getElementById('backtrack_days').value;
    if(backtrack_days != ""){
        back = parseInt(backtrack_days) + 1;
        fetch("http://localhost:8090/backtrack/" + back);
        LoadChart(back);
    }   
}